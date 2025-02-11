terraform {
  required_providers {
    vsphere = {
      source = "hashicorp/vsphere"
      version = "2.10.0"
    }
  }
}

provider "vsphere" {
  user = "administrator@vsphere.local"
  password = "Bpovtmg1!"
  vsphere_server = "10.100.100.2"
  allow_unverified_ssl = true
}

data "vsphere_datacenter" "dc1" {
  name = "my_prod_datacenter"
}
data "vsphere_datacenter" "dc2" {
  name = "Demo-Datacenter"
}
