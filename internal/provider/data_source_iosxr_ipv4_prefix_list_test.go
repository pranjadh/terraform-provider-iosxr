// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrIPv4PrefixList(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrIPv4PrefixListConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.sequence_number", "4096"),
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.remark", "REMARK"),
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.permission", "deny"),
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.prefix", "10.1.1.1"),
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.mask", "255.255.0.0"),
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.match_prefix_length_eq", "12"),
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.match_prefix_length_ge", "22"),
					resource.TestCheckResourceAttr("data.iosxr_ipv4_prefix_list.test", "sequences.0.match_prefix_length_le", "32"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrIPv4PrefixListConfig = `

resource "iosxr_ipv4_prefix_list" "test" {
	prefix_list_name = "LIST1"
	sequences = [{
		sequence_number = 4096
		remark = "REMARK"
		permission = "deny"
		prefix = "10.1.1.1"
		mask = "255.255.0.0"
		match_prefix_length_eq = 12
		match_prefix_length_ge = 22
		match_prefix_length_le = 32
	}]
}

data "iosxr_ipv4_prefix_list" "test" {
	prefix_list_name = "LIST1"
	depends_on = [iosxr_ipv4_prefix_list.test]
}
`