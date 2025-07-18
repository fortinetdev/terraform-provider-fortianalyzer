---
subcategory: "Fmupdate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_fwmsetting"
description: |-
  Configure firmware management settings.
---

# fortianalyzer_fmupdate_fwmsetting
Configure firmware management settings.

## Example Usage

```hcl
resource "fortianalyzer_fmupdate_fwmsetting" "trname" {
  auto_scan_fgt_disk      = "enable"
  check_fgt_disk          = "enable"
  fds_failover_fmg        = "enable"
  fds_image_timeout       = 1800
  immx_source             = "fmg"
  multiple_steps_interval = 60
}
```

## Argument Reference


The following arguments are supported:


* `auto_scan_fgt_disk` - auto scan fgt disk if needed. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `check_fgt_disk` - check fgt disk before upgrade image. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fds_failover_fmg` - using fmg local image file is download from fds fails. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fds_image_timeout` - timer for fgt download image from fortiguard (300-3600s default=1800)
* `health_check` - do health check after upgrade disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `immx_source` - Configure which of IMMX file to be used for choosing upgrade path. Default is file for FortiManager fmg - Use IMMX file for FortiManager fgt - Use IMMX file for FortiGate cloud - Use IMMX file for FortiCloud Valid values: `fmg`, `fgt`, `cloud`.

* `log` - Configure log setting for fwm daemon fwm - FWM daemon log fwm_dm - FWM and Deployment service log fwm_dm_json - FWM and Deployment service log with JSON data between FMG-FGT Valid values: `fwm`, `fwm_dm`, `fwm_dm_json`.

* `max_device_history` - max number of device upgrade report(1-10000, default=100)
* `max_profile_history` - max number of profile upgrade report(1-10000, default=100)
* `multiple_steps_interval` - waiting time between multiple steps upgrade (30-180s, default=60)
* `retrieve` - do retrieve after upgrade disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `retry_interval` - waiting time for resending request to device(1-360s, default=60)
* `retry_max` - max retry times(0-100, default=10)
* `revision_diff` - calculate diff script after upgrade disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `send_image_retry` - retry send image when failed(0-2, default=0(no retry))
* `upgrade_timeout` - Upgrade-Timeout. The structure of `upgrade_timeout` block is documented below.

The `upgrade_timeout` block supports:

* `check_status_timeout` - timeout for checking status after tunnnel is up.(1-6000s, default=600)
* `ctrl_check_status_timeout` - timeout for checking fap/fsw/fext status after request upgrade.(1-12000s, default=1200)
* `ctrl_put_image_by_fds_timeout` - timeout for waiting device get fap/fsw/fext image from fortiguard.(1-9000ss, default=900)
* `ha_sync_timeout` - timeout for waiting HA sync.(1-18000s, default=1800)
* `health_check_timeout` - timeout for waiting retrieve.(1-6000s, default=600)
* `license_check_timeout` - timeout for waiting fortigate check license.(1-6000s, default=600)
* `prepare_image_timeout` - timeout for preparing image.(1-6000s, default=600)
* `put_image_by_fds_timeout` - timeout for waiting device get image from fortiguard.(1-18000s, default=1800)
* `put_image_timeout` - timeout for waiting send image over tunnel.(1-18000s, default=1800)
* `reboot_of_fsck_timeout` - timeout for waiting fortigate reboot.(1-18000s, default=1800)
* `reboot_of_upgrade_timeout` - timeout for waiting fortigate reboot after image upgrade.(1-12000s, default=1200)
* `retrieve_timeout` - timeout for waiting retrieve.(1-18000s, default=1800)
* `rpc_timeout` - timeout for waiting fortigate rpc response.(1-1800s, default=180)
* `total_timeout` - timeout for the whole fortigate upgrade(1-86400s, default=3600)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FwmSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_fwmsetting.labelname FmupdateFwmSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

