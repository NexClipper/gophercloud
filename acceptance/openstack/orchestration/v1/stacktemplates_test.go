//go:build acceptance
// +build acceptance

package v1

import (
	"testing"

	"github.com/nexclipper/gophercloud/acceptance/clients"
	"github.com/nexclipper/gophercloud/acceptance/tools"
	"github.com/nexclipper/gophercloud/openstack/orchestration/v1/stacktemplates"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestStackTemplatesCRUD(t *testing.T) {
	client, err := clients.NewOrchestrationV1Client()
	th.AssertNoErr(t, err)

	stack, err := CreateStack(t, client)
	th.AssertNoErr(t, err)
	defer DeleteStack(t, client, stack.Name, stack.ID)

	tmpl, err := stacktemplates.Get(client, stack.Name, stack.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, tmpl)
}

func TestStackTemplatesValidate(t *testing.T) {
	client, err := clients.NewOrchestrationV1Client()
	th.AssertNoErr(t, err)

	validateOpts := stacktemplates.ValidateOpts{
		Template: validateTemplate,
	}

	validatedTemplate, err := stacktemplates.Validate(client, validateOpts).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, validatedTemplate)
}

func TestStackTemplateWithFile(t *testing.T) {
	client, err := clients.NewOrchestrationV1Client()
	th.AssertNoErr(t, err)

	stack, err := CreateStackWithFile(t, client)
	th.AssertNoErr(t, err)
	defer DeleteStack(t, client, stack.Name, stack.ID)

	tmpl, err := stacktemplates.Get(client, stack.Name, stack.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, tmpl)
}
