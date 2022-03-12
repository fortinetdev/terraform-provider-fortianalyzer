---
subcategory: "System Certificate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_certificate_remote"
description: |-
  Remote certificate.
---

# fortianalyzer_system_certificate_remote
Remote certificate.

## Argument Reference


The following arguments are supported:


* `cert` - Remote certificate.
* `comment` - Remote certificate comment.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateRemote can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_certificate_remote.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

