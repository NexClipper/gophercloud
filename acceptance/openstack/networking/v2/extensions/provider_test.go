//go:build acceptance || networking || provider
// +build acceptance networking provider

package extensions

import (
	"testing"

	"github.com/nexclipper/gophercloud/acceptance/clients"
	networking "github.com/nexclipper/gophercloud/acceptance/openstack/networking/v2"
	"github.com/nexclipper/gophercloud/acceptance/tools"
	"github.com/nexclipper/gophercloud/openstack/networking/v2/networks"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestNetworksProviderCRUD(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create a network
	network, err := networking.CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer networking.DeleteNetwork(t, client, network.ID)

	getResult := networks.Get(client, network.ID)
	newNetwork, err := getResult.Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newNetwork)
}
