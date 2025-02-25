package containers

import "github.com/nexclipper/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("containers")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("containers", id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("containers")
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("containers", id)
}

func listConsumersURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "consumers")
}

func createConsumerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "consumers")
}

func deleteConsumerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "consumers")
}

func createSecretRefURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "secrets")
}

func deleteSecretRefURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "secrets")
}
