// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: IPv4 local in policy configuration.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocalInPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocalInPolicyCreate,
		Read:   resourceSystemLocalInPolicyRead,
		Update: resourceSystemLocalInPolicyUpdate,
		Delete: resourceSystemLocalInPolicyDelete,

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
						"dst_ip": &schema.Schema{
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemLocalInPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocalInPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLocalInPolicy(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLocalInPolicyRead(d, m)
}

func resourceSystemLocalInPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocalInPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocalInPolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLocalInPolicyRead(d, m)
}

func resourceSystemLocalInPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocalInPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocalInPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocalInPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocalInPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocalInPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocalInPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyDportBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemLocalInPolicyDportBlockDportValue(i["dport-value"], d, pre_append)
			tmp["dport_value"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy-DportBlock-DportValue")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicyDportBlockDportValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyDstBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := i["dst-ip"]; ok {
			v := flattenSystemLocalInPolicyDstBlockDstIp(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy-DstBlock-DstIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicyDstBlockDstIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyDport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLocalInPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyIntfBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemLocalInPolicyIntfBlockIntfName(i["intf-name"], d, pre_append)
			tmp["intf_name"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy-IntfBlock-IntfName")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicyIntfBlockIntfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicySrcBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemLocalInPolicySrcBlockSrcIp(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "SystemLocalInPolicy-SrcBlock-SrcIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLocalInPolicySrcBlockSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicySrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemLocalInPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemLocalInPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemLocalInPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemLocalInPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemLocalInPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if _, ok := o["dport"].([]interface{}); ok {
		if isImportTable() {
			if err = d.Set("dport_block", flattenSystemLocalInPolicyDportBlock(o["dport"], d, "dport_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy-DportBlock"); ok {
					if err = d.Set("dport_block", vv); err != nil {
						return fmt.Errorf("Error reading dport_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dport_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("dport_block"); ok {
				if err = d.Set("dport_block", flattenSystemLocalInPolicyDportBlock(o["dport"], d, "dport_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy-DportBlock"); ok {
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
			if err = d.Set("dst_block", flattenSystemLocalInPolicyDstBlock(o["dst"], d, "dst_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy-DstBlock"); ok {
					if err = d.Set("dst_block", vv); err != nil {
						return fmt.Errorf("Error reading dst_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dst_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("dst_block"); ok {
				if err = d.Set("dst_block", flattenSystemLocalInPolicyDstBlock(o["dst"], d, "dst_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy-DstBlock"); ok {
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
		if err = d.Set("dport", flattenSystemLocalInPolicyDport(o["dport"], d, "dport")); err != nil {
			if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy-Dport"); ok {
				if err = d.Set("dport", vv); err != nil {
					return fmt.Errorf("Error reading dport: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dport: %v", err)
			}
		}
	}

	if err = d.Set("dst", flattenSystemLocalInPolicyDst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLocalInPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLocalInPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if _, ok := o["intf"].([]interface{}); ok {
		if isImportTable() {
			if err = d.Set("intf_block", flattenSystemLocalInPolicyIntfBlock(o["intf"], d, "intf_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy-IntfBlock"); ok {
					if err = d.Set("intf_block", vv); err != nil {
						return fmt.Errorf("Error reading intf_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading intf_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("intf_block"); ok {
				if err = d.Set("intf_block", flattenSystemLocalInPolicyIntfBlock(o["intf"], d, "intf_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy-IntfBlock"); ok {
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
		if err = d.Set("intf", flattenSystemLocalInPolicyIntf(o["intf"], d, "intf")); err != nil {
			if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy-Intf"); ok {
				if err = d.Set("intf", vv); err != nil {
					return fmt.Errorf("Error reading intf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading intf: %v", err)
			}
		}
	}

	if err = d.Set("protocol", flattenSystemLocalInPolicyProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemLocalInPolicy-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if _, ok := o["src"].([]interface{}); ok {
		if isImportTable() {
			if err = d.Set("src_block", flattenSystemLocalInPolicySrcBlock(o["src"], d, "src_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy-SrcBlock"); ok {
					if err = d.Set("src_block", vv); err != nil {
						return fmt.Errorf("Error reading src_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading src_block: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("src_block"); ok {
				if err = d.Set("src_block", flattenSystemLocalInPolicySrcBlock(o["src"], d, "src_block")); err != nil {
					if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy-SrcBlock"); ok {
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

	if err = d.Set("src", flattenSystemLocalInPolicySrc(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	return nil
}

func flattenSystemLocalInPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocalInPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyDportBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dport-value"], _ = expandSystemLocalInPolicyDportBlockDportValue(d, i["dport_value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicyDportBlockDportValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyDstBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-ip"], _ = expandSystemLocalInPolicyDstBlockDstIp(d, i["dst_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicyDstBlockDstIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyDport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemLocalInPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyIntfBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["intf-name"], _ = expandSystemLocalInPolicyIntfBlockIntfName(d, i["intf_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicyIntfBlockIntfName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicySrcBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["src-ip"], _ = expandSystemLocalInPolicySrcBlockSrcIp(d, i["src_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLocalInPolicySrcBlockSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicySrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectSystemLocalInPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemLocalInPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemLocalInPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dport_block"); ok || d.HasChange("dport_block") {
		t, err := expandSystemLocalInPolicyDportBlock(d, v, "dport_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dport"] = t
		}
	}

	if v, ok := d.GetOk("dst_block"); ok || d.HasChange("dst_block") {
		t, err := expandSystemLocalInPolicyDstBlock(d, v, "dst_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dport"); ok || d.HasChange("dport") {
		t, err := expandSystemLocalInPolicyDport(d, v, "dport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dport"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemLocalInPolicyDst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemLocalInPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("intf_block"); ok || d.HasChange("intf_block") {
		t, err := expandSystemLocalInPolicyIntfBlock(d, v, "intf_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("intf"); ok || d.HasChange("intf") {
		t, err := expandSystemLocalInPolicyIntf(d, v, "intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemLocalInPolicyProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("src_block"); ok || d.HasChange("src_block") {
		t, err := expandSystemLocalInPolicySrcBlock(d, v, "src_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandSystemLocalInPolicySrc(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	return &obj, nil
}
