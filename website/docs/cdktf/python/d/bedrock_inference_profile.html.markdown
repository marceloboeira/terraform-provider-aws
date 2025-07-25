---
subcategory: "Bedrock"
layout: "aws"
page_title: "AWS: aws_bedrock_inference_profile"
description: |-
  Terraform data source for managing an AWS Bedrock Inference Profile.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_bedrock_inference_profile

Terraform data source for managing an AWS Bedrock Inference Profile.

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
from imports.aws.data_aws_bedrock_inference_profile import DataAwsBedrockInferenceProfile
from imports.aws.data_aws_bedrock_inference_profiles import DataAwsBedrockInferenceProfiles
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        test = DataAwsBedrockInferenceProfiles(self, "test")
        data_aws_bedrock_inference_profile_test =
        DataAwsBedrockInferenceProfile(self, "test_1",
            inference_profile_id=Token.as_string(
                Fn.lookup_nested(test.inference_profile_summaries, ["0", "inference_profile_id"
                ]))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_bedrock_inference_profile_test.override_logical_id("test")
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `inference_profile_id` - (Required) Inference Profile identifier.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

- `inference_profile_arn` - The Amazon Resource Name (ARN) of the inference profile.
- `inference_profile_name` - The unique identifier of the inference profile.
- `models` - A list of information about each model in the inference profile. See [`models`](#models).
- `status` - The status of the inference profile. `ACTIVE` means that the inference profile is available to use.
- `type` - The type of the inference profile. `SYSTEM_DEFINED` means that the inference profile is defined by Amazon Bedrock. `APPLICATION` means that the inference profile is defined by the user.
- `created_at` - The time at which the inference profile was created.
- `description` - The description of the inference profile.
- `updated_at` - The time at which the inference profile was last updated.

### `models`

- `model_arn` - The Amazon Resource Name (ARN) of the model.

<!-- cache-key: cdktf-0.20.8 input-ce9c57bc602aee4011dc55abf62f1a45ad4a18bc33669a0e70b4e9d1ed631871 -->