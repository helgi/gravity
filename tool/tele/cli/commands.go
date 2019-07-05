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

package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/gravitational/gravity/lib/constants"
	"github.com/gravitational/gravity/lib/loc"
)

// Application represents the command-line "tele" application and contains
// definitions of all its flags, arguments and subcommands
type Application struct {
	*kingpin.Application
	// Debug allows to run the command in debug mode
	Debug *bool
	// Insecure turns off TLS hostname validation
	Insecure *bool
	// StateDir is the local state directory
	StateDir *string
	// VersionCmd outputs the binary version
	VersionCmd VersionCmd
	// BuildCmd builds app installer tarball
	BuildCmd BuildCmd
	// ListCmd lists available apps and runtimes
	ListCmd ListCmd
	// PullCmd downloads app installer from Ops Center
	PullCmd PullCmd
}

// VersionCmd outputs the binary version
type VersionCmd struct {
	*kingpin.CmdClause
	// Output is output format
	Output *constants.Format
}

// BuildCmd builds app installer tarball
type BuildCmd struct {
	*kingpin.CmdClause
	// ManifestPath is the path to app manifest file
	ManifestPath *string
	// OutFile is the output tarball file
	OutFile *string
	// Overwrite overwrites existing tarball
	Overwrite *bool
	// Repository is where packages are downloaded from
	Repository *string
	// Name allows to override app name
	Name *string
	// Version allows to override app version
	Version *string
	// VendorPatters is file pattern to search for images
	VendorPatterns *[]string
	// VendorIgnorePatterns if file pattern to ignore when searching for images
	VendorIgnorePatterns *[]string
	// SetImages rewrites images to specified versions
	SetImages *loc.DockerImages
	// SetDeps rewrites app dependencies to specified versions
	SetDeps *loc.Locators
	// SkipVersionCheck suppresses version mismatch check
	SkipVersionCheck *bool
	// Parallel defines the number of tasks to execute concurrently
	Parallel *int
	// SupportLegacyUpgrade specifies whether to build an installer with support
	// for upgrading 5.0.x version
	SupportLegacyUpgrade *bool
	// Quiet allows to suppress console output
	Quiet *bool
}

type ListCmd struct {
	*kingpin.CmdClause
	// Runtimes shows available runtimes
	Runtimes *bool
	// Format is the output format
	Format *constants.Format
	// All displays all available versions
	All *bool
}

// PullCmd downloads app installer from Ops Center
type PullCmd struct {
	*kingpin.CmdClause
	// App is app name
	App *string
	// OutFile is installer tarball file name
	OutFile *string
	// Force overwrites existing tarball
	Force *bool
	// Quiet allows to suppress console output
	Quiet *bool
}
