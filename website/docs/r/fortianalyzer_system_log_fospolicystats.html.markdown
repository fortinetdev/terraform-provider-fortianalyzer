---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_fospolicystats"
description: |-
  FortiOS policy statistics settings.
---

# fortianalyzer_system_log_fospolicystats
FortiOS policy statistics settings.

## Argument Reference


The following arguments are supported:


* `retention_days` - Number of days for FortiOS policy stats data storage.
* `sampling_interval` - Interval to request policy stats data from FortiOS in minutes.
* `status` - Disable/Enable FortiOS policy statistics feature. disable - Disable querying FortiOS policy stats. enable - Enable querying FortiOS policy stats. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogFosPolicyStats can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_fospolicystats.labelname SystemLogFosPolicyStats
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

