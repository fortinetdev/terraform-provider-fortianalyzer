---
subcategory: "System Certificate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_certificate_ssh"
description: |-
  SSH certificates and keys.
---

# fortianalyzer_system_certificate_ssh
SSH certificates and keys.

## Argument Reference


The following arguments are supported:


* `certificate` - SSH certificate.
* `comment` - SSH certificate comment.
* `name` - Name of SSH certificate.
* `private_key` - SSH private-key


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateSsh can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_certificate_ssh.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

