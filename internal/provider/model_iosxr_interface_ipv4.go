// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type InterfaceIPv4 struct {
	Device        types.String `tfsdk:"device"`
	Id            types.String `tfsdk:"id"`
	InterfaceName types.String `tfsdk:"interface_name"`
	Address       types.String `tfsdk:"address"`
	Netmask       types.String `tfsdk:"netmask"`
	Unnumbered    types.String `tfsdk:"unnumbered"`
}

func (data InterfaceIPv4) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name=%s]/ipv4", data.InterfaceName.Value)
}

func (data InterfaceIPv4) toBody() string {
	body := "{}"
	if !data.Address.Null && !data.Address.Unknown {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.address", data.Address.Value)
	}
	if !data.Netmask.Null && !data.Netmask.Unknown {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.netmask", data.Netmask.Value)
	}
	if !data.Unnumbered.Null && !data.Unnumbered.Unknown {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.unnumbered", data.Unnumbered.Value)
	}
	return body
}

func (data *InterfaceIPv4) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.address"); value.Exists() {
		data.Address.Value = value.String()
	} else {
		data.Address.Null = true
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.netmask"); value.Exists() {
		data.Netmask.Value = value.String()
	} else {
		data.Netmask.Null = true
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.unnumbered"); value.Exists() {
		data.Unnumbered.Value = value.String()
	} else {
		data.Unnumbered.Null = true
	}
}

func (data *InterfaceIPv4) fromPlan(plan InterfaceIPv4) {
	data.Device = plan.Device
	data.InterfaceName.Value = plan.InterfaceName.Value
}

func (data *InterfaceIPv4) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.InterfaceName.Unknown {
		data.InterfaceName.Unknown = false
		data.InterfaceName.Null = true
	}
	if data.Address.Unknown {
		data.Address.Unknown = false
		data.Address.Null = true
	}
	if data.Netmask.Unknown {
		data.Netmask.Unknown = false
		data.Netmask.Null = true
	}
	if data.Unnumbered.Unknown {
		data.Unnumbered.Unknown = false
		data.Unnumbered.Null = true
	}
}

func (data *InterfaceIPv4) getDeletedListItems(state InterfaceIPv4) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *InterfaceIPv4) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
