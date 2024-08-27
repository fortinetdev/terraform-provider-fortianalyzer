---
subcategory: "System FortiView"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_fortiview_autocache"
description: |-
  FortiView auto-cache settings.
---

# fortianalyzer_system_fortiview_autocache
FortiView auto-cache settings.

## Example Usage

```hcl
resource "fortianalyzer_system_fortiview_autocache" "trname" {
  aggressive_fortiview = "disable"
  interval             = 168
  status               = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `aggressive_fortiview` - Enable/disable auto-cache on fortiview aggressively. disable - Disable the aggressive fortiview auto-cache. enable - Enable the aggressive fortiview auto-cache. Valid values: `disable`, `enable`.

* `incr_fortiview` - Enable/disable fortiview incremental cache. disable - Disable the fortiview incremental auto cache. enable - Enable the fortiview incremental auto cache. Valid values: `disable`, `enable`.

* `interval` - The time interval in hours for fortiview auto-cache.
* `status` - Enable/disable fortiview auto-cache. disable - Disable the fortiview auto-cache. enable - Enable the fortiview auto-cache. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FortiviewAutoCache can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_fortiview_autocache.labelname SystemFortiviewAutoCache
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

