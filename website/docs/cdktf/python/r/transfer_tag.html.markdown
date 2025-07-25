---
subcategory: "Transfer Family"
layout: "aws"
page_title: "AWS: aws_transfer_tag"
description: |-
  Manages an individual Transfer Family resource tag
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_transfer_tag

Manages an individual Transfer Family resource tag. This resource should only be used in cases where Transfer Family resources are created outside Terraform (e.g., Servers without AWS Management Console) or the tag key has the `aws:` prefix.

~> **NOTE:** This tagging resource should not be combined with the Terraform resource for managing the parent resource. For example, using `aws_transfer_server` and `aws_transfer_tag` to manage tags of the same server will cause a perpetual difference where the `aws_transfer_server` resource will try to remove the tag being added by the `aws_transfer_tag` resource.

~> **NOTE:** This tagging resource does not use the [provider `ignore_tags` configuration](/docs/providers/aws/index.html#ignore_tags).

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.transfer_server import TransferServer
from imports.aws.transfer_tag import TransferTag
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = TransferServer(self, "example",
            identity_provider_type="SERVICE_MANAGED"
        )
        TransferTag(self, "hostname",
            key="transfer:customHostname",
            resource_arn=example.arn,
            value="example.com"
        )
        TransferTag(self, "zone_id",
            key="transfer:route53HostedZoneId",
            resource_arn=example.arn,
            value="/hostedzone/MyHostedZoneId"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resource_arn` - (Required) Amazon Resource Name (ARN) of the Transfer Family resource to tag.
* `key` - (Required) Tag name.
* `value` - (Required) Tag value.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Transfer Family resource identifier and key, separated by a comma (`,`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_transfer_tag` using the Transfer Family resource identifier and key, separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.transfer_tag import TransferTag
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        TransferTag.generate_config_for_import(self, "example", "arn:aws:transfer:us-east-1:123456789012:server/s-1234567890abcdef0,Name")
```

Using `terraform import`, import `aws_transfer_tag` using the Transfer Family resource identifier and key, separated by a comma (`,`). For example:

```console
% terraform import aws_transfer_tag.example arn:aws:transfer:us-east-1:123456789012:server/s-1234567890abcdef0,Name
```

<!-- cache-key: cdktf-0.20.8 input-dbff2d9c3045148139dd50b19e3b850f040e43d426f0e86074715746bb3c41ac -->