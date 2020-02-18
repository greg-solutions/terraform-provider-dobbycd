package dobbycd

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

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

func TestProject_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testDobbyProviders,
		Steps: []resource.TestStep{
			{
				Config: projectResource("test11"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("dobbycd_project.project", nameKey),
				),
			},
		},
	})
}
func TestResourceProjectJob(t *testing.T) {
	_ = resourceProjectCreate(nil, nil)
}

func projectResource(name string) string {
	return fmt.Sprintf(basicConfig(), name)
}

func basicConfig() string {
	return `
resource "dobbycd_project" "project" {
  name = "%s"
}
`
}
