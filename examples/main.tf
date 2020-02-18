provider "dobbycd" {
  url = "http://localhost:8080/v1"
  username = "user"
  password = "user"
}

resource "dobbycd_project" "project_example" {
  name = "vadim-proj4"
}

resource "dobbycd_pipeline" "project_pipeline" {
  path = "/server/example/"
  token = "Jjabke6ybh9DeyxYzGe2"
  branch = "develop"
  repository = "https://gitlab.com/gregsolutions/dobby-cd.git"
  project_id = dobbycd_project.project_example.id
}


variable "permissions" {
  default = [
    {
      group_dn = "cn=ldap-user1,ou=RealmRoles,dc=example,dc=org"
      permit_type = "Admin"
    },

    {
      group_dn = "cn=ldap-user2,ou=RealmRoles,dc=example,dc=org"
      permit_type = "Admin"
    },

    {
      group_dn = "cn=ldap-user3,ou=RealmRoles,dc=example,dc=org"
      permit_type = "Admin"
    },

    {
      group_dn = "cn=ldap-user4,ou=RealmRoles,dc=example,dc=org"
      permit_type = "Admin"
    }
  ]
}

resource "dobbycd_global_permissions" "test_global_permissions_dynamic" {
  dynamic "permission" {
    for_each = var.permissions
    content {
      group_dn = permission.value.group_dn
      permit_type = permission.value.permit_type
    }
  }
}
