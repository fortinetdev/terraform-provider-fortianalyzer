---
subcategory: "System Report"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_report_setting"
description: |-
  Report settings.
---

# fortianalyzer_system_report_setting
Report settings.

## Example Usage

```hcl
resource "fortianalyzer_system_report_setting" "trname" {
  aggregate_report      = "disable"
  capwap_port           = 5246
  week_start            = "sun"
  template_auto_install = "default"
  report_priority       = "auto"
  ldap_cache_timeout    = 60
}
```

## Argument Reference


The following arguments are supported:


* `aggregate_report` - Enable/disable including a group report along with the per-device reports. disable - Exclude a group report along with the per-device reports. enable - Include a group report along with the per-device reports. Valid values: `disable`, `enable`.

* `capwap_port` - Exclude capwap traffic by port.
* `capwap_service` - Exclude capwap traffic by service.
* `exclude_capwap` - Exclude capwap traffic. disable - Disable. by-port - By port. by-service - By service. Valid values: `disable`, `by-port`, `by-service`.

* `hcache_lossless` - Usableness of ready-with-loss hcaches. disable - Use ready-with-loss hcaches. enable - Do not use ready-with-loss hcaches. Valid values: `disable`, `enable`.

* `ldap_cache_timeout` - LDAP cache timeout in minutes, default 60, 0 means not use cache.
* `max_rpt_pdf_rows` - Maximum number of rows can be generated in a single pdf.
* `max_table_rows` - Maximum number of rows can be generated in a single table.
* `report_priority` - Priority of sql report. high - High low - Low auto - Auto Valid values: `high`, `low`, `auto`.

* `template_auto_install` - The language used for new ADOMs (default = default). default - Default. english - English. Valid values: `default`, `english`.

* `week_start` - Day of the week on which the week starts. sun - Sunday. mon - Monday. Valid values: `sun`, `mon`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ReportSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_report_setting.labelname SystemReportSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

