---
subcategory: "Device Manager"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_dvmdb_group"
description: |-
  Device group table.
---

# fortianalyzer_dvmdb_group
Device group table.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cluster_type` - Cluster_Type. Valid values: `unknown`, `vwan`, `sase`.

* `desc` - Desc.
* `fosid` - Id.
* `metafields` - Default metafields: none.
* `name` - Name.
* `os_type` - Os_Type. Valid values: `unknown`, `fos`, `fsw`, `foc`, `fml`, `faz`, `fwb`, `fch`, `fct`, `log`, `fmg`, `fsa`, `fdd`, `fac`, `fpx`, `fna`, `ffw`, `fsr`, `fad`, `fdc`, `fap`, `fxt`.

* `type` - Type. Valid values: `normal`, `default`, `auto`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dvmdb Group can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_dvmdb_group.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
