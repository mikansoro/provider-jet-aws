//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedSecurityOptionsObservation) DeepCopyInto(out *AdvancedSecurityOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedSecurityOptionsObservation.
func (in *AdvancedSecurityOptionsObservation) DeepCopy() *AdvancedSecurityOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(AdvancedSecurityOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedSecurityOptionsParameters) DeepCopyInto(out *AdvancedSecurityOptionsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.InternalUserDatabaseEnabled != nil {
		in, out := &in.InternalUserDatabaseEnabled, &out.InternalUserDatabaseEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MasterUserOptions != nil {
		in, out := &in.MasterUserOptions, &out.MasterUserOptions
		*out = make([]MasterUserOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedSecurityOptionsParameters.
func (in *AdvancedSecurityOptionsParameters) DeepCopy() *AdvancedSecurityOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(AdvancedSecurityOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigObservation) DeepCopyInto(out *ClusterConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigObservation.
func (in *ClusterConfigObservation) DeepCopy() *ClusterConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigParameters) DeepCopyInto(out *ClusterConfigParameters) {
	*out = *in
	if in.DedicatedMasterCount != nil {
		in, out := &in.DedicatedMasterCount, &out.DedicatedMasterCount
		*out = new(float64)
		**out = **in
	}
	if in.DedicatedMasterEnabled != nil {
		in, out := &in.DedicatedMasterEnabled, &out.DedicatedMasterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DedicatedMasterType != nil {
		in, out := &in.DedicatedMasterType, &out.DedicatedMasterType
		*out = new(string)
		**out = **in
	}
	if in.InstanceCount != nil {
		in, out := &in.InstanceCount, &out.InstanceCount
		*out = new(float64)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.WarmCount != nil {
		in, out := &in.WarmCount, &out.WarmCount
		*out = new(float64)
		**out = **in
	}
	if in.WarmEnabled != nil {
		in, out := &in.WarmEnabled, &out.WarmEnabled
		*out = new(bool)
		**out = **in
	}
	if in.WarmType != nil {
		in, out := &in.WarmType, &out.WarmType
		*out = new(string)
		**out = **in
	}
	if in.ZoneAwarenessConfig != nil {
		in, out := &in.ZoneAwarenessConfig, &out.ZoneAwarenessConfig
		*out = make([]ZoneAwarenessConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ZoneAwarenessEnabled != nil {
		in, out := &in.ZoneAwarenessEnabled, &out.ZoneAwarenessEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigParameters.
func (in *ClusterConfigParameters) DeepCopy() *ClusterConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoOptionsObservation) DeepCopyInto(out *CognitoOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoOptionsObservation.
func (in *CognitoOptionsObservation) DeepCopy() *CognitoOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(CognitoOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoOptionsParameters) DeepCopyInto(out *CognitoOptionsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IdentityPoolID != nil {
		in, out := &in.IdentityPoolID, &out.IdentityPoolID
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoOptionsParameters.
func (in *CognitoOptionsParameters) DeepCopy() *CognitoOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(CognitoOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainEndpointOptionsObservation) DeepCopyInto(out *DomainEndpointOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainEndpointOptionsObservation.
func (in *DomainEndpointOptionsObservation) DeepCopy() *DomainEndpointOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(DomainEndpointOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainEndpointOptionsParameters) DeepCopyInto(out *DomainEndpointOptionsParameters) {
	*out = *in
	if in.CustomEndpoint != nil {
		in, out := &in.CustomEndpoint, &out.CustomEndpoint
		*out = new(string)
		**out = **in
	}
	if in.CustomEndpointCertificateArn != nil {
		in, out := &in.CustomEndpointCertificateArn, &out.CustomEndpointCertificateArn
		*out = new(string)
		**out = **in
	}
	if in.CustomEndpointEnabled != nil {
		in, out := &in.CustomEndpointEnabled, &out.CustomEndpointEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EnforceHTTPS != nil {
		in, out := &in.EnforceHTTPS, &out.EnforceHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.TLSSecurityPolicy != nil {
		in, out := &in.TLSSecurityPolicy, &out.TLSSecurityPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainEndpointOptionsParameters.
func (in *DomainEndpointOptionsParameters) DeepCopy() *DomainEndpointOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(DomainEndpointOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainObservation) DeepCopyInto(out *DomainObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DomainID != nil {
		in, out := &in.DomainID, &out.DomainID
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KibanaEndpoint != nil {
		in, out := &in.KibanaEndpoint, &out.KibanaEndpoint
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainObservation.
func (in *DomainObservation) DeepCopy() *DomainObservation {
	if in == nil {
		return nil
	}
	out := new(DomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainParameters) DeepCopyInto(out *DomainParameters) {
	*out = *in
	if in.AccessPolicies != nil {
		in, out := &in.AccessPolicies, &out.AccessPolicies
		*out = new(string)
		**out = **in
	}
	if in.AdvancedOptions != nil {
		in, out := &in.AdvancedOptions, &out.AdvancedOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AdvancedSecurityOptions != nil {
		in, out := &in.AdvancedSecurityOptions, &out.AdvancedSecurityOptions
		*out = make([]AdvancedSecurityOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterConfig != nil {
		in, out := &in.ClusterConfig, &out.ClusterConfig
		*out = make([]ClusterConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CognitoOptions != nil {
		in, out := &in.CognitoOptions, &out.CognitoOptions
		*out = make([]CognitoOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DomainEndpointOptions != nil {
		in, out := &in.DomainEndpointOptions, &out.DomainEndpointOptions
		*out = make([]DomainEndpointOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EBSOptions != nil {
		in, out := &in.EBSOptions, &out.EBSOptions
		*out = make([]EBSOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ElasticsearchVersion != nil {
		in, out := &in.ElasticsearchVersion, &out.ElasticsearchVersion
		*out = new(string)
		**out = **in
	}
	if in.EncryptAtRest != nil {
		in, out := &in.EncryptAtRest, &out.EncryptAtRest
		*out = make([]EncryptAtRestParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LogPublishingOptions != nil {
		in, out := &in.LogPublishingOptions, &out.LogPublishingOptions
		*out = make([]LogPublishingOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeToNodeEncryption != nil {
		in, out := &in.NodeToNodeEncryption, &out.NodeToNodeEncryption
		*out = make([]NodeToNodeEncryptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SnapshotOptions != nil {
		in, out := &in.SnapshotOptions, &out.SnapshotOptions
		*out = make([]SnapshotOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCOptions != nil {
		in, out := &in.VPCOptions, &out.VPCOptions
		*out = make([]VPCOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainParameters.
func (in *DomainParameters) DeepCopy() *DomainParameters {
	if in == nil {
		return nil
	}
	out := new(DomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPolicy) DeepCopyInto(out *DomainPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPolicy.
func (in *DomainPolicy) DeepCopy() *DomainPolicy {
	if in == nil {
		return nil
	}
	out := new(DomainPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPolicyList) DeepCopyInto(out *DomainPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPolicyList.
func (in *DomainPolicyList) DeepCopy() *DomainPolicyList {
	if in == nil {
		return nil
	}
	out := new(DomainPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPolicyObservation) DeepCopyInto(out *DomainPolicyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPolicyObservation.
func (in *DomainPolicyObservation) DeepCopy() *DomainPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(DomainPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPolicyParameters) DeepCopyInto(out *DomainPolicyParameters) {
	*out = *in
	if in.AccessPolicies != nil {
		in, out := &in.AccessPolicies, &out.AccessPolicies
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.DomainRefs != nil {
		in, out := &in.DomainRefs, &out.DomainRefs
		*out = new(v1.Reference)
		**out = **in
	}
	if in.DomainSelector != nil {
		in, out := &in.DomainSelector, &out.DomainSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPolicyParameters.
func (in *DomainPolicyParameters) DeepCopy() *DomainPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(DomainPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPolicySpec) DeepCopyInto(out *DomainPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPolicySpec.
func (in *DomainPolicySpec) DeepCopy() *DomainPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DomainPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPolicyStatus) DeepCopyInto(out *DomainPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPolicyStatus.
func (in *DomainPolicyStatus) DeepCopy() *DomainPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DomainPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSAMLOptions) DeepCopyInto(out *DomainSAMLOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSAMLOptions.
func (in *DomainSAMLOptions) DeepCopy() *DomainSAMLOptions {
	if in == nil {
		return nil
	}
	out := new(DomainSAMLOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainSAMLOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSAMLOptionsList) DeepCopyInto(out *DomainSAMLOptionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainSAMLOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSAMLOptionsList.
func (in *DomainSAMLOptionsList) DeepCopy() *DomainSAMLOptionsList {
	if in == nil {
		return nil
	}
	out := new(DomainSAMLOptionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainSAMLOptionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSAMLOptionsObservation) DeepCopyInto(out *DomainSAMLOptionsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSAMLOptionsObservation.
func (in *DomainSAMLOptionsObservation) DeepCopy() *DomainSAMLOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(DomainSAMLOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSAMLOptionsParameters) DeepCopyInto(out *DomainSAMLOptionsParameters) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.DomainRefs != nil {
		in, out := &in.DomainRefs, &out.DomainRefs
		*out = new(v1.Reference)
		**out = **in
	}
	if in.DomainSelector != nil {
		in, out := &in.DomainSelector, &out.DomainSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SAMLOptions != nil {
		in, out := &in.SAMLOptions, &out.SAMLOptions
		*out = make([]SAMLOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSAMLOptionsParameters.
func (in *DomainSAMLOptionsParameters) DeepCopy() *DomainSAMLOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(DomainSAMLOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSAMLOptionsSpec) DeepCopyInto(out *DomainSAMLOptionsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSAMLOptionsSpec.
func (in *DomainSAMLOptionsSpec) DeepCopy() *DomainSAMLOptionsSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSAMLOptionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSAMLOptionsStatus) DeepCopyInto(out *DomainSAMLOptionsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSAMLOptionsStatus.
func (in *DomainSAMLOptionsStatus) DeepCopy() *DomainSAMLOptionsStatus {
	if in == nil {
		return nil
	}
	out := new(DomainSAMLOptionsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EBSOptionsObservation) DeepCopyInto(out *EBSOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBSOptionsObservation.
func (in *EBSOptionsObservation) DeepCopy() *EBSOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(EBSOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EBSOptionsParameters) DeepCopyInto(out *EBSOptionsParameters) {
	*out = *in
	if in.EBSEnabled != nil {
		in, out := &in.EBSEnabled, &out.EBSEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(float64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBSOptionsParameters.
func (in *EBSOptionsParameters) DeepCopy() *EBSOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(EBSOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptAtRestObservation) DeepCopyInto(out *EncryptAtRestObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptAtRestObservation.
func (in *EncryptAtRestObservation) DeepCopy() *EncryptAtRestObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptAtRestObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptAtRestParameters) DeepCopyInto(out *EncryptAtRestParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyIDRef != nil {
		in, out := &in.KMSKeyIDRef, &out.KMSKeyIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.KMSKeyIDSelector != nil {
		in, out := &in.KMSKeyIDSelector, &out.KMSKeyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptAtRestParameters.
func (in *EncryptAtRestParameters) DeepCopy() *EncryptAtRestParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptAtRestParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpObservation) DeepCopyInto(out *IdpObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpObservation.
func (in *IdpObservation) DeepCopy() *IdpObservation {
	if in == nil {
		return nil
	}
	out := new(IdpObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpParameters) DeepCopyInto(out *IdpParameters) {
	*out = *in
	if in.EntityID != nil {
		in, out := &in.EntityID, &out.EntityID
		*out = new(string)
		**out = **in
	}
	if in.MetadataContent != nil {
		in, out := &in.MetadataContent, &out.MetadataContent
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpParameters.
func (in *IdpParameters) DeepCopy() *IdpParameters {
	if in == nil {
		return nil
	}
	out := new(IdpParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogPublishingOptionsObservation) DeepCopyInto(out *LogPublishingOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogPublishingOptionsObservation.
func (in *LogPublishingOptionsObservation) DeepCopy() *LogPublishingOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(LogPublishingOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogPublishingOptionsParameters) DeepCopyInto(out *LogPublishingOptionsParameters) {
	*out = *in
	if in.CloudwatchLogGroupArn != nil {
		in, out := &in.CloudwatchLogGroupArn, &out.CloudwatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogType != nil {
		in, out := &in.LogType, &out.LogType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogPublishingOptionsParameters.
func (in *LogPublishingOptionsParameters) DeepCopy() *LogPublishingOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(LogPublishingOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterUserOptionsObservation) DeepCopyInto(out *MasterUserOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterUserOptionsObservation.
func (in *MasterUserOptionsObservation) DeepCopy() *MasterUserOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(MasterUserOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterUserOptionsParameters) DeepCopyInto(out *MasterUserOptionsParameters) {
	*out = *in
	if in.MasterUserArn != nil {
		in, out := &in.MasterUserArn, &out.MasterUserArn
		*out = new(string)
		**out = **in
	}
	if in.MasterUserName != nil {
		in, out := &in.MasterUserName, &out.MasterUserName
		*out = new(string)
		**out = **in
	}
	if in.MasterUserPasswordSecretRef != nil {
		in, out := &in.MasterUserPasswordSecretRef, &out.MasterUserPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterUserOptionsParameters.
func (in *MasterUserOptionsParameters) DeepCopy() *MasterUserOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(MasterUserOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeToNodeEncryptionObservation) DeepCopyInto(out *NodeToNodeEncryptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeToNodeEncryptionObservation.
func (in *NodeToNodeEncryptionObservation) DeepCopy() *NodeToNodeEncryptionObservation {
	if in == nil {
		return nil
	}
	out := new(NodeToNodeEncryptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeToNodeEncryptionParameters) DeepCopyInto(out *NodeToNodeEncryptionParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeToNodeEncryptionParameters.
func (in *NodeToNodeEncryptionParameters) DeepCopy() *NodeToNodeEncryptionParameters {
	if in == nil {
		return nil
	}
	out := new(NodeToNodeEncryptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLOptionsObservation) DeepCopyInto(out *SAMLOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLOptionsObservation.
func (in *SAMLOptionsObservation) DeepCopy() *SAMLOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(SAMLOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLOptionsParameters) DeepCopyInto(out *SAMLOptionsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Idp != nil {
		in, out := &in.Idp, &out.Idp
		*out = make([]IdpParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MasterBackendRole != nil {
		in, out := &in.MasterBackendRole, &out.MasterBackendRole
		*out = new(string)
		**out = **in
	}
	if in.MasterUserNameSecretRef != nil {
		in, out := &in.MasterUserNameSecretRef, &out.MasterUserNameSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.RolesKey != nil {
		in, out := &in.RolesKey, &out.RolesKey
		*out = new(string)
		**out = **in
	}
	if in.SessionTimeoutMinutes != nil {
		in, out := &in.SessionTimeoutMinutes, &out.SessionTimeoutMinutes
		*out = new(float64)
		**out = **in
	}
	if in.SubjectKey != nil {
		in, out := &in.SubjectKey, &out.SubjectKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLOptionsParameters.
func (in *SAMLOptionsParameters) DeepCopy() *SAMLOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(SAMLOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotOptionsObservation) DeepCopyInto(out *SnapshotOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotOptionsObservation.
func (in *SnapshotOptionsObservation) DeepCopy() *SnapshotOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(SnapshotOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotOptionsParameters) DeepCopyInto(out *SnapshotOptionsParameters) {
	*out = *in
	if in.AutomatedSnapshotStartHour != nil {
		in, out := &in.AutomatedSnapshotStartHour, &out.AutomatedSnapshotStartHour
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotOptionsParameters.
func (in *SnapshotOptionsParameters) DeepCopy() *SnapshotOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCOptionsObservation) DeepCopyInto(out *VPCOptionsObservation) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCOptionsObservation.
func (in *VPCOptionsObservation) DeepCopy() *VPCOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(VPCOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCOptionsParameters) DeepCopyInto(out *VPCOptionsParameters) {
	*out = *in
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupRefs != nil {
		in, out := &in.SecurityGroupRefs, &out.SecurityGroupRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupSelector != nil {
		in, out := &in.SecurityGroupSelector, &out.SecurityGroupSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetRefs != nil {
		in, out := &in.SubnetRefs, &out.SubnetRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SubnetSelector != nil {
		in, out := &in.SubnetSelector, &out.SubnetSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCOptionsParameters.
func (in *VPCOptionsParameters) DeepCopy() *VPCOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(VPCOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAwarenessConfigObservation) DeepCopyInto(out *ZoneAwarenessConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAwarenessConfigObservation.
func (in *ZoneAwarenessConfigObservation) DeepCopy() *ZoneAwarenessConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ZoneAwarenessConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAwarenessConfigParameters) DeepCopyInto(out *ZoneAwarenessConfigParameters) {
	*out = *in
	if in.AvailabilityZoneCount != nil {
		in, out := &in.AvailabilityZoneCount, &out.AvailabilityZoneCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAwarenessConfigParameters.
func (in *ZoneAwarenessConfigParameters) DeepCopy() *ZoneAwarenessConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ZoneAwarenessConfigParameters)
	in.DeepCopyInto(out)
	return out
}
