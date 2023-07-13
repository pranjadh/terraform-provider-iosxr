// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrL2VPNBridgeGroupBridgeDomain(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn_bridge_group_bridge_domain.test", "bridge_domain_name", "BD123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn_bridge_group_bridge_domain.test", "evis.0.vpn_id", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn_bridge_group_bridge_domain.test", "vnis.0.vni_id", "1234"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_minimum(),
			},
			{
				Config: testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_l2vpn_bridge_group_bridge_domain.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/bridge/groups/group[group-name=BG123]/bridge-domains/bridge-domain[bridge-domain-name=BD123]",
			},
		},
	})
}

func testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_minimum() string {
	config := `resource "iosxr_l2vpn_bridge_group_bridge_domain" "test" {` + "\n"
	config += `	bridge_group_name = "BG123"` + "\n"
	config += `	bridge_domain_name = "BD123"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrL2VPNBridgeGroupBridgeDomainConfig_all() string {
	config := `resource "iosxr_l2vpn_bridge_group_bridge_domain" "test" {` + "\n"
	config += `	bridge_group_name = "BG123"` + "\n"
	config += `	bridge_domain_name = "BD123"` + "\n"
	config += `	evis = [{` + "\n"
	config += `		vpn_id = 1234` + "\n"
	config += `	}]` + "\n"
	config += `	vnis = [{` + "\n"
	config += `		vni_id = 1234` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
