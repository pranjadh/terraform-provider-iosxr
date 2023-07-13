// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrLoggingVRF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv4_addresses.0.ipv4_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv4_addresses.0.severity", "info"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv6_addresses.0.ipv6_address", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_logging_vrf.test", "host_ipv6_addresses.0.severity", "info"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrLoggingVRFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrLoggingVRFConfig() string {
	config := `resource "iosxr_logging_vrf" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	host_ipv4_addresses = [{` + "\n"
	config += `		ipv4_address = "1.1.1.1"` + "\n"
	config += `		severity = "info"` + "\n"
	config += `	}]` + "\n"
	config += `	host_ipv6_addresses = [{` + "\n"
	config += `		ipv6_address = "2001::1"` + "\n"
	config += `		severity = "info"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_logging_vrf" "test" {
			vrf_name = "VRF1"
			depends_on = [iosxr_logging_vrf.test]
		}
	`
	return config
}
