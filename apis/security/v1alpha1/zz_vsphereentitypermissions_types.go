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

type PermissionsInitParameters struct {

	// Whether user_or_group field refers to a user or a group. True for a group and false for a user.
	IsGroup *bool `json:"isGroup,omitempty" tf:"is_group,omitempty"`

	// Whether or not this permission propagates down the hierarchy to sub-entities.
	Propagate *bool `json:"propagate,omitempty" tf:"propagate,omitempty"`

	// Reference to the role providing the access.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// User or group receiving access.
	UserOrGroup *string `json:"userOrGroup,omitempty" tf:"user_or_group,omitempty"`
}

type PermissionsObservation struct {

	// Whether user_or_group field refers to a user or a group. True for a group and false for a user.
	IsGroup *bool `json:"isGroup,omitempty" tf:"is_group,omitempty"`

	// Whether or not this permission propagates down the hierarchy to sub-entities.
	Propagate *bool `json:"propagate,omitempty" tf:"propagate,omitempty"`

	// Reference to the role providing the access.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// User or group receiving access.
	UserOrGroup *string `json:"userOrGroup,omitempty" tf:"user_or_group,omitempty"`
}

type PermissionsParameters struct {

	// Whether user_or_group field refers to a user or a group. True for a group and false for a user.
	// +kubebuilder:validation:Optional
	IsGroup *bool `json:"isGroup" tf:"is_group,omitempty"`

	// Whether or not this permission propagates down the hierarchy to sub-entities.
	// +kubebuilder:validation:Optional
	Propagate *bool `json:"propagate" tf:"propagate,omitempty"`

	// Reference to the role providing the access.
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId" tf:"role_id,omitempty"`

	// User or group receiving access.
	// +kubebuilder:validation:Optional
	UserOrGroup *string `json:"userOrGroup" tf:"user_or_group,omitempty"`
}

type VSphereEntityPermissionsInitParameters struct {

	// The managed object id or uuid of the entity.
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The entity managed object type.
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// Permissions to be given to the entity.
	Permissions []PermissionsInitParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type VSphereEntityPermissionsObservation struct {

	// The managed object id or uuid of the entity.
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The entity managed object type.
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Permissions to be given to the entity.
	Permissions []PermissionsObservation `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type VSphereEntityPermissionsParameters struct {

	// The managed object id or uuid of the entity.
	// +kubebuilder:validation:Optional
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The entity managed object type.
	// +kubebuilder:validation:Optional
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// Permissions to be given to the entity.
	// +kubebuilder:validation:Optional
	Permissions []PermissionsParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

// VSphereEntityPermissionsSpec defines the desired state of VSphereEntityPermissions
type VSphereEntityPermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereEntityPermissionsParameters `json:"forProvider"`
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
	InitProvider VSphereEntityPermissionsInitParameters `json:"initProvider,omitempty"`
}

// VSphereEntityPermissionsStatus defines the observed state of VSphereEntityPermissions.
type VSphereEntityPermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereEntityPermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VSphereEntityPermissions is the Schema for the VSphereEntityPermissionss API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereEntityPermissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entityId) || (has(self.initProvider) && has(self.initProvider.entityId))",message="spec.forProvider.entityId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entityType) || (has(self.initProvider) && has(self.initProvider.entityType))",message="spec.forProvider.entityType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissions) || (has(self.initProvider) && has(self.initProvider.permissions))",message="spec.forProvider.permissions is a required parameter"
	Spec   VSphereEntityPermissionsSpec   `json:"spec"`
	Status VSphereEntityPermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereEntityPermissionsList contains a list of VSphereEntityPermissionss
type VSphereEntityPermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereEntityPermissions `json:"items"`
}

// Repository type metadata.
var (
	VSphereEntityPermissions_Kind             = "VSphereEntityPermissions"
	VSphereEntityPermissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereEntityPermissions_Kind}.String()
	VSphereEntityPermissions_KindAPIVersion   = VSphereEntityPermissions_Kind + "." + CRDGroupVersion.String()
	VSphereEntityPermissions_GroupVersionKind = CRDGroupVersion.WithKind(VSphereEntityPermissions_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereEntityPermissions{}, &VSphereEntityPermissionsList{})
}
