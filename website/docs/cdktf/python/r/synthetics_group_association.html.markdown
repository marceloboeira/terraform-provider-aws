---
subcategory: "CloudWatch Synthetics"
layout: "aws"
page_title: "AWS: aws_synthetics_group_association"
description: |-
  Provides a Synthetics Group Association resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_synthetics_group_association

Provides a Synthetics Group Association resource.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.synthetics_group_association import SyntheticsGroupAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SyntheticsGroupAssociation(self, "example",
            canary_arn=Token.as_string(aws_synthetics_canary_example.arn),
            group_name=Token.as_string(aws_synthetics_group_example.name)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `group_name` - (Required) Name of the group that the canary will be associated with.
* `canary_arn` - (Required) ARN of the canary.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `group_name` - Name of the Group.
* `group_id` - ID of the Group.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CloudWatch Synthetics Group Association using the `canary_arn,group_name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.synthetics_group_association import SyntheticsGroupAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SyntheticsGroupAssociation.generate_config_for_import(self, "example", "arn:aws:synthetics:us-west-2:123456789012:canary:tf-acc-test-abcd1234,examplename")
```

Using `terraform import`, import CloudWatch Synthetics Group Association using the `canary_arn,group_name`. For example:

```console
% terraform import aws_synthetics_group_association.example arn:aws:synthetics:us-west-2:123456789012:canary:tf-acc-test-abcd1234,examplename
```

<!-- cache-key: cdktf-0.20.8 input-2c5ab784846b3e6de6ec8e8722c24c9a4e4637014d43ab79bfe6f754d98bed8e -->