package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	providerConfig = `
provider "quay" {}
`
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"quay": providerserver.NewProtocol6WithError(New()()),
	}
)

func TestAccProviderValidationWithoutENV(t *testing.T) {
	// unset QUAY_* environment variables
	t.Setenv("QUAY_URL", "")
	t.Setenv("QUAY_TOKEN", "")
	t.Setenv("QUAY_CLIENT_ID", "")
	t.Setenv("QUAY_CLIENT_SECRET", "")

	resource.Test(t, resource.TestCase{
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: "~> 3.6",
			},
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unknown config values
			{
				PlanOnly: true,
				Config: `
resource "random_string" "this" {
  length = 8
}

provider "quay" {
  url = "https://quay-${random_string.this}.example.com"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Unknown configuration values"),
			},
			// Token and client credentials
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  token = "asdf"
  client_id = "quay"
  client_secret = "asdf"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Cannot specify token and client credentials"),
			},
			// No config values
			{
				PlanOnly: true,
				Config: `
provider "quay" {}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing Quay URL"),
			},
			// Invalid URL
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "quay.example.com"
  token = "asdf"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Quay URL is not a valid URL"),
			},
			// Missing Quay token and client credentials
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing Quay token and client credentials"),
			},
			// Missing client secret
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  client_id = "quay"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing client secret"),
			},
			// Missing client ID
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  client_secret = "asdf"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing client ID"),
			},
		},
	})
}

func TestAccProviderValidationWithEnv(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// No config values
			{
				PlanOnly: true,
				Config: `
provider "quay" {}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Error reading Quay org"),
			},
		},
	})
}
