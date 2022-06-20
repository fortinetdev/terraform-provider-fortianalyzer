---
subcategory: "System SNMP"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_snmp_sysinfo"
description: |-
  SNMP configuration.
---

# fortianalyzer_system_snmp_sysinfo
SNMP configuration.

## Example Usage

```hcl
resource "fortianalyzer_system_snmp_sysinfo" "trname" {
  fortianalyzer_legacy_sysoid          = "disable"
  status                               = "disable"
  trap_cpu_high_exclude_nice_threshold = 80
  trap_high_cpu_threshold              = 80
  trap_low_memory_threshold            = 80
}
```

## Argument Reference


The following arguments are supported:


* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engineID string (maximum 24 characters).
* `fortianalyzer_legacy_sysoid` - Enable legacy FortiAnalyzer sysObjectOID. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `location` - System location.
* `status` - Enable/disable SNMP. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `trap_cpu_high_exclude_nice_threshold` - SNMP trap for CPU usage threshold (exclude NICE processes).
* `trap_high_cpu_threshold` - SNMP trap for CPU usage threshold.
* `trap_low_memory_threshold` - SNMP trap for memory usage threshold.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SnmpSysinfo can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_snmp_sysinfo.labelname SystemSnmpSysinfo
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

