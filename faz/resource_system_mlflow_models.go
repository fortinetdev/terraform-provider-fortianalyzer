// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: System MlflowModels

package fortianalyzer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemMlflowModels() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMlflowModelsCreate,
		Read:   resourceSystemMlflowModelsRead,
		Update: resourceSystemMlflowModelsUpdate,
		Delete: resourceSystemMlflowModelsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceSystemMlflowModelsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemMlflowModels(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemMlflowModels resource while getting object: %v", err)
	}

	_, err = c.CreateSystemMlflowModels(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemMlflowModels resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemMlflowModelsRead(d, m)
}

func resourceSystemMlflowModelsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemMlflowModels(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemMlflowModels resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMlflowModels(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemMlflowModels resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemMlflowModelsRead(d, m)
}

func resourceSystemMlflowModelsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemMlflowModels(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMlflowModels resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMlflowModelsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemMlflowModels(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemMlflowModels resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMlflowModels(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemMlflowModels resource from API: %v", err)
	}
	return nil
}

func flattenSystemMlflowModelsArtifactMaxCountSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsArtifactRetentionSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsAssetMaxCountSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsAutoDeploySmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsIdSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsInferenceIntervalSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsModelTypeSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsStatusSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainDataDaysSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainRepeatAtSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainScheduleRangeSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainTestDaysSmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMlflowModelsTrainTestSensitivitySmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemMlflowModels(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("artifact_max_count", flattenSystemMlflowModelsArtifactMaxCountSmm(o["artifact-max-count"], d, "artifact_max_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["artifact-max-count"], "SystemMlflowModels-ArtifactMaxCount"); ok {
			if err = d.Set("artifact_max_count", vv); err != nil {
				return fmt.Errorf("Error reading artifact_max_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading artifact_max_count: %v", err)
		}
	}

	if err = d.Set("artifact_retention", flattenSystemMlflowModelsArtifactRetentionSmm(o["artifact-retention"], d, "artifact_retention")); err != nil {
		if vv, ok := fortiAPIPatch(o["artifact-retention"], "SystemMlflowModels-ArtifactRetention"); ok {
			if err = d.Set("artifact_retention", vv); err != nil {
				return fmt.Errorf("Error reading artifact_retention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading artifact_retention: %v", err)
		}
	}

	if err = d.Set("asset_max_count", flattenSystemMlflowModelsAssetMaxCountSmm(o["asset-max-count"], d, "asset_max_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["asset-max-count"], "SystemMlflowModels-AssetMaxCount"); ok {
			if err = d.Set("asset_max_count", vv); err != nil {
				return fmt.Errorf("Error reading asset_max_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asset_max_count: %v", err)
		}
	}

	if err = d.Set("auto_deploy", flattenSystemMlflowModelsAutoDeploySmm(o["auto-deploy"], d, "auto_deploy")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-deploy"], "SystemMlflowModels-AutoDeploy"); ok {
			if err = d.Set("auto_deploy", vv); err != nil {
				return fmt.Errorf("Error reading auto_deploy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_deploy: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemMlflowModelsIdSmm(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemMlflowModels-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("inference_interval", flattenSystemMlflowModelsInferenceIntervalSmm(o["inference-interval"], d, "inference_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["inference-interval"], "SystemMlflowModels-InferenceInterval"); ok {
			if err = d.Set("inference_interval", vv); err != nil {
				return fmt.Errorf("Error reading inference_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inference_interval: %v", err)
		}
	}

	if err = d.Set("model_type", flattenSystemMlflowModelsModelTypeSmm(o["model-type"], d, "model_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["model-type"], "SystemMlflowModels-ModelType"); ok {
			if err = d.Set("model_type", vv); err != nil {
				return fmt.Errorf("Error reading model_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading model_type: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemMlflowModelsStatusSmm(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemMlflowModels-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("train_data_days", flattenSystemMlflowModelsTrainDataDaysSmm(o["train-data-days"], d, "train_data_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["train-data-days"], "SystemMlflowModels-TrainDataDays"); ok {
			if err = d.Set("train_data_days", vv); err != nil {
				return fmt.Errorf("Error reading train_data_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading train_data_days: %v", err)
		}
	}

	if err = d.Set("train_repeat_at", flattenSystemMlflowModelsTrainRepeatAtSmm(o["train-repeat-at"], d, "train_repeat_at")); err != nil {
		if vv, ok := fortiAPIPatch(o["train-repeat-at"], "SystemMlflowModels-TrainRepeatAt"); ok {
			if err = d.Set("train_repeat_at", vv); err != nil {
				return fmt.Errorf("Error reading train_repeat_at: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading train_repeat_at: %v", err)
		}
	}

	if err = d.Set("train_schedule_range", flattenSystemMlflowModelsTrainScheduleRangeSmm(o["train-schedule-range"], d, "train_schedule_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["train-schedule-range"], "SystemMlflowModels-TrainScheduleRange"); ok {
			if err = d.Set("train_schedule_range", vv); err != nil {
				return fmt.Errorf("Error reading train_schedule_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading train_schedule_range: %v", err)
		}
	}

	if err = d.Set("train_test_days", flattenSystemMlflowModelsTrainTestDaysSmm(o["train-test-days"], d, "train_test_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["train-test-days"], "SystemMlflowModels-TrainTestDays"); ok {
			if err = d.Set("train_test_days", vv); err != nil {
				return fmt.Errorf("Error reading train_test_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading train_test_days: %v", err)
		}
	}

	if err = d.Set("train_test_sensitivity", flattenSystemMlflowModelsTrainTestSensitivitySmm(o["train-test-sensitivity"], d, "train_test_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["train-test-sensitivity"], "SystemMlflowModels-TrainTestSensitivity"); ok {
			if err = d.Set("train_test_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading train_test_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading train_test_sensitivity: %v", err)
		}
	}

	return nil
}

func flattenSystemMlflowModelsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemMlflowModelsArtifactMaxCountSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsArtifactRetentionSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsAssetMaxCountSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsAutoDeploySmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsIdSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsInferenceIntervalSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsModelTypeSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsStatusSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainDataDaysSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainRepeatAtSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainScheduleRangeSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainTestDaysSmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMlflowModelsTrainTestSensitivitySmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMlflowModels(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("artifact_max_count"); ok || d.HasChange("artifact_max_count") {
		t, err := expandSystemMlflowModelsArtifactMaxCountSmm(d, v, "artifact_max_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["artifact-max-count"] = t
		}
	}

	if v, ok := d.GetOk("artifact_retention"); ok || d.HasChange("artifact_retention") {
		t, err := expandSystemMlflowModelsArtifactRetentionSmm(d, v, "artifact_retention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["artifact-retention"] = t
		}
	}

	if v, ok := d.GetOk("asset_max_count"); ok || d.HasChange("asset_max_count") {
		t, err := expandSystemMlflowModelsAssetMaxCountSmm(d, v, "asset_max_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asset-max-count"] = t
		}
	}

	if v, ok := d.GetOk("auto_deploy"); ok || d.HasChange("auto_deploy") {
		t, err := expandSystemMlflowModelsAutoDeploySmm(d, v, "auto_deploy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-deploy"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemMlflowModelsIdSmm(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("inference_interval"); ok || d.HasChange("inference_interval") {
		t, err := expandSystemMlflowModelsInferenceIntervalSmm(d, v, "inference_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inference-interval"] = t
		}
	}

	if v, ok := d.GetOk("model_type"); ok || d.HasChange("model_type") {
		t, err := expandSystemMlflowModelsModelTypeSmm(d, v, "model_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemMlflowModelsStatusSmm(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("train_data_days"); ok || d.HasChange("train_data_days") {
		t, err := expandSystemMlflowModelsTrainDataDaysSmm(d, v, "train_data_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["train-data-days"] = t
		}
	}

	if v, ok := d.GetOk("train_repeat_at"); ok || d.HasChange("train_repeat_at") {
		t, err := expandSystemMlflowModelsTrainRepeatAtSmm(d, v, "train_repeat_at")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["train-repeat-at"] = t
		}
	}

	if v, ok := d.GetOk("train_schedule_range"); ok || d.HasChange("train_schedule_range") {
		t, err := expandSystemMlflowModelsTrainScheduleRangeSmm(d, v, "train_schedule_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["train-schedule-range"] = t
		}
	}

	if v, ok := d.GetOk("train_test_days"); ok || d.HasChange("train_test_days") {
		t, err := expandSystemMlflowModelsTrainTestDaysSmm(d, v, "train_test_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["train-test-days"] = t
		}
	}

	if v, ok := d.GetOk("train_test_sensitivity"); ok || d.HasChange("train_test_sensitivity") {
		t, err := expandSystemMlflowModelsTrainTestSensitivitySmm(d, v, "train_test_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["train-test-sensitivity"] = t
		}
	}

	return &obj, nil
}
