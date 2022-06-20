---
subcategory: "Fmupdate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_avips_advancedlog"
description: |-
  Enable/disable logging of FortiGuard antivirus and IPS update packages received by FortiManager's built-in FortiGuard.
---

# fortianalyzer_fmupdate_avips_advancedlog
Enable/disable logging of FortiGuard antivirus and IPS update packages received by FortiManager's built-in FortiGuard.

## Example Usage

```hcl
resource "fortianalyzer_fmupdate_avips_advancedlog" "trname" {
  log_fortigate = "disable"
  log_server    = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `log_fortigate` - Enable/disable logging of FortiGuard antivirus and IPS service updates of FortiGate devices (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `log_server` - Enable/disable logging of update packages received by the build-in FortiGuard server (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate AvIpsAdvancedLog can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_avips_advancedlog.labelname FmupdateAvIpsAdvancedLog
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

