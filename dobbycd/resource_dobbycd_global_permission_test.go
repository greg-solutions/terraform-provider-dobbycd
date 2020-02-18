package dobbycd

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestGlobalPermission_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testDobbyProviders,
		Steps: []resource.TestStep{
			{
				Config: globalPermissionResource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("dobbycd_global_permissions.test_global_permissions", "permissions"),
				),
			},
		},
	})
}
func TestResourceGlobalPermission(t *testing.T) {
	_ = resourceGlobalPermissionCreate(nil, nil)
}

func globalPermissionResource() string {
	s := `
		resource "dobbycd_global_permissions" "test_global_permissions" {
		  permission {
			  group_dn = "cn=ldap-user,ou=RealmRoles,dc=example,dc=org"
			  permit_type = "Admin"
			}
		}
	`
	return s
}
