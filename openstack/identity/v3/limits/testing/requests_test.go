package testing

import (
	"testing"

	"github.com/nexclipper/gophercloud/openstack/identity/v3/limits"
	"github.com/nexclipper/gophercloud/pagination"
	th "github.com/nexclipper/gophercloud/testhelper"
	"github.com/nexclipper/gophercloud/testhelper/client"
)

func TestListLimits(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLimitsSuccessfully(t)

	count := 0
	err := limits.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := limits.ExtractLimits(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedLimitsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}

func TestListLimitsAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLimitsSuccessfully(t)

	allPages, err := limits.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := limits.ExtractLimits(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedLimitsSlice, actual)
}
