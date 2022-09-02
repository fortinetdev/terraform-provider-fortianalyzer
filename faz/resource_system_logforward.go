// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Log forwarding.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogForwardCreate,
		Read:   resourceSystemLogForwardRead,
		Update: resourceSystemLogForwardUpdate,
		Delete: resourceSystemLogForwardDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"agg_archive_types": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"agg_data_end_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"agg_data_start_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"agg_logtypes": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"agg_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"agg_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"agg_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"agg_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"fwd_archive_types": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fwd_archives": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_compression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_facility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_ha_bind_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_log_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_max_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_syslog_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"log_field_exclusion": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dev_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"field_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"log_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_field_exclusion_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"oper": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_filter_logic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_filter_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_masking_custom": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"field_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"log_masking_custom_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_masking_fields": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"log_masking_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_masking_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcapurl_domain_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pcapurl_enrich": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_cert_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_service_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sync_metadata": &schema.Schema{
				Type:     schema.TypeSet,
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

func resourceSystemLogForwardCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogForward(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogForward resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLogForward(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemLogForward resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogForwardRead(d, m)
}

func resourceSystemLogForwardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogForward(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogForward resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogForward(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogForward resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogForwardRead(d, m)
}

func resourceSystemLogForwardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogForward(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogForward resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogForwardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogForward(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogForward resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogForward(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogForward resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogForwardAggArchiveTypes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardAggDataEndTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardAggDataStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardAggLogtypes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardAggPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardAggSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardAggTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardAggUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardDeviceFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenSystemLogForwardDeviceFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SystemLogForward-DeviceFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom"
		if _, ok := i["adom"]; ok {
			v := flattenSystemLogForwardDeviceFilterAdom(i["adom"], d, pre_append)
			tmp["adom"] = fortiAPISubPartPatch(v, "SystemLogForward-DeviceFilter-Adom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenSystemLogForwardDeviceFilterDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "SystemLogForward-DeviceFilter-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogForwardDeviceFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogForward-DeviceFilter-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogForwardDeviceFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardDeviceFilterAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardDeviceFilterDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardDeviceFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdArchiveTypes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardFwdArchives(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdCompression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdFacility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdHaBindVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdLogSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdMaxDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardFwdSyslogFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFieldExclusion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dev_type"
		if _, ok := i["dev-type"]; ok {
			v := flattenSystemLogForwardLogFieldExclusionDevType(i["dev-type"], d, pre_append)
			tmp["dev_type"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFieldExclusion-DevType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_list"
		if _, ok := i["field-list"]; ok {
			v := flattenSystemLogForwardLogFieldExclusionFieldList(i["field-list"], d, pre_append)
			tmp["field_list"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFieldExclusion-FieldList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogForwardLogFieldExclusionId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFieldExclusion-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := i["log-type"]; ok {
			v := flattenSystemLogForwardLogFieldExclusionLogType(i["log-type"], d, pre_append)
			tmp["log_type"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFieldExclusion-LogType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogForwardLogFieldExclusionDevType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFieldExclusionFieldList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFieldExclusionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFieldExclusionLogType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFieldExclusionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field"
		if _, ok := i["field"]; ok {
			v := flattenSystemLogForwardLogFilterField(i["field"], d, pre_append)
			tmp["field"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFilter-Field")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogForwardLogFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFilter-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oper"
		if _, ok := i["oper"]; ok {
			v := flattenSystemLogForwardLogFilterOper(i["oper"], d, pre_append)
			tmp["oper"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFilter-Oper")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemLogForwardLogFilterValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemLogForward-LogFilter-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogForwardLogFilterField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFilterOper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFilterValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFilterLogic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogFilterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogMaskingCustom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_name"
		if _, ok := i["field-name"]; ok {
			v := flattenSystemLogForwardLogMaskingCustomFieldName(i["field-name"], d, pre_append)
			tmp["field_name"] = fortiAPISubPartPatch(v, "SystemLogForward-LogMaskingCustom-FieldName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_type"
		if _, ok := i["field-type"]; ok {
			v := flattenSystemLogForwardLogMaskingCustomFieldType(i["field-type"], d, pre_append)
			tmp["field_type"] = fortiAPISubPartPatch(v, "SystemLogForward-LogMaskingCustom-FieldType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogForwardLogMaskingCustomId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogForward-LogMaskingCustom-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogForwardLogMaskingCustomFieldName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogMaskingCustomFieldType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogMaskingCustomId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogMaskingCustomPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardLogMaskingFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardLogMaskingKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogForwardLogMaskingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardPcapurlDomainIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardPcapurlEnrich(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardPeerCertCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardProxyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardProxyServicePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardServerAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardServerDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogForwardSyncMetadata(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemLogForward(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("agg_archive_types", flattenSystemLogForwardAggArchiveTypes(o["agg-archive-types"], d, "agg_archive_types")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-archive-types"], "SystemLogForward-AggArchiveTypes"); ok {
			if err = d.Set("agg_archive_types", vv); err != nil {
				return fmt.Errorf("Error reading agg_archive_types: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_archive_types: %v", err)
		}
	}

	if err = d.Set("agg_data_end_time", flattenSystemLogForwardAggDataEndTime(o["agg-data-end-time"], d, "agg_data_end_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-data-end-time"], "SystemLogForward-AggDataEndTime"); ok {
			if err = d.Set("agg_data_end_time", vv); err != nil {
				return fmt.Errorf("Error reading agg_data_end_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_data_end_time: %v", err)
		}
	}

	if err = d.Set("agg_data_start_time", flattenSystemLogForwardAggDataStartTime(o["agg-data-start-time"], d, "agg_data_start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-data-start-time"], "SystemLogForward-AggDataStartTime"); ok {
			if err = d.Set("agg_data_start_time", vv); err != nil {
				return fmt.Errorf("Error reading agg_data_start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_data_start_time: %v", err)
		}
	}

	if err = d.Set("agg_logtypes", flattenSystemLogForwardAggLogtypes(o["agg-logtypes"], d, "agg_logtypes")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-logtypes"], "SystemLogForward-AggLogtypes"); ok {
			if err = d.Set("agg_logtypes", vv); err != nil {
				return fmt.Errorf("Error reading agg_logtypes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_logtypes: %v", err)
		}
	}

	if err = d.Set("agg_password", flattenSystemLogForwardAggPassword(o["agg-password"], d, "agg_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-password"], "SystemLogForward-AggPassword"); ok {
			if err = d.Set("agg_password", vv); err != nil {
				return fmt.Errorf("Error reading agg_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_password: %v", err)
		}
	}

	if err = d.Set("agg_schedule", flattenSystemLogForwardAggSchedule(o["agg-schedule"], d, "agg_schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-schedule"], "SystemLogForward-AggSchedule"); ok {
			if err = d.Set("agg_schedule", vv); err != nil {
				return fmt.Errorf("Error reading agg_schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_schedule: %v", err)
		}
	}

	if err = d.Set("agg_time", flattenSystemLogForwardAggTime(o["agg-time"], d, "agg_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-time"], "SystemLogForward-AggTime"); ok {
			if err = d.Set("agg_time", vv); err != nil {
				return fmt.Errorf("Error reading agg_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_time: %v", err)
		}
	}

	if err = d.Set("agg_user", flattenSystemLogForwardAggUser(o["agg-user"], d, "agg_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["agg-user"], "SystemLogForward-AggUser"); ok {
			if err = d.Set("agg_user", vv); err != nil {
				return fmt.Errorf("Error reading agg_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agg_user: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("device_filter", flattenSystemLogForwardDeviceFilter(o["device-filter"], d, "device_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["device-filter"], "SystemLogForward-DeviceFilter"); ok {
				if err = d.Set("device_filter", vv); err != nil {
					return fmt.Errorf("Error reading device_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_filter"); ok {
			if err = d.Set("device_filter", flattenSystemLogForwardDeviceFilter(o["device-filter"], d, "device_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["device-filter"], "SystemLogForward-DeviceFilter"); ok {
					if err = d.Set("device_filter", vv); err != nil {
						return fmt.Errorf("Error reading device_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("fwd_archive_types", flattenSystemLogForwardFwdArchiveTypes(o["fwd-archive-types"], d, "fwd_archive_types")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-archive-types"], "SystemLogForward-FwdArchiveTypes"); ok {
			if err = d.Set("fwd_archive_types", vv); err != nil {
				return fmt.Errorf("Error reading fwd_archive_types: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_archive_types: %v", err)
		}
	}

	if err = d.Set("fwd_archives", flattenSystemLogForwardFwdArchives(o["fwd-archives"], d, "fwd_archives")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-archives"], "SystemLogForward-FwdArchives"); ok {
			if err = d.Set("fwd_archives", vv); err != nil {
				return fmt.Errorf("Error reading fwd_archives: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_archives: %v", err)
		}
	}

	if err = d.Set("fwd_compression", flattenSystemLogForwardFwdCompression(o["fwd-compression"], d, "fwd_compression")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-compression"], "SystemLogForward-FwdCompression"); ok {
			if err = d.Set("fwd_compression", vv); err != nil {
				return fmt.Errorf("Error reading fwd_compression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_compression: %v", err)
		}
	}

	if err = d.Set("fwd_facility", flattenSystemLogForwardFwdFacility(o["fwd-facility"], d, "fwd_facility")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-facility"], "SystemLogForward-FwdFacility"); ok {
			if err = d.Set("fwd_facility", vv); err != nil {
				return fmt.Errorf("Error reading fwd_facility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_facility: %v", err)
		}
	}

	if err = d.Set("fwd_ha_bind_vip", flattenSystemLogForwardFwdHaBindVip(o["fwd-ha-bind-vip"], d, "fwd_ha_bind_vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-ha-bind-vip"], "SystemLogForward-FwdHaBindVip"); ok {
			if err = d.Set("fwd_ha_bind_vip", vv); err != nil {
				return fmt.Errorf("Error reading fwd_ha_bind_vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_ha_bind_vip: %v", err)
		}
	}

	if err = d.Set("fwd_log_source_ip", flattenSystemLogForwardFwdLogSourceIp(o["fwd-log-source-ip"], d, "fwd_log_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-log-source-ip"], "SystemLogForward-FwdLogSourceIp"); ok {
			if err = d.Set("fwd_log_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading fwd_log_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_log_source_ip: %v", err)
		}
	}

	if err = d.Set("fwd_max_delay", flattenSystemLogForwardFwdMaxDelay(o["fwd-max-delay"], d, "fwd_max_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-max-delay"], "SystemLogForward-FwdMaxDelay"); ok {
			if err = d.Set("fwd_max_delay", vv); err != nil {
				return fmt.Errorf("Error reading fwd_max_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_max_delay: %v", err)
		}
	}

	if err = d.Set("fwd_reliable", flattenSystemLogForwardFwdReliable(o["fwd-reliable"], d, "fwd_reliable")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-reliable"], "SystemLogForward-FwdReliable"); ok {
			if err = d.Set("fwd_reliable", vv); err != nil {
				return fmt.Errorf("Error reading fwd_reliable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_reliable: %v", err)
		}
	}

	if err = d.Set("fwd_secure", flattenSystemLogForwardFwdSecure(o["fwd-secure"], d, "fwd_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-secure"], "SystemLogForward-FwdSecure"); ok {
			if err = d.Set("fwd_secure", vv); err != nil {
				return fmt.Errorf("Error reading fwd_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_secure: %v", err)
		}
	}

	if err = d.Set("fwd_server_type", flattenSystemLogForwardFwdServerType(o["fwd-server-type"], d, "fwd_server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-server-type"], "SystemLogForward-FwdServerType"); ok {
			if err = d.Set("fwd_server_type", vv); err != nil {
				return fmt.Errorf("Error reading fwd_server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_server_type: %v", err)
		}
	}

	if err = d.Set("fwd_syslog_format", flattenSystemLogForwardFwdSyslogFormat(o["fwd-syslog-format"], d, "fwd_syslog_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-syslog-format"], "SystemLogForward-FwdSyslogFormat"); ok {
			if err = d.Set("fwd_syslog_format", vv); err != nil {
				return fmt.Errorf("Error reading fwd_syslog_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_syslog_format: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLogForwardId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLogForward-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("log_field_exclusion", flattenSystemLogForwardLogFieldExclusion(o["log-field-exclusion"], d, "log_field_exclusion")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-field-exclusion"], "SystemLogForward-LogFieldExclusion"); ok {
				if err = d.Set("log_field_exclusion", vv); err != nil {
					return fmt.Errorf("Error reading log_field_exclusion: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_field_exclusion: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_field_exclusion"); ok {
			if err = d.Set("log_field_exclusion", flattenSystemLogForwardLogFieldExclusion(o["log-field-exclusion"], d, "log_field_exclusion")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-field-exclusion"], "SystemLogForward-LogFieldExclusion"); ok {
					if err = d.Set("log_field_exclusion", vv); err != nil {
						return fmt.Errorf("Error reading log_field_exclusion: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_field_exclusion: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_field_exclusion_status", flattenSystemLogForwardLogFieldExclusionStatus(o["log-field-exclusion-status"], d, "log_field_exclusion_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-field-exclusion-status"], "SystemLogForward-LogFieldExclusionStatus"); ok {
			if err = d.Set("log_field_exclusion_status", vv); err != nil {
				return fmt.Errorf("Error reading log_field_exclusion_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_field_exclusion_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("log_filter", flattenSystemLogForwardLogFilter(o["log-filter"], d, "log_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-filter"], "SystemLogForward-LogFilter"); ok {
				if err = d.Set("log_filter", vv); err != nil {
					return fmt.Errorf("Error reading log_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_filter"); ok {
			if err = d.Set("log_filter", flattenSystemLogForwardLogFilter(o["log-filter"], d, "log_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-filter"], "SystemLogForward-LogFilter"); ok {
					if err = d.Set("log_filter", vv); err != nil {
						return fmt.Errorf("Error reading log_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_filter_logic", flattenSystemLogForwardLogFilterLogic(o["log-filter-logic"], d, "log_filter_logic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-filter-logic"], "SystemLogForward-LogFilterLogic"); ok {
			if err = d.Set("log_filter_logic", vv); err != nil {
				return fmt.Errorf("Error reading log_filter_logic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_filter_logic: %v", err)
		}
	}

	if err = d.Set("log_filter_status", flattenSystemLogForwardLogFilterStatus(o["log-filter-status"], d, "log_filter_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-filter-status"], "SystemLogForward-LogFilterStatus"); ok {
			if err = d.Set("log_filter_status", vv); err != nil {
				return fmt.Errorf("Error reading log_filter_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_filter_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("log_masking_custom", flattenSystemLogForwardLogMaskingCustom(o["log-masking-custom"], d, "log_masking_custom")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-masking-custom"], "SystemLogForward-LogMaskingCustom"); ok {
				if err = d.Set("log_masking_custom", vv); err != nil {
					return fmt.Errorf("Error reading log_masking_custom: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_masking_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_masking_custom"); ok {
			if err = d.Set("log_masking_custom", flattenSystemLogForwardLogMaskingCustom(o["log-masking-custom"], d, "log_masking_custom")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-masking-custom"], "SystemLogForward-LogMaskingCustom"); ok {
					if err = d.Set("log_masking_custom", vv); err != nil {
						return fmt.Errorf("Error reading log_masking_custom: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_masking_custom: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_masking_custom_priority", flattenSystemLogForwardLogMaskingCustomPriority(o["log-masking-custom-priority"], d, "log_masking_custom_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-masking-custom-priority"], "SystemLogForward-LogMaskingCustomPriority"); ok {
			if err = d.Set("log_masking_custom_priority", vv); err != nil {
				return fmt.Errorf("Error reading log_masking_custom_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_masking_custom_priority: %v", err)
		}
	}

	if err = d.Set("log_masking_fields", flattenSystemLogForwardLogMaskingFields(o["log-masking-fields"], d, "log_masking_fields")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-masking-fields"], "SystemLogForward-LogMaskingFields"); ok {
			if err = d.Set("log_masking_fields", vv); err != nil {
				return fmt.Errorf("Error reading log_masking_fields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_masking_fields: %v", err)
		}
	}

	if err = d.Set("log_masking_key", flattenSystemLogForwardLogMaskingKey(o["log-masking-key"], d, "log_masking_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-masking-key"], "SystemLogForward-LogMaskingKey"); ok {
			if err = d.Set("log_masking_key", vv); err != nil {
				return fmt.Errorf("Error reading log_masking_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_masking_key: %v", err)
		}
	}

	if err = d.Set("log_masking_status", flattenSystemLogForwardLogMaskingStatus(o["log-masking-status"], d, "log_masking_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-masking-status"], "SystemLogForward-LogMaskingStatus"); ok {
			if err = d.Set("log_masking_status", vv); err != nil {
				return fmt.Errorf("Error reading log_masking_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_masking_status: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemLogForwardMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemLogForward-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("pcapurl_domain_ip", flattenSystemLogForwardPcapurlDomainIp(o["pcapurl-domain-ip"], d, "pcapurl_domain_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["pcapurl-domain-ip"], "SystemLogForward-PcapurlDomainIp"); ok {
			if err = d.Set("pcapurl_domain_ip", vv); err != nil {
				return fmt.Errorf("Error reading pcapurl_domain_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pcapurl_domain_ip: %v", err)
		}
	}

	if err = d.Set("pcapurl_enrich", flattenSystemLogForwardPcapurlEnrich(o["pcapurl-enrich"], d, "pcapurl_enrich")); err != nil {
		if vv, ok := fortiAPIPatch(o["pcapurl-enrich"], "SystemLogForward-PcapurlEnrich"); ok {
			if err = d.Set("pcapurl_enrich", vv); err != nil {
				return fmt.Errorf("Error reading pcapurl_enrich: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pcapurl_enrich: %v", err)
		}
	}

	if err = d.Set("peer_cert_cn", flattenSystemLogForwardPeerCertCn(o["peer-cert-cn"], d, "peer_cert_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-cert-cn"], "SystemLogForward-PeerCertCn"); ok {
			if err = d.Set("peer_cert_cn", vv); err != nil {
				return fmt.Errorf("Error reading peer_cert_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_cert_cn: %v", err)
		}
	}

	if err = d.Set("proxy_service", flattenSystemLogForwardProxyService(o["proxy-service"], d, "proxy_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-service"], "SystemLogForward-ProxyService"); ok {
			if err = d.Set("proxy_service", vv); err != nil {
				return fmt.Errorf("Error reading proxy_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_service: %v", err)
		}
	}

	if err = d.Set("proxy_service_priority", flattenSystemLogForwardProxyServicePriority(o["proxy-service-priority"], d, "proxy_service_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-service-priority"], "SystemLogForward-ProxyServicePriority"); ok {
			if err = d.Set("proxy_service_priority", vv); err != nil {
				return fmt.Errorf("Error reading proxy_service_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_service_priority: %v", err)
		}
	}

	if err = d.Set("server_addr", flattenSystemLogForwardServerAddr(o["server-addr"], d, "server_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-addr"], "SystemLogForward-ServerAddr"); ok {
			if err = d.Set("server_addr", vv); err != nil {
				return fmt.Errorf("Error reading server_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_addr: %v", err)
		}
	}

	if err = d.Set("server_device", flattenSystemLogForwardServerDevice(o["server-device"], d, "server_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-device"], "SystemLogForward-ServerDevice"); ok {
			if err = d.Set("server_device", vv); err != nil {
				return fmt.Errorf("Error reading server_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_device: %v", err)
		}
	}

	if err = d.Set("server_name", flattenSystemLogForwardServerName(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "SystemLogForward-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("server_port", flattenSystemLogForwardServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "SystemLogForward-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("signature", flattenSystemLogForwardSignature(o["signature"], d, "signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature"], "SystemLogForward-Signature"); ok {
			if err = d.Set("signature", vv); err != nil {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("sync_metadata", flattenSystemLogForwardSyncMetadata(o["sync-metadata"], d, "sync_metadata")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-metadata"], "SystemLogForward-SyncMetadata"); ok {
			if err = d.Set("sync_metadata", vv); err != nil {
				return fmt.Errorf("Error reading sync_metadata: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_metadata: %v", err)
		}
	}

	return nil
}

func flattenSystemLogForwardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogForwardAggArchiveTypes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardAggDataEndTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardAggDataStartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardAggLogtypes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardAggPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardAggSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardAggTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardAggUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardDeviceFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandSystemLogForwardDeviceFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adom"], _ = expandSystemLogForwardDeviceFilterAdom(d, i["adom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device"], _ = expandSystemLogForwardDeviceFilterDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogForwardDeviceFilterId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogForwardDeviceFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardDeviceFilterAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardDeviceFilterDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardDeviceFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdArchiveTypes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardFwdArchives(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdCompression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdFacility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdHaBindVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdLogSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdMaxDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardFwdSyslogFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFieldExclusion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dev_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dev-type"], _ = expandSystemLogForwardLogFieldExclusionDevType(d, i["dev_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field-list"], _ = expandSystemLogForwardLogFieldExclusionFieldList(d, i["field_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogForwardLogFieldExclusionId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-type"], _ = expandSystemLogForwardLogFieldExclusionLogType(d, i["log_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogForwardLogFieldExclusionDevType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFieldExclusionFieldList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFieldExclusionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFieldExclusionLogType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFieldExclusionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field"], _ = expandSystemLogForwardLogFilterField(d, i["field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogForwardLogFilterId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oper"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oper"], _ = expandSystemLogForwardLogFilterOper(d, i["oper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemLogForwardLogFilterValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogForwardLogFilterField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFilterOper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFilterValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFilterLogic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogFilterStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogMaskingCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field-name"], _ = expandSystemLogForwardLogMaskingCustomFieldName(d, i["field_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field-type"], _ = expandSystemLogForwardLogMaskingCustomFieldType(d, i["field_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogForwardLogMaskingCustomId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogForwardLogMaskingCustomFieldName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogMaskingCustomFieldType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogMaskingCustomId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogMaskingCustomPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardLogMaskingFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardLogMaskingKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogForwardLogMaskingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardPcapurlDomainIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardPcapurlEnrich(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardPeerCertCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardProxyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardProxyServicePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardServerAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardServerDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogForwardSyncMetadata(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemLogForward(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("agg_archive_types"); ok || d.HasChange("agg_archive_types") {
		t, err := expandSystemLogForwardAggArchiveTypes(d, v, "agg_archive_types")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-archive-types"] = t
		}
	}

	if v, ok := d.GetOk("agg_data_end_time"); ok || d.HasChange("agg_data_end_time") {
		t, err := expandSystemLogForwardAggDataEndTime(d, v, "agg_data_end_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-data-end-time"] = t
		}
	}

	if v, ok := d.GetOk("agg_data_start_time"); ok || d.HasChange("agg_data_start_time") {
		t, err := expandSystemLogForwardAggDataStartTime(d, v, "agg_data_start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-data-start-time"] = t
		}
	}

	if v, ok := d.GetOk("agg_logtypes"); ok || d.HasChange("agg_logtypes") {
		t, err := expandSystemLogForwardAggLogtypes(d, v, "agg_logtypes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-logtypes"] = t
		}
	}

	if v, ok := d.GetOk("agg_password"); ok || d.HasChange("agg_password") {
		t, err := expandSystemLogForwardAggPassword(d, v, "agg_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-password"] = t
		}
	}

	if v, ok := d.GetOk("agg_schedule"); ok || d.HasChange("agg_schedule") {
		t, err := expandSystemLogForwardAggSchedule(d, v, "agg_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-schedule"] = t
		}
	}

	if v, ok := d.GetOk("agg_time"); ok || d.HasChange("agg_time") {
		t, err := expandSystemLogForwardAggTime(d, v, "agg_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-time"] = t
		}
	}

	if v, ok := d.GetOk("agg_user"); ok || d.HasChange("agg_user") {
		t, err := expandSystemLogForwardAggUser(d, v, "agg_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agg-user"] = t
		}
	}

	if v, ok := d.GetOk("device_filter"); ok || d.HasChange("device_filter") {
		t, err := expandSystemLogForwardDeviceFilter(d, v, "device_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-filter"] = t
		}
	}

	if v, ok := d.GetOk("fwd_archive_types"); ok || d.HasChange("fwd_archive_types") {
		t, err := expandSystemLogForwardFwdArchiveTypes(d, v, "fwd_archive_types")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-archive-types"] = t
		}
	}

	if v, ok := d.GetOk("fwd_archives"); ok || d.HasChange("fwd_archives") {
		t, err := expandSystemLogForwardFwdArchives(d, v, "fwd_archives")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-archives"] = t
		}
	}

	if v, ok := d.GetOk("fwd_compression"); ok || d.HasChange("fwd_compression") {
		t, err := expandSystemLogForwardFwdCompression(d, v, "fwd_compression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-compression"] = t
		}
	}

	if v, ok := d.GetOk("fwd_facility"); ok || d.HasChange("fwd_facility") {
		t, err := expandSystemLogForwardFwdFacility(d, v, "fwd_facility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-facility"] = t
		}
	}

	if v, ok := d.GetOk("fwd_ha_bind_vip"); ok || d.HasChange("fwd_ha_bind_vip") {
		t, err := expandSystemLogForwardFwdHaBindVip(d, v, "fwd_ha_bind_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-ha-bind-vip"] = t
		}
	}

	if v, ok := d.GetOk("fwd_log_source_ip"); ok || d.HasChange("fwd_log_source_ip") {
		t, err := expandSystemLogForwardFwdLogSourceIp(d, v, "fwd_log_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-log-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("fwd_max_delay"); ok || d.HasChange("fwd_max_delay") {
		t, err := expandSystemLogForwardFwdMaxDelay(d, v, "fwd_max_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-max-delay"] = t
		}
	}

	if v, ok := d.GetOk("fwd_reliable"); ok || d.HasChange("fwd_reliable") {
		t, err := expandSystemLogForwardFwdReliable(d, v, "fwd_reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-reliable"] = t
		}
	}

	if v, ok := d.GetOk("fwd_secure"); ok || d.HasChange("fwd_secure") {
		t, err := expandSystemLogForwardFwdSecure(d, v, "fwd_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-secure"] = t
		}
	}

	if v, ok := d.GetOk("fwd_server_type"); ok || d.HasChange("fwd_server_type") {
		t, err := expandSystemLogForwardFwdServerType(d, v, "fwd_server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-server-type"] = t
		}
	}

	if v, ok := d.GetOk("fwd_syslog_format"); ok || d.HasChange("fwd_syslog_format") {
		t, err := expandSystemLogForwardFwdSyslogFormat(d, v, "fwd_syslog_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-syslog-format"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandSystemLogForwardId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log_field_exclusion"); ok || d.HasChange("log_field_exclusion") {
		t, err := expandSystemLogForwardLogFieldExclusion(d, v, "log_field_exclusion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-field-exclusion"] = t
		}
	}

	if v, ok := d.GetOk("log_field_exclusion_status"); ok || d.HasChange("log_field_exclusion_status") {
		t, err := expandSystemLogForwardLogFieldExclusionStatus(d, v, "log_field_exclusion_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-field-exclusion-status"] = t
		}
	}

	if v, ok := d.GetOk("log_filter"); ok || d.HasChange("log_filter") {
		t, err := expandSystemLogForwardLogFilter(d, v, "log_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-filter"] = t
		}
	}

	if v, ok := d.GetOk("log_filter_logic"); ok || d.HasChange("log_filter_logic") {
		t, err := expandSystemLogForwardLogFilterLogic(d, v, "log_filter_logic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-filter-logic"] = t
		}
	}

	if v, ok := d.GetOk("log_filter_status"); ok || d.HasChange("log_filter_status") {
		t, err := expandSystemLogForwardLogFilterStatus(d, v, "log_filter_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-filter-status"] = t
		}
	}

	if v, ok := d.GetOk("log_masking_custom"); ok || d.HasChange("log_masking_custom") {
		t, err := expandSystemLogForwardLogMaskingCustom(d, v, "log_masking_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-masking-custom"] = t
		}
	}

	if v, ok := d.GetOk("log_masking_custom_priority"); ok || d.HasChange("log_masking_custom_priority") {
		t, err := expandSystemLogForwardLogMaskingCustomPriority(d, v, "log_masking_custom_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-masking-custom-priority"] = t
		}
	}

	if v, ok := d.GetOk("log_masking_fields"); ok || d.HasChange("log_masking_fields") {
		t, err := expandSystemLogForwardLogMaskingFields(d, v, "log_masking_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-masking-fields"] = t
		}
	}

	if v, ok := d.GetOk("log_masking_key"); ok || d.HasChange("log_masking_key") {
		t, err := expandSystemLogForwardLogMaskingKey(d, v, "log_masking_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-masking-key"] = t
		}
	}

	if v, ok := d.GetOk("log_masking_status"); ok || d.HasChange("log_masking_status") {
		t, err := expandSystemLogForwardLogMaskingStatus(d, v, "log_masking_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-masking-status"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemLogForwardMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("pcapurl_domain_ip"); ok || d.HasChange("pcapurl_domain_ip") {
		t, err := expandSystemLogForwardPcapurlDomainIp(d, v, "pcapurl_domain_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcapurl-domain-ip"] = t
		}
	}

	if v, ok := d.GetOk("pcapurl_enrich"); ok || d.HasChange("pcapurl_enrich") {
		t, err := expandSystemLogForwardPcapurlEnrich(d, v, "pcapurl_enrich")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcapurl-enrich"] = t
		}
	}

	if v, ok := d.GetOk("peer_cert_cn"); ok || d.HasChange("peer_cert_cn") {
		t, err := expandSystemLogForwardPeerCertCn(d, v, "peer_cert_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-cert-cn"] = t
		}
	}

	if v, ok := d.GetOk("proxy_service"); ok || d.HasChange("proxy_service") {
		t, err := expandSystemLogForwardProxyService(d, v, "proxy_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-service"] = t
		}
	}

	if v, ok := d.GetOk("proxy_service_priority"); ok || d.HasChange("proxy_service_priority") {
		t, err := expandSystemLogForwardProxyServicePriority(d, v, "proxy_service_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-service-priority"] = t
		}
	}

	if v, ok := d.GetOk("server_addr"); ok || d.HasChange("server_addr") {
		t, err := expandSystemLogForwardServerAddr(d, v, "server_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-addr"] = t
		}
	}

	if v, ok := d.GetOk("server_device"); ok || d.HasChange("server_device") {
		t, err := expandSystemLogForwardServerDevice(d, v, "server_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-device"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandSystemLogForwardServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok || d.HasChange("server_port") {
		t, err := expandSystemLogForwardServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandSystemLogForwardSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("sync_metadata"); ok || d.HasChange("sync_metadata") {
		t, err := expandSystemLogForwardSyncMetadata(d, v, "sync_metadata")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-metadata"] = t
		}
	}

	return &obj, nil
}
