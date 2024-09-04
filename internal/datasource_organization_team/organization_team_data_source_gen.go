// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_organization_team

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OrganizationTeamDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Computed:            true,
				Description:         "Markdown description",
				MarkdownDescription: "Markdown description",
			},
			"members": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "List of team members",
				MarkdownDescription: "List of team members",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Team name",
				MarkdownDescription: "Team name",
			},
			"orgname": schema.StringAttribute{
				Required:            true,
				Description:         "Organization name",
				MarkdownDescription: "Organization name",
			},
			"role": schema.StringAttribute{
				Computed:            true,
				Description:         "Team role",
				MarkdownDescription: "Team role",
			},
		},
	}
}

type OrganizationTeamModel struct {
	Description types.String `tfsdk:"description"`
	Members     types.List   `tfsdk:"members"`
	Name        types.String `tfsdk:"name"`
	Orgname     types.String `tfsdk:"orgname"`
	Role        types.String `tfsdk:"role"`
}