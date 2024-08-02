package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and read resource
			{
				Config: providerConfig + `
resource "quay_organization" "test" {
  name = "test"
  email = "quay+test@example.com"
}

data "quay_organization" "test" {
  name = "test"

  depends_on = [
    quay_organization.test
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.quay_organization.test", "name", "test"),
					resource.TestCheckResourceAttr("data.quay_organization.test", "email", "quay+test@example.com"),
				),
			},
		},
	})
}
