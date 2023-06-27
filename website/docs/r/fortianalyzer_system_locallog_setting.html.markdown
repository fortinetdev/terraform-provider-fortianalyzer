---
subcategory: "System LocalLog"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_locallog_setting"
description: |-
  Settings for locallog logging.
---

# fortianalyzer_system_locallog_setting
Settings for locallog logging.

## Example Usage

```hcl
resource "fortianalyzer_system_locallog_setting" "trname" {
  log_interval_dev_no_logging = 1440
  log_interval_disk_full      = 5
  log_interval_gbday_exceeded = 1440
}
```

## Argument Reference


The following arguments are supported:


* `log_daemon_crash` - Send a logmsg when a daemon crashes. enable/disable disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `log_interval_adom_perf_stats` - Interval in minute for logging the event of adom perf stats.
* `log_interval_dev_no_logging` - Interval in minute for logging the event of no logs received from a device.
* `log_interval_disk_full` - Interval in minute for logging the event of disk full.
* `log_interval_gbday_exceeded` - Interval in minute for logging the event of the GB/Day license exceeded.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_locallog_setting.labelname SystemLocallogSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

