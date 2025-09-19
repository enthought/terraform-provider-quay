package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/resource_organization"
)

var (
	_ resource.Resource                = (*organizationResource)(nil)
	_ resource.ResourceWithConfigure   = (*organizationResource)(nil)
	_ resource.ResourceWithImportState = (*organizationResource)(nil)
)

func NewOrganizationResource() resource.Resource {
	return &organizationResource{}
}

type organizationResource struct {
	client *quay_api.APIClient
}

type organizationModelJSON struct {
	Email string                               `json:"email"`
	Name  string                               `json:"name"`
	Teams map[string]organizationTeamModelJSON `json:"teams"`
}

type organizationTeamModelJSON struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	CanView     bool   `json:"can_view"`
	RepoCount   int    `json:"repo_count"`
	MemberCount int    `json:"member_count"`
	IsSynced    bool   `json:"is_synced"`
}

func (r *organizationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (r *organizationResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_organization.OrganizationResourceSchema(ctx)
}

func (r *organizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_organization.OrganizationModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	newOrg := quay_api.NewNewOrg(data.Name.ValueString(), data.Email.ValueString())
	_, err := r.client.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay org", "Could not create Quay org, unexpected error: "+errDetail)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_organization.OrganizationModel
	var resData organizationModelJSON

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get data from API
	httpRes, err := r.client.OrganizationAPI.GetOrganization(context.Background(), data.Name.ValueString()).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error read Quay org", "Could not read Quay org, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay org",
			"Could not read Quay org, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay org",
			"Could not read Quay org, unexpected error: "+err.Error())
		return
	}

	// Set values from API
	data.Name = types.StringValue(resData.Name)
	data.Email = types.StringValue(resData.Email)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var dataState resource_organization.OrganizationModel
	var dataPlan resource_organization.OrganizationModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &dataState)...)

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &dataPlan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	// Update email first since the org path will change if the org name changes
	if dataPlan.Email != dataState.Email {
		updateOrg := quay_api.UpdateOrg{
			Name:  dataState.Name.ValueStringPointer(),
			Email: dataPlan.Email.ValueStringPointer(),
		}
		_, err := r.client.OrganizationAPI.ChangeOrganizationDetails(context.Background(), dataState.Name.ValueString()).Body(updateOrg).Execute()
		if err != nil {
			errDetail := handleQuayAPIError(err)
			resp.Diagnostics.AddError("Error updating Quay org", "Could not update Quay org, unexpected error: "+errDetail)
			return
		}
	}

	if dataPlan.Name != dataState.Name {
		updateOrg := quay_api.UpdateOrg{
			Name:  dataPlan.Name.ValueStringPointer(),
			Email: dataPlan.Email.ValueStringPointer(),
		}
		_, err := r.client.SuperuserAPI.ChangeOrganization(context.Background(), dataState.Name.ValueString()).Body(updateOrg).Execute()
		if err != nil {
			errDetail := handleQuayAPIError(err)
			resp.Diagnostics.AddError("Error updating Quay org", "Could not update Quay org, unexpected error: "+errDetail)
			return
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &dataPlan)...)
}

func (r *organizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_organization.OrganizationModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	_, err := r.client.OrganizationAPI.DeleteAdminedOrganization(context.Background(), data.Name.ValueString()).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay org", "Could not delete Quay org, unexpected error: "+errDetail)
		return
	}
}

func (r *organizationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *organizationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
