package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	fake "github.com/nexclipper/gophercloud/openstack/networking/v2/common"
	"github.com/nexclipper/gophercloud/openstack/networking/v2/extensions/agents"
	"github.com/nexclipper/gophercloud/pagination"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/agents", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, AgentsListResult)
	})

	count := 0

	agents.List(fake.ServiceClient(), agents.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := agents.ExtractAgents(page)

		if err != nil {
			t.Errorf("Failed to extract agents: %v", err)
			return false, nil
		}

		expected := []agents.Agent{
			Agent1,
			Agent2,
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/agents/43583cf5-472e-4dc8-af5b-6aed4c94ee3a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, AgentsGetResult)
	})

	s, err := agents.Get(fake.ServiceClient(), "43583cf5-472e-4dc8-af5b-6aed4c94ee3a").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, s.ID, "43583cf5-472e-4dc8-af5b-6aed4c94ee3a")
	th.AssertEquals(t, s.Binary, "neutron-openvswitch-agent")
	th.AssertEquals(t, s.AdminStateUp, true)
	th.AssertEquals(t, s.Alive, true)
	th.AssertEquals(t, s.Topic, "N/A")
	th.AssertEquals(t, s.Host, "compute3")
	th.AssertEquals(t, s.AgentType, "Open vSwitch agent")
	th.AssertEquals(t, s.HeartbeatTimestamp, time.Date(2019, 1, 9, 11, 43, 01, 0, time.UTC))
	th.AssertEquals(t, s.StartedAt, time.Date(2018, 6, 26, 21, 46, 20, 0, time.UTC))
	th.AssertEquals(t, s.CreatedAt, time.Date(2017, 7, 26, 23, 2, 5, 0, time.UTC))
	th.AssertDeepEquals(t, s.Configurations, map[string]interface{}{
		"ovs_hybrid_plug":            false,
		"datapath_type":              "system",
		"vhostuser_socket_dir":       "/var/run/openvswitch",
		"log_agent_heartbeats":       false,
		"l2_population":              true,
		"enable_distributed_routing": false,
	})
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/agents/43583cf5-472e-4dc8-af5b-6aed4c94ee3a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, AgentUpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, AgentsUpdateResult)
	})

	iTrue := true
	description := "My OVS agent for OpenStack"
	updateOpts := &agents.UpdateOpts{
		Description:  &description,
		AdminStateUp: &iTrue,
	}
	s, err := agents.Update(fake.ServiceClient(), "43583cf5-472e-4dc8-af5b-6aed4c94ee3a", updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, *s, Agent)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/agents/43583cf5-472e-4dc8-af5b-6aed4c94ee3a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	err := agents.Delete(fake.ServiceClient(), "43583cf5-472e-4dc8-af5b-6aed4c94ee3a").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestListDHCPNetworks(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/agents/43583cf5-472e-4dc8-af5b-6aed4c94ee3a/dhcp-networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, AgentDHCPNetworksListResult)
	})

	s, err := agents.ListDHCPNetworks(fake.ServiceClient(), "43583cf5-472e-4dc8-af5b-6aed4c94ee3a").Extract()
	th.AssertNoErr(t, err)

	var nilSlice []string
	th.AssertEquals(t, len(s), 1)
	th.AssertEquals(t, s[0].ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertEquals(t, s[0].AdminStateUp, true)
	th.AssertEquals(t, s[0].ProjectID, "4fd44f30292945e481c7b8a0c8908869")
	th.AssertEquals(t, s[0].Shared, false)
	th.AssertEquals(t, s[0].Name, "net1")
	th.AssertEquals(t, s[0].Status, "ACTIVE")
	th.AssertDeepEquals(t, s[0].Tags, nilSlice)
	th.AssertEquals(t, s[0].TenantID, "4fd44f30292945e481c7b8a0c8908869")
	th.AssertDeepEquals(t, s[0].AvailabilityZoneHints, []string{})
	th.AssertDeepEquals(t, s[0].Subnets, []string{"54d6f61d-db07-451c-9ab3-b9609b6b6f0b"})

}

func TestScheduleDHCPNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/agents/43583cf5-472e-4dc8-af5b-6aed4c94ee3a/dhcp-networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, ScheduleDHCPNetworkRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	})

	opts := &agents.ScheduleDHCPNetworkOpts{
		NetworkID: "1ae075ca-708b-4e66-b4a7-b7698632f05f",
	}
	err := agents.ScheduleDHCPNetwork(fake.ServiceClient(), "43583cf5-472e-4dc8-af5b-6aed4c94ee3a", opts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestRemoveDHCPNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/agents/43583cf5-472e-4dc8-af5b-6aed4c94ee3a/dhcp-networks/1ae075ca-708b-4e66-b4a7-b7698632f05f", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	err := agents.RemoveDHCPNetwork(fake.ServiceClient(), "43583cf5-472e-4dc8-af5b-6aed4c94ee3a", "1ae075ca-708b-4e66-b4a7-b7698632f05f").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestListBGPSpeakers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	agentID := "30d76012-46de-4215-aaa1-a1630d01d891"

	th.Mux.HandleFunc("/v2.0/agents/"+agentID+"/bgp-drinstances",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, ListBGPSpeakersResult)
		})

	count := 0
	agents.ListBGPSpeakers(fake.ServiceClient(), agentID).EachPage(
		func(page pagination.Page) (bool, error) {
			count++
			actual, err := agents.ExtractBGPSpeakers(page)

			th.AssertNoErr(t, err)
			th.AssertEquals(t, len(actual), 1)
			th.AssertEquals(t, actual[0].ID, "cab00464-284d-4251-9798-2b27db7b1668")
			th.AssertEquals(t, actual[0].Name, "gophercloud-testing-speaker")
			th.AssertEquals(t, actual[0].LocalAS, 12345)
			th.AssertEquals(t, actual[0].IPVersion, 4)
			return true, nil
		})
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}
