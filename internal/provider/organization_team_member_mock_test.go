package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccMockOrganizationTeamMember_LifecycleAndImport(t *testing.T) {
	mock := newMockQuay(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create and Read
			{
				Config: providerConfig + `
resource "quay_organization_team_member" "test" {
  orgname    = "testorg"
  teamname   = "devteam"
  membername = "testorg+robot1"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "orgname", "testorg"),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "teamname", "devteam"),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "membername", "testorg+robot1"),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "id", "testorg/devteam/testorg+robot1"),
					func(_ *terraform.State) error {
						if !mock.HasTeamMember("testorg", "devteam", "testorg+robot1") {
							return fmt.Errorf("expected testorg+robot1 to be in mock devteam")
						}
						return nil
					},
				),
			},
			// 2. Import
			{
				ResourceName:                         "quay_organization_team_member.test",
				ImportState:                          true,
				ImportStateId:                        "testorg/devteam/testorg+robot1",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "id",
			},
			// 3. Replace when membername changes -> old member removed, new member added
			{
				Config: providerConfig + `
resource "quay_organization_team_member" "test" {
  orgname    = "testorg"
  teamname   = "devteam"
  membername = "testorg+robot2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "membername", "testorg+robot2"),
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "id", "testorg/devteam/testorg+robot2"),
					func(_ *terraform.State) error {
						if mock.HasTeamMember("testorg", "devteam", "testorg+robot1") {
							return fmt.Errorf("expected old member testorg+robot1 to be removed from mock devteam")
						}
						if !mock.HasTeamMember("testorg", "devteam", "testorg+robot2") {
							return fmt.Errorf("expected new member testorg+robot2 to be in mock devteam")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccMockOrganizationTeamMember_SharedTeamPreserved(t *testing.T) {
	mock := newMockQuay(t)

	// Pre-seed unrelated members in the team outside Terraform
	mock.AddTeamMember("testorg", "sharedteam", "existing_lead")
	mock.AddTeamMember("testorg", "sharedteam", "existing_robot")

	testConfig := providerConfig + `
resource "quay_organization_team_member" "new_member" {
  orgname    = "testorg"
  teamname   = "sharedteam"
  membername = "testorg+new_robot"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create: new member is added, existing members remain intact
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.new_member", "membername", "testorg+new_robot"),
					func(_ *terraform.State) error {
						members := mock.GetTeamMembers("testorg", "sharedteam")
						expected := map[string]bool{
							"existing_lead":     true,
							"existing_robot":    true,
							"testorg+new_robot": true,
						}
						if len(members) != 3 {
							return fmt.Errorf("expected 3 members in sharedteam, got %d: %v", len(members), members)
						}
						for _, m := range members {
							if !expected[m] {
								return fmt.Errorf("unexpected member %s in sharedteam", m)
							}
						}
						return nil
					},
				),
			},
			// 2. Destroy: only managed member is removed, pre-existing members stay untouched
			{
				Config:  testConfig,
				Destroy: true,
			},
		},
	})

	// Post-destroy verification: pre-existing members still present, managed member removed
	if !mock.HasTeamMember("testorg", "sharedteam", "existing_lead") {
		t.Errorf("expected existing_lead to remain in sharedteam after destroy")
	}
	if !mock.HasTeamMember("testorg", "sharedteam", "existing_robot") {
		t.Errorf("expected existing_robot to remain in sharedteam after destroy")
	}
	if mock.HasTeamMember("testorg", "sharedteam", "testorg+new_robot") {
		t.Errorf("expected testorg+new_robot to be deleted from sharedteam after destroy")
	}
}

func TestAccMockOrganizationTeamMember_MultipleMembersSameTeam(t *testing.T) {
	mock := newMockQuay(t)

	twoMembersConfig := providerConfig + `
resource "quay_organization_team_member" "member_a" {
  orgname    = "testorg"
  teamname   = "sharedteam"
  membername = "testorg+robot_a"
}

resource "quay_organization_team_member" "member_b" {
  orgname    = "testorg"
  teamname   = "sharedteam"
  membername = "testorg+robot_b"
}
`

	oneMemberConfig := providerConfig + `
resource "quay_organization_team_member" "member_a" {
  orgname    = "testorg"
  teamname   = "sharedteam"
  membername = "testorg+robot_a"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create two members in the same team
			{
				Config: twoMembersConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.member_a", "membername", "testorg+robot_a"),
					resource.TestCheckResourceAttr("quay_organization_team_member.member_b", "membername", "testorg+robot_b"),
					func(_ *terraform.State) error {
						if !mock.HasTeamMember("testorg", "sharedteam", "testorg+robot_a") {
							return fmt.Errorf("expected robot_a in team")
						}
						if !mock.HasTeamMember("testorg", "sharedteam", "testorg+robot_b") {
							return fmt.Errorf("expected robot_b in team")
						}
						return nil
					},
				),
			},
			// 2. Remove member_b from configuration -> member_b removed, member_a remains
			{
				Config: oneMemberConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.member_a", "membername", "testorg+robot_a"),
					func(_ *terraform.State) error {
						if !mock.HasTeamMember("testorg", "sharedteam", "testorg+robot_a") {
							return fmt.Errorf("expected robot_a to remain in team")
						}
						if mock.HasTeamMember("testorg", "sharedteam", "testorg+robot_b") {
							return fmt.Errorf("expected robot_b to be removed from team")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccMockOrganizationTeamMember_DriftMemberRemoved(t *testing.T) {
	mock := newMockQuay(t)

	testConfig := providerConfig + `
resource "quay_organization_team_member" "test" {
  orgname    = "testorg"
  teamname   = "devteam"
  membername = "testorg+robot1"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "membername", "testorg+robot1"),
				),
			},
			// 2. Drift: remove member in mock -> plan shows recreation
			{
				PreConfig: func() {
					mock.RemoveTeamMember("testorg", "devteam", "testorg+robot1")
				},
				Config:             testConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// 3. Apply again -> recreated
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "membername", "testorg+robot1"),
					func(_ *terraform.State) error {
						if !mock.HasTeamMember("testorg", "devteam", "testorg+robot1") {
							return fmt.Errorf("expected testorg+robot1 to be recreated in mock devteam")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccMockOrganizationTeamMember_DriftTeamDeleted(t *testing.T) {
	mock := newMockQuay(t)

	testConfig := providerConfig + `
resource "quay_organization_team_member" "test" {
  orgname    = "testorg"
  teamname   = "ephemeral_team"
  membername = "testorg+robot1"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("quay_organization_team_member.test", "membername", "testorg+robot1"),
				),
			},
			// 2. Drift: entire team deleted in mock (404 on GET members) -> Read removes from state gracefully, plan shows create
			{
				PreConfig: func() {
					mock.DeleteTeam("testorg", "ephemeral_team")
				},
				Config:             testConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMockOrganizationTeamMember_DeleteWhenAlreadyGone(t *testing.T) {
	mock := newMockQuay(t)

	testConfig := providerConfig + `
resource "quay_organization_team_member" "test" {
  orgname    = "testorg"
  teamname   = "devteam"
  membername = "testorg+robot1"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 1. Create
			{
				Config: testConfig,
			},
			// 2. Delete member outside Terraform, then run Destroy -> succeeds (404 is ignored)
			{
				PreConfig: func() {
					mock.RemoveTeamMember("testorg", "devteam", "testorg+robot1")
				},
				Destroy: true,
				Config:  testConfig,
			},
		},
	})
}

func TestAccMockOrganizationTeamMember_ImportInvalidID(t *testing.T) {
	_ = newMockQuay(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "quay_organization_team_member" "test" {
  orgname    = "testorg"
  teamname   = "devteam"
  membername = "testorg+robot1"
}
`,
			},
			// Invalid ID: only 2 parts
			{
				ResourceName:  "quay_organization_team_member.test",
				ImportState:   true,
				ImportStateId: "testorg/devteam",
				ExpectError:   regexp.MustCompile("Unexpected Import Identifier"),
			},
			// Invalid ID: 4 parts
			{
				ResourceName:  "quay_organization_team_member.test",
				ImportState:   true,
				ImportStateId: "testorg/devteam/robot1/extra",
				ExpectError:   regexp.MustCompile("Unexpected Import Identifier"),
			},
		},
	})
}
