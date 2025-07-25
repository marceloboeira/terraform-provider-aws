---
subcategory: "Lambda"
layout: "aws"
page_title: "AWS: aws_lambda_layer_version"
description: |-
  Manages an AWS Lambda Layer Version.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lambda_layer_version

Manages an AWS Lambda Layer Version. Use this resource to share code and dependencies across multiple Lambda functions.

For information about Lambda Layers and how to use them, see [AWS Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).

~> **Note:** Setting `skipDestroy` to `true` means that the AWS Provider will not destroy any layer version, even when running `terraform destroy`. Layer versions are thus intentional dangling resources that are not managed by Terraform and may incur extra expense in your AWS account.

## Example Usage

### Basic Layer

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LambdaLayerVersion } from "./.gen/providers/aws/lambda-layer-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new LambdaLayerVersion(this, "example", {
      compatibleRuntimes: ["nodejs20.x"],
      filename: "lambda_layer_payload.zip",
      layerName: "lambda_layer_name",
    });
  }
}

```

### Layer with S3 Source

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LambdaLayerVersion } from "./.gen/providers/aws/lambda-layer-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new LambdaLayerVersion(this, "example", {
      compatibleArchitectures: ["x86_64", "arm64"],
      compatibleRuntimes: ["nodejs20.x", "python3.12"],
      layerName: "lambda_layer_name",
      s3Bucket: lambdaLayerZip.bucket,
      s3Key: lambdaLayerZip.key,
    });
  }
}

```

### Layer with Multiple Runtimes and Architectures

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LambdaLayerVersion } from "./.gen/providers/aws/lambda-layer-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new LambdaLayerVersion(this, "example", {
      compatibleArchitectures: ["x86_64", "arm64"],
      compatibleRuntimes: [
        "nodejs18.x",
        "nodejs20.x",
        "python3.11",
        "python3.12",
      ],
      description: "Shared utilities for Lambda functions",
      filename: "lambda_layer_payload.zip",
      layerName: "multi_runtime_layer",
      licenseInfo: "MIT",
      sourceCodeHash: Token.asString(
        Fn.filebase64sha256("lambda_layer_payload.zip")
      ),
    });
  }
}

```

## Specifying the Deployment Package

AWS Lambda Layers expect source code to be provided as a deployment package whose structure varies depending on which `compatibleRuntimes` this layer specifies. See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) for the valid values of `compatibleRuntimes`.

Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or indirectly via Amazon S3 (using the `s3Bucket`, `s3Key` and `s3ObjectVersion` arguments). When providing the deployment package via S3 it may be useful to use [the `aws_s3_object` resource](s3_object.html) to upload it.

For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading large files efficiently.

## Argument Reference

The following arguments are required:

* `layerName` - (Required) Unique name for your Lambda Layer.

The following arguments are optional:

* `compatibleArchitectures` - (Optional) List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x86_64` and `arm64` can be specified.
* `compatibleRuntimes` - (Optional) List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 15 runtimes can be specified.
* `description` - (Optional) Description of what your Lambda Layer does.
* `filename` - (Optional) Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
* `licenseInfo` - (Optional) License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `s3Bucket` - (Optional) S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
* `s3Key` - (Optional) S3 key of an object containing the function's deployment package. Conflicts with `filename`.
* `s3ObjectVersion` - (Optional) Object version containing the function's deployment package. Conflicts with `filename`.
* `skipDestroy` - (Optional) Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
* `sourceCodeHash` - (Optional) Virtual attribute used to trigger replacement when source code changes. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `filebase64sha256("file.zip")` (Terraform 0.11.12 or later) or `base64sha256(file("file.zip"))` (Terraform 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Lambda Layer with version.
* `codeSha256` - Base64-encoded representation of raw SHA-256 sum of the zip file.
* `createdDate` - Date this resource was created.
* `layerArn` - ARN of the Lambda Layer without version.
* `signingJobArn` - ARN of a signing job.
* `signingProfileVersionArn` - ARN for a signing profile version.
* `sourceCodeSize` - Size in bytes of the function .zip file.
* `version` - Lambda Layer version.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Lambda Layers using `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LambdaLayerVersion } from "./.gen/providers/aws/lambda-layer-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    LambdaLayerVersion.generateConfigForImport(
      this,
      "example",
      "arn:aws:lambda:us-west-2:123456789012:layer:example:1"
    );
  }
}

```

Using `terraform import`, import Lambda Layers using `arn`. For example:

```console
% terraform import aws_lambda_layer_version.example arn:aws:lambda:us-west-2:123456789012:layer:example:1
```

<!-- cache-key: cdktf-0.20.8 input-17d6f85fd3e965491ff0999c83723c133d2ef2607b90dbd2bd3c2fac8fb23c56 -->