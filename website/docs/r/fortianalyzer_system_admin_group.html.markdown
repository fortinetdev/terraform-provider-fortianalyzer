---
subcategory: "System Admin"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_admin_group"
description: |-
  User group.
---

# fortianalyzer_system_admin_group
User group.

## Example Usage

```hcl
resource "fortianalyzer_system_admin_group" "trname" {
  name = "group"
  member {
    name = fortianalyzer_system_admin_ldap.member1.name
  }
  member {
    name = fortianalyzer_system_admin_radius.member2.name
  }
}

resource "fortianalyzer_system_admin_ldap" "member1" {
  cnid            = "uid"
  connect_timeout = 10
  dn              = "dc=example,dc=com"
  memberof_attr   = "memberOf"
  name            = "member1"
  password        = ["password"]
  port            = 389
  secure          = "disable"
  server          = "1.1.1.1"
  type            = "simple"
  username        = "admin"
}

resource "fortianalyzer_system_admin_radius" "member2" {
  auth_type        = "any"
  name             = "member2"
  port             = 1812
  secondary_secret = ["secondary_secret"]
  secondary_server = "2.2.2.2"
  secret           = ["secret"]
  server           = "1.1.1.1"
}
```

## Argument Reference


The following arguments are supported:


* `member` - Member. The structure of `member` block is documented below.
* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `member` block supports:

* `name` - Group member name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AdminGroup can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_admin_group.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

