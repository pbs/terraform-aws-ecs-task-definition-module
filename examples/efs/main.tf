module "efs1" {
  source = "github.com/pbs/terraform-aws-efs-module?ref=0.0.1"

  organization = var.organization
  environment  = var.environment
  product      = "${var.product}1"
  repo         = var.repo
}

module "efs2" {
  source = "github.com/pbs/terraform-aws-efs-module?ref=0.0.1"

  organization = var.organization
  environment  = var.environment
  product      = "${var.product}2"
  repo         = var.repo
}

module "task" {
  source = "../.."

  efs_mounts = [
    {
      file_system_id = module.efs1.id
      efs_path       = "/"
      container_path = "/mnt/efs1"
    },
    {
      file_system_id = module.efs2.id
      efs_path       = "/"
      container_path = "/mnt/efs2"
    }
  ]

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
