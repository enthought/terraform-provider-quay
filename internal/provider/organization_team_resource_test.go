package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationTeamResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_team" {
  name = "org_team"
  email = "quay+org_team@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_team.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team.name
  role = "member"
  description = "test"
  members = [
    quay_organization_robot.test.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "org_team+test"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization_team.test",
				ImportState:                          true,
				ImportStateId:                        "org_team+test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			// Update
			{
				Config: providerConfig + `
resource "quay_organization" "org_team" {
  name = "org_team"
  email = "quay+org_team@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_team.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team.name
  role = "admin"
  description = "test2"
  members = [
    quay_organization_robot.test.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "org_team+test"),
				),
			},
			// Replace resource
			{
				Config: providerConfig + `
resource "quay_organization" "org_team" {
  name = "org_team"
  email = "quay+org_team@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_team.name
}

resource "quay_organization_team" "test" {
  name = "test2"
  orgname = quay_organization.org_team.name
  role = "admin"
  description = "test2"
  members = [
    quay_organization_robot.test.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "admin"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "description", "test2"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "org_team+test"),
				),
			},
		},
	})
}

func TestAccOrganizationTeamResourceMembers(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create team with no members
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_members" {
  name = "org_team_members"
  email = "quay+org_team_members@example.com"
}
resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
  role = "member"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team_members"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
				),
			},
			// Set members to empty list
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_members" {
  name = "org_team_members"
  email = "quay+org_team_members@example.com"
}
resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
  role = "member"
  members = []
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team_members"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
				),
			},
			// Add team member
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_members" {
  name = "org_team_members"
  email = "quay+org_team_members@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
  role = "member"
  members = [
    quay_organization_robot.test.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team_members"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "org_team_members+test"),
				),
			},
			// Add team member
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_members" {
  name = "org_team_members"
  email = "quay+org_team_members@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
}

resource "quay_organization_robot" "test2" {
  name = "test2"
  orgname = quay_organization.org_team_members.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
  role = "member"
  members = [
    quay_organization_robot.test.fullname,
    quay_organization_robot.test2.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team_members"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "org_team_members+test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.1", "org_team_members+test2"),
				),
			},
			// Remove team member by removing robot resource first (Quay will automatically remove the robot from the team)
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_members" {
  name = "org_team_members"
  email = "quay+org_team_members@example.com"
}

resource "quay_organization_robot" "test2" {
  name = "test2"
  orgname = quay_organization.org_team_members.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
  role = "member"
  members = [
    quay_organization_robot.test2.fullname
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team_members"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "members.0", "org_team_members+test2"),
				),
			},
			// Remove team member without removing robot resource
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_members" {
  name = "org_team_members"
  email = "quay+org_team_members@example.com"
}

resource "quay_organization_robot" "test2" {
  name = "test2"
  orgname = quay_organization.org_team_members.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_members.name
  role = "member"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "orgname", "org_team_members"),
					resource.TestCheckResourceAttr("quay_organization_team.test", "role", "member"),
				),
			},
		},
	})
}
