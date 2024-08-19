package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationRobotResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_robot" {
  name = "org_robot"
  email = "quay+org_robot@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_robot.name
  description = "test"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_robot.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "orgname", "org_robot"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "description", "test"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "fullname", "org_robot+test"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization_robot.test",
				ImportState:                          true,
				ImportStateId:                        "org_robot+test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			// Replace resource
			{
				Config: providerConfig + `
resource "quay_organization" "org_robot" {
  name = "org_robot"
  email = "quay+org_robot@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test2"
  orgname = quay_organization.org_robot.name
  description = "test2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_robot.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "orgname", "org_robot"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "description", "test2"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "fullname", "org_robot+test2"),
				),
			},
		},
	})
}
