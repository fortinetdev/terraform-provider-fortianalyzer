---
subcategory: "System Certificate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_certificate_ca"
description: |-
  CA certificate.
---

# fortianalyzer_system_certificate_ca
CA certificate.

## Example Usage

```hcl
# Supply your own PEM-encoded certificate at the referenced path.
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_certificate_ca" "trname" {

  comment = "Default CA certificate"
  name    = "testca"
  ca      = [file("${path.module}/example-ca.pem")]
}
```

## Argument Reference


The following arguments are supported:


* `ca` - CA certificate.
* `comment` - CA certificate comment.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateCa can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_certificate_ca.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

