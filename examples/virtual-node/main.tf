module "task" {
  source = "../.."

  mesh_name    = "mesh_name"
  virtual_node = "virtual_node"

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
