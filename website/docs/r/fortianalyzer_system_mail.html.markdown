---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_mail"
description: |-
  Alert emails.
---

# fortianalyzer_system_mail
Alert emails.

## Example Usage

```hcl
resource "fortianalyzer_system_mail" "trname" {
  auth_type = "psk"
  fosid     = 1
  passwd    = ["password"]
  port      = 25
  server    = "1.1.1.1"
  user      = "user"
}
```

## Argument Reference


The following arguments are supported:


* `auth` - Enable authentication. disable - Disable authentication. enable - Enable authentication. Valid values: `disable`, `enable`.

* `auth_type` - SMTP authentication type. psk - Use username and password to authenticate. certificate - Use local certificate to authenticate. Valid values: `psk`, `certificate`.

* `from` - Username for MAIL FROM.
* `fosid` - Mail Service ID.
* `local_cert` - SMTP local certificate.
* `passwd` - SMTP account password.
* `port` - SMTP server port.
* `secure_option` - Communication secure option. default - Try STARTTLS, proceed as plain text communication otherwise. none - Communication will be in plain text format. smtps - Communication will be protected by SMTPS. starttls - Communication will be protected by STARTTLS. Valid values: `default`, `none`, `smtps`, `starttls`.

* `server` - SMTP server.
* `ssl_protocol` - set the lowest SSL protocol version for connection to mail server. follow-global-ssl-protocol - Follow system.global.global-ssl-protocol setting (default). sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version. tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: `follow-global-ssl-protocol`, `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.

* `user` - SMTP account username.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Mail can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_mail.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

