/*
Copyright 2018 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package localenv

import (
	"context"
	"time"

	"github.com/gravitational/gravity/lib/app"
	"github.com/gravitational/gravity/lib/app/service"
	"github.com/gravitational/gravity/lib/blob/fs"
	"github.com/gravitational/gravity/lib/defaults"
	"github.com/gravitational/gravity/lib/httplib"
	"github.com/gravitational/gravity/lib/ops/opsservice"
	"github.com/gravitational/gravity/lib/pack"
	"github.com/gravitational/gravity/lib/pack/localpack"
	"github.com/gravitational/gravity/lib/storage"
	"github.com/gravitational/gravity/lib/storage/keyval"
	"github.com/gravitational/gravity/lib/users"
	"github.com/gravitational/gravity/lib/users/usersservice"
	"github.com/gravitational/teleport/lib/events"

	"github.com/gravitational/trace"
	"k8s.io/client-go/kubernetes"
)

// NewClusterEnvironment returns a new instance of ClusterEnvironment
// with all services initialized
func (r *LocalEnvironment) NewClusterEnvironment(opts ...ClusterEnvironmentOption) (*ClusterEnvironment, error) {
	client, _, err := httplib.GetClusterKubeClient(r.DNS.Addr())
	if err != nil {
		log.Errorf("Failed to create Kubernetes client: %v.",
			trace.DebugReport(err))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), defaults.AuditLogClientTimeout)
	defer cancel()
	auditLog, err := r.AuditLog(ctx)
	if err != nil {
		log.Warnf("Failed to create audit log: %v.",
			trace.DebugReport(err))
	}
	config := clusterEnvironmentConfig{
		client:   client,
		auditLog: auditLog,
	}
	for _, opt := range opts {
		opt(&config)
	}
	return newClusterEnvironment(config)
}

// ClusterEnvironment provides access to local cluster services
type ClusterEnvironment struct {
	// Backend is the cluster etcd backend
	Backend storage.Backend
	// Packages is the package service that talks to local storage
	Packages pack.PackageService
	// ClusterPackages is the package service that talks to cluster API
	ClusterPackages pack.PackageService
	// Apps is the cluster apps service
	Apps app.Applications
	// Users is the cluster identity service
	Users users.Identity
	// Operator is the local operator service
	Operator *opsservice.Operator
	// Client is the cluster Kubernetes client
	Client *kubernetes.Clientset
}

// NewClusterEnvironment initializes local cluster services and
// returns a new instance of cluster environment.
// The resulting environment will not have a kubernetes client
func NewClusterEnvironment() (*ClusterEnvironment, error) {
	return newClusterEnvironment(clusterEnvironmentConfig{})
}

// WithClient is an option to override the kubernetes client to use
// in the cluster environment
func WithClient(client *kubernetes.Clientset) ClusterEnvironmentOption {
	return func(config *clusterEnvironmentConfig) {
		config.client = client
	}
}

// WithEtcdTimeout is an option to override the etcd timeout
func WithEtcdTimeout(timeout time.Duration) ClusterEnvironmentOption {
	return func(config *clusterEnvironmentConfig) {
		config.etcdTimeout = timeout
	}
}

// ClusterEnvironmentOption describes a functional option for customizing
// a cluster environment
type ClusterEnvironmentOption func(*clusterEnvironmentConfig)

func newClusterEnvironment(config clusterEnvironmentConfig) (*ClusterEnvironment, error) {
	etcdConfig, err := keyval.LocalEtcdConfig(config.etcdTimeout)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	backend, err := keyval.NewETCD(*etcdConfig)
	if err != nil {
		return nil, trace.Wrap(err, "failed to connect to etcd")
	}

	packagesDir, err := SitePackagesDir()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	objects, err := fs.New(packagesDir)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	unpackedDir, err := SiteUnpackedDir()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	packages, err := localpack.New(localpack.Config{
		Backend:     backend,
		Objects:     objects,
		UnpackedDir: unpackedDir,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	apps, err := service.New(service.Config{
		Backend:     backend,
		Packages:    packages,
		UnpackedDir: unpackedDir,
		Client:      config.client,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	users, err := usersservice.New(usersservice.Config{
		Backend: backend,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	siteDir, err := SiteDir()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	operator, err := opsservice.NewLocalOperator(opsservice.Config{
		Backend:  backend,
		Packages: packages,
		Apps:     apps,
		Users:    users,
		StateDir: siteDir,
		Local:    true,
		Client:   config.client,
		AuditLog: config.auditLog,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	clusterPackages, err := ClusterPackages()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return &ClusterEnvironment{
		Backend:         backend,
		Packages:        packages,
		ClusterPackages: clusterPackages,
		Apps:            apps,
		Users:           users,
		Operator:        operator,
		Client:          config.client,
	}, nil
}

type clusterEnvironmentConfig struct {
	client *kubernetes.Clientset
	// etcdTimeout specifies the timeout for etcd queries.
	// Falls back to defaults.EtcdRetryInterval if unspecified
	etcdTimeout time.Duration
	// auditLog provides API to the cluster audit log
	auditLog events.IAuditLog
}
