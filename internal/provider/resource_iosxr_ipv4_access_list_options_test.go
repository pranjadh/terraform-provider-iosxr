// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrIPv4AccessListOptions(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list_options.test", "log_update_threshold", "214748"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list_options.test", "log_update_rate", "1000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrIPv4AccessListOptionsConfig_minimum(),
			},
			{
				Config: testAccIosxrIPv4AccessListOptionsConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_ipv4_access_list_options.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-ipv4-access-list-cfg:/ipv4/access-list-options",
			},
		},
	})
}

func testAccIosxrIPv4AccessListOptionsConfig_minimum() string {
	config := `resource "iosxr_ipv4_access_list_options" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrIPv4AccessListOptionsConfig_all() string {
	config := `resource "iosxr_ipv4_access_list_options" "test" {` + "\n"
	config += `	log_update_threshold = 214748` + "\n"
	config += `	log_update_rate = 1000` + "\n"
	config += `}` + "\n"
	return config
}
