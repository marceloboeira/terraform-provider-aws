---
subcategory: "SES (Simple Email)"
layout: "aws"
page_title: "AWS: aws_ses_active_receipt_rule_set"
description: |-
  Retrieve the active SES receipt rule set
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ses_active_receipt_rule_set

Retrieve the active SES receipt rule set

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsSesActiveReceiptRuleSet } from "./.gen/providers/aws/data-aws-ses-active-receipt-rule-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsSesActiveReceiptRuleSet(this, "main", {});
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - SES receipt rule set ARN.
* `ruleSetName` - Name of the rule set

<!-- cache-key: cdktf-0.20.8 input-8c1cde8147ab60839e3b3ea77a29e3169a07c133077e254557d587b273702866 -->