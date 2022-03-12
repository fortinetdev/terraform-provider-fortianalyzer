---
subcategory: "System Log"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_logfetch_serversettings"
description: |-
  Log-fetch server settings.
---

# fortianalyzer_system_logfetch_serversettings
Log-fetch server settings.

## Argument Reference


The following arguments are supported:


* `max_conn_per_session` - Max concurrent file download connections per session.
* `max_sessions` - Max concurrent fetch sessions.
* `session_timeout` - Fetch session timeout in minute.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogFetchServerSettings can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_logfetch_serversettings.labelname SystemLogFetchServerSettings
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

