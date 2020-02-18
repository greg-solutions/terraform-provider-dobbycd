package dobbycd

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestPipeline_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testDobbyProviders,
		Steps: []resource.TestStep{
			{
				Config: pipelineResource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("dobbycd_pipeline.project_pipeline", pathKey),
				),
			},
		},
	})
}
func TestResource(t *testing.T) {
	_ = resourceProjectCreate(nil, nil)
}

func pipelineResource() string {
	s := `
		resource "dobbycd_pipeline" "project_pipeline" {
			path = "/server/example/"
			token = ""
			branch = "develop"
			repository = "https://gitlab.com/gregsolutions/dobby-cd.git"
			project_id = "8ded5c04-3ae6-444d-9142-ef7b71370ca4"
		}
	`
	return s
}
