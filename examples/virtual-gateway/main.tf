module "task" {
  source = "../.."

  mesh_name       = "mesh_name"
  virtual_gateway = "virtual_gateway"

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
