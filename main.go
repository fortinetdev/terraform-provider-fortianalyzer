package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-providers/terraform-provider-fortianalyzer/faz"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortianalyzer.Provider})
}
