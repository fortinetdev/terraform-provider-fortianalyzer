---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_socfabric_trustedlist"
description: |-
  Pre-authorized security fabric nodes
---

# fortianalyzer_system_socfabric_trustedlist
Pre-authorized security fabric nodes

## Argument Reference


The following arguments are supported:


* `fosid` - Trusted list ID.
* `serial` - FAZ serial number(support wildcard).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SocFabricTrustedList can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_socfabric_trustedlist.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

