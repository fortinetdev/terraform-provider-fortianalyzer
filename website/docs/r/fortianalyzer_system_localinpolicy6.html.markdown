---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_localinpolicy6"
description: |-
  IPv6 local in policy configuration.
---

# fortianalyzer_system_localinpolicy6
IPv6 local in policy configuration.

## Argument Reference


The following arguments are supported:


* `action` - Action performed on traffic matching this policy. drop - Drop traffic matching this policy (default). reject - Reject traffic matching this policy. accept - Allow traffic matching this policy. Valid values: `drop`, `reject`, `accept`.

* `description` - Description.
* `dport_block` - Dport. The structure of `dport_block` block is documented below.
* `dst_block` - Dst. The structure of `dst_block` block is documented below.
* `dport` - Destination port number (0 for all).
* `dst` - Destination IP and prefix.
* `fosid` - Entry number.
* `intf_block` - Intf. The structure of `intf_block` block is documented below.
* `intf` - Incoming interface name.
* `protocol` - Traffic protocol. tcp - TCP only. udp - UDP only. tcp_udp - Both TCP and UDP. Valid values: `tcp`, `udp`, `tcp_udp`.

* `src_block` - Src. The structure of `src_block` block is documented below.
* `src` - Source IP and prefix.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dport_block` block supports:

* `dport_value` - Dport-Value.

The `dst_block` block supports:

* `src_ip` - Src-Ip.

The `intf_block` block supports:

* `intf_name` - Intf-Name.

The `src_block` block supports:

* `src_ip` - Src-Ip.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LocalInPolicy6 can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_localinpolicy6.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

