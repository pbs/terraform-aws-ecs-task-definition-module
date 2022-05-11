# PBS TF ecs task definition module

## Installation

### Using the Repo Source

```hcl
module "task" {
    source = "github.com/pbs/terraform-aws-ecs-task-definition-module?ref=x.y.z"
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
  source = "github.com/pbs/terraform-aws-ecs-task-definition-module?ref=x.y.z"

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

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs
[ecs-service-module]: https://github.com/pbs/terraform-aws-ecs-service-module
