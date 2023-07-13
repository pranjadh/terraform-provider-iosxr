// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrEVPN(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_evpn.test", "source_interface", "Loopback0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrEVPNConfig_minimum(),
			},
			{
				Config: testAccIosxrEVPNConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_evpn.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/evpn",
			},
		},
	})
}

func testAccIosxrEVPNConfig_minimum() string {
	config := `resource "iosxr_evpn" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrEVPNConfig_all() string {
	config := `resource "iosxr_evpn" "test" {` + "\n"
	config += `	source_interface = "Loopback0"` + "\n"
	config += `}` + "\n"
	return config
}
