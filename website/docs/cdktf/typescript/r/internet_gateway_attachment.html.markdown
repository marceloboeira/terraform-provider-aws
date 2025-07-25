---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_internet_gateway_attachment"
description: |-
  Provides a resource to create a VPC Internet Gateway Attachment.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_internet_gateway_attachment

Provides a resource to create a VPC Internet Gateway Attachment.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { InternetGateway } from "./.gen/providers/aws/internet-gateway";
import { InternetGatewayAttachment } from "./.gen/providers/aws/internet-gateway-attachment";
import { Vpc } from "./.gen/providers/aws/vpc";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new InternetGateway(this, "example", {});
    const awsVpcExample = new Vpc(this, "example_1", {
      cidrBlock: "10.1.0.0/16",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsVpcExample.overrideLogicalId("example");
    const awsInternetGatewayAttachmentExample = new InternetGatewayAttachment(
      this,
      "example_2",
      {
        internetGatewayId: example.id,
        vpcId: Token.asString(awsVpcExample.id),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsInternetGatewayAttachmentExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `internetGatewayId` - (Required) The ID of the internet gateway.
* `vpcId` - (Required) The ID of the VPC.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the VPC and Internet Gateway separated by a colon.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `20m`)
- `delete` - (Default `20m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Internet Gateway Attachments using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { InternetGatewayAttachment } from "./.gen/providers/aws/internet-gateway-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    InternetGatewayAttachment.generateConfigForImport(
      this,
      "example",
      "igw-c0a643a9:vpc-123456"
    );
  }
}

```

Using `terraform import`, import Internet Gateway Attachments using the `id`. For example:

```console
% terraform import aws_internet_gateway_attachment.example igw-c0a643a9:vpc-123456
```

<!-- cache-key: cdktf-0.20.8 input-51e5d79bfa8bcf5ae3791841c9824423c3c25786f72567a790a8893b02862a87 -->