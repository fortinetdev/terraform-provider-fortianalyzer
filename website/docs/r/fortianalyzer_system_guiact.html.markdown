---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_guiact"
description: |-
  System settings through GUI.
---

# fortianalyzer_system_guiact
System settings through GUI.

## Argument Reference


The following arguments are supported:


* `backup_all` - Backup all settings.
* `backup_conf` - Backup config file.
* `eventlog_msg` - Write event log.
* `eventlog_path` - Event log path.
* `reboot` - Reboot system.
* `reset2default` - Reset to factory default.
* `restore_all` - Restore all settings.
* `restore_conf` - Restore config file.
* `time` - Time.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Guiact can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_guiact.labelname SystemGuiact
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

