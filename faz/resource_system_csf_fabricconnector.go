// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Fabric connector configuration.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsfFabricConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfFabricConnectorCreate,
		Read:   resourceSystemCsfFabricConnectorRead,
		Update: resourceSystemCsfFabricConnectorUpdate,
		Delete: resourceSystemCsfFabricConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration_write_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemCsfFabricConnectorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCsfFabricConnector(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfFabricConnector resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCsfFabricConnector(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemCsfFabricConnector resource: %v", err)
	}

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemCsfFabricConnectorRead(d, m)
}

func resourceSystemCsfFabricConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCsfFabricConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfFabricConnector resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCsfFabricConnector(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfFabricConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemCsfFabricConnectorRead(d, m)
}

func resourceSystemCsfFabricConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemCsfFabricConnector(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsfFabricConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfFabricConnectorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemCsfFabricConnector(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfFabricConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsfFabricConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfFabricConnector resource from API: %v", err)
	}
	return nil
}

func refreshObjectSystemCsfFabricConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accprofile", flattenSystemCsfFabricConnectorAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile"], "SystemCsfFabricConnector-Accprofile"); ok {
			if err = d.Set("accprofile", vv); err != nil {
				return fmt.Errorf("Error reading accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("configuration_write_access", flattenSystemCsfFabricConnectorConfigurationWriteAccess(o["configuration-write-access"], d, "configuration_write_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["configuration-write-access"], "SystemCsfFabricConnector-ConfigurationWriteAccess"); ok {
			if err = d.Set("configuration_write_access", vv); err != nil {
				return fmt.Errorf("Error reading configuration_write_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading configuration_write_access: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemCsfFabricConnectorSerial(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemCsfFabricConnector-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfFabricConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func getObjectSystemCsfFabricConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accprofile"); ok || d.HasChange("accprofile") {
		t, err := expandSystemCsfFabricConnectorAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("configuration_write_access"); ok || d.HasChange("configuration_write_access") {
		t, err := expandSystemCsfFabricConnectorConfigurationWriteAccess(d, v, "configuration_write_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configuration-write-access"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemCsfFabricConnectorSerial(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	return &obj, nil
}
