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
  token = ""
  branch = "develop"
  repository = "https://gitlab.com/gregsolutions/dobby-cd.git"
  project_id = dobbycd_project.project_example.id
}

resource "dobbycd_global_permissions" "test_global_permissions" {
  permission {
    group_dn = "cn=ldap-user,ou=RealmRoles,dc=example,dc=org"
    permit_type = "Admin"
  }
}
