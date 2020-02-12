package dobbycd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)


var testDobbyProviders map[string]terraform.ResourceProvider
var testDobbyProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testDobbyProvider *schema.Provider
var testDobbyProviderFunc func() *schema.Provider


func init() {
	testDobbyProvider = Provider().(*schema.Provider)
	testDobbyProviders = map[string]terraform.ResourceProvider{
		"dobbycd": testDobbyProvider,
	}
	testDobbyProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"dobbycd": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
	testDobbyProviderFunc = func() *schema.Provider { return testDobbyProvider }
}

