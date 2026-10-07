---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_soc_collector"
description: |-
  FAZ SOC collector settings.
---

# fortianalyzer_system_soc_collector
FAZ SOC collector settings.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_soc_collector" "test" {
  status    = "disable"
  perf_mode = "minimum"
}
```

## Argument Reference


The following arguments are supported:


* `client_id` - SIEM client ID.
* `client_secret` - SIEM client secret.
* `log_level` - Log-Level. The structure of `log_level` block is documented below.
* `log_queue` - Log-Queue. The structure of `log_queue` block is documented below.
* `name` - SIEM collector name.
* `org_name` - organization name.
* `perf_mode` - Set the resource limitation profile for soc collector processes. minimum - Restrict system resources for the SIEM collector. balanced - Moderate resource usage for balanced performance of the SIEM collector. aggressive - Allow more system resources for peak performance of the SIEM collector. Valid values: `minimum`, `balanced`, `aggressive`.

* `status` - Enable or Disable SIEM collector. disable - Disable SIEM collector. enable - Enable SIEM collector. Valid values: `disable`, `enable`.

* `super_addr` - SIEM supervisor IP/FQDN.
* `super_port` - SIEM supervisor port (1 - 65535).
* `verify_ssl` - Enable or Disable verify-ssl on SIEM collector connection. disable - Disable verify-ssl on connection. enable - Enable verify-ssl on connection. Valid values: `disable`, `enable`.


The `log_level` block supports:

* `all` - Set the log level for all. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phagentmanager` - Set the log level for phAgentManager. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phdiscover` - Set the log level for phDiscover. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `pheventforwarder` - Set the log level for phEventForwarder. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `pheventhandler` - Set the log level for phEventHandler. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `pheventpackager` - Set the log level for phEventPackager. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phmonitor` - Set the log level for phMonitor. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phparser` - Set the log level for phParser. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.

* `phperfmonitor` - Set the log level for phPerfMonitor. none - NONE. critical - CRITICAL. error - ERROR. warning - WARNING. notification - NOTIF. info - INFO. debug - DEBUG. Valid values: `none`, `critical`, `error`, `warning`, `notification`, `info`, `debug`.


The `log_queue` block supports:

* `config_queue_file_count_limit` - Set the maximum number of device-config files that can be queued by soc collector processes.
* `config_queue_size_limit_mb` - Set the maximum total size of device config files that can be queued by soc collector processes.
* `linux_queue_file_count_limit` - Set the maximum number of linux event files that can be queued by soc collector processes.
* `linux_queue_size_limit_mb` - Set the maximum total size of linux event files that can be queued by soc collector processes.
* `logs_queue_file_count_limit` - Set the maximum number of other (i.e. non-windows, non-linux) event files that can be queued by soc collector processes.
* `logs_queue_size_limit_mb` - Set the maximum total size of other (i.e. non-windows, non-linux) event files that can be queued by soc collector processes.
* `win_queue_file_count_limit` - Set the maximum number of windows event files that can be queued by soc collector processes.
* `win_queue_size_limit_mb` - Set the maximum total size of windows event files that can be queued by soc collector processes.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SocCollector can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_soc_collector.labelname SystemSocCollector
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

