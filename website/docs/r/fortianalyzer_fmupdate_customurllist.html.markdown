---
subcategory: "Fmupdate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_customurllist"
description: |-
  Configure the URL database for rating and filtering.
---

# fortianalyzer_fmupdate_customurllist
Configure the URL database for rating and filtering.

## Example Usage

```hcl
resource "fortianalyzer_fmupdate_customurllist" "trname" {
  db_selection = ["both"]
}
```

## Argument Reference


The following arguments are supported:


* `db_selection` - Manage the URL database (default = both). both - Support both custom-url and FortiGuard database. custom-url - Custom imported URL list. fortiguard-db - Fortinet's Fortiguard database. Valid values: `both`, `custom-url`, `fortiguard-db`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate CustomUrlList can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_customurllist.labelname FmupdateCustomUrlList
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

