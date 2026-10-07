---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_soc_collector_logqueue"
description: |-
  Log queue settings.
---

# fortianalyzer_system_soc_collector_logqueue
Log queue settings.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_soc_collector_logqueue" "test" {
  config_queue_file_count_limit = 1000
}
```

## Argument Reference


The following arguments are supported:


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

System SocCollectorLogQueue can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_soc_collector_logqueue.labelname SystemSocCollectorLogQueue
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

