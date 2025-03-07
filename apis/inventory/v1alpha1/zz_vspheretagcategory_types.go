// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VSphereTagCategoryInitParameters struct {

	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	// Object types to which this category's tags can be attached. Valid types include: Folder, ClusterComputeResource, Datacenter, Datastore, StoragePod, DistributedVirtualPortgroup, DistributedVirtualSwitch, VmwareDistributedVirtualSwitch, HostSystem, com.vmware.content.Library, com.vmware.content.library.Item, HostNetwork, Network, OpaqueNetwork, ResourcePool, VirtualApp, VirtualMachine.
	// +listType=set
	AssociableTypes []*string `json:"associableTypes,omitempty" tf:"associable_types,omitempty"`

	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of SINGLE (object can only
	// be assigned one tag in this category), to MULTIPLE (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	// The associated cardinality of the category. Can be one of SINGLE (object can only be assigned one tag in this category) or MULTIPLE (object can be assigned multiple tags in this category).
	Cardinality *string `json:"cardinality,omitempty" tf:"cardinality,omitempty"`

	// A description for the category.
	// The description of the category.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the category.
	// The display name of the category.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VSphereTagCategoryObservation struct {

	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	// Object types to which this category's tags can be attached. Valid types include: Folder, ClusterComputeResource, Datacenter, Datastore, StoragePod, DistributedVirtualPortgroup, DistributedVirtualSwitch, VmwareDistributedVirtualSwitch, HostSystem, com.vmware.content.Library, com.vmware.content.library.Item, HostNetwork, Network, OpaqueNetwork, ResourcePool, VirtualApp, VirtualMachine.
	// +listType=set
	AssociableTypes []*string `json:"associableTypes,omitempty" tf:"associable_types,omitempty"`

	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of SINGLE (object can only
	// be assigned one tag in this category), to MULTIPLE (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	// The associated cardinality of the category. Can be one of SINGLE (object can only be assigned one tag in this category) or MULTIPLE (object can be assigned multiple tags in this category).
	Cardinality *string `json:"cardinality,omitempty" tf:"cardinality,omitempty"`

	// A description for the category.
	// The description of the category.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the category.
	// The display name of the category.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VSphereTagCategoryParameters struct {

	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	// Object types to which this category's tags can be attached. Valid types include: Folder, ClusterComputeResource, Datacenter, Datastore, StoragePod, DistributedVirtualPortgroup, DistributedVirtualSwitch, VmwareDistributedVirtualSwitch, HostSystem, com.vmware.content.Library, com.vmware.content.library.Item, HostNetwork, Network, OpaqueNetwork, ResourcePool, VirtualApp, VirtualMachine.
	// +kubebuilder:validation:Optional
	// +listType=set
	AssociableTypes []*string `json:"associableTypes,omitempty" tf:"associable_types,omitempty"`

	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of SINGLE (object can only
	// be assigned one tag in this category), to MULTIPLE (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	// The associated cardinality of the category. Can be one of SINGLE (object can only be assigned one tag in this category) or MULTIPLE (object can be assigned multiple tags in this category).
	// +kubebuilder:validation:Optional
	Cardinality *string `json:"cardinality,omitempty" tf:"cardinality,omitempty"`

	// A description for the category.
	// The description of the category.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the category.
	// The display name of the category.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// VSphereTagCategorySpec defines the desired state of VSphereTagCategory
type VSphereTagCategorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereTagCategoryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VSphereTagCategoryInitParameters `json:"initProvider,omitempty"`
}

// VSphereTagCategoryStatus defines the observed state of VSphereTagCategory.
type VSphereTagCategoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereTagCategoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VSphereTagCategory is the Schema for the VSphereTagCategorys API. Provides a vSphere tag category resource. This can be used to manage tag categories in vSphere.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereTagCategory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.associableTypes) || (has(self.initProvider) && has(self.initProvider.associableTypes))",message="spec.forProvider.associableTypes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cardinality) || (has(self.initProvider) && has(self.initProvider.cardinality))",message="spec.forProvider.cardinality is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VSphereTagCategorySpec   `json:"spec"`
	Status VSphereTagCategoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereTagCategoryList contains a list of VSphereTagCategorys
type VSphereTagCategoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereTagCategory `json:"items"`
}

// Repository type metadata.
var (
	VSphereTagCategory_Kind             = "VSphereTagCategory"
	VSphereTagCategory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereTagCategory_Kind}.String()
	VSphereTagCategory_KindAPIVersion   = VSphereTagCategory_Kind + "." + CRDGroupVersion.String()
	VSphereTagCategory_GroupVersionKind = CRDGroupVersion.WithKind(VSphereTagCategory_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereTagCategory{}, &VSphereTagCategoryList{})
}
