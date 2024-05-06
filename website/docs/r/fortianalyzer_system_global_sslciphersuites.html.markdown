---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_global_sslciphersuites"
description: |-
  Configure preferred SSL/TLS cipher suites
---

# fortianalyzer_system_global_sslciphersuites
Configure preferred SSL/TLS cipher suites

## Argument Reference


The following arguments are supported:


* `cipher` - Cipher name
* `priority` - SSL/TLS cipher suites priority.
* `version` - SSL/TLS version the cipher suite can be used with. tls1.2-or-below - TLS 1.2 or below. tls1.3 - TLS 1.3 Valid values: `tls1.2-or-below`, `tls1.3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{priority}}.

## Import

System GlobalSslCipherSuites can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_global_sslciphersuites.labelname {{priority}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

