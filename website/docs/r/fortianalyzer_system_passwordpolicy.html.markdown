---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_passwordpolicy"
description: |-
  Password policy.
---

# fortianalyzer_system_passwordpolicy
Password policy.

## Example Usage

```hcl
resource "fortianalyzer_system_passwordpolicy" "trname" {
  change_4_characters = "enable"
  expire              = 0
  minimum_length      = 8
  must_contain        = ["upper-case-letter", "lower-case-letter", "number", "non-alphanumeric"]
  status              = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `change_4_characters` - Enable/disable changing at least 4 characters for new password. disable - Disable changing at least 4 characters for new password. enable - Enable changing at least 4 characters for new password. Valid values: `disable`, `enable`.

* `expire` - Number of days after which admin users' password will expire (0 - 3650, 0 = never expire).
* `login_lockout_upon_downgrade` - Enable/disable administrative user login lockout upon downgrade (defaut = disable). If enabled, downgrading the FortiOS firmware to a lower version where safer passwords are unsupported will lock out administrative users. disable - Disable administrative user login lockout upon downgrade. enable - Enable administrative user login lockout upon downgrade. Valid values: `disable`, `enable`.

* `minimum_length` - Minimum password length.
* `must_contain` - Password character requirements. upper-case-letter - Require password to contain upper case letter. lower-case-letter - Require password to contain lower case letter. number - Require password to contain number. non-alphanumeric - Require password to contain non-alphanumeric characters. Valid values: `upper-case-letter`, `lower-case-letter`, `number`, `non-alphanumeric`.

* `password_history` - Number of unique new passwords that must be used before old password can be reused (0 - 20).
* `status` - Enable/disable password policy. disable - Disable password policy. enable - Enable password policy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System PasswordPolicy can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_passwordpolicy.labelname SystemPasswordPolicy
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

