package firewall_test

/*
Depends on the new testing framework at https://github.com/browningluke/terraform-provider-opnsense/pull/186/

import (
	"fmt"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGroupDataSource(t *testing.T) {
	acctest.AccPreCheck(t)
	client := acctest.Client(t)

	const (
		testName        = "group_example"
		testSequence    = "100"
		testNoGroup     = "0"
		testDescription = "terraform-provider-opnsense acceptance test"
	)

	group := &firewall.Group{
		Name:        testName,
		Members:     api.SelectedMapList{"wan"},
		Sequence:    testSequence,
		NoGroup:     testNoGroup,
		Description: testDescription,
	}

	id := acctest.CreateDataSourceTestResource(t, acctest.ResourceLifecycle[*firewall.Group]{
		Create: client.Firewall().AddGroup,
		Delete: client.Firewall().DeleteGroup,
	}, group)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupDataSourceConfig(id),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "id", id),
					resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "name", testName),
					resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "members.#", "1"),
					resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "members.0", "wan"),
					resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "sequence", testSequence),
					resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "no_group", testNoGroup),
					resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "description", testDescription),
				),
			},
		},
	})
}

func testAccGroupDataSourceConfig(id string) string {
	return fmt.Sprintf(`
data "opnsense_firewall_group" "test" {
  id = "%s"
}
`, id)
}
*/