// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: System LogSettingsClient-Cert-Auth

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogSettingsClientCertAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogSettingsClientCertAuthUpdate,
		Read:   resourceSystemLogSettingsClientCertAuthRead,
		Update: resourceSystemLogSettingsClientCertAuthUpdate,
		Delete: resourceSystemLogSettingsClientCertAuthDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_client": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemLogSettingsClientCertAuthUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogSettingsClientCertAuth(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettingsClientCertAuth resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogSettingsClientCertAuth(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettingsClientCertAuth resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogSettingsClientCertAuth")

	return resourceSystemLogSettingsClientCertAuthRead(d, m)
}

func resourceSystemLogSettingsClientCertAuthDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogSettingsClientCertAuth(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogSettingsClientCertAuth resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogSettingsClientCertAuthRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogSettingsClientCertAuth(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettingsClientCertAuth resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogSettingsClientCertAuth(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettingsClientCertAuth resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogSettingsClientCertAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsClientCertAuthTlsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsClientCertAuthTrustedClient(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := i["certificate"]; ok {
			v := flattenSystemLogSettingsClientCertAuthTrustedClientCertificate(i["certificate"], d, pre_append)
			tmp["certificate"] = fortiAPISubPartPatch(v, "SystemLogSettingsClientCertAuth-TrustedClient-Certificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSystemLogSettingsClientCertAuthTrustedClientDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SystemLogSettingsClientCertAuth-TrustedClient-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenSystemLogSettingsClientCertAuthTrustedClientDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "SystemLogSettingsClientCertAuth-TrustedClient-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogSettingsClientCertAuthTrustedClientId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogSettingsClientCertAuth-TrustedClient-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemLogSettingsClientCertAuthTrustedClientType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemLogSettingsClientCertAuth-TrustedClient-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogSettingsClientCertAuthTrustedClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsClientCertAuthTrustedClientDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsClientCertAuthTrustedClientDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsClientCertAuthTrustedClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsClientCertAuthTrustedClientType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogSettingsClientCertAuth(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mode", flattenSystemLogSettingsClientCertAuthMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemLogSettingsClientCertAuth-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("tls_port", flattenSystemLogSettingsClientCertAuthTlsPort(o["tls-port"], d, "tls_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-port"], "SystemLogSettingsClientCertAuth-TlsPort"); ok {
			if err = d.Set("tls_port", vv); err != nil {
				return fmt.Errorf("Error reading tls_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("trusted_client", flattenSystemLogSettingsClientCertAuthTrustedClient(o["trusted-client"], d, "trusted_client")); err != nil {
			if vv, ok := fortiAPIPatch(o["trusted-client"], "SystemLogSettingsClientCertAuth-TrustedClient"); ok {
				if err = d.Set("trusted_client", vv); err != nil {
					return fmt.Errorf("Error reading trusted_client: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading trusted_client: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusted_client"); ok {
			if err = d.Set("trusted_client", flattenSystemLogSettingsClientCertAuthTrustedClient(o["trusted-client"], d, "trusted_client")); err != nil {
				if vv, ok := fortiAPIPatch(o["trusted-client"], "SystemLogSettingsClientCertAuth-TrustedClient"); ok {
					if err = d.Set("trusted_client", vv); err != nil {
						return fmt.Errorf("Error reading trusted_client: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading trusted_client: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemLogSettingsClientCertAuthFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogSettingsClientCertAuthMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsClientCertAuthTlsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsClientCertAuthTrustedClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["certificate"], _ = expandSystemLogSettingsClientCertAuthTrustedClientCertificate(d, i["certificate"], pre_append)
		} else {
			tmp["certificate"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSystemLogSettingsClientCertAuthTrustedClientDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandSystemLogSettingsClientCertAuthTrustedClientDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogSettingsClientCertAuthTrustedClientId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemLogSettingsClientCertAuthTrustedClientType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogSettingsClientCertAuthTrustedClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsClientCertAuthTrustedClientDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsClientCertAuthTrustedClientDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsClientCertAuthTrustedClientId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsClientCertAuthTrustedClientType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogSettingsClientCertAuth(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemLogSettingsClientCertAuthMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("tls_port"); ok || d.HasChange("tls_port") {
		t, err := expandSystemLogSettingsClientCertAuthTlsPort(d, v, "tls_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-port"] = t
		}
	}

	if v, ok := d.GetOk("trusted_client"); ok || d.HasChange("trusted_client") {
		t, err := expandSystemLogSettingsClientCertAuthTrustedClient(d, v, "trusted_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-client"] = t
		}
	}

	return &obj, nil
}
