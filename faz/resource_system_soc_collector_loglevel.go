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

func resourceSystemSocCollectorLogLevel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSocCollectorLogLevelUpdate,
		Read:   resourceSystemSocCollectorLogLevelRead,
		Update: resourceSystemSocCollectorLogLevelUpdate,
		Delete: resourceSystemSocCollectorLogLevelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phagentmanager": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phdiscover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pheventforwarder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pheventhandler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pheventpackager": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phparser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phperfmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSocCollectorLogLevelUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSocCollectorLogLevel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocCollectorLogLevel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSocCollectorLogLevel(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocCollectorLogLevel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSocCollectorLogLevel")

	return resourceSystemSocCollectorLogLevelRead(d, m)
}

func resourceSystemSocCollectorLogLevelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSocCollectorLogLevel(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSocCollectorLogLevel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSocCollectorLogLevelRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSocCollectorLogLevel(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocCollectorLogLevel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSocCollectorLogLevel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocCollectorLogLevel resource from API: %v", err)
	}
	return nil
}

func flattenSystemSocCollectorLogLevelAllSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhagentmanagerSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhdiscoverSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPheventforwarderSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPheventhandlerSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPheventpackagerSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhmonitorSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhparserSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhperfmonitorSscll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSocCollectorLogLevel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("all", flattenSystemSocCollectorLogLevelAllSscll(o["all"], d, "all")); err != nil {
		if vv, ok := fortiAPIPatch(o["all"], "SystemSocCollectorLogLevel-All"); ok {
			if err = d.Set("all", vv); err != nil {
				return fmt.Errorf("Error reading all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading all: %v", err)
		}
	}

	if err = d.Set("phagentmanager", flattenSystemSocCollectorLogLevelPhagentmanagerSscll(o["phAgentManager"], d, "phagentmanager")); err != nil {
		if vv, ok := fortiAPIPatch(o["phAgentManager"], "SystemSocCollectorLogLevel-Phagentmanager"); ok {
			if err = d.Set("phagentmanager", vv); err != nil {
				return fmt.Errorf("Error reading phagentmanager: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phagentmanager: %v", err)
		}
	}

	if err = d.Set("phdiscover", flattenSystemSocCollectorLogLevelPhdiscoverSscll(o["phDiscover"], d, "phdiscover")); err != nil {
		if vv, ok := fortiAPIPatch(o["phDiscover"], "SystemSocCollectorLogLevel-Phdiscover"); ok {
			if err = d.Set("phdiscover", vv); err != nil {
				return fmt.Errorf("Error reading phdiscover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phdiscover: %v", err)
		}
	}

	if err = d.Set("pheventforwarder", flattenSystemSocCollectorLogLevelPheventforwarderSscll(o["phEventForwarder"], d, "pheventforwarder")); err != nil {
		if vv, ok := fortiAPIPatch(o["phEventForwarder"], "SystemSocCollectorLogLevel-Pheventforwarder"); ok {
			if err = d.Set("pheventforwarder", vv); err != nil {
				return fmt.Errorf("Error reading pheventforwarder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pheventforwarder: %v", err)
		}
	}

	if err = d.Set("pheventhandler", flattenSystemSocCollectorLogLevelPheventhandlerSscll(o["phEventHandler"], d, "pheventhandler")); err != nil {
		if vv, ok := fortiAPIPatch(o["phEventHandler"], "SystemSocCollectorLogLevel-Pheventhandler"); ok {
			if err = d.Set("pheventhandler", vv); err != nil {
				return fmt.Errorf("Error reading pheventhandler: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pheventhandler: %v", err)
		}
	}

	if err = d.Set("pheventpackager", flattenSystemSocCollectorLogLevelPheventpackagerSscll(o["phEventPackager"], d, "pheventpackager")); err != nil {
		if vv, ok := fortiAPIPatch(o["phEventPackager"], "SystemSocCollectorLogLevel-Pheventpackager"); ok {
			if err = d.Set("pheventpackager", vv); err != nil {
				return fmt.Errorf("Error reading pheventpackager: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pheventpackager: %v", err)
		}
	}

	if err = d.Set("phmonitor", flattenSystemSocCollectorLogLevelPhmonitorSscll(o["phMonitor"], d, "phmonitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["phMonitor"], "SystemSocCollectorLogLevel-Phmonitor"); ok {
			if err = d.Set("phmonitor", vv); err != nil {
				return fmt.Errorf("Error reading phmonitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phmonitor: %v", err)
		}
	}

	if err = d.Set("phparser", flattenSystemSocCollectorLogLevelPhparserSscll(o["phParser"], d, "phparser")); err != nil {
		if vv, ok := fortiAPIPatch(o["phParser"], "SystemSocCollectorLogLevel-Phparser"); ok {
			if err = d.Set("phparser", vv); err != nil {
				return fmt.Errorf("Error reading phparser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phparser: %v", err)
		}
	}

	if err = d.Set("phperfmonitor", flattenSystemSocCollectorLogLevelPhperfmonitorSscll(o["phPerfMonitor"], d, "phperfmonitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["phPerfMonitor"], "SystemSocCollectorLogLevel-Phperfmonitor"); ok {
			if err = d.Set("phperfmonitor", vv); err != nil {
				return fmt.Errorf("Error reading phperfmonitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phperfmonitor: %v", err)
		}
	}

	return nil
}

func flattenSystemSocCollectorLogLevelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSocCollectorLogLevelAllSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhagentmanagerSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhdiscoverSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPheventforwarderSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPheventhandlerSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPheventpackagerSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhmonitorSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhparserSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhperfmonitorSscll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSocCollectorLogLevel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("all"); ok || d.HasChange("all") {
		t, err := expandSystemSocCollectorLogLevelAllSscll(d, v, "all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all"] = t
		}
	}

	if v, ok := d.GetOk("phagentmanager"); ok || d.HasChange("phagentmanager") {
		t, err := expandSystemSocCollectorLogLevelPhagentmanagerSscll(d, v, "phagentmanager")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phAgentManager"] = t
		}
	}

	if v, ok := d.GetOk("phdiscover"); ok || d.HasChange("phdiscover") {
		t, err := expandSystemSocCollectorLogLevelPhdiscoverSscll(d, v, "phdiscover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phDiscover"] = t
		}
	}

	if v, ok := d.GetOk("pheventforwarder"); ok || d.HasChange("pheventforwarder") {
		t, err := expandSystemSocCollectorLogLevelPheventforwarderSscll(d, v, "pheventforwarder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phEventForwarder"] = t
		}
	}

	if v, ok := d.GetOk("pheventhandler"); ok || d.HasChange("pheventhandler") {
		t, err := expandSystemSocCollectorLogLevelPheventhandlerSscll(d, v, "pheventhandler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phEventHandler"] = t
		}
	}

	if v, ok := d.GetOk("pheventpackager"); ok || d.HasChange("pheventpackager") {
		t, err := expandSystemSocCollectorLogLevelPheventpackagerSscll(d, v, "pheventpackager")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phEventPackager"] = t
		}
	}

	if v, ok := d.GetOk("phmonitor"); ok || d.HasChange("phmonitor") {
		t, err := expandSystemSocCollectorLogLevelPhmonitorSscll(d, v, "phmonitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phMonitor"] = t
		}
	}

	if v, ok := d.GetOk("phparser"); ok || d.HasChange("phparser") {
		t, err := expandSystemSocCollectorLogLevelPhparserSscll(d, v, "phparser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phParser"] = t
		}
	}

	if v, ok := d.GetOk("phperfmonitor"); ok || d.HasChange("phperfmonitor") {
		t, err := expandSystemSocCollectorLogLevelPhperfmonitorSscll(d, v, "phperfmonitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phPerfMonitor"] = t
		}
	}

	return &obj, nil
}
