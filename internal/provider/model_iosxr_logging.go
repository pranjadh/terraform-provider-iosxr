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

type Logging struct {
	Device                    types.String `tfsdk:"device"`
	Id                        types.String `tfsdk:"id"`
	Ipv4Dscp                  types.String `tfsdk:"ipv4_dscp"`
	Trap                      types.String `tfsdk:"trap"`
	EventsDisplayLocation     types.Bool   `tfsdk:"events_display_location"`
	EventsLevel               types.String `tfsdk:"events_level"`
	Console                   types.String `tfsdk:"console"`
	Monitor                   types.String `tfsdk:"monitor"`
	BufferedLoggingBufferSize types.Int64  `tfsdk:"buffered_logging_buffer_size"`
	BufferedLevel             types.String `tfsdk:"buffered_level"`
	FacilityLevel             types.String `tfsdk:"facility_level"`
	Hostnameprefix            types.String `tfsdk:"hostnameprefix"`
	SuppressDuplicates        types.Bool   `tfsdk:"suppress_duplicates"`
}

func (data Logging) getPath() string {
	return "Cisco-IOS-XR-um-logging-cfg:/logging"
}

func (data Logging) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Ipv4Dscp.IsNull() && !data.Ipv4Dscp.IsUnknown() {
		body, _ = sjson.Set(body, "ipv4.dscp", data.Ipv4Dscp.ValueString())
	}
	if !data.Trap.IsNull() && !data.Trap.IsUnknown() {
		body, _ = sjson.Set(body, "trap", data.Trap.ValueString())
	}
	if !data.EventsDisplayLocation.IsNull() && !data.EventsDisplayLocation.IsUnknown() {
		if data.EventsDisplayLocation.ValueBool() {
			body, _ = sjson.Set(body, "Cisco-IOS-XR-um-logging-events-cfg:events.display-location", map[string]string{})
		}
	}
	if !data.EventsLevel.IsNull() && !data.EventsLevel.IsUnknown() {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-logging-events-cfg:events.level", data.EventsLevel.ValueString())
	}
	if !data.Console.IsNull() && !data.Console.IsUnknown() {
		body, _ = sjson.Set(body, "console", data.Console.ValueString())
	}
	if !data.Monitor.IsNull() && !data.Monitor.IsUnknown() {
		body, _ = sjson.Set(body, "monitor", data.Monitor.ValueString())
	}
	if !data.BufferedLoggingBufferSize.IsNull() && !data.BufferedLoggingBufferSize.IsUnknown() {
		body, _ = sjson.Set(body, "buffered.logging-buffer-size", strconv.FormatInt(data.BufferedLoggingBufferSize.ValueInt64(), 10))
	}
	if !data.BufferedLevel.IsNull() && !data.BufferedLevel.IsUnknown() {
		body, _ = sjson.Set(body, "buffered.level", data.BufferedLevel.ValueString())
	}
	if !data.FacilityLevel.IsNull() && !data.FacilityLevel.IsUnknown() {
		body, _ = sjson.Set(body, "facility.level", data.FacilityLevel.ValueString())
	}
	if !data.Hostnameprefix.IsNull() && !data.Hostnameprefix.IsUnknown() {
		body, _ = sjson.Set(body, "hostnameprefix", data.Hostnameprefix.ValueString())
	}
	if !data.SuppressDuplicates.IsNull() && !data.SuppressDuplicates.IsUnknown() {
		if data.SuppressDuplicates.ValueBool() {
			body, _ = sjson.Set(body, "suppress.duplicates", map[string]string{})
		}
	}
	return body
}

func (data *Logging) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "ipv4.dscp"); value.Exists() && !data.Ipv4Dscp.IsNull() {
		data.Ipv4Dscp = types.StringValue(value.String())
	} else {
		data.Ipv4Dscp = types.StringNull()
	}
	if value := gjson.GetBytes(res, "trap"); value.Exists() && !data.Trap.IsNull() {
		data.Trap = types.StringValue(value.String())
	} else {
		data.Trap = types.StringNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-logging-events-cfg:events.display-location"); !data.EventsDisplayLocation.IsNull() {
		if value.Exists() {
			data.EventsDisplayLocation = types.BoolValue(true)
		} else {
			data.EventsDisplayLocation = types.BoolValue(false)
		}
	} else {
		data.EventsDisplayLocation = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-logging-events-cfg:events.level"); value.Exists() && !data.EventsLevel.IsNull() {
		data.EventsLevel = types.StringValue(value.String())
	} else {
		data.EventsLevel = types.StringNull()
	}
	if value := gjson.GetBytes(res, "console"); value.Exists() && !data.Console.IsNull() {
		data.Console = types.StringValue(value.String())
	} else {
		data.Console = types.StringNull()
	}
	if value := gjson.GetBytes(res, "monitor"); value.Exists() && !data.Monitor.IsNull() {
		data.Monitor = types.StringValue(value.String())
	} else {
		data.Monitor = types.StringNull()
	}
	if value := gjson.GetBytes(res, "buffered.logging-buffer-size"); value.Exists() && !data.BufferedLoggingBufferSize.IsNull() {
		data.BufferedLoggingBufferSize = types.Int64Value(value.Int())
	} else {
		data.BufferedLoggingBufferSize = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "buffered.level"); value.Exists() && !data.BufferedLevel.IsNull() {
		data.BufferedLevel = types.StringValue(value.String())
	} else {
		data.BufferedLevel = types.StringNull()
	}
	if value := gjson.GetBytes(res, "facility.level"); value.Exists() && !data.FacilityLevel.IsNull() {
		data.FacilityLevel = types.StringValue(value.String())
	} else {
		data.FacilityLevel = types.StringNull()
	}
	if value := gjson.GetBytes(res, "hostnameprefix"); value.Exists() && !data.Hostnameprefix.IsNull() {
		data.Hostnameprefix = types.StringValue(value.String())
	} else {
		data.Hostnameprefix = types.StringNull()
	}
	if value := gjson.GetBytes(res, "suppress.duplicates"); !data.SuppressDuplicates.IsNull() {
		if value.Exists() {
			data.SuppressDuplicates = types.BoolValue(true)
		} else {
			data.SuppressDuplicates = types.BoolValue(false)
		}
	} else {
		data.SuppressDuplicates = types.BoolNull()
	}
}

func (data *Logging) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "ipv4.dscp"); value.Exists() {
		data.Ipv4Dscp = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "trap"); value.Exists() {
		data.Trap = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-logging-events-cfg:events.display-location"); value.Exists() {
		data.EventsDisplayLocation = types.BoolValue(true)
	} else {
		data.EventsDisplayLocation = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-logging-events-cfg:events.level"); value.Exists() {
		data.EventsLevel = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "console"); value.Exists() {
		data.Console = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "monitor"); value.Exists() {
		data.Monitor = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "buffered.logging-buffer-size"); value.Exists() {
		data.BufferedLoggingBufferSize = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "buffered.level"); value.Exists() {
		data.BufferedLevel = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "facility.level"); value.Exists() {
		data.FacilityLevel = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "hostnameprefix"); value.Exists() {
		data.Hostnameprefix = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "suppress.duplicates"); value.Exists() {
		data.SuppressDuplicates = types.BoolValue(true)
	} else {
		data.SuppressDuplicates = types.BoolValue(false)
	}
}

func (data *Logging) getDeletedListItems(ctx context.Context, state Logging) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *Logging) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.EventsDisplayLocation.IsNull() && !data.EventsDisplayLocation.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XR-um-logging-events-cfg:events/display-location", data.getPath()))
	}
	if !data.SuppressDuplicates.IsNull() && !data.SuppressDuplicates.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/suppress/duplicates", data.getPath()))
	}
	return emptyLeafsDelete
}
