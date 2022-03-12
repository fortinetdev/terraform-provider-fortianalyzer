---
subcategory: "System Certificate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_certificate_crl"
description: |-
  Certificate Revocation List.
---

# fortianalyzer_system_certificate_crl
Certificate Revocation List.

## Argument Reference


The following arguments are supported:


* `comment` - Comment of this Certificate Revocation List.
* `crl` - Certificate Revocation List.
* `http_url` - HTTP server URL for CRL auto-update
* `name` - Name.
* `update_interval` - CRL auto-update interval (in minutes)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateCrl can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_certificate_crl.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

