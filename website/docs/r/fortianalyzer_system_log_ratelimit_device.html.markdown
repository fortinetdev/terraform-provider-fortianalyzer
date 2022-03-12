---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_ratelimit_device"
description: |-
  Device log rate limit.
---

# fortianalyzer_system_log_ratelimit_device
Device log rate limit.

## Argument Reference


The following arguments are supported:


* `device` - Device(s) filter according to filter-type setting, wildcard expression supported.
* `filter_type` - Device filter type. devid - Device ID. Valid values: `devid`.

* `fosid` - Device filter ID.
* `ratelimit` - Maximum device log rate limit.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogRatelimitDevice can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_ratelimit_device.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

