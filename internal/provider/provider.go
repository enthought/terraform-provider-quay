package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

var _ provider.Provider = (*quayProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &quayProvider{}
	}
}

type quayProvider struct{}

type quayProviderModel struct {
	Url          types.String `tfsdk:"url"`
	Token        types.String `tfsdk:"token"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

func (p *quayProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"url": schema.StringAttribute{
			Description: "Quay URL. May also be provided via the QUAY_URL environment variable. Example: https://quay.example.com",
			Optional:    true,
		},
		"token": schema.StringAttribute{
			Description: "Quay token. May also be provided via the QUAY_TOKEN environment variable.",
			Optional:    true,
			Sensitive:   true,
		},
		"client_id": schema.StringAttribute{
			Description: "Client ID. Used for generating a JWT OAuth2 access token with client credentials.",
			Optional:    true,
		},
		"client_secret": schema.StringAttribute{
			Description: "Client secret. Used for generating a JWT OAuth2 access token with client credentials.",
			Optional:    true,
			Sensitive:   true,
		},
	}}
}

func (p *quayProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Quay client")

	var config quayProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	quayURL := os.Getenv("QUAY_URL")
	quayToken := os.Getenv("QUAY_TOKEN")

	if !config.Url.IsNull() {
		quayURL = config.Url.ValueString()
	}

	if !config.Token.IsNull() {
		quayToken = config.Token.ValueString()
	}

	ctx = tflog.SetField(ctx, "quay_url", quayURL)
	ctx = tflog.SetField(ctx, "quay_token", quayToken)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "quay_token")

	tflog.Debug(ctx, "Creating Quay client")

	configuration := &quay_api.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: quay_api.ServerConfigurations{
			{
				URL:         quayURL,
				Description: "No description provided",
			},
		},
		OperationServers: map[string]quay_api.ServerConfigurations{},
	}
	configuration.AddDefaultHeader("Authorization", "Bearer "+quayToken)
	client := quay_api.NewAPIClient(configuration)

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Quay client", map[string]any{"success": true})
}

func (p *quayProvider) ValidateConfig(ctx context.Context, req provider.ValidateConfigRequest, resp *provider.ValidateConfigResponse) {
	var config quayProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Url.IsUnknown() || config.Token.IsUnknown() || config.ClientID.IsUnknown() || config.ClientSecret.IsUnknown() {
		resp.Diagnostics.AddError(
			"Unknown configuration values",
			"The provider cannot create the Quay client if any configuration values are unknown. "+
				"Either target apply the source of the unknown value(s) first, set the value(s) statically in the configuration, or set the appropriate environment variable(s).",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	quayURL := os.Getenv("QUAY_URL")
	quayToken := os.Getenv("QUAY_TOKEN")
	clientID := os.Getenv("QUAY_CLIENT_ID")
	clientSecret := os.Getenv("QUAY_CLIENT_SECRET")

	if !config.Url.IsNull() {
		quayURL = config.Url.ValueString()
	}

	if !config.Token.IsNull() {
		quayToken = config.Token.ValueString()
	}

	if !config.ClientID.IsNull() {
		clientID = config.ClientID.ValueString()
	}

	if !config.ClientSecret.IsNull() {
		clientSecret = config.ClientSecret.ValueString()
	}

	if (quayToken != "" && clientID != "") || (quayToken != "" && clientSecret != "") {
		resp.Diagnostics.AddError(
			"Cannot specify token and client credentials",
			"Token cannot be specified when a client ID or client secret are also specified. You must pick one authentication method.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	if quayURL == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Missing Quay URL",
			"The provider cannot create the Quay client as there is a missing or empty value for the Quay URL. "+
				"Set the URL in the configuration or use the QUAY_URL environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	if !isValidURL(quayURL) {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Quay URL is not a valid URL",
			"The provider cannot create the Quay client as the URL provided is not valid.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	if quayToken == "" && (clientID == "" && clientSecret == "") {
		resp.Diagnostics.AddError(
			"Missing Quay token and client credentials",
			"The provider cannot create the Quay client as both the Quay token and client credentials are missing or empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	if clientID != "" && clientSecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing client secret",
			"The provider cannot create the Quay client as there is a missing value or empty value for the client secret."+
				"Set the client secret in the configuration or use the QUAY_CLIENT_SECRET environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	if clientID == "" && clientSecret != "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing client ID",
			"The provider cannot create the Quay client as there is a missing value or empty value for the client ID."+
				"Set the client ID in the configuration or use the QUAY_CLIENT_ID environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *quayProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "quay"
}

func (p *quayProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOrganizationDataSource,
		NewRepositoryDataSource,
		NewOrganizationTeamPermissionDataSource,
		NewOrganizationTeamDataSource,
		NewOrganizationRobotDataSource,
	}
}

func (p *quayProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrganizationResource,
		NewOrganizationTeamResource,
		NewOrganizationRobotResource,
		NewRepositoryResource,
		NewOrganizationTeamPermissionResource,
	}
}
