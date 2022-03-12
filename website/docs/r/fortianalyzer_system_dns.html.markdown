---
subcategory: "System Global"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_dns"
description: |-
  DNS configuration.
---

# fortianalyzer_system_dns
DNS configuration.

## Argument Reference


The following arguments are supported:


* `ip6_primary` - IPv6 primary DNS IP.
* `ip6_secondary` - IPv6 secondary DNS IP.
* `primary` - Primary DNS IP.
* `secondary` - Secondary DNS IP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dns can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_dns.labelname SystemDns
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

