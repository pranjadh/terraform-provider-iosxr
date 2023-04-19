// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSNMPServer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSNMPServerConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "rf", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "bfd", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "ntp", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "ethernet_oam_events", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "copy_complete", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "traps_snmp_linkup", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "traps_snmp_linkdown", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "power", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "config", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "entity", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "system", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "bridgemib", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "entity_state_operstatus", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "entity_redundancy_all", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "trap_source_both", "Loopback10"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "l2vpn_all", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "l2vpn_vc_up", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "l2vpn_vc_down", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "sensor", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "fru_ctrl", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "isis_authentication_failure", "enable"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "bgp_cbgp2_updown", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "bgp_bgp4_mib_updown", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "users.0.user_name", "USER1"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "users.0.group_name", "GROUP1"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "users.0.v3_auth_md5_encryption_aes", "073C05626E2A4841141D"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.group_name", "GROUP12"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.v3_priv", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.v3_read", "VIEW1"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.v3_write", "VIEW2"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.v3_context", "CONTEXT1"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.v3_notify", "VIEW3"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.v3_ipv4", "ACL1"),
					resource.TestCheckResourceAttr("iosxr_snmp_server.test", "groups.0.v3_ipv6", "ACL2"),
				),
			},
			{
				ResourceName:  "iosxr_snmp_server.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-snmp-server-cfg:snmp-server",
			},
		},
	})
}

func testAccIosxrSNMPServerConfig_minimum() string {
	return `
	resource "iosxr_snmp_server" "test" {
	}
	`
}

func testAccIosxrSNMPServerConfig_all() string {
	return `
	resource "iosxr_snmp_server" "test" {
		rf = true
		bfd = true
		ntp = true
		ethernet_oam_events = true
		copy_complete = true
		traps_snmp_linkup = true
		traps_snmp_linkdown = true
		power = true
		config = true
		entity = true
		system = true
		bridgemib = true
		entity_state_operstatus = true
		entity_redundancy_all = true
		trap_source_both = "Loopback10"
		l2vpn_all = true
		l2vpn_vc_up = true
		l2vpn_vc_down = true
		sensor = true
		fru_ctrl = true
		isis_authentication_failure = "enable"
		bgp_cbgp2_updown = true
		bgp_bgp4_mib_updown = true
		users = [{
			user_name = "USER1"
			group_name = "GROUP1"
			v3_auth_md5_encryption_aes = "073C05626E2A4841141D"
		}]
		groups = [{
			group_name = "GROUP12"
			v3_priv = true
			v3_read = "VIEW1"
			v3_write = "VIEW2"
			v3_context = "CONTEXT1"
			v3_notify = "VIEW3"
			v3_ipv4 = "ACL1"
			v3_ipv6 = "ACL2"
		}]
	}
	`
}
