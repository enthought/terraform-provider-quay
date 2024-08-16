package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTeamResource(t *testing.T) {
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

resource "quay_team" "test" {
  name = "test"
  orgname = quay_organization.test.name
  role = "member"
  description = "test"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_team.test", "role", "member"),
					resource.TestCheckResourceAttr("quay_team.test", "description", "test"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_team.test",
				ImportState:                          true,
				ImportStateId:                        "test+test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			// Update
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}

resource "quay_team" "test" {
  name = "test"
  orgname = quay_organization.test.name
  role = "admin"
  description = "test2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_team.test", "description", "test2"),
				),
			},
			// Replace resource
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}

resource "quay_team" "test" {
  name = "test2"
  orgname = quay_organization.test.name
  role = "admin"
  description = "test2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_team.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_team.test", "description", "test2"),
				),
			},
		},
	})
}
