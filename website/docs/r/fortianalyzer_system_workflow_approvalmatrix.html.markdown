---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_workflow_approvalmatrix"
description: |-
  workflow approval matrix.
---

# fortianalyzer_system_workflow_approvalmatrix
workflow approval matrix.

## Argument Reference


The following arguments are supported:


* `adom_name` - Adom Name
* `approver` - Approver. The structure of `approver` block is documented below.
* `mail_server` - Notify mail server id.
* `notify` - Notify users
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `approver` block supports:

* `member` - Member of approver.
* `seq_num` - Entry number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{adom_name}}.

## Import

System WorkflowApprovalMatrix can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_workflow_approvalmatrix.labelname {{adom_name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

