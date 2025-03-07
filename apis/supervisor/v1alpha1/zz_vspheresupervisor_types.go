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

type EgressCidrInitParameters struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type EgressCidrObservation struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type EgressCidrParameters struct {

	// Network address.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// Subnet prefix.
	// +kubebuilder:validation:Optional
	Prefix *float64 `json:"prefix" tf:"prefix,omitempty"`
}

type IngressCidrInitParameters struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type IngressCidrObservation struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type IngressCidrParameters struct {

	// Network address.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// Subnet prefix.
	// +kubebuilder:validation:Optional
	Prefix *float64 `json:"prefix" tf:"prefix,omitempty"`
}

type ManagementNetworkInitParameters struct {

	// Number of addresses to allocate. Starts from starting_address
	// Number of addresses to allocate. Starts from 'starting_address'
	AddressCount *float64 `json:"addressCount,omitempty" tf:"address_count,omitempty"`

	// Gateway IP address.
	// Gateway IP address.
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// ID of the network. (e.g. a distributed port group).
	// ID of the network. (e.g. a distributed port group).
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Starting address of the management network range.
	// Starting address of the management network range.
	StartingAddress *string `json:"startingAddress,omitempty" tf:"starting_address,omitempty"`

	// Subnet mask.
	// Subnet mask.
	SubnetMask *string `json:"subnetMask,omitempty" tf:"subnet_mask,omitempty"`
}

type ManagementNetworkObservation struct {

	// Number of addresses to allocate. Starts from starting_address
	// Number of addresses to allocate. Starts from 'starting_address'
	AddressCount *float64 `json:"addressCount,omitempty" tf:"address_count,omitempty"`

	// Gateway IP address.
	// Gateway IP address.
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// ID of the network. (e.g. a distributed port group).
	// ID of the network. (e.g. a distributed port group).
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Starting address of the management network range.
	// Starting address of the management network range.
	StartingAddress *string `json:"startingAddress,omitempty" tf:"starting_address,omitempty"`

	// Subnet mask.
	// Subnet mask.
	SubnetMask *string `json:"subnetMask,omitempty" tf:"subnet_mask,omitempty"`
}

type ManagementNetworkParameters struct {

	// Number of addresses to allocate. Starts from starting_address
	// Number of addresses to allocate. Starts from 'starting_address'
	// +kubebuilder:validation:Optional
	AddressCount *float64 `json:"addressCount" tf:"address_count,omitempty"`

	// Gateway IP address.
	// Gateway IP address.
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`

	// ID of the network. (e.g. a distributed port group).
	// ID of the network. (e.g. a distributed port group).
	// +kubebuilder:validation:Optional
	Network *string `json:"network" tf:"network,omitempty"`

	// Starting address of the management network range.
	// Starting address of the management network range.
	// +kubebuilder:validation:Optional
	StartingAddress *string `json:"startingAddress" tf:"starting_address,omitempty"`

	// Subnet mask.
	// Subnet mask.
	// +kubebuilder:validation:Optional
	SubnetMask *string `json:"subnetMask" tf:"subnet_mask,omitempty"`
}

type NamespaceInitParameters struct {

	// The list of content libraries to associate with the namespace
	// A list of content libraries.
	ContentLibraries []*string `json:"contentLibraries,omitempty" tf:"content_libraries,omitempty"`

	// The name of the namespace
	// The name of the namespace.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The list of virtual machine classes to add to the namespace
	// A list of virtual machine classes.
	VMClasses []*string `json:"vmClasses,omitempty" tf:"vm_classes,omitempty"`
}

type NamespaceObservation struct {

	// The list of content libraries to associate with the namespace
	// A list of content libraries.
	ContentLibraries []*string `json:"contentLibraries,omitempty" tf:"content_libraries,omitempty"`

	// The name of the namespace
	// The name of the namespace.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The list of virtual machine classes to add to the namespace
	// A list of virtual machine classes.
	VMClasses []*string `json:"vmClasses,omitempty" tf:"vm_classes,omitempty"`
}

type NamespaceParameters struct {

	// The list of content libraries to associate with the namespace
	// A list of content libraries.
	// +kubebuilder:validation:Optional
	ContentLibraries []*string `json:"contentLibraries,omitempty" tf:"content_libraries,omitempty"`

	// The name of the namespace
	// The name of the namespace.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The list of virtual machine classes to add to the namespace
	// A list of virtual machine classes.
	// +kubebuilder:validation:Optional
	VMClasses []*string `json:"vmClasses,omitempty" tf:"vm_classes,omitempty"`
}

type PodCidrInitParameters struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type PodCidrObservation struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type PodCidrParameters struct {

	// Network address.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// Subnet prefix.
	// +kubebuilder:validation:Optional
	Prefix *float64 `json:"prefix" tf:"prefix,omitempty"`
}

type ServiceCidrInitParameters struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type ServiceCidrObservation struct {

	// Network address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Subnet prefix.
	Prefix *float64 `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type ServiceCidrParameters struct {

	// Network address.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// Subnet prefix.
	// +kubebuilder:validation:Optional
	Prefix *float64 `json:"prefix" tf:"prefix,omitempty"`
}

type VSphereSupervisorInitParameters struct {

	// The identifier of the compute cluster.
	// ID of the vSphere cluster on which workload management will be enabled.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The identifier of the subscribed content library.
	// ID of the subscribed content library.
	ContentLibrary *string `json:"contentLibrary,omitempty" tf:"content_library,omitempty"`

	// The UUID of the distributed switch.
	// The UUID (not ID) of the distributed switch.
	DvsUUID *string `json:"dvsUuid,omitempty" tf:"dvs_uuid,omitempty"`

	// The identifier of the NSX Edge Cluster.
	// ID of the NSX Edge Cluster.
	EdgeCluster *string `json:"edgeCluster,omitempty" tf:"edge_cluster,omitempty"`

	// CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
	// CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
	EgressCidr []EgressCidrInitParameters `json:"egressCidr,omitempty" tf:"egress_cidr,omitempty"`

	// CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
	// CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
	IngressCidr []IngressCidrInitParameters `json:"ingressCidr,omitempty" tf:"ingress_cidr,omitempty"`

	// The list of addresses of the primary DNS servers.
	// List of DNS servers to use on the Kubernetes API server.
	MainDNS []*string `json:"mainDns,omitempty" tf:"main_dns,omitempty"`

	// The configuration for the management network which the control plane VMs will be connected to.
	// The configuration for the management network which the control plane VMs will be connected to.
	ManagementNetwork []ManagementNetworkInitParameters `json:"managementNetwork,omitempty" tf:"management_network,omitempty"`

	// The list of namespaces to create in the Supervisor cluster
	// List of namespaces associated with the cluster.
	Namespace []NamespaceInitParameters `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
	// CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
	PodCidr []PodCidrInitParameters `json:"podCidr,omitempty" tf:"pod_cidr,omitempty"`

	// List of DNS search domains.
	// List of DNS search domains.
	SearchDomains []*string `json:"searchDomains,omitempty" tf:"search_domains,omitempty"`

	// CIDR block from which Kubernetes allocates service cluster IP addresses.
	// CIDR block from which Kubernetes allocates service cluster IP addresses.
	ServiceCidr []ServiceCidrInitParameters `json:"serviceCidr,omitempty" tf:"service_cidr,omitempty"`

	// The size of the Kubernetes API server.
	// Size of the Kubernetes API server.
	SizingHint *string `json:"sizingHint,omitempty" tf:"sizing_hint,omitempty"`

	// The name of the storage policy.
	// The name of a storage policy associated with the datastore where the container images will be stored.
	StoragePolicy *string `json:"storagePolicy,omitempty" tf:"storage_policy,omitempty"`

	// The list of addresses of the DNS servers to use for the worker nodes.
	// List of DNS servers to use on the worker nodes.
	WorkerDNS []*string `json:"workerDns,omitempty" tf:"worker_dns,omitempty"`
}

type VSphereSupervisorObservation struct {

	// The identifier of the compute cluster.
	// ID of the vSphere cluster on which workload management will be enabled.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The identifier of the subscribed content library.
	// ID of the subscribed content library.
	ContentLibrary *string `json:"contentLibrary,omitempty" tf:"content_library,omitempty"`

	// The UUID of the distributed switch.
	// The UUID (not ID) of the distributed switch.
	DvsUUID *string `json:"dvsUuid,omitempty" tf:"dvs_uuid,omitempty"`

	// The identifier of the NSX Edge Cluster.
	// ID of the NSX Edge Cluster.
	EdgeCluster *string `json:"edgeCluster,omitempty" tf:"edge_cluster,omitempty"`

	// CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
	// CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
	EgressCidr []EgressCidrObservation `json:"egressCidr,omitempty" tf:"egress_cidr,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
	// CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
	IngressCidr []IngressCidrObservation `json:"ingressCidr,omitempty" tf:"ingress_cidr,omitempty"`

	// The list of addresses of the primary DNS servers.
	// List of DNS servers to use on the Kubernetes API server.
	MainDNS []*string `json:"mainDns,omitempty" tf:"main_dns,omitempty"`

	// The configuration for the management network which the control plane VMs will be connected to.
	// The configuration for the management network which the control plane VMs will be connected to.
	ManagementNetwork []ManagementNetworkObservation `json:"managementNetwork,omitempty" tf:"management_network,omitempty"`

	// The list of namespaces to create in the Supervisor cluster
	// List of namespaces associated with the cluster.
	Namespace []NamespaceObservation `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
	// CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
	PodCidr []PodCidrObservation `json:"podCidr,omitempty" tf:"pod_cidr,omitempty"`

	// List of DNS search domains.
	// List of DNS search domains.
	SearchDomains []*string `json:"searchDomains,omitempty" tf:"search_domains,omitempty"`

	// CIDR block from which Kubernetes allocates service cluster IP addresses.
	// CIDR block from which Kubernetes allocates service cluster IP addresses.
	ServiceCidr []ServiceCidrObservation `json:"serviceCidr,omitempty" tf:"service_cidr,omitempty"`

	// The size of the Kubernetes API server.
	// Size of the Kubernetes API server.
	SizingHint *string `json:"sizingHint,omitempty" tf:"sizing_hint,omitempty"`

	// The name of the storage policy.
	// The name of a storage policy associated with the datastore where the container images will be stored.
	StoragePolicy *string `json:"storagePolicy,omitempty" tf:"storage_policy,omitempty"`

	// The list of addresses of the DNS servers to use for the worker nodes.
	// List of DNS servers to use on the worker nodes.
	WorkerDNS []*string `json:"workerDns,omitempty" tf:"worker_dns,omitempty"`
}

type VSphereSupervisorParameters struct {

	// The identifier of the compute cluster.
	// ID of the vSphere cluster on which workload management will be enabled.
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The identifier of the subscribed content library.
	// ID of the subscribed content library.
	// +kubebuilder:validation:Optional
	ContentLibrary *string `json:"contentLibrary,omitempty" tf:"content_library,omitempty"`

	// The UUID of the distributed switch.
	// The UUID (not ID) of the distributed switch.
	// +kubebuilder:validation:Optional
	DvsUUID *string `json:"dvsUuid,omitempty" tf:"dvs_uuid,omitempty"`

	// The identifier of the NSX Edge Cluster.
	// ID of the NSX Edge Cluster.
	// +kubebuilder:validation:Optional
	EdgeCluster *string `json:"edgeCluster,omitempty" tf:"edge_cluster,omitempty"`

	// CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
	// CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
	// +kubebuilder:validation:Optional
	EgressCidr []EgressCidrParameters `json:"egressCidr,omitempty" tf:"egress_cidr,omitempty"`

	// CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
	// CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
	// +kubebuilder:validation:Optional
	IngressCidr []IngressCidrParameters `json:"ingressCidr,omitempty" tf:"ingress_cidr,omitempty"`

	// The list of addresses of the primary DNS servers.
	// List of DNS servers to use on the Kubernetes API server.
	// +kubebuilder:validation:Optional
	MainDNS []*string `json:"mainDns,omitempty" tf:"main_dns,omitempty"`

	// The configuration for the management network which the control plane VMs will be connected to.
	// The configuration for the management network which the control plane VMs will be connected to.
	// +kubebuilder:validation:Optional
	ManagementNetwork []ManagementNetworkParameters `json:"managementNetwork,omitempty" tf:"management_network,omitempty"`

	// The list of namespaces to create in the Supervisor cluster
	// List of namespaces associated with the cluster.
	// +kubebuilder:validation:Optional
	Namespace []NamespaceParameters `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
	// CIDR blocks from which Kubernetes allocates pod IP addresses. Minimum subnet size is 23.
	// +kubebuilder:validation:Optional
	PodCidr []PodCidrParameters `json:"podCidr,omitempty" tf:"pod_cidr,omitempty"`

	// List of DNS search domains.
	// List of DNS search domains.
	// +kubebuilder:validation:Optional
	SearchDomains []*string `json:"searchDomains,omitempty" tf:"search_domains,omitempty"`

	// CIDR block from which Kubernetes allocates service cluster IP addresses.
	// CIDR block from which Kubernetes allocates service cluster IP addresses.
	// +kubebuilder:validation:Optional
	ServiceCidr []ServiceCidrParameters `json:"serviceCidr,omitempty" tf:"service_cidr,omitempty"`

	// The size of the Kubernetes API server.
	// Size of the Kubernetes API server.
	// +kubebuilder:validation:Optional
	SizingHint *string `json:"sizingHint,omitempty" tf:"sizing_hint,omitempty"`

	// The name of the storage policy.
	// The name of a storage policy associated with the datastore where the container images will be stored.
	// +kubebuilder:validation:Optional
	StoragePolicy *string `json:"storagePolicy,omitempty" tf:"storage_policy,omitempty"`

	// The list of addresses of the DNS servers to use for the worker nodes.
	// List of DNS servers to use on the worker nodes.
	// +kubebuilder:validation:Optional
	WorkerDNS []*string `json:"workerDns,omitempty" tf:"worker_dns,omitempty"`
}

// VSphereSupervisorSpec defines the desired state of VSphereSupervisor
type VSphereSupervisorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereSupervisorParameters `json:"forProvider"`
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
	InitProvider VSphereSupervisorInitParameters `json:"initProvider,omitempty"`
}

// VSphereSupervisorStatus defines the observed state of VSphereSupervisor.
type VSphereSupervisorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereSupervisorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VSphereSupervisor is the Schema for the VSphereSupervisors API. Provides a VMware vSphere Supervisor resource..
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereSupervisor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cluster) || (has(self.initProvider) && has(self.initProvider.cluster))",message="spec.forProvider.cluster is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.contentLibrary) || (has(self.initProvider) && has(self.initProvider.contentLibrary))",message="spec.forProvider.contentLibrary is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dvsUuid) || (has(self.initProvider) && has(self.initProvider.dvsUuid))",message="spec.forProvider.dvsUuid is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.edgeCluster) || (has(self.initProvider) && has(self.initProvider.edgeCluster))",message="spec.forProvider.edgeCluster is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.egressCidr) || (has(self.initProvider) && has(self.initProvider.egressCidr))",message="spec.forProvider.egressCidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ingressCidr) || (has(self.initProvider) && has(self.initProvider.ingressCidr))",message="spec.forProvider.ingressCidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mainDns) || (has(self.initProvider) && has(self.initProvider.mainDns))",message="spec.forProvider.mainDns is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.managementNetwork) || (has(self.initProvider) && has(self.initProvider.managementNetwork))",message="spec.forProvider.managementNetwork is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.podCidr) || (has(self.initProvider) && has(self.initProvider.podCidr))",message="spec.forProvider.podCidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.searchDomains) || (has(self.initProvider) && has(self.initProvider.searchDomains))",message="spec.forProvider.searchDomains is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceCidr) || (has(self.initProvider) && has(self.initProvider.serviceCidr))",message="spec.forProvider.serviceCidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sizingHint) || (has(self.initProvider) && has(self.initProvider.sizingHint))",message="spec.forProvider.sizingHint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storagePolicy) || (has(self.initProvider) && has(self.initProvider.storagePolicy))",message="spec.forProvider.storagePolicy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.workerDns) || (has(self.initProvider) && has(self.initProvider.workerDns))",message="spec.forProvider.workerDns is a required parameter"
	Spec   VSphereSupervisorSpec   `json:"spec"`
	Status VSphereSupervisorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereSupervisorList contains a list of VSphereSupervisors
type VSphereSupervisorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereSupervisor `json:"items"`
}

// Repository type metadata.
var (
	VSphereSupervisor_Kind             = "VSphereSupervisor"
	VSphereSupervisor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereSupervisor_Kind}.String()
	VSphereSupervisor_KindAPIVersion   = VSphereSupervisor_Kind + "." + CRDGroupVersion.String()
	VSphereSupervisor_GroupVersionKind = CRDGroupVersion.WithKind(VSphereSupervisor_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereSupervisor{}, &VSphereSupervisorList{})
}
