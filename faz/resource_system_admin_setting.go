// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Admin setting.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminSettingUpdate,
		Read:   resourceSystemAdminSettingRead,
		Update: resourceSystemAdminSettingUpdate,
		Delete: resourceSystemAdminSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"access_banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_login_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"banner_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgt_gui_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgt_gui_proxy_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"firmware_upgrade_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsw_ignore_platform_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"idle_timeout_api": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"idle_timeout_gui": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"idle_timeout_sso": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"object_threshold_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_threshold_limit_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"objects_force_deletion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_fgfm_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shell_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shell_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
			},
			"show_add_multiple": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"show_checkbox_in_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"show_device_import_export": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"show_fct_manager": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"show_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"show_log_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unreg_dev_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webadmin_language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAdminSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAdminSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdminSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAdminSetting")

	return resourceSystemAdminSettingRead(d, m)
}

func resourceSystemAdminSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAdminSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAdminSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminSettingAccessBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingAdminHttpsRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingAdminLoginMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingAdminServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingAuthAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingAuthPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingBannerMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingFgtGuiProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingFgtGuiProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingFirmwareUpgradeCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingFswIgnorePlatformCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingGuiTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingIdleTimeoutApi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingIdleTimeoutGui(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingIdleTimeoutSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingObjectThresholdLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingObjectThresholdLimitValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingObjectsForceDeletion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingPreferredFgfmIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingShellAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingShellPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminSettingShowAddMultiple(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingShowCheckboxInTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingShowDeviceImportExport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingShowFctManager(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingShowHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingShowLogForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingUnregDevOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSettingWebadminLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("access_banner", flattenSystemAdminSettingAccessBanner(o["access-banner"], d, "access_banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-banner"], "SystemAdminSetting-AccessBanner"); ok {
			if err = d.Set("access_banner", vv); err != nil {
				return fmt.Errorf("Error reading access_banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_banner: %v", err)
		}
	}

	if err = d.Set("admin_https_redirect", flattenSystemAdminSettingAdminHttpsRedirect(o["admin-https-redirect"], d, "admin_https_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-https-redirect"], "SystemAdminSetting-AdminHttpsRedirect"); ok {
			if err = d.Set("admin_https_redirect", vv); err != nil {
				return fmt.Errorf("Error reading admin_https_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_https_redirect: %v", err)
		}
	}

	if err = d.Set("admin_login_max", flattenSystemAdminSettingAdminLoginMax(o["admin-login-max"], d, "admin_login_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-login-max"], "SystemAdminSetting-AdminLoginMax"); ok {
			if err = d.Set("admin_login_max", vv); err != nil {
				return fmt.Errorf("Error reading admin_login_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_login_max: %v", err)
		}
	}

	if err = d.Set("admin_server_cert", flattenSystemAdminSettingAdminServerCert(o["admin_server_cert"], d, "admin_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin_server_cert"], "SystemAdminSetting-AdminServerCert"); ok {
			if err = d.Set("admin_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading admin_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_server_cert: %v", err)
		}
	}

	if err = d.Set("auth_addr", flattenSystemAdminSettingAuthAddr(o["auth-addr"], d, "auth_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-addr"], "SystemAdminSetting-AuthAddr"); ok {
			if err = d.Set("auth_addr", vv); err != nil {
				return fmt.Errorf("Error reading auth_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_addr: %v", err)
		}
	}

	if err = d.Set("auth_port", flattenSystemAdminSettingAuthPort(o["auth-port"], d, "auth_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-port"], "SystemAdminSetting-AuthPort"); ok {
			if err = d.Set("auth_port", vv); err != nil {
				return fmt.Errorf("Error reading auth_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_port: %v", err)
		}
	}

	if err = d.Set("banner_message", flattenSystemAdminSettingBannerMessage(o["banner-message"], d, "banner_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["banner-message"], "SystemAdminSetting-BannerMessage"); ok {
			if err = d.Set("banner_message", vv); err != nil {
				return fmt.Errorf("Error reading banner_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banner_message: %v", err)
		}
	}

	if err = d.Set("fgt_gui_proxy", flattenSystemAdminSettingFgtGuiProxy(o["fgt-gui-proxy"], d, "fgt_gui_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-gui-proxy"], "SystemAdminSetting-FgtGuiProxy"); ok {
			if err = d.Set("fgt_gui_proxy", vv); err != nil {
				return fmt.Errorf("Error reading fgt_gui_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_gui_proxy: %v", err)
		}
	}

	if err = d.Set("fgt_gui_proxy_port", flattenSystemAdminSettingFgtGuiProxyPort(o["fgt-gui-proxy-port"], d, "fgt_gui_proxy_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-gui-proxy-port"], "SystemAdminSetting-FgtGuiProxyPort"); ok {
			if err = d.Set("fgt_gui_proxy_port", vv); err != nil {
				return fmt.Errorf("Error reading fgt_gui_proxy_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_gui_proxy_port: %v", err)
		}
	}

	if err = d.Set("firmware_upgrade_check", flattenSystemAdminSettingFirmwareUpgradeCheck(o["firmware-upgrade-check"], d, "firmware_upgrade_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-upgrade-check"], "SystemAdminSetting-FirmwareUpgradeCheck"); ok {
			if err = d.Set("firmware_upgrade_check", vv); err != nil {
				return fmt.Errorf("Error reading firmware_upgrade_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_upgrade_check: %v", err)
		}
	}

	if err = d.Set("fsw_ignore_platform_check", flattenSystemAdminSettingFswIgnorePlatformCheck(o["fsw-ignore-platform-check"], d, "fsw_ignore_platform_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsw-ignore-platform-check"], "SystemAdminSetting-FswIgnorePlatformCheck"); ok {
			if err = d.Set("fsw_ignore_platform_check", vv); err != nil {
				return fmt.Errorf("Error reading fsw_ignore_platform_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsw_ignore_platform_check: %v", err)
		}
	}

	if err = d.Set("gui_theme", flattenSystemAdminSettingGuiTheme(o["gui-theme"], d, "gui_theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-theme"], "SystemAdminSetting-GuiTheme"); ok {
			if err = d.Set("gui_theme", vv); err != nil {
				return fmt.Errorf("Error reading gui_theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_theme: %v", err)
		}
	}

	if err = d.Set("http_port", flattenSystemAdminSettingHttpPort(o["http_port"], d, "http_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["http_port"], "SystemAdminSetting-HttpPort"); ok {
			if err = d.Set("http_port", vv); err != nil {
				return fmt.Errorf("Error reading http_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_port: %v", err)
		}
	}

	if err = d.Set("https_port", flattenSystemAdminSettingHttpsPort(o["https_port"], d, "https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["https_port"], "SystemAdminSetting-HttpsPort"); ok {
			if err = d.Set("https_port", vv); err != nil {
				return fmt.Errorf("Error reading https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenSystemAdminSettingIdleTimeout(o["idle_timeout"], d, "idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle_timeout"], "SystemAdminSetting-IdleTimeout"); ok {
			if err = d.Set("idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("idle_timeout_api", flattenSystemAdminSettingIdleTimeoutApi(o["idle_timeout_api"], d, "idle_timeout_api")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle_timeout_api"], "SystemAdminSetting-IdleTimeoutApi"); ok {
			if err = d.Set("idle_timeout_api", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout_api: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout_api: %v", err)
		}
	}

	if err = d.Set("idle_timeout_gui", flattenSystemAdminSettingIdleTimeoutGui(o["idle_timeout_gui"], d, "idle_timeout_gui")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle_timeout_gui"], "SystemAdminSetting-IdleTimeoutGui"); ok {
			if err = d.Set("idle_timeout_gui", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout_gui: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout_gui: %v", err)
		}
	}

	if err = d.Set("idle_timeout_sso", flattenSystemAdminSettingIdleTimeoutSso(o["idle_timeout_sso"], d, "idle_timeout_sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle_timeout_sso"], "SystemAdminSetting-IdleTimeoutSso"); ok {
			if err = d.Set("idle_timeout_sso", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout_sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout_sso: %v", err)
		}
	}

	if err = d.Set("object_threshold_limit", flattenSystemAdminSettingObjectThresholdLimit(o["object-threshold-limit"], d, "object_threshold_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-threshold-limit"], "SystemAdminSetting-ObjectThresholdLimit"); ok {
			if err = d.Set("object_threshold_limit", vv); err != nil {
				return fmt.Errorf("Error reading object_threshold_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_threshold_limit: %v", err)
		}
	}

	if err = d.Set("object_threshold_limit_value", flattenSystemAdminSettingObjectThresholdLimitValue(o["object-threshold-limit-value"], d, "object_threshold_limit_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-threshold-limit-value"], "SystemAdminSetting-ObjectThresholdLimitValue"); ok {
			if err = d.Set("object_threshold_limit_value", vv); err != nil {
				return fmt.Errorf("Error reading object_threshold_limit_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_threshold_limit_value: %v", err)
		}
	}

	if err = d.Set("objects_force_deletion", flattenSystemAdminSettingObjectsForceDeletion(o["objects-force-deletion"], d, "objects_force_deletion")); err != nil {
		if vv, ok := fortiAPIPatch(o["objects-force-deletion"], "SystemAdminSetting-ObjectsForceDeletion"); ok {
			if err = d.Set("objects_force_deletion", vv); err != nil {
				return fmt.Errorf("Error reading objects_force_deletion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading objects_force_deletion: %v", err)
		}
	}

	if err = d.Set("preferred_fgfm_intf", flattenSystemAdminSettingPreferredFgfmIntf(o["preferred-fgfm-intf"], d, "preferred_fgfm_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-fgfm-intf"], "SystemAdminSetting-PreferredFgfmIntf"); ok {
			if err = d.Set("preferred_fgfm_intf", vv); err != nil {
				return fmt.Errorf("Error reading preferred_fgfm_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_fgfm_intf: %v", err)
		}
	}

	if err = d.Set("shell_access", flattenSystemAdminSettingShellAccess(o["shell-access"], d, "shell_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["shell-access"], "SystemAdminSetting-ShellAccess"); ok {
			if err = d.Set("shell_access", vv); err != nil {
				return fmt.Errorf("Error reading shell_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shell_access: %v", err)
		}
	}

	if err = d.Set("show_add_multiple", flattenSystemAdminSettingShowAddMultiple(o["show-add-multiple"], d, "show_add_multiple")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-add-multiple"], "SystemAdminSetting-ShowAddMultiple"); ok {
			if err = d.Set("show_add_multiple", vv); err != nil {
				return fmt.Errorf("Error reading show_add_multiple: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_add_multiple: %v", err)
		}
	}

	if err = d.Set("show_checkbox_in_table", flattenSystemAdminSettingShowCheckboxInTable(o["show-checkbox-in-table"], d, "show_checkbox_in_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-checkbox-in-table"], "SystemAdminSetting-ShowCheckboxInTable"); ok {
			if err = d.Set("show_checkbox_in_table", vv); err != nil {
				return fmt.Errorf("Error reading show_checkbox_in_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_checkbox_in_table: %v", err)
		}
	}

	if err = d.Set("show_device_import_export", flattenSystemAdminSettingShowDeviceImportExport(o["show-device-import-export"], d, "show_device_import_export")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-device-import-export"], "SystemAdminSetting-ShowDeviceImportExport"); ok {
			if err = d.Set("show_device_import_export", vv); err != nil {
				return fmt.Errorf("Error reading show_device_import_export: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_device_import_export: %v", err)
		}
	}

	if err = d.Set("show_fct_manager", flattenSystemAdminSettingShowFctManager(o["show-fct-manager"], d, "show_fct_manager")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-fct-manager"], "SystemAdminSetting-ShowFctManager"); ok {
			if err = d.Set("show_fct_manager", vv); err != nil {
				return fmt.Errorf("Error reading show_fct_manager: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_fct_manager: %v", err)
		}
	}

	if err = d.Set("show_hostname", flattenSystemAdminSettingShowHostname(o["show-hostname"], d, "show_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-hostname"], "SystemAdminSetting-ShowHostname"); ok {
			if err = d.Set("show_hostname", vv); err != nil {
				return fmt.Errorf("Error reading show_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_hostname: %v", err)
		}
	}

	if err = d.Set("show_log_forwarding", flattenSystemAdminSettingShowLogForwarding(o["show-log-forwarding"], d, "show_log_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-log-forwarding"], "SystemAdminSetting-ShowLogForwarding"); ok {
			if err = d.Set("show_log_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading show_log_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_log_forwarding: %v", err)
		}
	}

	if err = d.Set("unreg_dev_opt", flattenSystemAdminSettingUnregDevOpt(o["unreg_dev_opt"], d, "unreg_dev_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["unreg_dev_opt"], "SystemAdminSetting-UnregDevOpt"); ok {
			if err = d.Set("unreg_dev_opt", vv); err != nil {
				return fmt.Errorf("Error reading unreg_dev_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unreg_dev_opt: %v", err)
		}
	}

	if err = d.Set("webadmin_language", flattenSystemAdminSettingWebadminLanguage(o["webadmin_language"], d, "webadmin_language")); err != nil {
		if vv, ok := fortiAPIPatch(o["webadmin_language"], "SystemAdminSetting-WebadminLanguage"); ok {
			if err = d.Set("webadmin_language", vv); err != nil {
				return fmt.Errorf("Error reading webadmin_language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webadmin_language: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminSettingAccessBanner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingAdminHttpsRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingAdminLoginMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingAdminServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingAuthAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingAuthPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingBannerMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingFgtGuiProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingFgtGuiProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingFirmwareUpgradeCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingFswIgnorePlatformCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingGuiTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingIdleTimeoutApi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingIdleTimeoutGui(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingIdleTimeoutSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingObjectThresholdLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingObjectThresholdLimitValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingObjectsForceDeletion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingPreferredFgfmIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingShellAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingShellPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminSettingShowAddMultiple(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingShowCheckboxInTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingShowDeviceImportExport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingShowFctManager(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingShowHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingShowLogForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingUnregDevOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSettingWebadminLanguage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_banner"); ok || d.HasChange("access_banner") {
		t, err := expandSystemAdminSettingAccessBanner(d, v, "access_banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-banner"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_redirect"); ok || d.HasChange("admin_https_redirect") {
		t, err := expandSystemAdminSettingAdminHttpsRedirect(d, v, "admin_https_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-redirect"] = t
		}
	}

	if v, ok := d.GetOk("admin_login_max"); ok || d.HasChange("admin_login_max") {
		t, err := expandSystemAdminSettingAdminLoginMax(d, v, "admin_login_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-login-max"] = t
		}
	}

	if v, ok := d.GetOk("admin_server_cert"); ok || d.HasChange("admin_server_cert") {
		t, err := expandSystemAdminSettingAdminServerCert(d, v, "admin_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin_server_cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_addr"); ok || d.HasChange("auth_addr") {
		t, err := expandSystemAdminSettingAuthAddr(d, v, "auth_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-addr"] = t
		}
	}

	if v, ok := d.GetOk("auth_port"); ok || d.HasChange("auth_port") {
		t, err := expandSystemAdminSettingAuthPort(d, v, "auth_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-port"] = t
		}
	}

	if v, ok := d.GetOk("banner_message"); ok || d.HasChange("banner_message") {
		t, err := expandSystemAdminSettingBannerMessage(d, v, "banner_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner-message"] = t
		}
	}

	if v, ok := d.GetOk("fgt_gui_proxy"); ok || d.HasChange("fgt_gui_proxy") {
		t, err := expandSystemAdminSettingFgtGuiProxy(d, v, "fgt_gui_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-gui-proxy"] = t
		}
	}

	if v, ok := d.GetOk("fgt_gui_proxy_port"); ok || d.HasChange("fgt_gui_proxy_port") {
		t, err := expandSystemAdminSettingFgtGuiProxyPort(d, v, "fgt_gui_proxy_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-gui-proxy-port"] = t
		}
	}

	if v, ok := d.GetOk("firmware_upgrade_check"); ok || d.HasChange("firmware_upgrade_check") {
		t, err := expandSystemAdminSettingFirmwareUpgradeCheck(d, v, "firmware_upgrade_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-upgrade-check"] = t
		}
	}

	if v, ok := d.GetOk("fsw_ignore_platform_check"); ok || d.HasChange("fsw_ignore_platform_check") {
		t, err := expandSystemAdminSettingFswIgnorePlatformCheck(d, v, "fsw_ignore_platform_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-ignore-platform-check"] = t
		}
	}

	if v, ok := d.GetOk("gui_theme"); ok || d.HasChange("gui_theme") {
		t, err := expandSystemAdminSettingGuiTheme(d, v, "gui_theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme"] = t
		}
	}

	if v, ok := d.GetOk("http_port"); ok || d.HasChange("http_port") {
		t, err := expandSystemAdminSettingHttpPort(d, v, "http_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http_port"] = t
		}
	}

	if v, ok := d.GetOk("https_port"); ok || d.HasChange("https_port") {
		t, err := expandSystemAdminSettingHttpsPort(d, v, "https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https_port"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok || d.HasChange("idle_timeout") {
		t, err := expandSystemAdminSettingIdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle_timeout"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout_api"); ok || d.HasChange("idle_timeout_api") {
		t, err := expandSystemAdminSettingIdleTimeoutApi(d, v, "idle_timeout_api")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle_timeout_api"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout_gui"); ok || d.HasChange("idle_timeout_gui") {
		t, err := expandSystemAdminSettingIdleTimeoutGui(d, v, "idle_timeout_gui")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle_timeout_gui"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout_sso"); ok || d.HasChange("idle_timeout_sso") {
		t, err := expandSystemAdminSettingIdleTimeoutSso(d, v, "idle_timeout_sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle_timeout_sso"] = t
		}
	}

	if v, ok := d.GetOk("object_threshold_limit"); ok || d.HasChange("object_threshold_limit") {
		t, err := expandSystemAdminSettingObjectThresholdLimit(d, v, "object_threshold_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-threshold-limit"] = t
		}
	}

	if v, ok := d.GetOk("object_threshold_limit_value"); ok || d.HasChange("object_threshold_limit_value") {
		t, err := expandSystemAdminSettingObjectThresholdLimitValue(d, v, "object_threshold_limit_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-threshold-limit-value"] = t
		}
	}

	if v, ok := d.GetOk("objects_force_deletion"); ok || d.HasChange("objects_force_deletion") {
		t, err := expandSystemAdminSettingObjectsForceDeletion(d, v, "objects_force_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["objects-force-deletion"] = t
		}
	}

	if v, ok := d.GetOk("preferred_fgfm_intf"); ok || d.HasChange("preferred_fgfm_intf") {
		t, err := expandSystemAdminSettingPreferredFgfmIntf(d, v, "preferred_fgfm_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-fgfm-intf"] = t
		}
	}

	if v, ok := d.GetOk("shell_access"); ok || d.HasChange("shell_access") {
		t, err := expandSystemAdminSettingShellAccess(d, v, "shell_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shell-access"] = t
		}
	}

	if v, ok := d.GetOk("shell_password"); ok || d.HasChange("shell_password") {
		t, err := expandSystemAdminSettingShellPassword(d, v, "shell_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shell-password"] = t
		}
	}

	if v, ok := d.GetOk("show_add_multiple"); ok || d.HasChange("show_add_multiple") {
		t, err := expandSystemAdminSettingShowAddMultiple(d, v, "show_add_multiple")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-add-multiple"] = t
		}
	}

	if v, ok := d.GetOk("show_checkbox_in_table"); ok || d.HasChange("show_checkbox_in_table") {
		t, err := expandSystemAdminSettingShowCheckboxInTable(d, v, "show_checkbox_in_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-checkbox-in-table"] = t
		}
	}

	if v, ok := d.GetOk("show_device_import_export"); ok || d.HasChange("show_device_import_export") {
		t, err := expandSystemAdminSettingShowDeviceImportExport(d, v, "show_device_import_export")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-device-import-export"] = t
		}
	}

	if v, ok := d.GetOk("show_fct_manager"); ok || d.HasChange("show_fct_manager") {
		t, err := expandSystemAdminSettingShowFctManager(d, v, "show_fct_manager")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-fct-manager"] = t
		}
	}

	if v, ok := d.GetOk("show_hostname"); ok || d.HasChange("show_hostname") {
		t, err := expandSystemAdminSettingShowHostname(d, v, "show_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-hostname"] = t
		}
	}

	if v, ok := d.GetOk("show_log_forwarding"); ok || d.HasChange("show_log_forwarding") {
		t, err := expandSystemAdminSettingShowLogForwarding(d, v, "show_log_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-log-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("unreg_dev_opt"); ok || d.HasChange("unreg_dev_opt") {
		t, err := expandSystemAdminSettingUnregDevOpt(d, v, "unreg_dev_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unreg_dev_opt"] = t
		}
	}

	if v, ok := d.GetOk("webadmin_language"); ok || d.HasChange("webadmin_language") {
		t, err := expandSystemAdminSettingWebadminLanguage(d, v, "webadmin_language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webadmin_language"] = t
		}
	}

	return &obj, nil
}
