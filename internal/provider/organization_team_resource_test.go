package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationTeamResource(t *testing.T) {
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

resource "quay_organization_robot" "test" {
 name = "test"
 orgname = quay_organization.test.name
}

resource "quay_organization_team" "test" {
 name = "test"
 orgname = quay_organization.test.name
 role = "member"
 description = "test"
 members = [
   quay_organization_robot.test.fullname
 ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "test+test"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization_team.test",
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

resource "quay_organization_robot" "test" {
 name = "test"
 orgname = quay_organization.test.name
}

resource "quay_organization_team" "test" {
 name = "test"
 orgname = quay_organization.test.name
 role = "admin"
 description = "test2"
 members = [
   quay_organization_robot.test.fullname
 ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "test+test"),
				),
			},
			// Replace resource
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
 name = "test"
 email = "quay+test@example.com"
}

resource "quay_organization_robot" "test" {
 name = "test"
 orgname = quay_organization.test.name
}

resource "quay_organization_team" "test" {
 name = "test2"
 orgname = quay_organization.test.name
 role = "admin"
 description = "test2"
 members = [
   quay_organization_robot.test.fullname
 ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "test+test"),
				),
			},
			// Add team member
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.test.name
}

resource "quay_organization_robot" "test2" {
  name = "test2"
  orgname = quay_organization.test.name
}

resource "quay_organization_team" "test" {
  name = "test2"
  orgname = quay_organization.test.name
  role = "admin"
  description = "test2"
  members = [
    quay_organization_robot.test.fullname,
    quay_organization_robot.test2.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "test+test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.1", "test+test2"),
				),
			},
			// Remove team member by removing robot resource first (Quay will automatically remove the robot from the team)
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}

resource "quay_organization_robot" "test2" {
  name = "test2"
  orgname = quay_organization.test.name
}

resource "quay_organization_team" "test" {
  name = "test2"
  orgname = quay_organization.test.name
  role = "admin"
  description = "test2"
  members = [
    quay_organization_robot.test2.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "test+test2"),
				),
			},
			// Remove team member without removing robot resource
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}

resource "quay_organization_robot" "test2" {
  name = "test2"
  orgname = quay_organization.test.name
}

resource "quay_organization_team" "test" {
  name = "test2"
  orgname = quay_organization.test.name
  role = "admin"
  description = "test2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test2"),
				),
			},
		},
	})
}
