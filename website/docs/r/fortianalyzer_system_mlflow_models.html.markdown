---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_mlflow_models"
description: |-
  System MlflowModels
---

# fortianalyzer_system_mlflow_models
System MlflowModels

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_mlflow_models" "test" {
  fosid                  = 1
  model_type             = "login-anomaly"
  status                 = "disable"
  train_test_sensitivity = "low"
}
```

## Argument Reference


The following arguments are supported:


* `artifact_max_count` - Artifact-Max-Count.
* `artifact_retention` - Artifact-Retention.
* `asset_max_count` - Asset-Max-Count.
* `auto_deploy` - Auto-Deploy. Valid values: `disable`, `enable`.

* `fosid` - Id.
* `inference_interval` - Inference-Interval.
* `model_type` - Model-Type. Valid values: `login-anomaly`, `traffic-download-anomaly`, `traffic-upload-anomaly`.

* `status` - Status. Valid values: `disable`, `enable`.

* `train_data_days` - Train-Data-Days.
* `train_repeat_at` - Train-Repeat-At.
* `train_schedule_range` - Train-Schedule-Range.
* `train_test_days` - Train-Test-Days.
* `train_test_sensitivity` - Train-Test-Sensitivity. Valid values: `low`, `medium`, `high`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System MlflowModels can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_mlflow_models.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

