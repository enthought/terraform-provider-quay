package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/datasource_organization_robot"
)

var (
	_ datasource.DataSource              = (*organizationRobotDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*organizationRobotDataSource)(nil)
)

func NewOrganizationRobotDataSource() datasource.DataSource {
	return &organizationRobotDataSource{}
}

type organizationRobotDataSource struct {
	client *quay_api.APIClient
}

func (d *organizationRobotDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_robot"
}

func (d *organizationRobotDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_organization_robot.OrganizationRobotDataSourceSchema(ctx)
}

func (d *organizationRobotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_organization_robot.OrganizationRobotModel
	var resRobotData organizationRobotModelJSON

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	orgName := data.Orgname.ValueString()
	robotName := data.Name.ValueString()

	// Get robot
	httpRes, err := d.client.RobotAPI.GetOrgRobot(context.Background(), orgName, robotName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError(
			"Error reading Quay org robot",
			"Could not read Quay org robot, unexpected error: "+errDetail)
		return
	}
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team",
			"Could not read Quay team, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resRobotData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay team",
			"Could not read Quay team, unexpected error: "+err.Error())
		return
	}

	// Set Description
	data.Description = types.StringValue(resRobotData.Description)

	// Set robot full name
	data.Fullname = types.StringValue(orgName + "+" + robotName)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *organizationRobotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
