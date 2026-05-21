// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: System LocallogTacacs+AccountingFilter

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocallogTacacsAccountingFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogTacacsAccountingFilterUpdate,
		Read:   resourceSystemLocallogTacacsAccountingFilterRead,
		Update: resourceSystemLocallogTacacsAccountingFilterUpdate,
		Delete: resourceSystemLocallogTacacsAccountingFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cli_cmd_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_change_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogTacacsAccountingFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogTacacsAccountingFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogTacacsAccountingFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogTacacsAccountingFilter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogTacacsAccountingFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogTacacsAccountingFilter")

	return resourceSystemLocallogTacacsAccountingFilterRead(d, m)
}

func resourceSystemLocallogTacacsAccountingFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogTacacsAccountingFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogTacacsAccountingFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogTacacsAccountingFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogTacacsAccountingFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogTacacsAccountingFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogTacacsAccountingFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogTacacsAccountingFilter resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogTacacsAccountingFilterCliCmdAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogTacacsAccountingFilterConfigChangeAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogTacacsAccountingFilterLoginAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogTacacsAccountingFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cli_cmd_audit", flattenSystemLocallogTacacsAccountingFilterCliCmdAudit(o["cli-cmd-audit"], d, "cli_cmd_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-cmd-audit"], "SystemLocallogTacacsAccountingFilter-CliCmdAudit"); ok {
			if err = d.Set("cli_cmd_audit", vv); err != nil {
				return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
		}
	}

	if err = d.Set("config_change_audit", flattenSystemLocallogTacacsAccountingFilterConfigChangeAudit(o["config-change-audit"], d, "config_change_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-change-audit"], "SystemLocallogTacacsAccountingFilter-ConfigChangeAudit"); ok {
			if err = d.Set("config_change_audit", vv); err != nil {
				return fmt.Errorf("Error reading config_change_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_change_audit: %v", err)
		}
	}

	if err = d.Set("login_audit", flattenSystemLocallogTacacsAccountingFilterLoginAudit(o["login-audit"], d, "login_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-audit"], "SystemLocallogTacacsAccountingFilter-LoginAudit"); ok {
			if err = d.Set("login_audit", vv); err != nil {
				return fmt.Errorf("Error reading login_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_audit: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogTacacsAccountingFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogTacacsAccountingFilterCliCmdAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogTacacsAccountingFilterConfigChangeAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogTacacsAccountingFilterLoginAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogTacacsAccountingFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cli_cmd_audit"); ok || d.HasChange("cli_cmd_audit") {
		t, err := expandSystemLocallogTacacsAccountingFilterCliCmdAudit(d, v, "cli_cmd_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-cmd-audit"] = t
		}
	}

	if v, ok := d.GetOk("config_change_audit"); ok || d.HasChange("config_change_audit") {
		t, err := expandSystemLocallogTacacsAccountingFilterConfigChangeAudit(d, v, "config_change_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-change-audit"] = t
		}
	}

	if v, ok := d.GetOk("login_audit"); ok || d.HasChange("login_audit") {
		t, err := expandSystemLocallogTacacsAccountingFilterLoginAudit(d, v, "login_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-audit"] = t
		}
	}

	return &obj, nil
}
