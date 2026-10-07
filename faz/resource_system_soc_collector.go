// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: FAZ SOC collector settings.

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSocCollector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSocCollectorUpdate,
		Read:   resourceSystemSocCollectorRead,
		Update: resourceSystemSocCollectorUpdate,
		Delete: resourceSystemSocCollectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"log_level": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
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
				},
			},
			"log_queue": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_queue_file_count_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"config_queue_size_limit_mb": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"linux_queue_file_count_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"linux_queue_size_limit_mb": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"logs_queue_file_count_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"logs_queue_size_limit_mb": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"win_queue_file_count_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"win_queue_size_limit_mb": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"perf_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"super_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"super_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"verify_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSocCollectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSocCollector(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocCollector resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSocCollector(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocCollector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSocCollector")

	return resourceSystemSocCollectorRead(d, m)
}

func resourceSystemSocCollectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSocCollector(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSocCollector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSocCollectorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSocCollector(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocCollector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSocCollector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocCollector resource from API: %v", err)
	}
	return nil
}

func flattenSystemSocCollectorClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorClientSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSocCollectorLogLevel(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all"
	if _, ok := i["all"]; ok {
		result["all"] = flattenSystemSocCollectorLogLevelAll(i["all"], d, pre_append)
	}

	pre_append = pre + ".0." + "phagentmanager"
	if _, ok := i["phAgentManager"]; ok {
		result["phagentmanager"] = flattenSystemSocCollectorLogLevelPhagentmanager(i["phAgentManager"], d, pre_append)
	}

	pre_append = pre + ".0." + "phdiscover"
	if _, ok := i["phDiscover"]; ok {
		result["phdiscover"] = flattenSystemSocCollectorLogLevelPhdiscover(i["phDiscover"], d, pre_append)
	}

	pre_append = pre + ".0." + "pheventforwarder"
	if _, ok := i["phEventForwarder"]; ok {
		result["pheventforwarder"] = flattenSystemSocCollectorLogLevelPheventforwarder(i["phEventForwarder"], d, pre_append)
	}

	pre_append = pre + ".0." + "pheventhandler"
	if _, ok := i["phEventHandler"]; ok {
		result["pheventhandler"] = flattenSystemSocCollectorLogLevelPheventhandler(i["phEventHandler"], d, pre_append)
	}

	pre_append = pre + ".0." + "pheventpackager"
	if _, ok := i["phEventPackager"]; ok {
		result["pheventpackager"] = flattenSystemSocCollectorLogLevelPheventpackager(i["phEventPackager"], d, pre_append)
	}

	pre_append = pre + ".0." + "phmonitor"
	if _, ok := i["phMonitor"]; ok {
		result["phmonitor"] = flattenSystemSocCollectorLogLevelPhmonitor(i["phMonitor"], d, pre_append)
	}

	pre_append = pre + ".0." + "phparser"
	if _, ok := i["phParser"]; ok {
		result["phparser"] = flattenSystemSocCollectorLogLevelPhparser(i["phParser"], d, pre_append)
	}

	pre_append = pre + ".0." + "phperfmonitor"
	if _, ok := i["phPerfMonitor"]; ok {
		result["phperfmonitor"] = flattenSystemSocCollectorLogLevelPhperfmonitor(i["phPerfMonitor"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemSocCollectorLogLevelAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhagentmanager(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhdiscover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPheventforwarder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPheventhandler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPheventpackager(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhmonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhparser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogLevelPhperfmonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueue(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config_queue_file_count_limit"
	if _, ok := i["config-queue-file-count-limit"]; ok {
		result["config_queue_file_count_limit"] = flattenSystemSocCollectorLogQueueConfigQueueFileCountLimit(i["config-queue-file-count-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "config_queue_size_limit_mb"
	if _, ok := i["config-queue-size-limit-mb"]; ok {
		result["config_queue_size_limit_mb"] = flattenSystemSocCollectorLogQueueConfigQueueSizeLimitMb(i["config-queue-size-limit-mb"], d, pre_append)
	}

	pre_append = pre + ".0." + "linux_queue_file_count_limit"
	if _, ok := i["linux-queue-file-count-limit"]; ok {
		result["linux_queue_file_count_limit"] = flattenSystemSocCollectorLogQueueLinuxQueueFileCountLimit(i["linux-queue-file-count-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "linux_queue_size_limit_mb"
	if _, ok := i["linux-queue-size-limit-mb"]; ok {
		result["linux_queue_size_limit_mb"] = flattenSystemSocCollectorLogQueueLinuxQueueSizeLimitMb(i["linux-queue-size-limit-mb"], d, pre_append)
	}

	pre_append = pre + ".0." + "logs_queue_file_count_limit"
	if _, ok := i["logs-queue-file-count-limit"]; ok {
		result["logs_queue_file_count_limit"] = flattenSystemSocCollectorLogQueueLogsQueueFileCountLimit(i["logs-queue-file-count-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "logs_queue_size_limit_mb"
	if _, ok := i["logs-queue-size-limit-mb"]; ok {
		result["logs_queue_size_limit_mb"] = flattenSystemSocCollectorLogQueueLogsQueueSizeLimitMb(i["logs-queue-size-limit-mb"], d, pre_append)
	}

	pre_append = pre + ".0." + "win_queue_file_count_limit"
	if _, ok := i["win-queue-file-count-limit"]; ok {
		result["win_queue_file_count_limit"] = flattenSystemSocCollectorLogQueueWinQueueFileCountLimit(i["win-queue-file-count-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "win_queue_size_limit_mb"
	if _, ok := i["win-queue-size-limit-mb"]; ok {
		result["win_queue_size_limit_mb"] = flattenSystemSocCollectorLogQueueWinQueueSizeLimitMb(i["win-queue-size-limit-mb"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemSocCollectorLogQueueConfigQueueFileCountLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueConfigQueueSizeLimitMb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLinuxQueueFileCountLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLinuxQueueSizeLimitMb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLogsQueueFileCountLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueLogsQueueSizeLimitMb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueWinQueueFileCountLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorLogQueueWinQueueSizeLimitMb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorOrgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorPerfMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorSuperAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorSuperPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocCollectorVerifySsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSocCollector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("client_id", flattenSystemSocCollectorClientId(o["client-id"], d, "client_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-id"], "SystemSocCollector-ClientId"); ok {
			if err = d.Set("client_id", vv); err != nil {
				return fmt.Errorf("Error reading client_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if err = d.Set("client_secret", flattenSystemSocCollectorClientSecret(o["client-secret"], d, "client_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-secret"], "SystemSocCollector-ClientSecret"); ok {
			if err = d.Set("client_secret", vv); err != nil {
				return fmt.Errorf("Error reading client_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_secret: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("log_level", flattenSystemSocCollectorLogLevel(o["log-level"], d, "log_level")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-level"], "SystemSocCollector-LogLevel"); ok {
				if err = d.Set("log_level", vv); err != nil {
					return fmt.Errorf("Error reading log_level: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_level: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_level"); ok {
			if err = d.Set("log_level", flattenSystemSocCollectorLogLevel(o["log-level"], d, "log_level")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-level"], "SystemSocCollector-LogLevel"); ok {
					if err = d.Set("log_level", vv); err != nil {
						return fmt.Errorf("Error reading log_level: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_level: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("log_queue", flattenSystemSocCollectorLogQueue(o["log-queue"], d, "log_queue")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-queue"], "SystemSocCollector-LogQueue"); ok {
				if err = d.Set("log_queue", vv); err != nil {
					return fmt.Errorf("Error reading log_queue: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_queue: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_queue"); ok {
			if err = d.Set("log_queue", flattenSystemSocCollectorLogQueue(o["log-queue"], d, "log_queue")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-queue"], "SystemSocCollector-LogQueue"); ok {
					if err = d.Set("log_queue", vv); err != nil {
						return fmt.Errorf("Error reading log_queue: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_queue: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemSocCollectorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSocCollector-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("org_name", flattenSystemSocCollectorOrgName(o["org-name"], d, "org_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["org-name"], "SystemSocCollector-OrgName"); ok {
			if err = d.Set("org_name", vv); err != nil {
				return fmt.Errorf("Error reading org_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading org_name: %v", err)
		}
	}

	if err = d.Set("perf_mode", flattenSystemSocCollectorPerfMode(o["perf-mode"], d, "perf_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["perf-mode"], "SystemSocCollector-PerfMode"); ok {
			if err = d.Set("perf_mode", vv); err != nil {
				return fmt.Errorf("Error reading perf_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading perf_mode: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSocCollectorStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSocCollector-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("super_addr", flattenSystemSocCollectorSuperAddr(o["super-addr"], d, "super_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["super-addr"], "SystemSocCollector-SuperAddr"); ok {
			if err = d.Set("super_addr", vv); err != nil {
				return fmt.Errorf("Error reading super_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading super_addr: %v", err)
		}
	}

	if err = d.Set("super_port", flattenSystemSocCollectorSuperPort(o["super-port"], d, "super_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["super-port"], "SystemSocCollector-SuperPort"); ok {
			if err = d.Set("super_port", vv); err != nil {
				return fmt.Errorf("Error reading super_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading super_port: %v", err)
		}
	}

	if err = d.Set("verify_ssl", flattenSystemSocCollectorVerifySsl(o["verify-ssl"], d, "verify_ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-ssl"], "SystemSocCollector-VerifySsl"); ok {
			if err = d.Set("verify_ssl", vv); err != nil {
				return fmt.Errorf("Error reading verify_ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_ssl: %v", err)
		}
	}

	return nil
}

func flattenSystemSocCollectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSocCollectorClientId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorClientSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSocCollectorLogLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["all"], _ = expandSystemSocCollectorLogLevelAll(d, i["all"], pre_append)
	}
	pre_append = pre + ".0." + "phagentmanager"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phAgentManager"], _ = expandSystemSocCollectorLogLevelPhagentmanager(d, i["phagentmanager"], pre_append)
	}
	pre_append = pre + ".0." + "phdiscover"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phDiscover"], _ = expandSystemSocCollectorLogLevelPhdiscover(d, i["phdiscover"], pre_append)
	}
	pre_append = pre + ".0." + "pheventforwarder"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phEventForwarder"], _ = expandSystemSocCollectorLogLevelPheventforwarder(d, i["pheventforwarder"], pre_append)
	}
	pre_append = pre + ".0." + "pheventhandler"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phEventHandler"], _ = expandSystemSocCollectorLogLevelPheventhandler(d, i["pheventhandler"], pre_append)
	}
	pre_append = pre + ".0." + "pheventpackager"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phEventPackager"], _ = expandSystemSocCollectorLogLevelPheventpackager(d, i["pheventpackager"], pre_append)
	}
	pre_append = pre + ".0." + "phmonitor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phMonitor"], _ = expandSystemSocCollectorLogLevelPhmonitor(d, i["phmonitor"], pre_append)
	}
	pre_append = pre + ".0." + "phparser"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phParser"], _ = expandSystemSocCollectorLogLevelPhparser(d, i["phparser"], pre_append)
	}
	pre_append = pre + ".0." + "phperfmonitor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phPerfMonitor"], _ = expandSystemSocCollectorLogLevelPhperfmonitor(d, i["phperfmonitor"], pre_append)
	}

	return result, nil
}

func expandSystemSocCollectorLogLevelAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhagentmanager(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhdiscover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPheventforwarder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPheventhandler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPheventpackager(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhmonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhparser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogLevelPhperfmonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config_queue_file_count_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["config-queue-file-count-limit"], _ = expandSystemSocCollectorLogQueueConfigQueueFileCountLimit(d, i["config_queue_file_count_limit"], pre_append)
	}
	pre_append = pre + ".0." + "config_queue_size_limit_mb"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["config-queue-size-limit-mb"], _ = expandSystemSocCollectorLogQueueConfigQueueSizeLimitMb(d, i["config_queue_size_limit_mb"], pre_append)
	}
	pre_append = pre + ".0." + "linux_queue_file_count_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["linux-queue-file-count-limit"], _ = expandSystemSocCollectorLogQueueLinuxQueueFileCountLimit(d, i["linux_queue_file_count_limit"], pre_append)
	}
	pre_append = pre + ".0." + "linux_queue_size_limit_mb"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["linux-queue-size-limit-mb"], _ = expandSystemSocCollectorLogQueueLinuxQueueSizeLimitMb(d, i["linux_queue_size_limit_mb"], pre_append)
	}
	pre_append = pre + ".0." + "logs_queue_file_count_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["logs-queue-file-count-limit"], _ = expandSystemSocCollectorLogQueueLogsQueueFileCountLimit(d, i["logs_queue_file_count_limit"], pre_append)
	}
	pre_append = pre + ".0." + "logs_queue_size_limit_mb"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["logs-queue-size-limit-mb"], _ = expandSystemSocCollectorLogQueueLogsQueueSizeLimitMb(d, i["logs_queue_size_limit_mb"], pre_append)
	}
	pre_append = pre + ".0." + "win_queue_file_count_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["win-queue-file-count-limit"], _ = expandSystemSocCollectorLogQueueWinQueueFileCountLimit(d, i["win_queue_file_count_limit"], pre_append)
	}
	pre_append = pre + ".0." + "win_queue_size_limit_mb"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["win-queue-size-limit-mb"], _ = expandSystemSocCollectorLogQueueWinQueueSizeLimitMb(d, i["win_queue_size_limit_mb"], pre_append)
	}

	return result, nil
}

func expandSystemSocCollectorLogQueueConfigQueueFileCountLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueConfigQueueSizeLimitMb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLinuxQueueFileCountLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLinuxQueueSizeLimitMb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLogsQueueFileCountLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueLogsQueueSizeLimitMb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueWinQueueFileCountLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorLogQueueWinQueueSizeLimitMb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorOrgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorPerfMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorSuperAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorSuperPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocCollectorVerifySsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSocCollector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("client_id"); ok || d.HasChange("client_id") {
		t, err := expandSystemSocCollectorClientId(d, v, "client_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	}

	if v, ok := d.GetOk("client_secret"); ok || d.HasChange("client_secret") {
		t, err := expandSystemSocCollectorClientSecret(d, v, "client_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	}

	if v, ok := d.GetOk("log_level"); ok || d.HasChange("log_level") {
		t, err := expandSystemSocCollectorLogLevel(d, v, "log_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-level"] = t
		}
	}

	if v, ok := d.GetOk("log_queue"); ok || d.HasChange("log_queue") {
		t, err := expandSystemSocCollectorLogQueue(d, v, "log_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-queue"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSocCollectorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("org_name"); ok || d.HasChange("org_name") {
		t, err := expandSystemSocCollectorOrgName(d, v, "org_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["org-name"] = t
		}
	}

	if v, ok := d.GetOk("perf_mode"); ok || d.HasChange("perf_mode") {
		t, err := expandSystemSocCollectorPerfMode(d, v, "perf_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["perf-mode"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSocCollectorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("super_addr"); ok || d.HasChange("super_addr") {
		t, err := expandSystemSocCollectorSuperAddr(d, v, "super_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["super-addr"] = t
		}
	}

	if v, ok := d.GetOk("super_port"); ok || d.HasChange("super_port") {
		t, err := expandSystemSocCollectorSuperPort(d, v, "super_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["super-port"] = t
		}
	}

	if v, ok := d.GetOk("verify_ssl"); ok || d.HasChange("verify_ssl") {
		t, err := expandSystemSocCollectorVerifySsl(d, v, "verify_ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-ssl"] = t
		}
	}

	return &obj, nil
}
