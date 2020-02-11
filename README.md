# Terraform Dobby CD provider

Terraform allows you to keep your infrastructure as a code (HCL language).
This provider supports controlling dobbycd over its http api.

## Configuration of provider

```hcl
provider "dobbycd" {
  url = "http://localhost:8080/v1"
  username = "user"
  password = "user"
}
```

##Project
```hcl
resource "dobbycd_project" "project_example" {
  name = "project_example"
}
```