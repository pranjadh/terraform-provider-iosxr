// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VRFRouteTargetTwoByteASFormat struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	VrfName          types.String `tfsdk:"vrf_name"`
	AddressFamily    types.String `tfsdk:"address_family"`
	SubAddressFamily types.String `tfsdk:"sub_address_family"`
	Direction        types.String `tfsdk:"direction"`
	AsNumber         types.Int64  `tfsdk:"as_number"`
	Index            types.Int64  `tfsdk:"index"`
	Stitching        types.Bool   `tfsdk:"stitching"`
}

func (data VRFRouteTargetTwoByteASFormat) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=%s]/address-family/%s/%s/Cisco-IOS-XR-um-router-bgp-cfg:%s/Cisco-IOS-XR-um-router-bgp-cfg:route-target/Cisco-IOS-XR-um-router-bgp-cfg:two-byte-as-rts/Cisco-IOS-XR-um-router-bgp-cfg:two-byte-as-rt[as-number=%v][index=%v][stitching=%v]", data.VrfName.Value, data.AddressFamily.Value, data.SubAddressFamily.Value, data.Direction.Value, data.AsNumber.Value, data.Index.Value, data.Stitching.Value)
}

func (data VRFRouteTargetTwoByteASFormat) toBody() string {
	body := "{}"
	return body
}

func (data *VRFRouteTargetTwoByteASFormat) fromBody(res []byte) {
}

func (data *VRFRouteTargetTwoByteASFormat) fromPlan(plan VRFRouteTargetTwoByteASFormat) {
	data.Device = plan.Device
	data.VrfName.Value = plan.VrfName.Value
	data.AddressFamily.Value = plan.AddressFamily.Value
	data.SubAddressFamily.Value = plan.SubAddressFamily.Value
	data.Direction.Value = plan.Direction.Value
	data.AsNumber.Value = plan.AsNumber.Value
	data.Index.Value = plan.Index.Value
	data.Stitching.Value = plan.Stitching.Value
}

func (data *VRFRouteTargetTwoByteASFormat) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.VrfName.Unknown {
		data.VrfName.Unknown = false
		data.VrfName.Null = true
	}
	if data.AddressFamily.Unknown {
		data.AddressFamily.Unknown = false
		data.AddressFamily.Null = true
	}
	if data.SubAddressFamily.Unknown {
		data.SubAddressFamily.Unknown = false
		data.SubAddressFamily.Null = true
	}
	if data.Direction.Unknown {
		data.Direction.Unknown = false
		data.Direction.Null = true
	}
	if data.AsNumber.Unknown {
		data.AsNumber.Unknown = false
		data.AsNumber.Null = true
	}
	if data.Index.Unknown {
		data.Index.Unknown = false
		data.Index.Null = true
	}
	if data.Stitching.Unknown {
		data.Stitching.Unknown = false
		data.Stitching.Null = true
	}
}

func (data *VRFRouteTargetTwoByteASFormat) getDeletedListItems(state VRFRouteTargetTwoByteASFormat) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *VRFRouteTargetTwoByteASFormat) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
