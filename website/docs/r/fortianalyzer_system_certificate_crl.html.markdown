---
subcategory: "System Certificate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_certificate_crl"
description: |-
  Certificate Revocation List.
---

# fortianalyzer_system_certificate_crl
Certificate Revocation List.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_certificate_crl" "test" {
  comment = "test"
  name    = "example-crl"
  # crl = [""]

  http_url        = "https://crl.example.com/example.crl"
  update_interval = 10
}
```

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

