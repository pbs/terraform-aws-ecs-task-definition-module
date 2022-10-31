variable "role_policy_json" {
  description = "(optional) IAM policy to attach to role used for this task"
  default     = null
  type        = string
}

variable "task_execution_role_policy_json" {
  description = "(optional) IAM policy to attach to task execution role used for this task"
  default     = null
  type        = string
}

variable "runtime_platform" {
  description = "(optional) Runtime platform for the task. Defaults to LINUX operating system family w/ CPU architecture x86_64."
  default = {
    operating_system_family = "LINUX"
    cpu_architecture        = "x86_64"
  }
  type = object({
    operating_system_family = optional(string, "LINUX")
    cpu_architecture        = optional(string, "x86_64")
  })
}
