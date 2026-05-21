// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: System LocallogTacacs+AccountingSetting

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocallogTacacsAccountingSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogTacacsAccountingSettingUpdate,
		Read:   resourceSystemLocallogTacacsAccountingSettingRead,
		Update: resourceSystemLocallogTacacsAccountingSettingUpdate,
		Delete: resourceSystemLocallogTacacsAccountingSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tacacs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogTacacsAccountingSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogTacacsAccountingSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogTacacsAccountingSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogTacacsAccountingSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogTacacsAccountingSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogTacacsAccountingSetting")

	return resourceSystemLocallogTacacsAccountingSettingRead(d, m)
}

func resourceSystemLocallogTacacsAccountingSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogTacacsAccountingSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogTacacsAccountingSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogTacacsAccountingSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogTacacsAccountingSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogTacacsAccountingSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogTacacsAccountingSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogTacacsAccountingSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogTacacsAccountingSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogTacacsAccountingSettingTacacsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogTacacsAccountingSettingTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogTacacsAccountingSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemLocallogTacacsAccountingSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLocallogTacacsAccountingSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tacacs_name", flattenSystemLocallogTacacsAccountingSettingTacacsName(o["tacacs-name"], d, "tacacs_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["tacacs-name"], "SystemLocallogTacacsAccountingSetting-TacacsName"); ok {
			if err = d.Set("tacacs_name", vv); err != nil {
				return fmt.Errorf("Error reading tacacs_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tacacs_name: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemLocallogTacacsAccountingSettingTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "SystemLocallogTacacsAccountingSetting-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogTacacsAccountingSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogTacacsAccountingSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogTacacsAccountingSettingTacacsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogTacacsAccountingSettingTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogTacacsAccountingSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLocallogTacacsAccountingSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tacacs_name"); ok || d.HasChange("tacacs_name") {
		t, err := expandSystemLocallogTacacsAccountingSettingTacacsName(d, v, "tacacs_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs-name"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandSystemLocallogTacacsAccountingSettingTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
