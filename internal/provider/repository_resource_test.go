package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRepositoryResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_repo" {
  name = "org_repo"
  email = "quay+repo@example.com"
}

resource "quay_repository" "test" {
  name = "test"
  namespace = quay_organization.org_repo.name
  visibility = "private"
  description = "test"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_repository.test", "namespace", "org_repo"),
					resource.TestCheckResourceAttr("quay_repository.test", "visibility", "private"),
					resource.TestCheckResourceAttr("quay_repository.test", "description", "test"),
				),
			},
			// Import
			{
				ResourceName:                         "quay_repository.test",
				ImportState:                          true,
				ImportStateId:                        "org_repo/test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
		},
	})
}
