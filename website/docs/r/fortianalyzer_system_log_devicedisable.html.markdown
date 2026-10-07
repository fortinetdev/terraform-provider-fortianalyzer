---
subcategory: "System Log"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_devicedisable"
description: |-
  Disable client device logging.
---

# fortianalyzer_system_log_devicedisable
Disable client device logging.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_log_devicedisable" "test" {
  # ttl    = 20000
  device = "test"
  expire = "2025-06-24 15:07:39"
  fosid  = 2
}
```

## Argument Reference


The following arguments are supported:


* `ttl` - Time to Live
* `device` - Device to be disabled logging
* `expire` - Expire.
* `fosid` - ID of device logging disable entry.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogDeviceDisable can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_devicedisable.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

