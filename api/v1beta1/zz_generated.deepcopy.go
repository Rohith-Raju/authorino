// +build !ignore_autogenerated

/*
Copyright 2020 Red Hat, Inc.

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

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization) DeepCopyInto(out *Authorization) {
	*out = *in
	if in.OPAPolicy != nil {
		in, out := &in.OPAPolicy, &out.OPAPolicy
		*out = new(Authorization_OPAPolicy)
		**out = **in
	}
	if in.JWTClaimSet != nil {
		in, out := &in.JWTClaimSet, &out.JWTClaimSet
		*out = new(Authorization_JWTClaimSet)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization.
func (in *Authorization) DeepCopy() *Authorization {
	if in == nil {
		return nil
	}
	out := new(Authorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_JWTClaimSet) DeepCopyInto(out *Authorization_JWTClaimSet) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(Authorization_JWTClaimSet_Match)
		(*in).DeepCopyInto(*out)
	}
	if in.Claim != nil {
		in, out := &in.Claim, &out.Claim
		*out = new(Authorization_JWTClaimSet_Claim)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_JWTClaimSet.
func (in *Authorization_JWTClaimSet) DeepCopy() *Authorization_JWTClaimSet {
	if in == nil {
		return nil
	}
	out := new(Authorization_JWTClaimSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_JWTClaimSet_Claim) DeepCopyInto(out *Authorization_JWTClaimSet_Claim) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_JWTClaimSet_Claim.
func (in *Authorization_JWTClaimSet_Claim) DeepCopy() *Authorization_JWTClaimSet_Claim {
	if in == nil {
		return nil
	}
	out := new(Authorization_JWTClaimSet_Claim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_JWTClaimSet_HTTPMatch) DeepCopyInto(out *Authorization_JWTClaimSet_HTTPMatch) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_JWTClaimSet_HTTPMatch.
func (in *Authorization_JWTClaimSet_HTTPMatch) DeepCopy() *Authorization_JWTClaimSet_HTTPMatch {
	if in == nil {
		return nil
	}
	out := new(Authorization_JWTClaimSet_HTTPMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_JWTClaimSet_Match) DeepCopyInto(out *Authorization_JWTClaimSet_Match) {
	*out = *in
	if in.Http != nil {
		in, out := &in.Http, &out.Http
		*out = new(Authorization_JWTClaimSet_HTTPMatch)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_JWTClaimSet_Match.
func (in *Authorization_JWTClaimSet_Match) DeepCopy() *Authorization_JWTClaimSet_Match {
	if in == nil {
		return nil
	}
	out := new(Authorization_JWTClaimSet_Match)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization_OPAPolicy) DeepCopyInto(out *Authorization_OPAPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization_OPAPolicy.
func (in *Authorization_OPAPolicy) DeepCopy() *Authorization_OPAPolicy {
	if in == nil {
		return nil
	}
	out := new(Authorization_OPAPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credentials) DeepCopyInto(out *Credentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credentials.
func (in *Credentials) DeepCopy() *Credentials {
	if in == nil {
		return nil
	}
	out := new(Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	out.Credentials = in.Credentials
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = new(Identity_OidcConfig)
		**out = **in
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(Identity_APIKey)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_APIKey) DeepCopyInto(out *Identity_APIKey) {
	*out = *in
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_APIKey.
func (in *Identity_APIKey) DeepCopy() *Identity_APIKey {
	if in == nil {
		return nil
	}
	out := new(Identity_APIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_OidcConfig) DeepCopyInto(out *Identity_OidcConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_OidcConfig.
func (in *Identity_OidcConfig) DeepCopy() *Identity_OidcConfig {
	if in == nil {
		return nil
	}
	out := new(Identity_OidcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		*out = new(Metadata_UserInfo)
		**out = **in
	}
	if in.UMA != nil {
		in, out := &in.UMA, &out.UMA
		*out = new(Metadata_UMA)
		(*in).DeepCopyInto(*out)
	}
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
func (in *Metadata_UMA) DeepCopyInto(out *Metadata_UMA) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata_UMA.
func (in *Metadata_UMA) DeepCopy() *Metadata_UMA {
	if in == nil {
		return nil
	}
	out := new(Metadata_UMA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata_UserInfo) DeepCopyInto(out *Metadata_UserInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata_UserInfo.
func (in *Metadata_UserInfo) DeepCopy() *Metadata_UserInfo {
	if in == nil {
		return nil
	}
	out := new(Metadata_UserInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Service) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceList) DeepCopyInto(out *ServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceList.
func (in *ServiceList) DeepCopy() *ServiceList {
	if in == nil {
		return nil
	}
	out := new(ServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]*Identity, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Identity)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]*Metadata, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Metadata)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = make([]*Authorization, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Authorization)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}
