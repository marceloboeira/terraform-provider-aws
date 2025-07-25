---
subcategory: "SageMaker AI"
layout: "aws"
page_title: "AWS: aws_sagemaker_notebook_instance_lifecycle_configuration"
description: |-
  Provides a lifecycle configuration for SageMaker AI Notebook Instances.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sagemaker_notebook_instance_lifecycle_configuration

Provides a lifecycle configuration for SageMaker AI Notebook Instances.

## Example Usage

Usage:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sagemaker_notebook_instance_lifecycle_configuration import SagemakerNotebookInstanceLifecycleConfiguration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SagemakerNotebookInstanceLifecycleConfiguration(self, "lc",
            name="foo",
            on_create=Token.as_string(Fn.base64encode("echo foo")),
            on_start=Token.as_string(Fn.base64encode("echo bar"))
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Optional) The name of the lifecycle configuration (must be unique). If omitted, Terraform will assign a random, unique name.
* `on_create` - (Optional) A shell script (base64-encoded) that runs only once when the SageMaker AI Notebook Instance is created.
* `on_start` - (Optional) A shell script (base64-encoded) that runs every time the SageMaker AI Notebook Instance is started including the time it's created.
* `tags` - (Optional) A mapping of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import models using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sagemaker_notebook_instance_lifecycle_configuration import SagemakerNotebookInstanceLifecycleConfiguration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SagemakerNotebookInstanceLifecycleConfiguration.generate_config_for_import(self, "lc", "foo")
```

Using `terraform import`, import models using the `name`. For example:

```console
% terraform import aws_sagemaker_notebook_instance_lifecycle_configuration.lc foo
```

<!-- cache-key: cdktf-0.20.8 input-eaa75171c1e8ec6bf10d5608f06924201a24c18d88774cda961ad5311596b7db -->