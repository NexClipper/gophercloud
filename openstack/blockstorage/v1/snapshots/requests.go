package snapshots

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/utils"
	"github.com/rackspace/gophercloud/pagination"

	"github.com/racker/perigee"
)

// CreateOpts contains options for creating a Snapshot. This object is passed to
// the snapshots.Create function. For more information about these parameters,
// see the Snapshot object.
type CreateOpts struct {
	Description string                 // OPTIONAL
	Force       bool                   // OPTIONAL
	Metadata    map[string]interface{} // OPTIONAL
	Name        string                 // OPTIONAL
	VolumeID    string                 // REQUIRED
}

// Create will create a new Snapshot based on the values in CreateOpts. To extract
// the Snapshot object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts *CreateOpts) CreateResult {
	type snapshot struct {
		Description *string                `json:"display_description,omitempty"`
		Force       bool                   `json:"force,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Name        *string                `json:"display_name,omitempty"`
		VolumeID    *string                `json:"volume_id,omitempty"`
	}

	type request struct {
		Snapshot snapshot `json:"snapshot"`
	}

	reqBody := request{
		Snapshot: snapshot{},
	}

	reqBody.Snapshot.Description = utils.MaybeString(opts.Description)
	reqBody.Snapshot.Name = utils.MaybeString(opts.Name)
	reqBody.Snapshot.VolumeID = utils.MaybeString(opts.VolumeID)

	reqBody.Snapshot.Force = opts.Force

	var res CreateResult
	_, res.Err = perigee.Request("POST", createURL(client), perigee.Options{
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		OkCodes:     []int{200, 201},
		ReqBody:     &reqBody,
		Results:     &res.Resp,
	})
	return res
}

// Delete will delete the existing Snapshot with the provided ID.
func Delete(client *gophercloud.ServiceClient, id string) error {
	_, err := perigee.Request("DELETE", deleteURL(client, id), perigee.Options{
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		OkCodes:     []int{202, 204},
	})
	return err
}

// Get retrieves the Snapshot with the provided ID. To extract the Snapshot object
// from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = perigee.Request("GET", getURL(client, id), perigee.Options{
		Results:     &res.Resp,
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		OkCodes:     []int{200},
	})
	return res
}

// ListOpts hold options for listing Snapshots. It is passed to the
// snapshots.List function.
type ListOpts struct {
	Name     string `q:"display_name"`
	Status   string `q:"status"`
	VolumeID string `q:"volume_id"`
}

// List returns Snapshots optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts *ListOpts) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := gophercloud.BuildQueryString(opts)
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query.String()
	}

	createPage := func(r pagination.LastHTTPResponse) pagination.Page {
		return ListResult{pagination.SinglePageBase(r)}
	}
	return pagination.NewPager(client, url, createPage)
}

// UpdateOpts contain options for updating an existing Snapshot. This object is
// passed to the snapshots.Update function. For more information about the
// parameters, see the Snapshot object.
type UpdateOpts struct {
	Description string
	Name        string
}

// Update will update the Snapshot with provided information. To extract the updated
// Snapshot from the response, call the Extract method on the UpdateResult.
func Update(client *gophercloud.ServiceClient, id string, opts *UpdateOpts) UpdateResult {
	type update struct {
		Description *string `json:"display_description,omitempty"`
		Name        *string `json:"display_name,omitempty"`
	}

	type request struct {
		Volume update `json:"snapshot"`
	}

	reqBody := request{
		Volume: update{},
	}

	reqBody.Volume.Description = utils.MaybeString(opts.Description)
	reqBody.Volume.Name = utils.MaybeString(opts.Name)

	var res UpdateResult

	_, res.Err = perigee.Request("PUT", updateURL(client, id), perigee.Options{
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		OkCodes:     []int{202, 204},
		ReqBody:     &reqBody,
		Results:     &res.Resp,
	})
	return res
}
