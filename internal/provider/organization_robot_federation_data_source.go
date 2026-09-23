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

	"github.com/enthought/terraform-provider-quay/quay_api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = (*organizationRobotFederationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*organizationRobotFederationDataSource)(nil)
)

func NewOrganizationRobotFederationDataSource() datasource.DataSource {
	return &organizationRobotFederationDataSource{}
}

type organizationRobotFederationDataSource struct {
	client *quay_api.APIClient
}

func (d *organizationRobotFederationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_robot_federation"
}

func (d *organizationRobotFederationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Reads the OIDC federation configuration for an organization robot account.",
		Attributes: map[string]schema.Attribute{
			"orgname": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the organization that owns the robot.",
			},
			"robotname": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The short name of the robot (without the organization prefix).",
			},
			"federation": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "The list of OIDC federation rules configured for the robot.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"issuer": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The issuer (`iss` claim) of the OIDC token.",
						},
						"subject": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The subject (`sub` claim) of the OIDC token.",
						},
					},
				},
			},
		},
	}
}

func (d *organizationRobotFederationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data organizationRobotFederationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgName := data.Orgname.ValueString()
	robotName := data.Robotname.ValueString()

	httpRes, err := d.client.RobotAPI.GetOrgRobotFederation(ctx, robotName, orgName).Execute()
	if err != nil {
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

func (d *organizationRobotFederationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*quay_api.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *quay_api.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}
