variable "service_name" {
  description = "(optional) name of the service running this task. Only important here because the AWS console defaults to `/ecs/service_name` when displaying logs for a service"
  default     = null
  type        = string
}

variable "log_group_name" {
  description = "(optional) name for the log group"
  default     = null
  type        = string
}

variable "retention_in_days" {
  description = "(optional) log retention in days"
  default     = 7
  type        = number
}
