package dobbycd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client"
)

func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://localhost:8080/v1",
				Description: "host address of dobbycd instance",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOBBYCD_USERNAME", ""),
				Description: "username which should be used to loginto instance",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOBBYCD_PASSWORD", ""),
				Description: "password which should be used to login to instance",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dobbycd_project":            resourceProjectJob(),
			"dobbycd_pipeline":           resourcePipeline(),
			"dobbycd_global_permissions": resourceGlobalPermissions(),
		},
		ConfigureFunc: configureFunc,
	}
}

func configureFunc(rd *schema.ResourceData) (interface{}, error) {

	config := DobbyCdConfig{
		URL:      rd.Get("url").(string),
		Username: rd.Get("username").(string),
		Password: rd.Get("password").(string),
	}

	capi := client.NewApi(config.URL, config.Username, config.Password)
	return capi, nil
}

type DobbyCdConfig struct {
	URL      string
	Username string
	Password string
}
