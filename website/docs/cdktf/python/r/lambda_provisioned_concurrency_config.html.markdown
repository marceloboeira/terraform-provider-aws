---
subcategory: "Lambda"
layout: "aws"
page_title: "AWS: aws_lambda_provisioned_concurrency_config"
description: |-
  Manages an AWS Lambda Provisioned Concurrency Configuration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lambda_provisioned_concurrency_config

Manages an AWS Lambda Provisioned Concurrency Configuration. Use this resource to configure provisioned concurrency for Lambda functions.

~> **Note:** Setting `skip_destroy` to `true` means that the AWS Provider will not destroy a provisioned concurrency configuration, even when running `terraform destroy`. The configuration is thus an intentional dangling resource that is not managed by Terraform and may incur extra expense in your AWS account.

## Example Usage

### Alias Name

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lambda_provisioned_concurrency_config import LambdaProvisionedConcurrencyConfig
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        LambdaProvisionedConcurrencyConfig(self, "example",
            function_name=Token.as_string(aws_lambda_alias_example.function_name),
            provisioned_concurrent_executions=1,
            qualifier=Token.as_string(aws_lambda_alias_example.name)
        )
```

### Function Version

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lambda_provisioned_concurrency_config import LambdaProvisionedConcurrencyConfig
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        LambdaProvisionedConcurrencyConfig(self, "example",
            function_name=Token.as_string(aws_lambda_function_example.function_name),
            provisioned_concurrent_executions=1,
            qualifier=Token.as_string(aws_lambda_function_example.version)
        )
```

## Argument Reference

The following arguments are required:

* `function_name` - (Required) Name or Amazon Resource Name (ARN) of the Lambda Function.
* `provisioned_concurrent_executions` - (Required) Amount of capacity to allocate. Must be greater than or equal to 1.
* `qualifier` - (Required) Lambda Function version or Lambda Alias name.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `skip_destroy` - (Optional) Whether to retain the provisioned concurrency configuration upon destruction. Defaults to `false`. If set to `true`, the resource is simply removed from state instead.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Lambda Function name and qualifier separated by a comma (`,`).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `15m`)
* `update` - (Default `15m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a Lambda Provisioned Concurrency Configuration using the `function_name` and `qualifier` separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lambda_provisioned_concurrency_config import LambdaProvisionedConcurrencyConfig
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        LambdaProvisionedConcurrencyConfig.generate_config_for_import(self, "example", "example,production")
```

Using `terraform import`, import a Lambda Provisioned Concurrency Configuration using the `function_name` and `qualifier` separated by a comma (`,`). For example:

```console
% terraform import aws_lambda_provisioned_concurrency_config.example example,production
```

<!-- cache-key: cdktf-0.20.8 input-9e96be72ca128434e95b56fd02de74603cea7d34217500c43a5176e8976a8232 -->