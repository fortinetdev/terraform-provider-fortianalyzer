// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Log forwarding service.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogForwardService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogForwardServiceUpdate,
		Read:   resourceSystemLogForwardServiceRead,
		Update: resourceSystemLogForwardServiceUpdate,
		Delete: resourceSystemLogForwardServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"accept_aggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregation_disk_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"collector_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogForwardServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogForwardService(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogForwardService resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogForwardService(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogForwardService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogForwardService")

	return resourceSystemLogForwardServiceRead(d, m)
}

func resourceSystemLogForwardServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogForwardService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogForwardService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogForwardServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogForwardService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogForwardService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogForwardService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogForwardService resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogForwardServiceAcceptAggregation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardServiceAggregationDiskQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardServiceCollectorAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogForwardService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accept_aggregation", flattenSystemLogForwardServiceAcceptAggregation(o["accept-aggregation"], d, "accept_aggregation")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-aggregation"], "SystemLogForwardService-AcceptAggregation"); ok {
			if err = d.Set("accept_aggregation", vv); err != nil {
				return fmt.Errorf("Error reading accept_aggregation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_aggregation: %v", err)
		}
	}

	if err = d.Set("aggregation_disk_quota", flattenSystemLogForwardServiceAggregationDiskQuota(o["aggregation-disk-quota"], d, "aggregation_disk_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregation-disk-quota"], "SystemLogForwardService-AggregationDiskQuota"); ok {
			if err = d.Set("aggregation_disk_quota", vv); err != nil {
				return fmt.Errorf("Error reading aggregation_disk_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregation_disk_quota: %v", err)
		}
	}

	if err = d.Set("collector_auth", flattenSystemLogForwardServiceCollectorAuth(o["collector-auth"], d, "collector_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-auth"], "SystemLogForwardService-CollectorAuth"); ok {
			if err = d.Set("collector_auth", vv); err != nil {
				return fmt.Errorf("Error reading collector_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_auth: %v", err)
		}
	}

	return nil
}

func flattenSystemLogForwardServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogForwardServiceAcceptAggregation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardServiceAggregationDiskQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardServiceCollectorAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogForwardService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_aggregation"); ok || d.HasChange("accept_aggregation") {
		t, err := expandSystemLogForwardServiceAcceptAggregation(d, v, "accept_aggregation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-aggregation"] = t
		}
	}

	if v, ok := d.GetOk("aggregation_disk_quota"); ok || d.HasChange("aggregation_disk_quota") {
		t, err := expandSystemLogForwardServiceAggregationDiskQuota(d, v, "aggregation_disk_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregation-disk-quota"] = t
		}
	}

	if v, ok := d.GetOk("collector_auth"); ok || d.HasChange("collector_auth") {
		t, err := expandSystemLogForwardServiceCollectorAuth(d, v, "collector_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-auth"] = t
		}
	}

	return &obj, nil
}
