---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_bgp_neighbors_address_family Resource - terraform-provider-iosxr"
subcategory: "BGP"
description: |-
  This resource can manage the Router BGP Neighbors Address Family configuration.
---

# iosxr_router_bgp_neighbors_address_family (Resource)

This resource can manage the Router BGP Neighbors Address Family configuration.

## Example Usage

```terraform
resource "iosxr_router_bgp_neighbors_address_family" "example" {
  as_number                                                 = "65001"
  neighbor_address                                          = "10.1.1.2"
  af_name                                                   = "vpnv4-unicast"
  import_stitching_rt_re_originate_stitching_rt             = true
  route_reflector_client_inheritance_disable                = true
  advertise_vpnv4_unicast_enable_re_originated_stitching_rt = true
  next_hop_self_inheritance_disable                         = true
  encapsulation_type_srv6                                   = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `af_name` (String) Enter Address Family command mode
  - Choices: `all-address-family`, `ipv4-flowspec`, `ipv4-labeled-unicast`, `ipv4-mdt`, `ipv4-multicast`, `ipv4-mvpn`, `ipv4-rt-filter`, `ipv4-sr-policy`, `ipv4-tunnel`, `ipv4-unicast`, `ipv6-flowspec`, `ipv6-labeled-unicast`, `ipv6-multicast`, `ipv6-mvpn`, `ipv6-sr-policy`, `ipv6-unicast`, `l2vpn-evpn`, `l2vpn-mspw`, `l2vpn-vpls-vpws`, `link-state-link-state`, `vpnv4-flowspec`, `vpnv4-multicast`, `vpnv4-unicast`, `vpnv6-flowspec`, `vpnv6-multicast`, `vpnv6-unicast`
- `as_number` (String) bgp as-number
- `neighbor_address` (String) Neighbor address

### Optional

- `advertise_vpnv4_unicast_enable_re_originated_stitching_rt` (Boolean) Advertise re-originated and local routes with stitching Route-Targets
- `device` (String) A device name from the provider configuration.
- `encapsulation_type_srv6` (Boolean) SRv6 encapsulation
- `import_stitching_rt_re_originate_stitching_rt` (Boolean) Reoriginate imported routes by attaching stitching RTs
- `next_hop_self_inheritance_disable` (Boolean) Prevent next-hop-self from being inherited from the parent
- `route_reflector_client_inheritance_disable` (Boolean) Prevent route-reflector-client from being inherited from the parent

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_router_bgp_neighbors_address_family.example "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/neighbors/neighbor[neighbor-address=10.1.1.2]/address-families/address-family[af-name=vpnv4-unicast]"
```