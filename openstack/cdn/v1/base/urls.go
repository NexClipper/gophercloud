package base

import "github.com/nexclipper/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL()
}

func pingURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("ping")
}
