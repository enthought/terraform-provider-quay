package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationTeamMemberResource(t *testing.T) {
	org, orgRef, orgHCL := testAccOrg("org_team_member", "org_team_member")
	team := testAccName("team")
	robot := testAccName("one")
	robot2 := testAccName("two")

	config := func(member string) string {
		return providerConfig + orgHCL + fmt.Sprintf(`
resource "quay_organization_team" "test" {
  name = %[1]q
  orgname = %[2]s
  role = "member"

  # quay_organization_team_member adds members outside this resource.
  lifecycle {
    ignore_changes = [members]
  }
}

resource "quay_organization_robot" "test" {
  name = %[3]q
  orgname = %[2]s
}

resource "quay_organization_robot" "test2" {
  name = %[4]q
  orgname = %[2]s
}

resource "quay_organization_team_member" "test" {
  orgname = %[2]s
  teamname = quay_organization_team.test.name
  membername = quay_organization_robot.%[5]s.fullname
}
`, team, orgRef, robot, robot2, member)
	}

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: config("test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "orgname", org),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "teamname", team),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "membername", org+"+"+robot),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "id", org+"/"+team+"/"+org+"+"+robot),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization_team_member.test",
				ImportState:                          true,
				ImportStateId:                        org + "/" + team + "/" + org + "+" + robot,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "id",
			},
			// Replace
			{
				Config: config("test2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "membername", org+"+"+robot2),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "id", org+"/"+team+"/"+org+"+"+robot2),
				),
			},
		},
	})
}
