// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterISISInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterISISInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "circuit_type", "level-1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "hello_padding_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "hello_padding_sometimes", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "priority", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "point_to_point", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "passive", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "suppressed", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "hello_password_keychain", "KEY_CHAIN_1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterISISInterfaceConfig = `

resource "iosxr_router_isis_interface" "test" {
	process_id = "P1"
	interface_name = "GigabitEthernet0/0/0/1"
	circuit_type = "level-1"
	hello_padding_disable = true
	hello_padding_sometimes = false
	priority = 10
	point_to_point = false
	passive = false
	suppressed = false
	shutdown = false
	hello_password_keychain = "KEY_CHAIN_1"
}

data "iosxr_router_isis_interface" "test" {
	process_id = "P1"
	interface_name = "GigabitEthernet0/0/0/1"
	depends_on = [iosxr_router_isis_interface.test]
}
`