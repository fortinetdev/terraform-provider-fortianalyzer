// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Pre-authorized security fabric nodes

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSocFabricTrustedList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSocFabricTrustedListCreate,
		Read:   resourceSystemSocFabricTrustedListRead,
		Update: resourceSystemSocFabricTrustedListUpdate,
		Delete: resourceSystemSocFabricTrustedListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSocFabricTrustedListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSocFabricTrustedList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSocFabricTrustedList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSocFabricTrustedList(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemSocFabricTrustedList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSocFabricTrustedListRead(d, m)
}

func resourceSystemSocFabricTrustedListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSocFabricTrustedList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabricTrustedList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSocFabricTrustedList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabricTrustedList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSocFabricTrustedListRead(d, m)
}

func resourceSystemSocFabricTrustedListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSocFabricTrustedList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSocFabricTrustedList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSocFabricTrustedListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSocFabricTrustedList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabricTrustedList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSocFabricTrustedList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabricTrustedList resource from API: %v", err)
	}
	return nil
}

func refreshObjectSystemSocFabricTrustedList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemSocFabricTrustedListId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSocFabricTrustedList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemSocFabricTrustedListSerial(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemSocFabricTrustedList-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	return nil
}

func flattenSystemSocFabricTrustedListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func getObjectSystemSocFabricTrustedList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSocFabricTrustedListId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemSocFabricTrustedListSerial(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	return &obj, nil
}
