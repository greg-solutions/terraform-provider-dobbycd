package dobbycd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.com/gregsolutions/dobby-cd/api"
)

func resourceProjectJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.DobbyCDApi)
	name := d.Get("name").(string)

	p, err := client.CreateProject(name)
	if err != nil {
		return err
	}

	d.SetId(p.ID)
	_ = d.Set("name", p.Name)

	_ = resourceProjectRead(d, m)
	return nil
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.DobbyCDApi)
	projectId := d.Id()
	
	p, err := client.GetProject(projectId)
	if err != nil {
		return err
	}
	_ = d.Set("name", p.Name)
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.DobbyCDApi)
	projectId := d.Id()

	name := d.Get("name").(string)
	p, err := client.UpdateProject(projectId,name)
	if err != nil {
		return err
	}
	_ = d.Set("name", p.Name)
	return nil
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.DobbyCDApi)
	projectId := d.Id()

	 err := client.DeleteProject(projectId)
	if err != nil {
		return err
	}
	return nil
}
