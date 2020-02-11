provider "dobbycd" {
  url = "http://localhost:8080/v1"
  username = "user"
  password = "user"
}

resource "dobbycd_project" "project_example" {
  name = "vadim-proj4"
}