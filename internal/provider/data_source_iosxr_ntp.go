// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

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
	_ datasource.DataSource              = &NTPDataSource{}
	_ datasource.DataSourceWithConfigure = &NTPDataSource{}
)

func NewNTPDataSource() datasource.DataSource {
	return &NTPDataSource{}
}

type NTPDataSource struct {
	client *client.Client
}

func (d *NTPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntp"
}

func (d *NTPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the NTP configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"ipv4_dscp": schema.StringAttribute{
				MarkdownDescription: "Set IP DSCP (DiffServ CodePoint)",
				Computed:            true,
			},
			"ipv4_precedence": schema.StringAttribute{
				MarkdownDescription: "Set precedence",
				Computed:            true,
			},
			"ipv6_dscp": schema.StringAttribute{
				MarkdownDescription: "Set IP DSCP (DiffServ CodePoint)",
				Computed:            true,
			},
			"ipv6_precedence": schema.StringAttribute{
				MarkdownDescription: "Set precedence",
				Computed:            true,
			},
			"access_group_ipv6_peer": schema.StringAttribute{
				MarkdownDescription: "Provide full access",
				Computed:            true,
			},
			"access_group_ipv6_query_only": schema.StringAttribute{
				MarkdownDescription: "Allow only control queries",
				Computed:            true,
			},
			"access_group_ipv6_serve": schema.StringAttribute{
				MarkdownDescription: "Provide server and query access",
				Computed:            true,
			},
			"access_group_ipv6_serve_only": schema.StringAttribute{
				MarkdownDescription: "Provide only server access",
				Computed:            true,
			},
			"access_group_ipv4_peer": schema.StringAttribute{
				MarkdownDescription: "Provide full access",
				Computed:            true,
			},
			"access_group_ipv4_query_only": schema.StringAttribute{
				MarkdownDescription: "Allow only control queries",
				Computed:            true,
			},
			"access_group_ipv4_serve": schema.StringAttribute{
				MarkdownDescription: "Provide server and query access",
				Computed:            true,
			},
			"access_group_ipv4_serve_only": schema.StringAttribute{
				MarkdownDescription: "Provide only server access",
				Computed:            true,
			},
			"vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "Specify non-default VRF",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf_name": schema.StringAttribute{
							MarkdownDescription: "Specify non-default VRF",
							Computed:            true,
						},
						"ipv6_peer": schema.StringAttribute{
							MarkdownDescription: "Provide full access",
							Computed:            true,
						},
						"ipv6_query_only": schema.StringAttribute{
							MarkdownDescription: "Allow only control queries",
							Computed:            true,
						},
						"ipv6_serve": schema.StringAttribute{
							MarkdownDescription: "Provide server and query access",
							Computed:            true,
						},
						"ipv6_serve_only": schema.StringAttribute{
							MarkdownDescription: "Provide only server access",
							Computed:            true,
						},
						"ipv4_peer": schema.StringAttribute{
							MarkdownDescription: "Provide full access",
							Computed:            true,
						},
						"ipv4_query_only": schema.StringAttribute{
							MarkdownDescription: "Allow only control queries",
							Computed:            true,
						},
						"ipv4_serve": schema.StringAttribute{
							MarkdownDescription: "Provide server and query access",
							Computed:            true,
						},
						"ipv4_serve_only": schema.StringAttribute{
							MarkdownDescription: "Provide only server access",
							Computed:            true,
						},
					},
				},
			},
			"authenticate": schema.BoolAttribute{
				MarkdownDescription: "Authenticate time sources",
				Computed:            true,
			},
			"auth_keys": schema.ListNestedAttribute{
				MarkdownDescription: "Authentication key for trusted time sources",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key_number": schema.Int64Attribute{
							MarkdownDescription: "Authentication key for trusted time sources",
							Computed:            true,
						},
						"md5_encrypted": schema.StringAttribute{
							MarkdownDescription: "Specify an encrypted key",
							Computed:            true,
						},
					},
				},
			},
			"broadcastdelay": schema.Int64Attribute{
				MarkdownDescription: "Estimated round-trip delay",
				Computed:            true,
			},
			"max_associations": schema.Int64Attribute{
				MarkdownDescription: "Set maximum number of associations",
				Computed:            true,
			},
			"trusted_keys": schema.ListNestedAttribute{
				MarkdownDescription: "Key numbers for trusted time sources",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key_number": schema.Int64Attribute{
							MarkdownDescription: "Key numbers for trusted time sources",
							Computed:            true,
						},
					},
				},
			},
			"update_calendar": schema.BoolAttribute{
				MarkdownDescription: "Periodically update calendar with NTP time",
				Computed:            true,
			},
			"log_internal_sync": schema.BoolAttribute{
				MarkdownDescription: "Logs internal synchronization changes",
				Computed:            true,
			},
			"source_interface_name": schema.StringAttribute{
				MarkdownDescription: "default interface",
				Computed:            true,
			},
			"source_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "Specify non-default VRF",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf_name": schema.StringAttribute{
							MarkdownDescription: "Specify non-default VRF",
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "default interface for the VRF",
							Computed:            true,
						},
					},
				},
			},
			"passive": schema.BoolAttribute{
				MarkdownDescription: "Enable the passive associations",
				Computed:            true,
			},
		},
	}
}

func (d *NTPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *NTPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NTPData

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
