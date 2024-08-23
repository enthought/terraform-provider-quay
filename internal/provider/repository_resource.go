// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Modifications copyright (c) Enthought, Inc.
// SPDX-License-Identifier:	BSD-3-Clause

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"io"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/resource_repository"
)

var (
	_ resource.Resource                = (*repositoryResource)(nil)
	_ resource.ResourceWithConfigure   = (*repositoryResource)(nil)
	_ resource.ResourceWithImportState = (*repositoryResource)(nil)
)

func NewRepositoryResource() resource.Resource {
	return &repositoryResource{}
}

type repositoryResource struct {
	client *quay_api.APIClient
}

type repositoryModelJSON struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
}

func (r *repositoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (r *repositoryResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_repository.RepositoryResourceSchema(ctx)
}

func (r *repositoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_repository.RepositoryModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	repoName := data.Name.ValueString()
	visibility := data.Visibility.ValueString()
	namespace := data.Namespace.ValueString()
	description := data.Description.ValueString()

	// Create repository
	newRepo := quay_api.NewRepo{
		Repository:  repoName,
		Visibility:  visibility,
		Namespace:   &namespace,
		Description: description,
	}
	_, err := r.client.RepositoryAPI.CreateRepo(context.Background()).Body(newRepo).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error creating Quay repository", "Could not create Quay repository, unexpected error: "+errDetail)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *repositoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_repository.RepositoryModel
	var resRepoData repositoryModelJSON

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	repoName := data.Name.ValueString()
	namespace := data.Namespace.ValueString()

	// Get repository
	httpRes, err := r.client.RepositoryAPI.GetRepo(context.Background(), namespace+"/"+repoName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error reading Quay repository", "Could not create Quay repository, unexpected error: "+errDetail)
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

func (r *repositoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_repository.RepositoryModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *repositoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_repository.RepositoryModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	repoName := data.Name.ValueString()
	namespace := data.Namespace.ValueString()

	_, err := r.client.RepositoryAPI.DeleteRepository(context.Background(), namespace+"/"+repoName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay repository", "Could not delete Quay repository, unexpected error: "+errDetail)
		return
	}
}

func (r *repositoryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *repositoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idSplit := strings.Split(req.ID, "/")
	if len(idSplit) != 2 || idSplit[0] == "" || idSplit[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format namespace/repository. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), idSplit[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idSplit[1])...)
}
