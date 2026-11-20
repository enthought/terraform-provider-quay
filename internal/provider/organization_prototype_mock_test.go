package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccMockOrganizationPrototype_LifecycleAndImport(t *testing.T) {
	mock := newMockQuay(t)
	var firstID string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create and Read
			{
				Config: providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull"
  delegate_kind   = "user"
  role            = "read"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "orgname", "testorg"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "activating_user", "testorg+push"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "delegate_name", "testorg+pull"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "delegate_kind", "user"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "read"),
					resource.TestCheckResourceAttrSet("quay_organization_prototype.test", "id"),
					func(s *terraform.State) error {
						rs := s.RootModule().Resources["quay_organization_prototype.test"]
						firstID = rs.Primary.ID
						if firstID == "" {
							return fmt.Errorf("expected non-empty id")
						}
						return nil
					},
				),
			},
			// 2. Import
			{
				ResourceName:      "quay_organization_prototype.test",
				ImportState:       true,
				ImportStateIdFunc: testAccOrganizationPrototypeImportStateID("quay_organization_prototype.test"),
				ImportStateVerify: true,
			},
			// 3. Update role in place (ID must not change)
			{
				Config: providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull"
  delegate_kind   = "user"
  role            = "write"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "write"),
					func(s *terraform.State) error {
						rs := s.RootModule().Resources["quay_organization_prototype.test"]
						if rs.Primary.ID != firstID {
							return fmt.Errorf("expected id to remain %s, got %s", firstID, rs.Primary.ID)
						}
						return nil
					},
				),
			},
			// 4. Change delegate_name -> triggers replace (new ID)
			{
				Config: providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull2"
  delegate_kind   = "user"
  role            = "write"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "delegate_name", "testorg+pull2"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "write"),
					func(s *terraform.State) error {
						rs := s.RootModule().Resources["quay_organization_prototype.test"]
						if rs.Primary.ID == firstID {
							return fmt.Errorf("expected replacement with a new id, got same %s", rs.Primary.ID)
						}
						return nil
					},
				),
			},
			// 5. Change activating_user -> triggers replace (new ID)
			{
				Config: providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push2"
  delegate_name   = "testorg+pull2"
  delegate_kind   = "user"
  role            = "write"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "activating_user", "testorg+push2"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "delegate_name", "testorg+pull2"),
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "write"),
				),
			},
		},
	})
	_ = mock
}

func TestAccMockOrganizationPrototype_DriftAndRecreate(t *testing.T) {
	mock := newMockQuay(t)
	var protoID string

	testConfig := providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull"
  delegate_kind   = "user"
  role            = "read"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "read"),
					func(s *terraform.State) error {
						rs := s.RootModule().Resources["quay_organization_prototype.test"]
						protoID = rs.Primary.ID
						return nil
					},
				),
			},
			// 2. Drift: delete prototype in mock outside Terraform -> Plan shows non-empty recreate
			{
				PreConfig: func() {
					mock.DeletePrototype("testorg", protoID)
				},
				Config:             testConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// 3. Apply again to recreate
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_prototype.test", "role", "read"),
					resource.TestCheckResourceAttrSet("quay_organization_prototype.test", "id"),
				),
			},
		},
	})
}

func TestAccMockOrganizationPrototype_DeleteWhenAlreadyGone(t *testing.T) {
	mock := newMockQuay(t)
	var protoID string

	testConfig := providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull"
  delegate_kind   = "user"
  role            = "read"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create
			{
				Config: testConfig,
				Check: func(s *terraform.State) error {
					rs := s.RootModule().Resources["quay_organization_prototype.test"]
					protoID = rs.Primary.ID
					return nil
				},
			},
			// 2. Delete outside Terraform, then run Destroy -> succeeds (404 is ignored)
			{
				PreConfig: func() {
					mock.DeletePrototype("testorg", protoID)
				},
				Destroy: true,
				Config:  testConfig,
			},
		},
	})
}

func TestAccMockOrganizationPrototype_ValidationErrors(t *testing.T) {
	_ = newMockQuay(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Invalid role
			{
				PlanOnly: true,
				Config: providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull"
  delegate_kind   = "user"
  role            = "superadmin"
}
`,
				ExpectError: regexp.MustCompile("Attribute role value must be one of"),
			},
			// Invalid delegate_kind
			{
				PlanOnly: true,
				Config: providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull"
  delegate_kind   = "group"
  role            = "read"
}
`,
				ExpectError: regexp.MustCompile("Attribute delegate_kind value must be one of"),
			},
		},
	})
}

func TestAccMockOrganizationPrototype_ImportInvalidID(t *testing.T) {
	_ = newMockQuay(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "quay_organization_prototype" "test" {
  orgname         = "testorg"
  activating_user = "testorg+push"
  delegate_name   = "testorg+pull"
  delegate_kind   = "user"
  role            = "read"
}
`,
			},
			{
				ResourceName:  "quay_organization_prototype.test",
				ImportState:   true,
				ImportStateId: "invalid-import-id-no-slash",
				ExpectError:   regexp.MustCompile("Unexpected Import Identifier"),
			},
		},
	})
}
