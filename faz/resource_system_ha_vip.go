// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: VIPs.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaVip() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaVipCreate,
		Read:   resourceSystemHaVipRead,
		Update: resourceSystemHaVipUpdate,
		Delete: resourceSystemHaVipDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemHaVipCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaVip(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaVip resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaVip(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemHaVip resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaVipRead(d, m)
}

func resourceSystemHaVipUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaVip(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaVip resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaVip(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaVip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaVipRead(d, m)
}

func resourceSystemHaVipDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemHaVip(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaVip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaVipRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemHaVip(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaVip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaVip(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaVip resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaVipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipVipInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaVip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemHaVipId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemHaVip-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemHaVipStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemHaVip-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vip", flattenSystemHaVipVip(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "SystemHaVip-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("vip_interface", flattenSystemHaVipVipInterface(o["vip-interface"], d, "vip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip-interface"], "SystemHaVip-VipInterface"); ok {
			if err = d.Set("vip_interface", vv); err != nil {
				return fmt.Errorf("Error reading vip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip_interface: %v", err)
		}
	}

	return nil
}

func flattenSystemHaVipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaVipId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipVipInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaVip(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandSystemHaVipId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemHaVipStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandSystemHaVipVip(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("vip_interface"); ok || d.HasChange("vip_interface") {
		t, err := expandSystemHaVipVipInterface(d, v, "vip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip-interface"] = t
		}
	}

	return &obj, nil
}
