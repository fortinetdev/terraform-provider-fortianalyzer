---
subcategory: "System Admin"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_admin_setting"
description: |-
  Admin setting.
---

# fortianalyzer_system_admin_setting
Admin setting.

## Example Usage

```hcl
resource "fortianalyzer_system_admin_setting" "trname" {
  access_banner             = "disable"
  admin_https_redirect      = "enable"
  admin_login_max           = 256
  gui_theme                 = "blue"
  http_port                 = 80
  https_port                = 443
  idle_timeout              = 900
  idle_timeout_api          = 900
  idle_timeout_gui          = 900
  objects_force_deletion    = "enable"
  shell_access              = "disable"
  show_add_multiple         = "disable"
  show_checkbox_in_table    = "disable"
  show_device_import_export = "disable"
  show_fct_manager          = "disable"
  show_hostname             = "disable"
  show_log_forwarding       = "enable"
  unreg_dev_opt             = "add_allow_service"
  webadmin_language         = "auto_detect"
}
```

## Argument Reference


The following arguments are supported:


* `access_banner` - Enable/disable access banner. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `admin_https_redirect` - Enable/disable redirection of HTTP admin traffic to HTTPS. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `admin_login_max` - Maximum number admin users logged in at one time (1 - 256).
* `admin_scp` - Enable/disable admin SCP. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `admin_server_cert` - HTTPS & Web Service server certificate.
* `auth_addr` - IP which is used by FGT to authorize FAZ.
* `auth_port` - Port which is used by FGT to authorize FAZ.
* `banner_message` - Banner message.
* `fgt_gui_proxy` - Enable/disable FortiGate GUI proxy. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fgt_gui_proxy_port` - FortiGate GUI proxy port.
* `firmware_upgrade_check` - Enable/disable firmware upgrade check. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fsw_ignore_platform_check` - Enable/disable FortiSwitch Manager switch platform support check. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `gui_theme` - Color scheme to use for the administration GUI. blue - Blueberry green - Kiwi red - Cherry melongene - Plum spring - Spring summer - Summer autumn - Autumn winter - Winter circuit-board - Circuit Board calla-lily - Calla Lily binary-tunnel - Binary Tunnel mars - Mars blue-sea - Blue Sea technology - Technology landscape - Landscape twilight - Twilight canyon - Canyon northern-light - Northern Light astronomy - Astronomy fish - Fish penguin - Penguin mountain - Mountain panda - Panda parrot - Parrot cave - Cave zebra - Zebra contrast-dark - High Contrast Dark Valid values: `blue`, `green`, `red`, `melongene`, `spring`, `summer`, `autumn`, `winter`, `circuit-board`, `calla-lily`, `binary-tunnel`, `mars`, `blue-sea`, `technology`, `landscape`, `twilight`, `canyon`, `northern-light`, `astronomy`, `fish`, `penguin`, `mountain`, `panda`, `parrot`, `cave`, `zebra`, `contrast-dark`.

* `http_port` - HTTP port.
* `https_port` - HTTPS port.
* `idle_timeout` - Idle timeout (60 - 28800 sec).
* `idle_timeout_api` - Idle timeout for API sessions (1 - 28800 sec).
* `idle_timeout_gui` - Idle timeout for GUI sessions (60 - 28800 sec).
* `idle_timeout_sso` - Idle timeout for SSO sessions (60 - 28800 sec).
* `object_threshold_limit` - Object-Threshold-Limit. Valid values: `disable`, `enable`.

* `object_threshold_limit_value` - Object-Threshold-Limit-Value.
* `objects_force_deletion` - Enable/disable used objects force deletion. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `preferred_fgfm_intf` - Preferred interface for FGFM connection.
* `shell_access` - Enable/disable shell access. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `shell_password` - Password for shell access.
* `show_add_multiple` - Show add multiple button. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_checkbox_in_table` - Show checkboxs in tables on GUI. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_device_import_export` - Enable/disable import/export of ADOM, device, and group lists. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_fct_manager` - Enable/disable FCT manager. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_hostname` - Enable/disable hostname display in the GUI login page. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_log_forwarding` - Show log forwarding tab in regular mode. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_sdwan_manager` - Show-Sdwan-Manager. Valid values: `disable`, `enable`.

* `unreg_dev_opt` - Action to take when unregistered device connects to FortiAnalyzer. add_no_service - Add unregistered devices but deny service requests. ignore - Ignore unregistered devices. add_allow_service - Add unregistered devices and allow service requests. Valid values: `add_no_service`, `ignore`, `add_allow_service`.

* `webadmin_language` - Web admin language. auto_detect - Automatically detect language. english - English. simplified_chinese - Simplified Chinese. traditional_chinese - Traditional Chinese. japanese - Japanese. korean - Korean. spanish - Spanish. Valid values: `auto_detect`, `english`, `simplified_chinese`, `traditional_chinese`, `japanese`, `korean`, `spanish`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AdminSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_admin_setting.labelname SystemAdminSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

