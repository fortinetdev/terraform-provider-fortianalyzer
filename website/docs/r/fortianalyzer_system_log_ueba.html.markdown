---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_ueba"
description: |-
  UEBAsettings.
---

# fortianalyzer_system_log_ueba
UEBAsettings.

## Argument Reference


The following arguments are supported:


* `hostname_ep_unifier` - Hostname-Ep-Unifier. Valid values: `disable`, `enable`.

* `ip_only_ep` - Disable/Enable IP-only endpoint identification. disable - Disable IP-only endpoint identification. enable - Enable IP-only endpoint identification. Valid values: `disable`, `enable`.

* `ip_unique_scope` - set ip-unique-scope. adom - set ip-unique-scope to adom. effective only when ip-only-endpoint is enable vdom - set ip-unique-scope to vdom. effective only when ip-only-endpoint is enable Valid values: `adom`, `vdom`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogUeba can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_ueba.labelname SystemLogUeba
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

