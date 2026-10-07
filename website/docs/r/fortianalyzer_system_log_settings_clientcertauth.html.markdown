---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_settings_clientcertauth"
description: |-
  System LogSettingsClient-Cert-Auth
---

# fortianalyzer_system_log_settings_clientcertauth
System LogSettingsClient-Cert-Auth

## Example Usage

```hcl
resource "fortianalyzer_system_log_settings_clientcertauth" "basic" {
  mode = "basic"
}
```

## Argument Reference


The following arguments are supported:


* `mode` - Mode. Valid values: `basic`, `strict`.

* `tls_port` - Tls-Port. Valid values: `both`, `514`, `6514`.

* `trusted_client` - Trusted-Client. The structure of `trusted_client` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `trusted_client` block supports:

* `certificate` - Certificate.
* `description` - Description.
* `domain` - Domain.
* `id` - Id.
* `type` - Type. Valid values: `certificate`, `domain`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogSettingsClientCertAuth can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_settings_clientcertauth.labelname SystemLogSettingsClientCertAuth
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

