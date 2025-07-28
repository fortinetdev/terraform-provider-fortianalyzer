---
subcategory: "No Category"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_logforward"
description: |-
  Log forwarding.
---

# fortianalyzer_system_logforward
Log forwarding.

## Example Usage

```hcl
resource "fortianalyzer_system_logforward" "trname" {
  fwd_archive_types          = ["Web_Archive", "Email_Archive", "IM_Archive", "File_Transfer_Archive", "MMS_Archive", "AV_Quarantine", "IPS_Packets", "EDISC_Archive"]
  fwd_archives               = "enable"
  fwd_compression            = "disable"
  fwd_facility               = "local7"
  fwd_log_source_ip          = "local_ip"
  fwd_max_delay              = "5min"
  fwd_reliable               = "disable"
  fwd_server_type            = "fortianalyzer"
  fosid                      = 2
  log_field_exclusion_status = "disable"
  log_filter_status          = "disable"
  log_masking_status         = "disable"
  mode                       = "forwarding"
  server_name                = "server"
  server_addr                = "1.1.1.1"
}
```

## Argument Reference


The following arguments are supported:


* `agg_archive_types` - Archive types. Web_Archive -  Secure_Web_Archive -  Email_Archive -  File_Transfer_Archive -  IM_Archive -  MMS_Archive -  AV_Quarantine -  IPS_Packets -  Valid values: `Web_Archive`, `Secure_Web_Archive`, `Email_Archive`, `File_Transfer_Archive`, `IM_Archive`, `MMS_Archive`, `AV_Quarantine`, `IPS_Packets`.

* `agg_data_end_time` - End date and time of the data-range &lt;hh:mm yyyy/mm/dd&gt;.
* `agg_data_start_time` - Start date and time of the data-range &lt;hh:mm yyyy/mm/dd&gt;.
* `agg_logtypes` - Log types. none - none app-ctrl -  attack -  content -  dlp -  emailfilter -  event -  generic -  history -  traffic -  virus -  webfilter -  netscan -  fct-event -  fct-traffic -  fct-netscan -  waf -  gtp -  dns -  ssh -  ssl -  file-filter -  asset -  protocol -  siem -  Valid values: `none`, `app-ctrl`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `webfilter`, `netscan`, `fct-event`, `fct-traffic`, `fct-netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `asset`, `protocol`, `siem`.

* `agg_password` - Log aggregation access password for server.
* `agg_schedule` - Schedule log aggregation mode. daily - Run daily log aggregation on-demand - Run log aggregation on demand Valid values: `daily`, `on-demand`.

* `agg_time` - Daily at.
* `agg_user` - Log aggregation access user name for server.
* `device_filter` - Device-Filter. The structure of `device_filter` block is documented below.
* `fwd_archive_types` - forwarding archive types. Web_Archive -  Email_Archive -  IM_Archive -  File_Transfer_Archive -  MMS_Archive -  AV_Quarantine -  IPS_Packets -  EDISC_Archive -  Valid values: `Web_Archive`, `Email_Archive`, `IM_Archive`, `File_Transfer_Archive`, `MMS_Archive`, `AV_Quarantine`, `IPS_Packets`, `EDISC_Archive`.

* `fwd_archives` - Enable/disable forwarding archives. disable - Disable forwarding archives. enable - Enable forwarding archives. Valid values: `disable`, `enable`.

* `fwd_compression` - Enable/disable compression for better bandwidth efficiency. disable - Disable compression of messages. enable - Enable compression of messages. Valid values: `disable`, `enable`.

* `fwd_facility` - Facility for remote syslog. kernel - kernel messages user - random user level messages mail - Mail system. daemon - System daemons. auth - Security/authorization messages. syslog - Messages generated internally by syslog daemon. lpr - Line printer subsystem. news - Network news subsystem. uucp - Network news subsystem. clock - Clock daemon. authpriv - Security/authorization messages (private). ftp - FTP daemon. ntp - NTP daemon. audit - Log audit. alert - Log alert. cron - Clock daemon. local0 - Reserved for local use. local1 - Reserved for local use. local2 - Reserved for local use. local3 - Reserved for local use. local4 - Reserved for local use. local5 - Reserved for local use. local6 - Reserved for local use. local7 - Reserved for local use. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `clock`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `cron`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.

* `fwd_ha_bind_vip` - When HA is enabled, always use vip as forwarding port disable - Disable bind forwarding to vip interface. enable - Enable bind forwarding to vip interface. Valid values: `disable`, `enable`.

* `fwd_log_source_ip` - Log's source IP address (no effect for reliable forwarding). local_ip - Use FAZVM64 local ip. original_ip - Use original source ip. Valid values: `local_ip`, `original_ip`.

* `fwd_max_delay` - Max delay for near realtime log forwarding. realtime - Realtime forwarding, no delay. 1min - Near realtime forwarding with up to one miniute delay. 5min - Near realtime forwarding with up to five miniutes delay. Valid values: `realtime`, `1min`, `5min`.

* `fwd_output_plugin_id` - Name of the output plugin profile
* `fwd_reliable` - Enable/disable reliable logging. disable - Disable reliable logging. enable - Enable reliable logging. Valid values: `disable`, `enable`.

* `fwd_secure` - Enable/disable TLS/SSL secured reliable logging. disable - Disable TLS/SSL secured reliable logging. enable - Enable TLS/SSL secured reliable logging. Valid values: `disable`, `enable`.

* `fwd_server_type` - Forwarding all logs to syslog server or FortiAnalyzer. syslog - Forward logs to generic syslog server. fortianalyzer - Forward logs to FortiAnalyzer. cef - Forward logs to a CEF (Common Event Format) server. syslog-pack - Forward logs to a FortiAnalyzer which support packed syslog message. Valid values: `syslog`, `fortianalyzer`, `cef`, `syslog-pack`.

* `fwd_syslog_decode_b64` - Fwd-Syslog-Decode-B64. Valid values: `disable`, `enable`.

* `fwd_syslog_enrich_cve` - Fwd-Syslog-Enrich-Cve. Valid values: `disable`, `enable`.

* `fwd_syslog_format` - Forwarding format for syslog. fgt - fgt syslog format rfc-5424 - rfc-5424 syslog format Valid values: `fgt`, `rfc-5424`.

* `fwd_syslog_transparent` - Enable/disable transparently forwarding logs from syslog devices to syslog server. disable - Disable syslog transparent forward mode. Received syslogs becomes part of a FAZ syslog when forwarded out. enable - Enable syslog transparent forward mode. Received syslogs are forwarded without modifications. faz-enrich - Disable syslog transparent forward mode. Additional FAZ fields are added to the end of syslog. Valid values: `disable`, `enable`, `faz-enrich`.

* `fosid` - Log forwarding ID.
* `log_field_exclusion` - Log-Field-Exclusion. The structure of `log_field_exclusion` block is documented below.
* `log_field_exclusion_status` - Enable or disable log field exclusion. disable - Disable log field exclusion. enable - Enable log field exclusion. Valid values: `disable`, `enable`.

* `log_filter` - Log-Filter. The structure of `log_filter` block is documented below.
* `log_filter_logic` - Logic operator used to connect filters. and - Conjunctive filters. or - Disjunctive filters. Valid values: `and`, `or`.

* `log_filter_status` - Enable or disable log filtering. disable - Disable log filtering. enable - Enable log filtering. Valid values: `disable`, `enable`.

* `log_masking_custom` - Log-Masking-Custom. The structure of `log_masking_custom` block is documented below.
* `log_masking_custom_priority` - Prioritize custom fields. disable - Disable custom field search priority.  - Prioritize custom fields. Valid values: `disable`, ``.

* `log_masking_fields` - Log field masking fields. user - User name. srcip - Source IP. srcname - Source name. srcmac - Source MAC. dstip - Destination IP. dstname - Dst name. email - Email. message - Message. domain - Domain. Valid values: `user`, `srcip`, `srcname`, `srcmac`, `dstip`, `dstname`, `email`, `message`, `domain`.

* `log_masking_key` - Log field masking key.
* `log_masking_status` - Enable or disable log field masking. disable - Disable log field masking. enable - Enable log field masking. Valid values: `disable`, `enable`.

* `mode` - Log forwarding mode. forwarding - Realtime or near realtime forwarding logs to servers. aggregation - Aggregate logs and archives to Analyzer. disable - Do not forward or aggregate logs. Valid values: `forwarding`, `aggregation`, `disable`.

* `pcapurl_domain_ip` - The domain name or ip for forming a pcapurl. This pcapurl will be appended to applicable forwarded logs for downloading a pcap file.
* `pcapurl_enrich` - Enable/disable enriching pcapurl. disable - Disable enriching pcapurl. enable - Enable enriching pcapurl. It will append a &apos;pcapurl&apos; field to the forwarded syslogs. Valid values: `disable`, `enable`.

* `peer_cert_cn` - Certificate common name of log-forward server.
* `proxy_service` - Enable/disable proxy service under collector mode. disable - Disable proxy service. enable - Enable proxy service. Valid values: `disable`, `enable`.

* `proxy_service_priority` - Proxy service priority from 1 (lowest) to 20 (highest).
* `server_addr` - Remote server address.
* `server_device` - Log forwarding server device ID.
* `server_name` - Log forwarding server name.
* `server_port` - Server listen port (1 - 65535).
* `signature` - Aggregation cfg hash token.
* `sync_metadata` - Synchronizing meta data types. sf-topology - Security Fabric topology interface-role - Interface Role device - Device information endusr-avatar - End-user avatar Valid values: `sf-topology`, `interface-role`, `device`, `endusr-avatar`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `device_filter` block supports:

* `action` - Include or exclude the specified device. include - Include specified device. exclude - Exclude specified device. include-like - Include specified device matching the given wildcard expression. exclude-like - Exclude specified device matching the given wildcard expression. Valid values: `include`, `exclude`, `include-like`, `exclude-like`.

* `adom` - Adom name or (null) for all adoms, or a wildcard expression matching adom(s) if action is a like action.
* `device` - Device ID of log client device, or a wildcard expression matching log client device(s) if action is a like action.
* `id` - Device filter ID.

The `log_field_exclusion` block supports:

* `dev_type` - Device type. FortiGate - FortiGate Device FortiManager - FortiManager Device Syslog - Syslog Device FortiMail - FortiMail Device FortiWeb - FortiWeb Device FortiCache - FortiCache Device FortiAnalyzer - FortiAnalyzer Device FortiSandbox - FortiSandbox Device FortiDDoS - FortiDDoS Device FortiNAC - FortiNAC Device FortiDeceptor - FortiDeceptor Device FortiADC - FortiADC Device FortiFirewall - FortiFirewall Device Valid values: `FortiGate`, `FortiManager`, `Syslog`, `FortiMail`, `FortiWeb`, `FortiCache`, `FortiAnalyzer`, `FortiSandbox`, `FortiDDoS`, `FortiNAC`, `FortiDeceptor`, `FortiADC`, `FortiFirewall`.

* `field_list` - List of fields to be excluded.
* `id` - Log field exclusion ID.
* `log_type` - Log type. app-ctrl - Application Control appevent - APPEVENT attack - Attack content - DLP Archive dlp - Data Leak Prevention emailfilter - Email Filter event - Event generic - Generic history - Mail Statistics traffic - Traffic virus - Virus voip - VoIP webfilter - Web Filter netscan - Network Scan waf - WAF gtp - GTP dns - Domain Name System ssh - SSH ssl - SSL file-filter - FFLT Asset - Asset protocol - PROTOCOL ANY-TYPE - Any log type Valid values: `app-ctrl`, `appevent`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `Asset`, `protocol`, `ANY-TYPE`.


The `log_filter` block supports:

* `field` - Field name. type - Log type logid - Log ID level - Level devid - Device ID vd - Vdom ID srcip - Source IP srcintf - Source Interface dstip - Destination IP dstintf - Destination Interface dstport - Destination Port user - User group - Group free-text - General free-text filter Valid values: `type`, `logid`, `level`, `devid`, `vd`, `srcip`, `srcintf`, `dstip`, `dstintf`, `dstport`, `user`, `group`, `free-text`.

* `id` - Log filter ID.
* `oper` - Field filter operator. &lt; - =Less than or equal to &gt; - =Greater than or equal to contain - Contain not-contain - Not contain match - Match (expression) Valid values: `=`, `!=`, `<`, `>`, `<=`, `>=`, `contain`, `not-contain`, `match`.

* `value` - Field filter operand or free-text matching expression.

The `log_masking_custom` block supports:

* `field_name` - Field name.
* `field_type` - Field type. string - String. ip - IP. mac - MAC address. email - Email address. unknown - Unknown. Valid values: `string`, `ip`, `mac`, `email`, `unknown`.

* `id` - Field masking id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogForward can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_logforward.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

