---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_locallog_tacacsaccounting_setting"
description: |-
  System LocallogTacacs+AccountingSetting
---

# fortianalyzer_system_locallog_tacacsaccounting_setting
System LocallogTacacs+AccountingSetting

## Argument Reference


The following arguments are supported:


* `status` - Status. Valid values: `disable`, `enable`.

* `tacacs_name` - Tacacs-Name.
* `timeout` - Timeout.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogTacacsAccountingSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_locallog_tacacsaccounting_setting.labelname SystemLocallogTacacsAccountingSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

