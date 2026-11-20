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
	"terraform-provider-quay/internal/resource_organization_prototype"
)

var (
	_ resource.Resource                = (*organizationPrototypeResource)(nil)
	_ resource.ResourceWithConfigure   = (*organizationPrototypeResource)(nil)
	_ resource.ResourceWithImportState = (*organizationPrototypeResource)(nil)
)

func NewOrganizationPrototypeResource() resource.Resource {
	return &organizationPrototypeResource{}
}

type organizationPrototypeResource struct {
	client *quay_api.APIClient
}

type prototypeModelJSON struct {
	ID             string                  `json:"id"`
	Role           string                  `json:"role"`
	Delegate       prototypeUserModelJSON  `json:"delegate"`
	ActivatingUser *prototypeUserModelJSON `json:"activating_user"`
}

type prototypeUserModelJSON struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type prototypeListModelJSON struct {
	Prototypes []prototypeModelJSON `json:"prototypes"`
}

func (r *organizationPrototypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_prototype"
}

func (r *organizationPrototypeResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_organization_prototype.OrganizationPrototypeResourceSchema(ctx)
}

func (r *organizationPrototypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_organization_prototype.OrganizationPrototypeModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()

	// Create prototype
	newPrototype := quay_api.NewPrototype{
		Role:           data.Role.ValueString(),
		ActivatingUser: quay_api.NewNewPrototypeActivatingUser(data.ActivatingUser.ValueString()),
		Delegate: quay_api.NewPrototypeDelegate{
			Name: data.DelegateName.ValueString(),
			Kind: data.DelegateKind.ValueString(),
		},
	}
	httpRes, err := r.client.PrototypeAPI.CreateOrganizationPrototypePermission(context.Background(), orgName).Body(newPrototype).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay prototype", "Could not create Quay prototype, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	// Parse response to get prototype ID
	var resPrototypeData prototypeModelJSON
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay prototype response",
			"Could not read Quay prototype response, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resPrototypeData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error parsing Quay prototype response",
			"Could not parse Quay prototype response, unexpected error: "+err.Error())
		return
	}

	data.Id = types.StringValue(resPrototypeData.ID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationPrototypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_organization_prototype.OrganizationPrototypeModel
	var resPrototypeList prototypeListModelJSON

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	prototypeID := data.Id.ValueString()

	// Get prototypes
	httpRes, err := r.client.PrototypeAPI.GetOrganizationPrototypePermissions(context.Background(), orgName).Execute()
	if err != nil {
		if isQuayNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay prototype", "Could not read Quay prototype, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay prototype",
			"Could not read Quay prototype, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resPrototypeList)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay prototype",
			"Could not read Quay prototype, unexpected error: "+err.Error())
		return
	}

	// Find prototype by ID
	for _, prototype := range resPrototypeList.Prototypes {
		if prototype.ID != prototypeID {
			continue
		}

		data.Role = types.StringValue(prototype.Role)
		data.DelegateName = types.StringValue(prototype.Delegate.Name)
		data.DelegateKind = types.StringValue(prototype.Delegate.Kind)
		if prototype.ActivatingUser != nil {
			data.ActivatingUser = types.StringValue(prototype.ActivatingUser.Name)
		}

		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// Prototype was deleted outside of Terraform
	resp.State.RemoveResource(ctx)
}

func (r *organizationPrototypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var dataState resource_organization_prototype.OrganizationPrototypeModel
	var dataPlan resource_organization_prototype.OrganizationPrototypeModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &dataState)...)

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &dataPlan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := dataPlan.Orgname.ValueString()
	prototypeID := dataState.Id.ValueString()

	// Update prototype role
	if !dataPlan.Role.Equal(dataState.Role) {
		updatePrototype := quay_api.NewPrototypeUpdate(dataPlan.Role.ValueString())
		_, err := r.client.PrototypeAPI.UpdateOrganizationPrototypePermission(context.Background(), orgName, prototypeID).Body(*updatePrototype).Execute()
		if err != nil {
			errDetail := handleQuayAPIError(err)
			resp.Diagnostics.AddError("Error updating Quay prototype", "Could not update Quay prototype, unexpected error: "+errDetail)
			return
		}
	}

	dataPlan.Id = dataState.Id

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &dataPlan)...)
}

func (r *organizationPrototypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_organization_prototype.OrganizationPrototypeModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	prototypeID := data.Id.ValueString()

	// Delete prototype
	_, err := r.client.PrototypeAPI.DeleteOrganizationPrototypePermission(context.Background(), orgName, prototypeID).Execute()
	if err != nil {
		if isQuayNotFound(err) {
			return
		}
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay prototype", "Could not delete Quay prototype, unexpected error: "+errDetail)
		return
	}
}

func (r *organizationPrototypeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *organizationPrototypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idSplit := strings.Split(req.ID, "/")
	if len(idSplit) != 2 || idSplit[0] == "" || idSplit[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format orgname/prototypeid. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("orgname"), idSplit[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idSplit[1])...)
}
