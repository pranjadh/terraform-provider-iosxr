// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type BGPASFormat struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	Asdot   types.Bool   `tfsdk:"asdot"`
	Asplain types.Bool   `tfsdk:"asplain"`
}

func (data BGPASFormat) getPath() string {
	return "Cisco-IOS-XR-um-router-bgp-cfg:/as-format"
}

func (data BGPASFormat) toBody() string {
	body := "{}"
	if !data.Asdot.Null && !data.Asdot.Unknown {
		if data.Asdot.Value {
			body, _ = sjson.Set(body, "asdot", map[string]string{})
		}
	}
	if !data.Asplain.Null && !data.Asplain.Unknown {
		if data.Asplain.Value {
			body, _ = sjson.Set(body, "asplain", map[string]string{})
		}
	}
	return body
}

func (data *BGPASFormat) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "asdot"); value.Exists() {
		data.Asdot.Value = true
	} else {
		data.Asdot.Value = false
	}
	if value := gjson.GetBytes(res, "asplain"); value.Exists() {
		data.Asplain.Value = true
	} else {
		data.Asplain.Value = false
	}
}

func (data *BGPASFormat) fromPlan(plan BGPASFormat) {
	data.Device = plan.Device
}

func (data *BGPASFormat) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Asdot.Unknown {
		data.Asdot.Unknown = false
		data.Asdot.Null = true
	}
	if data.Asplain.Unknown {
		data.Asplain.Unknown = false
		data.Asplain.Null = true
	}
}

func (data *BGPASFormat) getDeletedListItems(state BGPASFormat) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *BGPASFormat) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
