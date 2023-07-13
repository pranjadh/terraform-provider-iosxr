// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrL2VPN(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_l2vpn.test", "description", "My L2VPN Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_l2vpn.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_l2vpn.test", "xconnect_groups.0.group_name", "P2P"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrL2VPNConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrL2VPNConfig() string {
	config := `resource "iosxr_l2vpn" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	description = "My L2VPN Description"` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	xconnect_groups = [{` + "\n"
	config += `		group_name = "P2P"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_l2vpn" "test" {
			depends_on = [iosxr_l2vpn.test]
		}
	`
	return config
}
