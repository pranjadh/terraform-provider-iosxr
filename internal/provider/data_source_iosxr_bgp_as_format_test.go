// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrBGPASFormat(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_bgp_as_format.test", "asdot", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_bgp_as_format.test", "asplain", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrBGPASFormatConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrBGPASFormatConfig() string {
	config := `resource "iosxr_bgp_as_format" "test" {` + "\n"
	config += `	asdot = false` + "\n"
	config += `	asplain = true` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_bgp_as_format" "test" {
			depends_on = [iosxr_bgp_as_format.test]
		}
	`
	return config
}
