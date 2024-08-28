package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationTeamDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_data" {
  name = "org_team_data"
  email = "quay+org_team_data@example.com"
}

resource "quay_repository" "test" {
  name = "test"
  namespace = quay_organization.org_team_data.name
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_team_data.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_data.name
  role = "member"
  description = "test"
  members = [
    quay_organization_robot.test.fullname
  ]
}

data "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_data.name
  
  depends_on = [
    quay_organization_team.test
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.quay_organization_team.test", "name", "test"),
					resource.TestCheckResourceAttr("data.quay_organization_team.test", "orgname", "org_team_data"),
					resource.TestCheckResourceAttr("data.quay_organization_team.test", "role", "member"),
					resource.TestCheckResourceAttr("data.quay_organization_team.test", "description", "test"),
					resource.TestCheckResourceAttr("data.quay_organization_team.test", "members.0", "org_team_data+test"),
				),
			},
		},
	})
}
