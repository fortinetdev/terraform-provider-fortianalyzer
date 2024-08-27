---
subcategory: "System Global"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_ha_peer"
description: |-
  Peers.
---

# fortianalyzer_system_ha_peer
Peers.

## Argument Reference


The following arguments are supported:


* `addr` - Address of peer for management and data.
* `addr_hb` - Address of peer&apos;s interface for heartbeat, set if different from ip. (needed only when using unicast)
* `fosid` - Id.
* `ip` - IP address of peer for management and data.
* `ip_hb` - IP address of peer's VIP interface for heartbeat, set if different from ip. (needed only when using unicast)
* `serial_number` - Serial number of peer.
* `status` - Peer enabled status. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System HaPeer can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_ha_peer.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

