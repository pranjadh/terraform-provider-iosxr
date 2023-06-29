// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrLACP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrLACPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_lacp.test", "mac", "00:11:00:11:00:11"),
					resource.TestCheckResourceAttr("iosxr_lacp.test", "priority", "1"),
				),
			},
			{
				ResourceName:  "iosxr_lacp.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-lacp-cfg:/lacp/system",
			},
		},
	})
}

func testAccIosxrLACPConfig_minimum() string {
	return `
	resource "iosxr_lacp" "test" {
	}
	`
}

func testAccIosxrLACPConfig_all() string {
	return `
	resource "iosxr_lacp" "test" {
		mac = "00:11:00:11:00:11"
		priority = 1
	}
	`
}