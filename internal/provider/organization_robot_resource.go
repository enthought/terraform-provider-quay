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

	"github.com/enthought/terraform-provider-quay/quay_api"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-quay/internal/resource_organization_robot"
)

var (
	_ resource.Resource                = (*organizationRobotResource)(nil)
	_ resource.ResourceWithConfigure   = (*organizationRobotResource)(nil)
	_ resource.ResourceWithImportState = (*organizationRobotResource)(nil)
)

func NewOrganizationRobotResource() resource.Resource {
	return &organizationRobotResource{}
}

type organizationRobotResource struct {
	client *quay_api.APIClient
}

type organizationRobotModelJSON struct {
	Created      string `json:"created"`
	Description  string `json:"description"`
	LastAccessed string `json:"last_accessed"`
	Name         string `json:"name"`
	Token        string `json:"token"`
}

func (r *organizationRobotResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_robot"
}

func (r *organizationRobotResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_organization_robot.OrganizationRobotResourceSchema(ctx)
}

func (r *organizationRobotResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_organization_robot.OrganizationRobotModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	robotName := data.Name.ValueString()
	robotDescription := data.Description.ValueStringPointer()

	// Create robot
	newRobot := quay_api.CreateRobot{
		Description:          robotDescription,
		UnstructuredMetadata: nil,
	}
	httpRes, err := r.client.RobotAPI.CreateOrgRobot(context.Background(), orgName, robotName).Body(newRobot).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay org robot", "Could not create Quay org robot, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	// Parse response to get token
	var resRobotData organizationRobotModelJSON
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay org robot response",
			"Could not read Quay org robot response, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resRobotData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error parsing Quay org robot response",
			"Could not parse Quay org robot response, unexpected error: "+err.Error())
		return
	}

	// Set token if available
	if resRobotData.Token != "" {
		data.Token = types.StringValue(resRobotData.Token)
	}

	// Set robot full name
	data.Fullname = types.StringValue(orgName + "+" + robotName)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationRobotResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_organization_robot.OrganizationRobotModel
	var resRobotData organizationRobotModelJSON

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	robotName := data.Name.ValueString()

	// Get robot
	httpRes, err := r.client.RobotAPI.GetOrgRobot(context.Background(), orgName, robotName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay org robot", "Could not read Quay org robot, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay org robot",
			"Could not read Quay org robot, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resRobotData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay org robot",
			"Could not read Quay org robot, unexpected error: "+err.Error())
		return
	}

	// Set Description
	data.Description = types.StringValue(resRobotData.Description)

	// Set Token if available
	if resRobotData.Token != "" {
		data.Token = types.StringValue(resRobotData.Token)
	}

	// Set robot full name
	data.Fullname = types.StringValue(orgName + "+" + robotName)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationRobotResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_organization_robot.OrganizationRobotModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationRobotResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_organization_robot.OrganizationRobotModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	robotName := data.Name.ValueString()

	// Delete robot
	_, err := r.client.RobotAPI.DeleteOrgRobot(context.Background(), orgName, robotName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay org robot", "Could not delete Quay org robot, unexpected error: "+errDetail)
		return
	}
}

func (r *organizationRobotResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *organizationRobotResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idSplit := strings.Split(req.ID, "+")
	if len(idSplit) != 2 || idSplit[0] == "" || idSplit[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format orgname+robotname. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("orgname"), idSplit[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idSplit[1])...)
}
