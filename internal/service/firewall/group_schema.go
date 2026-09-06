package firewall

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type groupResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Members     types.Set    `tfsdk:"members"`
	NoGroup     types.String `tfsdk:"no_group"`
	Sequence    types.Int64  `tfsdk:"sequence"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func groupResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Firewall groups allow multiple interfaces to be grouped together and referenced as a single firewall interface.",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the firewall group.",
				Required:            true,
			},

			"members": schema.SetAttribute{
				MarkdownDescription: "Interfaces that belong to this firewall group.",
				Required:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
				},
			},

			"no_group": schema.StringAttribute{
				MarkdownDescription: "Whether the group should not be included in automatically generated interface groups. Defaults to `\"0\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0"),
			},

			"sequence": schema.Int64Attribute{
				MarkdownDescription: "The sequence number used to order firewall groups. Must be between `0` and `9999`. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},

			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description for this firewall group.",
				Optional:            true,
			},

			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func groupDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Firewall groups allow multiple interfaces to be grouped together and referenced as a single firewall interface.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},

			"name": dschema.StringAttribute{
				MarkdownDescription: "The name of the firewall group.",
				Computed:            true,
			},

			"members": dschema.SetAttribute{
				MarkdownDescription: "Interfaces that belong to this firewall group.",
				Computed:            true,
				ElementType:         types.StringType,
			},

			"no_group": dschema.StringAttribute{
				MarkdownDescription: "Whether the group is excluded from automatically generated interface groups.",
				Computed:            true,
			},

			"sequence": dschema.Int64Attribute{
				MarkdownDescription: "The sequence number used to order firewall groups.",
				Computed:            true,
			},

			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this firewall group.",
				Computed:            true,
			},
		},
	}
}

func convertGroupSchemaToStruct(d *groupResourceModel) (*firewall.Group, error) {
	var members []string
	d.Members.ElementsAs(context.Background(), &members, false)

	return &firewall.Group{
		Name:        d.Name.ValueString(),
		Members:     members,
		NoGroup:     d.NoGroup.ValueString(),
		Sequence:    tools.Int64ToString(d.Sequence.ValueInt64()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertGroupStructToSchema(d *firewall.Group) (*groupResourceModel, error) {
	membersSet := tools.StringSliceToSet(d.Members)

	return &groupResourceModel{
		Name:        types.StringValue(d.Name),
		Members:     membersSet,
		NoGroup:     types.StringValue(d.NoGroup),
		Sequence:    types.Int64Value(tools.StringToInt64(d.Sequence)),
		Description: types.StringValue(d.Description),
	}, nil
}
