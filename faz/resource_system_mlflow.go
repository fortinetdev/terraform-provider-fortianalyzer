// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: System Mlflow

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemMlflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMlflowUpdate,
		Read:   resourceSystemMlflowRead,
		Update: resourceSystemMlflowUpdate,
		Delete: resourceSystemMlflowDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"models": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"artifact_max_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"artifact_retention": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"asset_max_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_deploy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"inference_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"model_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"train_data_days": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"train_repeat_at": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"train_schedule_range": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"train_test_days": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"train_test_sensitivity": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemMlflowUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemMlflow(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemMlflow resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMlflow(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemMlflow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemMlflow")

	return resourceSystemMlflowRead(d, m)
}

func resourceSystemMlflowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemMlflow(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMlflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMlflowRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemMlflow(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemMlflow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMlflow(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemMlflow resource from API: %v", err)
	}
	return nil
}

func flattenSystemMlflowModels(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "artifact_max_count"
		if _, ok := i["artifact-max-count"]; ok {
			v := flattenSystemMlflowModelsArtifactMaxCount(i["artifact-max-count"], d, pre_append)
			tmp["artifact_max_count"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-ArtifactMaxCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "artifact_retention"
		if _, ok := i["artifact-retention"]; ok {
			v := flattenSystemMlflowModelsArtifactRetention(i["artifact-retention"], d, pre_append)
			tmp["artifact_retention"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-ArtifactRetention")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "asset_max_count"
		if _, ok := i["asset-max-count"]; ok {
			v := flattenSystemMlflowModelsAssetMaxCount(i["asset-max-count"], d, pre_append)
			tmp["asset_max_count"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-AssetMaxCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auto_deploy"
		if _, ok := i["auto-deploy"]; ok {
			v := flattenSystemMlflowModelsAutoDeploy(i["auto-deploy"], d, pre_append)
			tmp["auto_deploy"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-AutoDeploy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemMlflowModelsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "inference_interval"
		if _, ok := i["inference-interval"]; ok {
			v := flattenSystemMlflowModelsInferenceInterval(i["inference-interval"], d, pre_append)
			tmp["inference_interval"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-InferenceInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "model_type"
		if _, ok := i["model-type"]; ok {
			v := flattenSystemMlflowModelsModelType(i["model-type"], d, pre_append)
			tmp["model_type"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-ModelType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemMlflowModelsStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_data_days"
		if _, ok := i["train-data-days"]; ok {
			v := flattenSystemMlflowModelsTrainDataDays(i["train-data-days"], d, pre_append)
			tmp["train_data_days"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-TrainDataDays")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_repeat_at"
		if _, ok := i["train-repeat-at"]; ok {
			v := flattenSystemMlflowModelsTrainRepeatAt(i["train-repeat-at"], d, pre_append)
			tmp["train_repeat_at"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-TrainRepeatAt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_schedule_range"
		if _, ok := i["train-schedule-range"]; ok {
			v := flattenSystemMlflowModelsTrainScheduleRange(i["train-schedule-range"], d, pre_append)
			tmp["train_schedule_range"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-TrainScheduleRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_test_days"
		if _, ok := i["train-test-days"]; ok {
			v := flattenSystemMlflowModelsTrainTestDays(i["train-test-days"], d, pre_append)
			tmp["train_test_days"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-TrainTestDays")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_test_sensitivity"
		if _, ok := i["train-test-sensitivity"]; ok {
			v := flattenSystemMlflowModelsTrainTestSensitivity(i["train-test-sensitivity"], d, pre_append)
			tmp["train_test_sensitivity"] = fortiAPISubPartPatch(v, "SystemMlflow-Models-TrainTestSensitivity")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemMlflowModelsArtifactMaxCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsArtifactRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsAssetMaxCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsAutoDeploy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsInferenceInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsModelType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainDataDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainRepeatAt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainScheduleRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainTestDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainTestSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemMlflow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("models", flattenSystemMlflowModels(o["models"], d, "models")); err != nil {
			if vv, ok := fortiAPIPatch(o["models"], "SystemMlflow-Models"); ok {
				if err = d.Set("models", vv); err != nil {
					return fmt.Errorf("Error reading models: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading models: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("models"); ok {
			if err = d.Set("models", flattenSystemMlflowModels(o["models"], d, "models")); err != nil {
				if vv, ok := fortiAPIPatch(o["models"], "SystemMlflow-Models"); ok {
					if err = d.Set("models", vv); err != nil {
						return fmt.Errorf("Error reading models: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading models: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenSystemMlflowStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemMlflow-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemMlflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemMlflowModels(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "artifact_max_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["artifact-max-count"], _ = expandSystemMlflowModelsArtifactMaxCount(d, i["artifact_max_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "artifact_retention"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["artifact-retention"], _ = expandSystemMlflowModelsArtifactRetention(d, i["artifact_retention"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "asset_max_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["asset-max-count"], _ = expandSystemMlflowModelsAssetMaxCount(d, i["asset_max_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auto_deploy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auto-deploy"], _ = expandSystemMlflowModelsAutoDeploy(d, i["auto_deploy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemMlflowModelsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "inference_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["inference-interval"], _ = expandSystemMlflowModelsInferenceInterval(d, i["inference_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "model_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["model-type"], _ = expandSystemMlflowModelsModelType(d, i["model_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemMlflowModelsStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_data_days"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["train-data-days"], _ = expandSystemMlflowModelsTrainDataDays(d, i["train_data_days"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_repeat_at"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["train-repeat-at"], _ = expandSystemMlflowModelsTrainRepeatAt(d, i["train_repeat_at"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_schedule_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["train-schedule-range"], _ = expandSystemMlflowModelsTrainScheduleRange(d, i["train_schedule_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_test_days"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["train-test-days"], _ = expandSystemMlflowModelsTrainTestDays(d, i["train_test_days"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "train_test_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["train-test-sensitivity"], _ = expandSystemMlflowModelsTrainTestSensitivity(d, i["train_test_sensitivity"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemMlflowModelsArtifactMaxCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsArtifactRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsAssetMaxCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsAutoDeploy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsInferenceInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsModelType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainDataDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainRepeatAt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainScheduleRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainTestDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainTestSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMlflow(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("models"); ok || d.HasChange("models") {
		t, err := expandSystemMlflowModels(d, v, "models")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["models"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemMlflowStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
