package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"terraform-provider-quay/internal/datasource_organization_team_permission"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

var (
	_ datasource.DataSource              = (*organizationTeamPermissionDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*organizationTeamPermissionDataSource)(nil)
)

func NewOrganizationTeamPermissionDataSource() datasource.DataSource {
	return &organizationTeamPermissionDataSource{}
}

type organizationTeamPermissionDataSource struct {
	client *quay_api.APIClient
}

func (d *organizationTeamPermissionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_team_permission"
}

func (d *organizationTeamPermissionDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_organization_team_permission.OrganizationTeamPermissionDataSourceSchema(ctx)
}

func (d *organizationTeamPermissionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_organization_team_permission.OrganizationTeamPermissionModel
	var resData teamPermissionModelJSON

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	repoName := data.Reponame.ValueString()
	teamName := data.Teamname.ValueString()

	// Get team permission
	httpRes, err := d.client.PermissionAPI.GetTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Execute()
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *organizationTeamPermissionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
