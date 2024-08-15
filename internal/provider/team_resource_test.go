package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Required parameters only
func TestAccTeamResourceRequired(t *testing.T) {
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
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_team.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_team.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_team.test", "role", "member"),
				),
			},
		},
	})
}
