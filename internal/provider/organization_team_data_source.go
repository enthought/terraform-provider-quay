package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/datasource_organization_team"
)

var (
	_ datasource.DataSource              = (*organizationTeamDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*organizationTeamDataSource)(nil)
)

func NewOrganizationTeamDataSource() datasource.DataSource {
	return &organizationTeamDataSource{}
}

type organizationTeamDataSource struct {
	client *quay_api.APIClient
}

func (d *organizationTeamDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_team"
}

func (d *organizationTeamDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_organization_team.OrganizationTeamDataSourceSchema(ctx)
}

func (d *organizationTeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_organization_team.OrganizationTeamModel
	var resTeamData teamModelJSON
	var resOrgData organizationModelJSON

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	teamName := data.Name.ValueString()
	orgName := data.Orgname.ValueString()

	// Get members
	httpRes, err := d.client.TeamAPI.GetOrganizationTeamMembers(context.Background(), orgName, teamName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay team", "Could not read Quay team, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

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
	members, diags := types.SetValueFrom(ctx, types.StringType, memList)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Members = members

	// Get org data
	httpRes, err = d.client.OrganizationAPI.GetOrganization(context.Background(), orgName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay team", "Could not read Quay team, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *organizationTeamDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*quay_api.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *quay_api.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
	}

	d.client = client
}
