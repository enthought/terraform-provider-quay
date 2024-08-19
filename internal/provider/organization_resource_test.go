package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org" {
  name = "org"
  email = "quay+org@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.org", "name", "org"),
					resource.TestCheckResourceAttr("quay_organization.org", "email", "quay+org@example.com"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization.org",
				ImportState:                          true,
				ImportStateId:                        "org",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			// Update name
			{
				Config: providerConfig + `
resource "quay_organization" "org" {
  name = "org2"
  email = "quay+org@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.org", "name", "org2"),
					resource.TestCheckResourceAttr("quay_organization.org", "email", "quay+org@example.com"),
				),
			},
			// Update email
			{
				Config: providerConfig + `
resource "quay_organization" "org" {
  name = "org2"
  email = "quay+org2@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.org", "name", "org2"),
					resource.TestCheckResourceAttr("quay_organization.org", "email", "quay+org2@example.com"),
				),
			},
			// Update name and email
			{
				Config: providerConfig + `
resource "quay_organization" "org" {
  name = "org3"
  email = "quay+org3@example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization.org", "name", "org3"),
					resource.TestCheckResourceAttr("quay_organization.org", "email", "quay+org3@example.com"),
				),
			},
		},
	})
}
