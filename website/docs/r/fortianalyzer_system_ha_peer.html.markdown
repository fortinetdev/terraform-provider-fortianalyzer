---
subcategory: "System Global"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_ha_peer"
description: |-
  Peers.
---

# fortianalyzer_system_ha_peer
Peers.

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_ha" "test" {
  mode                      = "a-p"
  cfg_sync_hb_interval      = 80
  group_id                  = 255
  group_name                = "example-ha-group"
  password                  = ["ExamplePassword123!"]
  hb_interface              = "port1"
  hb_interval               = 20
  healthcheck               = ["DB"]
  initial_sync              = "true"
  initial_sync_threads      = 15
  load_balance              = "round-robin"
  log_sync                  = "enable"
  preferred_role            = "secondary"
  priority                  = 120
  private_clusterid         = 64
  private_file_quota        = 20480
  private_hb_interval       = 255
  private_hb_lost_threshold = 255
  private_mode              = "primary"
  unicast                   = "enable"
  vip_block {
    id            = 1
    vip           = "192.0.2.3"
    vip_interface = "port5"
  }

}

resource "fortianalyzer_system_ha_peer" "test" {
  fosid         = 3
  serial_number = "FAZVM00000000001"
  #addr = "test2"
  #addr_hb = "test2"
  #ip         = "192.0.2.4"
  depends_on = [fortianalyzer_system_ha.test]
}
```

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

