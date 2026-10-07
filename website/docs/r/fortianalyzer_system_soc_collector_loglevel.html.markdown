---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_soc_collector_loglevel"
description: |-
  Log queue settings.
---

# fortianalyzer_system_soc_collector_loglevel
Log queue settings.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_soc_collector_loglevel" "test" {
  phmonitor = "warning"
}
```

## Argument Reference


The following arguments are supported:


* `all` - Set the log level for all. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phagentmanager` - Set the log level for phAgentManager. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phdiscover` - Set the log level for phDiscover. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `pheventforwarder` - Set the log level for phEventForwarder. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `pheventhandler` - Set the log level for phEventHandler. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `pheventpackager` - Set the log level for phEventPackager. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phmonitor` - Set the log level for phMonitor. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phparser` - Set the log level for phParser. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phperfmonitor` - Set the log level for phPerfMonitor. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SocCollectorLogLevel can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_soc_collector_loglevel.labelname SystemSocCollectorLogLevel
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

