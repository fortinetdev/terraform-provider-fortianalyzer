---
subcategory: "System Admin"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_admin_profile"
description: |-
  Admin profile.
---

# fortianalyzer_system_admin_profile
Admin profile.

## Example Usage

```hcl
resource "fortianalyzer_system_admin_profile" "trname" {
  profileid          = "profile"
  realtime_monitor   = "read-write"
  report_viewer      = "read-write"
  run_report         = "read-write"
  scope              = "global"
  script_access      = "read-write"
  super_user_profile = "enable"
  system_setting     = "read-write"
  triage_events      = "read-write"
  update_incidents   = "read-write"
}
```

## Argument Reference


The following arguments are supported:


* `adom_lock` - ADOM locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `adom_switch` - Administrator domain. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `allow_to_install` - Enable/disable the restricted user to install objects to the devices. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `change_password` - Enable/disable the user to change self password. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `datamask` - Enable/disable data masking. disable - Disable data masking. enable - Enable data masking. Valid values: `disable`, `enable`.

* `datamask_custom_fields` - Datamask-Custom-Fields. The structure of `datamask_custom_fields` block is documented below.
* `datamask_custom_priority` - Prioritize custom fields. disable - Disable custom field search priority. enable - Enable custom field search priority. Valid values: `disable`, `enable`.

* `datamask_fields` - Data masking fields. user - User name. srcip - Source IP. srcname - Source name. srcmac - Source MAC. dstip - Destination IP. dstname - Dst name. email - Email. message - Message. domain - Domain. Valid values: `user`, `srcip`, `srcname`, `srcmac`, `dstip`, `dstname`, `email`, `message`, `domain`.

* `datamask_key` - Data masking encryption key.
* `datamask_unmasked_time` - Time in days without data masking.
* `description` - Description.
* `device_ap` - Manage AP. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_forticlient` - Manage FortiClient. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_fortiextender` - Manage FortiExtender. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_fortiswitch` - Manage FortiSwitch. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_manager` - Device manager. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_op` - Device add/delete/edit. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_policy_package_lock` - Device/Policy Package locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_wan_link_load_balance` - Manage WAN link load balance. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `event_management` - Event management. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `execute_playbook` - Execute playbook. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `extension_access` - Manage extension access. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fabric_viewer` - Fabric viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fortirecorder_setting` - FortiRecorder settings. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `ips_lock` - IPS locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `ipv6_trusthost1` - Admin user trusted host IPv6, default ::/0 for all.
* `ipv6_trusthost10` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost2` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost3` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost4` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost5` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost6` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost7` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost8` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost9` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ips_baseline_ovrd` - Enable/disable override baseline ips sensor. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `log_viewer` - Log viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `profileid` - Profile ID.
* `realtime_monitor` - Realtime monitor. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `report_viewer` - Report viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `rpc_permit` - Set none/read/read-write rpc-permission read-write - Read-write permission. none - No permission. read - Read-only permission. Valid values: `read-write`, `none`, `read`.

* `run_report` - Run reports. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `scope` - Scope. global - Global scope. adom - ADOM scope. Valid values: `global`, `adom`.

* `script_access` - Script access. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `super_user_profile` - Enable/disable super user profile disable - Disable super user profile enable - Enable super user profile Valid values: `disable`, `enable`.

* `system_setting` - System setting. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `triage_events` - Triage events. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `trusthost1` - Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
* `trusthost10` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost2` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost3` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost4` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost5` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost6` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost7` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost8` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost9` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `update_incidents` - Create/update incidents. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `datamask_custom_fields` block supports:

* `field_category` - Field categories. log - Log. fortiview - FortiView. alert - Event management. ueba - UEBA. all - All. Valid values: `log`, `fortiview`, `alert`, `ueba`, `all`.

* `field_name` - Field name.
* `field_status` - Field status. disable - Disable field. enable - Enable field. Valid values: `disable`, `enable`.

* `field_type` - Field type. string - String. ip - IP. mac - MAC address. email - Email address. unknown - Unknown. Valid values: `string`, `ip`, `mac`, `email`, `unknown`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profileid}}.

## Import

System AdminProfile can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_admin_profile.labelname {{profileid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

