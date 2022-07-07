# PBS TF ECS Task Definition Module

## Installation

### Using the Repo Source

```hcl
module "task" {
    source = "github.com/pbs/terraform-aws-ecs-task-definition-module?ref=0.0.2"
}
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

This provisions a task definition for use with an ECS service.

Most of the time, you shouldn't have to use this module directly. It is integrated into the [ECS service module][ecs-service-module], allowing you to specify most important configurations that are relevant to your service there unless you have very particular needs for your task definition.

Integrate this module like so:

```hcl
module "task" {
  source = "github.com/pbs/terraform-aws-ecs-task-definition-module?ref=0.0.2"

  # Tagging Parameters
  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo

  # Optional Parameters
  image_repo = "nginx"
  image_tag = "latest"
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`0.0.2`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs
[ecs-service-module]: https://github.com/pbs/terraform-aws-ecs-service-module

## Requirements

| Name                                                                     | Version  |
| ------------------------------------------------------------------------ | -------- |
| <a name="requirement_terraform"></a> [terraform](#requirement_terraform) | >= 1.1.7 |
| <a name="requirement_aws"></a> [aws](#requirement_aws)                   | >= 4.5.0 |

## Providers

| Name                                             | Version |
| ------------------------------------------------ | ------- |
| <a name="provider_aws"></a> [aws](#provider_aws) | 4.5.0   |

## Modules

No modules.

## Resources

| Name                                                                                                                                                         | Type        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------- |
| [aws_cloudwatch_log_group.logs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group)                            | resource    |
| [aws_ecs_task_definition.task_def](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_task_definition)                          | resource    |
| [aws_iam_role.task_execution_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role)                                     | resource    |
| [aws_iam_role.task_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role)                                               | resource    |
| [aws_iam_role_policy.task_execution_role_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy)                | resource    |
| [aws_iam_role_policy.task_role_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy)                          | resource    |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity)                                | data source |
| [aws_iam_policy_document.assume_role_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document)             | data source |
| [aws_iam_policy_document.policy_doc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document)                     | data source |
| [aws_iam_policy_document.task_execution_role_policy_doc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_iam_policy_document.vgw_policy_doc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document)                 | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region)                                                  | data source |
| [aws_secretsmanager_secret.newrelic_secret](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/secretsmanager_secret)            | data source |

## Inputs

| Name                                                                                                                           | Description                                                                                                                                                      | Type                                                                                                          | Default                                                                              | Required |
| ------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ | :------: |
| <a name="input_environment"></a> [environment](#input_environment)                                                             | Environment (sharedtools, dev, staging, prod)                                                                                                                    | `string`                                                                                                      | n/a                                                                                  |   yes    |
| <a name="input_organization"></a> [organization](#input_organization)                                                          | Organization using this module. Used to prefix tags so that they are easily identified as being from your organization                                           | `string`                                                                                                      | n/a                                                                                  |   yes    |
| <a name="input_product"></a> [product](#input_product)                                                                         | Tag used to group resources according to product                                                                                                                 | `string`                                                                                                      | n/a                                                                                  |   yes    |
| <a name="input_repo"></a> [repo](#input_repo)                                                                                  | Tag used to point to the repo using this module                                                                                                                  | `string`                                                                                                      | n/a                                                                                  |   yes    |
| <a name="input_command"></a> [command](#input_command)                                                                         | (optional) command to run in the container as an array. e.g. ["sleep", "10"]. If null, does not set a command in the task definition.                            | `list(string)`                                                                                                | `null`                                                                               |    no    |
| <a name="input_container_definitions"></a> [container_definitions](#input_container_definitions)                               | (optional) JSON container definitions for task                                                                                                                   | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_container_name"></a> [container_name](#input_container_name)                                                    | (optional) name for the container to have                                                                                                                        | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_container_port"></a> [container_port](#input_container_port)                                                    | (optional) port the container is exposing                                                                                                                        | `number`                                                                                                      | `80`                                                                                 |    no    |
| <a name="input_cpu_reservation"></a> [cpu_reservation](#input_cpu_reservation)                                                 | (optional) CPU reservation for task                                                                                                                              | `number`                                                                                                      | `256`                                                                                |    no    |
| <a name="input_efs_mounts"></a> [efs_mounts](#input_efs_mounts)                                                                | (optional) efs mount set of objects. Components should include dns_name, container_mount_point, efs_mount_point                                                  | <pre>set(object({<br> file_system_id = string<br> efs_path = string<br> container_path = string<br> }))</pre> | `[]`                                                                                 |    no    |
| <a name="input_entrypoint"></a> [entrypoint](#input_entrypoint)                                                                | (optional) entrypoint to run in the container as an array. e.g. ["sleep", "10"]. If null, does not set an entrypoint in the task definition.                     | `list(string)`                                                                                                | `null`                                                                               |    no    |
| <a name="input_env_vars"></a> [env_vars](#input_env_vars)                                                                      | (optional) environment variables to be passed to the container. By default, only passes SSM_PATH                                                                 | `set(map(any))`                                                                                               | `null`                                                                               |    no    |
| <a name="input_envoy_tag"></a> [envoy_tag](#input_envoy_tag)                                                                   | (optional) tag for envoy. Update periodically if using App Mesh.                                                                                                 | `string`                                                                                                      | `"v1.18.3.0-prod"`                                                                   |    no    |
| <a name="input_image_repo"></a> [image_repo](#input_image_repo)                                                                | (optional) image repo. e.g. image_repo = nginx --> nginx:image_tag                                                                                               | `string`                                                                                                      | `"nginx"`                                                                            |    no    |
| <a name="input_image_tag"></a> [image_tag](#input_image_tag)                                                                   | (optional) tag of the image. e.g. image_tag = latest --> image_repo:latest                                                                                       | `string`                                                                                                      | `"alpine"`                                                                           |    no    |
| <a name="input_log_group_name"></a> [log_group_name](#input_log_group_name)                                                    | (optional) name for the log group                                                                                                                                | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_memory_reservation"></a> [memory_reservation](#input_memory_reservation)                                        | (optional) memory reservation for task                                                                                                                           | `number`                                                                                                      | `512`                                                                                |    no    |
| <a name="input_mesh_name"></a> [mesh_name](#input_mesh_name)                                                                   | (optional) the name for the App Mesh this task is associated with. If null, ignored                                                                              | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_name"></a> [name](#input_name)                                                                                  | Name of the Ecs Task Definition Module. If null, will default to product.                                                                                        | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_network_mode"></a> [network_mode](#input_network_mode)                                                          | (optional) network mode for the task                                                                                                                             | `string`                                                                                                      | `"awsvpc"`                                                                           |    no    |
| <a name="input_newrelic_firelens_image"></a> [newrelic_firelens_image](#input_newrelic_firelens_image)                         | Name for AWS Secrets Manager secret of New Relic Insights insert key.                                                                                            | `string`                                                                                                      | `"533243300146.dkr.ecr.us-east-1.amazonaws.com/newrelic/logging-firelens-fluentbit"` |    no    |
| <a name="input_newrelic_secret_arn"></a> [newrelic_secret_arn](#input_newrelic_secret_arn)                                     | ARN for AWS Secrets Manager secret of New Relic Insights insert key.                                                                                             | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_newrelic_secret_name"></a> [newrelic_secret_name](#input_newrelic_secret_name)                                  | Name for AWS Secrets Manager secret of New Relic Insights insert key.                                                                                            | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_requires_compatibilities"></a> [requires_compatibilities](#input_requires_compatibilities)                      | (optional) capabilities that the task requires                                                                                                                   | `set(string)`                                                                                                 | <pre>[<br> "FARGATE"<br>]</pre>                                                      |    no    |
| <a name="input_retention_in_days"></a> [retention_in_days](#input_retention_in_days)                                           | (optional) log retention in days                                                                                                                                 | `number`                                                                                                      | `7`                                                                                  |    no    |
| <a name="input_role_policy_json"></a> [role_policy_json](#input_role_policy_json)                                              | (optional) IAM policy to attach to role used for this task                                                                                                       | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_service_name"></a> [service_name](#input_service_name)                                                          | (optional) name of the service running this task. Only important here because the AWS console defaults to `/ecs/service_name` when displaying logs for a service | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_ssm_path"></a> [ssm_path](#input_ssm_path)                                                                      | (optional) path to the ssm parameters you want pulled into your container during execution of the entrypoint                                                     | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_tags"></a> [tags](#input_tags)                                                                                  | Extra tags                                                                                                                                                       | `map(string)`                                                                                                 | `{}`                                                                                 |    no    |
| <a name="input_task_execution_role_policy_json"></a> [task_execution_role_policy_json](#input_task_execution_role_policy_json) | (optional) IAM policy to attach to task execution role used for this task                                                                                        | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_task_family"></a> [task_family](#input_task_family)                                                             | (optional) task family for task. This is effectively the name of the task, without qualification of revision                                                     | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_use_xray_sidecar"></a> [use_xray_sidecar](#input_use_xray_sidecar)                                              | (optional) if set to null, will use the sidecar to trace the task if envoy is used, as that automatically implements tracing configs.                            | `bool`                                                                                                        | `null`                                                                               |    no    |
| <a name="input_virtual_gateway"></a> [virtual_gateway](#input_virtual_gateway)                                                 | (optional) the name of the virtual gateway associated with this task definition. If null, ignored                                                                | `string`                                                                                                      | `null`                                                                               |    no    |
| <a name="input_virtual_node"></a> [virtual_node](#input_virtual_node)                                                          | (optional) the name of the virtual node associated with this task definition. Ignored if virtual_gateway set. If null, ignored                                   | `string`                                                                                                      | `null`                                                                               |    no    |

## Outputs

| Name                                                                                               | Description                           |
| -------------------------------------------------------------------------------------------------- | ------------------------------------- |
| <a name="output_arn"></a> [arn](#output_arn)                                                       | Task definition ARN                   |
| <a name="output_container_definitions"></a> [container_definitions](#output_container_definitions) | Task definition container definitions |
| <a name="output_role_arn"></a> [role_arn](#output_role_arn)                                        | IAM role ARN                          |
