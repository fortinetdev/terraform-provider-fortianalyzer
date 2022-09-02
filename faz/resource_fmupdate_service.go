// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Enable/disable services provided by the built-in FortiGuard.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateService() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateServiceUpdate,
		Read:   resourceFmupdateServiceRead,
		Update: resourceFmupdateServiceUpdate,
		Delete: resourceFmupdateServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"avips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateService(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateService resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateService(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateService")

	return resourceFmupdateServiceRead(d, m)
}

func resourceFmupdateServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateService resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateServiceAvips(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("avips", flattenFmupdateServiceAvips(o["avips"], d, "avips")); err != nil {
		if vv, ok := fortiAPIPatch(o["avips"], "FmupdateService-Avips"); ok {
			if err = d.Set("avips", vv); err != nil {
				return fmt.Errorf("Error reading avips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading avips: %v", err)
		}
	}

	return nil
}

func flattenFmupdateServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateServiceAvips(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("avips"); ok || d.HasChange("avips") {
		t, err := expandFmupdateServiceAvips(d, v, "avips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["avips"] = t
		}
	}

	return &obj, nil
}
