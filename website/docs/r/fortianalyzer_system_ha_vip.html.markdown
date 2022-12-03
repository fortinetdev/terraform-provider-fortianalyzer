---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_ha_vip"
description: |-
  VIPs.
---

# fortianalyzer_system_ha_vip
VIPs.

## Argument Reference


The following arguments are supported:


* `fosid` - Id.
* `status` - VIP enabled status. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `vip` - Virtual IP address for the HA
* `vip_interface` - Interface for configuring virtual IP address


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System HaVip can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_ha_vip.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

