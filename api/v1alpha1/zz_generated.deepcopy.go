// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AffinitySettings) DeepCopyInto(out *AffinitySettings) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AffinitySettings.
func (in *AffinitySettings) DeepCopy() *AffinitySettings {
	if in == nil {
		return nil
	}
	out := new(AffinitySettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntime) DeepCopyInto(out *ContainerRuntime) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntime.
func (in *ContainerRuntime) DeepCopy() *ContainerRuntime {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JVB) DeepCopyInto(out *JVB) {
	*out = *in
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(ContainerRuntime)
		(*in).DeepCopyInto(*out)
	}
	in.AffinitySettings.DeepCopyInto(&out.AffinitySettings)
	in.Strategy.DeepCopyInto(&out.Strategy)
	in.Ports.DeepCopyInto(&out.Ports)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JVB.
func (in *JVB) DeepCopy() *JVB {
	if in == nil {
		return nil
	}
	out := new(JVB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JVBPorts) DeepCopyInto(out *JVBPorts) {
	*out = *in
	if in.UDP != nil {
		in, out := &in.UDP, &out.UDP
		*out = new(int32)
		**out = **in
	}
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JVBPorts.
func (in *JVBPorts) DeepCopy() *JVBPorts {
	if in == nil {
		return nil
	}
	out := new(JVBPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JVBStatus) DeepCopyInto(out *JVBStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JVBStatus.
func (in *JVBStatus) DeepCopy() *JVBStatus {
	if in == nil {
		return nil
	}
	out := new(JVBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JVBStrategy) DeepCopyInto(out *JVBStrategy) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JVBStrategy.
func (in *JVBStrategy) DeepCopy() *JVBStrategy {
	if in == nil {
		return nil
	}
	out := new(JVBStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jibri) DeepCopyInto(out *Jibri) {
	*out = *in
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(ContainerRuntime)
		(*in).DeepCopyInto(*out)
	}
	in.AffinitySettings.DeepCopyInto(&out.AffinitySettings)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jibri.
func (in *Jibri) DeepCopy() *Jibri {
	if in == nil {
		return nil
	}
	out := new(Jibri)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jicofo) DeepCopyInto(out *Jicofo) {
	*out = *in
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(ContainerRuntime)
		(*in).DeepCopyInto(*out)
	}
	in.AffinitySettings.DeepCopyInto(&out.AffinitySettings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jicofo.
func (in *Jicofo) DeepCopy() *Jicofo {
	if in == nil {
		return nil
	}
	out := new(Jicofo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jitsi) DeepCopyInto(out *Jitsi) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jitsi.
func (in *Jitsi) DeepCopy() *Jitsi {
	if in == nil {
		return nil
	}
	out := new(Jitsi)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jitsi) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JitsiList) DeepCopyInto(out *JitsiList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jitsi, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JitsiList.
func (in *JitsiList) DeepCopy() *JitsiList {
	if in == nil {
		return nil
	}
	out := new(JitsiList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JitsiList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JitsiSpec) DeepCopyInto(out *JitsiSpec) {
	*out = *in
	in.JVB.DeepCopyInto(&out.JVB)
	in.Prosody.DeepCopyInto(&out.Prosody)
	in.Jicofo.DeepCopyInto(&out.Jicofo)
	in.Jibri.DeepCopyInto(&out.Jibri)
	in.Web.DeepCopyInto(&out.Web)
	out.Version = in.Version
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Ingress.DeepCopyInto(&out.Ingress)
	if in.TURN != nil {
		in, out := &in.TURN, &out.TURN
		*out = new(TURN)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JitsiSpec.
func (in *JitsiSpec) DeepCopy() *JitsiSpec {
	if in == nil {
		return nil
	}
	out := new(JitsiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JitsiStatus) DeepCopyInto(out *JitsiStatus) {
	*out = *in
	out.JVBStatus = in.JVBStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JitsiStatus.
func (in *JitsiStatus) DeepCopy() *JitsiStatus {
	if in == nil {
		return nil
	}
	out := new(JitsiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prosody) DeepCopyInto(out *Prosody) {
	*out = *in
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(ContainerRuntime)
		(*in).DeepCopyInto(*out)
	}
	in.AffinitySettings.DeepCopyInto(&out.AffinitySettings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prosody.
func (in *Prosody) DeepCopy() *Prosody {
	if in == nil {
		return nil
	}
	out := new(Prosody)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TURN) DeepCopyInto(out *TURN) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TURN.
func (in *TURN) DeepCopy() *TURN {
	if in == nil {
		return nil
	}
	out := new(TURN)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Version) DeepCopyInto(out *Version) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Version.
func (in *Version) DeepCopy() *Version {
	if in == nil {
		return nil
	}
	out := new(Version)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Web) DeepCopyInto(out *Web) {
	*out = *in
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(ContainerRuntime)
		(*in).DeepCopyInto(*out)
	}
	in.AffinitySettings.DeepCopyInto(&out.AffinitySettings)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.CustomConfig != nil {
		in, out := &in.CustomConfig, &out.CustomConfig
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CustomInterfaceConfig != nil {
		in, out := &in.CustomInterfaceConfig, &out.CustomInterfaceConfig
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Web.
func (in *Web) DeepCopy() *Web {
	if in == nil {
		return nil
	}
	out := new(Web)
	in.DeepCopyInto(out)
	return out
}
