## 1.2.0 (Unreleased)

## 1.1.0 (Dec 1, 2022)

IMPROVEMENTS:

* Support FortiAnalyzer v7.0.4, v7.0.5, v7.2.0, v7.2.1;
* Add computed feature for some arguments;
* Fix argument type change between different FortiAnalyzer version of resource resource_system_ha;
* Fix mkey issue for some resources; 
* Fix data type convert issue;
* Fix mkey issue for resource system_log_ratelimit_ratelimits;

FEATURES:

* **New Resource:** `fortianalyzer_fmupdate_fwmsetting_upgradetimeout`
* **New Resource:** `fortianalyzer_system_ha_vip`
* **New Resource:** `fortianalyzer_system_localinpolicy`
* **New Resource:** `fortianalyzer_system_localinpolicy6`

## 1.0.2 (Aug 31, 2022)

IMPROVEMENTS:

* Upgrade Terraform SDK to v2;

## 1.0.1 (Jun 20, 2022)

IMPROVEMENTS:

* Update sensitive keys
* Fix type convert issue
* Fix captial letters issue on Doc
* Update delete operation for arguments
* Fix argument of 'fazadom' issue
* Add examples in Doc

## 1.0.0 (Mar 11, 2022)

FEATURES:

* **New Resource:** `fortianalyzer_json_generic_api`
* **New Resource:** `fortianalyzer_dvm_cmd_add_device`
* **New Resource:** `fortianalyzer_dvm_cmd_del_device`
* **New Resource:** `fortianalyzer_dvmdb_adom`
* **New Resource:** `fortianalyzer_dvmdb_group`
* **New Resource:** `fortianalyzer_fmupdate_analyzer_virusreport`
* **New Resource:** `fortianalyzer_fmupdate_avips_advancedlog`
* **New Resource:** `fortianalyzer_fmupdate_avips_webproxy`
* **New Resource:** `fortianalyzer_fmupdate_customurllist`
* **New Resource:** `fortianalyzer_fmupdate_diskquota`
* **New Resource:** `fortianalyzer_fmupdate_fctservices`
* **New Resource:** `fortianalyzer_fmupdate_fdssetting`
* **New Resource:** `fortianalyzer_fmupdate_fdssetting_pushoverride`
* **New Resource:** `fortianalyzer_fmupdate_fdssetting_pushoverridetoclient`
* **New Resource:** `fortianalyzer_fmupdate_fdssetting_serveroverride`
* **New Resource:** `fortianalyzer_fmupdate_fdssetting_updateschedule`
* **New Resource:** `fortianalyzer_fmupdate_fwmsetting`
* **New Resource:** `fortianalyzer_fmupdate_multilayer`
* **New Resource:** `fortianalyzer_fmupdate_publicnetwork`
* **New Resource:** `fortianalyzer_fmupdate_serveraccesspriorities`
* **New Resource:** `fortianalyzer_fmupdate_serveroverridestatus`
* **New Resource:** `fortianalyzer_fmupdate_service`
* **New Resource:** `fortianalyzer_fmupdate_webspam_fgdsetting`
* **New Resource:** `fortianalyzer_fmupdate_webspam_webproxy`
* **New Resource:** `fortianalyzer_system_admin_group`
* **New Resource:** `fortianalyzer_system_admin_ldap`
* **New Resource:** `fortianalyzer_system_admin_profile`
* **New Resource:** `fortianalyzer_system_admin_radius`
* **New Resource:** `fortianalyzer_system_admin_setting`
* **New Resource:** `fortianalyzer_system_admin_tacacs`
* **New Resource:** `fortianalyzer_system_admin_user`
* **New Resource:** `fortianalyzer_system_alertconsole`
* **New Resource:** `fortianalyzer_system_alertevent`
* **New Resource:** `fortianalyzer_system_alertemail`
* **New Resource:** `fortianalyzer_system_autodelete`
* **New Resource:** `fortianalyzer_system_autodelete_dlpfilesautodeletion`
* **New Resource:** `fortianalyzer_system_autodelete_logautodeletion`
* **New Resource:** `fortianalyzer_system_autodelete_quarantinefilesautodeletion`
* **New Resource:** `fortianalyzer_system_autodelete_reportautodeletion`
* **New Resource:** `fortianalyzer_system_backup_allsettings`
* **New Resource:** `fortianalyzer_system_centralmanagement`
* **New Resource:** `fortianalyzer_system_certificate_ca`
* **New Resource:** `fortianalyzer_system_certificate_crl`
* **New Resource:** `fortianalyzer_system_certificate_local`
* **New Resource:** `fortianalyzer_system_certificate_oftp`
* **New Resource:** `fortianalyzer_system_certificate_remote`
* **New Resource:** `fortianalyzer_system_certificate_ssh`
* **New Resource:** `fortianalyzer_system_connector`
* **New Resource:** `fortianalyzer_system_dns`
* **New Resource:** `fortianalyzer_system_docker`
* **New Resource:** `fortianalyzer_system_fips`
* **New Resource:** `fortianalyzer_system_fortiview_autocache`
* **New Resource:** `fortianalyzer_system_fortiview_setting`
* **New Resource:** `fortianalyzer_system_global`
* **New Resource:** `fortianalyzer_system_global_sslciphersuites`
* **New Resource:** `fortianalyzer_system_guiact`
* **New Resource:** `fortianalyzer_system_ha`
* **New Resource:** `fortianalyzer_system_ha_peer`
* **New Resource:** `fortianalyzer_system_ha_privatepeer`
* **New Resource:** `fortianalyzer_system_interface`
* **New Resource:** `fortianalyzer_system_locallog_disk_filter`
* **New Resource:** `fortianalyzer_system_locallog_disk_setting`
* **New Resource:** `fortianalyzer_system_locallog_fortianalyzer2_filter`
* **New Resource:** `fortianalyzer_system_locallog_fortianalyzer2_setting`
* **New Resource:** `fortianalyzer_system_locallog_fortianalyzer3_filter`
* **New Resource:** `fortianalyzer_system_locallog_fortianalyzer3_setting`
* **New Resource:** `fortianalyzer_system_locallog_fortianalyzer_filter`
* **New Resource:** `fortianalyzer_system_locallog_fortianalyzer_setting`
* **New Resource:** `fortianalyzer_system_locallog_memory_filter`
* **New Resource:** `fortianalyzer_system_locallog_memory_setting`
* **New Resource:** `fortianalyzer_system_locallog_setting`
* **New Resource:** `fortianalyzer_system_locallog_syslogd2_filter`
* **New Resource:** `fortianalyzer_system_locallog_syslogd2_setting`
* **New Resource:** `fortianalyzer_system_locallog_syslogd3_filter`
* **New Resource:** `fortianalyzer_system_locallog_syslogd3_setting`
* **New Resource:** `fortianalyzer_system_locallog_syslogd_filter`
* **New Resource:** `fortianalyzer_system_locallog_syslogd_setting`
* **New Resource:** `fortianalyzer_system_logfetch_clientprofile`
* **New Resource:** `fortianalyzer_system_logfetch_serversettings`
* **New Resource:** `fortianalyzer_system_logforward`
* **New Resource:** `fortianalyzer_system_logforwardservice`
* **New Resource:** `fortianalyzer_system_log_alert`
* **New Resource:** `fortianalyzer_system_log_devicedisable`
* **New Resource:** `fortianalyzer_system_log_fospolicystats`
* **New Resource:** `fortianalyzer_system_log_interfacestats`
* **New Resource:** `fortianalyzer_system_log_ioc`
* **New Resource:** `fortianalyzer_system_log_ratelimit_ratelimits`
* **New Resource:** `fortianalyzer_system_log_maildomain`
* **New Resource:** `fortianalyzer_system_log_ratelimit`
* **New Resource:** `fortianalyzer_system_log_ratelimit_device`
* **New Resource:** `fortianalyzer_system_log_settings`
* **New Resource:** `fortianalyzer_system_log_settings_rollinganalyzer`
* **New Resource:** `fortianalyzer_system_log_settings_rollinglocal`
* **New Resource:** `fortianalyzer_system_log_settings_rollingregular`
* **New Resource:** `fortianalyzer_system_log_topology`
* **New Resource:** `fortianalyzer_system_mail`
* **New Resource:** `fortianalyzer_system_metadata_admins`
* **New Resource:** `fortianalyzer_system_ntp`
* **New Resource:** `fortianalyzer_system_ntp_ntpserver`
* **New Resource:** `fortianalyzer_system_passwordpolicy`
* **New Resource:** `fortianalyzer_system_report_autocache`
* **New Resource:** `fortianalyzer_system_report_estbrowsetime`
* **New Resource:** `fortianalyzer_system_report_setting`
* **New Resource:** `fortianalyzer_system_route`
* **New Resource:** `fortianalyzer_system_route6`
* **New Resource:** `fortianalyzer_system_saml`
* **New Resource:** `fortianalyzer_system_saml_fabricidp`
* **New Resource:** `fortianalyzer_system_saml_serviceproviders`
* **New Resource:** `fortianalyzer_system_sniffer`
* **New Resource:** `fortianalyzer_system_snmp_community`
* **New Resource:** `fortianalyzer_system_snmp_sysinfo`
* **New Resource:** `fortianalyzer_system_snmp_user`
* **New Resource:** `fortianalyzer_system_socfabric`
* **New Resource:** `fortianalyzer_system_sql`
* **New Resource:** `fortianalyzer_system_sql_customindex`
* **New Resource:** `fortianalyzer_system_sql_customskipidx`
* **New Resource:** `fortianalyzer_system_webproxy`
* **New Resource:** `fortianalyzer_system_sql_tsindexfield`
* **New Resource:** `fortianalyzer_system_syslog`
* **New Resource:** `fortianalyzer_system_workflow_approvalmatrix`