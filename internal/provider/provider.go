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
	"golang.org/x/oauth2"

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
	Url            types.String `tfsdk:"url"`
	Token          types.String `tfsdk:"token"`
	OAuth2Username types.String `tfsdk:"oauth2_username"`
	OAuth2Password types.String `tfsdk:"oauth2_password"`
	OAuth2ClientID types.String `tfsdk:"oauth2_client_id"`
	OAuth2TokenURL types.String `tfsdk:"oauth2_token_url"`
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
		"oauth2_username": schema.StringAttribute{
			Description: "OAuth2 username. Used for generating a JWT OAuth2 access token with password grant.",
			Optional:    true,
		},
		"oauth2_password": schema.StringAttribute{
			Description: "OAuth2 password. Used for generating a JWT OAuth2 access token with password grant.",
			Optional:    true,
			Sensitive:   true,
		},
		"oauth2_client_id": schema.StringAttribute{
			Description: "OAuth2 client ID. Used for generating a JWT OAuth2 access token with password grant.",
			Optional:    true,
		},
		"oauth2_token_url": schema.StringAttribute{
			Description: "OAuth2 token endpoint URL. Used for generating a JWT OAuth2 access token with password grant.",
			Optional:    true,
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

	url := os.Getenv("QUAY_URL")
	token := os.Getenv("QUAY_TOKEN")
	oauth2Username := os.Getenv("QUAY_OAUTH2_USERNAME")
	oauth2Password := os.Getenv("QUAY_OAUTH2_PASSWORD")
	oauth2ClientID := os.Getenv("QUAY_OAUTH2_CLIENT_ID")
	oauth2TokenURL := os.Getenv("QUAY_OAUTH2_TOKEN_URL")

	if !config.Url.IsNull() {
		url = config.Url.ValueString()
	}

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if !config.OAuth2Username.IsNull() {
		oauth2Username = config.OAuth2Username.ValueString()
	}

	if !config.OAuth2Password.IsNull() {
		oauth2Password = config.OAuth2Password.ValueString()
	}

	if !config.OAuth2ClientID.IsNull() {
		oauth2ClientID = config.OAuth2ClientID.ValueString()
	}

	if !config.OAuth2TokenURL.IsNull() {
		oauth2TokenURL = config.OAuth2TokenURL.ValueString()
	}

	ctx = tflog.SetField(ctx, "url", url)
	ctx = tflog.SetField(ctx, "token", token)
	ctx = tflog.SetField(ctx, "oauth2_username", oauth2Username)
	ctx = tflog.SetField(ctx, "oauth2_password", oauth2Password)
	ctx = tflog.SetField(ctx, "oauth2_client_id", oauth2ClientID)
	ctx = tflog.SetField(ctx, "oauth2_token_url", oauth2TokenURL)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "token")
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "oauth2_password")

	tflog.Debug(ctx, "Creating Quay client")

	configuration := &quay_api.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: quay_api.ServerConfigurations{
			{
				URL:         url,
				Description: "No description provided",
			},
		},
		OperationServers: map[string]quay_api.ServerConfigurations{},
	}

	if token == "" {
		oauth2Config := &oauth2.Config{
			ClientID: oauth2ClientID,
			Endpoint: oauth2.Endpoint{
				TokenURL: oauth2TokenURL,
			},
		}

		oauth2Token, err := oauth2Config.PasswordCredentialsToken(ctx, oauth2Username, oauth2Password)
		if err != nil {
			resp.Diagnostics.AddError("Error retrieving OAuth2 access token",
				"Error retrieving OAuth2 access token: "+err.Error())
			return
		}

		token = oauth2Token.AccessToken
	}

	configuration.AddDefaultHeader("Authorization", "Bearer "+token)
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

	if config.Url.IsUnknown() || config.Token.IsUnknown() || config.OAuth2Username.IsUnknown() || config.OAuth2Password.IsUnknown() ||
		config.OAuth2ClientID.IsUnknown() || config.OAuth2TokenURL.IsUnknown() {
		resp.Diagnostics.AddError(
			"Unknown configuration values",
			"The provider cannot create the Quay client if any configuration values are unknown. "+
				"Either target apply the source of the unknown value(s) first, set the value(s) statically in the configuration, or set the appropriate environment variable(s).",
		)
		return
	}

	url := os.Getenv("QUAY_URL")
	token := os.Getenv("QUAY_TOKEN")
	oauth2Username := os.Getenv("QUAY_OAUTH2_USERNAME")
	oauth2Password := os.Getenv("QUAY_OAUTH2_PASSWORD")
	oauth2ClientID := os.Getenv("QUAY_OAUTH2_CLIENT_ID")
	oauth2TokenURL := os.Getenv("QUAY_OAUTH2_TOKEN_URL")

	if !config.Url.IsNull() {
		url = config.Url.ValueString()
	}

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if !config.OAuth2Username.IsNull() {
		oauth2Username = config.OAuth2Username.ValueString()
	}

	if !config.OAuth2Password.IsNull() {
		oauth2Password = config.OAuth2Password.ValueString()
	}

	if !config.OAuth2ClientID.IsNull() {
		oauth2ClientID = config.OAuth2ClientID.ValueString()
	}

	if !config.OAuth2TokenURL.IsNull() {
		oauth2TokenURL = config.OAuth2TokenURL.ValueString()
	}
	if token != "" && (oauth2Username != "" || oauth2Password != "" || oauth2ClientID != "" || oauth2TokenURL != "") {
		resp.Diagnostics.AddError(
			"Cannot specify token and OAuth2 credentials",
			"Token cannot be specified when OAuth2 credentials are also specified. You must pick one authentication method.",
		)
		return
	}

	if url == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Missing Quay URL",
			"The provider cannot create the Quay client as there is a missing or empty value for the Quay URL. "+
				"Set the URL in the configuration or use the QUAY_URL environment variable.",
		)
		return
	}

	if !isValidURL(url) {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Quay URL is not a valid URL",
			"The provider cannot create the Quay client as the URL provided is not valid.",
		)
		return
	}

	if token == "" && oauth2Username == "" && oauth2Password == "" && oauth2ClientID == "" && oauth2TokenURL == "" {
		resp.Diagnostics.AddError(
			"Missing Quay token and OAuth2 credentials",
			"The provider cannot create the Quay client as both the Quay token and OAuth2 credentials are missing or empty.",
		)
		return
	}

	if token == "" && oauth2Username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("oauth2_username"),
			"Missing OAuth2 username",
			"The provider cannot create the Quay client as there is a missing value or empty value for the OAuth2 username."+
				"Set the OAuth2 username in the configuration or use the QUAY_OAUTH2_USERNAME environment variable.",
		)
		return
	}

	if token == "" && oauth2Password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("oauth2_password"),
			"Missing OAuth2 password",
			"The provider cannot create the Quay client as there is a missing value or empty value for the OAuth2 password."+
				"Set the OAuth2 password in the configuration or use the QUAY_OAUTH2_PASSWORD environment variable.",
		)
		return
	}

	if token == "" && oauth2ClientID == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("oauth2_client_id"),
			"Missing OAuth2 client ID",
			"The provider cannot create the Quay client as there is a missing value or empty value for the OAuth2 client ID."+
				"Set the OAuth2 client ID in the configuration or use the QUAY_OAUTH2_CLIENT_ID environment variable.",
		)
		return
	}

	if token == "" && oauth2TokenURL == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("oauth2_token_url"),
			"Missing OAuth2 token URL",
			"The provider cannot create the Quay client as there is a missing value or empty value for the OAuth2 token URL."+
				"Set the OAuth2 token URL in the configuration or use the QUAY_OAUTH2_TOKEN_URL environment variable.",
		)
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
