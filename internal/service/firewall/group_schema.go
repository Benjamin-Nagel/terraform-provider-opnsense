package firewall

import (
	"context"
	"regexp"

	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
		Version:             1,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Only texts containing letters, digits and underscores with a maximum length of 15 characters are allowed and the name may not end with a digit.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[0-9a-zA-Z_]{0,14}[a-zA-Z_]$`),
						"Only texts containing letters, digits and underscores with a maximum length of 15 characters are allowed and the name may not end with a digit.",
					),
					stringvalidator.LengthBetween(1, 15),
				},
			},

			"members": schema.SetAttribute{
				MarkdownDescription: "Rules for WAN type interfaces in groups do not contain the reply-to mechanism upon which Multi-WAN typically relies.",
				Required:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
				},
			},

			"sequence": schema.Int64Attribute{
				MarkdownDescription: "Priority sequence used in sorting the groups",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 9999),
				},
			},

			"no_group": schema.StringAttribute{
				MarkdownDescription: "Prevent grouping these members in the interfaces menu section",
				Optional:            true,
			},

			"description": schema.StringAttribute{
				MarkdownDescription: "You may enter a description here for your reference (not parsed).",
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
				MarkdownDescription: "Only texts containing letters, digits and underscores with a maximum length of 15 characters are allowed and the name may not end with a digit.",
				Computed:            true,
			},

			"members": dschema.SetAttribute{
				MarkdownDescription: "Rules for WAN type interfaces in groups do not contain the reply-to mechanism upon which Multi-WAN typically relies.",
				Computed:            true,
				ElementType:         types.StringType,
			},

			"sequence": dschema.Int64Attribute{
				MarkdownDescription: "Priority sequence used in sorting the groups",
				Computed:            true,
			},

			"no_group": dschema.StringAttribute{
				MarkdownDescription: "Prevent grouping these members in the interfaces menu section",
				Computed:            true,
			},

			"description": dschema.StringAttribute{
				MarkdownDescription: "You may enter a description here for your reference (not parsed).",
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
