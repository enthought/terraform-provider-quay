package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization.test", "email", "quay+test@example.com"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization.test",
				ImportState:                          true,
				ImportStateId:                        "test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			// Update name
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test2"
  email = "quay+test@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization.test", "email", "quay+test@example.com"),
				),
			},
			// Update email
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test2"
  email = "quay+test2@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization.test", "email", "quay+test2@example.com"),
				),
			},
			// Update name and email
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test3"
  email = "quay+test3@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.test", "name", "test3"),
					resource.TestCheckResourceAttr("quay_organization.test", "email", "quay+test3@example.com"),
				),
			},
		},
	})
}
