---
subcategory: "System Alert"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_alertconsole"
description: |-
  Alert console.
---

# fortianalyzer_system_alertconsole
Alert console.

## Example Usage

```hcl
resource "fortianalyzer_system_alertconsole" "trname" {
  period         = 7
  severity_level = ["emergency"]
}
```

## Argument Reference


The following arguments are supported:


* `period` - Alert console keeps alerts for this period. 1 - 1 day. 2 - 2 days. 3 - 3 days. 4 - 4 days. 5 - 5 days. 6 - 6 days. 7 - 7 days. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`.

* `severity_level_unitary` - Alert console keeps alerts of this and higher severity. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notify - Notify level. information - Information level. debug - Debug level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notify`, `information`, `debug`.

* `severity_level` - Alert console keeps alerts of this and higher severity. debug - Debug level. information - Information level. notify - Notify level. warning - Warning level. error - Error level. critical - Critical level. alert - Alert level. emergency - Emergency level. Valid values: `debug`, `information`, `notify`, `warning`, `error`, `critical`, `alert`, `emergency`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AlertConsole can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_alertconsole.labelname SystemAlertConsole
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

