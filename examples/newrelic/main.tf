module "task" {
  source = "../.."

  newrelic_secret_arn = "arn:aws:secretsmanager:*:*:secret:fake-newrelic-secret-arn"

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
