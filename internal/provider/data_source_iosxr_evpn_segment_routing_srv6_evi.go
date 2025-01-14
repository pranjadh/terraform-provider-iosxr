// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &EVPNSegmentRoutingSRv6EVIDataSource{}
	_ datasource.DataSourceWithConfigure = &EVPNSegmentRoutingSRv6EVIDataSource{}
)

func NewEVPNSegmentRoutingSRv6EVIDataSource() datasource.DataSource {
	return &EVPNSegmentRoutingSRv6EVIDataSource{}
}

type EVPNSegmentRoutingSRv6EVIDataSource struct {
	client *client.Client
}

func (d *EVPNSegmentRoutingSRv6EVIDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_evpn_segment_routing_srv6_evi"
}

func (d *EVPNSegmentRoutingSRv6EVIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the EVPN Segment Routing SRv6 EVI configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"vpn_id": schema.Int64Attribute{
				MarkdownDescription: "Configure EVPN Instance VPN ID",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this EVPN Instance",
				Computed:            true,
			},
			"bgp_route_target_import_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Two Byte AS Number",
							Computed:            true,
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: "AS:nn (hex or decimal format)",
							Computed:            true,
						},
					},
				},
			},
			"bgp_route_target_import_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Four Byte AS number",
							Computed:            true,
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: "AS:nn (hex or decimal format)",
							Computed:            true,
						},
					},
				},
			},
			"bgp_route_target_import_ipv4_address_format": schema.ListNestedAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: "IP-address:nn (hex or decimal format)",
							Computed:            true,
						},
					},
				},
			},
			"bgp_route_target_export_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Two Byte AS Number",
							Computed:            true,
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: "AS:nn (hex or decimal format)",
							Computed:            true,
						},
					},
				},
			},
			"bgp_route_target_export_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Four Byte AS number",
							Computed:            true,
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: "AS:nn (hex or decimal format)",
							Computed:            true,
						},
					},
				},
			},
			"bgp_route_target_export_ipv4_address_format": schema.ListNestedAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: "IP-address:nn (hex or decimal format)",
							Computed:            true,
						},
					},
				},
			},
			"advertise_mac": schema.BoolAttribute{
				MarkdownDescription: "Configure EVPN Instance MAC advertisement",
				Computed:            true,
			},
			"locator": schema.StringAttribute{
				MarkdownDescription: "EVI locator to use for EVPN SID allocation",
				Computed:            true,
			},
		},
	}
}

func (d *EVPNSegmentRoutingSRv6EVIDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *EVPNSegmentRoutingSRv6EVIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config EVPNSegmentRoutingSRv6EVIData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
