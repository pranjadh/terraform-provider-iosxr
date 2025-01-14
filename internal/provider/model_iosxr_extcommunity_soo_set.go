// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ExtcommunitySOOSet struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

type ExtcommunitySOOSetData struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

func (data ExtcommunitySOOSet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/extended-community-soo-sets/extended-community-soo-set[set-name=%s]", data.SetName.ValueString())
}

func (data ExtcommunitySOOSetData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/extended-community-soo-sets/extended-community-soo-set[set-name=%s]", data.SetName.ValueString())
}

func (data ExtcommunitySOOSet) toBody(ctx context.Context) string {
	body := "{}"
	if !data.SetName.IsNull() && !data.SetName.IsUnknown() {
		body, _ = sjson.Set(body, "set-name", data.SetName.ValueString())
	}
	if !data.Rpl.IsNull() && !data.Rpl.IsUnknown() {
		body, _ = sjson.Set(body, "rpl-extended-community-soo-set", data.Rpl.ValueString())
	}
	return body
}

func (data *ExtcommunitySOOSet) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rpl-extended-community-soo-set"); value.Exists() && !data.Rpl.IsNull() {
		data.Rpl = types.StringValue(value.String())
	} else {
		data.Rpl = types.StringNull()
	}
}

func (data *ExtcommunitySOOSetData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rpl-extended-community-soo-set"); value.Exists() {
		data.Rpl = types.StringValue(value.String())
	}
}

func (data *ExtcommunitySOOSet) getDeletedListItems(ctx context.Context, state ExtcommunitySOOSet) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *ExtcommunitySOOSet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *ExtcommunitySOOSet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Rpl.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/rpl-extended-community-soo-set", data.getPath()))
	}
	return deletePaths
}
