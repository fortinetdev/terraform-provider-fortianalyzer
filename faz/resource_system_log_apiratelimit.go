// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: System LogApi-Ratelimit

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogApiRatelimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogApiRatelimitUpdate,
		Read:   resourceSystemLogApiRatelimitRead,
		Update: resourceSystemLogApiRatelimitUpdate,
		Delete: resourceSystemLogApiRatelimitDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"read_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"write_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogApiRatelimitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogApiRatelimit(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogApiRatelimit resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogApiRatelimit(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogApiRatelimit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogApiRatelimit")

	return resourceSystemLogApiRatelimitRead(d, m)
}

func resourceSystemLogApiRatelimitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogApiRatelimit(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogApiRatelimit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogApiRatelimitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogApiRatelimit(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogApiRatelimit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogApiRatelimit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogApiRatelimit resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogApiRatelimitReadLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogApiRatelimitWriteLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogApiRatelimit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("read_limit", flattenSystemLogApiRatelimitReadLimit(o["read-limit"], d, "read_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["read-limit"], "SystemLogApiRatelimit-ReadLimit"); ok {
			if err = d.Set("read_limit", vv); err != nil {
				return fmt.Errorf("Error reading read_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading read_limit: %v", err)
		}
	}

	if err = d.Set("write_limit", flattenSystemLogApiRatelimitWriteLimit(o["write-limit"], d, "write_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["write-limit"], "SystemLogApiRatelimit-WriteLimit"); ok {
			if err = d.Set("write_limit", vv); err != nil {
				return fmt.Errorf("Error reading write_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading write_limit: %v", err)
		}
	}

	return nil
}

func flattenSystemLogApiRatelimitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogApiRatelimitReadLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogApiRatelimitWriteLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogApiRatelimit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("read_limit"); ok || d.HasChange("read_limit") {
		t, err := expandSystemLogApiRatelimitReadLimit(d, v, "read_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["read-limit"] = t
		}
	}

	if v, ok := d.GetOk("write_limit"); ok || d.HasChange("write_limit") {
		t, err := expandSystemLogApiRatelimitWriteLimit(d, v, "write_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["write-limit"] = t
		}
	}

	return &obj, nil
}
