// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type EVPN struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	DeleteMode      types.String `tfsdk:"delete_mode"`
	SourceInterface types.String `tfsdk:"source_interface"`
}

type EVPNData struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	SourceInterface types.String `tfsdk:"source_interface"`
}

func (data EVPN) getPath() string {
	return "Cisco-IOS-XR-um-l2vpn-cfg:/evpn"
}

func (data EVPNData) getPath() string {
	return "Cisco-IOS-XR-um-l2vpn-cfg:/evpn"
}

func (data EVPN) toBody(ctx context.Context) string {
	body := "{}"
	if !data.SourceInterface.IsNull() && !data.SourceInterface.IsUnknown() {
		body, _ = sjson.Set(body, "source.interface", data.SourceInterface.ValueString())
	}
	return body
}

func (data *EVPN) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "source.interface"); value.Exists() && !data.SourceInterface.IsNull() {
		data.SourceInterface = types.StringValue(value.String())
	} else {
		data.SourceInterface = types.StringNull()
	}
}

func (data *EVPNData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "source.interface"); value.Exists() {
		data.SourceInterface = types.StringValue(value.String())
	}
}

func (data *EVPN) getDeletedListItems(ctx context.Context, state EVPN) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *EVPN) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *EVPN) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.SourceInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/source/interface", data.getPath()))
	}
	return deletePaths
}
