output "arn" {
  description = "Task definition ARN"
  value       = module.task.arn
}

output "role_arn" {
  description = "IAM role ARN"
  value       = module.task.role_arn
}

output "container_definitions" {
  description = "Task definition container definitions"
  value       = module.task.container_definitions
}

output "efs1_id" {
  description = "EFS file system 1 ID"
  value       = module.efs1.id
}

output "efs2_id" {
  description = "EFS file system 2 ID"
  value       = module.efs2.id
}
