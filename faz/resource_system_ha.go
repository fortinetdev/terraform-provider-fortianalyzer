// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),

// Description: HA configuration.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaUpdate,
		Read:   resourceSystemHaRead,
		Update: resourceSystemHaUpdate,
		Delete: resourceSystemHaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cfg_sync_hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hb_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"initial_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"initial_sync_threads": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"load_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_hb": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"preferred_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_clusterid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_file_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
			},
			"private_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemHaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHa(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemHa")

	return resourceSystemHaRead(d, m)
}

func resourceSystemHaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemHa(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemHa(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaCfgSyncHbIntervalSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaGroupIdSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaGroupNameSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbInterfaceSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbIntervalSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHealthcheckSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaInitialSyncSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaInitialSyncThreadsSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLoadBalanceSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLogSyncSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaModeSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPasswordSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaPeerSha(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemHaPeerIdSha(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemHaPeerIpSha(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_hb"
		if _, ok := i["ip-hb"]; ok {
			v := flattenSystemHaPeerIpHbSha(i["ip-hb"], d, pre_append)
			tmp["ip_hb"] = fortiAPISubPartPatch(v, "SystemHa-Peer-IpHb")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := i["serial-number"]; ok {
			v := flattenSystemHaPeerSerialNumberSha(i["serial-number"], d, pre_append)
			tmp["serial_number"] = fortiAPISubPartPatch(v, "SystemHa-Peer-SerialNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemHaPeerStatusSha(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemHaPeerIdSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIpSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIpHbSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerSerialNumberSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerStatusSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPreferredRoleSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrioritySha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivateClusteridSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivateFileQuotaSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivateHbIntervalSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivateHbLostThresholdSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivateModeSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivatePasswordSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaPrivatePeerSha(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemHaPrivatePeerIdSha(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-PrivatePeer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemHaPrivatePeerIpSha(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemHa-PrivatePeer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenSystemHaPrivatePeerIp6Sha(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "SystemHa-PrivatePeer-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := i["serial-number"]; ok {
			v := flattenSystemHaPrivatePeerSerialNumberSha(i["serial-number"], d, pre_append)
			tmp["serial_number"] = fortiAPISubPartPatch(v, "SystemHa-PrivatePeer-SerialNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemHaPrivatePeerStatusSha(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemHa-PrivatePeer-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemHaPrivatePeerIdSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivatePeerIpSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivatePeerIp6Sha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivatePeerSerialNumberSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrivatePeerStatusSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipInterfaceSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cfg_sync_hb_interval", flattenSystemHaCfgSyncHbIntervalSha(o["cfg-sync-hb-interval"], d, "cfg_sync_hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["cfg-sync-hb-interval"], "SystemHa-CfgSyncHbInterval"); ok {
			if err = d.Set("cfg_sync_hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading cfg_sync_hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cfg_sync_hb_interval: %v", err)
		}
	}

	if err = d.Set("group_id", flattenSystemHaGroupIdSha(o["group-id"], d, "group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-id"], "SystemHa-GroupId"); ok {
			if err = d.Set("group_id", vv); err != nil {
				return fmt.Errorf("Error reading group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_id: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemHaGroupNameSha(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "SystemHa-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("hb_interface", flattenSystemHaHbInterfaceSha(o["hb-interface"], d, "hb_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interface"], "SystemHa-HbInterface"); ok {
			if err = d.Set("hb_interface", vv); err != nil {
				return fmt.Errorf("Error reading hb_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interface: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemHaHbIntervalSha(o["hb-interval"], d, "hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interval"], "SystemHa-HbInterval"); ok {
			if err = d.Set("hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("healthcheck", flattenSystemHaHealthcheckSha(o["healthcheck"], d, "healthcheck")); err != nil {
		if vv, ok := fortiAPIPatch(o["healthcheck"], "SystemHa-Healthcheck"); ok {
			if err = d.Set("healthcheck", vv); err != nil {
				return fmt.Errorf("Error reading healthcheck: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading healthcheck: %v", err)
		}
	}

	if err = d.Set("initial_sync", flattenSystemHaInitialSyncSha(o["initial-sync"], d, "initial_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["initial-sync"], "SystemHa-InitialSync"); ok {
			if err = d.Set("initial_sync", vv); err != nil {
				return fmt.Errorf("Error reading initial_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading initial_sync: %v", err)
		}
	}

	if err = d.Set("initial_sync_threads", flattenSystemHaInitialSyncThreadsSha(o["initial-sync-threads"], d, "initial_sync_threads")); err != nil {
		if vv, ok := fortiAPIPatch(o["initial-sync-threads"], "SystemHa-InitialSyncThreads"); ok {
			if err = d.Set("initial_sync_threads", vv); err != nil {
				return fmt.Errorf("Error reading initial_sync_threads: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading initial_sync_threads: %v", err)
		}
	}

	if err = d.Set("load_balance", flattenSystemHaLoadBalanceSha(o["load-balance"], d, "load_balance")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance"], "SystemHa-LoadBalance"); ok {
			if err = d.Set("load_balance", vv); err != nil {
				return fmt.Errorf("Error reading load_balance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance: %v", err)
		}
	}

	if err = d.Set("log_sync", flattenSystemHaLogSyncSha(o["log-sync"], d, "log_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-sync"], "SystemHa-LogSync"); ok {
			if err = d.Set("log_sync", vv); err != nil {
				return fmt.Errorf("Error reading log_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_sync: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemHaModeSha(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemHa-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("peer", flattenSystemHaPeerSha(o["peer"], d, "peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["peer"], "SystemHa-Peer"); ok {
				if err = d.Set("peer", vv); err != nil {
					return fmt.Errorf("Error reading peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("peer"); ok {
			if err = d.Set("peer", flattenSystemHaPeerSha(o["peer"], d, "peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["peer"], "SystemHa-Peer"); ok {
					if err = d.Set("peer", vv); err != nil {
						return fmt.Errorf("Error reading peer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading peer: %v", err)
				}
			}
		}
	}

	if err = d.Set("preferred_role", flattenSystemHaPreferredRoleSha(o["preferred-role"], d, "preferred_role")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-role"], "SystemHa-PreferredRole"); ok {
			if err = d.Set("preferred_role", vv); err != nil {
				return fmt.Errorf("Error reading preferred_role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_role: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemHaPrioritySha(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemHa-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("private_clusterid", flattenSystemHaPrivateClusteridSha(o["private-clusterid"], d, "private_clusterid")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-clusterid"], "SystemHa-PrivateClusterid"); ok {
			if err = d.Set("private_clusterid", vv); err != nil {
				return fmt.Errorf("Error reading private_clusterid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_clusterid: %v", err)
		}
	}

	if err = d.Set("private_file_quota", flattenSystemHaPrivateFileQuotaSha(o["private-file-quota"], d, "private_file_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-file-quota"], "SystemHa-PrivateFileQuota"); ok {
			if err = d.Set("private_file_quota", vv); err != nil {
				return fmt.Errorf("Error reading private_file_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_file_quota: %v", err)
		}
	}

	if err = d.Set("private_hb_interval", flattenSystemHaPrivateHbIntervalSha(o["private-hb-interval"], d, "private_hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-hb-interval"], "SystemHa-PrivateHbInterval"); ok {
			if err = d.Set("private_hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading private_hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_hb_interval: %v", err)
		}
	}

	if err = d.Set("private_hb_lost_threshold", flattenSystemHaPrivateHbLostThresholdSha(o["private-hb-lost-threshold"], d, "private_hb_lost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-hb-lost-threshold"], "SystemHa-PrivateHbLostThreshold"); ok {
			if err = d.Set("private_hb_lost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading private_hb_lost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("private_mode", flattenSystemHaPrivateModeSha(o["private-mode"], d, "private_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-mode"], "SystemHa-PrivateMode"); ok {
			if err = d.Set("private_mode", vv); err != nil {
				return fmt.Errorf("Error reading private_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("private_peer", flattenSystemHaPrivatePeerSha(o["private-peer"], d, "private_peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["private-peer"], "SystemHa-PrivatePeer"); ok {
				if err = d.Set("private_peer", vv); err != nil {
					return fmt.Errorf("Error reading private_peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading private_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("private_peer"); ok {
			if err = d.Set("private_peer", flattenSystemHaPrivatePeerSha(o["private-peer"], d, "private_peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["private-peer"], "SystemHa-PrivatePeer"); ok {
					if err = d.Set("private_peer", vv); err != nil {
						return fmt.Errorf("Error reading private_peer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading private_peer: %v", err)
				}
			}
		}
	}

	if err = d.Set("unicast", flattenSystemHaUnicastSha(o["unicast"], d, "unicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast"], "SystemHa-Unicast"); ok {
			if err = d.Set("unicast", vv); err != nil {
				return fmt.Errorf("Error reading unicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast: %v", err)
		}
	}

	if err = d.Set("vip", flattenSystemHaVipSha(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "SystemHa-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("vip_interface", flattenSystemHaVipInterfaceSha(o["vip-interface"], d, "vip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip-interface"], "SystemHa-VipInterface"); ok {
			if err = d.Set("vip_interface", vv); err != nil {
				return fmt.Errorf("Error reading vip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip_interface: %v", err)
		}
	}

	return nil
}

func flattenSystemHaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaCfgSyncHbIntervalSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupIdSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupNameSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbInterfaceSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbIntervalSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHealthcheckSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaInitialSyncSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaInitialSyncThreadsSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLoadBalanceSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLogSyncSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaModeSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPasswordSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaPeerSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemHaPeerIdSha(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemHaPeerIpSha(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_hb"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-hb"], _ = expandSystemHaPeerIpHbSha(d, i["ip_hb"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial-number"], _ = expandSystemHaPeerSerialNumberSha(d, i["serial_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemHaPeerStatusSha(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaPeerIdSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIpSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIpHbSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerSerialNumberSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerStatusSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPreferredRoleSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrioritySha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivateClusteridSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivateFileQuotaSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivateHbIntervalSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivateHbLostThresholdSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivateModeSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivatePasswordSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaPrivatePeerSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemHaPrivatePeerIdSha(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemHaPrivatePeerIpSha(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandSystemHaPrivatePeerIp6Sha(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial-number"], _ = expandSystemHaPrivatePeerSerialNumberSha(d, i["serial_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemHaPrivatePeerStatusSha(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaPrivatePeerIdSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivatePeerIpSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivatePeerIp6Sha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivatePeerSerialNumberSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrivatePeerStatusSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipInterfaceSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cfg_sync_hb_interval"); ok || d.HasChange("cfg_sync_hb_interval") {
		t, err := expandSystemHaCfgSyncHbIntervalSha(d, v, "cfg_sync_hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cfg-sync-hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("group_id"); ok || d.HasChange("group_id") {
		t, err := expandSystemHaGroupIdSha(d, v, "group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-id"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandSystemHaGroupNameSha(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("hb_interface"); ok || d.HasChange("hb_interface") {
		t, err := expandSystemHaHbInterfaceSha(d, v, "hb_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interface"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok || d.HasChange("hb_interval") {
		t, err := expandSystemHaHbIntervalSha(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck"); ok || d.HasChange("healthcheck") {
		t, err := expandSystemHaHealthcheckSha(d, v, "healthcheck")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck"] = t
		}
	}

	if v, ok := d.GetOk("initial_sync"); ok || d.HasChange("initial_sync") {
		t, err := expandSystemHaInitialSyncSha(d, v, "initial_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initial-sync"] = t
		}
	}

	if v, ok := d.GetOk("initial_sync_threads"); ok || d.HasChange("initial_sync_threads") {
		t, err := expandSystemHaInitialSyncThreadsSha(d, v, "initial_sync_threads")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initial-sync-threads"] = t
		}
	}

	if v, ok := d.GetOk("load_balance"); ok || d.HasChange("load_balance") {
		t, err := expandSystemHaLoadBalanceSha(d, v, "load_balance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance"] = t
		}
	}

	if v, ok := d.GetOk("log_sync"); ok || d.HasChange("log_sync") {
		t, err := expandSystemHaLogSyncSha(d, v, "log_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-sync"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemHaModeSha(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemHaPasswordSha(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandSystemHaPeerSha(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("preferred_role"); ok || d.HasChange("preferred_role") {
		t, err := expandSystemHaPreferredRoleSha(d, v, "preferred_role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-role"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemHaPrioritySha(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("private_clusterid"); ok || d.HasChange("private_clusterid") {
		t, err := expandSystemHaPrivateClusteridSha(d, v, "private_clusterid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-clusterid"] = t
		}
	}

	if v, ok := d.GetOk("private_file_quota"); ok || d.HasChange("private_file_quota") {
		t, err := expandSystemHaPrivateFileQuotaSha(d, v, "private_file_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-file-quota"] = t
		}
	}

	if v, ok := d.GetOk("private_hb_interval"); ok || d.HasChange("private_hb_interval") {
		t, err := expandSystemHaPrivateHbIntervalSha(d, v, "private_hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("private_hb_lost_threshold"); ok || d.HasChange("private_hb_lost_threshold") {
		t, err := expandSystemHaPrivateHbLostThresholdSha(d, v, "private_hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("private_mode"); ok || d.HasChange("private_mode") {
		t, err := expandSystemHaPrivateModeSha(d, v, "private_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-mode"] = t
		}
	}

	if v, ok := d.GetOk("private_password"); ok || d.HasChange("private_password") {
		t, err := expandSystemHaPrivatePasswordSha(d, v, "private_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-password"] = t
		}
	}

	if v, ok := d.GetOk("private_peer"); ok || d.HasChange("private_peer") {
		t, err := expandSystemHaPrivatePeerSha(d, v, "private_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-peer"] = t
		}
	}

	if v, ok := d.GetOk("unicast"); ok || d.HasChange("unicast") {
		t, err := expandSystemHaUnicastSha(d, v, "unicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandSystemHaVipSha(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("vip_interface"); ok || d.HasChange("vip_interface") {
		t, err := expandSystemHaVipInterfaceSha(d, v, "vip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip-interface"] = t
		}
	}

	return &obj, nil
}
