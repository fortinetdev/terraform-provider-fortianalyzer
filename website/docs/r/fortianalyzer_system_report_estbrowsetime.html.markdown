---
subcategory: "System Report"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_report_estbrowsetime"
description: |-
  Report estimated browse time settings
---

# fortianalyzer_system_report_estbrowsetime
Report estimated browse time settings

## Example Usage

```hcl
resource "fortianalyzer_system_report_estbrowsetime" "trname" {
  max_read_time = 180
  status        = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `max_read_time` - Read time threshold for each page view.
* `status` - Estimate browse time status. disable - Disable estimating browse time. enable - Enable estimating browse time. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ReportEstBrowseTime can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_report_estbrowsetime.labelname SystemReportEstBrowseTime
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

