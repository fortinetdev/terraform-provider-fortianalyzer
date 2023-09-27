---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_pcapfile"
description: |-
  Log pcap-file settings.
---

# fortianalyzer_system_log_pcapfile
Log pcap-file settings.

## Argument Reference


The following arguments are supported:


* `download_mode` - Download mode for pcap files. plain - Download original file. zip - Download zip file without password. zip-with-password - Download zip file with password. Valid values: `plain`, `zip`, `zip-with-password`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogPcapFile can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_pcapfile.labelname SystemLogPcapFile
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

