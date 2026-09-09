package firewall

import (
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertGroupStructToSchema(t *testing.T) {
	upstream := &firewall.Group{
		Name:        "test_group",
		Members:     []string{"wan"},
		Sequence:    "100",
		NoGroup:     "0",
		Description: "test-description",
	}

	model, err := convertGroupStructToSchema(upstream)

	require.NoError(t, err)
	require.NotNil(t, model)

	assert.Equal(t, types.StringValue("test_group"), model.Name)
	assert.Equal(t, tools.StringSliceToSet([]string{"wan"}), model.Members)
	assert.Equal(t, types.StringValue("0"), model.NoGroup)
	assert.Equal(t, types.Int64Value(100), model.Sequence)
	assert.Equal(t, types.StringValue("test-description"), model.Description)
}

func TestConvertGroupSchemaToStruct(t *testing.T) {
	model := &groupResourceModel{
		Name:        types.StringValue("test_group"),
		Members:     tools.StringSliceToSet([]string{"wan"}),
		Sequence:    types.Int64Value(100),
		NoGroup:     types.StringValue("0"),
		Description: types.StringValue("test-description"),
	}

	upstream, err := convertGroupSchemaToStruct(model)

	require.NoError(t, err)
	require.NotNil(t, upstream)

	assert.Equal(t, "test_group", upstream.Name)
	assert.Equal(t, api.SelectedMapList{"wan"}, upstream.Members)
	assert.Equal(t, "100", upstream.Sequence)
	assert.Equal(t, "0", upstream.NoGroup)
	assert.Equal(t, "test-description", upstream.Description)
}
