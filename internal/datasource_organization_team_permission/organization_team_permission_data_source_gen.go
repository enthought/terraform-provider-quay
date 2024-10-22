// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_organization_team_permission

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OrganizationTeamPermissionDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"orgname": schema.StringAttribute{
				Required:            true,
				Description:         "Organization name",
				MarkdownDescription: "Organization name",
			},
			"permission": schema.StringAttribute{
				Computed:            true,
				Description:         "Team permission",
				MarkdownDescription: "Team permission",
			},
			"reponame": schema.StringAttribute{
				Required:            true,
				Description:         "Repository name",
				MarkdownDescription: "Repository name",
			},
			"teamname": schema.StringAttribute{
				Required:            true,
				Description:         "Team name",
				MarkdownDescription: "Team name",
			},
		},
	}
}

type OrganizationTeamPermissionModel struct {
	Orgname    types.String `tfsdk:"orgname"`
	Permission types.String `tfsdk:"permission"`
	Reponame   types.String `tfsdk:"reponame"`
	Teamname   types.String `tfsdk:"teamname"`
}
