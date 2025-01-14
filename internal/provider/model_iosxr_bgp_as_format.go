// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPASFormat struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	Asdot   types.Bool   `tfsdk:"asdot"`
	Asplain types.Bool   `tfsdk:"asplain"`
}

type BGPASFormatData struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	Asdot   types.Bool   `tfsdk:"asdot"`
	Asplain types.Bool   `tfsdk:"asplain"`
}

func (data BGPASFormat) getPath() string {
	return "Cisco-IOS-XR-um-router-bgp-cfg:/as-format"
}

func (data BGPASFormatData) getPath() string {
	return "Cisco-IOS-XR-um-router-bgp-cfg:/as-format"
}

func (data BGPASFormat) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Asdot.IsNull() && !data.Asdot.IsUnknown() {
		if data.Asdot.ValueBool() {
			body, _ = sjson.Set(body, "asdot", map[string]string{})
		}
	}
	if !data.Asplain.IsNull() && !data.Asplain.IsUnknown() {
		if data.Asplain.ValueBool() {
			body, _ = sjson.Set(body, "asplain", map[string]string{})
		}
	}
	return body
}

func (data *BGPASFormat) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "asdot"); !data.Asdot.IsNull() {
		if value.Exists() {
			data.Asdot = types.BoolValue(true)
		} else {
			data.Asdot = types.BoolValue(false)
		}
	} else {
		data.Asdot = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "asplain"); !data.Asplain.IsNull() {
		if value.Exists() {
			data.Asplain = types.BoolValue(true)
		} else {
			data.Asplain = types.BoolValue(false)
		}
	} else {
		data.Asplain = types.BoolNull()
	}
}

func (data *BGPASFormatData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "asdot"); value.Exists() {
		data.Asdot = types.BoolValue(true)
	} else {
		data.Asdot = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "asplain"); value.Exists() {
		data.Asplain = types.BoolValue(true)
	} else {
		data.Asplain = types.BoolValue(false)
	}
}

func (data *BGPASFormat) getDeletedListItems(ctx context.Context, state BGPASFormat) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *BGPASFormat) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Asdot.IsNull() && !data.Asdot.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/asdot", data.getPath()))
	}
	if !data.Asplain.IsNull() && !data.Asplain.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/asplain", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *BGPASFormat) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Asdot.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/asdot", data.getPath()))
	}
	if !data.Asplain.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/asplain", data.getPath()))
	}
	return deletePaths
}
