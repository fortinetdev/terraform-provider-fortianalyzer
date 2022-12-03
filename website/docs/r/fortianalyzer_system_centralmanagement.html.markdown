---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_centralmanagement"
description: |-
  Central management configuration.
---

# fortianalyzer_system_centralmanagement
Central management configuration.

## Example Usage

```hcl
resource "fortianalyzer_system_centralmanagement" "trname" {
  allow_monitor           = "disable"
  authorized_manager_only = "enable"
  enc_algorithm           = "default"
  fmg                     = "1.1.1.1"
  type                    = "fortimanager"
}
```

## Argument Reference


The following arguments are supported:


* `acctid` - Account ID.
* `allow_monitor` - Enable/disable remote monitor of device. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `authorized_manager_only` - Enable/disable restricted to authorized manager only. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `enc_algorithm` - SSL communication encryption algorithms. default - SSL communication with high and medium encryption algorithms low - SSL communication with low encryption algorithms high - SSL communication with high encryption algorithms Valid values: `default`, `low`, `high`.

* `fmg` - Address of Fortimanager (IP or FQDN).
* `mgmtid` - Management ID.
* `serial_number` - Serial number.
* `type` - Type of management server. fortimanager - FortiManager. Valid values: `fortimanager`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System CentralManagement can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_centralmanagement.labelname SystemCentralManagement
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

