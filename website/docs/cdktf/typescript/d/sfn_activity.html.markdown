---
subcategory: "SFN (Step Functions)"
layout: "aws"
page_title: "AWS: aws_sfn_activity"
description: |-
  Use this data source to get information about a Step Functions Activity.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_sfn_activity

Provides a Step Functions Activity data source

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsSfnActivity } from "./.gen/providers/aws/data-aws-sfn-activity";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsSfnActivity(this, "sfn_activity", {
      name: "my-activity",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Optional) Name that identifies the activity.
* `arn` - (Optional) ARN that identifies the activity.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - ARN that identifies the activity.
* `creationDate` - Date the activity was created.

<!-- cache-key: cdktf-0.20.8 input-df253935e50b8adfc1992d59a02b1ae7dde44de3aa93f4b7f5899db93b2ecb77 -->