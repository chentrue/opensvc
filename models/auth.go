package models

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"log"
)

func StackAuth (region string) (provider *gophercloud.ProviderClient) {
	reg, err := GetClouds(region)
	if err != nil {
		return nil
	}
	auth := gophercloud.AuthOptions{
		IdentityEndpoint: reg.Auth_url,
		Username: reg.Username,
		Password: reg.Password,
		DomainName: reg.Domain,
		TenantName: reg.Project,
	}
	prov, err := openstack.AuthenticatedClient(auth)
	if err != nil {
		log.Fatalf("some worng:", err.Error())
	}
	return prov
}
