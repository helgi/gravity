// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAction) DeepCopyInto(out *AWSAction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAction.
func (in *AWSAction) DeepCopy() *AWSAction {
	if in == nil {
		return nil
	}
	out := new(AWSAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSStatement) DeepCopyInto(out *AWSStatement) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make(AWSActions, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSStatement.
func (in *AWSStatement) DeepCopy() *AWSStatement {
	if in == nil {
		return nil
	}
	out := new(AWSStatement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSTerraformProvisioner) DeepCopyInto(out *AWSTerraformProvisioner) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSTerraformProvisioner.
func (in *AWSTerraformProvisioner) DeepCopy() *AWSTerraformProvisioner {
	if in == nil {
		return nil
	}
	out := new(AWSTerraformProvisioner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSTerraformSpec) DeepCopyInto(out *AWSTerraformSpec) {
	*out = *in
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Statement.DeepCopyInto(&out.Statement)
	in.Docker.DeepCopyInto(&out.Docker)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSTerraformSpec.
func (in *AWSTerraformSpec) DeepCopy() *AWSTerraformSpec {
	if in == nil {
		return nil
	}
	out := new(AWSTerraformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDependency) DeepCopyInto(out *AppDependency) {
	*out = *in
	out.Package = in.Package
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDependency.
func (in *AppDependency) DeepCopy() *AppDependency {
	if in == nil {
		return nil
	}
	out := new(AppDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPU) DeepCopyInto(out *CPU) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPU.
func (in *CPU) DeepCopy() *CPU {
	if in == nil {
		return nil
	}
	out := new(CPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependencies) DeepCopyInto(out *Dependencies) {
	*out = *in
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make(PackageDependencies, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Apps != nil {
		in, out := &in.Apps, &out.Apps
		*out = make(AppDependencies, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependencies.
func (in *Dependencies) DeepCopy() *Dependencies {
	if in == nil {
		return nil
	}
	out := new(Dependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Directory) DeepCopyInto(out *Directory) {
	*out = *in
	if in.FSTypes != nil {
		in, out := &in.FSTypes, &out.FSTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Directory.
func (in *Directory) DeepCopy() *Directory {
	if in == nil {
		return nil
	}
	out := new(Directory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Docker) DeepCopyInto(out *Docker) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Docker.
func (in *Docker) DeepCopy() *Docker {
	if in == nil {
		return nil
	}
	out := new(Docker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EULA) DeepCopyInto(out *EULA) {
	*out = *in
	out.Source = in.Source
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EULA.
func (in *EULA) DeepCopy() *EULA {
	if in == nil {
		return nil
	}
	out := new(EULA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionExtension) DeepCopyInto(out *EncryptionExtension) {
	*out = *in
	out.EncryptionKey = in.EncryptionKey
	out.CACert = in.CACert
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionExtension.
func (in *EncryptionExtension) DeepCopy() *EncryptionExtension {
	if in == nil {
		return nil
	}
	out := new(EncryptionExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extensions) DeepCopyInto(out *Extensions) {
	*out = *in
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		if *in == nil {
			*out = nil
		} else {
			*out = new(EncryptionExtension)
			**out = **in
		}
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		if *in == nil {
			*out = nil
		} else {
			*out = new(MonitoringExtension)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extensions.
func (in *Extensions) DeepCopy() *Extensions {
	if in == nil {
		return nil
	}
	out := new(Extensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinalInstallStep) DeepCopyInto(out *FinalInstallStep) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinalInstallStep.
func (in *FinalInstallStep) DeepCopy() *FinalInstallStep {
	if in == nil {
		return nil
	}
	out := new(FinalInstallStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flavor) DeepCopyInto(out *Flavor) {
	*out = *in
	out.Threshold = in.Threshold
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flavor.
func (in *Flavor) DeepCopy() *Flavor {
	if in == nil {
		return nil
	}
	out := new(Flavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flavors) DeepCopyInto(out *Flavors) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Flavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flavors.
func (in *Flavors) DeepCopy() *Flavors {
	if in == nil {
		return nil
	}
	out := new(Flavors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Header) DeepCopyInto(out *Header) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Header.
func (in *Header) DeepCopy() *Header {
	if in == nil {
		return nil
	}
	out := new(Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hooks) DeepCopyInto(out *Hooks) {
	*out = *in
	if in.Install != nil {
		in, out := &in.Install, &out.Install
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Installed != nil {
		in, out := &in.Installed, &out.Installed
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Uninstall != nil {
		in, out := &in.Uninstall, &out.Uninstall
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Uninstalling != nil {
		in, out := &in.Uninstalling, &out.Uninstalling
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeAdding != nil {
		in, out := &in.NodeAdding, &out.NodeAdding
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeAdded != nil {
		in, out := &in.NodeAdded, &out.NodeAdded
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeRemoving != nil {
		in, out := &in.NodeRemoving, &out.NodeRemoving
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeRemoved != nil {
		in, out := &in.NodeRemoved, &out.NodeRemoved
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Updating != nil {
		in, out := &in.Updating, &out.Updating
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Rollback != nil {
		in, out := &in.Rollback, &out.Rollback
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RolledBack != nil {
		in, out := &in.RolledBack, &out.RolledBack
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.LicenseUpdated != nil {
		in, out := &in.LicenseUpdated, &out.LicenseUpdated
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Stop != nil {
		in, out := &in.Stop, &out.Stop
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Dump != nil {
		in, out := &in.Dump, &out.Dump
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		if *in == nil {
			*out = nil
		} else {
			*out = new(HooksBase)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hooks.
func (in *Hooks) DeepCopy() *Hooks {
	if in == nil {
		return nil
	}
	out := new(Hooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HooksBase) DeepCopyInto(out *HooksBase) {
	*out = *in
	in.JobSpec.DeepCopyInto(&out.JobSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HooksBase.
func (in *HooksBase) DeepCopy() *HooksBase {
	if in == nil {
		return nil
	}
	out := new(HooksBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Installer) DeepCopyInto(out *Installer) {
	*out = *in
	in.Provisioners.DeepCopyInto(&out.Provisioners)
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make(map[string]ServerProfile, len(*in))
		for key, val := range *in {
			newVal := new(ServerProfile)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	in.Flavors.DeepCopyInto(&out.Flavors)
	in.License.DeepCopyInto(&out.License)
	out.EULA = in.EULA
	if in.FinalInstallStep != nil {
		in, out := &in.FinalInstallStep, &out.FinalInstallStep
		if *in == nil {
			*out = nil
		} else {
			*out = new(FinalInstallStep)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Installer.
func (in *Installer) DeepCopy() *Installer {
	if in == nil {
		return nil
	}
	out := new(Installer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfig) DeepCopyInto(out *KubeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfig.
func (in *KubeConfig) DeepCopy() *KubeConfig {
	if in == nil {
		return nil
	}
	out := new(KubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *License) DeepCopyInto(out *License) {
	*out = *in
	if in.TrialFlavors != nil {
		in, out := &in.TrialFlavors, &out.TrialFlavors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new License.
func (in *License) DeepCopy() *License {
	if in == nil {
		return nil
	}
	out := new(License)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manifest) DeepCopyInto(out *Manifest) {
	*out = *in
	in.Header.DeepCopyInto(&out.Header)
	if in.Installer != nil {
		in, out := &in.Installer, &out.Installer
		if *in == nil {
			*out = nil
		} else {
			*out = new(Installer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Orchestration != nil {
		in, out := &in.Orchestration, &out.Orchestration
		if *in == nil {
			*out = nil
		} else {
			*out = new(Orchestration)
			**out = **in
		}
	}
	in.Dependencies.DeepCopyInto(&out.Dependencies)
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hooks)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Base != nil {
		in, out := &in.Base, &out.Base
		if *in == nil {
			*out = nil
		} else {
			*out = new(ManifestRef)
			**out = **in
		}
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebConfig != nil {
		in, out := &in.WebConfig, &out.WebConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.OpsCenterConfig != nil {
		in, out := &in.OpsCenterConfig, &out.OpsCenterConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(OpsCenterConfig)
			**out = **in
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		if *in == nil {
			*out = nil
		} else {
			*out = new(Extensions)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manifest.
func (in *Manifest) DeepCopy() *Manifest {
	if in == nil {
		return nil
	}
	out := new(Manifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Manifest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestRef) DeepCopyInto(out *ManifestRef) {
	*out = *in
	out.Package = in.Package
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestRef.
func (in *ManifestRef) DeepCopy() *ManifestRef {
	if in == nil {
		return nil
	}
	out := new(ManifestRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Logo != nil {
		in, out := &in.Logo, &out.Logo
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringExtension) DeepCopyInto(out *MonitoringExtension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringExtension.
func (in *MonitoringExtension) DeepCopy() *MonitoringExtension {
	if in == nil {
		return nil
	}
	out := new(MonitoringExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mount) DeepCopyInto(out *Mount) {
	*out = *in
	if in.FSTypes != nil {
		in, out := &in.FSTypes, &out.FSTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mount.
func (in *Mount) DeepCopy() *Mount {
	if in == nil {
		return nil
	}
	out := new(Mount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OS) DeepCopyInto(out *OS) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OS.
func (in *OS) DeepCopy() *OS {
	if in == nil {
		return nil
	}
	out := new(OS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnPremProvisioner) DeepCopyInto(out *OnPremProvisioner) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnPremProvisioner.
func (in *OnPremProvisioner) DeepCopy() *OnPremProvisioner {
	if in == nil {
		return nil
	}
	out := new(OnPremProvisioner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnPremSpec) DeepCopyInto(out *OnPremSpec) {
	*out = *in
	in.Docker.DeepCopyInto(&out.Docker)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnPremSpec.
func (in *OnPremSpec) DeepCopy() *OnPremSpec {
	if in == nil {
		return nil
	}
	out := new(OnPremSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsCenterConfig) DeepCopyInto(out *OpsCenterConfig) {
	*out = *in
	out.Address = in.Address
	out.Token = in.Token
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsCenterConfig.
func (in *OpsCenterConfig) DeepCopy() *OpsCenterConfig {
	if in == nil {
		return nil
	}
	out := new(OpsCenterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Orchestration) DeepCopyInto(out *Orchestration) {
	*out = *in
	out.KubeConfig = in.KubeConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Orchestration.
func (in *Orchestration) DeepCopy() *Orchestration {
	if in == nil {
		return nil
	}
	out := new(Orchestration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageDependency) DeepCopyInto(out *PackageDependency) {
	*out = *in
	out.Package = in.Package
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageDependency.
func (in *PackageDependency) DeepCopy() *PackageDependency {
	if in == nil {
		return nil
	}
	out := new(PackageDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provisioners) DeepCopyInto(out *Provisioners) {
	*out = *in
	if in.Virsh != nil {
		in, out := &in.Virsh, &out.Virsh
		if *in == nil {
			*out = nil
		} else {
			*out = new(VirshProvisioner)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.OnPrem != nil {
		in, out := &in.OnPrem, &out.OnPrem
		if *in == nil {
			*out = nil
		} else {
			*out = new(OnPremProvisioner)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AWSTerraform != nil {
		in, out := &in.AWSTerraform, &out.AWSTerraform
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSTerraformProvisioner)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provisioners.
func (in *Provisioners) DeepCopy() *Provisioners {
	if in == nil {
		return nil
	}
	out := new(Provisioners)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RAM) DeepCopyInto(out *RAM) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RAM.
func (in *RAM) DeepCopy() *RAM {
	if in == nil {
		return nil
	}
	out := new(RAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPorts) DeepCopyInto(out *ServerPorts) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPorts.
func (in *ServerPorts) DeepCopy() *ServerPorts {
	if in == nil {
		return nil
	}
	out := new(ServerPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfile) DeepCopyInto(out *ServerProfile) {
	*out = *in
	out.CPU = in.CPU
	out.RAM = in.RAM
	if in.OS != nil {
		in, out := &in.OS, &out.OS
		*out = make([]OS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ServerPorts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Disk = in.Disk
	out.Network = in.Network
	if in.Directories != nil {
		in, out := &in.Directories, &out.Directories
		*out = make(Directories, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make(Mounts, len(*in))
		for key, val := range *in {
			newVal := new(Mount)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	out.Request = in.Request
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfile.
func (in *ServerProfile) DeepCopy() *ServerProfile {
	if in == nil {
		return nil
	}
	out := new(ServerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerProfileRequest) DeepCopyInto(out *ServerProfileRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerProfileRequest.
func (in *ServerProfileRequest) DeepCopy() *ServerProfileRequest {
	if in == nil {
		return nil
	}
	out := new(ServerProfileRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Threshold) DeepCopyInto(out *Threshold) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Threshold.
func (in *Threshold) DeepCopy() *Threshold {
	if in == nil {
		return nil
	}
	out := new(Threshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserExtension) DeepCopyInto(out *UserExtension) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserExtension.
func (in *UserExtension) DeepCopy() *UserExtension {
	if in == nil {
		return nil
	}
	out := new(UserExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirshProvisioner) DeepCopyInto(out *VirshProvisioner) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirshProvisioner.
func (in *VirshProvisioner) DeepCopy() *VirshProvisioner {
	if in == nil {
		return nil
	}
	out := new(VirshProvisioner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirshSpec) DeepCopyInto(out *VirshSpec) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]Device, len(*in))
		copy(*out, *in)
	}
	in.Docker.DeepCopyInto(&out.Docker)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirshSpec.
func (in *VirshSpec) DeepCopy() *VirshSpec {
	if in == nil {
		return nil
	}
	out := new(VirshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebConfig) DeepCopyInto(out *WebConfig) {
	*out = *in
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebModules)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebConfig.
func (in *WebConfig) DeepCopy() *WebConfig {
	if in == nil {
		return nil
	}
	out := new(WebConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebFeature) DeepCopyInto(out *WebFeature) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebFeature.
func (in *WebFeature) DeepCopy() *WebFeature {
	if in == nil {
		return nil
	}
	out := new(WebFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebFeatures) DeepCopyInto(out *WebFeatures) {
	*out = *in
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebFeature)
			**out = **in
		}
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebFeature)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebFeatures.
func (in *WebFeatures) DeepCopy() *WebFeatures {
	if in == nil {
		return nil
	}
	out := new(WebFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebInstaller) DeepCopyInto(out *WebInstaller) {
	*out = *in
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebInstaller.
func (in *WebInstaller) DeepCopy() *WebInstaller {
	if in == nil {
		return nil
	}
	out := new(WebInstaller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebModules) DeepCopyInto(out *WebModules) {
	*out = *in
	if in.Installer != nil {
		in, out := &in.Installer, &out.Installer
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebInstaller)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Site != nil {
		in, out := &in.Site, &out.Site
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebSite)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebModules.
func (in *WebModules) DeepCopy() *WebModules {
	if in == nil {
		return nil
	}
	out := new(WebModules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebSite) DeepCopyInto(out *WebSite) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebFeatures)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebSite.
func (in *WebSite) DeepCopy() *WebSite {
	if in == nil {
		return nil
	}
	out := new(WebSite)
	in.DeepCopyInto(out)
	return out
}