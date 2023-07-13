// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrServiceTimestamps(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "debug_datetime_localtime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "debug_datetime_msec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "debug_datetime_show_timezone", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "debug_datetime_year", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "debug_uptime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "debug_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "log_datetime_localtime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "log_datetime_msec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "log_datetime_show_timezone", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "log_datetime_year", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "log_uptime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_service_timestamps.test", "log_disable", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrServiceTimestampsConfig_minimum(),
			},
			{
				Config: testAccIosxrServiceTimestampsConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_service_timestamps.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-service-timestamps-cfg:/service/timestamps",
			},
		},
	})
}

func testAccIosxrServiceTimestampsConfig_minimum() string {
	config := `resource "iosxr_service_timestamps" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrServiceTimestampsConfig_all() string {
	config := `resource "iosxr_service_timestamps" "test" {` + "\n"
	config += `	debug_datetime_localtime = true` + "\n"
	config += `	debug_datetime_msec = true` + "\n"
	config += `	debug_datetime_show_timezone = true` + "\n"
	config += `	debug_datetime_year = true` + "\n"
	config += `	debug_uptime = true` + "\n"
	config += `	debug_disable = true` + "\n"
	config += `	log_datetime_localtime = true` + "\n"
	config += `	log_datetime_msec = true` + "\n"
	config += `	log_datetime_show_timezone = true` + "\n"
	config += `	log_datetime_year = true` + "\n"
	config += `	log_uptime = true` + "\n"
	config += `	log_disable = true` + "\n"
	config += `}` + "\n"
	return config
}
