// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: IPv6 local in policy configuration.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocalInPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocalInPolicy6Create,
		Read:   resourceSystemLocalInPolicy6Read,
		Update: resourceSystemLocalInPolicy6Update,
		Delete: resourceSystemLocalInPolicy6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dport_block": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dport_value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dst_block": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"intf_block": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"intf_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_block": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemLocalInPolicy6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLocalInPolicy6(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLocalInPolicy6Read(d, m)
}

func resourceSystemLocalInPolicy6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocalInPolicy6(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLocalInPolicy6Read(d, m)
}

func resourceSystemLocalInPolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocalInPolicy6(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocalInPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocalInPolicy6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocalInPolicy6(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocalInPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocalInPolicy6Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Description(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6DportBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport_value"
		if _, ok := i["dport-value"]; ok {
			v := flattenSystemLocalInPolicy6DportBlockDportValue(i["dport-value"], d, pre_append)
			tmp["dport_value"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy6-DportBlock-DportValue")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicy6DportBlockDportValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6DstBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenSystemLocalInPolicy6DstBlockSrcIp(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy6-DstBlock-SrcIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicy6DstBlockSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Dport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6IntfBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intf_name"
		if _, ok := i["intf-name"]; ok {
			v := flattenSystemLocalInPolicy6IntfBlockIntfName(i["intf-name"], d, pre_append)
			tmp["intf_name"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy6-IntfBlock-IntfName")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicy6IntfBlockIntfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Intf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6SrcBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenSystemLocalInPolicy6SrcBlockSrcIp(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy6-SrcBlock-SrcIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicy6SrcBlockSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocalInPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemLocalInPolicy6Action(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemLocalInPolicy6-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemLocalInPolicy6Description(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemLocalInPolicy6-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if _, ok := o["dport"].([]interface{}); ok {
		if isImportTable() {
			if err = d.Set("dport_block", flattenSystemLocalInPolicy6DportBlock(o["dport"], d, "dport_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy6-DportBlock"); ok {
					if err = d.Set("dport_block", vv); err != nil {
						return fmt.Errorf("Error reading dport_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dport_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("dport_block"); ok {
				if err = d.Set("dport_block", flattenSystemLocalInPolicy6DportBlock(o["dport"], d, "dport_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy6-DportBlock"); ok {
						if err = d.Set("dport_block", vv); err != nil {
							return fmt.Errorf("Error reading dport_block: %v", err)
						}
					} else {
						return fmt.Errorf("Error reading dport_block: %v", err)
					}
				}
			}
		}
	}

	if _, ok := o["dst"].([]interface{}); ok {
		if isImportTable() {
			if err = d.Set("dst_block", flattenSystemLocalInPolicy6DstBlock(o["dst"], d, "dst_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy6-DstBlock"); ok {
					if err = d.Set("dst_block", vv); err != nil {
						return fmt.Errorf("Error reading dst_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dst_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("dst_block"); ok {
				if err = d.Set("dst_block", flattenSystemLocalInPolicy6DstBlock(o["dst"], d, "dst_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy6-DstBlock"); ok {
						if err = d.Set("dst_block", vv); err != nil {
							return fmt.Errorf("Error reading dst_block: %v", err)
						}
					} else {
						return fmt.Errorf("Error reading dst_block: %v", err)
					}
				}
			}
		}
	}

	if _, ok := o["dport"].(float64); ok {
		if err = d.Set("dport", flattenSystemLocalInPolicy6Dport(o["dport"], d, "dport")); err != nil {
			if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy6-Dport"); ok {
				if err = d.Set("dport", vv); err != nil {
					return fmt.Errorf("Error reading dport: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dport: %v", err)
			}
		}
	}

	if _, ok := o["dst"].(string); ok {
		if err = d.Set("dst", flattenSystemLocalInPolicy6Dst(o["dst"], d, "dst")); err != nil {
			if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy6-Dst"); ok {
				if err = d.Set("dst", vv); err != nil {
					return fmt.Errorf("Error reading dst: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		}
	}

	if err = d.Set("fosid", flattenSystemLocalInPolicy6Id(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLocalInPolicy6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if _, ok := o["intf"].([]interface{}); ok {
		if isImportTable() {
			if err = d.Set("intf_block", flattenSystemLocalInPolicy6IntfBlock(o["intf"], d, "intf_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy6-IntfBlock"); ok {
					if err = d.Set("intf_block", vv); err != nil {
						return fmt.Errorf("Error reading intf_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading intf_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("intf_block"); ok {
				if err = d.Set("intf_block", flattenSystemLocalInPolicy6IntfBlock(o["intf"], d, "intf_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy6-IntfBlock"); ok {
						if err = d.Set("intf_block", vv); err != nil {
							return fmt.Errorf("Error reading intf_block: %v", err)
						}
					} else {
						return fmt.Errorf("Error reading intf_block: %v", err)
					}
				}
			}
		}
	}

	if _, ok := o["intf"].(string); ok {
		if err = d.Set("intf", flattenSystemLocalInPolicy6Intf(o["intf"], d, "intf")); err != nil {
			if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy6-Intf"); ok {
				if err = d.Set("intf", vv); err != nil {
					return fmt.Errorf("Error reading intf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading intf: %v", err)
			}
		}
	}

	if err = d.Set("protocol", flattenSystemLocalInPolicy6Protocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemLocalInPolicy6-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if _, ok := o["src"].([]interface{}); ok {
		if isImportTable() {
			if err = d.Set("src_block", flattenSystemLocalInPolicy6SrcBlock(o["src"], d, "src_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy6-SrcBlock"); ok {
					if err = d.Set("src_block", vv); err != nil {
						return fmt.Errorf("Error reading src_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading src_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("src_block"); ok {
				if err = d.Set("src_block", flattenSystemLocalInPolicy6SrcBlock(o["src"], d, "src_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy6-SrcBlock"); ok {
						if err = d.Set("src_block", vv); err != nil {
							return fmt.Errorf("Error reading src_block: %v", err)
						}
					} else {
						return fmt.Errorf("Error reading src_block: %v", err)
					}
				}
			}
		}
	}

	if _, ok := o["src"].(string); ok {
		if err = d.Set("src", flattenSystemLocalInPolicy6Src(o["src"], d, "src")); err != nil {
			if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy6-Src"); ok {
				if err = d.Set("src", vv); err != nil {
					return fmt.Errorf("Error reading src: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading src: %v", err)
			}
		}
	}

	return nil
}

func flattenSystemLocalInPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocalInPolicy6Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Description(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6DportBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dport-value"], _ = expandSystemLocalInPolicy6DportBlockDportValue(d, i["dport_value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicy6DportBlockDportValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6DstBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandSystemLocalInPolicy6DstBlockSrcIp(d, i["src_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicy6DstBlockSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Dport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Dst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6IntfBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intf_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["intf-name"], _ = expandSystemLocalInPolicy6IntfBlockIntfName(d, i["intf_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicy6IntfBlockIntfName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Intf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6SrcBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandSystemLocalInPolicy6SrcBlockSrcIp(d, i["src_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicy6SrcBlockSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Src(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocalInPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemLocalInPolicy6Action(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemLocalInPolicy6Description(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dport_block"); ok || d.HasChange("dport_block") {
		t, err := expandSystemLocalInPolicy6DportBlock(d, v, "dport_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dport"] = t
		}
	}

	if v, ok := d.GetOk("dst_block"); ok || d.HasChange("dst_block") {
		t, err := expandSystemLocalInPolicy6DstBlock(d, v, "dst_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dport"); ok || d.HasChange("dport") {
		t, err := expandSystemLocalInPolicy6Dport(d, v, "dport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dport"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemLocalInPolicy6Dst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemLocalInPolicy6Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("intf_block"); ok || d.HasChange("intf_block") {
		t, err := expandSystemLocalInPolicy6IntfBlock(d, v, "intf_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("intf"); ok || d.HasChange("intf") {
		t, err := expandSystemLocalInPolicy6Intf(d, v, "intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemLocalInPolicy6Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("src_block"); ok || d.HasChange("src_block") {
		t, err := expandSystemLocalInPolicy6SrcBlock(d, v, "src_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandSystemLocalInPolicy6Src(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	return &obj, nil
}
