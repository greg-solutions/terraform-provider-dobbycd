package dobbycd

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client/model"
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
	permissions := d.Get(permissionsKey)
	perms := make([]*model.Permit, 0)

	err := mapstructure.Decode(permissions, &perms)
	if err != nil {
		return err
	}
	resp, err := client.SetGlobalPermissions(perms)
	if err != nil {
		return err
	}
	_ = resp
	d.SetId(uuid.New().String())
	_ = d.Set(permissionsKey, permissions)
	return nil
}

func resourceGlobalPermissionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)
	perm, err := client.GetGlobalPermissions()
	if err != nil {
		return err
	}
	_ = perm
	perms := d.Get(permissionsKey)

	_ = d.Set(permissionsKey, perms)
	return nil
}

func resourceGlobalPermissionUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceGlobalPermissionCreate(d, m)
}

func resourceGlobalPermissionDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.DobbyCDApi)

	perms := make([]*model.Permit, 0)
	_, err := client.SetGlobalPermissions(perms)
	if err != nil {
		return err
	}
	return nil
}
