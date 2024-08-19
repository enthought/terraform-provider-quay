package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrganizationDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and read resource
			{
				Config: providerConfig + `
resource "quay_organization" "org_data" {
  name = "org_data"
  email = "quay+org_data@example.com"
}

data "quay_organization" "org_data" {
  name = "org_data"

  depends_on = [
    quay_organization.org_data
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.quay_organization.org_data", "name", "org_data"),
					resource.TestCheckResourceAttr("data.quay_organization.org_data", "email", "quay+org_data@example.com"),
				),
			},
		},
	})
}
