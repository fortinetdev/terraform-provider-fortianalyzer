---
subcategory: "System Global"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_ha"
description: |-
  HA configuration.
---

# fortianalyzer_system_ha
HA configuration.

## Argument Reference


The following arguments are supported:


* `cfg_sync_hb_interval` - Config sync heartbeat interval (1 - 80).
* `group_id` - HA group ID (1 - 255).
* `group_name` - HA group name.
* `hb_interface` - Interface for heartbeat.
* `hb_interval` - Heartbeat interval (1 - 20).
* `healthcheck` - Healthchecking options. DB - Check database is running fault-test - temp fault test Valid values: `DB`, `fault-test`.

* `initial_sync` - Need to sync data from primary before up as an HA member. false - False. true - True. Valid values: `false`, `true`.

* `initial_sync_threads` - Number of threads used for initial sync (1-15).
* `load_balance` - Load balance to secondaries. disable - Disable load-balance to secondaries. round-robin - Round-Robin mode. Valid values: `disable`, `round-robin`.

* `log_sync` - Sync logs to secondary FortiAnalyzers. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `mode` - Standalone or HA (a-p) mode standalone - Standalone mode. a-p - Active-Passive mode. Valid values: `standalone`, `a-p`.

* `password` - HA group password.
* `peer` - Peer. The structure of `peer` block is documented below.
* `preferred_role` - Preferred role, runtime role may be different. secondary - Prefer secondary mode, FAZ can only become primary after data-sync is done. primary - Prefer primary mode, FAZ can become primary if there's no existing primary. Valid values: `secondary`, `primary`.

* `priority` - Set the runtime priority between 80 (lowest) - 120 (highest).
* `private_clusterid` - Cluster ID range (1 - 64).
* `private_file_quota` - File quota in MB (2048 - 20480).
* `private_hb_interval` - Heartbeat interval (1 - 255).
* `private_hb_lost_threshold` - Heartbeat lost threshold (1 - 255).
* `private_mode` - Mode. standalone - Standalone. primary - Primary. secondary - Secondary. Valid values: `standalone`, `primary`, `secondary`.

* `private_password` - Group password.
* `private_peer` - Private-Peer. The structure of `private_peer` block is documented below.
* `unicast` - Use unicast for HA heartbeat. disable - HA heartbeat through multicast. enable - HA heartbeat through unicast. Valid values: `disable`, `enable`.

* `vip` - Virtual IP address for the HA
* `vip_interface` - Interface for configuring virtual IP address
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `peer` block supports:

* `id` - Id.
* `ip` - IP address of peer for management and data.
* `ip_hb` - IP address of peer's VIP interface for heartbeat, set if different from ip. (needed only when using unicast)
* `serial_number` - Serial number of peer.
* `status` - Peer enabled status. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.


The `private_peer` block supports:

* `id` - Id.
* `ip` - IP address of peer.
* `ip6` - IP address (V6) of peer.
* `serial_number` - Serial number of peer.
* `status` - Peer admin status. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ha can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_ha.labelname SystemHa
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

