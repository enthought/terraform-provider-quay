package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccOrganizationPrototypeImportStateID(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		return rs.Primary.Attributes["orgname"] + "/" + rs.Primary.ID, nil
	}
}

func TestAccOrganizationPrototypeResource(t *testing.T) {
	org, orgRef, orgHCL := testAccOrg("org_prototype", "org_prototype")
	push := testAccName("push")
	pull := testAccName("pull")
	pull2 := testAccName("pulltwo")

	config := func(role, delegate string, withPull2 bool) string {
		c := providerConfig + orgHCL + fmt.Sprintf(`
resource "quay_organization_robot" "push" {
  name = %[1]q
  orgname = %[2]s
}

resource "quay_organization_robot" "pull" {
  name = %[3]q
  orgname = %[2]s
}
`, push, orgRef, pull)
		if withPull2 {
			c += fmt.Sprintf(`
resource "quay_organization_robot" "pull2" {
  name = %[1]q
  orgname = %[2]s
}
`, pull2, orgRef)
		}
		c += fmt.Sprintf(`
resource "quay_organization_prototype" "test" {
  orgname = %[1]s
  activating_user = quay_organization_robot.push.fullname
  delegate_name = quay_organization_robot.%[2]s.fullname
  delegate_kind = "user"
  role = %[3]q
}
`, orgRef, delegate, role)
		return c
	}

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: config("read", "pull", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "orgname", org),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "activating_user", org+"+"+push),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "delegate_name", org+"+"+pull),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "delegate_kind", "user"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "read"),
					resource.TestCheckResourceAttrSet("quay_organization_prototype.test", "id"),
				),
			},
			// Import
			{
				ResourceName:      "quay_organization_prototype.test",
				ImportState:       true,
				ImportStateIdFunc: testAccOrganizationPrototypeImportStateID("quay_organization_prototype.test"),
				ImportStateVerify: true,
			},
			// Update
			{
				Config: config("write", "pull", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "write"),
					resource.TestCheckResourceAttrSet("quay_organization_prototype.test", "id"),
				),
			},
			// Replace
			{
				Config: config("write", "pull2", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "delegate_name", org+"+"+pull2),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "write"),
				),
			},
		},
	})
}
