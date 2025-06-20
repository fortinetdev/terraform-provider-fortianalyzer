---
subcategory: "System Admin"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_admin_ldap"
description: |-
  LDAP server entry configuration.
---

# fortianalyzer_system_admin_ldap
LDAP server entry configuration.

## Example Usage

```hcl
resource "fortianalyzer_system_admin_ldap" "trname" {
  cnid            = "uid"
  connect_timeout = 10
  dn              = "dc=example,dc=com"
  memberof_attr   = "memberOf"
  name            = "ldap_user"
  password        = ["password"]
  port            = 389
  secure          = "disable"
  server          = "1.1.1.1"
  type            = "simple"
  username        = "admin"
}
```

## Argument Reference


The following arguments are supported:


* `fazadom` - Adom. The structure of `fazadom` block is documented below.
* `adom_access` - set all or specify adom access type. all - All ADOMs access. specify - Specify ADOMs access. Valid values: `all`, `specify`.

* `adom_attr` - Attribute used to retrieve adom
* `attributes` - Attributes used for group searching.
* `ca_cert` - CA certificate name.
* `cnid` - Common Name Identifier (default = CN).
* `connect_timeout` - LDAP connection timeout (msec).
* `dn` - Distinguished Name.
* `filter` - Filter used for group searching.
* `group` - Full base DN used for group searching.
* `memberof_attr` - Attribute used to retrieve memeberof.
* `name` - LDAP server entry name.
* `password` - Password for initial binding.
* `port` - Port number of LDAP server (default = 389).
* `profile_attr` - Attribute used to retrieve admin profile.
* `secondary_server` - {<name_str|ip_str>} secondary LDAP server domain name or IP.
* `secure` - SSL connection. disable - No SSL. starttls - Use StartTLS. ldaps - Use LDAPS. Valid values: `disable`, `starttls`, `ldaps`.

* `server` - {<name_str|ip_str>} LDAP server domain name or IP.
* `ssl_protocol` - set the lowest SSL protocol version for connection to ldap server. follow-global-ssl-protocol - Follow system.global.global-ssl-protocol setting (default). sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version. tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: `follow-global-ssl-protocol`, `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.

* `tertiary_server` - {<name_str|ip_str>} tertiary LDAP server domain name or IP.
* `type` - Type of LDAP binding. simple - Simple password authentication without search. anonymous - Bind using anonymous user search. regular - Bind using username/password and then search. Valid values: `simple`, `anonymous`, `regular`.

* `username` - Username (full DN) for initial binding.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `fazadom` block supports:

* `adom_name` - Admin domain names.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AdminLdap can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_admin_ldap.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

