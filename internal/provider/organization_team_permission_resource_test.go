package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationTeamPermissionResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_team_perm" {
  name = "org_team_perm"
  email = "quay+org_team@example.com"
}

resource "quay_repository" "test" {
  name = "test"
  namespace = quay_organization.org_team_perm.name
}

resource "quay_organization_team" "test" {
  name = "test"
  orgname = quay_organization.org_team_perm.name
  role = "member"
}

resource "quay_organization_team_permission" "test" {
  orgname = quay_organization.org_team_perm.name
  reponame = quay_repository.test.name
  teamname = quay_organization_team.test.name
  permission = "read"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_permission.test", "orgname", "org_team_perm"),
					resource.TestCheckResourceAttr("quay_organization_team_permission.test", "reponame", "test"),
					resource.TestCheckResourceAttr("quay_organization_team_permission.test", "teamname", "test"),
					resource.TestCheckResourceAttr("quay_organization_team_permission.test", "permission", "read"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization_team_permission.test",
				ImportState:                          true,
				ImportStateId:                        "org_team_perm test test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "teamname",
			},
		},
	})
}
