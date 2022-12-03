---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_syslog"
description: |-
  Syslog servers.
---

# fortianalyzer_system_syslog
Syslog servers.

## Example Usage

```hcl
resource "fortianalyzer_system_syslog" "trname" {
  ip   = "1.1.1.1"
  port = 514
  name = "syslog"
}
```

## Argument Reference


The following arguments are supported:


* `ip` - Syslog server IP address or hostname.
* `local_cert` - Select local certificate used for secure connection.
* `name` - Syslog server name.
* `peer_cert_cn` - Certificate common name of syslog server. null or &apos;-&apos; means not check certificate CN of syslog server
* `port` - Syslog server port.
* `reliable` - Enable/disable reliable connection with syslog server. disable - Disable reliable connection with syslog server. enable - Enable reliable connection with syslog server. Valid values: `disable`, `enable`.

* `secure_connection` - Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Syslog can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_syslog.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

