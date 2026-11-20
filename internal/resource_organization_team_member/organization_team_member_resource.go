package resource_organization_team_member

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func OrganizationTeamMemberResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description:         "Manages a single Quay team membership.",
		MarkdownDescription: "Manages a single Quay team membership without owning the team or its other members. Use this resource for shared teams; use `quay_organization_team` to manage the full membership list of a team.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Membership ID in the format orgname/teamname/membername",
				MarkdownDescription: "Membership ID in the format `orgname/teamname/membername`",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"membername": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the user or robot (full name, e.g. org+robot) to add to the team",
				MarkdownDescription: "Name of the user or robot (full name, e.g. `org+robot`) to add to the team",
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
			"teamname": schema.StringAttribute{
				Required:            true,
				Description:         "Team name",
				MarkdownDescription: "Team name",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

type OrganizationTeamMemberModel struct {
	Id         types.String `tfsdk:"id"`
	Membername types.String `tfsdk:"membername"`
	Orgname    types.String `tfsdk:"orgname"`
	Teamname   types.String `tfsdk:"teamname"`
}
