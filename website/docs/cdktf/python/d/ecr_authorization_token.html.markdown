---
subcategory: "ECR (Elastic Container Registry)"
layout: "aws"
page_title: "AWS: aws_ecr_authorization_token"
description: |-
    Provides details about an ECR Authorization Token
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ecr_authorization_token

The ECR Authorization Token data source allows the authorization token, proxy endpoint, token expiration date, user name and password to be retrieved for an ECR repository.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ecr_authorization_token import DataAwsEcrAuthorizationToken
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsEcrAuthorizationToken(self, "token")
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `registry_id` - (Optional) AWS account ID of the ECR Repository. If not specified the default account is assumed.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `authorization_token` - Temporary IAM authentication credentials to access the ECR repository encoded in base64 in the form of `user_name:password`.
* `expires_at` - Time in UTC RFC3339 format when the authorization token expires.
* `id` - Region of the authorization token.
* `password` - Password decoded from the authorization token.
* `proxy_endpoint` - Registry URL to use in the docker login command.
* `user_name` - User name decoded from the authorization token.

<!-- cache-key: cdktf-0.20.8 input-428ea462ea1aa068d61da6676b5421245994a7db5fc507fa86411952d346563e -->