package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationRobotResource(t *testing.T) {
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
  description = "test"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_robot.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "description", "test"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_organization_robot.test",
				ImportState:                          true,
				ImportStateId:                        "test+test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			// Replace resource
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test2"
  orgname = quay_organization.test.name
  description = "test2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_robot.test", "name", "test2"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "orgname", "test"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "description", "test2"),
				),
			},
		},
	})
}
