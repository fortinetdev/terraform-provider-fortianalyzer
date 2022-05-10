---
subcategory: "System Log"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_devicedisable"
description: |-
  Disable client device logging.
---

# fortianalyzer_system_log_devicedisable
Disable client device logging.

## Argument Reference


The following arguments are supported:


* `ttl` - Time to Live
* `device` - Device to be disabled logging
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

