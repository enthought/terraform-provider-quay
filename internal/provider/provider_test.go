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
	t.Setenv("QUAY_OAUTH2_USERNAME", "")
	t.Setenv("QUAY_OAUTH2_PASSWORD", "")
	t.Setenv("QUAY_OAUTH2_CLIENT_ID", "")
	t.Setenv("QUAY_OAUTH2_TOKEN_URL", "")

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
			// Token and OAuth2 credentials
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  token = "asdf"
  oauth2_client_id = "quay"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Cannot specify token and OAuth2 credentials"),
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
			// Missing Quay token and OAuth2 credentials
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
				ExpectError: regexp.MustCompile(".*Error: Missing Quay token and OAuth2 credentials"),
			},
			// Missing OAuth2 username
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  oauth2_password = "test"
  oauth2_client_id = "quay"
  oauth2_token_url = "https://auth.example.com/token"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing OAuth2 username"),
			},
			// Missing OAuth2 password
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  oauth2_username = "test"
  oauth2_client_id = "quay"
  oauth2_token_url = "https://auth.example.com/token"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing OAuth2 password"),
			},
			// Missing OAuth2 client ID
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  oauth2_username = "test"
  oauth2_password = "test"
  oauth2_token_url = "https://auth.example.com/token"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing OAuth2 client ID"),
			},
			// Missing OAuth2 token URL
			{
				PlanOnly: true,
				Config: `
provider "quay" {
  url = "https://quay.example.com"
  oauth2_username = "test"
  oauth2_password = "test"
  oauth2_client_id = "quay"
}

data "quay_organization" "provider_test" {
  name = "provider_test"
}
`,
				ExpectError: regexp.MustCompile(".*Error: Missing OAuth2 token URL"),
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
