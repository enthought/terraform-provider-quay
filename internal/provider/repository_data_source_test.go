package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRepositoryDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig + `
resource "quay_organization" "org_repo_data" {
  name = "org_repo_data"
  email = "quay+org_repo_data@example.com"
}

resource "quay_repository" "test" {
  name = "test"
  namespace = quay_organization.org_repo_data.name
  visibility = "private"
  description = "test"
}

data "quay_repository" "test" {
  name = "test"
  namespace = quay_organization.org_repo_data.name
  
  depends_on = [
    quay_repository.test
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.quay_repository.test", "name", "test"),
					resource.TestCheckResourceAttr("data.quay_repository.test", "namespace", "org_repo_data"),
					resource.TestCheckResourceAttr("data.quay_repository.test", "visibility", "private"),
					resource.TestCheckResourceAttr("data.quay_repository.test", "description", "test"),
				),
			},
		},
	})
}
