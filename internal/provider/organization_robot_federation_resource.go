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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = (*organizationRobotFederationResource)(nil)
	_ resource.ResourceWithConfigure   = (*organizationRobotFederationResource)(nil)
	_ resource.ResourceWithImportState = (*organizationRobotFederationResource)(nil)
)

func NewOrganizationRobotFederationResource() resource.Resource {
	return &organizationRobotFederationResource{}
}

type organizationRobotFederationResource struct {
	client *quay_api.APIClient
}

// organizationRobotFederationModel is the Terraform state model for the resource.
type organizationRobotFederationModel struct {
	Orgname    types.String                      `tfsdk:"orgname"`
	Robotname  types.String                      `tfsdk:"robotname"`
	Federation []organizationRobotFederationRule `tfsdk:"federation"`
}

// organizationRobotFederationRule is a single issuer/subject federation entry.
type organizationRobotFederationRule struct {
	Issuer  types.String `tfsdk:"issuer"`
	Subject types.String `tfsdk:"subject"`
}

func (r *organizationRobotFederationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_robot_federation"
}

func (r *organizationRobotFederationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages the OIDC federation configuration for an organization robot account. " +
			"This allows a robot to be authenticated using tokens issued by an external OIDC provider " +
			"(keyless / SSO robot federation). The set of federation rules is managed as a whole; " +
			"applying this resource replaces the robot's entire federation configuration.",
		Attributes: map[string]schema.Attribute{
			"orgname": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the organization that owns the robot.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"robotname": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The short name of the robot (without the organization prefix).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"federation": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "The list of OIDC federation rules for the robot.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"issuer": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "The issuer (`iss` claim) of the OIDC token, e.g. `https://token.actions.githubusercontent.com`.",
						},
						"subject": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "The subject (`sub` claim) of the OIDC token that is permitted to authenticate as this robot.",
						},
					},
				},
			},
		},
	}
}

// apiBody converts the state model's federation rules into the API request body.
func (m organizationRobotFederationModel) apiBody() []quay_api.CreateRobotFederationInner {
	body := make([]quay_api.CreateRobotFederationInner, 0, len(m.Federation))
	for _, rule := range m.Federation {
		body = append(body, quay_api.CreateRobotFederationInner{
			Issuer:  rule.Issuer.ValueString(),
			Subject: rule.Subject.ValueString(),
		})
	}
	return body
}

// federationEntryJSON matches the JSON shape returned by the federation GET endpoint.
type federationEntryJSON struct {
	Issuer  string `json:"issuer"`
	Subject string `json:"subject"`
}

func (r *organizationRobotFederationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data organizationRobotFederationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgName := data.Orgname.ValueString()
	robotName := data.Robotname.ValueString()

	httpRes, err := r.client.RobotAPI.CreateOrgRobotFederation(ctx, robotName, orgName).Body(data.apiBody()).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay robot federation", "Could not create Quay robot federation, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationRobotFederationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data organizationRobotFederationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgName := data.Orgname.ValueString()
	robotName := data.Robotname.ValueString()

	httpRes, err := r.client.RobotAPI.GetOrgRobotFederation(ctx, robotName, orgName).Execute()
	if err != nil {
		// If the robot (or its federation config) no longer exists, remove from state.
		if httpRes != nil && httpRes.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay robot federation", "Could not read Quay robot federation, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Quay robot federation", "Could not read Quay robot federation response, unexpected error: "+err.Error())
		return
	}

	var entries []federationEntryJSON
	if err := json.Unmarshal(body, &entries); err != nil {
		resp.Diagnostics.AddError("Error parsing Quay robot federation", "Could not parse Quay robot federation response, unexpected error: "+err.Error())
		return
	}

	rules := make([]organizationRobotFederationRule, 0, len(entries))
	for _, e := range entries {
		rules = append(rules, organizationRobotFederationRule{
			Issuer:  types.StringValue(e.Issuer),
			Subject: types.StringValue(e.Subject),
		})
	}
	data.Federation = rules

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationRobotFederationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data organizationRobotFederationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgName := data.Orgname.ValueString()
	robotName := data.Robotname.ValueString()

	// POST replaces the entire federation configuration, so it also serves as update.
	httpRes, err := r.client.RobotAPI.CreateOrgRobotFederation(ctx, robotName, orgName).Body(data.apiBody()).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error updating Quay robot federation", "Could not update Quay robot federation, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *organizationRobotFederationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data organizationRobotFederationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgName := data.Orgname.ValueString()
	robotName := data.Robotname.ValueString()

	httpRes, err := r.client.RobotAPI.DeleteOrgRobotFederation(ctx, robotName, orgName).Execute()
	if err != nil {
		// Treat an already-absent federation config as a successful delete.
		if httpRes != nil && httpRes.StatusCode == 404 {
			return
		}
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay robot federation", "Could not delete Quay robot federation, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()
}

func (r *organizationRobotFederationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import ID format: "<orgname>/<robotname>".
	parts := strings.SplitN(req.ID, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier in the format \"orgname/robotname\", got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("orgname"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("robotname"), parts[1])...)
}

func (r *organizationRobotFederationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*quay_api.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *quay_api.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}
