// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/sjson"
)

type L2VPNBridgeGroup struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	DeleteMode types.String `tfsdk:"delete_mode"`
	GroupName  types.String `tfsdk:"group_name"`
}

type L2VPNBridgeGroupData struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	GroupName types.String `tfsdk:"group_name"`
}

func (data L2VPNBridgeGroup) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/bridge/groups/group[group-name=%s]", data.GroupName.ValueString())
}

func (data L2VPNBridgeGroupData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/bridge/groups/group[group-name=%s]", data.GroupName.ValueString())
}

func (data L2VPNBridgeGroup) toBody(ctx context.Context) string {
	body := "{}"
	if !data.GroupName.IsNull() && !data.GroupName.IsUnknown() {
		body, _ = sjson.Set(body, "group-name", data.GroupName.ValueString())
	}
	return body
}

func (data *L2VPNBridgeGroup) updateFromBody(ctx context.Context, res []byte) {
}

func (data *L2VPNBridgeGroupData) fromBody(ctx context.Context, res []byte) {
}

func (data *L2VPNBridgeGroup) getDeletedListItems(ctx context.Context, state L2VPNBridgeGroup) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *L2VPNBridgeGroup) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *L2VPNBridgeGroup) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	return deletePaths
}
