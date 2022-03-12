---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_connector"
description: |-
  Configure connector.
---

# fortianalyzer_system_connector
Configure connector.

## Argument Reference


The following arguments are supported:


* `conn_refresh_interval` - connector refresh interval (60 - 1800 seconds).
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

