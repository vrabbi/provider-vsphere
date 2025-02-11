package vspherevmfsdatastore

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_vmfs_datastore", func(r *config.Resource) {
		r.ShortGroup = "Storage"
		r.Kind = "VSphereVmfsDatastore"
		r.Version = "v1alpha1"
	})
}
