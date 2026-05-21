---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_locallog_tacacsaccounting_filter"
description: |-
  System LocallogTacacs+AccountingFilter
---

# fortianalyzer_system_locallog_tacacsaccounting_filter
System LocallogTacacs+AccountingFilter

## Argument Reference


The following arguments are supported:


* `cli_cmd_audit` - Cli-Cmd-Audit. Valid values: `disable`, `enable`.

* `config_change_audit` - Config-Change-Audit. Valid values: `disable`, `enable`.

* `login_audit` - Login-Audit. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogTacacsAccountingFilter can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_locallog_tacacsaccounting_filter.labelname SystemLocallogTacacsAccountingFilter
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

