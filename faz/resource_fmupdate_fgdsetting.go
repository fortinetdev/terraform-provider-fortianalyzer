// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Fmupdate Fgd-Setting

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFgdSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFgdSettingUpdate,
		Read:   resourceFmupdateFgdSettingRead,
		Update: resourceFmupdateFgdSettingUpdate,
		Delete: resourceFmupdateFgdSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"as_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"as_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"av_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av2_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"av2_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av2_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eventlog_query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd_pull_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fq_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fq_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fq_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iot_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iot_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iotv_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linkd_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_client_worker": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_log_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_unrated_site": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"restrict_as1_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_as2_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_as4_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_av_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_av2_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_fq_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_iots_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_wf_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"servlist": &schema.Schema{
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
										Computed: true,
									},
									"ip6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"service_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"stat_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stat_log_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stat_sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"update_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wf_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wf_dn_cache_expire_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wf_dn_cache_max_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wf_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wf_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFgdSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFgdSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFgdSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFgdSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFgdSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFgdSetting")

	return resourceFmupdateFgdSettingRead(d, m)
}

func resourceFmupdateFgdSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateFgdSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFgdSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFgdSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateFgdSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFgdSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFgdSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFgdSetting resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFgdSettingAsCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAsPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAvCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAvLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAvPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAv2Cache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAv2Log(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingAv2Preload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingEventlogQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingFgdPullInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingFqCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingFqLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingFqPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingIotCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingIotLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingIotPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingIotvPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingLinkdLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingMaxClientWorker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingMaxLogQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingMaxUnratedSite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictAs1Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictAs2Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictAs4Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictAvDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictAv2Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictFqDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictIotsDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingRestrictWfDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingServerOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := i["servlist"]; ok {
		result["servlist"] = flattenFmupdateFgdSettingServerOverrideServlist(i["servlist"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFgdSettingServerOverrideStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFgdSettingServerOverrideServlist(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateFgdSettingServerOverrideServlistId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFgdSettingServerOverride-Servlist-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFgdSettingServerOverrideServlistIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFgdSettingServerOverride-Servlist-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateFgdSettingServerOverrideServlistIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateFgdSettingServerOverride-Servlist-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFgdSettingServerOverrideServlistPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFgdSettingServerOverride-Servlist-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := i["service-type"]; ok {
			v := flattenFmupdateFgdSettingServerOverrideServlistServiceType(i["service-type"], d, pre_append)
			tmp["service_type"] = fortiAPISubPartPatch(v, "FmupdateFgdSettingServerOverride-Servlist-ServiceType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFmupdateFgdSettingServerOverrideServlistId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingServerOverrideServlistIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingServerOverrideServlistIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingServerOverrideServlistPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingServerOverrideServlistServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingServerOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingStatLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingStatLogInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingStatSyncInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingUpdateLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingWfCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingWfDnCacheExpireTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingWfDnCacheMaxNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingWfLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFgdSettingWfPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFgdSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("as_cache", flattenFmupdateFgdSettingAsCache(o["as-cache"], d, "as_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-cache"], "FmupdateFgdSetting-AsCache"); ok {
			if err = d.Set("as_cache", vv); err != nil {
				return fmt.Errorf("Error reading as_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_cache: %v", err)
		}
	}

	if err = d.Set("as_log", flattenFmupdateFgdSettingAsLog(o["as-log"], d, "as_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-log"], "FmupdateFgdSetting-AsLog"); ok {
			if err = d.Set("as_log", vv); err != nil {
				return fmt.Errorf("Error reading as_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_log: %v", err)
		}
	}

	if err = d.Set("as_preload", flattenFmupdateFgdSettingAsPreload(o["as-preload"], d, "as_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-preload"], "FmupdateFgdSetting-AsPreload"); ok {
			if err = d.Set("as_preload", vv); err != nil {
				return fmt.Errorf("Error reading as_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_preload: %v", err)
		}
	}

	if err = d.Set("av_cache", flattenFmupdateFgdSettingAvCache(o["av-cache"], d, "av_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-cache"], "FmupdateFgdSetting-AvCache"); ok {
			if err = d.Set("av_cache", vv); err != nil {
				return fmt.Errorf("Error reading av_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_cache: %v", err)
		}
	}

	if err = d.Set("av_log", flattenFmupdateFgdSettingAvLog(o["av-log"], d, "av_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-log"], "FmupdateFgdSetting-AvLog"); ok {
			if err = d.Set("av_log", vv); err != nil {
				return fmt.Errorf("Error reading av_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_log: %v", err)
		}
	}

	if err = d.Set("av_preload", flattenFmupdateFgdSettingAvPreload(o["av-preload"], d, "av_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-preload"], "FmupdateFgdSetting-AvPreload"); ok {
			if err = d.Set("av_preload", vv); err != nil {
				return fmt.Errorf("Error reading av_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_preload: %v", err)
		}
	}

	if err = d.Set("av2_cache", flattenFmupdateFgdSettingAv2Cache(o["av2-cache"], d, "av2_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-cache"], "FmupdateFgdSetting-Av2Cache"); ok {
			if err = d.Set("av2_cache", vv); err != nil {
				return fmt.Errorf("Error reading av2_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_cache: %v", err)
		}
	}

	if err = d.Set("av2_log", flattenFmupdateFgdSettingAv2Log(o["av2-log"], d, "av2_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-log"], "FmupdateFgdSetting-Av2Log"); ok {
			if err = d.Set("av2_log", vv); err != nil {
				return fmt.Errorf("Error reading av2_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_log: %v", err)
		}
	}

	if err = d.Set("av2_preload", flattenFmupdateFgdSettingAv2Preload(o["av2-preload"], d, "av2_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-preload"], "FmupdateFgdSetting-Av2Preload"); ok {
			if err = d.Set("av2_preload", vv); err != nil {
				return fmt.Errorf("Error reading av2_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_preload: %v", err)
		}
	}

	if err = d.Set("eventlog_query", flattenFmupdateFgdSettingEventlogQuery(o["eventlog-query"], d, "eventlog_query")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventlog-query"], "FmupdateFgdSetting-EventlogQuery"); ok {
			if err = d.Set("eventlog_query", vv); err != nil {
				return fmt.Errorf("Error reading eventlog_query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventlog_query: %v", err)
		}
	}

	if err = d.Set("fgd_pull_interval", flattenFmupdateFgdSettingFgdPullInterval(o["fgd-pull-interval"], d, "fgd_pull_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-pull-interval"], "FmupdateFgdSetting-FgdPullInterval"); ok {
			if err = d.Set("fgd_pull_interval", vv); err != nil {
				return fmt.Errorf("Error reading fgd_pull_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_pull_interval: %v", err)
		}
	}

	if err = d.Set("fq_cache", flattenFmupdateFgdSettingFqCache(o["fq-cache"], d, "fq_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-cache"], "FmupdateFgdSetting-FqCache"); ok {
			if err = d.Set("fq_cache", vv); err != nil {
				return fmt.Errorf("Error reading fq_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_cache: %v", err)
		}
	}

	if err = d.Set("fq_log", flattenFmupdateFgdSettingFqLog(o["fq-log"], d, "fq_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-log"], "FmupdateFgdSetting-FqLog"); ok {
			if err = d.Set("fq_log", vv); err != nil {
				return fmt.Errorf("Error reading fq_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_log: %v", err)
		}
	}

	if err = d.Set("fq_preload", flattenFmupdateFgdSettingFqPreload(o["fq-preload"], d, "fq_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-preload"], "FmupdateFgdSetting-FqPreload"); ok {
			if err = d.Set("fq_preload", vv); err != nil {
				return fmt.Errorf("Error reading fq_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_preload: %v", err)
		}
	}

	if err = d.Set("iot_cache", flattenFmupdateFgdSettingIotCache(o["iot-cache"], d, "iot_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-cache"], "FmupdateFgdSetting-IotCache"); ok {
			if err = d.Set("iot_cache", vv); err != nil {
				return fmt.Errorf("Error reading iot_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_cache: %v", err)
		}
	}

	if err = d.Set("iot_log", flattenFmupdateFgdSettingIotLog(o["iot-log"], d, "iot_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-log"], "FmupdateFgdSetting-IotLog"); ok {
			if err = d.Set("iot_log", vv); err != nil {
				return fmt.Errorf("Error reading iot_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_log: %v", err)
		}
	}

	if err = d.Set("iot_preload", flattenFmupdateFgdSettingIotPreload(o["iot-preload"], d, "iot_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-preload"], "FmupdateFgdSetting-IotPreload"); ok {
			if err = d.Set("iot_preload", vv); err != nil {
				return fmt.Errorf("Error reading iot_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_preload: %v", err)
		}
	}

	if err = d.Set("iotv_preload", flattenFmupdateFgdSettingIotvPreload(o["iotv-preload"], d, "iotv_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["iotv-preload"], "FmupdateFgdSetting-IotvPreload"); ok {
			if err = d.Set("iotv_preload", vv); err != nil {
				return fmt.Errorf("Error reading iotv_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iotv_preload: %v", err)
		}
	}

	if err = d.Set("linkd_log", flattenFmupdateFgdSettingLinkdLog(o["linkd-log"], d, "linkd_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["linkd-log"], "FmupdateFgdSetting-LinkdLog"); ok {
			if err = d.Set("linkd_log", vv); err != nil {
				return fmt.Errorf("Error reading linkd_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linkd_log: %v", err)
		}
	}

	if err = d.Set("max_client_worker", flattenFmupdateFgdSettingMaxClientWorker(o["max-client-worker"], d, "max_client_worker")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-client-worker"], "FmupdateFgdSetting-MaxClientWorker"); ok {
			if err = d.Set("max_client_worker", vv); err != nil {
				return fmt.Errorf("Error reading max_client_worker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_client_worker: %v", err)
		}
	}

	if err = d.Set("max_log_quota", flattenFmupdateFgdSettingMaxLogQuota(o["max-log-quota"], d, "max_log_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-quota"], "FmupdateFgdSetting-MaxLogQuota"); ok {
			if err = d.Set("max_log_quota", vv); err != nil {
				return fmt.Errorf("Error reading max_log_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_quota: %v", err)
		}
	}

	if err = d.Set("max_unrated_site", flattenFmupdateFgdSettingMaxUnratedSite(o["max-unrated-site"], d, "max_unrated_site")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-unrated-site"], "FmupdateFgdSetting-MaxUnratedSite"); ok {
			if err = d.Set("max_unrated_site", vv); err != nil {
				return fmt.Errorf("Error reading max_unrated_site: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_unrated_site: %v", err)
		}
	}

	if err = d.Set("restrict_as1_dbver", flattenFmupdateFgdSettingRestrictAs1Dbver(o["restrict-as1-dbver"], d, "restrict_as1_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as1-dbver"], "FmupdateFgdSetting-RestrictAs1Dbver"); ok {
			if err = d.Set("restrict_as1_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as1_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as1_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_as2_dbver", flattenFmupdateFgdSettingRestrictAs2Dbver(o["restrict-as2-dbver"], d, "restrict_as2_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as2-dbver"], "FmupdateFgdSetting-RestrictAs2Dbver"); ok {
			if err = d.Set("restrict_as2_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as2_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as2_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_as4_dbver", flattenFmupdateFgdSettingRestrictAs4Dbver(o["restrict-as4-dbver"], d, "restrict_as4_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as4-dbver"], "FmupdateFgdSetting-RestrictAs4Dbver"); ok {
			if err = d.Set("restrict_as4_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as4_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as4_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_av_dbver", flattenFmupdateFgdSettingRestrictAvDbver(o["restrict-av-dbver"], d, "restrict_av_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-av-dbver"], "FmupdateFgdSetting-RestrictAvDbver"); ok {
			if err = d.Set("restrict_av_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_av_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_av_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_av2_dbver", flattenFmupdateFgdSettingRestrictAv2Dbver(o["restrict-av2-dbver"], d, "restrict_av2_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-av2-dbver"], "FmupdateFgdSetting-RestrictAv2Dbver"); ok {
			if err = d.Set("restrict_av2_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_av2_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_av2_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_fq_dbver", flattenFmupdateFgdSettingRestrictFqDbver(o["restrict-fq-dbver"], d, "restrict_fq_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-fq-dbver"], "FmupdateFgdSetting-RestrictFqDbver"); ok {
			if err = d.Set("restrict_fq_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_fq_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_fq_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_iots_dbver", flattenFmupdateFgdSettingRestrictIotsDbver(o["restrict-iots-dbver"], d, "restrict_iots_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-iots-dbver"], "FmupdateFgdSetting-RestrictIotsDbver"); ok {
			if err = d.Set("restrict_iots_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_iots_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_iots_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_wf_dbver", flattenFmupdateFgdSettingRestrictWfDbver(o["restrict-wf-dbver"], d, "restrict_wf_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-wf-dbver"], "FmupdateFgdSetting-RestrictWfDbver"); ok {
			if err = d.Set("restrict_wf_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_wf_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_wf_dbver: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_override", flattenFmupdateFgdSettingServerOverride(o["server-override"], d, "server_override")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-override"], "FmupdateFgdSetting-ServerOverride"); ok {
				if err = d.Set("server_override", vv); err != nil {
					return fmt.Errorf("Error reading server_override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_override"); ok {
			if err = d.Set("server_override", flattenFmupdateFgdSettingServerOverride(o["server-override"], d, "server_override")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-override"], "FmupdateFgdSetting-ServerOverride"); ok {
					if err = d.Set("server_override", vv); err != nil {
						return fmt.Errorf("Error reading server_override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_override: %v", err)
				}
			}
		}
	}

	if err = d.Set("stat_log", flattenFmupdateFgdSettingStatLog(o["stat-log"], d, "stat_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-log"], "FmupdateFgdSetting-StatLog"); ok {
			if err = d.Set("stat_log", vv); err != nil {
				return fmt.Errorf("Error reading stat_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_log: %v", err)
		}
	}

	if err = d.Set("stat_log_interval", flattenFmupdateFgdSettingStatLogInterval(o["stat-log-interval"], d, "stat_log_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-log-interval"], "FmupdateFgdSetting-StatLogInterval"); ok {
			if err = d.Set("stat_log_interval", vv); err != nil {
				return fmt.Errorf("Error reading stat_log_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_log_interval: %v", err)
		}
	}

	if err = d.Set("stat_sync_interval", flattenFmupdateFgdSettingStatSyncInterval(o["stat-sync-interval"], d, "stat_sync_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-sync-interval"], "FmupdateFgdSetting-StatSyncInterval"); ok {
			if err = d.Set("stat_sync_interval", vv); err != nil {
				return fmt.Errorf("Error reading stat_sync_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_sync_interval: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenFmupdateFgdSettingUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "FmupdateFgdSetting-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("update_log", flattenFmupdateFgdSettingUpdateLog(o["update-log"], d, "update_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-log"], "FmupdateFgdSetting-UpdateLog"); ok {
			if err = d.Set("update_log", vv); err != nil {
				return fmt.Errorf("Error reading update_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_log: %v", err)
		}
	}

	if err = d.Set("wf_cache", flattenFmupdateFgdSettingWfCache(o["wf-cache"], d, "wf_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-cache"], "FmupdateFgdSetting-WfCache"); ok {
			if err = d.Set("wf_cache", vv); err != nil {
				return fmt.Errorf("Error reading wf_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_cache: %v", err)
		}
	}

	if err = d.Set("wf_dn_cache_expire_time", flattenFmupdateFgdSettingWfDnCacheExpireTime(o["wf-dn-cache-expire-time"], d, "wf_dn_cache_expire_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-dn-cache-expire-time"], "FmupdateFgdSetting-WfDnCacheExpireTime"); ok {
			if err = d.Set("wf_dn_cache_expire_time", vv); err != nil {
				return fmt.Errorf("Error reading wf_dn_cache_expire_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_dn_cache_expire_time: %v", err)
		}
	}

	if err = d.Set("wf_dn_cache_max_number", flattenFmupdateFgdSettingWfDnCacheMaxNumber(o["wf-dn-cache-max-number"], d, "wf_dn_cache_max_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-dn-cache-max-number"], "FmupdateFgdSetting-WfDnCacheMaxNumber"); ok {
			if err = d.Set("wf_dn_cache_max_number", vv); err != nil {
				return fmt.Errorf("Error reading wf_dn_cache_max_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_dn_cache_max_number: %v", err)
		}
	}

	if err = d.Set("wf_log", flattenFmupdateFgdSettingWfLog(o["wf-log"], d, "wf_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-log"], "FmupdateFgdSetting-WfLog"); ok {
			if err = d.Set("wf_log", vv); err != nil {
				return fmt.Errorf("Error reading wf_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_log: %v", err)
		}
	}

	if err = d.Set("wf_preload", flattenFmupdateFgdSettingWfPreload(o["wf-preload"], d, "wf_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-preload"], "FmupdateFgdSetting-WfPreload"); ok {
			if err = d.Set("wf_preload", vv); err != nil {
				return fmt.Errorf("Error reading wf_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_preload: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFgdSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFgdSettingAsCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAsPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAvCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAvLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAvPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAv2Cache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAv2Log(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingAv2Preload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingEventlogQuery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingFgdPullInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingFqCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingFqLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingFqPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingIotCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingIotLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingIotPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingIotvPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingLinkdLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingMaxClientWorker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingMaxLogQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingMaxUnratedSite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictAs1Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictAs2Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictAs4Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictAvDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictAv2Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictFqDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictIotsDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingRestrictWfDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["servlist"], _ = expandFmupdateFgdSettingServerOverrideServlist(d, i["servlist"], pre_append)
	} else {
		result["servlist"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFgdSettingServerOverrideStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateFgdSettingServerOverrideServlist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFmupdateFgdSettingServerOverrideServlistId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateFgdSettingServerOverrideServlistIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandFmupdateFgdSettingServerOverrideServlistIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateFgdSettingServerOverrideServlistPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-type"], _ = expandFmupdateFgdSettingServerOverrideServlistServiceType(d, i["service_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFmupdateFgdSettingServerOverrideServlistId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingServerOverrideServlistIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingServerOverrideServlistIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingServerOverrideServlistPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingServerOverrideServlistServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingServerOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingStatLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingStatLogInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingStatSyncInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingUpdateLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingWfCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingWfDnCacheExpireTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingWfDnCacheMaxNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingWfLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFgdSettingWfPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFgdSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("as_cache"); ok || d.HasChange("as_cache") {
		t, err := expandFmupdateFgdSettingAsCache(d, v, "as_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-cache"] = t
		}
	}

	if v, ok := d.GetOk("as_log"); ok || d.HasChange("as_log") {
		t, err := expandFmupdateFgdSettingAsLog(d, v, "as_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-log"] = t
		}
	}

	if v, ok := d.GetOk("as_preload"); ok || d.HasChange("as_preload") {
		t, err := expandFmupdateFgdSettingAsPreload(d, v, "as_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-preload"] = t
		}
	}

	if v, ok := d.GetOk("av_cache"); ok || d.HasChange("av_cache") {
		t, err := expandFmupdateFgdSettingAvCache(d, v, "av_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-cache"] = t
		}
	}

	if v, ok := d.GetOk("av_log"); ok || d.HasChange("av_log") {
		t, err := expandFmupdateFgdSettingAvLog(d, v, "av_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-log"] = t
		}
	}

	if v, ok := d.GetOk("av_preload"); ok || d.HasChange("av_preload") {
		t, err := expandFmupdateFgdSettingAvPreload(d, v, "av_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-preload"] = t
		}
	}

	if v, ok := d.GetOk("av2_cache"); ok || d.HasChange("av2_cache") {
		t, err := expandFmupdateFgdSettingAv2Cache(d, v, "av2_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-cache"] = t
		}
	}

	if v, ok := d.GetOk("av2_log"); ok || d.HasChange("av2_log") {
		t, err := expandFmupdateFgdSettingAv2Log(d, v, "av2_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-log"] = t
		}
	}

	if v, ok := d.GetOk("av2_preload"); ok || d.HasChange("av2_preload") {
		t, err := expandFmupdateFgdSettingAv2Preload(d, v, "av2_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-preload"] = t
		}
	}

	if v, ok := d.GetOk("eventlog_query"); ok || d.HasChange("eventlog_query") {
		t, err := expandFmupdateFgdSettingEventlogQuery(d, v, "eventlog_query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventlog-query"] = t
		}
	}

	if v, ok := d.GetOk("fgd_pull_interval"); ok || d.HasChange("fgd_pull_interval") {
		t, err := expandFmupdateFgdSettingFgdPullInterval(d, v, "fgd_pull_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-pull-interval"] = t
		}
	}

	if v, ok := d.GetOk("fq_cache"); ok || d.HasChange("fq_cache") {
		t, err := expandFmupdateFgdSettingFqCache(d, v, "fq_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-cache"] = t
		}
	}

	if v, ok := d.GetOk("fq_log"); ok || d.HasChange("fq_log") {
		t, err := expandFmupdateFgdSettingFqLog(d, v, "fq_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-log"] = t
		}
	}

	if v, ok := d.GetOk("fq_preload"); ok || d.HasChange("fq_preload") {
		t, err := expandFmupdateFgdSettingFqPreload(d, v, "fq_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-preload"] = t
		}
	}

	if v, ok := d.GetOk("iot_cache"); ok || d.HasChange("iot_cache") {
		t, err := expandFmupdateFgdSettingIotCache(d, v, "iot_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-cache"] = t
		}
	}

	if v, ok := d.GetOk("iot_log"); ok || d.HasChange("iot_log") {
		t, err := expandFmupdateFgdSettingIotLog(d, v, "iot_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-log"] = t
		}
	}

	if v, ok := d.GetOk("iot_preload"); ok || d.HasChange("iot_preload") {
		t, err := expandFmupdateFgdSettingIotPreload(d, v, "iot_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-preload"] = t
		}
	}

	if v, ok := d.GetOk("iotv_preload"); ok || d.HasChange("iotv_preload") {
		t, err := expandFmupdateFgdSettingIotvPreload(d, v, "iotv_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iotv-preload"] = t
		}
	}

	if v, ok := d.GetOk("linkd_log"); ok || d.HasChange("linkd_log") {
		t, err := expandFmupdateFgdSettingLinkdLog(d, v, "linkd_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linkd-log"] = t
		}
	}

	if v, ok := d.GetOk("max_client_worker"); ok || d.HasChange("max_client_worker") {
		t, err := expandFmupdateFgdSettingMaxClientWorker(d, v, "max_client_worker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-client-worker"] = t
		}
	}

	if v, ok := d.GetOk("max_log_quota"); ok || d.HasChange("max_log_quota") {
		t, err := expandFmupdateFgdSettingMaxLogQuota(d, v, "max_log_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-quota"] = t
		}
	}

	if v, ok := d.GetOk("max_unrated_site"); ok || d.HasChange("max_unrated_site") {
		t, err := expandFmupdateFgdSettingMaxUnratedSite(d, v, "max_unrated_site")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-unrated-site"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as1_dbver"); ok || d.HasChange("restrict_as1_dbver") {
		t, err := expandFmupdateFgdSettingRestrictAs1Dbver(d, v, "restrict_as1_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as1-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as2_dbver"); ok || d.HasChange("restrict_as2_dbver") {
		t, err := expandFmupdateFgdSettingRestrictAs2Dbver(d, v, "restrict_as2_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as2-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as4_dbver"); ok || d.HasChange("restrict_as4_dbver") {
		t, err := expandFmupdateFgdSettingRestrictAs4Dbver(d, v, "restrict_as4_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as4-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_av_dbver"); ok || d.HasChange("restrict_av_dbver") {
		t, err := expandFmupdateFgdSettingRestrictAvDbver(d, v, "restrict_av_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-av-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_av2_dbver"); ok || d.HasChange("restrict_av2_dbver") {
		t, err := expandFmupdateFgdSettingRestrictAv2Dbver(d, v, "restrict_av2_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-av2-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_fq_dbver"); ok || d.HasChange("restrict_fq_dbver") {
		t, err := expandFmupdateFgdSettingRestrictFqDbver(d, v, "restrict_fq_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-fq-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_iots_dbver"); ok || d.HasChange("restrict_iots_dbver") {
		t, err := expandFmupdateFgdSettingRestrictIotsDbver(d, v, "restrict_iots_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-iots-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_wf_dbver"); ok || d.HasChange("restrict_wf_dbver") {
		t, err := expandFmupdateFgdSettingRestrictWfDbver(d, v, "restrict_wf_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-wf-dbver"] = t
		}
	}

	if v, ok := d.GetOk("server_override"); ok || d.HasChange("server_override") {
		t, err := expandFmupdateFgdSettingServerOverride(d, v, "server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-override"] = t
		}
	}

	if v, ok := d.GetOk("stat_log"); ok || d.HasChange("stat_log") {
		t, err := expandFmupdateFgdSettingStatLog(d, v, "stat_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-log"] = t
		}
	}

	if v, ok := d.GetOk("stat_log_interval"); ok || d.HasChange("stat_log_interval") {
		t, err := expandFmupdateFgdSettingStatLogInterval(d, v, "stat_log_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-log-interval"] = t
		}
	}

	if v, ok := d.GetOk("stat_sync_interval"); ok || d.HasChange("stat_sync_interval") {
		t, err := expandFmupdateFgdSettingStatSyncInterval(d, v, "stat_sync_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-sync-interval"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok || d.HasChange("update_interval") {
		t, err := expandFmupdateFgdSettingUpdateInterval(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("update_log"); ok || d.HasChange("update_log") {
		t, err := expandFmupdateFgdSettingUpdateLog(d, v, "update_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-log"] = t
		}
	}

	if v, ok := d.GetOk("wf_cache"); ok || d.HasChange("wf_cache") {
		t, err := expandFmupdateFgdSettingWfCache(d, v, "wf_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-cache"] = t
		}
	}

	if v, ok := d.GetOk("wf_dn_cache_expire_time"); ok || d.HasChange("wf_dn_cache_expire_time") {
		t, err := expandFmupdateFgdSettingWfDnCacheExpireTime(d, v, "wf_dn_cache_expire_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-dn-cache-expire-time"] = t
		}
	}

	if v, ok := d.GetOk("wf_dn_cache_max_number"); ok || d.HasChange("wf_dn_cache_max_number") {
		t, err := expandFmupdateFgdSettingWfDnCacheMaxNumber(d, v, "wf_dn_cache_max_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-dn-cache-max-number"] = t
		}
	}

	if v, ok := d.GetOk("wf_log"); ok || d.HasChange("wf_log") {
		t, err := expandFmupdateFgdSettingWfLog(d, v, "wf_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-log"] = t
		}
	}

	if v, ok := d.GetOk("wf_preload"); ok || d.HasChange("wf_preload") {
		t, err := expandFmupdateFgdSettingWfPreload(d, v, "wf_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-preload"] = t
		}
	}

	return &obj, nil
}
