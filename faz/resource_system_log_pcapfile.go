// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Log pcap-file settings.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogPcapFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogPcapFileUpdate,
		Read:   resourceSystemLogPcapFileRead,
		Update: resourceSystemLogPcapFileUpdate,
		Delete: resourceSystemLogPcapFileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"download_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogPcapFileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogPcapFile(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogPcapFile resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogPcapFile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogPcapFile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogPcapFile")

	return resourceSystemLogPcapFileRead(d, m)
}

func resourceSystemLogPcapFileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogPcapFile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogPcapFile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogPcapFileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogPcapFile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogPcapFile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogPcapFile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogPcapFile resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogPcapFileDownloadMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogPcapFile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("download_mode", flattenSystemLogPcapFileDownloadMode(o["download-mode"], d, "download_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["download-mode"], "SystemLogPcapFile-DownloadMode"); ok {
			if err = d.Set("download_mode", vv); err != nil {
				return fmt.Errorf("Error reading download_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading download_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemLogPcapFileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogPcapFileDownloadMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogPcapFile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("download_mode"); ok || d.HasChange("download_mode") {
		t, err := expandSystemLogPcapFileDownloadMode(d, v, "download_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["download-mode"] = t
		}
	}

	return &obj, nil
}
