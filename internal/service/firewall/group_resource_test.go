package firewall_test

import (
	"strconv"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFirewallGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: testAccFirewallGroupResourceConfig(
					"tfacc_group",
					`["wan"]`,
					"0",
					100,
					"test-description",
				),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "name", "tfacc_group"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "members.#", "1"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "members.0", "wan"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "sequence", "100"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "no_group", "0"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "description", "test-description"),
					resource.TestCheckResourceAttrSet("opnsense_firewall_group.test", "id"),
				),
			},

			{
				ResourceName:      "opnsense_firewall_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},

			{
				Config: testAccFirewallGroupResourceConfig(
					"tfacc_changed",
					`["wan"]`,
					"1",
					200,
					"test-description-updated",
				),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "name", "tfacc_changed"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "members.#", "1"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "members.0", "wan"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "sequence", "200"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "no_group", "1"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "description", "test-description-updated"),
					resource.TestCheckResourceAttrSet("opnsense_firewall_group.test", "id"),
				),
			},

			{
				ResourceName:      "opnsense_firewall_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFirewallGroupResourceConfig(
	name, members, noGroup string,
	sequence int,
	description string,
) string {
	return `
resource "opnsense_firewall_group" "test" {
	name        = "` + name + `"
	members     = ` + members + `
	no_group    = "` + noGroup + `"
	sequence    = ` + strconv.Itoa(sequence) + `
	description = "` + description + `"
}
`
}
