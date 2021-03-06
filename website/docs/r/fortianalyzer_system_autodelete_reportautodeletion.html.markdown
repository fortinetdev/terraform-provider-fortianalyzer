---
subcategory: "System AutoDelete"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_autodelete_reportautodeletion"
description: |-
  Automatic deletion policy for reports.
---

# fortianalyzer_system_autodelete_reportautodeletion
Automatic deletion policy for reports.

## Example Usage

```hcl
resource "fortianalyzer_system_autodelete_reportautodeletion" "trname" {
  retention = "days"
  runat     = 1
  status    = "enable"
  value     = 1
}
```

## Argument Reference


The following arguments are supported:


* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoDeleteReportAutoDeletion can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_autodelete_reportautodeletion.labelname SystemAutoDeleteReportAutoDeletion
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

