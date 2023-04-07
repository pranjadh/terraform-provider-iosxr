// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type L2VPNBridgeGroupBridgeDomain struct {
	Device                         types.String                             `tfsdk:"device"`
	Id                             types.String                             `tfsdk:"id"`
	BridgeGroupName                types.String                             `tfsdk:"bridge_group_name"`
	BridgeDomainName               types.String                             `tfsdk:"bridge_domain_name"`
	Evis                           []L2VPNBridgeGroupBridgeDomainEvis       `tfsdk:"evis"`
	Vnis                           []L2VPNBridgeGroupBridgeDomainVnis       `tfsdk:"vnis"`
	Mtu                            types.Int64                              `tfsdk:"mtu"`
	StormControlBroadcastPps       types.Int64                              `tfsdk:"storm_control_broadcast_pps"`
	StormControlBroadcastKbps      types.Int64                              `tfsdk:"storm_control_broadcast_kbps"`
	StormControlMulticastPps       types.Int64                              `tfsdk:"storm_control_multicast_pps"`
	StormControlMulticastKbps      types.Int64                              `tfsdk:"storm_control_multicast_kbps"`
	StormControlUnknownUnicastPps  types.Int64                              `tfsdk:"storm_control_unknown_unicast_pps"`
	StormControlUnknownUnicastKbps types.Int64                              `tfsdk:"storm_control_unknown_unicast_kbps"`
	Interfaces                     []L2VPNBridgeGroupBridgeDomainInterfaces `tfsdk:"interfaces"`
}
type L2VPNBridgeGroupBridgeDomainEvis struct {
	VpnId types.Int64 `tfsdk:"vpn_id"`
}
type L2VPNBridgeGroupBridgeDomainVnis struct {
	VniId types.Int64 `tfsdk:"vni_id"`
}
type L2VPNBridgeGroupBridgeDomainInterfaces struct {
	InterfaceName     types.String `tfsdk:"interface_name"`
	SplitHorizonGroup types.Bool   `tfsdk:"split_horizon_group"`
}

func (data L2VPNBridgeGroupBridgeDomain) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/bridge/groups/group[group-name=%s]/bridge-domains/bridge-domain[bridge-domain-name=%s]", data.BridgeGroupName.ValueString(), data.BridgeDomainName.ValueString())
}

func (data L2VPNBridgeGroupBridgeDomain) toBody(ctx context.Context) string {
	body := "{}"
	if !data.BridgeDomainName.IsNull() && !data.BridgeDomainName.IsUnknown() {
		body, _ = sjson.Set(body, "bridge-domain-name", data.BridgeDomainName.ValueString())
	}
	if !data.Mtu.IsNull() && !data.Mtu.IsUnknown() {
		body, _ = sjson.Set(body, "mtu", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}
	if !data.StormControlBroadcastPps.IsNull() && !data.StormControlBroadcastPps.IsUnknown() {
		body, _ = sjson.Set(body, "storm-control.broadcast.pps", strconv.FormatInt(data.StormControlBroadcastPps.ValueInt64(), 10))
	}
	if !data.StormControlBroadcastKbps.IsNull() && !data.StormControlBroadcastKbps.IsUnknown() {
		body, _ = sjson.Set(body, "storm-control.broadcast.kbps", strconv.FormatInt(data.StormControlBroadcastKbps.ValueInt64(), 10))
	}
	if !data.StormControlMulticastPps.IsNull() && !data.StormControlMulticastPps.IsUnknown() {
		body, _ = sjson.Set(body, "storm-control.multicast.pps", strconv.FormatInt(data.StormControlMulticastPps.ValueInt64(), 10))
	}
	if !data.StormControlMulticastKbps.IsNull() && !data.StormControlMulticastKbps.IsUnknown() {
		body, _ = sjson.Set(body, "storm-control.multicast.kbps", strconv.FormatInt(data.StormControlMulticastKbps.ValueInt64(), 10))
	}
	if !data.StormControlUnknownUnicastPps.IsNull() && !data.StormControlUnknownUnicastPps.IsUnknown() {
		body, _ = sjson.Set(body, "storm-control.unknown-unicast.pps", strconv.FormatInt(data.StormControlUnknownUnicastPps.ValueInt64(), 10))
	}
	if !data.StormControlUnknownUnicastKbps.IsNull() && !data.StormControlUnknownUnicastKbps.IsUnknown() {
		body, _ = sjson.Set(body, "storm-control.unknown-unicast.kbps", strconv.FormatInt(data.StormControlUnknownUnicastKbps.ValueInt64(), 10))
	}
	if len(data.Evis) > 0 {
		body, _ = sjson.Set(body, "evis.evi", []interface{}{})
		for index, item := range data.Evis {
			if !item.VpnId.IsNull() && !item.VpnId.IsUnknown() {
				body, _ = sjson.Set(body, "evis.evi"+"."+strconv.Itoa(index)+"."+"vpn-id", strconv.FormatInt(item.VpnId.ValueInt64(), 10))
			}
		}
	}
	if len(data.Vnis) > 0 {
		body, _ = sjson.Set(body, "vnis.vni", []interface{}{})
		for index, item := range data.Vnis {
			if !item.VniId.IsNull() && !item.VniId.IsUnknown() {
				body, _ = sjson.Set(body, "vnis.vni"+"."+strconv.Itoa(index)+"."+"vni-id", strconv.FormatInt(item.VniId.ValueInt64(), 10))
			}
		}
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces.interface", []interface{}{})
		for index, item := range data.Interfaces {
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
			if !item.SplitHorizonGroup.IsNull() && !item.SplitHorizonGroup.IsUnknown() {
				if item.SplitHorizonGroup.ValueBool() {
					body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"split-horizon.group", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *L2VPNBridgeGroupBridgeDomain) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.Evis {
		keys := [...]string{"vpn-id"}
		keyValues := [...]string{strconv.FormatInt(data.Evis[i].VpnId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "evis.evi").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("vpn-id"); value.Exists() && !data.Evis[i].VpnId.IsNull() {
			data.Evis[i].VpnId = types.Int64Value(value.Int())
		} else {
			data.Evis[i].VpnId = types.Int64Null()
		}
	}
	for i := range data.Vnis {
		keys := [...]string{"vni-id"}
		keyValues := [...]string{strconv.FormatInt(data.Vnis[i].VniId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "vnis.vni").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("vni-id"); value.Exists() && !data.Vnis[i].VniId.IsNull() {
			data.Vnis[i].VniId = types.Int64Value(value.Int())
		} else {
			data.Vnis[i].VniId = types.Int64Null()
		}
	}
	if value := gjson.GetBytes(res, "mtu"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "storm-control.broadcast.pps"); value.Exists() && !data.StormControlBroadcastPps.IsNull() {
		data.StormControlBroadcastPps = types.Int64Value(value.Int())
	} else {
		data.StormControlBroadcastPps = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "storm-control.broadcast.kbps"); value.Exists() && !data.StormControlBroadcastKbps.IsNull() {
		data.StormControlBroadcastKbps = types.Int64Value(value.Int())
	} else {
		data.StormControlBroadcastKbps = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "storm-control.multicast.pps"); value.Exists() && !data.StormControlMulticastPps.IsNull() {
		data.StormControlMulticastPps = types.Int64Value(value.Int())
	} else {
		data.StormControlMulticastPps = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "storm-control.multicast.kbps"); value.Exists() && !data.StormControlMulticastKbps.IsNull() {
		data.StormControlMulticastKbps = types.Int64Value(value.Int())
	} else {
		data.StormControlMulticastKbps = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "storm-control.unknown-unicast.pps"); value.Exists() && !data.StormControlUnknownUnicastPps.IsNull() {
		data.StormControlUnknownUnicastPps = types.Int64Value(value.Int())
	} else {
		data.StormControlUnknownUnicastPps = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "storm-control.unknown-unicast.kbps"); value.Exists() && !data.StormControlUnknownUnicastKbps.IsNull() {
		data.StormControlUnknownUnicastKbps = types.Int64Value(value.Int())
	} else {
		data.StormControlUnknownUnicastKbps = types.Int64Null()
	}
	for i := range data.Interfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "interfaces.interface").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("interface-name"); value.Exists() && !data.Interfaces[i].InterfaceName.IsNull() {
			data.Interfaces[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.Interfaces[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("split-horizon.group"); !data.Interfaces[i].SplitHorizonGroup.IsNull() {
			if value.Exists() {
				data.Interfaces[i].SplitHorizonGroup = types.BoolValue(true)
			} else {
				data.Interfaces[i].SplitHorizonGroup = types.BoolValue(false)
			}
		} else {
			data.Interfaces[i].SplitHorizonGroup = types.BoolNull()
		}
	}
}

func (data *L2VPNBridgeGroupBridgeDomain) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "evis.evi"); value.Exists() {
		data.Evis = make([]L2VPNBridgeGroupBridgeDomainEvis, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNBridgeGroupBridgeDomainEvis{}
			if cValue := v.Get("vpn-id"); cValue.Exists() {
				item.VpnId = types.Int64Value(cValue.Int())
			}
			data.Evis = append(data.Evis, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "vnis.vni"); value.Exists() {
		data.Vnis = make([]L2VPNBridgeGroupBridgeDomainVnis, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNBridgeGroupBridgeDomainVnis{}
			if cValue := v.Get("vni-id"); cValue.Exists() {
				item.VniId = types.Int64Value(cValue.Int())
			}
			data.Vnis = append(data.Vnis, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "mtu"); value.Exists() {
		data.Mtu = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "storm-control.broadcast.pps"); value.Exists() {
		data.StormControlBroadcastPps = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "storm-control.broadcast.kbps"); value.Exists() {
		data.StormControlBroadcastKbps = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "storm-control.multicast.pps"); value.Exists() {
		data.StormControlMulticastPps = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "storm-control.multicast.kbps"); value.Exists() {
		data.StormControlMulticastKbps = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "storm-control.unknown-unicast.pps"); value.Exists() {
		data.StormControlUnknownUnicastPps = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "storm-control.unknown-unicast.kbps"); value.Exists() {
		data.StormControlUnknownUnicastKbps = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "interfaces.interface"); value.Exists() {
		data.Interfaces = make([]L2VPNBridgeGroupBridgeDomainInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNBridgeGroupBridgeDomainInterfaces{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("split-horizon.group"); cValue.Exists() {
				item.SplitHorizonGroup = types.BoolValue(true)
			} else {
				item.SplitHorizonGroup = types.BoolValue(false)
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
}

func (data *L2VPNBridgeGroupBridgeDomain) fromPlan(ctx context.Context, plan L2VPNBridgeGroupBridgeDomain) {
	data.Device = plan.Device
	data.BridgeGroupName = types.StringValue(plan.BridgeGroupName.ValueString())
	data.BridgeDomainName = types.StringValue(plan.BridgeDomainName.ValueString())
}

func (data *L2VPNBridgeGroupBridgeDomain) getDeletedListItems(ctx context.Context, state L2VPNBridgeGroupBridgeDomain) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Evis {
		keys := [...]string{"vpn-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.Evis[i].VpnId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Evis[i].VpnId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Evis {
			found = true
			if state.Evis[i].VpnId.ValueInt64() != data.Evis[j].VpnId.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/evis/evi%v", state.getPath(), keyString))
		}
	}
	for i := range state.Vnis {
		keys := [...]string{"vni-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.Vnis[i].VniId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Vnis[i].VniId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Vnis {
			found = true
			if state.Vnis[i].VniId.ValueInt64() != data.Vnis[j].VniId.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/vnis/vni%v", state.getPath(), keyString))
		}
	}
	for i := range state.Interfaces {
		keys := [...]string{"interface-name"}
		stateKeyValues := [...]string{state.Interfaces[i].InterfaceName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Interfaces[i].InterfaceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Interfaces {
			found = true
			if state.Interfaces[i].InterfaceName.ValueString() != data.Interfaces[j].InterfaceName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/interfaces/interface%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *L2VPNBridgeGroupBridgeDomain) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}