// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrSNMPServerMIB(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server_mib.test", "ifmib_ifalias_long", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server_mib.test", "ifindex_persist", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSNMPServerMIBConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrSNMPServerMIBConfig() string {
	config := `resource "iosxr_snmp_server_mib" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	ifmib_ifalias_long = true` + "\n"
	config += `	ifindex_persist = true` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_snmp_server_mib" "test" {
			depends_on = [iosxr_snmp_server_mib.test]
		}
	`
	return config
}
