// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type EVPNInterface struct {
	Device                                           types.String `tfsdk:"device"`
	Id                                               types.String `tfsdk:"id"`
	InterfaceName                                    types.String `tfsdk:"interface_name"`
	CoreIsolationGroup                               types.Int64  `tfsdk:"core_isolation_group"`
	EthernetSegmentIdentifierTypeZeroBytes1          types.String `tfsdk:"ethernet_segment_identifier_type_zero_bytes_1"`
	EthernetSegmentIdentifierTypeZeroBytes23         types.String `tfsdk:"ethernet_segment_identifier_type_zero_bytes_23"`
	EthernetSegmentIdentifierTypeZeroBytes45         types.String `tfsdk:"ethernet_segment_identifier_type_zero_bytes_45"`
	EthernetSegmentIdentifierTypeZeroBytes67         types.String `tfsdk:"ethernet_segment_identifier_type_zero_bytes_67"`
	EthernetSegmentIdentifierTypeZeroBytes89         types.String `tfsdk:"ethernet_segment_identifier_type_zero_bytes_89"`
	EthernetSegmentIdentifierTypeZeroEsi             types.String `tfsdk:"ethernet_segment_identifier_type_zero_esi"`
	EthernetSegmentLoadBalancingModeAllActive        types.Bool   `tfsdk:"ethernet_segment_load_balancing_mode_all_active"`
	EthernetSegmentLoadBalancingModePortActive       types.Bool   `tfsdk:"ethernet_segment_load_balancing_mode_port_active"`
	EthernetSegmentLoadBalancingModeSingleActive     types.Bool   `tfsdk:"ethernet_segment_load_balancing_mode_single_active"`
	EthernetSegmentLoadBalancingModeSingleFlowActive types.Bool   `tfsdk:"ethernet_segment_load_balancing_mode_single_flow_active"`
}

func (data EVPNInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/evpn/interface/interface[interface-name=%s]", data.InterfaceName.ValueString())
}

func (data EVPNInterface) toBody(ctx context.Context) string {
	body := "{}"
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "interface-name", data.InterfaceName.ValueString())
	}
	if !data.CoreIsolationGroup.IsNull() && !data.CoreIsolationGroup.IsUnknown() {
		body, _ = sjson.Set(body, "core-isolation-group", strconv.FormatInt(data.CoreIsolationGroup.ValueInt64(), 10))
	}
	if !data.EthernetSegmentIdentifierTypeZeroBytes1.IsNull() && !data.EthernetSegmentIdentifierTypeZeroBytes1.IsUnknown() {
		body, _ = sjson.Set(body, "ethernet-segment.identifier.type.zero.bytes-1", data.EthernetSegmentIdentifierTypeZeroBytes1.ValueString())
	}
	if !data.EthernetSegmentIdentifierTypeZeroBytes23.IsNull() && !data.EthernetSegmentIdentifierTypeZeroBytes23.IsUnknown() {
		body, _ = sjson.Set(body, "ethernet-segment.identifier.type.zero.bytes-23", data.EthernetSegmentIdentifierTypeZeroBytes23.ValueString())
	}
	if !data.EthernetSegmentIdentifierTypeZeroBytes45.IsNull() && !data.EthernetSegmentIdentifierTypeZeroBytes45.IsUnknown() {
		body, _ = sjson.Set(body, "ethernet-segment.identifier.type.zero.bytes-45", data.EthernetSegmentIdentifierTypeZeroBytes45.ValueString())
	}
	if !data.EthernetSegmentIdentifierTypeZeroBytes67.IsNull() && !data.EthernetSegmentIdentifierTypeZeroBytes67.IsUnknown() {
		body, _ = sjson.Set(body, "ethernet-segment.identifier.type.zero.bytes-67", data.EthernetSegmentIdentifierTypeZeroBytes67.ValueString())
	}
	if !data.EthernetSegmentIdentifierTypeZeroBytes89.IsNull() && !data.EthernetSegmentIdentifierTypeZeroBytes89.IsUnknown() {
		body, _ = sjson.Set(body, "ethernet-segment.identifier.type.zero.bytes-89", data.EthernetSegmentIdentifierTypeZeroBytes89.ValueString())
	}
	if !data.EthernetSegmentIdentifierTypeZeroEsi.IsNull() && !data.EthernetSegmentIdentifierTypeZeroEsi.IsUnknown() {
		body, _ = sjson.Set(body, "ethernet-segment.identifier.type.zero.esi", data.EthernetSegmentIdentifierTypeZeroEsi.ValueString())
	}
	if !data.EthernetSegmentLoadBalancingModeAllActive.IsNull() && !data.EthernetSegmentLoadBalancingModeAllActive.IsUnknown() {
		if data.EthernetSegmentLoadBalancingModeAllActive.ValueBool() {
			body, _ = sjson.Set(body, "ethernet-segment.load-balancing-mode.all-active", map[string]string{})
		}
	}
	if !data.EthernetSegmentLoadBalancingModePortActive.IsNull() && !data.EthernetSegmentLoadBalancingModePortActive.IsUnknown() {
		if data.EthernetSegmentLoadBalancingModePortActive.ValueBool() {
			body, _ = sjson.Set(body, "ethernet-segment.load-balancing-mode.port-active", map[string]string{})
		}
	}
	if !data.EthernetSegmentLoadBalancingModeSingleActive.IsNull() && !data.EthernetSegmentLoadBalancingModeSingleActive.IsUnknown() {
		if data.EthernetSegmentLoadBalancingModeSingleActive.ValueBool() {
			body, _ = sjson.Set(body, "ethernet-segment.load-balancing-mode.single-active", map[string]string{})
		}
	}
	if !data.EthernetSegmentLoadBalancingModeSingleFlowActive.IsNull() && !data.EthernetSegmentLoadBalancingModeSingleFlowActive.IsUnknown() {
		if data.EthernetSegmentLoadBalancingModeSingleFlowActive.ValueBool() {
			body, _ = sjson.Set(body, "ethernet-segment.load-balancing-mode.single-flow-active", map[string]string{})
		}
	}
	return body
}

func (data *EVPNInterface) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "core-isolation-group"); value.Exists() && !data.CoreIsolationGroup.IsNull() {
		data.CoreIsolationGroup = types.Int64Value(value.Int())
	} else {
		data.CoreIsolationGroup = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-1"); value.Exists() && !data.EthernetSegmentIdentifierTypeZeroBytes1.IsNull() {
		data.EthernetSegmentIdentifierTypeZeroBytes1 = types.StringValue(value.String())
	} else {
		data.EthernetSegmentIdentifierTypeZeroBytes1 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-23"); value.Exists() && !data.EthernetSegmentIdentifierTypeZeroBytes23.IsNull() {
		data.EthernetSegmentIdentifierTypeZeroBytes23 = types.StringValue(value.String())
	} else {
		data.EthernetSegmentIdentifierTypeZeroBytes23 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-45"); value.Exists() && !data.EthernetSegmentIdentifierTypeZeroBytes45.IsNull() {
		data.EthernetSegmentIdentifierTypeZeroBytes45 = types.StringValue(value.String())
	} else {
		data.EthernetSegmentIdentifierTypeZeroBytes45 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-67"); value.Exists() && !data.EthernetSegmentIdentifierTypeZeroBytes67.IsNull() {
		data.EthernetSegmentIdentifierTypeZeroBytes67 = types.StringValue(value.String())
	} else {
		data.EthernetSegmentIdentifierTypeZeroBytes67 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-89"); value.Exists() && !data.EthernetSegmentIdentifierTypeZeroBytes89.IsNull() {
		data.EthernetSegmentIdentifierTypeZeroBytes89 = types.StringValue(value.String())
	} else {
		data.EthernetSegmentIdentifierTypeZeroBytes89 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.esi"); value.Exists() && !data.EthernetSegmentIdentifierTypeZeroEsi.IsNull() {
		data.EthernetSegmentIdentifierTypeZeroEsi = types.StringValue(value.String())
	} else {
		data.EthernetSegmentIdentifierTypeZeroEsi = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.all-active"); !data.EthernetSegmentLoadBalancingModeAllActive.IsNull() {
		if value.Exists() {
			data.EthernetSegmentLoadBalancingModeAllActive = types.BoolValue(true)
		} else {
			data.EthernetSegmentLoadBalancingModeAllActive = types.BoolValue(false)
		}
	} else {
		data.EthernetSegmentLoadBalancingModeAllActive = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.port-active"); !data.EthernetSegmentLoadBalancingModePortActive.IsNull() {
		if value.Exists() {
			data.EthernetSegmentLoadBalancingModePortActive = types.BoolValue(true)
		} else {
			data.EthernetSegmentLoadBalancingModePortActive = types.BoolValue(false)
		}
	} else {
		data.EthernetSegmentLoadBalancingModePortActive = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.single-active"); !data.EthernetSegmentLoadBalancingModeSingleActive.IsNull() {
		if value.Exists() {
			data.EthernetSegmentLoadBalancingModeSingleActive = types.BoolValue(true)
		} else {
			data.EthernetSegmentLoadBalancingModeSingleActive = types.BoolValue(false)
		}
	} else {
		data.EthernetSegmentLoadBalancingModeSingleActive = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.single-flow-active"); !data.EthernetSegmentLoadBalancingModeSingleFlowActive.IsNull() {
		if value.Exists() {
			data.EthernetSegmentLoadBalancingModeSingleFlowActive = types.BoolValue(true)
		} else {
			data.EthernetSegmentLoadBalancingModeSingleFlowActive = types.BoolValue(false)
		}
	} else {
		data.EthernetSegmentLoadBalancingModeSingleFlowActive = types.BoolNull()
	}
}

func (data *EVPNInterface) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "core-isolation-group"); value.Exists() {
		data.CoreIsolationGroup = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-1"); value.Exists() {
		data.EthernetSegmentIdentifierTypeZeroBytes1 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-23"); value.Exists() {
		data.EthernetSegmentIdentifierTypeZeroBytes23 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-45"); value.Exists() {
		data.EthernetSegmentIdentifierTypeZeroBytes45 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-67"); value.Exists() {
		data.EthernetSegmentIdentifierTypeZeroBytes67 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.bytes-89"); value.Exists() {
		data.EthernetSegmentIdentifierTypeZeroBytes89 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ethernet-segment.identifier.type.zero.esi"); value.Exists() {
		data.EthernetSegmentIdentifierTypeZeroEsi = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.all-active"); value.Exists() {
		data.EthernetSegmentLoadBalancingModeAllActive = types.BoolValue(true)
	} else {
		data.EthernetSegmentLoadBalancingModeAllActive = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.port-active"); value.Exists() {
		data.EthernetSegmentLoadBalancingModePortActive = types.BoolValue(true)
	} else {
		data.EthernetSegmentLoadBalancingModePortActive = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.single-active"); value.Exists() {
		data.EthernetSegmentLoadBalancingModeSingleActive = types.BoolValue(true)
	} else {
		data.EthernetSegmentLoadBalancingModeSingleActive = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "ethernet-segment.load-balancing-mode.single-flow-active"); value.Exists() {
		data.EthernetSegmentLoadBalancingModeSingleFlowActive = types.BoolValue(true)
	} else {
		data.EthernetSegmentLoadBalancingModeSingleFlowActive = types.BoolValue(false)
	}
}

func (data *EVPNInterface) getDeletedListItems(ctx context.Context, state EVPNInterface) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *EVPNInterface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.EthernetSegmentLoadBalancingModeAllActive.IsNull() && !data.EthernetSegmentLoadBalancingModeAllActive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ethernet-segment/load-balancing-mode/all-active", data.getPath()))
	}
	if !data.EthernetSegmentLoadBalancingModePortActive.IsNull() && !data.EthernetSegmentLoadBalancingModePortActive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ethernet-segment/load-balancing-mode/port-active", data.getPath()))
	}
	if !data.EthernetSegmentLoadBalancingModeSingleActive.IsNull() && !data.EthernetSegmentLoadBalancingModeSingleActive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ethernet-segment/load-balancing-mode/single-active", data.getPath()))
	}
	if !data.EthernetSegmentLoadBalancingModeSingleFlowActive.IsNull() && !data.EthernetSegmentLoadBalancingModeSingleFlowActive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ethernet-segment/load-balancing-mode/single-flow-active", data.getPath()))
	}
	return emptyLeafsDelete
}
