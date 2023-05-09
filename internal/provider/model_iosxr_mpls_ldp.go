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

type MPLSLDP struct {
	Device          types.String             `tfsdk:"device"`
	Id              types.String             `tfsdk:"id"`
	RouterId        types.String             `tfsdk:"router_id"`
	AddressFamilies []MPLSLDPAddressFamilies `tfsdk:"address_families"`
	Interfaces      []MPLSLDPInterfaces      `tfsdk:"interfaces"`
}
type MPLSLDPAddressFamilies struct {
	AfName types.String `tfsdk:"af_name"`
}
type MPLSLDPInterfaces struct {
	InterfaceName types.String `tfsdk:"interface_name"`
}

func (data MPLSLDP) getPath() string {
	return "Cisco-IOS-XR-um-mpls-ldp-cfg:/mpls/ldp"
}

func (data MPLSLDP) toBody(ctx context.Context) string {
	body := "{}"
	if !data.RouterId.IsNull() && !data.RouterId.IsUnknown() {
		body, _ = sjson.Set(body, "router-id", data.RouterId.ValueString())
	}
	if len(data.AddressFamilies) > 0 {
		body, _ = sjson.Set(body, "address-families.address-family", []interface{}{})
		for index, item := range data.AddressFamilies {
			if !item.AfName.IsNull() && !item.AfName.IsUnknown() {
				body, _ = sjson.Set(body, "address-families.address-family"+"."+strconv.Itoa(index)+"."+"af-name", item.AfName.ValueString())
			}
		}
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces.interface", []interface{}{})
		for index, item := range data.Interfaces {
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
		}
	}
	return body
}

func (data *MPLSLDP) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "router-id"); value.Exists() && !data.RouterId.IsNull() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	for i := range data.AddressFamilies {
		keys := [...]string{"af-name"}
		keyValues := [...]string{data.AddressFamilies[i].AfName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "address-families.address-family").ForEach(
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
		if value := r.Get("af-name"); value.Exists() && !data.AddressFamilies[i].AfName.IsNull() {
			data.AddressFamilies[i].AfName = types.StringValue(value.String())
		} else {
			data.AddressFamilies[i].AfName = types.StringNull()
		}
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
	}
}

func (data *MPLSLDP) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "router-id"); value.Exists() {
		data.RouterId = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "address-families.address-family"); value.Exists() {
		data.AddressFamilies = make([]MPLSLDPAddressFamilies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MPLSLDPAddressFamilies{}
			if cValue := v.Get("af-name"); cValue.Exists() {
				item.AfName = types.StringValue(cValue.String())
			}
			data.AddressFamilies = append(data.AddressFamilies, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "interfaces.interface"); value.Exists() {
		data.Interfaces = make([]MPLSLDPInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MPLSLDPInterfaces{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
}

func (data *MPLSLDP) getDeletedListItems(ctx context.Context, state MPLSLDP) []string {
	deletedListItems := make([]string, 0)
	for i := range state.AddressFamilies {
		keys := [...]string{"af-name"}
		stateKeyValues := [...]string{state.AddressFamilies[i].AfName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.AddressFamilies[i].AfName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.AddressFamilies {
			found = true
			if state.AddressFamilies[i].AfName.ValueString() != data.AddressFamilies[j].AfName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/address-families/address-family%v", state.getPath(), keyString))
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

func (data *MPLSLDP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}
