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

type KeyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type KeyParameters struct {

	// +kubebuilder:validation:Optional
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerMasterKeySpec *string `json:"customerMasterKeySpec,omitempty" tf:"customer_master_key_spec,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionWindowInDays *float64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EnableKeyRotation *bool `json:"enableKeyRotation,omitempty" tf:"enable_key_rotation,omitempty"`

	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KeyUsage *string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// +kubebuilder:validation:Optional
	MultiRegion *bool `json:"multiRegion,omitempty" tf:"multi_region,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// KeySpec defines the desired state of Key
type KeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyParameters `json:"forProvider"`
}

// KeyStatus defines the observed state of Key.
type KeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Key is the Schema for the Keys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeySpec   `json:"spec"`
	Status            KeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyList contains a list of Keys
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

// Repository type metadata.
var (
	Key_Kind             = "Key"
	Key_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Key_Kind}.String()
	Key_KindAPIVersion   = Key_Kind + "." + CRDGroupVersion.String()
	Key_GroupVersionKind = CRDGroupVersion.WithKind(Key_Kind)
)

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
