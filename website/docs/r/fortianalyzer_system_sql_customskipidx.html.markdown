---
subcategory: "System SQL"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_sql_customskipidx"
description: |-
  List of aditional SQL skip index fields.
---

# fortianalyzer_system_sql_customskipidx
List of aditional SQL skip index fields.

## Argument Reference


The following arguments are supported:


* `device_type` - Device type. FortiGate - Set device type to FortiGate. FortiManager - Set device type to FortiManager FortiClient - Set device type to FortiClient. FortiMail - Set device type to FortiMail. FortiWeb - Set device type to FortiWeb. FortiSandbox - Set device type to FortiSandbox FortiProxy - Set device type to FortiProxy Valid values: `FortiGate`, `FortiManager`, `FortiClient`, `FortiMail`, `FortiWeb`, `FortiSandbox`, `FortiProxy`.

* `fosid` - Add or Edit log index fields.
* `index_field` - Field to be added to skip index.
* `log_type` - Log type. app-ctrl -  attack -  content -  dlp -  emailfilter -  event -  generic -  history -  traffic -  virus -  voip -  webfilter -  netscan -  fct-event -  fct-traffic -  fct-netscan -  waf -  gtp -  dns -  ssh -  ssl -  file-filter -  asset -  protocol -  siem -  Valid values: `app-ctrl`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `fct-event`, `fct-traffic`, `fct-netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `asset`, `protocol`, `siem`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SqlCustomSkipidx can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_sql_customskipidx.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

