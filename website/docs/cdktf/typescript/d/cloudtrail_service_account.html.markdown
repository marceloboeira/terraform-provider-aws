---
subcategory: "CloudTrail"
layout: "aws"
page_title: "AWS: aws_cloudtrail_service_account"
description: |-
  Get AWS CloudTrail Service Account ID for storing trail data in S3.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_cloudtrail_service_account

Use this data source to get the Account ID of the [AWS CloudTrail Service Account](http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-supported-regions.html)
in a given region for the purpose of allowing CloudTrail to store trail data in S3.

~> **Warning:** This data source is deprecated. The AWS documentation [states that](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create-s3-bucket-policy-for-cloudtrail.html#troubleshooting-s3-bucket-policy) a [service principal name](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services) should be used instead of an AWS account ID in any relevant IAM policy.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsCloudtrailServiceAccount } from "./.gen/providers/aws/data-aws-cloudtrail-service-account";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { S3Bucket } from "./.gen/providers/aws/s3-bucket";
import { S3BucketPolicy } from "./.gen/providers/aws/s3-bucket-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const bucket = new S3Bucket(this, "bucket", {
      bucket: "tf-cloudtrail-logging-test-bucket",
      forceDestroy: true,
    });
    const main = new DataAwsCloudtrailServiceAccount(this, "main", {});
    const allowCloudtrailLogging = new DataAwsIamPolicyDocument(
      this,
      "allow_cloudtrail_logging",
      {
        statement: [
          {
            actions: ["s3:PutObject"],
            effect: "Allow",
            principals: [
              {
                identifiers: [Token.asString(main.arn)],
                type: "AWS",
              },
            ],
            resources: ["${" + bucket.arn + "}/*"],
            sid: "Put bucket policy needed for trails",
          },
          {
            actions: ["s3:GetBucketAcl"],
            effect: "Allow",
            principals: [
              {
                identifiers: [Token.asString(main.arn)],
                type: "AWS",
              },
            ],
            resources: [bucket.arn],
            sid: "Get bucket policy needed for trails",
          },
        ],
      }
    );
    const awsS3BucketPolicyAllowCloudtrailLogging = new S3BucketPolicy(
      this,
      "allow_cloudtrail_logging_3",
      {
        bucket: bucket.id,
        policy: Token.asString(allowCloudtrailLogging.json),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3BucketPolicyAllowCloudtrailLogging.overrideLogicalId(
      "allow_cloudtrail_logging"
    );
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Name of the Region whose AWS CloudTrail account ID is desired. Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - ID of the AWS CloudTrail service account in the selected Region.
* `arn` - ARN of the AWS CloudTrail service account in the selected Region.

<!-- cache-key: cdktf-0.20.8 input-7d03bf056d83d24ff4a8e8c0186dc2c6012fafd4991666a1b6ca9425a8ef7a52 -->