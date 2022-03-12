---
subcategory: "System NTP"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_ntp_ntpserver"
description: |-
  NTP server.
---

# fortianalyzer_system_ntp_ntpserver
NTP server.

## Argument Reference


The following arguments are supported:


* `authentication` - Enable/disable MD5 authentication. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fosid` - Time server ID.
* `key` - Key for authentication.
* `key_id` - Key ID for authentication.
* `maxpoll` - Maximum poll interval in seconds as power of 2 (e.g. 6 means 64 seconds).
* `minpoll` - Minimum poll interval in seconds as power of 2 (e.g. 6 means 64 seconds).
* `ntpv3` - Enable/disable NTPv3. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `server` - IP address/hostname of NTP Server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System NtpNtpserver can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_ntp_ntpserver.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

