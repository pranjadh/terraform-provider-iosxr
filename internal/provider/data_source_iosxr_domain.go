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
	_ datasource.DataSource              = &DomainDataSource{}
	_ datasource.DataSourceWithConfigure = &DomainDataSource{}
)

func NewDomainDataSource() datasource.DataSource {
	return &DomainDataSource{}
}

type DomainDataSource struct {
	client *client.Client
}

func (d *DomainDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (d *DomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Domain configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"domains": schema.ListNestedAttribute{
				MarkdownDescription: "A domain name",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"domain_name": schema.StringAttribute{
							MarkdownDescription: "A domain name",
							Computed:            true,
						},
						"order": schema.Int64Attribute{
							MarkdownDescription: "This is used to sort the servers in the order of precedence",
							Computed:            true,
						},
					},
				},
			},
			"lookup_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable Domain Name System hostname translation",
				Computed:            true,
			},
			"lookup_source_interface": schema.StringAttribute{
				MarkdownDescription: "Specify source interface for DNS resolver",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Define the default domain name",
				Computed:            true,
			},
			"ipv4_hosts": schema.ListNestedAttribute{
				MarkdownDescription: "Name of host",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							MarkdownDescription: "Name of host",
							Computed:            true,
						},
						"ip_address": schema.ListAttribute{
							MarkdownDescription: "Host IP address (maximum of 8)",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"name_servers": schema.ListNestedAttribute{
				MarkdownDescription: "Specify address of name server to use",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "Specify address of name server to use",
							Computed:            true,
						},
						"order": schema.Int64Attribute{
							MarkdownDescription: "This is used to sort the servers in the order of precedence",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_hosts": schema.ListNestedAttribute{
				MarkdownDescription: "Name of host",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							MarkdownDescription: "Name of host",
							Computed:            true,
						},
						"ipv6_address": schema.ListAttribute{
							MarkdownDescription: "IPv6 name or address (maximum four addresses)",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"multicast": schema.StringAttribute{
				MarkdownDescription: "Define the domain name for multicast address lookups",
				Computed:            true,
			},
			"default_flows_disable": schema.BoolAttribute{
				MarkdownDescription: "disables default flows programming",
				Computed:            true,
			},
		},
	}
}

func (d *DomainDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *DomainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DomainData

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
