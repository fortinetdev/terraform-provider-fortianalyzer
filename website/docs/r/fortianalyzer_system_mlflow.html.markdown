---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_mlflow"
description: |-
  System Mlflow
---

# fortianalyzer_system_mlflow
System Mlflow

## Example Usage

```hcl
# Replace example credentials, serial numbers, and addresses with your own values.
resource "fortianalyzer_system_mlflow" "test" {
  status = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `models` - Models. The structure of `models` block is documented below.
* `status` - Status. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `models` block supports:

* `artifact_max_count` - Artifact-Max-Count.
* `artifact_retention` - Artifact-Retention.
* `asset_max_count` - Asset-Max-Count.
* `auto_deploy` - Auto-Deploy. Valid values: `disable`, `enable`.

* `id` - Id.
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
* `id` - an identifier for the resource.

## Import

System Mlflow can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_mlflow.labelname SystemMlflow
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

