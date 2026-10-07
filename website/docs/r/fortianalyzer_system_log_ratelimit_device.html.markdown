---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_ratelimit_device"
description: |-
  Device log rate limit.
---

# fortianalyzer_system_log_ratelimit_device
Device log rate limit.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
# 6.4.8, 6.4.9
resource "fortianalyzer_system_log_ratelimit_device" "test" {
  device      = "root"
  filter_type = "adom"
  fosid       = 1
  ratelimit   = 100
  depends_on  = [fortianalyzer_system_log_ratelimit.trname]
}

resource "fortianalyzer_system_log_ratelimit" "trname" {
  mode = "manual"
}
```

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

