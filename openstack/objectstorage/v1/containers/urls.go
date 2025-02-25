package containers

import "github.com/nexclipper/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}

func createURL(c *gophercloud.ServiceClient, container string) string {
	return c.ServiceURL(container)
}

func getURL(c *gophercloud.ServiceClient, container string) string {
	return createURL(c, container)
}

func deleteURL(c *gophercloud.ServiceClient, container string) string {
	return createURL(c, container)
}

func updateURL(c *gophercloud.ServiceClient, container string) string {
	return createURL(c, container)
}

func bulkDeleteURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint + "?bulk-delete=true"
}
