package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/datasource_organization"
)

var (
	_ datasource.DataSource              = (*organizationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*organizationDataSource)(nil)
)

func NewOrganizationDataSource() datasource.DataSource {
	return &organizationDataSource{}
}

type organizationDataSource struct {
	client *quay_api.APIClient
}

func (d *organizationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (d *organizationDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_organization.OrganizationDataSourceSchema(ctx)
}

func (d *organizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_organization.OrganizationModel
	var resData organizationModelJSON
	var apiErr *quay_api.GenericOpenAPIError

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get data from API
	httpRes, err := d.client.OrganizationAPI.GetOrganization(context.Background(), data.Name.ValueString()).Execute()
	if err != nil {
		errDetail := ""
		if errors.As(err, &apiErr) {
			errDetail = string(apiErr.Body())
		} else {
			errDetail = err.Error()
		}
		resp.Diagnostics.AddError(
			"Error reading Quay org",
			"Could not read Quay org, unexpected error: "+errDetail)
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *organizationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
