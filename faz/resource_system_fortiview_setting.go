// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: FortiView settings.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFortiviewSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortiviewSettingUpdate,
		Read:   resourceSystemFortiviewSettingRead,
		Update: resourceSystemFortiviewSettingUpdate,
		Delete: resourceSystemFortiviewSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"data_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"not_scanned_apps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_run_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolve_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemFortiviewSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemFortiviewSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiviewSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFortiviewSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiviewSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFortiviewSetting")

	return resourceSystemFortiviewSettingRead(d, m)
}

func resourceSystemFortiviewSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemFortiviewSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFortiviewSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortiviewSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemFortiviewSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiviewSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortiviewSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiviewSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortiviewSettingDataSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiviewSettingNotScannedApps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiviewSettingQueryRunMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiviewSettingResolveIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFortiviewSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("data_source", flattenSystemFortiviewSettingDataSource(o["data-source"], d, "data_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-source"], "SystemFortiviewSetting-DataSource"); ok {
			if err = d.Set("data_source", vv); err != nil {
				return fmt.Errorf("Error reading data_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_source: %v", err)
		}
	}

	if err = d.Set("not_scanned_apps", flattenSystemFortiviewSettingNotScannedApps(o["not-scanned-apps"], d, "not_scanned_apps")); err != nil {
		if vv, ok := fortiAPIPatch(o["not-scanned-apps"], "SystemFortiviewSetting-NotScannedApps"); ok {
			if err = d.Set("not_scanned_apps", vv); err != nil {
				return fmt.Errorf("Error reading not_scanned_apps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading not_scanned_apps: %v", err)
		}
	}

	if err = d.Set("query_run_mode", flattenSystemFortiviewSettingQueryRunMode(o["query-run-mode"], d, "query_run_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-run-mode"], "SystemFortiviewSetting-QueryRunMode"); ok {
			if err = d.Set("query_run_mode", vv); err != nil {
				return fmt.Errorf("Error reading query_run_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_run_mode: %v", err)
		}
	}

	if err = d.Set("resolve_ip", flattenSystemFortiviewSettingResolveIp(o["resolve-ip"], d, "resolve_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["resolve-ip"], "SystemFortiviewSetting-ResolveIp"); ok {
			if err = d.Set("resolve_ip", vv); err != nil {
				return fmt.Errorf("Error reading resolve_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resolve_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemFortiviewSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFortiviewSettingDataSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiviewSettingNotScannedApps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiviewSettingQueryRunMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiviewSettingResolveIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortiviewSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("data_source"); ok || d.HasChange("data_source") {
		t, err := expandSystemFortiviewSettingDataSource(d, v, "data_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-source"] = t
		}
	}

	if v, ok := d.GetOk("not_scanned_apps"); ok || d.HasChange("not_scanned_apps") {
		t, err := expandSystemFortiviewSettingNotScannedApps(d, v, "not_scanned_apps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["not-scanned-apps"] = t
		}
	}

	if v, ok := d.GetOk("query_run_mode"); ok || d.HasChange("query_run_mode") {
		t, err := expandSystemFortiviewSettingQueryRunMode(d, v, "query_run_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-run-mode"] = t
		}
	}

	if v, ok := d.GetOk("resolve_ip"); ok || d.HasChange("resolve_ip") {
		t, err := expandSystemFortiviewSettingResolveIp(d, v, "resolve_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-ip"] = t
		}
	}

	return &obj, nil
}
