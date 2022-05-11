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
