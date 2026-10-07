---
subcategory: "System SQL"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_sql_tsindexfield"
description: |-
  List of SQL text search index fields.
---

# fortianalyzer_system_sql_tsindexfield
List of SQL text search index fields.

## Example Usage

```hcl
//import first: terraform import fortianalyzer_system_sql_tsindexfield.trname FGT-app-ctrl
resource "fortianalyzer_system_sql_tsindexfield" "trname" {
  category = "FGT-app-ctrl"
  value    = "action,app,dstip,dstport,group,hostname,service,srcip,user"
}
```

## Argument Reference


The following arguments are supported:


* `category` - Category of text search index fields.
* `value` - Fields of text search index.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{category}}.

## Import

System SqlTsIndexField can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_sql_tsindexfield.labelname {{category}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

