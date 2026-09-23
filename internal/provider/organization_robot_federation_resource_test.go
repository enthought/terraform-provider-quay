package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationRobotFederationResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_fed" {
  name  = "org_fed"
  email = "quay+org_fed@example.com"
}

resource "quay_organization_robot" "fed" {
  name        = "fed"
  orgname     = quay_organization.org_fed.name
  description = "fed"
}

resource "quay_organization_robot_federation" "test" {
  orgname   = quay_organization.org_fed.name
  robotname = quay_organization_robot.fed.name
  federation = [
    {
      issuer  = "https://token.actions.githubusercontent.com"
      subject = "repo:example/example:ref:refs/heads/main"
    },
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "orgname", "org_fed"),
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "robotname", "fed"),
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "federation.#", "1"),
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "federation.0.issuer", "https://token.actions.githubusercontent.com"),
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "federation.0.subject", "repo:example/example:ref:refs/heads/main"),
				),
			},
			// Import
			{
				ResourceName:      "quay_organization_robot_federation.test",
				ImportState:       true,
				ImportStateId:     "org_fed/fed",
				ImportStateVerify: true,
			},
			// Update (replace the federation rules)
			{
				Config: providerConfig + `
resource "quay_organization" "org_fed" {
  name  = "org_fed"
  email = "quay+org_fed@example.com"
}

resource "quay_organization_robot" "fed" {
  name        = "fed"
  orgname     = quay_organization.org_fed.name
  description = "fed"
}

resource "quay_organization_robot_federation" "test" {
  orgname   = quay_organization.org_fed.name
  robotname = quay_organization_robot.fed.name
  federation = [
    {
      issuer  = "https://token.actions.githubusercontent.com"
      subject = "repo:example/example:ref:refs/heads/develop"
    },
    {
      issuer  = "https://accounts.google.com"
      subject = "1234567890"
    },
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "federation.#", "2"),
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "federation.0.subject", "repo:example/example:ref:refs/heads/develop"),
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "federation.1.issuer", "https://accounts.google.com"),
					resource.TestCheckResourceAttr("quay_organization_robot_federation.test", "federation.1.subject", "1234567890"),
				),
			},
		},
	})
}
