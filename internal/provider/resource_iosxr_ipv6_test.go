// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrIPv6(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "hop_limit", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "icmp_error_interval", "2111"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "icmp_error_interval_bucket_size", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "source_route", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_timeout", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_max_packets", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_reassembler_drop_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_frag_hdr_incomplete_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_overlap_frag_drop_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "path_mtu_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "path_mtu_timeout", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrIPv6Config_minimum(),
			},
			{
				Config: testAccIosxrIPv6Config_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_ipv6.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-ipv6-cfg:/ipv6",
			},
		},
	})
}

func testAccIosxrIPv6Config_minimum() string {
	config := `resource "iosxr_ipv6" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrIPv6Config_all() string {
	config := `resource "iosxr_ipv6" "test" {` + "\n"
	config += `	hop_limit = 123` + "\n"
	config += `	icmp_error_interval = 2111` + "\n"
	config += `	icmp_error_interval_bucket_size = 123` + "\n"
	config += `	source_route = true` + "\n"
	config += `	assembler_timeout = 50` + "\n"
	config += `	assembler_max_packets = 40` + "\n"
	config += `	assembler_reassembler_drop_enable = true` + "\n"
	config += `	assembler_frag_hdr_incomplete_enable = true` + "\n"
	config += `	assembler_overlap_frag_drop_enable = true` + "\n"
	config += `	path_mtu_enable = true` + "\n"
	config += `	path_mtu_timeout = 10` + "\n"
	config += `}` + "\n"
	return config
}
