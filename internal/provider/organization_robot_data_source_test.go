package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationRobotDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "quay_organization" "org_robot_data" {
  name = "org_robot_data"
  email = "quay+org_robot_data@example.com"
}

resource "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_robot_data.name
  description = "test"
}

data "quay_organization_robot" "test" {
  name = "test"
  orgname = quay_organization.org_robot_data.name
  
  depends_on = [
    quay_organization_robot.test
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_robot.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "orgname", "org_robot_data"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "description", "test"),
					resource.TestCheckResourceAttr("quay_organization_robot.test", "fullname", "org_robot_data+test"),
				),
			},
		},
	})
}
