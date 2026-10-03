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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-quay/internal/resource_repository"

	"github.com/enthought/terraform-provider-quay/quay_api"
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
	State       string `json:"state"`
}

func (r *repositoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (r *repositoryResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_repository.RepositoryResourceSchema(ctx)
}

// ValidateConfig enforces state/mirror consistency at plan time so misconfiguration is
// caught before any API call. The mirror block's required attributes are declared optional
// in the schema (the framework validates required attributes even inside an absent block,
// which would block "transition out of mirror"), so their presence is enforced here and at
// Create/Update via validateMirrorConsistency.
func (r *repositoryResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_repository.RepositoryModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(validateMirrorConsistency(data.State, data.Mirror)...)
}

func (r *repositoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_repository.RepositoryModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Enforce state/mirror consistency before any API call
	resp.Diagnostics.Append(validateMirrorConsistency(data.State, data.Mirror)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	repoName := data.Name.ValueString()
	visibility := data.Visibility.ValueString()
	namespace := data.Namespace.ValueString()
	description := data.Description.ValueString()
	repoRef := namespace + "/" + repoName

	// Create repository
	newRepo := quay_api.NewRepo{
		Repository:  repoName,
		Visibility:  visibility,
		Namespace:   &namespace,
		Description: description,
	}
	_, err := r.client.RepositoryAPI.CreateRepo(context.Background()).Body(newRepo).Execute()
	if err != nil {
		if !isRepoAlreadyExists(err) {
			errDetail := handleQuayAPIError(err)
			resp.Diagnostics.AddError("Error creating Quay repository", "Could not create Quay repository, unexpected error: "+errDetail)
			return
		}
		// The repository already exists (created by an earlier, interrupted apply).
		// Adopt it: the state/mirror reconciliation below brings it up to date.
	}

	state := data.State.ValueString()
	if state == repoStateMirror || state == repoStateReadonly {
		// A just-created repository is NORMAL by default. When adopting an existing
		// repository (created by an earlier, interrupted apply) it may already be in
		// the target state, so only transition when it actually differs.
		current, ok := repoState(r.client, ctx, repoRef)
		if !ok {
			resp.Diagnostics.AddError("Error creating Quay repository", "Could not read Quay repository state after creation: "+repoRef)
			return
		}
		if current != state {
			resp.Diagnostics.Append(r.setRepoState(ctx, repoRef, state)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
	}

	if state == repoStateMirror {
		// Configure the mirror (no-op on a fresh NORMAL->MIRROR create, where the
		// record does not yet exist, and a PUT when an interrupted apply left a
		// partial record behind).
		resp.Diagnostics.Append(r.applyMirror(ctx, repoRef, data.Mirror)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Trigger an immediate sync if requested
		if data.Mirror.SyncOnChange.ValueBool() {
			resp.Diagnostics.Append(r.triggerMirrorSync(ctx, repoRef)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
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
		resp.Diagnostics.AddError("Error reading Quay repository", "Could not read Quay repository, unexpected error: "+errDetail)
		return
	}
	defer httpRes.Body.Close()

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
			"Could not parse Quay repository, unexpected error: "+err.Error())
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
	data.State = types.StringValue(mirrorStateFromAPI(resRepoData.State))

	// Read the mirror configuration when the repository is a mirror
	if data.State.ValueString() == repoStateMirror {
		mirror, diags := getMirror(r.client, ctx, namespace+"/"+repoName)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		mv, d := buildMirrorValueFromAPI(ctx, mirror, data.Mirror.UpstreamPassword, syncOnChangeFromMirror(data.Mirror), data.Mirror.SyncTimeout)
		resp.Diagnostics.Append(d...)
		data.Mirror = mv
	} else {
		// Not a mirror: the mirror block is null. A dormant mirror record, if any,
		// is intentionally ignored.
		data.Mirror = resource_repository.NewMirrorValueNull()
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *repositoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var dataState resource_repository.RepositoryModel
	var dataPlan resource_repository.RepositoryModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &dataState)...)

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &dataPlan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create variables
	repoName := dataPlan.Name.ValueString()
	namespace := dataPlan.Namespace.ValueString()
	description := dataPlan.Description.ValueString()
	repoRef := namespace + "/" + repoName

	resp.Diagnostics.Append(validateMirrorConsistency(dataPlan.State, dataPlan.Mirror)...)
	if resp.Diagnostics.HasError() {
		return
	}

	prevState := dataState.State.ValueString()
	newState := dataPlan.State.ValueString()

	// Update description
	if dataPlan.Description != dataState.Description {
		updateRepo := quay_api.NewRepoUpdate(description)
		_, err := r.client.RepositoryAPI.UpdateRepo(context.Background(), repoRef).Body(*updateRepo).Execute()
		if err != nil {
			errDetail := handleQuayAPIError(err)
			resp.Diagnostics.AddError("Error updating Quay repository", "Could not update Quay repository, unexpected error: "+errDetail)
			return
		}
	}

	// Update visibility
	visibility := dataPlan.Visibility.ValueString()
	prevVisibility := dataState.Visibility.ValueString()
	if visibility != prevVisibility {
		changeVisibility := quay_api.NewChangeVisibility(visibility)
		_, err := r.client.RepositoryAPI.ChangeRepoVisibility(context.Background(), repoRef).Body(*changeVisibility).Execute()
		if err != nil {
			errDetail := handleQuayAPIError(err)
			resp.Diagnostics.AddError("Error updating Quay repository", "Could not update Quay repository visibility, unexpected error: "+errDetail)
			return
		}
	}

	// Update state / mirror configuration
	stateChanged := prevState != newState
	mirrorChanged := !dataPlan.Mirror.Equal(dataState.Mirror)

	switch {
	case stateChanged && newState == repoStateMirror:
		// Transition into mirror
		resp.Diagnostics.Append(r.setRepoState(ctx, repoRef, newState)...)
		if resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(r.applyMirror(ctx, repoRef, dataPlan.Mirror)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if dataPlan.Mirror.SyncOnChange.ValueBool() {
			resp.Diagnostics.Append(r.triggerMirrorSync(ctx, repoRef)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}

	case stateChanged && prevState == repoStateMirror && newState != repoStateMirror:
		// Transition out of mirror. The platform keeps the mirror record dormant,
		// so only the state change is needed.
		resp.Diagnostics.Append(r.setRepoState(ctx, repoRef, newState)...)
		if resp.Diagnostics.HasError() {
			return
		}

	case stateChanged:
		// Change between normal and readonly (no mirror involved)
		resp.Diagnostics.Append(r.setRepoState(ctx, repoRef, newState)...)
		if resp.Diagnostics.HasError() {
			return
		}

	case !stateChanged && newState == repoStateMirror && mirrorChanged:
		// Still a mirror but the configuration changed
		resp.Diagnostics.Append(r.applyMirror(ctx, repoRef, dataPlan.Mirror)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if dataPlan.Mirror.SyncOnChange.ValueBool() {
			resp.Diagnostics.Append(r.triggerMirrorSync(ctx, repoRef)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &dataPlan)...)
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

	// A mirror record is removed together with the repository; no separate cleanup is needed.
	_, err := r.client.RepositoryAPI.DeleteRepository(context.Background(), namespace+"/"+repoName).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		resp.Diagnostics.AddError("Error deleting Quay repository", "Could not delete Quay repository, unexpected error: "+errDetail)
		return
	}
}

// setRepoState transitions the repository to the given state via the dedicated
// state-change endpoint (PUT /repository/{repository}/changestate).
func (r *repositoryResource) setRepoState(ctx context.Context, repoRef string, state string) diag.Diagnostics {
	var diags diag.Diagnostics
	change := quay_api.NewChangeState(mirrorStateToAPI(state))
	_, err := r.client.RepositoryAPI.ChangeRepoState(ctx, repoRef).Body(*change).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		diags.AddError("Error updating Quay repository state", "Could not update Quay repository state, unexpected error: "+errDetail)
	}
	return diags
}

// applyMirror writes the mirror configuration for a repository. On this platform the
// mirror record, once created, persists and cannot be deleted, so the provider updates it
// via PUT when it already exists and creates it via POST (which additionally requires
// sync_start_date) otherwise.
func (r *repositoryResource) applyMirror(ctx context.Context, repoRef string, m resource_repository.MirrorValue) diag.Diagnostics {
	var diags diag.Diagnostics
	includeSyncTimeout := syncTimeoutSupported(r.client, ctx)
	mirror, d := mirrorValueToAPI(ctx, m, includeSyncTimeout)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	if mirror == nil {
		diags.AddError(
			"Invalid mirror configuration",
			"A mirror block is required when the repository state is \"mirror\".",
		)
		return diags
	}
	if mirrorRecordExists(r.client, ctx, repoRef) {
		_, err := r.client.RepositoryAPI.PutRepoMirror(ctx, repoRef).Body(*mirror).Execute()
		if err != nil {
			errDetail := handleQuayAPIError(err)
			diags.AddError("Error configuring Quay repository mirror", "Could not configure Quay repository mirror, unexpected error: "+errDetail)
		}
		return diags
	}
	mirror.SetSyncStartDate(time.Now().UTC().Format(time.RFC3339))
	_, err := r.client.RepositoryAPI.PostRepoMirror(ctx, repoRef).Body(*mirror).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		diags.AddError("Error creating Quay repository mirror", "Could not create Quay repository mirror, unexpected error: "+errDetail)
	}
	return diags
}

// triggerMirrorSync asks the platform to start an immediate mirror sync. The operation is
// asynchronous: the provider returns as soon as the trigger is accepted.
func (r *repositoryResource) triggerMirrorSync(ctx context.Context, repoRef string) diag.Diagnostics {
	var diags diag.Diagnostics
	_, err := r.client.RepositoryAPI.SyncRepoMirrorNow(ctx, repoRef).Execute()
	if err != nil {
		if isSyncAlreadyRunning(err) {
			return diags
		}
		errDetail := handleQuayAPIError(err)
		diags.AddError("Error triggering Quay repository mirror sync", "Could not trigger Quay repository mirror sync, unexpected error: "+errDetail)
	}
	return diags
}

// isSyncAlreadyRunning reports whether the error from sync-now means a sync is already
// in progress. This Quay fork answers the sync-now endpoint with 404 when a sync has
// not finished yet, so a new sync cannot be started.
func isSyncAlreadyRunning(err error) bool {
	apiError, ok := apiErrorModel(err)
	return ok && apiError.GetStatus() == 404
}

// isRepoAlreadyExists reports whether the error from create-repo means the repository
// already exists, so the create can be adopted instead of failed.
func isRepoAlreadyExists(err error) bool {
	apiError, ok := apiErrorModel(err)
	return ok && apiError.GetStatus() == 400 &&
		strings.Contains(apiError.GetDetail(), "Repository already exists")
}

// repoState reads the current state of a repository, mapped to a provider state value.
// Returns ("", false) when the repository cannot be read.
func repoState(client *quay_api.APIClient, ctx context.Context, repoRef string) (string, bool) {
	httpRes, err := client.RepositoryAPI.GetRepo(ctx, repoRef).Execute()
	if err != nil {
		return "", false
	}
	defer httpRes.Body.Close()
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return "", false
	}
	var res repositoryModelJSON
	if err := json.Unmarshal(body, &res); err != nil {
		return "", false
	}
	return mirrorStateFromAPI(res.State), true
}

// getMirror fetches the mirror configuration for a repository.
func getMirror(client *quay_api.APIClient, ctx context.Context, repoRef string) (*quay_api.Mirror, diag.Diagnostics) {
	var diags diag.Diagnostics
	httpRes, err := client.RepositoryAPI.GetRepoMirror(ctx, repoRef).Execute()
	if err != nil {
		errDetail := handleQuayAPIError(err)
		diags.AddError("Error reading Quay repository mirror", "Could not read Quay repository mirror, unexpected error: "+errDetail)
		return nil, diags
	}
	defer httpRes.Body.Close()
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		diags.AddError(
			"Error reading Quay repository mirror",
			"Could not read Quay repository mirror, unexpected error: "+err.Error())
		return nil, diags
	}
	mirror := quay_api.Mirror{}
	if err := json.Unmarshal(body, &mirror); err != nil {
		diags.AddError(
			"Error reading Quay repository mirror",
			"Could not parse Quay repository mirror, unexpected error: "+err.Error())
		return nil, diags
	}
	return &mirror, diags
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
	idSplit := strings.SplitN(req.ID, "/", 2)
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
