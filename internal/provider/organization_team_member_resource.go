// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Modifications copyright (c) Enthought, Inc.
// SPDX-License-Identifier:	BSD-3-Clause

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/resource_organization_team_member"
)

var (
	_ resource.Resource                = (*organizationTeamMemberResource)(nil)
	_ resource.ResourceWithConfigure   = (*organizationTeamMemberResource)(nil)
	_ resource.ResourceWithImportState = (*organizationTeamMemberResource)(nil)
)

func NewOrganizationTeamMemberResource() resource.Resource {
	return &organizationTeamMemberResource{}
}

type organizationTeamMemberResource struct {
	client *quay_api.APIClient
}

func teamMemberID(orgName, teamName, memberName string) string {
	return orgName + "/" + teamName + "/" + memberName
}

func (r *organizationTeamMemberResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_team_member"
}

func (r *organizationTeamMemberResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_organization_team_member.OrganizationTeamMemberResourceSchema(ctx)
}

func (r *organizationTeamMemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_organization_team_member.OrganizationTeamMemberModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	teamName := data.Teamname.ValueString()
	memberName := data.Membername.ValueString()

	// Add team member
	_, err := r.client.TeamAPI.UpdateOrganizationTeamMember(context.Background(), orgName, memberName, teamName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay team member", "Could not create Quay team member, unexpected error: "+errDetail)
		return
	}

	data.Id = types.StringValue(teamMemberID(orgName, teamName, memberName))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationTeamMemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_organization_team_member.OrganizationTeamMemberModel
	var resTeamData teamModelJSON

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	teamName := data.Teamname.ValueString()
	memberName := data.Membername.ValueString()

	// Get members
	httpRes, err := r.client.TeamAPI.GetOrganizationTeamMembers(context.Background(), orgName, teamName).Execute()
	if err != nil {
		// Team was deleted outside of Terraform
		if isQuayNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay team member", "Could not read Quay team member, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team member",
			"Could not read Quay team member, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resTeamData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team member",
			"Could not read Quay team member, unexpected error: "+err.Error())
		return
	}

	// Check that the member is present
	for _, member := range resTeamData.Members {
		if member.Name == memberName {
			data.Id = types.StringValue(teamMemberID(orgName, teamName, memberName))

			// Save updated data into Terraform state
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	// Member was removed outside of Terraform
	resp.State.RemoveResource(ctx)
}

func (r *organizationTeamMemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_organization_team_member.OrganizationTeamMemberModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// All attributes require replacement, so there is nothing to update
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationTeamMemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_organization_team_member.OrganizationTeamMemberModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	teamName := data.Teamname.ValueString()
	memberName := data.Membername.ValueString()

	// Remove team member
	_, err := r.client.TeamAPI.DeleteOrganizationTeamMember(context.Background(), orgName, memberName, teamName).Execute()
	if err != nil {
		if isQuayNotFound(err) {
			return
		}
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay team member", "Could not delete Quay team member, unexpected error: "+errDetail)
		return
	}
}

func (r *organizationTeamMemberResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*quay_api.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *quay_api.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
	}

	r.client = client
}

func (r *organizationTeamMemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idSplit := strings.Split(req.ID, "/")
	if len(idSplit) != 3 || idSplit[0] == "" || idSplit[1] == "" || idSplit[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format orgname/teamname/membername. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("orgname"), idSplit[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("teamname"), idSplit[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("membername"), idSplit[2])...)
}
