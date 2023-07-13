// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrPrefixSet(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_prefix_set.test", "rpl", "prefix-set PREFIX_SET_1\n  10.1.1.0/26 ge 26,\n  10.1.2.0/26 ge 26\nend-set\n"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrPrefixSetConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrPrefixSetConfig() string {
	config := `resource "iosxr_prefix_set" "test" {` + "\n"
	config += `	set_name = "PREFIX_SET_1"` + "\n"
	config += `	rpl = "prefix-set PREFIX_SET_1\n  10.1.1.0/26 ge 26,\n  10.1.2.0/26 ge 26\nend-set\n"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_prefix_set" "test" {
			set_name = "PREFIX_SET_1"
			depends_on = [iosxr_prefix_set.test]
		}
	`
	return config
}
