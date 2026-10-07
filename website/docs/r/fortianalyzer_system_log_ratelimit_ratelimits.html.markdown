---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_ratelimit_ratelimits"
description: |-
  Per device or ADOM log rate limits.
---

# fortianalyzer_system_log_ratelimit_ratelimits
Per device or ADOM log rate limits.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_log_ratelimit_ratelimits" "test" {
  filter_type = "adom"
  fosid       = 3
  ratelimit   = 100
  filter      = "root"
  depends_on  = [fortianalyzer_system_log_ratelimit.trname]
}

resource "fortianalyzer_system_log_ratelimit" "trname" {
  mode = "manual"
}
```

## Argument Reference


The following arguments are supported:


* `filter` - Device or ADOM filter according to filter-type setting, wildcard expression supported.
* `filter_type` - Device filter type. devid - Device ID. adom - ADOM name. Valid values: `devid`, `adom`.

* `fosid` - Filter ID.
* `ratelimit` - Maximum log rate limit.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogRatelimitRatelimits can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_ratelimit_ratelimits.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

