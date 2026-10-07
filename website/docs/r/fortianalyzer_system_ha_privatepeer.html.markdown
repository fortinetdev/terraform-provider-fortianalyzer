---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_ha_privatepeer"
description: |-
  Peer.
---

# fortianalyzer_system_ha_privatepeer
Peer.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_ha_privatepeer" "test" {
  fosid         = 1
  ip            = "192.0.2.5"
  serial_number = "FAZVM00000000001"
}
```

## Argument Reference


The following arguments are supported:


* `fosid` - Id.
* `ip` - IP address of peer.
* `ip6` - IP address (V6) of peer.
* `serial_number` - Serial number of peer.
* `status` - Peer admin status. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System HaPrivatePeer can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_ha_privatepeer.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

