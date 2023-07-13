// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrIPv6(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "hop_limit", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "icmp_error_interval", "2111"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "icmp_error_interval_bucket_size", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "source_route", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "assembler_timeout", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "assembler_max_packets", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "assembler_reassembler_drop_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "assembler_frag_hdr_incomplete_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "assembler_overlap_frag_drop_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "path_mtu_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6.test", "path_mtu_timeout", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrIPv6Config(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrIPv6Config() string {
	config := `resource "iosxr_ipv6" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
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

	config += `
		data "iosxr_ipv6" "test" {
			depends_on = [iosxr_ipv6.test]
		}
	`
	return config
}
