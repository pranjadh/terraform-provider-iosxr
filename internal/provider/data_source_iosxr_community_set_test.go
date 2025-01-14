// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrCommunitySet(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_community_set.test", "rpl", "community-set TEST11\nend-set\n"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrCommunitySetConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrCommunitySetConfig() string {
	config := `resource "iosxr_community_set" "test" {` + "\n"
	config += `	set_name = "TEST11"` + "\n"
	config += `	rpl = "community-set TEST11\nend-set\n"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_community_set" "test" {
			set_name = "TEST11"
			depends_on = [iosxr_community_set.test]
		}
	`
	return config
}
