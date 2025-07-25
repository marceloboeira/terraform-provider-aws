---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "AWS: aws_s3_bucket_policy"
description: |-
    Provides IAM policy of an S3 bucket
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_s3_bucket_policy

The bucket policy data source returns IAM policy of an S3 bucket.

## Example Usage

The following example retrieves IAM policy of a specified S3 bucket.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformOutput, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_s3_bucket_policy import DataAwsS3BucketPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsS3BucketPolicy(self, "example",
            bucket="example-bucket-name"
        )
        TerraformOutput(self, "foo",
            value=example.policy
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `bucket` - (Required) Bucket name.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `policy` - IAM bucket policy.

<!-- cache-key: cdktf-0.20.8 input-c52fdb1fa6a13925ab126d46f133c636dd54d9336f0071f3a8eb2bf5c9078ed4 -->