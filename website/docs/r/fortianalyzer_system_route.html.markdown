---
subcategory: "System Route"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_route"
description: |-
  Routing table configuration.
---

# fortianalyzer_system_route
Routing table configuration.

## Example Usage

```hcl
resource "fortianalyzer_system_route" "trname" {
  device  = "port1"
  dst     = ["172.10.21.0", "255.255.255.0"]
  gateway = "172.10.21.1"
  seq_num = 2
}
```

## Argument Reference


The following arguments are supported:


* `device` - Gateway out interface.
* `dst` - Destination IP and mask for this route.
* `gateway` - Gateway IP for this route.
* `seq_num` - Entry number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

System Route can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_route.labelname {{seq_num}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

