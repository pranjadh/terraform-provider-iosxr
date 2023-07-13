// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrRouterStaticIPv4Multicast(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interfaces.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interfaces.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interfaces.0.distance_metric", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interfaces.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interfaces.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interface_addresses.0.address", "11.11.11.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interface_addresses.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interface_addresses.0.tag", "103"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interface_addresses.0.distance_metric", "144"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interface_addresses.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_interface_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_addresses.0.address", "100.0.2.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_addresses.0.description", "ip-description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_addresses.0.tag", "104"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_addresses.0.distance_metric", "155"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_addresses.0.track", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "nexthop_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.vrf_name", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.distance_metric", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interfaces.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.address", "11.11.11.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.tag", "103"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.distance_metric", "144"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_interface_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.address", "100.0.2.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.description", "ip-description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.tag", "104"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.distance_metric", "155"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.track", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_static_ipv4_multicast.test", "vrfs.0.nexthop_addresses.0.metric", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterStaticIPv4MulticastConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrRouterStaticIPv4MulticastConfig() string {
	config := `resource "iosxr_router_static_ipv4_multicast" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	prefix_address = "100.0.1.0"` + "\n"
	config += `	prefix_length = 24` + "\n"
	config += `	nexthop_interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		description = "interface-description"` + "\n"
	config += `		tag = 100` + "\n"
	config += `		distance_metric = 122` + "\n"
	config += `		permanent = true` + "\n"
	config += `		metric = 10` + "\n"
	config += `	}]` + "\n"
	config += `	nexthop_interface_addresses = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/2"` + "\n"
	config += `		address = "11.11.11.1"` + "\n"
	config += `		description = "interface-description"` + "\n"
	config += `		tag = 103` + "\n"
	config += `		distance_metric = 144` + "\n"
	config += `		permanent = true` + "\n"
	config += `		metric = 10` + "\n"
	config += `	}]` + "\n"
	config += `	nexthop_addresses = [{` + "\n"
	config += `		address = "100.0.2.0"` + "\n"
	config += `		description = "ip-description"` + "\n"
	config += `		tag = 104` + "\n"
	config += `		distance_metric = 155` + "\n"
	config += `		track = "TRACK1"` + "\n"
	config += `		metric = 10` + "\n"
	config += `	}]` + "\n"
	config += `	vrfs = [{` + "\n"
	config += `		vrf_name = "VRF1"` + "\n"
	config += `		nexthop_interfaces = [{` + "\n"
	config += `			interface_name = "GigabitEthernet0/0/0/3"` + "\n"
	config += `			description = "interface-description"` + "\n"
	config += `			tag = 100` + "\n"
	config += `			distance_metric = 122` + "\n"
	config += `			permanent = true` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `		nexthop_interface_addresses = [{` + "\n"
	config += `			interface_name = "GigabitEthernet0/0/0/4"` + "\n"
	config += `			address = "11.11.11.1"` + "\n"
	config += `			description = "interface-description"` + "\n"
	config += `			tag = 103` + "\n"
	config += `			distance_metric = 144` + "\n"
	config += `			permanent = true` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `		nexthop_addresses = [{` + "\n"
	config += `			address = "100.0.2.0"` + "\n"
	config += `			description = "ip-description"` + "\n"
	config += `			tag = 104` + "\n"
	config += `			distance_metric = 155` + "\n"
	config += `			track = "TRACK1"` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_router_static_ipv4_multicast" "test" {
			prefix_address = "100.0.1.0"
			prefix_length = 24
			depends_on = [iosxr_router_static_ipv4_multicast.test]
		}
	`
	return config
}
