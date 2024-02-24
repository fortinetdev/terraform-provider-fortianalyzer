---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_connector"
description: |-
  Configure connector.
---

# fortianalyzer_system_connector
Configure connector.

## Example Usage

```hcl
resource "fortianalyzer_system_connector" "trname" {
  fsso_refresh_interval = 180
  fsso_sess_timeout     = 300
  px_svr_timeout        = 300
}
```

## Argument Reference


The following arguments are supported:


* `cloud_orchest_refresh_interval` - Cloud Orchestration refresh interval (300 - 1800 seconds).
* `conn_refresh_interval` - connector refresh interval (60 - 1800 seconds).
* `faznotify_msg_queue_max` - faznotify max queued message per connector (10 - 10000).
* `faznotify_msg_timeout` - faznotify message timeout (1 - 720 hours).
* `fsso_refresh_interval` - FSSO refresh interval (60 - 600 seconds).
* `fsso_sess_timeout` - FSSO session timeout (60 - 600 seconds).
* `px_refresh_interval` - pxGrid refresh interval (60 - 1800 seconds).
* `px_svr_timeout` - pxGrid server timeout (30 - 600 seconds).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Connector can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_connector.labelname SystemConnector
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

