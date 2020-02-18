package dobbycd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client"
)

const (
	nameKey = "name"
)

func resourceProjectJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			nameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)
	name := d.Get(nameKey).(string)

	p, err := client.CreateProject(name)
	if err != nil {
		return err
	}

	d.SetId(p.ID)
	_ = d.Set(nameKey, p.Name)

	_ = resourceProjectRead(d, m)
	return nil
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)
	projectId := d.Id()

	p, err := client.GetProject(projectId)
	if err != nil {
		return err
	}
	_ = d.Set(nameKey, p.Name)
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)
	projectId := d.Id()

	name := d.Get(nameKey).(string)
	p, err := client.UpdateProject(projectId, name)
	if err != nil {
		return err
	}
	_ = d.Set(nameKey, p.Name)
	return nil
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)
	projectId := d.Id()

	err := client.DeleteProject(projectId)
	if err != nil {
		return err
	}
	return nil
}
