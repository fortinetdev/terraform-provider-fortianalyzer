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

* `ssl_protocol` - set the lowest SSL protocol version for connection to syslog server. follow-global-ssl-protocol - Follow system.global.global-ssl-protocol setting (default). sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version. tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: `follow-global-ssl-protocol`, `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.



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

