package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/datasource_repository"
)

var (
	_ datasource.DataSource              = (*repositoryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*repositoryDataSource)(nil)
)

func NewRepositoryDataSource() datasource.DataSource {
	return &repositoryDataSource{}
}

type repositoryDataSource struct {
	client *quay_api.APIClient
}

func (d *repositoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (d *repositoryDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_repository.RepositoryDataSourceSchema(ctx)
}

func (d *repositoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_repository.RepositoryModel

	var resRepoData repositoryModelJSON

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	repoName := data.Name.ValueString()
	namespace := data.Namespace.ValueString()

	// Get repository
	httpRes, err := d.client.RepositoryAPI.GetRepo(context.Background(), namespace+"/"+repoName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay repository",
			"Could not create Quay repository, unexpected error: "+errDetail)
		return
	}

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay repository",
			"Could not read Quay repository, unexpected error: "+err.Error())
		return
	}
	err = json.Unmarshal(body, &resRepoData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Quay repository",
			"Could not read Quay repository, unexpected error: "+err.Error())
		return
	}

	// Set visibility
	if resRepoData.IsPublic {
		data.Visibility = types.StringValue("public")
	} else {
		data.Visibility = types.StringValue("private")
	}

	// Set data
	data.Name = types.StringValue(resRepoData.Name)
	data.Namespace = types.StringValue(resRepoData.Namespace)
	data.Description = types.StringValue(resRepoData.Description)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *repositoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
