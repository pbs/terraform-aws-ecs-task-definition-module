module "task" {
  source = "../.."

  entrypoint = ["echo"]
  command    = ["hello"]

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
