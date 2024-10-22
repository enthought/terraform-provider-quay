// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_organization_team

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrganizationTeamResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Markdown description",
				MarkdownDescription: "Markdown description",
				Default:             stringdefault.StaticString(""),
			},
			"members": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "List of team members",
				MarkdownDescription: "List of team members",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Team name",
				MarkdownDescription: "Team name",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"orgname": schema.StringAttribute{
				Required:            true,
				Description:         "Organization name",
				MarkdownDescription: "Organization name",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"role": schema.StringAttribute{
				Required:            true,
				Description:         "Team role. Should be admin, creator, or member.",
				MarkdownDescription: "Team role. Should be admin, creator, or member.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"admin", "creator", "member"}...),
				},
			},
		},
	}
}

type OrganizationTeamModel struct {
	Description types.String `tfsdk:"description"`
	Members     types.Set    `tfsdk:"members"`
	Name        types.String `tfsdk:"name"`
	Orgname     types.String `tfsdk:"orgname"`
	Role        types.String `tfsdk:"role"`
}
