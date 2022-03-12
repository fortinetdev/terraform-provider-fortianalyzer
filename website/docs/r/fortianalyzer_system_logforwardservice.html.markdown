---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_logforwardservice"
description: |-
  Log forwarding service.
---

# fortianalyzer_system_logforwardservice
Log forwarding service.

## Argument Reference


The following arguments are supported:


* `accept_aggregation` - Enable/disable accept log aggregation option. disable - Disable attribute function enable - Enable attribute function Valid values: `disable`, `enable`.

* `aggregation_disk_quota` - Aggregated device disk quota (MB) on server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogForwardService can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_logforwardservice.labelname SystemLogForwardService
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

