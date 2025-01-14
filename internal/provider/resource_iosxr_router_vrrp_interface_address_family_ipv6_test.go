// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterVRRPInterfaceAddressFamilyIPv6(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "vrrp_id", "124"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "address_linklocal_autoconfig", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "priority", "250"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "name", "TEST2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "timer_advertisement_seconds", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "timer_force", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "preempt_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "preempt_delay", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "accept_mode_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_interfaces.0.interface_name", "GigabitEthernet0/0/0/5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_interfaces.0.priority_decrement", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_objects.0.object_name", "OBJECT"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_objects.0.priority_decrement", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "bfd_fast_detect_peer_ipv6", "3::3"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterVRRPInterfaceAddressFamilyIPv6Config_minimum(),
			},
			{
				Config: testAccIosxrRouterVRRPInterfaceAddressFamilyIPv6Config_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_router_vrrp_interface_address_family_ipv6.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=GigabitEthernet0/0/0/2]/address-family/ipv6/vrrps/vrrp[vrrp-id=%!d(string=124)]",
			},
		},
	})
}

func testAccIosxrRouterVRRPInterfaceAddressFamilyIPv6Config_minimum() string {
	config := `resource "iosxr_router_vrrp_interface_address_family_ipv6" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/2"` + "\n"
	config += `	vrrp_id = 124` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterVRRPInterfaceAddressFamilyIPv6Config_all() string {
	config := `resource "iosxr_router_vrrp_interface_address_family_ipv6" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/2"` + "\n"
	config += `	vrrp_id = 124` + "\n"
	config += `	address_linklocal_autoconfig = true` + "\n"
	config += `	priority = 250` + "\n"
	config += `	name = "TEST2"` + "\n"
	config += `	timer_advertisement_seconds = 10` + "\n"
	config += `	timer_force = true` + "\n"
	config += `	preempt_disable = false` + "\n"
	config += `	preempt_delay = 255` + "\n"
	config += `	accept_mode_disable = true` + "\n"
	config += `	track_interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/5"` + "\n"
	config += `		priority_decrement = 12` + "\n"
	config += `	}]` + "\n"
	config += `	track_objects = [{` + "\n"
	config += `		object_name = "OBJECT"` + "\n"
	config += `		priority_decrement = 22` + "\n"
	config += `	}]` + "\n"
	config += `	bfd_fast_detect_peer_ipv6 = "3::3"` + "\n"
	config += `}` + "\n"
	return config
}
