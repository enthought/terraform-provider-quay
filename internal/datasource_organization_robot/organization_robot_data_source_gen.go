// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_organization_robot

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OrganizationRobotDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Computed:            true,
				Description:         "Text description",
				MarkdownDescription: "Text description",
			},
			"fullname": schema.StringAttribute{
				Computed:            true,
				Description:         "Robot full name",
				MarkdownDescription: "Robot full name",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Robot short name",
				MarkdownDescription: "Robot short name",
			},
			"orgname": schema.StringAttribute{
				Required:            true,
				Description:         "Organization name",
				MarkdownDescription: "Organization name",
			},
		},
	}
}

type OrganizationRobotModel struct {
	Description types.String `tfsdk:"description"`
	Fullname    types.String `tfsdk:"fullname"`
	Name        types.String `tfsdk:"name"`
	Orgname     types.String `tfsdk:"orgname"`
}