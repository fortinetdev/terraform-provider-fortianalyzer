// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Log queue settings.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSocCollectorLogQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSocCollectorLogQueueUpdate,
		Read:   resourceSystemSocCollectorLogQueueRead,
		Update: resourceSystemSocCollectorLogQueueUpdate,
		Delete: resourceSystemSocCollectorLogQueueDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"config_queue_file_count_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"config_queue_size_limit_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"linux_queue_file_count_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"linux_queue_size_limit_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"logs_queue_file_count_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"logs_queue_size_limit_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"win_queue_file_count_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"win_queue_size_limit_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSocCollectorLogQueueUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSocCollectorLogQueue(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocCollectorLogQueue resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSocCollectorLogQueue(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocCollectorLogQueue resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSocCollectorLogQueue")

	return resourceSystemSocCollectorLogQueueRead(d, m)
}

func resourceSystemSocCollectorLogQueueDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSocCollectorLogQueue(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSocCollectorLogQueue resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSocCollectorLogQueueRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSocCollectorLogQueue(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocCollectorLogQueue resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSocCollectorLogQueue(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocCollectorLogQueue resource from API: %v", err)
	}
	return nil
}

func flattenSystemSocCollectorLogQueueConfigQueueFileCountLimitSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueConfigQueueSizeLimitMbSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLinuxQueueFileCountLimitSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLinuxQueueSizeLimitMbSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLogsQueueFileCountLimitSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLogsQueueSizeLimitMbSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueWinQueueFileCountLimitSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueWinQueueSizeLimitMbSsclq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSocCollectorLogQueue(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("config_queue_file_count_limit", flattenSystemSocCollectorLogQueueConfigQueueFileCountLimitSsclq(o["config-queue-file-count-limit"], d, "config_queue_file_count_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-queue-file-count-limit"], "SystemSocCollectorLogQueue-ConfigQueueFileCountLimit"); ok {
			if err = d.Set("config_queue_file_count_limit", vv); err != nil {
				return fmt.Errorf("Error reading config_queue_file_count_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_queue_file_count_limit: %v", err)
		}
	}

	if err = d.Set("config_queue_size_limit_mb", flattenSystemSocCollectorLogQueueConfigQueueSizeLimitMbSsclq(o["config-queue-size-limit-mb"], d, "config_queue_size_limit_mb")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-queue-size-limit-mb"], "SystemSocCollectorLogQueue-ConfigQueueSizeLimitMb"); ok {
			if err = d.Set("config_queue_size_limit_mb", vv); err != nil {
				return fmt.Errorf("Error reading config_queue_size_limit_mb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_queue_size_limit_mb: %v", err)
		}
	}

	if err = d.Set("linux_queue_file_count_limit", flattenSystemSocCollectorLogQueueLinuxQueueFileCountLimitSsclq(o["linux-queue-file-count-limit"], d, "linux_queue_file_count_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["linux-queue-file-count-limit"], "SystemSocCollectorLogQueue-LinuxQueueFileCountLimit"); ok {
			if err = d.Set("linux_queue_file_count_limit", vv); err != nil {
				return fmt.Errorf("Error reading linux_queue_file_count_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linux_queue_file_count_limit: %v", err)
		}
	}

	if err = d.Set("linux_queue_size_limit_mb", flattenSystemSocCollectorLogQueueLinuxQueueSizeLimitMbSsclq(o["linux-queue-size-limit-mb"], d, "linux_queue_size_limit_mb")); err != nil {
		if vv, ok := fortiAPIPatch(o["linux-queue-size-limit-mb"], "SystemSocCollectorLogQueue-LinuxQueueSizeLimitMb"); ok {
			if err = d.Set("linux_queue_size_limit_mb", vv); err != nil {
				return fmt.Errorf("Error reading linux_queue_size_limit_mb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linux_queue_size_limit_mb: %v", err)
		}
	}

	if err = d.Set("logs_queue_file_count_limit", flattenSystemSocCollectorLogQueueLogsQueueFileCountLimitSsclq(o["logs-queue-file-count-limit"], d, "logs_queue_file_count_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["logs-queue-file-count-limit"], "SystemSocCollectorLogQueue-LogsQueueFileCountLimit"); ok {
			if err = d.Set("logs_queue_file_count_limit", vv); err != nil {
				return fmt.Errorf("Error reading logs_queue_file_count_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logs_queue_file_count_limit: %v", err)
		}
	}

	if err = d.Set("logs_queue_size_limit_mb", flattenSystemSocCollectorLogQueueLogsQueueSizeLimitMbSsclq(o["logs-queue-size-limit-mb"], d, "logs_queue_size_limit_mb")); err != nil {
		if vv, ok := fortiAPIPatch(o["logs-queue-size-limit-mb"], "SystemSocCollectorLogQueue-LogsQueueSizeLimitMb"); ok {
			if err = d.Set("logs_queue_size_limit_mb", vv); err != nil {
				return fmt.Errorf("Error reading logs_queue_size_limit_mb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logs_queue_size_limit_mb: %v", err)
		}
	}

	if err = d.Set("win_queue_file_count_limit", flattenSystemSocCollectorLogQueueWinQueueFileCountLimitSsclq(o["win-queue-file-count-limit"], d, "win_queue_file_count_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["win-queue-file-count-limit"], "SystemSocCollectorLogQueue-WinQueueFileCountLimit"); ok {
			if err = d.Set("win_queue_file_count_limit", vv); err != nil {
				return fmt.Errorf("Error reading win_queue_file_count_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading win_queue_file_count_limit: %v", err)
		}
	}

	if err = d.Set("win_queue_size_limit_mb", flattenSystemSocCollectorLogQueueWinQueueSizeLimitMbSsclq(o["win-queue-size-limit-mb"], d, "win_queue_size_limit_mb")); err != nil {
		if vv, ok := fortiAPIPatch(o["win-queue-size-limit-mb"], "SystemSocCollectorLogQueue-WinQueueSizeLimitMb"); ok {
			if err = d.Set("win_queue_size_limit_mb", vv); err != nil {
				return fmt.Errorf("Error reading win_queue_size_limit_mb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading win_queue_size_limit_mb: %v", err)
		}
	}

	return nil
}

func flattenSystemSocCollectorLogQueueFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSocCollectorLogQueueConfigQueueFileCountLimitSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueConfigQueueSizeLimitMbSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLinuxQueueFileCountLimitSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLinuxQueueSizeLimitMbSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLogsQueueFileCountLimitSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLogsQueueSizeLimitMbSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueWinQueueFileCountLimitSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueWinQueueSizeLimitMbSsclq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSocCollectorLogQueue(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("config_queue_file_count_limit"); ok || d.HasChange("config_queue_file_count_limit") {
		t, err := expandSystemSocCollectorLogQueueConfigQueueFileCountLimitSsclq(d, v, "config_queue_file_count_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-queue-file-count-limit"] = t
		}
	}

	if v, ok := d.GetOk("config_queue_size_limit_mb"); ok || d.HasChange("config_queue_size_limit_mb") {
		t, err := expandSystemSocCollectorLogQueueConfigQueueSizeLimitMbSsclq(d, v, "config_queue_size_limit_mb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-queue-size-limit-mb"] = t
		}
	}

	if v, ok := d.GetOk("linux_queue_file_count_limit"); ok || d.HasChange("linux_queue_file_count_limit") {
		t, err := expandSystemSocCollectorLogQueueLinuxQueueFileCountLimitSsclq(d, v, "linux_queue_file_count_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linux-queue-file-count-limit"] = t
		}
	}

	if v, ok := d.GetOk("linux_queue_size_limit_mb"); ok || d.HasChange("linux_queue_size_limit_mb") {
		t, err := expandSystemSocCollectorLogQueueLinuxQueueSizeLimitMbSsclq(d, v, "linux_queue_size_limit_mb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linux-queue-size-limit-mb"] = t
		}
	}

	if v, ok := d.GetOk("logs_queue_file_count_limit"); ok || d.HasChange("logs_queue_file_count_limit") {
		t, err := expandSystemSocCollectorLogQueueLogsQueueFileCountLimitSsclq(d, v, "logs_queue_file_count_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logs-queue-file-count-limit"] = t
		}
	}

	if v, ok := d.GetOk("logs_queue_size_limit_mb"); ok || d.HasChange("logs_queue_size_limit_mb") {
		t, err := expandSystemSocCollectorLogQueueLogsQueueSizeLimitMbSsclq(d, v, "logs_queue_size_limit_mb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logs-queue-size-limit-mb"] = t
		}
	}

	if v, ok := d.GetOk("win_queue_file_count_limit"); ok || d.HasChange("win_queue_file_count_limit") {
		t, err := expandSystemSocCollectorLogQueueWinQueueFileCountLimitSsclq(d, v, "win_queue_file_count_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["win-queue-file-count-limit"] = t
		}
	}

	if v, ok := d.GetOk("win_queue_size_limit_mb"); ok || d.HasChange("win_queue_size_limit_mb") {
		t, err := expandSystemSocCollectorLogQueueWinQueueSizeLimitMbSsclq(d, v, "win_queue_size_limit_mb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["win-queue-size-limit-mb"] = t
		}
	}

	return &obj, nil
}
