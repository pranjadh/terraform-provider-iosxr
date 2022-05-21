// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type RouterOSPFVRF struct {
	Device                                types.String `tfsdk:"device"`
	Id                                    types.String `tfsdk:"id"`
	ProcessName                           types.String `tfsdk:"process_name"`
	VrfName                               types.String `tfsdk:"vrf_name"`
	MplsLdpSync                           types.Bool   `tfsdk:"mpls_ldp_sync"`
	HelloInterval                         types.Int64  `tfsdk:"hello_interval"`
	DeadInterval                          types.Int64  `tfsdk:"dead_interval"`
	Priority                              types.Int64  `tfsdk:"priority"`
	MtuIgnoreEnable                       types.Bool   `tfsdk:"mtu_ignore_enable"`
	MtuIgnoreDisable                      types.Bool   `tfsdk:"mtu_ignore_disable"`
	PassiveEnable                         types.Bool   `tfsdk:"passive_enable"`
	PassiveDisable                        types.Bool   `tfsdk:"passive_disable"`
	RouterId                              types.String `tfsdk:"router_id"`
	RedistributeConnected                 types.Bool   `tfsdk:"redistribute_connected"`
	RedistributeConnectedTag              types.Int64  `tfsdk:"redistribute_connected_tag"`
	RedistributeConnectedMetricType       types.String `tfsdk:"redistribute_connected_metric_type"`
	RedistributeStatic                    types.Bool   `tfsdk:"redistribute_static"`
	RedistributeStaticTag                 types.Int64  `tfsdk:"redistribute_static_tag"`
	RedistributeStaticMetricType          types.String `tfsdk:"redistribute_static_metric_type"`
	BfdFastDetect                         types.Bool   `tfsdk:"bfd_fast_detect"`
	BfdMinimumInterval                    types.Int64  `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier                         types.Int64  `tfsdk:"bfd_multiplier"`
	DefaultInformationOriginate           types.Bool   `tfsdk:"default_information_originate"`
	DefaultInformationOriginateAlways     types.Bool   `tfsdk:"default_information_originate_always"`
	DefaultInformationOriginateMetricType types.Int64  `tfsdk:"default_information_originate_metric_type"`
}

func (data RouterOSPFVRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=%s]/vrfs/vrf[vrf-name=%s]", data.ProcessName.Value, data.VrfName.Value)
}

func (data RouterOSPFVRF) toBody() string {
	body := "{}"
	if !data.MplsLdpSync.Null && !data.MplsLdpSync.Unknown {
		if data.MplsLdpSync.Value {
			body, _ = sjson.Set(body, "mpls.ldp.sync", map[string]string{})
		}
	}
	if !data.HelloInterval.Null && !data.HelloInterval.Unknown {
		body, _ = sjson.Set(body, "hello-interval", strconv.FormatInt(data.HelloInterval.Value, 10))
	}
	if !data.DeadInterval.Null && !data.DeadInterval.Unknown {
		body, _ = sjson.Set(body, "dead-interval", strconv.FormatInt(data.DeadInterval.Value, 10))
	}
	if !data.Priority.Null && !data.Priority.Unknown {
		body, _ = sjson.Set(body, "priority", strconv.FormatInt(data.Priority.Value, 10))
	}
	if !data.MtuIgnoreEnable.Null && !data.MtuIgnoreEnable.Unknown {
		if data.MtuIgnoreEnable.Value {
			body, _ = sjson.Set(body, "mtu-ignore.enable", map[string]string{})
		}
	}
	if !data.MtuIgnoreDisable.Null && !data.MtuIgnoreDisable.Unknown {
		if data.MtuIgnoreDisable.Value {
			body, _ = sjson.Set(body, "mtu-ignore.disable", map[string]string{})
		}
	}
	if !data.PassiveEnable.Null && !data.PassiveEnable.Unknown {
		if data.PassiveEnable.Value {
			body, _ = sjson.Set(body, "passive.enable", map[string]string{})
		}
	}
	if !data.PassiveDisable.Null && !data.PassiveDisable.Unknown {
		if data.PassiveDisable.Value {
			body, _ = sjson.Set(body, "passive.disable", map[string]string{})
		}
	}
	if !data.RouterId.Null && !data.RouterId.Unknown {
		body, _ = sjson.Set(body, "router-id", data.RouterId.Value)
	}
	if !data.RedistributeConnected.Null && !data.RedistributeConnected.Unknown {
		if data.RedistributeConnected.Value {
			body, _ = sjson.Set(body, "redistribute.connected", map[string]string{})
		}
	}
	if !data.RedistributeConnectedTag.Null && !data.RedistributeConnectedTag.Unknown {
		body, _ = sjson.Set(body, "redistribute.connected.tag", strconv.FormatInt(data.RedistributeConnectedTag.Value, 10))
	}
	if !data.RedistributeConnectedMetricType.Null && !data.RedistributeConnectedMetricType.Unknown {
		body, _ = sjson.Set(body, "redistribute.connected.metric-type", data.RedistributeConnectedMetricType.Value)
	}
	if !data.RedistributeStatic.Null && !data.RedistributeStatic.Unknown {
		if data.RedistributeStatic.Value {
			body, _ = sjson.Set(body, "redistribute.static", map[string]string{})
		}
	}
	if !data.RedistributeStaticTag.Null && !data.RedistributeStaticTag.Unknown {
		body, _ = sjson.Set(body, "redistribute.static.tag", strconv.FormatInt(data.RedistributeStaticTag.Value, 10))
	}
	if !data.RedistributeStaticMetricType.Null && !data.RedistributeStaticMetricType.Unknown {
		body, _ = sjson.Set(body, "redistribute.static.metric-type", data.RedistributeStaticMetricType.Value)
	}
	if !data.BfdFastDetect.Null && !data.BfdFastDetect.Unknown {
		if data.BfdFastDetect.Value {
			body, _ = sjson.Set(body, "bfd.fast-detect", map[string]string{})
		}
	}
	if !data.BfdMinimumInterval.Null && !data.BfdMinimumInterval.Unknown {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.Value, 10))
	}
	if !data.BfdMultiplier.Null && !data.BfdMultiplier.Unknown {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.Value, 10))
	}
	if !data.DefaultInformationOriginate.Null && !data.DefaultInformationOriginate.Unknown {
		if data.DefaultInformationOriginate.Value {
			body, _ = sjson.Set(body, "default-information.originate", map[string]string{})
		}
	}
	if !data.DefaultInformationOriginateAlways.Null && !data.DefaultInformationOriginateAlways.Unknown {
		if data.DefaultInformationOriginateAlways.Value {
			body, _ = sjson.Set(body, "default-information.originate.always", map[string]string{})
		}
	}
	if !data.DefaultInformationOriginateMetricType.Null && !data.DefaultInformationOriginateMetricType.Unknown {
		body, _ = sjson.Set(body, "default-information.originate.metric-type", strconv.FormatInt(data.DefaultInformationOriginateMetricType.Value, 10))
	}
	return body
}

func (data *RouterOSPFVRF) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "mpls.ldp.sync"); value.Exists() {
		data.MplsLdpSync.Value = true
	} else {
		data.MplsLdpSync.Value = false
	}
	if value := gjson.GetBytes(res, "hello-interval"); value.Exists() {
		data.HelloInterval.Value = value.Int()
	} else {
		data.HelloInterval.Null = true
	}
	if value := gjson.GetBytes(res, "dead-interval"); value.Exists() {
		data.DeadInterval.Value = value.Int()
	} else {
		data.DeadInterval.Null = true
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority.Value = value.Int()
	} else {
		data.Priority.Null = true
	}
	if value := gjson.GetBytes(res, "mtu-ignore.enable"); value.Exists() {
		data.MtuIgnoreEnable.Value = true
	} else {
		data.MtuIgnoreEnable.Value = false
	}
	if value := gjson.GetBytes(res, "mtu-ignore.disable"); value.Exists() {
		data.MtuIgnoreDisable.Value = true
	} else {
		data.MtuIgnoreDisable.Value = false
	}
	if value := gjson.GetBytes(res, "passive.enable"); value.Exists() {
		data.PassiveEnable.Value = true
	} else {
		data.PassiveEnable.Value = false
	}
	if value := gjson.GetBytes(res, "passive.disable"); value.Exists() {
		data.PassiveDisable.Value = true
	} else {
		data.PassiveDisable.Value = false
	}
	if value := gjson.GetBytes(res, "router-id"); value.Exists() {
		data.RouterId.Value = value.String()
	} else {
		data.RouterId.Null = true
	}
	if value := gjson.GetBytes(res, "redistribute.connected"); value.Exists() {
		data.RedistributeConnected.Value = true
	} else {
		data.RedistributeConnected.Value = false
	}
	if value := gjson.GetBytes(res, "redistribute.connected.tag"); value.Exists() {
		data.RedistributeConnectedTag.Value = value.Int()
	} else {
		data.RedistributeConnectedTag.Null = true
	}
	if value := gjson.GetBytes(res, "redistribute.connected.metric-type"); value.Exists() {
		data.RedistributeConnectedMetricType.Value = value.String()
	} else {
		data.RedistributeConnectedMetricType.Null = true
	}
	if value := gjson.GetBytes(res, "redistribute.static"); value.Exists() {
		data.RedistributeStatic.Value = true
	} else {
		data.RedistributeStatic.Value = false
	}
	if value := gjson.GetBytes(res, "redistribute.static.tag"); value.Exists() {
		data.RedistributeStaticTag.Value = value.Int()
	} else {
		data.RedistributeStaticTag.Null = true
	}
	if value := gjson.GetBytes(res, "redistribute.static.metric-type"); value.Exists() {
		data.RedistributeStaticMetricType.Value = value.String()
	} else {
		data.RedistributeStaticMetricType.Null = true
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect"); value.Exists() {
		data.BfdFastDetect.Value = true
	} else {
		data.BfdFastDetect.Value = false
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval.Value = value.Int()
	} else {
		data.BfdMinimumInterval.Null = true
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier.Value = value.Int()
	} else {
		data.BfdMultiplier.Null = true
	}
	if value := gjson.GetBytes(res, "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate.Value = true
	} else {
		data.DefaultInformationOriginate.Value = false
	}
	if value := gjson.GetBytes(res, "default-information.originate.always"); value.Exists() {
		data.DefaultInformationOriginateAlways.Value = true
	} else {
		data.DefaultInformationOriginateAlways.Value = false
	}
	if value := gjson.GetBytes(res, "default-information.originate.metric-type"); value.Exists() {
		data.DefaultInformationOriginateMetricType.Value = value.Int()
	} else {
		data.DefaultInformationOriginateMetricType.Null = true
	}
}

func (data *RouterOSPFVRF) fromPlan(plan RouterOSPFVRF) {
	data.Device = plan.Device
	data.ProcessName.Value = plan.ProcessName.Value
	data.VrfName.Value = plan.VrfName.Value
}

func (data *RouterOSPFVRF) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.ProcessName.Unknown {
		data.ProcessName.Unknown = false
		data.ProcessName.Null = true
	}
	if data.VrfName.Unknown {
		data.VrfName.Unknown = false
		data.VrfName.Null = true
	}
	if data.MplsLdpSync.Unknown {
		data.MplsLdpSync.Unknown = false
		data.MplsLdpSync.Null = true
	}
	if data.HelloInterval.Unknown {
		data.HelloInterval.Unknown = false
		data.HelloInterval.Null = true
	}
	if data.DeadInterval.Unknown {
		data.DeadInterval.Unknown = false
		data.DeadInterval.Null = true
	}
	if data.Priority.Unknown {
		data.Priority.Unknown = false
		data.Priority.Null = true
	}
	if data.MtuIgnoreEnable.Unknown {
		data.MtuIgnoreEnable.Unknown = false
		data.MtuIgnoreEnable.Null = true
	}
	if data.MtuIgnoreDisable.Unknown {
		data.MtuIgnoreDisable.Unknown = false
		data.MtuIgnoreDisable.Null = true
	}
	if data.PassiveEnable.Unknown {
		data.PassiveEnable.Unknown = false
		data.PassiveEnable.Null = true
	}
	if data.PassiveDisable.Unknown {
		data.PassiveDisable.Unknown = false
		data.PassiveDisable.Null = true
	}
	if data.RouterId.Unknown {
		data.RouterId.Unknown = false
		data.RouterId.Null = true
	}
	if data.RedistributeConnected.Unknown {
		data.RedistributeConnected.Unknown = false
		data.RedistributeConnected.Null = true
	}
	if data.RedistributeConnectedTag.Unknown {
		data.RedistributeConnectedTag.Unknown = false
		data.RedistributeConnectedTag.Null = true
	}
	if data.RedistributeConnectedMetricType.Unknown {
		data.RedistributeConnectedMetricType.Unknown = false
		data.RedistributeConnectedMetricType.Null = true
	}
	if data.RedistributeStatic.Unknown {
		data.RedistributeStatic.Unknown = false
		data.RedistributeStatic.Null = true
	}
	if data.RedistributeStaticTag.Unknown {
		data.RedistributeStaticTag.Unknown = false
		data.RedistributeStaticTag.Null = true
	}
	if data.RedistributeStaticMetricType.Unknown {
		data.RedistributeStaticMetricType.Unknown = false
		data.RedistributeStaticMetricType.Null = true
	}
	if data.BfdFastDetect.Unknown {
		data.BfdFastDetect.Unknown = false
		data.BfdFastDetect.Null = true
	}
	if data.BfdMinimumInterval.Unknown {
		data.BfdMinimumInterval.Unknown = false
		data.BfdMinimumInterval.Null = true
	}
	if data.BfdMultiplier.Unknown {
		data.BfdMultiplier.Unknown = false
		data.BfdMultiplier.Null = true
	}
	if data.DefaultInformationOriginate.Unknown {
		data.DefaultInformationOriginate.Unknown = false
		data.DefaultInformationOriginate.Null = true
	}
	if data.DefaultInformationOriginateAlways.Unknown {
		data.DefaultInformationOriginateAlways.Unknown = false
		data.DefaultInformationOriginateAlways.Null = true
	}
	if data.DefaultInformationOriginateMetricType.Unknown {
		data.DefaultInformationOriginateMetricType.Unknown = false
		data.DefaultInformationOriginateMetricType.Null = true
	}
}

func (data *RouterOSPFVRF) getDeletedListItems(state RouterOSPFVRF) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *RouterOSPFVRF) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
