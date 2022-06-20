---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_socfabric"
description: |-
  SOC Fabric.
---

# fortianalyzer_system_socfabric
SOC Fabric.

## Example Usage

```hcl
resource "fortianalyzer_system_socfabric" "trname" {
  name       = "socfabric"
  port       = 6443
  psk        = ["paaword"]
  status     = "enable"
  role       = "member"
  supervisor = "1.1.1.1"
}
```

## Argument Reference


The following arguments are supported:


* `name` - Fabric name.
* `port` - communication port (1 - 65535).
* `psk` - Fabric auth pwd.
* `role` - Enable or Disable SOC Fabric. member - SOC Fabric member. supervisor - SOC Fabric supervisor. Valid values: `member`, `supervisor`.

* `secure_connection` - Enable or Disable SSL/TLS. disable - Disable SSL/TLS. enable - Enable SSL/TLS. Valid values: `disable`, `enable`.

* `status` - Enable or Disable SOC Fabric. disable - Disable SOC Fabric. enable - Enable SOC Fabric. Valid values: `disable`, `enable`.

* `supervisor` - IP/FQDN of supervisor.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SocFabric can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_socfabric.labelname SystemSocFabric
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

