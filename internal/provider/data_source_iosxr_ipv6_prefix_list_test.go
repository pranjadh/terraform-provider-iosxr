// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrIPv6PrefixList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.sequence_number", "4096"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.remark", "REMARK"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.permission", "permit"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.prefix", "2001:db8:3333:4444:5555:6666:7777:8888"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.mask", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.match_prefix_length_eq", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.match_prefix_length_ge", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_ipv6_prefix_list.test", "sequences.0.match_prefix_length_le", "20"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrIPv6PrefixListConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrIPv6PrefixListConfig() string {
	config := `resource "iosxr_ipv6_prefix_list" "test" {` + "\n"
	config += `	prefix_list_name = "LIST1"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `		sequence_number = 4096` + "\n"
	config += `		remark = "REMARK"` + "\n"
	config += `		permission = "permit"` + "\n"
	config += `		prefix = "2001:db8:3333:4444:5555:6666:7777:8888"` + "\n"
	config += `		mask = 64` + "\n"
	config += `		match_prefix_length_eq = 10` + "\n"
	config += `		match_prefix_length_ge = 20` + "\n"
	config += `		match_prefix_length_le = 20` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_ipv6_prefix_list" "test" {
			prefix_list_name = "LIST1"
			depends_on = [iosxr_ipv6_prefix_list.test]
		}
	`
	return config
}
