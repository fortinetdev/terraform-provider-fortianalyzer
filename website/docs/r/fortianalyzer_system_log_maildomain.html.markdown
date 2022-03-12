---
subcategory: "System Log"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_maildomain"
description: |-
  FortiMail domain setting.
---

# fortianalyzer_system_log_maildomain
FortiMail domain setting.

## Argument Reference


The following arguments are supported:


* `devices` - Devices for domain to vdom mapping
* `domain` - FortiMail domain
* `fosid` - ID of FortiMail domain.
* `vdom` - Virtual Domain name mapping to FortiMail domain


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogMailDomain can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_maildomain.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

