---
subcategory: "Fmupdate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_publicnetwork"
description: |-
  Enable/disable access to the public FortiGuard.
---

# fortianalyzer_fmupdate_publicnetwork
Enable/disable access to the public FortiGuard.

## Example Usage

```hcl
resource "fortianalyzer_fmupdate_publicnetwork" "trname" {
  status = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `status` - Enable/disable public network (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `update_server_location` - Update-Server-Location. Valid values: `global`, `usa`, `eu`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate Publicnetwork can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_publicnetwork.labelname FmupdatePublicnetwork
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

