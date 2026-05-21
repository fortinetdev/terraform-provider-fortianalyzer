---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_apiratelimit"
description: |-
  System LogApi-Ratelimit
---

# fortianalyzer_system_log_apiratelimit
System LogApi-Ratelimit

## Argument Reference


The following arguments are supported:


* `read_limit` - Read-Limit.
* `write_limit` - Write-Limit.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogApiRatelimit can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_apiratelimit.labelname SystemLogApiRatelimit
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

