---
subcategory: "Fmupdate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_fdssetting_pushoverride"
description: |-
  Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.
---

# fortianalyzer_fmupdate_fdssetting_pushoverride
Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.

## Example Usage

```hcl
resource "fortianalyzer_fmupdate_fdssetting_pushoverride" "trname" {
  ip     = "0.0.0.0"
  port   = 9443
  status = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `ip` - External or virtual IP address of the NAT device that will forward push messages to the FortiManager unit.
* `port` - Receiving port number on the NAT device (1 - 65535, default = 9443).
* `status` - Enable/disable push updates for clients (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FdsSettingPushOverride can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_fdssetting_pushoverride.labelname FmupdateFdsSettingPushOverride
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

