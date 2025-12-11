---
layout: "fortianalyzer"
page_title: "Provider: FortiAnalyzer"
sidebar_current: "docs-fortianalyzer-index"
description: |-
  The FortiAnalyzer provider interacts with FortiAnalyzer.
---

# FortiAnalyzer Provider

The FortiAnalyzer provider is used to interact with the resources supported by FortiAnalyzer. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources.

## Example Usage

```hcl
# Configure the Provider for FortiAnalyzer
provider "fortianalyzer" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"

  scopetype    = "adom"
  adom         = "root"
}

# Create an admin user
resource "fortianalyzer_system_admin_user" "trname" {
    userid      = "terr-user"
    password    = ["fortinet"]
    description = "This is a Terraform test"
}

```

Before using this provider, the permission level for rpc-permit need to be set. See `Guides->To Set the Permission Level for RPC-Permit` for details.

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortianalyzer" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "true"

  scopetype    = "adom"
  adom         = "root"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.


## Authentication

The FortiAnalyzer provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by `username` and `password` parameters in the FortiAnalyzer provider block.

Usage:

```hcl
provider "fortianalyzer" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "true"

  scopetype    = "adom"
  adom         = "root"
}
```

### Token authentication

You can authenticate using API tokens instead of username/password. This method is more secure and recommended for production environments.

Usage with static token:

```hcl
provider "fortianalyzer" {
  hostname = "192.168.52.178"
  token    = "your-api-token-here"
  insecure = "true"

  scopetype = "adom"
  adom      = "root"
}
```

### Environment variables

You can provide your credentials via environment variables. For username/password authentication, use `FORTIANALYZER_ACCESS_HOSTNAME`, `FORTIANALYZER_ACCESS_USERNAME`, `FORTIANALYZER_ACCESS_PASSWORD`, `FORTIANALYZER_INSECURE` and `FORTIANALYZER_CA_CABUNDLE`. For token authentication, use `FORTIANALYZER_ACCESS_TOKEN`. Note that setting your FortiAnalyzer credentials using static credentials variables will override the environment variables.

Usage with username/password:

```shell
$ export "FORTIANALYZER_ACCESS_HOSTNAME"="192.168.52.178"
$ export "FORTIANALYZER_ACCESS_USERNAME"="admin"
$ export "FORTIANALYZER_ACCESS_PASSWORD"="admin"
$ export "FORTIANALYZER_INSECURE"="false"
$ export "FORTIANALYZER_CA_CABUNDLE"="/path/yourCA.crt"
```

Usage with token:

```shell
$ export "FORTIANALYZER_ACCESS_HOSTNAME"="192.168.52.178"
$ export "FORTIANALYZER_ACCESS_TOKEN"="your-api-token-here"
$ export "FORTIANALYZER_INSECURE"="true"
```

Then configure the FortiAnalyzer Provider as following:

```hcl
provider "fortianalyzer" {
  scopetype = "adom"
  adom      = "root"
}
```

## Argument Reference

The following arguments are supported:

* `hostname` - (Optional) The hostname or IP address of FortiAnalyzer unit. It must be provided, but it can also be sourced from the `FORTIANALYZER_ACCESS_HOSTNAME` environment variable.

* `username` - (Optional) Your username. It must be provided, but it can also be sourced from the `FORTIANALYZER_ACCESS_USERNAME` environment variable.

* `password` - (Optional) Your password. It must be provided, but it can also be sourced from the `FORTIANALYZER_ACCESS_PASSWORD` environment variable.

* `token` - (Optional) API token for authentication. When provided, username and password are not required. It can also be sourced from the `FORTIANALYZER_ACCESS_TOKEN` environment variable.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIANALYZER_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIANALYZER_CA_CABUNDLE` environment variable.

* `scopetype` - (Optional) The option is used to set the default scope of application of those resources managed by the provider. Valid values: `adom`, `global`. The default value is `adom`. Each resource can also set its own scope as needed, see the description of each resource for details.

* `adom` - (Optional) Adom. This value is valid only when the `scopetype` is set to `adom`. The option is used to set the default adom of the resources managed by the provider. The default value is `root`. Each resource can also set its own adom as needed, see the description of each resource for details.

* `import_options` - (Optional) This parameter is only used for import in some special cases. When the resource to be imported includes pkg parameter, you need to assign a value to the parameter here, for example:

    ```hcl
    provider "fortianalyzer" {
      hostname       = "192.168.52.178"
      username       = "admin"
      password       = "admin"
      insecure       = "true"

      scopetype      = "adom"
      adom           = "root"

      import_options = ["pkg=default"]
    }
    ```


## Release
Check out the FortiAnalyzer provider release notes and additional information from: [the FortiAnalyzer provider releases](https://github.com/fortinetdev/terraform-provider-fortianalyzer/releases).

## Versioning

The provider can cover FortiAnalyzer 7.0 versions, the configuration of all parameters should be based on the relevant FortiAnalyzer version manual and FortiAnalyzer API guide.
