//go:build acceptance || identity
// +build acceptance identity

package v2

import (
	"testing"

	"github.com/nexclipper/gophercloud/acceptance/clients"
	"github.com/nexclipper/gophercloud/acceptance/tools"
	"github.com/nexclipper/gophercloud/openstack/identity/v2/extensions"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestExtensionsList(t *testing.T) {
	clients.RequireIdentityV2(t)
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV2Client()
	th.AssertNoErr(t, err)

	allPages, err := extensions.List(client).AllPages()
	th.AssertNoErr(t, err)

	allExtensions, err := extensions.ExtractExtensions(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, extension := range allExtensions {
		tools.PrintResource(t, extension)
		if extension.Name == "OS-KSCRUD" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestExtensionsGet(t *testing.T) {
	clients.RequireIdentityV2(t)
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV2Client()
	th.AssertNoErr(t, err)

	extension, err := extensions.Get(client, "OS-KSCRUD").Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, extension)
}
