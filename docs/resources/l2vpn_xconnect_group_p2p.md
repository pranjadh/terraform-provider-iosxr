---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_l2vpn_xconnect_group_p2p Resource - terraform-provider-iosxr"
subcategory: "L2VPN"
description: |-
  This resource can manage the L2VPN Xconnect Group P2P configuration.
---

# iosxr_l2vpn_xconnect_group_p2p (Resource)

This resource can manage the L2VPN Xconnect Group P2P configuration.

## Example Usage

```terraform
resource "iosxr_l2vpn_xconnect_group_p2p" "example" {
  group_name        = "P2P"
  p2p_xconnect_name = "XC"
  description       = "My P2P Description"
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/2"
    }
  ]
  neighbor_evpn_evi_segment_routing_services = [
    {
      vpn_id                       = 4600
      service_id                   = 600
      segment_routing_srv6_locator = "LOC11"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_name` (String) Specify the group the cross connects belong to
- `p2p_xconnect_name` (String) Configure point to point cross connect commands

### Optional

- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `description` (String) Description for cross connect
- `device` (String) A device name from the provider configuration.
- `evpn_service_neighbors` (Attributes List) Specify service ID (used as local and remote ac-id) (see [below for nested schema](#nestedatt--evpn_service_neighbors))
- `evpn_target_neighbors` (Attributes List) Specify remote attachment circuit identifier (see [below for nested schema](#nestedatt--evpn_target_neighbors))
- `interfaces` (Attributes List) Specify (sub-)interface name to cross connect (see [below for nested schema](#nestedatt--interfaces))
- `ipv4_neighbors` (Attributes List) IPv4 (see [below for nested schema](#nestedatt--ipv4_neighbors))
- `ipv6_neighbors` (Attributes List) IPv6 (see [below for nested schema](#nestedatt--ipv6_neighbors))
- `neighbor_evpn_evi_segment_routing_services` (Attributes List) Specify service ID (used as local and remote ac-id) (see [below for nested schema](#nestedatt--neighbor_evpn_evi_segment_routing_services))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--evpn_service_neighbors"></a>
### Nested Schema for `evpn_service_neighbors`

Required:

- `service_id` (Number) Specify service ID (used as local and remote ac-id)
  - Range: `1`-`4294967294`
- `vpn_id` (Number) Ethernet VPN Identifier
  - Range: `1`-`65534`

Optional:

- `pw_class` (String) PW class template name to use


<a id="nestedatt--evpn_target_neighbors"></a>
### Nested Schema for `evpn_target_neighbors`

Required:

- `remote_ac_id` (Number) Specify remote attachment circuit identifier
  - Range: `1`-`4294967294`
- `source` (Number) Specify source attachment circuit identifier
  - Range: `1`-`4294967294`
- `vpn_id` (Number) Ethernet VPN Identifier
  - Range: `1`-`65534`

Optional:

- `pw_class` (String) PW class template name to use


<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Required:

- `interface_name` (String) Specify (sub-)interface name to cross connect


<a id="nestedatt--ipv4_neighbors"></a>
### Nested Schema for `ipv4_neighbors`

Required:

- `address` (String) IPv4
- `pw_id` (Number) Specify the pseudowire id
  - Range: `1`-`4294967295`

Optional:

- `pw_class` (String) PW class template name to use for this XC


<a id="nestedatt--ipv6_neighbors"></a>
### Nested Schema for `ipv6_neighbors`

Required:

- `address` (String) IPv6
- `pw_id` (Number) Specify the pseudowire id
  - Range: `1`-`4294967295`

Optional:

- `pw_class` (String) PW class template name to use for this XC


<a id="nestedatt--neighbor_evpn_evi_segment_routing_services"></a>
### Nested Schema for `neighbor_evpn_evi_segment_routing_services`

Required:

- `service_id` (Number) Specify service ID (used as local and remote ac-id)
  - Range: `1`-`4294967294`
- `vpn_id` (Number) Ethernet VPN Identifier
  - Range: `1`-`65534`

Optional:

- `segment_routing_srv6_locator` (String) PW locator to use for EVPN SID allocation

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_l2vpn_xconnect_group_p2p.example "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=P2P]/p2ps/p2p[p2p-xconnect-name=XC]"
```
