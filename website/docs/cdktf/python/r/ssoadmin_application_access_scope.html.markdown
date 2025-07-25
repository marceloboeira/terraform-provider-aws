---
subcategory: "SSO Admin"
layout: "aws"
page_title: "AWS: aws_ssoadmin_application_access_scope"
description: |-
  Terraform resource for managing an AWS SSO Admin Application Access Scope.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssoadmin_application_access_scope

Terraform resource for managing an AWS SSO Admin Application Access Scope.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ssoadmin_instances import DataAwsSsoadminInstances
from imports.aws.ssoadmin_application import SsoadminApplication
from imports.aws.ssoadmin_application_access_scope import SsoadminApplicationAccessScope
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsSsoadminInstances(self, "example")
        aws_ssoadmin_application_example = SsoadminApplication(self, "example_1",
            application_provider_arn="arn:aws:sso::aws:applicationProvider/custom",
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.arns), ["0"])),
            name="example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssoadmin_application_example.override_logical_id("example")
        aws_ssoadmin_application_access_scope_example =
        SsoadminApplicationAccessScope(self, "example_2",
            application_arn=Token.as_string(aws_ssoadmin_application_example.arn),
            authorized_targets=["arn:aws:sso::123456789012:application/ssoins-123456789012/apl-123456789012"
            ],
            scope="sso:account:access"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssoadmin_application_access_scope_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `application_arn` - (Required) Specifies the ARN of the application with the access scope with the targets to add or update.
* `scope` - (Required) Specifies the name of the access scope to be associated with the specified targets.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `authorized_targets` - (Optional) Specifies an array list of ARNs that represent the authorized targets for this access scope.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - A comma-delimited string concatenating `application_arn` and `scope`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SSO Admin Application Access Scope using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ssoadmin_application_access_scope import SsoadminApplicationAccessScope
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SsoadminApplicationAccessScope.generate_config_for_import(self, "example", "arn:aws:sso::123456789012:application/ssoins-123456789012/apl-123456789012,sso:account:access")
```

Using `terraform import`, import SSO Admin Application Access Scope using the `id`. For example:

```console
% terraform import aws_ssoadmin_application_access_scope.example arn:aws:sso::123456789012:application/ssoins-123456789012/apl-123456789012,sso:account:access
```

<!-- cache-key: cdktf-0.20.8 input-f3dec092d86f98863e4857dc79555ffc7d6e54579d8186572e056c8c788a15c6 -->