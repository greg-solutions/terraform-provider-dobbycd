package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/vadimDidenko/terraform-provider-jenkins/dobbycd"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dobbycd.Provider})
}
