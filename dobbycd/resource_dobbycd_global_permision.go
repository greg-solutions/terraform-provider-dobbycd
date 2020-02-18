package dobbycd

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client"
	"gitlab.com/gregsolutions/dobby-cd/api"
)

const (
	permissionsKey = "permission"
	groupKey       = "group_dn"
	permitTypeKey  = "permit_type"
)

func resourceGlobalPermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlobalPermissionCreate,
		Read:   resourceGlobalPermissionRead,
		Update: resourceGlobalPermissionUpdate,
		Delete: resourceGlobalPermissionDelete,

		Schema: map[string]*schema.Schema{
			permissionsKey: {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						groupKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						permitTypeKey: {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceGlobalPermissionCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)
	_ = client
	d.SetId(uuid.New().String())
	return nil
}

func resourceGlobalPermissionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.DobbyCDApi)
	perm, err := client.GetGlobalPermissions()
	if err != nil {
		return err
	}
	_ = perm
	return nil
}

func resourceGlobalPermissionUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceGlobalPermissionDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}
