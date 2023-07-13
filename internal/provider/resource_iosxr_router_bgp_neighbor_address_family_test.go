// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterBGPNeighborAddressFamily(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "af_name", "vpnv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "import_stitching_rt_re_originate_stitching_rt", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "route_reflector_client_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "advertise_vpnv4_unicast_enable_re_originated_stitching_rt", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "next_hop_self_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "encapsulation_type_srv6", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPNeighborAddressFamilyConfig_minimum(),
			},
			{
				Config: testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPNeighborAddressFamilyConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_router_bgp_neighbor_address_family.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/neighbors/neighbor[neighbor-address=10.1.1.2]/address-families/address-family[af-name=vpnv4-unicast]",
			},
		},
	})
}

const testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
	lists = [
		{
			name = "address-families/address-family"
			key = "af-name"
			items = [
				{
					"af-name" = "vpnv4-unicast"
				},
			] 
		},
		{
			name = "neighbors/neighbor"
			key = "neighbor-address"
			items = [
				{
					"neighbor-address" = "10.1.1.2"
					"remote-as" = "65001"
				},
			] 
		},
	]
}

`

func testAccIosxrRouterBGPNeighborAddressFamilyConfig_minimum() string {
	config := `resource "iosxr_router_bgp_neighbor_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	neighbor_address = "10.1.1.2"` + "\n"
	config += `	af_name = "vpnv4-unicast"` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterBGPNeighborAddressFamilyConfig_all() string {
	config := `resource "iosxr_router_bgp_neighbor_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	neighbor_address = "10.1.1.2"` + "\n"
	config += `	af_name = "vpnv4-unicast"` + "\n"
	config += `	import_stitching_rt_re_originate_stitching_rt = true` + "\n"
	config += `	route_reflector_client_inheritance_disable = true` + "\n"
	config += `	advertise_vpnv4_unicast_enable_re_originated_stitching_rt = true` + "\n"
	config += `	next_hop_self_inheritance_disable = true` + "\n"
	config += `	encapsulation_type_srv6 = true` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
