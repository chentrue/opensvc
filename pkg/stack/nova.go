package stack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	servers2 "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"opensvc/models"
)


func Server_list (region string) (map[string][]interface{}, error){
	provider := models.StackAuth(region)
	sc, serviceErr := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if serviceErr != nil {
		return nil, serviceErr
	}
	allpages, err := servers2.List(sc, servers2.ListOpts{}).AllPages()
	if err != nil {
		return nil, err
	}
	servers, _ := allpages.GetBody().(map[string][]interface{})
	return servers, nil
}

