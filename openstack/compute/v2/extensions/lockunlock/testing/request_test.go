package testing

import (
	"testing"

	"github.com/nexclipper/gophercloud/openstack/compute/v2/extensions/lockunlock"
	th "github.com/nexclipper/gophercloud/testhelper"
	"github.com/nexclipper/gophercloud/testhelper/client"
)

const serverID = "{serverId}"

func TestLock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStartServerResponse(t, serverID)

	err := lockunlock.Lock(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnlock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStopServerResponse(t, serverID)

	err := lockunlock.Unlock(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
