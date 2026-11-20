package resource_organization_prototype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func OrganizationPrototypeResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description:         "Manages a Quay default permission prototype.",
		MarkdownDescription: "Manages a Quay default permission prototype. A prototype automatically grants a role to a user, robot or team on every new repository created by the activating user.",
		Attributes: map[string]schema.Attribute{
			"activating_user": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the user or robot (full name for robots, e.g. org+robot) whose repository creation triggers the rule",
				MarkdownDescription: "Name of the user or robot (full name for robots, e.g. `org+robot`) whose repository creation triggers the rule",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"delegate_kind": schema.StringAttribute{
				Required:            true,
				Description:         "Whether the delegate is a user or a team. Should be user or team.",
				MarkdownDescription: "Whether the delegate is a user or a team. Should be `user` or `team`.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"user", "team"}...),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"delegate_name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the user, robot (full name, e.g. org+robot) or team that receives the permission",
				MarkdownDescription: "Name of the user, robot (full name, e.g. `org+robot`) or team that receives the permission",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Prototype ID assigned by Quay",
				MarkdownDescription: "Prototype ID assigned by Quay",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
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
				Description:         "Permission role. Should be read, write or admin.",
				MarkdownDescription: "Permission role. Should be `read`, `write` or `admin`.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"read", "write", "admin"}...),
				},
			},
		},
	}
}

type OrganizationPrototypeModel struct {
	ActivatingUser types.String `tfsdk:"activating_user"`
	DelegateKind   types.String `tfsdk:"delegate_kind"`
	DelegateName   types.String `tfsdk:"delegate_name"`
	Id             types.String `tfsdk:"id"`
	Orgname        types.String `tfsdk:"orgname"`
	Role           types.String `tfsdk:"role"`
}
