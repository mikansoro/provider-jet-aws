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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DomainPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainPolicyParameters struct {

	// +kubebuilder:validation:Required
	AccessPolicies *string `json:"accessPolicies" tf:"access_policies,omitempty"`

	// +crossplane:generate:reference:type=Domain
	// +crossplane:generate:reference:refFieldName=DomainRefs
	// +crossplane:generate:reference:selectorFieldName=DomainSelector
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	DomainRefs *v1.Reference `json:"domainRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DomainSelector *v1.Selector `json:"domainSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainPolicySpec defines the desired state of DomainPolicy
type DomainPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainPolicyParameters `json:"forProvider"`
}

// DomainPolicyStatus defines the observed state of DomainPolicy.
type DomainPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainPolicy is the Schema for the DomainPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DomainPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainPolicySpec   `json:"spec"`
	Status            DomainPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainPolicyList contains a list of DomainPolicys
type DomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainPolicy `json:"items"`
}

// Repository type metadata.
var (
	DomainPolicy_Kind             = "DomainPolicy"
	DomainPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainPolicy_Kind}.String()
	DomainPolicy_KindAPIVersion   = DomainPolicy_Kind + "." + CRDGroupVersion.String()
	DomainPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DomainPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainPolicy{}, &DomainPolicyList{})
}
