---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_webproxy"
description: |-
  Configure system web proxy.
---

# fortianalyzer_system_webproxy
Configure system web proxy.

## Argument Reference


The following arguments are supported:


* `address` - web proxy address.
* `mode` - Web proxy mode (default = tunnel) proxy - proxy mode tunnel - tunnel mode (default) Valid values: `proxy`, `tunnel`.

* `password` - The password for the user name used for authentication.
* `port` - The port number of the web proxy (1 - 65535, default = 1080).
* `status` - Enable/disable system web proxy (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `username` - The user name used for authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System WebProxy can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_webproxy.labelname SystemWebProxy
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

