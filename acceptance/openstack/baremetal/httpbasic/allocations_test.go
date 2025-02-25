//go:build acceptance || baremetal || allocations
// +build acceptance baremetal allocations

package httpbasic

import (
	"testing"

	"github.com/nexclipper/gophercloud/acceptance/clients"
	v1 "github.com/nexclipper/gophercloud/acceptance/openstack/baremetal/v1"
	"github.com/nexclipper/gophercloud/openstack/baremetal/v1/allocations"
	"github.com/nexclipper/gophercloud/pagination"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestAllocationsCreateDestroy(t *testing.T) {
	clients.RequireLong(t)
	clients.RequireIronicHTTPBasic(t)

	client, err := clients.NewBareMetalV1HTTPBasic()
	th.AssertNoErr(t, err)

	client.Microversion = "1.52"

	allocation, err := v1.CreateAllocation(t, client)
	th.AssertNoErr(t, err)
	defer v1.DeleteAllocation(t, client, allocation)

	found := false
	err = allocations.List(client, allocations.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		allocationList, err := allocations.ExtractAllocations(page)
		if err != nil {
			return false, err
		}

		for _, a := range allocationList {
			if a.UUID == allocation.UUID {
				found = true
				return true, nil
			}
		}

		return false, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, found, true)
}
