---
subcategory: "Device Manager"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_dvmdb_adom"
description: |-
  ADOM table, most attributes are read-only and can only be changed internally.
---

# fortianalyzer_dvmdb_adom
ADOM table, most attributes are read-only and can only be changed internally.

## Argument Reference


The following arguments are supported:


* `create_time` - Create_Time.
* `desc` - Desc.
* `flags` - Flags. Valid values: `migration`, `db_export`, `no_vpn_console`, `backup`, `other_devices`, `is_autosync`, `per_device_wtp`, `policy_check_on_install`, `install_on_policy_check_fail`, `auto_push_cfg`, `per_device_fsw`.

* `lock_override` - Lock_Override.
* `log_db_retention_hours` - Log_Db_Retention_Hours.
* `log_disk_quota` - Log_Disk_Quota.
* `log_disk_quota_alert_thres` - Log_Disk_Quota_Alert_Thres.
* `log_disk_quota_split_ratio` - Log_Disk_Quota_Split_Ratio.
* `log_file_retention_hours` - Log_File_Retention_Hours.
* `metafields` - Default metafields: none.
* `mig_mr` - Mig_Mr.
* `mig_os_ver` - Mig_Os_Ver. Valid values: `unknown`, `0.0`, `1.0`, `2.0`, `3.0`, `4.0`, `5.0`, `6.0`, `7.0`, `8.0`.

* `mode` - ems - (Value no longer used as of 4.3) provider - Global database. Valid values: `ems`, `gms`, `provider`.

* `mr` - Mr.
* `name` - Name.
* `os_ver` - Os_Ver. Valid values: `unknown`, `0.0`, `1.0`, `2.0`, `3.0`, `4.0`, `5.0`, `6.0`, `7.0`, `8.0`.

* `primary_dns_ip4` - Primary_Dns_Ip4.
* `primary_dns_ip6_1` - Primary_Dns_Ip6_1.
* `primary_dns_ip6_2` - Primary_Dns_Ip6_2.
* `primary_dns_ip6_3` - Primary_Dns_Ip6_3.
* `primary_dns_ip6_4` - Primary_Dns_Ip6_4.
* `restricted_prds_unitary` - Restricted_Prds. Valid values: `fos`, `foc`, `fml`, `fch`, `fwb`, `log`, `fct`, `faz`, `fsa`, `fsw`, `fmg`, `fdd`, `fac`, `fpx`, `fna`, `ffw`, `fsr`, `fad`, `fdc`, `fap`, `fxt`, `fts`, `fai`, `fwc`, `fis`, `fed`, `fabric`.

* `restricted_prds` - Restricted_Prds. Valid values: `fos`, `foc`, `fml`, `fch`, `fwb`, `log`, `fct`, `faz`, `fsa`, `fsw`, `fmg`, `fdd`, `fac`, `fpx`, `fna`, `ffw`, `fsr`, `fad`, `fdc`, `fap`, `fxt`.

* `secondary_dns_ip4` - Secondary_Dns_Ip4.
* `secondary_dns_ip6_1` - Secondary_Dns_Ip6_1.
* `secondary_dns_ip6_2` - Secondary_Dns_Ip6_2.
* `secondary_dns_ip6_3` - Secondary_Dns_Ip6_3.
* `secondary_dns_ip6_4` - Secondary_Dns_Ip6_4.
* `state` - State.
* `tz` - Tz.
* `uuid` - Uuid.
* `workspace_mode` - Workspace_Mode.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dvmdb Adom can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_dvmdb_adom.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

