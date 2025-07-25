---
subcategory: "Bedrock"
layout: "aws"
page_title: "AWS: aws_bedrock_foundation_model"
description: |-
  Terraform data source for managing an AWS Bedrock Foundation Model.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_bedrock_foundation_model

Terraform data source for managing an AWS Bedrock Foundation Model.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsBedrockFoundationModel } from "./.gen/providers/aws/data-aws-bedrock-foundation-model";
import { DataAwsBedrockFoundationModels } from "./.gen/providers/aws/data-aws-bedrock-foundation-models";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const test = new DataAwsBedrockFoundationModels(this, "test", {});
    const dataAwsBedrockFoundationModelTest = new DataAwsBedrockFoundationModel(
      this,
      "test_1",
      {
        modelId: Token.asString(
          Fn.lookupNested(test.modelSummaries, ["0", "model_id"])
        ),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsBedrockFoundationModelTest.overrideLogicalId("test");
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `modelId` - (Required) Model identifier.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `customizationsSupported` - Customizations that the model supports.
* `inferenceTypesSupported` - Inference types that the model supports.
* `inputModalities` - Input modalities that the model supports.
* `modelArn` - Model ARN.
* `modelName` - Model name.
* `outputModalities` - Output modalities that the model supports.
* `providerName` - Model provider name.
* `responseStreamingSupported` - Indicates whether the model supports streaming.

<!-- cache-key: cdktf-0.20.8 input-a67af68ab2d739928980a3a40e7ff283ddbc049ff3b090f23bdfa033811a21b2 -->