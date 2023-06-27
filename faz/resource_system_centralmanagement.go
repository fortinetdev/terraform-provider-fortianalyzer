// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Central management configuration.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCentralManagementUpdate,
		Read:   resourceSystemCentralManagementRead,
		Update: resourceSystemCentralManagementUpdate,
		Delete: resourceSystemCentralManagementDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"acctid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorized_manager_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"elite_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmtid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCentralManagementUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCentralManagement(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCentralManagement(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemCentralManagement")

	return resourceSystemCentralManagementRead(d, m)
}

func resourceSystemCentralManagementDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemCentralManagement(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCentralManagement resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemCentralManagement(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCentralManagement(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource from API: %v", err)
	}
	return nil
}

func flattenSystemCentralManagementAcctid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAuthorizedManagerOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementEliteService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFmg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementMgmtid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2int(v)
}

func flattenSystemCentralManagementSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCentralManagementType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("acctid", flattenSystemCentralManagementAcctid(o["acctid"], d, "acctid")); err != nil {
		if vv, ok := fortiAPIPatch(o["acctid"], "SystemCentralManagement-Acctid"); ok {
			if err = d.Set("acctid", vv); err != nil {
				return fmt.Errorf("Error reading acctid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acctid: %v", err)
		}
	}

	if err = d.Set("allow_monitor", flattenSystemCentralManagementAllowMonitor(o["allow-monitor"], d, "allow_monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-monitor"], "SystemCentralManagement-AllowMonitor"); ok {
			if err = d.Set("allow_monitor", vv); err != nil {
				return fmt.Errorf("Error reading allow_monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_monitor: %v", err)
		}
	}

	if err = d.Set("authorized_manager_only", flattenSystemCentralManagementAuthorizedManagerOnly(o["authorized-manager-only"], d, "authorized_manager_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorized-manager-only"], "SystemCentralManagement-AuthorizedManagerOnly"); ok {
			if err = d.Set("authorized_manager_only", vv); err != nil {
				return fmt.Errorf("Error reading authorized_manager_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorized_manager_only: %v", err)
		}
	}

	if err = d.Set("elite_service", flattenSystemCentralManagementEliteService(o["elite-service"], d, "elite_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["elite-service"], "SystemCentralManagement-EliteService"); ok {
			if err = d.Set("elite_service", vv); err != nil {
				return fmt.Errorf("Error reading elite_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading elite_service: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemCentralManagementEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "SystemCentralManagement-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("fmg", flattenSystemCentralManagementFmg(o["fmg"], d, "fmg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmg"], "SystemCentralManagement-Fmg"); ok {
			if err = d.Set("fmg", vv); err != nil {
				return fmt.Errorf("Error reading fmg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmg: %v", err)
		}
	}

	if err = d.Set("mgmtid", flattenSystemCentralManagementMgmtid(o["mgmtid"], d, "mgmtid")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgmtid"], "SystemCentralManagement-Mgmtid"); ok {
			if err = d.Set("mgmtid", vv); err != nil {
				return fmt.Errorf("Error reading mgmtid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgmtid: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemCentralManagementSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-number"], "SystemCentralManagement-SerialNumber"); ok {
			if err = d.Set("serial_number", vv); err != nil {
				return fmt.Errorf("Error reading serial_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemCentralManagementType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemCentralManagement-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCentralManagementAcctid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAuthorizedManagerOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementEliteService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementMgmtid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCentralManagementType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCentralManagement(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acctid"); ok || d.HasChange("acctid") {
		t, err := expandSystemCentralManagementAcctid(d, v, "acctid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acctid"] = t
		}
	}

	if v, ok := d.GetOk("allow_monitor"); ok || d.HasChange("allow_monitor") {
		t, err := expandSystemCentralManagementAllowMonitor(d, v, "allow_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-monitor"] = t
		}
	}

	if v, ok := d.GetOk("authorized_manager_only"); ok || d.HasChange("authorized_manager_only") {
		t, err := expandSystemCentralManagementAuthorizedManagerOnly(d, v, "authorized_manager_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized-manager-only"] = t
		}
	}

	if v, ok := d.GetOk("elite_service"); ok || d.HasChange("elite_service") {
		t, err := expandSystemCentralManagementEliteService(d, v, "elite_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["elite-service"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandSystemCentralManagementEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("fmg"); ok || d.HasChange("fmg") {
		t, err := expandSystemCentralManagementFmg(d, v, "fmg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg"] = t
		}
	}

	if v, ok := d.GetOk("mgmtid"); ok || d.HasChange("mgmtid") {
		t, err := expandSystemCentralManagementMgmtid(d, v, "mgmtid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmtid"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok || d.HasChange("serial_number") {
		t, err := expandSystemCentralManagementSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemCentralManagementType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
