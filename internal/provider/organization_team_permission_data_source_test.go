package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationTeamPermissionDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_perm_data" {
  name = "org_team_perm_data"
  email = "quay+org_team_perm_data@example.com"
}

resource "quay_repository" "test" {
  name = "test"
  namespace = quay_organization.org_team_perm_data.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_perm_data.name
  role = "member"
}

resource "quay_organization_team_permission" "test" {
  orgname = quay_organization.org_team_perm_data.name
  reponame = quay_repository.test.name
  teamname = quay_organization_team.test.name
  permission = "read"
}

data "quay_organization_team_permission" "test" {
  orgname = quay_organization.org_team_perm_data.name
  reponame = quay_repository.test.name
  teamname = quay_organization_team.test.name
  
  depends_on = [
    "quay_organization_team_permission.test"
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.quay_organization_team_permission.test", "orgname", "org_team_perm_data"),
					resource.TestCheckResourceAttr("data.quay_organization_team_permission.test", "reponame", "test"),
					resource.TestCheckResourceAttr("data.quay_organization_team_permission.test", "teamname", "test"),
					resource.TestCheckResourceAttr("data.quay_organization_team_permission.test", "permission", "read"),
				),
			},
		},
	})
}
