---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_fgdsetting"
description: |-
  Fmupdate Fgd-Setting
---

# fortianalyzer_fmupdate_fgdsetting
Fmupdate Fgd-Setting

## Argument Reference


The following arguments are supported:


* `as_cache` - As-Cache.
* `as_log` - As-Log. Valid values: `disable`, `nospam`, `all`.

* `as_preload` - As-Preload. Valid values: `disable`, `enable`.

* `av_cache` - Av-Cache.
* `av_log` - Av-Log. Valid values: `disable`, `novirus`, `all`.

* `av_preload` - Av-Preload. Valid values: `disable`, `enable`.

* `av2_cache` - Av2-Cache.
* `av2_log` - Av2-Log. Valid values: `disable`, `noav2`, `all`.

* `av2_preload` - Av2-Preload. Valid values: `disable`, `enable`.

* `eventlog_query` - Eventlog-Query. Valid values: `disable`, `enable`.

* `fgd_pull_interval` - Fgd-Pull-Interval.
* `fq_cache` - Fq-Cache.
* `fq_log` - Fq-Log. Valid values: `disable`, `nofilequery`, `all`.

* `fq_preload` - Fq-Preload. Valid values: `disable`, `enable`.

* `iot_cache` - Iot-Cache.
* `iot_log` - Iot-Log. Valid values: `disable`, `noiot`, `all`.

* `iot_preload` - Iot-Preload. Valid values: `disable`, `enable`.

* `iotv_preload` - Iotv-Preload. Valid values: `disable`, `enable`.

* `linkd_log` - Linkd-Log. Valid values: `emergency`, `alert`, `critical`, `error`, `warn`, `notice`, `info`, `debug`, `disable`.

* `max_client_worker` - Max-Client-Worker.
* `max_log_quota` - Max-Log-Quota.
* `max_unrated_site` - Max-Unrated-Site.
* `restrict_as1_dbver` - Restrict-As1-Dbver.
* `restrict_as2_dbver` - Restrict-As2-Dbver.
* `restrict_as4_dbver` - Restrict-As4-Dbver.
* `restrict_av_dbver` - Restrict-Av-Dbver.
* `restrict_av2_dbver` - Restrict-Av2-Dbver.
* `restrict_fq_dbver` - Restrict-Fq-Dbver.
* `restrict_iots_dbver` - Restrict-Iots-Dbver.
* `restrict_wf_dbver` - Restrict-Wf-Dbver.
* `server_override` - Server-Override. The structure of `server_override` block is documented below.
* `stat_log` - Stat-Log. Valid values: `emergency`, `alert`, `critical`, `error`, `warn`, `notice`, `info`, `debug`, `disable`.

* `stat_log_interval` - Stat-Log-Interval.
* `stat_sync_interval` - Stat-Sync-Interval.
* `update_interval` - Update-Interval.
* `update_log` - Update-Log. Valid values: `disable`, `enable`.

* `wf_cache` - Wf-Cache.
* `wf_dn_cache_expire_time` - Wf-Dn-Cache-Expire-Time.
* `wf_dn_cache_max_number` - Wf-Dn-Cache-Max-Number.
* `wf_log` - Wf-Log. Valid values: `disable`, `nourl`, `all`.

* `wf_preload` - Wf-Preload. Valid values: `disable`, `enable`.


The `server_override` block supports:

* `servlist` - Servlist. The structure of `servlist` block is documented below.
* `status` - Status. Valid values: `disable`, `enable`.


The `servlist` block supports:

* `id` - Id.
* `ip` - Ip.
* `ip6` - Ip6.
* `port` - Port.
* `service_type` - Service-Type. Valid values: `fgd`, `fsa`, `fgfq`, `geoip`, `iot-collect`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FgdSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_fgdsetting.labelname FmupdateFgdSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

