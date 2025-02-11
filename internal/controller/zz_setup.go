// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	vspherelicense "github.com/vrabbi/provider-vsphere/internal/controller/administration/vspherelicense"
	vspherecomputecluster "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherecomputecluster"
	vspherecomputeclusterhostgroup "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherecomputeclusterhostgroup"
	vspherecomputeclustervmaffinityrule "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherecomputeclustervmaffinityrule"
	vspherecomputeclustervmantiaffinityrule "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherecomputeclustervmantiaffinityrule"
	vspherecomputeclustervmdependencyrule "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherecomputeclustervmdependencyrule"
	vspherecomputeclustervmgroup "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherecomputeclustervmgroup"
	vspherecomputeclustervmhostrule "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherecomputeclustervmhostrule"
	vspheredpmhostoverride "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspheredpmhostoverride"
	vspheredrsvmoverride "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspheredrsvmoverride"
	vspherehavmoverride "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherehavmoverride"
	vspherehost "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherehost"
	vsphereresourcepool "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vsphereresourcepool"
	vspherevnic "github.com/vrabbi/provider-vsphere/internal/controller/hostandclustermanagement/vspherevnic"
	vspherecustomattribute "github.com/vrabbi/provider-vsphere/internal/controller/inventory/vspherecustomattribute"
	vspheredatacenter "github.com/vrabbi/provider-vsphere/internal/controller/inventory/vspheredatacenter"
	vspherefolder "github.com/vrabbi/provider-vsphere/internal/controller/inventory/vspherefolder"
	vspheretag "github.com/vrabbi/provider-vsphere/internal/controller/inventory/vspheretag"
	vspheretagcategory "github.com/vrabbi/provider-vsphere/internal/controller/inventory/vspheretagcategory"
	vsphereofflinesoftwaredepot "github.com/vrabbi/provider-vsphere/internal/controller/lifecycle/vsphereofflinesoftwaredepot"
	vspheredistributedportgroup "github.com/vrabbi/provider-vsphere/internal/controller/networking/vspheredistributedportgroup"
	vspheredistributedvirtualswitch "github.com/vrabbi/provider-vsphere/internal/controller/networking/vspheredistributedvirtualswitch"
	vspherehostportgroup "github.com/vrabbi/provider-vsphere/internal/controller/networking/vspherehostportgroup"
	vspherehostvirtualswitch "github.com/vrabbi/provider-vsphere/internal/controller/networking/vspherehostvirtualswitch"
	providerconfig "github.com/vrabbi/provider-vsphere/internal/controller/providerconfig"
	vsphereentitypermissions "github.com/vrabbi/provider-vsphere/internal/controller/security/vsphereentitypermissions"
	vsphererole "github.com/vrabbi/provider-vsphere/internal/controller/security/vsphererole"
	vspheredatastorecluster "github.com/vrabbi/provider-vsphere/internal/controller/storage/vspheredatastorecluster"
	vspheredatastoreclustervmantiaffinityrule "github.com/vrabbi/provider-vsphere/internal/controller/storage/vspheredatastoreclustervmantiaffinityrule"
	vspherefile "github.com/vrabbi/provider-vsphere/internal/controller/storage/vspherefile"
	vspherenasdatastore "github.com/vrabbi/provider-vsphere/internal/controller/storage/vspherenasdatastore"
	vspherestoragedrsvmoverride "github.com/vrabbi/provider-vsphere/internal/controller/storage/vspherestoragedrsvmoverride"
	vspherevmfsdatastore "github.com/vrabbi/provider-vsphere/internal/controller/storage/vspherevmfsdatastore"
	vspherevmstoragepolicy "github.com/vrabbi/provider-vsphere/internal/controller/storage/vspherevmstoragepolicy"
	vspheresupervisor "github.com/vrabbi/provider-vsphere/internal/controller/supervisor/vspheresupervisor"
	vspherevirtualmachineclass "github.com/vrabbi/provider-vsphere/internal/controller/supervisor/vspherevirtualmachineclass"
	vspherecontentlibrary "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vspherecontentlibrary"
	vspherecontentlibraryitem "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vspherecontentlibraryitem"
	vsphereguestoscustomization "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vsphereguestoscustomization"
	vspherevappcontainer "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vspherevappcontainer"
	vspherevappentity "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vspherevappentity"
	vspherevirtualdisk "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vspherevirtualdisk"
	vspherevirtualmachine "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vspherevirtualmachine"
	vspherevirtualmachinesnapshot "github.com/vrabbi/provider-vsphere/internal/controller/virtualmachine/vspherevirtualmachinesnapshot"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vspherelicense.Setup,
		vspherecomputecluster.Setup,
		vspherecomputeclusterhostgroup.Setup,
		vspherecomputeclustervmaffinityrule.Setup,
		vspherecomputeclustervmantiaffinityrule.Setup,
		vspherecomputeclustervmdependencyrule.Setup,
		vspherecomputeclustervmgroup.Setup,
		vspherecomputeclustervmhostrule.Setup,
		vspheredpmhostoverride.Setup,
		vspheredrsvmoverride.Setup,
		vspherehavmoverride.Setup,
		vspherehost.Setup,
		vsphereresourcepool.Setup,
		vspherevnic.Setup,
		vspherecustomattribute.Setup,
		vspheredatacenter.Setup,
		vspherefolder.Setup,
		vspheretag.Setup,
		vspheretagcategory.Setup,
		vsphereofflinesoftwaredepot.Setup,
		vspheredistributedportgroup.Setup,
		vspheredistributedvirtualswitch.Setup,
		vspherehostportgroup.Setup,
		vspherehostvirtualswitch.Setup,
		providerconfig.Setup,
		vsphereentitypermissions.Setup,
		vsphererole.Setup,
		vspheredatastorecluster.Setup,
		vspheredatastoreclustervmantiaffinityrule.Setup,
		vspherefile.Setup,
		vspherenasdatastore.Setup,
		vspherestoragedrsvmoverride.Setup,
		vspherevmfsdatastore.Setup,
		vspherevmstoragepolicy.Setup,
		vspheresupervisor.Setup,
		vspherevirtualmachineclass.Setup,
		vspherecontentlibrary.Setup,
		vspherecontentlibraryitem.Setup,
		vsphereguestoscustomization.Setup,
		vspherevappcontainer.Setup,
		vspherevappentity.Setup,
		vspherevirtualdisk.Setup,
		vspherevirtualmachine.Setup,
		vspherevirtualmachinesnapshot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
