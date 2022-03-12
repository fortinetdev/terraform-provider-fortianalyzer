---
subcategory: "System Route"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_route6"
description: |-
  Routing table configuration.
---

# fortianalyzer_system_route6
Routing table configuration.

## Argument Reference


The following arguments are supported:


* `device` - Gateway out interface.
* `dst` - Destination IP and mask for this route.
* `gateway` - Gateway IP for this route.
* `prio` - Entry number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{prio}}.

## Import

System Route6 can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_route6.labelname {{prio}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

