// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: UEBAsettings.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogUeba() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogUebaUpdate,
		Read:   resourceSystemLogUebaRead,
		Update: resourceSystemLogUebaUpdate,
		Delete: resourceSystemLogUebaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"hostname_ep_unifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_only_ep": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_unique_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogUebaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogUeba(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogUeba resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogUeba(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogUeba resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogUeba")

	return resourceSystemLogUebaRead(d, m)
}

func resourceSystemLogUebaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogUeba(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogUeba resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogUebaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogUeba(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogUeba resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogUeba(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogUeba resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogUebaHostnameEpUnifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogUebaIpOnlyEp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogUebaIpUniqueScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogUeba(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("hostname_ep_unifier", flattenSystemLogUebaHostnameEpUnifier(o["hostname-ep-unifier"], d, "hostname_ep_unifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname-ep-unifier"], "SystemLogUeba-HostnameEpUnifier"); ok {
			if err = d.Set("hostname_ep_unifier", vv); err != nil {
				return fmt.Errorf("Error reading hostname_ep_unifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname_ep_unifier: %v", err)
		}
	}

	if err = d.Set("ip_only_ep", flattenSystemLogUebaIpOnlyEp(o["ip-only-ep"], d, "ip_only_ep")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-only-ep"], "SystemLogUeba-IpOnlyEp"); ok {
			if err = d.Set("ip_only_ep", vv); err != nil {
				return fmt.Errorf("Error reading ip_only_ep: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_only_ep: %v", err)
		}
	}

	if err = d.Set("ip_unique_scope", flattenSystemLogUebaIpUniqueScope(o["ip-unique-scope"], d, "ip_unique_scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-unique-scope"], "SystemLogUeba-IpUniqueScope"); ok {
			if err = d.Set("ip_unique_scope", vv); err != nil {
				return fmt.Errorf("Error reading ip_unique_scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_unique_scope: %v", err)
		}
	}

	return nil
}

func flattenSystemLogUebaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogUebaHostnameEpUnifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogUebaIpOnlyEp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogUebaIpUniqueScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogUeba(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hostname_ep_unifier"); ok || d.HasChange("hostname_ep_unifier") {
		t, err := expandSystemLogUebaHostnameEpUnifier(d, v, "hostname_ep_unifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname-ep-unifier"] = t
		}
	}

	if v, ok := d.GetOk("ip_only_ep"); ok || d.HasChange("ip_only_ep") {
		t, err := expandSystemLogUebaIpOnlyEp(d, v, "ip_only_ep")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-only-ep"] = t
		}
	}

	if v, ok := d.GetOk("ip_unique_scope"); ok || d.HasChange("ip_unique_scope") {
		t, err := expandSystemLogUebaIpUniqueScope(d, v, "ip_unique_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-unique-scope"] = t
		}
	}

	return &obj, nil
}
