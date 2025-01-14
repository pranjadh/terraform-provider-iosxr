// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrEVPNInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "core_isolation_group", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_identifier_type_zero_bytes_1", "01"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_identifier_type_zero_bytes_23", "0100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_identifier_type_zero_bytes_45", "0100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_identifier_type_zero_bytes_67", "0100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_identifier_type_zero_bytes_89", "0100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_all_active", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_port_active", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_single_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_single_flow_active", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrEVPNInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrEVPNInterfaceConfig() string {
	config := `resource "iosxr_evpn_interface" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	interface_name = "Bundle-Ether12"` + "\n"
	config += `	core_isolation_group = 11` + "\n"
	config += `	ethernet_segment_identifier_type_zero_bytes_1 = "01"` + "\n"
	config += `	ethernet_segment_identifier_type_zero_bytes_23 = "0100"` + "\n"
	config += `	ethernet_segment_identifier_type_zero_bytes_45 = "0100"` + "\n"
	config += `	ethernet_segment_identifier_type_zero_bytes_67 = "0100"` + "\n"
	config += `	ethernet_segment_identifier_type_zero_bytes_89 = "0100"` + "\n"
	config += `	ethernet_segment_load_balancing_mode_all_active = false` + "\n"
	config += `	ethernet_segment_load_balancing_mode_port_active = false` + "\n"
	config += `	ethernet_segment_load_balancing_mode_single_active = true` + "\n"
	config += `	ethernet_segment_load_balancing_mode_single_flow_active = false` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_evpn_interface" "test" {
			interface_name = "Bundle-Ether12"
			depends_on = [iosxr_evpn_interface.test]
		}
	`
	return config
}
