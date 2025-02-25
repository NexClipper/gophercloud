//go:build acceptance || compute || rescueunrescue
// +build acceptance compute rescueunrescue

package v2

import (
	"testing"

	"github.com/nexclipper/gophercloud/acceptance/clients"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestServerRescueUnrescue(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	err = RescueServer(t, client, server)
	th.AssertNoErr(t, err)

	err = UnrescueServer(t, client, server)
	th.AssertNoErr(t, err)
}
