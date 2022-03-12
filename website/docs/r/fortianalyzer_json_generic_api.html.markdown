---
subcategory: "Generic"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_json_feneric_api"
description: |-
  FortiAnalyzer API Generic Interface.

---

# fortianalyzer_json_generic_api

FortiAnalyzer API Generic Interface.

## Example Usage

```hcl
provider "fortianalyzer" {
  hostname = "192.168.52.178"
  username = "admin"
  password = "admin"
  insecure = "true"

  scopetype = "adom"
  adom      = "root"
}

resource "fortianalyzer_json_generic_api" "test1" {
  json_content = <<JSON
{
    "method": "add",
    "params": [
        {
            "url": "/cli/global/system/admin/user",
            "data": {
                "userid": "terr-user",
                "password": ["fortinet"],
                "description": "This is a Terraform test"
            }
        }
    ]
}
JSON
}

output name1 {
  value       = jsondecode(fortianalyzer_json_generic_api.test1.response)
}

resource "fortianalyzer_json_generic_api" "test2" {
  json_content = <<JSON
{
    "method": "get",
    "params": [
        {
            "url": "/cli/global/system/admin/user"
        }
    ]
}
JSON
}

output name2 {
  value       = jsondecode(fortianalyzer_json_generic_api.test2.response)
}

```

## Argument Reference


The following arguments are supported:

* `force_recreate` - The argument is optional, if it is set, when its value changes, the resource will be re-created. It is usually used when the return value needs to be forced to update.
* `json_content` - Body data in JSON format.
* `comment` - Comment.
* `response` - API returns results.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - The resource id.
