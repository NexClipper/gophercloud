//go:build acceptance || compute || availabilityzones
// +build acceptance compute availabilityzones

package v2

import (
	"testing"

	"github.com/nexclipper/gophercloud/acceptance/clients"
	"github.com/nexclipper/gophercloud/acceptance/tools"
	"github.com/nexclipper/gophercloud/openstack/compute/v2/extensions/availabilityzones"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestAvailabilityZonesList(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	allPages, err := availabilityzones.List(client).AllPages()
	th.AssertNoErr(t, err)

	availabilityZoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, zoneInfo := range availabilityZoneInfo {
		tools.PrintResource(t, zoneInfo)

		if zoneInfo.ZoneName == "nova" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestAvailabilityZonesListDetail(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	allPages, err := availabilityzones.ListDetail(client).AllPages()
	th.AssertNoErr(t, err)

	availabilityZoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, zoneInfo := range availabilityZoneInfo {
		tools.PrintResource(t, zoneInfo)

		if zoneInfo.ZoneName == "nova" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}
