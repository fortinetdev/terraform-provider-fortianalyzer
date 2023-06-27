// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Provider for FortiAnalyzer

package fortianalyzer

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Provider creates and returns the FortiAnalyzer *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The hostname/IP address of the FORTIANALYZER to be connected",
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},

			"cabundlefile": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "CA Bundle file",
			},

			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "adom",
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
				}, false),
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},

			"import_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"fortianalyzer_json_generic_api": resourceJsonGenericAPI(),

			"fortianalyzer_dvm_cmd_add_device":                            resourceDvmCmdAddDevice(),
			"fortianalyzer_dvm_cmd_del_device":                            resourceDvmCmdDelDevice(),
			"fortianalyzer_dvmdb_adom":                                    resourceDvmdbAdom(),
			"fortianalyzer_dvmdb_group":                                   resourceDvmdbGroup(),
			"fortianalyzer_fmupdate_analyzer_virusreport":                 resourceFmupdateAnalyzerVirusreport(),
			"fortianalyzer_fmupdate_avips_advancedlog":                    resourceFmupdateAvIpsAdvancedLog(),
			"fortianalyzer_fmupdate_avips_webproxy":                       resourceFmupdateAvIpsWebProxy(),
			"fortianalyzer_fmupdate_customurllist":                        resourceFmupdateCustomUrlList(),
			"fortianalyzer_fmupdate_diskquota":                            resourceFmupdateDiskQuota(),
			"fortianalyzer_fmupdate_fctservices":                          resourceFmupdateFctServices(),
			"fortianalyzer_fmupdate_fdssetting":                           resourceFmupdateFdsSetting(),
			"fortianalyzer_fmupdate_fdssetting_pushoverride":              resourceFmupdateFdsSettingPushOverride(),
			"fortianalyzer_fmupdate_fdssetting_pushoverridetoclient":      resourceFmupdateFdsSettingPushOverrideToClient(),
			"fortianalyzer_fmupdate_fdssetting_serveroverride":            resourceFmupdateFdsSettingServerOverride(),
			"fortianalyzer_fmupdate_fdssetting_updateschedule":            resourceFmupdateFdsSettingUpdateSchedule(),
			"fortianalyzer_fmupdate_fwmsetting":                           resourceFmupdateFwmSetting(),
			"fortianalyzer_fmupdate_fwmsetting_upgradetimeout":            resourceFmupdateFwmSettingUpgradeTimeout(),
			"fortianalyzer_fmupdate_multilayer":                           resourceFmupdateMultilayer(),
			"fortianalyzer_fmupdate_publicnetwork":                        resourceFmupdatePublicnetwork(),
			"fortianalyzer_fmupdate_serveraccesspriorities":               resourceFmupdateServerAccessPriorities(),
			"fortianalyzer_fmupdate_serveroverridestatus":                 resourceFmupdateServerOverrideStatus(),
			"fortianalyzer_fmupdate_service":                              resourceFmupdateService(),
			"fortianalyzer_fmupdate_webspam_fgdsetting":                   resourceFmupdateWebSpamFgdSetting(),
			"fortianalyzer_fmupdate_webspam_webproxy":                     resourceFmupdateWebSpamWebProxy(),
			"fortianalyzer_system_admin_group":                            resourceSystemAdminGroup(),
			"fortianalyzer_system_admin_ldap":                             resourceSystemAdminLdap(),
			"fortianalyzer_system_admin_profile":                          resourceSystemAdminProfile(),
			"fortianalyzer_system_admin_radius":                           resourceSystemAdminRadius(),
			"fortianalyzer_system_admin_setting":                          resourceSystemAdminSetting(),
			"fortianalyzer_system_admin_tacacs":                           resourceSystemAdminTacacs(),
			"fortianalyzer_system_admin_user":                             resourceSystemAdminUser(),
			"fortianalyzer_system_alertconsole":                           resourceSystemAlertConsole(),
			"fortianalyzer_system_alertevent":                             resourceSystemAlertEvent(),
			"fortianalyzer_system_alertemail":                             resourceSystemAlertemail(),
			"fortianalyzer_system_autodelete":                             resourceSystemAutoDelete(),
			"fortianalyzer_system_autodelete_dlpfilesautodeletion":        resourceSystemAutoDeleteDlpFilesAutoDeletion(),
			"fortianalyzer_system_autodelete_logautodeletion":             resourceSystemAutoDeleteLogAutoDeletion(),
			"fortianalyzer_system_autodelete_quarantinefilesautodeletion": resourceSystemAutoDeleteQuarantineFilesAutoDeletion(),
			"fortianalyzer_system_autodelete_reportautodeletion":          resourceSystemAutoDeleteReportAutoDeletion(),
			"fortianalyzer_system_backup_allsettings":                     resourceSystemBackupAllSettings(),
			"fortianalyzer_system_centralmanagement":                      resourceSystemCentralManagement(),
			"fortianalyzer_system_certificate_ca":                         resourceSystemCertificateCa(),
			"fortianalyzer_system_certificate_crl":                        resourceSystemCertificateCrl(),
			"fortianalyzer_system_certificate_local":                      resourceSystemCertificateLocal(),
			"fortianalyzer_system_certificate_oftp":                       resourceSystemCertificateOftp(),
			"fortianalyzer_system_certificate_remote":                     resourceSystemCertificateRemote(),
			"fortianalyzer_system_certificate_ssh":                        resourceSystemCertificateSsh(),
			"fortianalyzer_system_connector":                              resourceSystemConnector(),
			"fortianalyzer_system_dns":                                    resourceSystemDns(),
			"fortianalyzer_system_docker":                                 resourceSystemDocker(),
			"fortianalyzer_system_fips":                                   resourceSystemFips(),
			"fortianalyzer_system_fortiview_autocache":                    resourceSystemFortiviewAutoCache(),
			"fortianalyzer_system_fortiview_setting":                      resourceSystemFortiviewSetting(),
			"fortianalyzer_system_global":                                 resourceSystemGlobal(),
			"fortianalyzer_system_global_sslciphersuites":                 resourceSystemGlobalSslCipherSuites(),
			"fortianalyzer_system_guiact":                                 resourceSystemGuiact(),
			"fortianalyzer_system_ha":                                     resourceSystemHa(),
			"fortianalyzer_system_ha_peer":                                resourceSystemHaPeer(),
			"fortianalyzer_system_ha_privatepeer":                         resourceSystemHaPrivatePeer(),
			"fortianalyzer_system_ha_vip":                                 resourceSystemHaVip(),
			"fortianalyzer_system_interface":                              resourceSystemInterface(),
			"fortianalyzer_system_localinpolicy":                          resourceSystemLocalInPolicy(),
			"fortianalyzer_system_localinpolicy6":                         resourceSystemLocalInPolicy6(),
			"fortianalyzer_system_locallog_disk_filter":                   resourceSystemLocallogDiskFilter(),
			"fortianalyzer_system_locallog_disk_setting":                  resourceSystemLocallogDiskSetting(),
			"fortianalyzer_system_locallog_fortianalyzer2_filter":         resourceSystemLocallogFortianalyzer2Filter(),
			"fortianalyzer_system_locallog_fortianalyzer2_setting":        resourceSystemLocallogFortianalyzer2Setting(),
			"fortianalyzer_system_locallog_fortianalyzer3_filter":         resourceSystemLocallogFortianalyzer3Filter(),
			"fortianalyzer_system_locallog_fortianalyzer3_setting":        resourceSystemLocallogFortianalyzer3Setting(),
			"fortianalyzer_system_locallog_fortianalyzer_filter":          resourceSystemLocallogFortianalyzerFilter(),
			"fortianalyzer_system_locallog_fortianalyzer_setting":         resourceSystemLocallogFortianalyzerSetting(),
			"fortianalyzer_system_locallog_memory_filter":                 resourceSystemLocallogMemoryFilter(),
			"fortianalyzer_system_locallog_memory_setting":                resourceSystemLocallogMemorySetting(),
			"fortianalyzer_system_locallog_setting":                       resourceSystemLocallogSetting(),
			"fortianalyzer_system_locallog_syslogd2_filter":               resourceSystemLocallogSyslogd2Filter(),
			"fortianalyzer_system_locallog_syslogd2_setting":              resourceSystemLocallogSyslogd2Setting(),
			"fortianalyzer_system_locallog_syslogd3_filter":               resourceSystemLocallogSyslogd3Filter(),
			"fortianalyzer_system_locallog_syslogd3_setting":              resourceSystemLocallogSyslogd3Setting(),
			"fortianalyzer_system_locallog_syslogd_filter":                resourceSystemLocallogSyslogdFilter(),
			"fortianalyzer_system_locallog_syslogd_setting":               resourceSystemLocallogSyslogdSetting(),
			"fortianalyzer_system_logfetch_clientprofile":                 resourceSystemLogFetchClientProfile(),
			"fortianalyzer_system_logfetch_serversettings":                resourceSystemLogFetchServerSettings(),
			"fortianalyzer_system_logforward":                             resourceSystemLogForward(),
			"fortianalyzer_system_logforwardservice":                      resourceSystemLogForwardService(),
			"fortianalyzer_system_log_alert":                              resourceSystemLogAlert(),
			"fortianalyzer_system_log_devicedisable":                      resourceSystemLogDeviceDisable(),
			"fortianalyzer_system_log_fospolicystats":                     resourceSystemLogFosPolicyStats(),
			"fortianalyzer_system_log_interfacestats":                     resourceSystemLogInterfaceStats(),
			"fortianalyzer_system_log_ioc":                                resourceSystemLogIoc(),
			"fortianalyzer_system_log_maildomain":                         resourceSystemLogMailDomain(),
			"fortianalyzer_system_log_ratelimit":                          resourceSystemLogRatelimit(),
			"fortianalyzer_system_log_ratelimit_device":                   resourceSystemLogRatelimitDevice(),
			"fortianalyzer_system_log_ratelimit_ratelimits":               resourceSystemLogRatelimitRatelimits(),
			"fortianalyzer_system_log_settings":                           resourceSystemLogSettings(),
			"fortianalyzer_system_log_settings_rollinganalyzer":           resourceSystemLogSettingsRollingAnalyzer(),
			"fortianalyzer_system_log_settings_rollinglocal":              resourceSystemLogSettingsRollingLocal(),
			"fortianalyzer_system_log_settings_rollingregular":            resourceSystemLogSettingsRollingRegular(),
			"fortianalyzer_system_log_topology":                           resourceSystemLogTopology(),
			"fortianalyzer_system_mail":                                   resourceSystemMail(),
			"fortianalyzer_system_metadata_admins":                        resourceSystemMetadataAdmins(),
			"fortianalyzer_system_ntp":                                    resourceSystemNtp(),
			"fortianalyzer_system_ntp_ntpserver":                          resourceSystemNtpNtpserver(),
			"fortianalyzer_system_passwordpolicy":                         resourceSystemPasswordPolicy(),
			"fortianalyzer_system_report_autocache":                       resourceSystemReportAutoCache(),
			"fortianalyzer_system_report_estbrowsetime":                   resourceSystemReportEstBrowseTime(),
			"fortianalyzer_system_report_setting":                         resourceSystemReportSetting(),
			"fortianalyzer_system_route":                                  resourceSystemRoute(),
			"fortianalyzer_system_route6":                                 resourceSystemRoute6(),
			"fortianalyzer_system_saml":                                   resourceSystemSaml(),
			"fortianalyzer_system_saml_fabricidp":                         resourceSystemSamlFabricIdp(),
			"fortianalyzer_system_saml_serviceproviders":                  resourceSystemSamlServiceProviders(),
			"fortianalyzer_system_sniffer":                                resourceSystemSniffer(),
			"fortianalyzer_system_snmp_community":                         resourceSystemSnmpCommunity(),
			"fortianalyzer_system_snmp_sysinfo":                           resourceSystemSnmpSysinfo(),
			"fortianalyzer_system_socfabric_trustedlist":                  resourceSystemSocFabricTrustedList(),
			"fortianalyzer_system_snmp_user":                              resourceSystemSnmpUser(),
			"fortianalyzer_system_socfabric":                              resourceSystemSocFabric(),
			"fortianalyzer_system_sql":                                    resourceSystemSql(),
			"fortianalyzer_system_sql_customindex":                        resourceSystemSqlCustomIndex(),
			"fortianalyzer_system_sql_customskipidx":                      resourceSystemSqlCustomSkipidx(),
			"fortianalyzer_system_sql_tsindexfield":                       resourceSystemSqlTsIndexField(),
			"fortianalyzer_system_syslog":                                 resourceSystemSyslog(),
			"fortianalyzer_system_webproxy":                               resourceSystemWebProxy(),
			"fortianalyzer_system_workflow_approvalmatrix":                resourceSystemWorkflowApprovalMatrix(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Hostname:      d.Get("hostname").(string),
		User:          d.Get("username").(string),
		Passwd:        d.Get("password").(string),
		CABundle:      d.Get("cabundlefile").(string),
		ScopeType:     d.Get("scopetype").(string),
		Adom:          d.Get("adom").(string),
		ImportOptions: d.Get("import_options").(*schema.Set),
	}

	v1, ok1 := d.GetOkExists("insecure")
	if ok1 {
		insecure := v1.(bool)
		config.Insecure = &insecure
	}

	// Create Client for later connections
	return config.CreateClient()
}
