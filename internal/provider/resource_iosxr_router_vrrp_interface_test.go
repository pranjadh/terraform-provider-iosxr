// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterVRRPInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "mac_refresh", "14"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "delay_minimum", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "delay_reload", "4321"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "bfd_minimum_interval", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "bfd_multiplier", "33"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterVRRPInterfaceConfig_minimum(),
			},
			{
				Config: testAccIosxrRouterVRRPInterfaceConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_router_vrrp_interface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]",
			},
		},
	})
}

func testAccIosxrRouterVRRPInterfaceConfig_minimum() string {
	config := `resource "iosxr_router_vrrp_interface" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterVRRPInterfaceConfig_all() string {
	config := `resource "iosxr_router_vrrp_interface" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	mac_refresh = 14` + "\n"
	config += `	delay_minimum = 1234` + "\n"
	config += `	delay_reload = 4321` + "\n"
	config += `	bfd_minimum_interval = 255` + "\n"
	config += `	bfd_multiplier = 33` + "\n"
	config += `}` + "\n"
	return config
}
