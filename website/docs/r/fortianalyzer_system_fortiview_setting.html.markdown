---
subcategory: "System FortiView"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_fortiview_setting"
description: |-
  FortiView settings.
---

# fortianalyzer_system_fortiview_setting
FortiView settings.

## Example Usage

```hcl
resource "fortianalyzer_system_fortiview_setting" "trname" {
  not_scanned_apps = "include"
  resolve_ip       = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `data_source` - Data soure of the fortiview query. auto - Data from hcache, and from logs in a flexible way. cache-only - Data from hcache only. log-and-cache - Data from logs and hcache. Valid values: `auto`, `cache-only`, `log-and-cache`.

* `not_scanned_apps` - Include/Exclude 'Not.Scanned' applications in FortiView. Set as 'exclude' if you want to filter out never scanned applications. exclude - Exclude 'Not.Scanned' applications in FortiView. include - Include 'Not.Scanned' applications in FortiView. Valid values: `exclude`, `include`.

* `query_run_mode` - Query-Run-Mode. Valid values: `auto`, `boost`.

* `resolve_ip` - Enable or disable resolving IP address to hostname in FortiView.  disable - Disable resolving IP address to hostname. enable - Enable resolving IP address to hostname. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FortiviewSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_fortiview_setting.labelname SystemFortiviewSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

