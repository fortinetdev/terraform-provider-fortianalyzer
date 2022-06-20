---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_docker"
description: |-
  Docker host.
---

# fortianalyzer_system_docker
Docker host.

## Example Usage

```hcl
resource "fortianalyzer_system_docker" "trname" {
  cpu                       = 50
  default_address_pool_base = ["172.17.0.0", "255.255.0.0"]
  default_address_pool_size = 24
  docker_user_login_max     = 32
  mem                       = 50
  status                    = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `cpu` - Cpu.
* `default_address_pool_base` - Set default-address-pool CIDR.
* `default_address_pool_size` - Set default-address-pool size.
* `docker_user_login_max` - Max login session for docker users.
* `fortiaiops` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortiauthenticator` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortiportal` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortisigconverter` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortisoar` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortiwlm` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fsmcollector` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `mem` - Max % RAM usage.
* `sdwancontroller` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `status` - Enable and set registry. disable - Disable docker host service. enable - Enable production registry. qa - Enable QA test registry. dev - Enable QA test registry (without signature). Valid values: `disable`, `enable`, `qa`, `dev`.

* `universalconnector` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Docker can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_docker.labelname SystemDocker
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

