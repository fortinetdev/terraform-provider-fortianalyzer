// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Configure firmware management settings.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFwmSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFwmSettingUpdate,
		Read:   resourceFmupdateFwmSettingRead,
		Update: resourceFmupdateFwmSettingUpdate,
		Delete: resourceFmupdateFwmSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auto_scan_fgt_disk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_fgt_disk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_failover_fmg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_image_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"immx_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_device_history": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_profile_history": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"multiple_steps_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retrieve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retry_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retry_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"revision_diff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_image_retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upgrade_timeout": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"check_status_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ctrl_check_status_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ctrl_put_image_by_fds_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ha_sync_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"health_check_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"license_check_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prepare_image_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"put_image_by_fds_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"put_image_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"reboot_of_fsck_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"reboot_of_upgrade_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"retrieve_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rpc_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"total_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFmupdateFwmSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFwmSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFwmSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFwmSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFwmSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFwmSetting")

	return resourceFmupdateFwmSettingRead(d, m)
}

func resourceFmupdateFwmSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateFwmSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFwmSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFwmSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateFwmSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFwmSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFwmSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFwmSetting resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFwmSettingAutoScanFgtDiskFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingCheckFgtDiskFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingFdsFailoverFmgFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingFdsImageTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingHealthCheckFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingImmxSourceFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingLogFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingMaxDeviceHistoryFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingMaxProfileHistoryFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingMultipleStepsIntervalFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingRetrieveFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingRetryIntervalFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingRetryMaxFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingRevisionDiffFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingSendImageRetryFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "check_status_timeout"
	if _, ok := i["check-status-timeout"]; ok {
		result["check_status_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfb(i["check-status-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ctrl_check_status_timeout"
	if _, ok := i["ctrl-check-status-timeout"]; ok {
		result["ctrl_check_status_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfb(i["ctrl-check-status-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ctrl_put_image_by_fds_timeout"
	if _, ok := i["ctrl-put-image-by-fds-timeout"]; ok {
		result["ctrl_put_image_by_fds_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfb(i["ctrl-put-image-by-fds-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ha_sync_timeout"
	if _, ok := i["ha-sync-timeout"]; ok {
		result["ha_sync_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfb(i["ha-sync-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "health_check_timeout"
	if _, ok := i["health-check-timeout"]; ok {
		result["health_check_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutHealthCheckTimeoutFfb(i["health-check-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "license_check_timeout"
	if _, ok := i["license-check-timeout"]; ok {
		result["license_check_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfb(i["license-check-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "prepare_image_timeout"
	if _, ok := i["prepare-image-timeout"]; ok {
		result["prepare_image_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfb(i["prepare-image-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "put_image_by_fds_timeout"
	if _, ok := i["put-image-by-fds-timeout"]; ok {
		result["put_image_by_fds_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfb(i["put-image-by-fds-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "put_image_timeout"
	if _, ok := i["put-image-timeout"]; ok {
		result["put_image_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfb(i["put-image-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "reboot_of_fsck_timeout"
	if _, ok := i["reboot-of-fsck-timeout"]; ok {
		result["reboot_of_fsck_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfb(i["reboot-of-fsck-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "reboot_of_upgrade_timeout"
	if _, ok := i["reboot-of-upgrade-timeout"]; ok {
		result["reboot_of_upgrade_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfb(i["reboot-of-upgrade-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "retrieve_timeout"
	if _, ok := i["retrieve-timeout"]; ok {
		result["retrieve_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfb(i["retrieve-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "rpc_timeout"
	if _, ok := i["rpc-timeout"]; ok {
		result["rpc_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfb(i["rpc-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "total_timeout"
	if _, ok := i["total-timeout"]; ok {
		result["total_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfb(i["total-timeout"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutHealthCheckTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFwmSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_scan_fgt_disk", flattenFmupdateFwmSettingAutoScanFgtDiskFfb(o["auto-scan-fgt-disk"], d, "auto_scan_fgt_disk")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-scan-fgt-disk"], "FmupdateFwmSetting-AutoScanFgtDisk"); ok {
			if err = d.Set("auto_scan_fgt_disk", vv); err != nil {
				return fmt.Errorf("Error reading auto_scan_fgt_disk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_scan_fgt_disk: %v", err)
		}
	}

	if err = d.Set("check_fgt_disk", flattenFmupdateFwmSettingCheckFgtDiskFfb(o["check-fgt-disk"], d, "check_fgt_disk")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-fgt-disk"], "FmupdateFwmSetting-CheckFgtDisk"); ok {
			if err = d.Set("check_fgt_disk", vv); err != nil {
				return fmt.Errorf("Error reading check_fgt_disk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_fgt_disk: %v", err)
		}
	}

	if err = d.Set("fds_failover_fmg", flattenFmupdateFwmSettingFdsFailoverFmgFfb(o["fds-failover-fmg"], d, "fds_failover_fmg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-failover-fmg"], "FmupdateFwmSetting-FdsFailoverFmg"); ok {
			if err = d.Set("fds_failover_fmg", vv); err != nil {
				return fmt.Errorf("Error reading fds_failover_fmg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_failover_fmg: %v", err)
		}
	}

	if err = d.Set("fds_image_timeout", flattenFmupdateFwmSettingFdsImageTimeoutFfb(o["fds-image-timeout"], d, "fds_image_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-image-timeout"], "FmupdateFwmSetting-FdsImageTimeout"); ok {
			if err = d.Set("fds_image_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fds_image_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_image_timeout: %v", err)
		}
	}

	if err = d.Set("health_check", flattenFmupdateFwmSettingHealthCheckFfb(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "FmupdateFwmSetting-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("immx_source", flattenFmupdateFwmSettingImmxSourceFfb(o["immx-source"], d, "immx_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["immx-source"], "FmupdateFwmSetting-ImmxSource"); ok {
			if err = d.Set("immx_source", vv); err != nil {
				return fmt.Errorf("Error reading immx_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading immx_source: %v", err)
		}
	}

	if err = d.Set("log", flattenFmupdateFwmSettingLogFfb(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "FmupdateFwmSetting-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("max_device_history", flattenFmupdateFwmSettingMaxDeviceHistoryFfb(o["max-device-history"], d, "max_device_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-device-history"], "FmupdateFwmSetting-MaxDeviceHistory"); ok {
			if err = d.Set("max_device_history", vv); err != nil {
				return fmt.Errorf("Error reading max_device_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_device_history: %v", err)
		}
	}

	if err = d.Set("max_profile_history", flattenFmupdateFwmSettingMaxProfileHistoryFfb(o["max-profile-history"], d, "max_profile_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-profile-history"], "FmupdateFwmSetting-MaxProfileHistory"); ok {
			if err = d.Set("max_profile_history", vv); err != nil {
				return fmt.Errorf("Error reading max_profile_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_profile_history: %v", err)
		}
	}

	if err = d.Set("multiple_steps_interval", flattenFmupdateFwmSettingMultipleStepsIntervalFfb(o["multiple-steps-interval"], d, "multiple_steps_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-steps-interval"], "FmupdateFwmSetting-MultipleStepsInterval"); ok {
			if err = d.Set("multiple_steps_interval", vv); err != nil {
				return fmt.Errorf("Error reading multiple_steps_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_steps_interval: %v", err)
		}
	}

	if err = d.Set("retrieve", flattenFmupdateFwmSettingRetrieveFfb(o["retrieve"], d, "retrieve")); err != nil {
		if vv, ok := fortiAPIPatch(o["retrieve"], "FmupdateFwmSetting-Retrieve"); ok {
			if err = d.Set("retrieve", vv); err != nil {
				return fmt.Errorf("Error reading retrieve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retrieve: %v", err)
		}
	}

	if err = d.Set("retry_interval", flattenFmupdateFwmSettingRetryIntervalFfb(o["retry-interval"], d, "retry_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry-interval"], "FmupdateFwmSetting-RetryInterval"); ok {
			if err = d.Set("retry_interval", vv); err != nil {
				return fmt.Errorf("Error reading retry_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry_interval: %v", err)
		}
	}

	if err = d.Set("retry_max", flattenFmupdateFwmSettingRetryMaxFfb(o["retry-max"], d, "retry_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry-max"], "FmupdateFwmSetting-RetryMax"); ok {
			if err = d.Set("retry_max", vv); err != nil {
				return fmt.Errorf("Error reading retry_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry_max: %v", err)
		}
	}

	if err = d.Set("revision_diff", flattenFmupdateFwmSettingRevisionDiffFfb(o["revision-diff"], d, "revision_diff")); err != nil {
		if vv, ok := fortiAPIPatch(o["revision-diff"], "FmupdateFwmSetting-RevisionDiff"); ok {
			if err = d.Set("revision_diff", vv); err != nil {
				return fmt.Errorf("Error reading revision_diff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revision_diff: %v", err)
		}
	}

	if err = d.Set("send_image_retry", flattenFmupdateFwmSettingSendImageRetryFfb(o["send-image-retry"], d, "send_image_retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-image-retry"], "FmupdateFwmSetting-SendImageRetry"); ok {
			if err = d.Set("send_image_retry", vv); err != nil {
				return fmt.Errorf("Error reading send_image_retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_image_retry: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("upgrade_timeout", flattenFmupdateFwmSettingUpgradeTimeoutFfb(o["upgrade-timeout"], d, "upgrade_timeout")); err != nil {
			if vv, ok := fortiAPIPatch(o["upgrade-timeout"], "FmupdateFwmSetting-UpgradeTimeout"); ok {
				if err = d.Set("upgrade_timeout", vv); err != nil {
					return fmt.Errorf("Error reading upgrade_timeout: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading upgrade_timeout: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("upgrade_timeout"); ok {
			if err = d.Set("upgrade_timeout", flattenFmupdateFwmSettingUpgradeTimeoutFfb(o["upgrade-timeout"], d, "upgrade_timeout")); err != nil {
				if vv, ok := fortiAPIPatch(o["upgrade-timeout"], "FmupdateFwmSetting-UpgradeTimeout"); ok {
					if err = d.Set("upgrade_timeout", vv); err != nil {
						return fmt.Errorf("Error reading upgrade_timeout: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading upgrade_timeout: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFmupdateFwmSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFwmSettingAutoScanFgtDiskFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingCheckFgtDiskFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingFdsFailoverFmgFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingFdsImageTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingHealthCheckFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingImmxSourceFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingLogFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingMaxDeviceHistoryFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingMaxProfileHistoryFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingMultipleStepsIntervalFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingRetrieveFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingRetryIntervalFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingRetryMaxFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingRevisionDiffFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingSendImageRetryFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "check_status_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-status-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfb(d, i["check_status_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ctrl_check_status_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ctrl-check-status-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfb(d, i["ctrl_check_status_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ctrl_put_image_by_fds_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ctrl-put-image-by-fds-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfb(d, i["ctrl_put_image_by_fds_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ha_sync_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ha-sync-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfb(d, i["ha_sync_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "health_check_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["health-check-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutHealthCheckTimeoutFfb(d, i["health_check_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "license_check_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["license-check-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfb(d, i["license_check_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "prepare_image_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["prepare-image-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfb(d, i["prepare_image_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "put_image_by_fds_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["put-image-by-fds-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfb(d, i["put_image_by_fds_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "put_image_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["put-image-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfb(d, i["put_image_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "reboot_of_fsck_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reboot-of-fsck-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfb(d, i["reboot_of_fsck_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "reboot_of_upgrade_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reboot-of-upgrade-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfb(d, i["reboot_of_upgrade_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "retrieve_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retrieve-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfb(d, i["retrieve_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "rpc_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rpc-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfb(d, i["rpc_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "total_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["total-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfb(d, i["total_timeout"], pre_append)
	}

	return result, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutHealthCheckTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFwmSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_scan_fgt_disk"); ok || d.HasChange("auto_scan_fgt_disk") {
		t, err := expandFmupdateFwmSettingAutoScanFgtDiskFfb(d, v, "auto_scan_fgt_disk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-scan-fgt-disk"] = t
		}
	}

	if v, ok := d.GetOk("check_fgt_disk"); ok || d.HasChange("check_fgt_disk") {
		t, err := expandFmupdateFwmSettingCheckFgtDiskFfb(d, v, "check_fgt_disk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-fgt-disk"] = t
		}
	}

	if v, ok := d.GetOk("fds_failover_fmg"); ok || d.HasChange("fds_failover_fmg") {
		t, err := expandFmupdateFwmSettingFdsFailoverFmgFfb(d, v, "fds_failover_fmg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-failover-fmg"] = t
		}
	}

	if v, ok := d.GetOk("fds_image_timeout"); ok || d.HasChange("fds_image_timeout") {
		t, err := expandFmupdateFwmSettingFdsImageTimeoutFfb(d, v, "fds_image_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-image-timeout"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandFmupdateFwmSettingHealthCheckFfb(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("immx_source"); ok || d.HasChange("immx_source") {
		t, err := expandFmupdateFwmSettingImmxSourceFfb(d, v, "immx_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["immx-source"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandFmupdateFwmSettingLogFfb(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("max_device_history"); ok || d.HasChange("max_device_history") {
		t, err := expandFmupdateFwmSettingMaxDeviceHistoryFfb(d, v, "max_device_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-device-history"] = t
		}
	}

	if v, ok := d.GetOk("max_profile_history"); ok || d.HasChange("max_profile_history") {
		t, err := expandFmupdateFwmSettingMaxProfileHistoryFfb(d, v, "max_profile_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-profile-history"] = t
		}
	}

	if v, ok := d.GetOk("multiple_steps_interval"); ok || d.HasChange("multiple_steps_interval") {
		t, err := expandFmupdateFwmSettingMultipleStepsIntervalFfb(d, v, "multiple_steps_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-steps-interval"] = t
		}
	}

	if v, ok := d.GetOk("retrieve"); ok || d.HasChange("retrieve") {
		t, err := expandFmupdateFwmSettingRetrieveFfb(d, v, "retrieve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retrieve"] = t
		}
	}

	if v, ok := d.GetOk("retry_interval"); ok || d.HasChange("retry_interval") {
		t, err := expandFmupdateFwmSettingRetryIntervalFfb(d, v, "retry_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry-interval"] = t
		}
	}

	if v, ok := d.GetOk("retry_max"); ok || d.HasChange("retry_max") {
		t, err := expandFmupdateFwmSettingRetryMaxFfb(d, v, "retry_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry-max"] = t
		}
	}

	if v, ok := d.GetOk("revision_diff"); ok || d.HasChange("revision_diff") {
		t, err := expandFmupdateFwmSettingRevisionDiffFfb(d, v, "revision_diff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-diff"] = t
		}
	}

	if v, ok := d.GetOk("send_image_retry"); ok || d.HasChange("send_image_retry") {
		t, err := expandFmupdateFwmSettingSendImageRetryFfb(d, v, "send_image_retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-image-retry"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_timeout"); ok || d.HasChange("upgrade_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutFfb(d, v, "upgrade_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade-timeout"] = t
		}
	}

	return &obj, nil
}
