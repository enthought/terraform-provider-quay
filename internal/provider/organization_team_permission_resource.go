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
	"terraform-provider-quay/internal/resource_organization_team_permission"
)

var (
	_ resource.Resource                = (*organizationTeamPermissionResource)(nil)
	_ resource.ResourceWithConfigure   = (*organizationTeamPermissionResource)(nil)
	_ resource.ResourceWithImportState = (*organizationTeamPermissionResource)(nil)
)

func NewOrganizationTeamPermissionResource() resource.Resource {
	return &organizationTeamPermissionResource{}
}

type organizationTeamPermissionResource struct {
	client *quay_api.APIClient
}

type teamPermissionModelJSON struct {
	Permission string `json:"role"`
}

func (r *organizationTeamPermissionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_team_permission"
}

func (r *organizationTeamPermissionResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_organization_team_permission.OrganizationTeamPermissionResourceSchema(ctx)
}

func (r *organizationTeamPermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_organization_team_permission.OrganizationTeamPermissionModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	repoName := data.Reponame.ValueString()
	teamName := data.Teamname.ValueString()
	permission := data.Permission.ValueString()

	// Create team permission
	newTeamPermission := quay_api.NewTeamPermission(permission)
	_, err := r.client.PermissionAPI.ChangeTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Body(*newTeamPermission).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError(
			"Error creating Quay team permission",
			"Could not create Quay team permission, unexpected error: "+errDetail)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationTeamPermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_organization_team_permission.OrganizationTeamPermissionModel
	var resData teamPermissionModelJSON

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	repoName := data.Reponame.ValueString()
	teamName := data.Teamname.ValueString()

	// Get team permission
	httpRes, err := r.client.PermissionAPI.GetTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError(
			"Error reading Quay team permission",
			"Could not read Quay team permission, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team permission",
			"Could not read Quay team permission, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team permission",
			"Could not read Quay team permission, unexpected error: "+err.Error())
		return
	}

	// Set permission
	data.Permission = types.StringValue(resData.Permission)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationTeamPermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var dataState resource_organization_team_permission.OrganizationTeamPermissionModel
	var dataPlan resource_organization_team_permission.OrganizationTeamPermissionModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &dataState)...)

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &dataPlan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := dataPlan.Orgname.ValueString()
	repoName := dataPlan.Reponame.ValueString()
	teamName := dataPlan.Teamname.ValueString()
	permission := dataPlan.Permission.ValueString()

	// Update team permission
	if dataPlan.Permission != dataState.Permission {
		newTeamPermission := quay_api.NewTeamPermission(permission)
		_, err := r.client.PermissionAPI.ChangeTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Body(*newTeamPermission).Execute()
		if err != nil {
			errDetail := handleQuayAPIError(err)
			resp.Diagnostics.AddError(
				"Error updating Quay team permission",
				"Could not update Quay team permission, unexpected error: "+errDetail)
			return
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &dataPlan)...)
}

func (r *organizationTeamPermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_organization_team_permission.OrganizationTeamPermissionModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	repoName := data.Reponame.ValueString()
	teamName := data.Teamname.ValueString()

	// Delete team permission
	_, err := r.client.PermissionAPI.DeleteTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError(
			"Error deleting Quay team permission",
			"Could not delete Quay team permission, unexpected error: "+errDetail)
		return
	}
}

func (r *organizationTeamPermissionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *organizationTeamPermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idSplit := strings.Split(req.ID, " ")
	if len(idSplit) != 3 || idSplit[0] == "" || idSplit[1] == "" || idSplit[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format orgname reponame teamname. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("orgname"), idSplit[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("reponame"), idSplit[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("teamname"), idSplit[2])...)
}
