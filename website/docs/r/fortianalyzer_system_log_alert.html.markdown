---
subcategory: "System Log"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_alert"
description: |-
  Log based alert settings.
---

# fortianalyzer_system_log_alert
Log based alert settings.

## Example Usage

```hcl
resource "fortianalyzer_system_log_alert" "trname" {
  max_alert_count = 10000
}
```

## Argument Reference


The following arguments are supported:


* `max_alert_count` - Maximum number of alerts supported.
* `min_severity_to_raise_incident_by_grouping` - Min-Severity-To-Raise-Incident-By-Grouping. Valid values: `none`, `critical`, `high`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogAlert can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_alert.labelname SystemLogAlert
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

