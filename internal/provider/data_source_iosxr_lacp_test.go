// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrLACP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_lacp.test", "mac", "00:11:00:11:00:11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_lacp.test", "priority", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrLACPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrLACPConfig() string {
	config := `resource "iosxr_lacp" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	mac = "00:11:00:11:00:11"` + "\n"
	config += `	priority = 1` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_lacp" "test" {
			depends_on = [iosxr_lacp.test]
		}
	`
	return config
}
