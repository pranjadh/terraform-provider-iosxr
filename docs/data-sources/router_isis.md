---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_isis Data Source - terraform-provider-iosxr"
subcategory: "ISIS"
description: |-
  This data source can read the Router ISIS configuration.
---

# iosxr_router_isis (Data Source)

This data source can read the Router ISIS configuration.

## Example Usage

```terraform
data "iosxr_router_isis" "example" {
  process_id = "P1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `process_id` (String) Process ID

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `affinity_maps` (Attributes List) Affinity map configuration (see [below for nested schema](#nestedatt--affinity_maps))
- `distribute_link_state_instance_id` (Number) Set distribution process instance identifier
- `distribute_link_state_level` (Number) Set distribution for one level only
- `distribute_link_state_throttle` (Number) Set throttle update in seconds
- `flex_algos` (Attributes List) Flex Algorithm definition (see [below for nested schema](#nestedatt--flex_algos))
- `id` (String) The path of the retrieved object.
- `interfaces` (Attributes List) Enter the IS-IS interface configuration submode (see [below for nested schema](#nestedatt--interfaces))
- `is_type` (String) Area type (level)
- `log_adjacency_changes` (Boolean) Enable logging adjacency state changes
- `lsp_gen_interval_initial_wait` (Number) Initial delay before generating an LSP
- `lsp_gen_interval_maximum_wait` (Number) Maximum delay before generating an LSP
- `lsp_gen_interval_secondary_wait` (Number) Secondary delay before generating an LSP
- `lsp_password_keychain` (String) Specifies a Key Chain name will follow
- `lsp_refresh_interval` (Number) Set LSP refresh interval
- `max_lsp_lifetime` (Number) Set maximum LSP lifetime
- `nets` (Attributes List) A Network Entity Title (NET) for this process (see [below for nested schema](#nestedatt--nets))
- `nsf_cisco` (Boolean) Cisco Proprietary NSF restart
- `nsf_ietf` (Boolean) IETF NSF restar
- `nsf_interface_expires` (Number) # of times T1 can expire waiting for the restart ACK
- `nsf_interface_timer` (Number) Timer used to wait for a restart ACK (seconds)
- `nsf_lifetime` (Number) Maximum route lifetime following restart (seconds)
- `nsr` (Boolean) Enable NSR
- `set_overload_bit_levels` (Attributes List) Set overload-bit for one level only (see [below for nested schema](#nestedatt--set_overload_bit_levels))

<a id="nestedatt--affinity_maps"></a>
### Nested Schema for `affinity_maps`

Read-Only:

- `bit_position` (Number) Bit position for affinity attribute value
- `name` (String) Affinity map configuration


<a id="nestedatt--flex_algos"></a>
### Nested Schema for `flex_algos`

Read-Only:

- `advertise_definition` (Boolean) Advertise the Flex-Algo Definition
- `algorithm_number` (Number) Flex Algorithm definition
- `metric_type_delay` (Boolean) Use delay as metric


<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Read-Only:

- `circuit_type` (String) Configure circuit type for interface
- `hello_padding_disable` (Boolean) Disable hello-padding
- `hello_padding_sometimes` (Boolean) Enable hello-padding during adjacency formation only
- `interface_name` (String) Enter the IS-IS interface configuration submode
- `passive` (Boolean) Do not establish adjacencies over this interface
- `point_to_point` (Boolean) Treat active LAN interface as point-to-point
- `priority` (Number) Set priority for Designated Router election
- `shutdown` (Boolean) Shutdown IS-IS on this interface
- `suppressed` (Boolean) Do not advertise connected prefixes of this interface


<a id="nestedatt--nets"></a>
### Nested Schema for `nets`

Read-Only:

- `net_id` (String) A Network Entity Title (NET) for this process


<a id="nestedatt--set_overload_bit_levels"></a>
### Nested Schema for `set_overload_bit_levels`

Read-Only:

- `advertise_external` (Boolean) If overload-bit set advertise IP prefixes learned from other protocols
- `advertise_interlevel` (Boolean) If overload-bit set advertise IP prefixes learned from another ISIS level
- `level_id` (Number) Set overload-bit for one level only
- `on_startup_advertise_as_overloaded` (Boolean) Time in seconds to advertise ourself as overloaded after reboot
- `on_startup_advertise_as_overloaded_time_to_advertise` (Number) Time in seconds to advertise ourself as overloaded after reboot
- `on_startup_wait_for_bgp` (Boolean) Set overload bit on startup until BGP signals convergence, or timeout
