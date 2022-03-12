---
subcategory: "System LocalLog"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_locallog_fortianalyzer3_setting"
description: |-
  Settings for locallog to fortianalyzer.
---

# fortianalyzer_system_locallog_fortianalyzer3_setting
Settings for locallog to fortianalyzer.

## Argument Reference


The following arguments are supported:


* `peer_cert_cn` - Certificate common name of remote fortianalyzer.
* `reliable` - Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: `disable`, `enable`.

* `secure_connection` - Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: `disable`, `enable`.

* `server` - Remote FortiAnalyzer server FQDN, hostname, or IP address
* `severity` - Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Log to FortiAnalyzer status. disable - Log to FortiAnalyzer disabled. realtime - Log to FortiAnalyzer in realtime. upload - Log to FortiAnalyzer at schedule time. Valid values: `disable`, `realtime`, `upload`.

* `upload_time` - Time to upload local log files (hh:mm).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogFortianalyzer3Setting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_locallog_fortianalyzer3_setting.labelname SystemLocallogFortianalyzer3Setting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

