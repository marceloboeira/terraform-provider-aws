---
subcategory: "EventBridge"
layout: "aws"
page_title: "AWS: aws_cloudwatch_event_connection"
description: |-
  Provides an EventBridge connection resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cloudwatch_event_connection

Provides an EventBridge connection resource.

~> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_event_connection import CloudwatchEventConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchEventConnection(self, "test",
            auth_parameters=CloudwatchEventConnectionAuthParameters(
                api_key=CloudwatchEventConnectionAuthParametersApiKey(
                    key="x-signature",
                    value="1234"
                )
            ),
            authorization_type="API_KEY",
            description="A connection description",
            name="ngrok-connection"
        )
```

## Example Usage Basic Authorization

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_event_connection import CloudwatchEventConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchEventConnection(self, "test",
            auth_parameters=CloudwatchEventConnectionAuthParameters(
                basic=CloudwatchEventConnectionAuthParametersBasic(
                    password="Pass1234!",
                    username="user"
                )
            ),
            authorization_type="BASIC",
            description="A connection description",
            name="ngrok-connection"
        )
```

## Example Usage OAuth Authorization

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_event_connection import CloudwatchEventConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchEventConnection(self, "test",
            auth_parameters=CloudwatchEventConnectionAuthParameters(
                oauth=CloudwatchEventConnectionAuthParametersOauth(
                    authorization_endpoint="https://auth.url.com/endpoint",
                    client_parameters=CloudwatchEventConnectionAuthParametersOauthClientParameters(
                        client_id="1234567890",
                        client_secret="Pass1234!"
                    ),
                    http_method="GET",
                    oauth_http_parameters=CloudwatchEventConnectionAuthParametersOauthOauthHttpParameters(
                        body=[CloudwatchEventConnectionAuthParametersOauthOauthHttpParametersBody(
                            is_value_secret=False,
                            key="body-parameter-key",
                            value="body-parameter-value"
                        )
                        ],
                        header=[CloudwatchEventConnectionAuthParametersOauthOauthHttpParametersHeader(
                            is_value_secret=False,
                            key="header-parameter-key",
                            value="header-parameter-value"
                        )
                        ],
                        query_string=[CloudwatchEventConnectionAuthParametersOauthOauthHttpParametersQueryString(
                            is_value_secret=False,
                            key="query-string-parameter-key",
                            value="query-string-parameter-value"
                        )
                        ]
                    )
                )
            ),
            authorization_type="OAUTH_CLIENT_CREDENTIALS",
            description="A connection description",
            name="ngrok-connection"
        )
```

## Example Usage Invocation Http Parameters

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_event_connection import CloudwatchEventConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchEventConnection(self, "test",
            auth_parameters=CloudwatchEventConnectionAuthParameters(
                basic=CloudwatchEventConnectionAuthParametersBasic(
                    password="Pass1234!",
                    username="user"
                ),
                invocation_http_parameters=CloudwatchEventConnectionAuthParametersInvocationHttpParameters(
                    body=[CloudwatchEventConnectionAuthParametersInvocationHttpParametersBody(
                        is_value_secret=False,
                        key="body-parameter-key",
                        value="body-parameter-value"
                    ), CloudwatchEventConnectionAuthParametersInvocationHttpParametersBody(
                        is_value_secret=True,
                        key="body-parameter-key2",
                        value="body-parameter-value2"
                    )
                    ],
                    header=[CloudwatchEventConnectionAuthParametersInvocationHttpParametersHeader(
                        is_value_secret=False,
                        key="header-parameter-key",
                        value="header-parameter-value"
                    )
                    ],
                    query_string=[CloudwatchEventConnectionAuthParametersInvocationHttpParametersQueryString(
                        is_value_secret=False,
                        key="query-string-parameter-key",
                        value="query-string-parameter-value"
                    )
                    ]
                )
            ),
            authorization_type="BASIC",
            description="A connection description",
            name="ngrok-connection"
        )
```

## Example Usage CMK Encryption

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_event_connection import CloudwatchEventConnection
from imports.aws.data_aws_caller_identity import DataAwsCallerIdentity
from imports.aws.data_aws_partition import DataAwsPartition
from imports.aws.kms_key import KmsKey
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchEventConnection(self, "test",
            auth_parameters=CloudwatchEventConnectionAuthParameters(
                basic=CloudwatchEventConnectionAuthParametersBasic(
                    password="Pass1234!",
                    username="user"
                )
            ),
            authorization_type="BASIC",
            description="A connection description",
            kms_key_identifier=example.id,
            name="ngrok-connection"
        )
        current = DataAwsCallerIdentity(self, "current")
        data_aws_partition_current = DataAwsPartition(self, "current_2")
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_partition_current.override_logical_id("current")
        aws_kms_key_test = KmsKey(self, "test_3",
            deletion_window_in_days=7,
            policy=Token.as_string(
                Fn.jsonencode({
                    "Id": "key-policy-example",
                    "Statement": [{
                        "Action": "kms:*",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:${" + data_aws_partition_current.partition + "}:iam::${" + current.account_id + "}:root"
                        },
                        "Resource": "*",
                        "Sid": "Enable IAM User Permissions"
                    }, {
                        "Action": ["kms:DescribeKey", "kms:Decrypt", "kms:GenerateDataKey"],
                        "Condition": {
                            "StringLike": {
                                "kms:_encryption_context:_secret_aRN": ["arn:aws:secretsmanager:*:*:secret:events!connection/*"
                                ],
                                "kms:_via_service": "secretsmanager.*.amazonaws.com"
                            }
                        },
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:${" + data_aws_partition_current.partition + "}:iam::${" + current.account_id + "}:root"
                        },
                        "Resource": "*",
                        "Sid": "Allow use of the key"
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            tags={
                "EventBridgeApiDestinations": "true"
            }
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_kms_key_test.override_logical_id("test")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name for the connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
* `description` - (Optional) Description for the connection. Maximum of 512 characters.
* `authorization_type` - (Required) Type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
* `auth_parameters` - (Required) Parameters used for authorization. A maximum of 1 are allowed. Documented below.
* `invocation_connectivity_parameters` - (Optional) Parameters to use for invoking a private API. Documented below.
* `kms_key_identifier` - (Optional) Identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

`auth_parameters` support the following:

* `api_key` - (Optional) Parameters used for API_KEY authorization. An API key to include in the header for each authentication request. A maximum of 1 are allowed. Conflicts with `basic` and `oauth`. Documented below.
* `basic` - (Optional) Parameters used for BASIC authorization. A maximum of 1 are allowed. Conflicts with `api_key` and `oauth`. Documented below.
* `invocation_http_parameters` - (Optional) Invocation Http Parameters are additional credentials used to sign each Invocation of the ApiDestination created from this Connection. If the ApiDestination Rule Target has additional HttpParameters, the values will be merged together, with the Connection Invocation Http Parameters taking precedence. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
* `oauth` - (Optional) Parameters used for OAUTH_CLIENT_CREDENTIALS authorization. A maximum of 1 are allowed. Conflicts with `basic` and `api_key`. Documented below.

`api_key` support the following:

* `key` - (Required) Header Name.
* `value` - (Required) Header Value. Created and stored in AWS Secrets Manager.

`basic` support the following:

* `username` - (Required) A username for the authorization.
* `password` - (Required) A password for the authorization. Created and stored in AWS Secrets Manager.

`oauth` support the following:

* `authorization_endpoint` - (Required) The URL to the authorization endpoint.
* `http_method` - (Required) A password for the authorization. Created and stored in AWS Secrets Manager.
* `client_parameters` - (Required) Contains the client parameters for OAuth authorization. Contains the following two parameters.
    * `client_id` - (Required) The client ID for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
    * `client_secret` - (Required) The client secret for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
* `oauth_http_parameters` - (Required) OAuth Http Parameters are additional credentials used to sign the request to the authorization endpoint to exchange the OAuth Client information for an access token. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.

`invocation_http_parameters` and `oauth_http_parameters` support the following:

* `body` - (Optional) Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
    * `key` - (Required) The key for the parameter.
    * `value` - (Required) The value associated with the key. Created and stored in AWS Secrets Manager if is secret.
    * `is_value_secret` - (Optional) Specified whether the value is secret.

* `header` - (Optional) Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
    * `key` - (Required) The key for the parameter.
    * `value` - (Required) The value associated with the key. Created and stored in AWS Secrets Manager if is secret.
    * `is_value_secret` - (Optional) Specified whether the value is secret.

* `query_string` - (Optional) Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
    * `key` - (Required) The key for the parameter.
    * `value` - (Required) The value associated with the key. Created and stored in AWS Secrets Manager if is secret.
    * `is_value_secret` - (Optional) Specified whether the value is secret.

`invocation_connectivity_parameters` supports the following:

* `resource_parameters` - (Required) The parameters for EventBridge to use when invoking the resource endpoint. Documented below.

`resource_parameters` supports the following:

* `resource_configuration_arn` - (Required) ARN of the Amazon VPC Lattice [resource configuration](vpclattice_resource_configuration) for the resource endpoint.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the connection.
* `secret_arn` - The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EventBridge connection using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_event_connection import CloudwatchEventConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchEventConnection.generate_config_for_import(self, "test", "ngrok-connection")
```

Using `terraform import`, import EventBridge EventBridge connection using the `name`. For example:

```console
% terraform import aws_cloudwatch_event_connection.test ngrok-connection
```

<!-- cache-key: cdktf-0.20.8 input-89571db5a946668df93cfaee7902913652ae1dd15f0f4ecd31915846dd20405b -->