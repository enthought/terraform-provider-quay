package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationRobotFederationDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "quay_organization" "org_fed_data" {
  name  = "org_fed_data"
  email = "quay+org_fed_data@example.com"
}

resource "quay_organization_robot" "fed_data" {
  name        = "fed_data"
  orgname     = quay_organization.org_fed_data.name
  description = "fed_data"
}

resource "quay_organization_robot_federation" "test" {
  orgname   = quay_organization.org_fed_data.name
  robotname = quay_organization_robot.fed_data.name
  federation = [
    {
      issuer  = "https://token.actions.githubusercontent.com"
      subject = "repo:example/example:ref:refs/heads/main"
    },
  ]
}

data "quay_organization_robot_federation" "test" {
  orgname   = quay_organization.org_fed_data.name
  robotname = quay_organization_robot.fed_data.name

  depends_on = [
    quay_organization_robot_federation.test
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.quay_organization_robot_federation.test", "orgname", "org_fed_data"),
					resource.TestCheckResourceAttr("data.quay_organization_robot_federation.test", "robotname", "fed_data"),
					resource.TestCheckResourceAttr("data.quay_organization_robot_federation.test", "federation.#", "1"),
					resource.TestCheckResourceAttr("data.quay_organization_robot_federation.test", "federation.0.issuer", "https://token.actions.githubusercontent.com"),
					resource.TestCheckResourceAttr("data.quay_organization_robot_federation.test", "federation.0.subject", "repo:example/example:ref:refs/heads/main"),
				),
			},
		},
	})
}
