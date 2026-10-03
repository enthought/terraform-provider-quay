package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRepositoryResourceStates(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create as readonly
			{
				Config: providerConfig + `
resource "quay_organization" "org_state" {
  name = "org_state"
  email = "quay+state@example.com"
}

resource "quay_repository" "test" {
  name = "state-test"
  namespace = quay_organization.org_state.name
  visibility = "private"
  description = "state test"
  state = "readonly"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.test", "state", "readonly"),
					resource.TestCheckNoResourceAttr("quay_repository.test", "mirror"),
				),
			},
			// Transition to normal
			{
				Config: providerConfig + `
resource "quay_organization" "org_state" {
  name = "org_state"
  email = "quay+state@example.com"
}

resource "quay_repository" "test" {
  name = "state-test"
  namespace = quay_organization.org_state.name
  visibility = "private"
  description = "state test"
  state = "normal"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.test", "state", "normal"),
				),
			},
			// Default state is normal
			{
				Config: providerConfig + `
resource "quay_organization" "org_state" {
  name = "org_state"
  email = "quay+state@example.com"
}

resource "quay_repository" "default" {
  name = "state-default"
  namespace = quay_organization.org_state.name
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.default", "state", "normal"),
				),
			},
		},
	})
}

func TestAccRepositoryResourceMirror(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// state=mirror without a mirror block is rejected at plan time
			{
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`mirror.*block.*state is "mirror"`),
				Config: providerConfig + `
resource "quay_organization" "org_mirror" {
  name = "org_mirror"
  email = "quay+mirror@example.com"
}

resource "quay_repository" "bad" {
  name = "mirror-test"
  namespace = quay_organization.org_mirror.name
  state = "mirror"
}
`,
			},
			// mirror block present with state=normal is rejected at plan time
			{
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`mirror.*block.*only valid when state is "mirror"`),
				Config: providerConfig + `
resource "quay_organization" "org_mirror" {
  name = "org_mirror"
  email = "quay+mirror@example.com"
}

resource "quay_repository" "bad" {
  name = "mirror-test"
  namespace = quay_organization.org_mirror.name
  state = "normal"

  mirror {
    upstream_repository = "quay.io/argoproj/argocd"
  }
}
`,
			},
			// a mirror block without required attributes is rejected at plan time
			{
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`upstream_repository.*is required when state is "mirror"`),
				Config: providerConfig + `
resource "quay_organization" "org_mirror" {
  name = "org_mirror"
  email = "quay+mirror@example.com"
}

resource "quay_repository" "bad" {
  name = "mirror-test"
  namespace = quay_organization.org_mirror.name
  state = "mirror"

  mirror {
    verify_tls = true
  }
}
`,
			},
			// Create a mirror repository
			{
				Config: providerConfig + `
resource "quay_organization" "org_mirror" {
  name = "org_mirror"
  email = "quay+mirror@example.com"
}

resource "quay_organization_robot" "mirror" {
  name = "mirror-bot"
  orgname = quay_organization.org_mirror.name
}

resource "quay_repository" "mirror" {
  name = "mirror-test"
  namespace = quay_organization.org_mirror.name
  visibility = "private"
  description = "mirror test"
  state = "mirror"

  mirror {
    upstream_repository = "quay.io/argoproj/argocd"
    robot_user = quay_organization_robot.mirror.fullname
    sync_interval = 86400
    sync_on_change = false
    verify_tls = true
    unsigned_images = false
    root_rule = { rule_kind = "tag_glob_csv", rule_value = ["v3.4.4"] }
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.mirror", "state", "mirror"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.sync_on_change", "false"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.upstream_repository", "quay.io/argoproj/argocd"),
					resource.TestCheckResourceAttrSet("quay_repository.mirror", "mirror.robot_user"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.sync_interval", "86400"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.verify_tls", "true"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.unsigned_images", "false"),
				),
			},
			// Import a mirrored repository
			{
				ResourceName:                         "quay_repository.mirror",
				ImportState:                          true,
				ImportStateId:                        "org_mirror/mirror-test",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "name",
				ImportStateVerifyIgnore: []string{
					"mirror.upstream_password",
					"mirror.sync_on_change",
				},
			},
			// Update mirror configuration and enable sync_on_change
			{
				Config: providerConfig + `
resource "quay_organization" "org_mirror" {
  name = "org_mirror"
  email = "quay+mirror@example.com"
}

resource "quay_organization_robot" "mirror" {
  name = "mirror-bot"
  orgname = quay_organization.org_mirror.name
}

resource "quay_repository" "mirror" {
  name = "mirror-test"
  namespace = quay_organization.org_mirror.name
  visibility = "private"
  description = "mirror test updated"
  state = "mirror"

  mirror {
    upstream_repository = "quay.io/argoproj/argocd"
    robot_user = quay_organization_robot.mirror.fullname
    sync_interval = 172800
    sync_on_change = true
    verify_tls = false
    unsigned_images = true
    root_rule = { rule_kind = "tag_glob_csv", rule_value = ["v3.4.4"] }
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.mirror", "state", "mirror"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.sync_on_change", "true"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.sync_interval", "172800"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.verify_tls", "false"),
					resource.TestCheckResourceAttr("quay_repository.mirror", "mirror.unsigned_images", "true"),
				),
			},
			// Transition out of mirror to normal (mirror config removed)
			{
				Config: providerConfig + `
resource "quay_organization" "org_mirror" {
  name = "org_mirror"
  email = "quay+mirror@example.com"
}

resource "quay_organization_robot" "mirror" {
  name = "mirror-bot"
  orgname = quay_organization.org_mirror.name
}

resource "quay_repository" "mirror" {
  name = "mirror-test"
  namespace = quay_organization.org_mirror.name
  visibility = "private"
  description = "mirror test updated"
  state = "normal"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.mirror", "state", "normal"),
					resource.TestCheckNoResourceAttr("quay_repository.mirror", "mirror"),
				),
			},
		},
	})
}

func TestAccRepositoryResourceMirrorValidation(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Invalid state value rejected at plan time
			{
				PlanOnly: true,
				Config: providerConfig + `
resource "quay_repository" "bad" {
  name = "bad-state"
  namespace = "someorg"
  state = "bogus"
}
`,
				ExpectError: regexp.MustCompile(".*value must be one of.*"),
			},
		},
	})
}

func TestAccRepositoryResourceMirrorBlockValidation(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// mirror block present without state=mirror should error (config consistency)
			{
				Config: providerConfig + `
resource "quay_repository" "bad" {
  name = "bad-mirror"
  namespace = "someorg"
  state = "normal"

  mirror {
    upstream_repository = "quay.io/argoproj/argocd"
  }
}
`,
				ExpectError: regexp.MustCompile(".*mirror.*block.*"),
			},
		},
	})
}

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
			// Update
			{
				Config: providerConfig + `
resource "quay_organization" "org_repo" {
  name = "org_repo"
  email = "quay+repo@example.com"
}

resource "quay_repository" "test" {
  name = "test"
  namespace = quay_organization.org_repo.name
  visibility = "public"
  description = "test2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.test", "name", "test"),
					resource.TestCheckResourceAttr("quay_repository.test", "namespace", "org_repo"),
					resource.TestCheckResourceAttr("quay_repository.test", "visibility", "public"),
					resource.TestCheckResourceAttr("quay_repository.test", "description", "test2"),
				),
			},
			// Replace resource
			{
				Config: providerConfig + `
resource "quay_organization" "org_repo" {
  name = "org_repo"
  email = "quay+repo@example.com"
}

resource "quay_repository" "test2" {
  name = "test2"
  namespace = quay_organization.org_repo.name
  visibility = "public"
  description = "test2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.test2", "name", "test2"),
					resource.TestCheckResourceAttr("quay_repository.test2", "namespace", "org_repo"),
					resource.TestCheckResourceAttr("quay_repository.test2", "visibility", "public"),
					resource.TestCheckResourceAttr("quay_repository.test2", "description", "test2"),
				),
			},
			// Create with only required values
			{
				Config: providerConfig + `
resource "quay_organization" "org_repo" {
  name = "org_repo"
  email = "quay+repo@example.com"
}

resource "quay_repository" "test3" {
  name = "test3"
  namespace = quay_organization.org_repo.name
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_repository.test3", "name", "test3"),
					resource.TestCheckResourceAttr("quay_repository.test3", "namespace", "org_repo"),
					resource.TestCheckResourceAttr("quay_repository.test3", "visibility", "private"),
					resource.TestCheckResourceAttr("quay_repository.test3", "description", ""),
				),
			},
		},
	})
}
