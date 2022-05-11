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
