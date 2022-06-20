---
subcategory: "Fmupdate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_service"
description: |-
  Enable/disable services provided by the built-in FortiGuard.
---

# fortianalyzer_fmupdate_service
Enable/disable services provided by the built-in FortiGuard.

## Example Usage

```hcl
resource "fortianalyzer_fmupdate_service" "trname" {
  avips = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `avips` - Enable/disable the built-in FortiGuard to provide FortiGuard antivirus and IPS updates (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate Service can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_service.labelname FmupdateService
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

