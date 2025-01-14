// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrSNMPServerVRFHost(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server_vrf_host.test", "unencrypted_strings.0.community_string", "COMMUNITY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server_vrf_host.test", "unencrypted_strings.0.version_v3_security_level", "auth"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSNMPServerVRFHostConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrSNMPServerVRFHostConfig() string {
	config := `resource "iosxr_snmp_server_vrf_host" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	address = "11.11.11.11"` + "\n"
	config += `	unencrypted_strings = [{` + "\n"
	config += `		community_string = "COMMUNITY1"` + "\n"
	config += `		version_v3_security_level = "auth"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_snmp_server_vrf_host" "test" {
			vrf_name = "VRF1"
			address = "11.11.11.11"
			depends_on = [iosxr_snmp_server_vrf_host.test]
		}
	`
	return config
}
