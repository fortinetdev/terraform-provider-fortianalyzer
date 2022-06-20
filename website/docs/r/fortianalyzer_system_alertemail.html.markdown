---
subcategory: "System Alert"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_alertemail"
description: |-
  Configure alertemail.
---

# fortianalyzer_system_alertemail
Configure alertemail.

## Example Usage

```hcl
resource "fortianalyzer_system_alertemail" "trname" {
  authentication = "enable"
  fromaddress    = "from@gamil.com"
  fromname       = "from"
  smtppassword   = ["password"]
  smtpport       = 25
  smtpserver     = "1.1.1.1"
  smtpuser       = "user"
}
```

## Argument Reference


The following arguments are supported:


* `authentication` - Enable/disable authentication. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fromaddress` - SMTP from address.
* `fromname` - SMTP from user.
* `smtppassword` - SMTP server password.
* `smtpport` - SMTP server port.
* `smtpserver` - SMTP server address.
* `smtpuser` - SMTP server user.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Alertemail can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_alertemail.labelname SystemAlertemail
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

