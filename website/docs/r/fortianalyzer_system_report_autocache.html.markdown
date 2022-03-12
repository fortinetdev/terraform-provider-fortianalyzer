---
subcategory: "System Report"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_report_autocache"
description: |-
  Report auto-cache settings.
---

# fortianalyzer_system_report_autocache
Report auto-cache settings.

## Argument Reference


The following arguments are supported:


* `aggressive_schedule` - Enable/disable auto-cache on schedule reports aggressively. disable - Disable the aggressive schedule auto-cache. enable - Enable the aggressive schedule auto-cache. Valid values: `disable`, `enable`.

* `order` - The order of which SQL log table is processed first. oldest-first - The oldest SQL log table is processed first. Valid values: `oldest-first`.

* `status` - Enable/disable sql report auto cache. disable - Disable the sql report auto-cache. enable - Enable the sql report auto-cache. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ReportAutoCache can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_report_autocache.labelname SystemReportAutoCache
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

