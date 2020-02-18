package dobbycd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client/model"
)

const (
	repositoryKey = "repository"
	projectIdKey  = "project_id"
	branchKey     = "branch"
	tokenKey      = "token"
	pathKey       = "path"
)

func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

		Schema: map[string]*schema.Schema{
			repositoryKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			branchKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			tokenKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			pathKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			projectIdKey: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePipelineCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)
	repository := d.Get(repositoryKey).(string)
	branch := d.Get(branchKey).(string)
	token := d.Get(tokenKey).(string)
	path := d.Get(pathKey).(string)
	projectId := d.Get(projectIdKey).(string)

	data := &model.PipelineRequest{
		Path:       path,
		Token:      token,
		Branch:     branch,
		Repository: repository,
	}
	p, err := client.CreatePipeline(projectId, data)
	if err != nil {
		return err
	}

	d.SetId(p.PipelineId)
	_ = d.Set(repositoryKey, repository)
	_ = d.Set(branchKey, branch)
	_ = d.Set(tokenKey, token)
	_ = d.Set(pathKey, path)

	_ = resourceProjectRead(d, m)
	return nil
}

func resourcePipelineRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}
