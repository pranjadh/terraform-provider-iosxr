// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrRouterISISInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "circuit_type", "level-1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "hello_padding_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "hello_padding_sometimes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "point_to_point", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "passive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "suppressed", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "hello_password_keychain", "KEY_CHAIN_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_interface.test", "bfd_fast_detect_ipv6", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterISISInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrRouterISISInterfaceConfig() string {
	config := `resource "iosxr_router_isis_interface" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	process_id = "P1"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	circuit_type = "level-1"` + "\n"
	config += `	hello_padding_disable = true` + "\n"
	config += `	hello_padding_sometimes = false` + "\n"
	config += `	priority = 10` + "\n"
	config += `	point_to_point = false` + "\n"
	config += `	passive = false` + "\n"
	config += `	suppressed = false` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	hello_password_keychain = "KEY_CHAIN_1"` + "\n"
	config += `	bfd_fast_detect_ipv6 = true` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_router_isis_interface" "test" {
			process_id = "P1"
			interface_name = "GigabitEthernet0/0/0/1"
			depends_on = [iosxr_router_isis_interface.test]
		}
	`
	return config
}
