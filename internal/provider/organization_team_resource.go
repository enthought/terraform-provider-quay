// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Modifications copyright (c) Enthought, Inc.
// SPDX-License-Identifier:	BSD-3-Clause

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/resource_organization_team"
)

var (
	_ resource.Resource                = (*organizationTeamResource)(nil)
	_ resource.ResourceWithConfigure   = (*organizationTeamResource)(nil)
	_ resource.ResourceWithImportState = (*organizationTeamResource)(nil)
)

func NewOrganizationTeamResource() resource.Resource {
	return &organizationTeamResource{}
}

type organizationTeamResource struct {
	client *quay_api.APIClient
}

type teamModelJSON struct {
	Name    string                `json:"name"`
	Members []teamMemberModelJSON `json:"members"`
	CanEdit bool                  `json:"can_edit"`
}

type teamMemberModelJSON struct {
	Name    string `json:"name"`
	Kind    string `json:"kind"`
	IsRobot bool   `json:"is_robot"`
	Invited bool   `json:"invited"`
}

func (r *organizationTeamResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_team"
}

func (r *organizationTeamResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_organization_team.OrganizationTeamResourceSchema(ctx)
}

func (r *organizationTeamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_organization_team.OrganizationTeamModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	teamName := data.Name.ValueString()
	orgName := data.Orgname.ValueString()
	role := data.Role.ValueString()
	description := data.Description.ValueString()

	// Create team
	newTeam := quay_api.TeamDescription{
		Role:        role,
		Description: &description,
	}
	_, err := r.client.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay team", "Could not create Quay team, unexpected error: "+errDetail)
		return
	}

	// Add team members
	if !data.Members.IsNull() {
		elements := make([]types.String, 0, len(data.Members.Elements()))
		diags := data.Members.ElementsAs(ctx, &elements, false)

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		for _, member := range elements {
			_, err = r.client.TeamAPI.UpdateOrganizationTeamMember(context.Background(), orgName, member.ValueString(), teamName).Execute()
			if err != nil {
				errDetail := handleQuayAPIError(err)
				resp.Diagnostics.AddError("Error creating Quay team", "Could not create Quay team, unexpected error: "+errDetail)
				return
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationTeamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_organization_team.OrganizationTeamModel
	var resTeamData teamModelJSON
	var resOrgData organizationModelJSON

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	teamName := data.Name.ValueString()
	orgName := data.Orgname.ValueString()

	// Get members
	httpRes, err := r.client.TeamAPI.GetOrganizationTeamMembers(context.Background(), orgName, teamName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay team", "Could not read Quay team, unexpected error: "+errDetail)
		return
	}

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team",
			"Could not read Quay team, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resTeamData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team",
			"Could not read Quay team, unexpected error: "+err.Error())
		return
	}

	// Set members
	var memList []string
	for _, member := range resTeamData.Members {
		memList = append(memList, member.Name)
	}
	members, diags := types.ListValueFrom(ctx, types.StringType, memList)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// handle case where `members` is set to []
	if len(data.Members.Elements()) > 0 || len(members.Elements()) > 0 {
		data.Members = members
	}

	// Get org data
	httpRes, err = r.client.OrganizationAPI.GetOrganization(context.Background(), orgName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay team", "Could not read Quay team, unexpected error: "+errDetail)
		return
	}

	body, err = io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team",
			"Could not read Quay team, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resOrgData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team",
			"Could not read Quay team, unexpected error: "+err.Error())
		return
	}

	// Set Role
	data.Role = types.StringValue(resOrgData.Teams[teamName].Role)

	// Set Description
	data.Description = types.StringValue(resOrgData.Teams[teamName].Description)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationTeamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var dataState resource_organization_team.OrganizationTeamModel
	var dataPlan resource_organization_team.OrganizationTeamModel
	var apiErr *quay_api.GenericOpenAPIError

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &dataState)...)

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &dataPlan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	teamName := dataPlan.Name.ValueString()
	orgName := dataPlan.Orgname.ValueString()

	// Update members
	if !dataPlan.Members.Equal(dataState.Members) {
		elementsPlan := make([]string, 0, len(dataPlan.Members.Elements()))
		diags := dataPlan.Members.ElementsAs(ctx, &elementsPlan, false)

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		elementsState := make([]string, 0, len(dataState.Members.Elements()))
		diags = dataState.Members.ElementsAs(ctx, &elementsState, false)

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// add team members
		for _, member := range subtractStringSlice(elementsPlan, elementsState) {
			_, err := r.client.TeamAPI.UpdateOrganizationTeamMember(context.Background(), orgName, member, teamName).Execute()
			if err != nil {
				errDetail := handleQuayAPIError(err)
				resp.Diagnostics.AddError("Error updating Quay team", "Could not update Quay team, unexpected error: "+errDetail)
				return
			}
		}

		// remove team members
		for _, member := range subtractStringSlice(elementsState, elementsPlan) {
			_, err := r.client.TeamAPI.DeleteOrganizationTeamMember(context.Background(), orgName, member, teamName).Execute()
			if err != nil {
				// handle case where team member no longer exists
				if errors.As(err, &apiErr) {
					if model, ok := apiErr.Model().(quay_api.ApiError); ok {
						if model.Status == 404 {
							continue
						}
					}
				}

				errDetail := handleQuayAPIError(err)
				resp.Diagnostics.AddError("Error updating Quay team", "Could not update Quay team, unexpected error: "+errDetail)
				return
			}
		}
	}

	// Update role and description
	updateTeam := quay_api.TeamDescription{
		Role:        dataPlan.Role.ValueString(),
		Description: dataPlan.Description.ValueStringPointer(),
	}
	_, err := r.client.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(updateTeam).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay team", "Could not create Quay team, unexpected error: "+errDetail)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &dataPlan)...)
}

func (r *organizationTeamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_organization_team.OrganizationTeamModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	teamName := data.Name.ValueString()
	orgName := data.Orgname.ValueString()

	// Delete API call logic
	_, err := r.client.TeamAPI.DeleteOrganizationTeam(context.Background(), orgName, teamName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay team", "Could not delete Quay team, unexpected error: "+errDetail)
		return
	}
}

func (r *organizationTeamResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *organizationTeamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idSplit := strings.Split(req.ID, "+")
	if len(idSplit) != 2 || idSplit[0] == "" || idSplit[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format orgname+teamname. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("orgname"), idSplit[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idSplit[1])...)
}
