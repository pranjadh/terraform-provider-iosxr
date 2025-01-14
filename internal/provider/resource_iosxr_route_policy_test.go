// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRoutePolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_route_policy.test", "route_policy_name", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_route_policy.test", "rpl", "route-policy ROUTE_POLICY_1\n  if destination in PREFIX_SET_1 then\n    set extcommunity rt (12345:1) additive\n  endif\n  pass\nend-policy\n"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRoutePolicyConfig_minimum(),
			},
			{
				Config: testAccIosxrRoutePolicyConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_route_policy.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]",
			},
		},
	})
}

func testAccIosxrRoutePolicyConfig_minimum() string {
	config := `resource "iosxr_route_policy" "test" {` + "\n"
	config += `	route_policy_name = "ROUTE_POLICY_1"` + "\n"
	config += `	rpl = "route-policy ROUTE_POLICY_1\n  if destination in PREFIX_SET_1 then\n    set extcommunity rt (12345:1) additive\n  endif\n  pass\nend-policy\n"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRoutePolicyConfig_all() string {
	config := `resource "iosxr_route_policy" "test" {` + "\n"
	config += `	route_policy_name = "ROUTE_POLICY_1"` + "\n"
	config += `	rpl = "route-policy ROUTE_POLICY_1\n  if destination in PREFIX_SET_1 then\n    set extcommunity rt (12345:1) additive\n  endif\n  pass\nend-policy\n"` + "\n"
	config += `}` + "\n"
	return config
}
