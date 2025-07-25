---
subcategory: "FIS (Fault Injection Simulator)"
layout: "aws"
page_title: "AWS: aws_fis_experiment_templates"
description: |-
  Get information about a set of FIS Experiment Templates
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_fis_experiment_templates

This resource can be useful for getting back a set of FIS experiment template IDs.

## Example Usage

The following shows outputting a list of all FIS experiment template IDs

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformOutput, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsFisExperimentTemplates } from "./.gen/providers/aws/data-aws-fis-experiment-templates";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const all = new DataAwsFisExperimentTemplates(this, "all", {});
    const cdktfTerraformOutputAll = new TerraformOutput(this, "all_1", {
      value: all.ids,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    cdktfTerraformOutputAll.overrideLogicalId("all");
  }
}

```

The following shows filtering FIS experiment templates by tag

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsFisExperimentTemplates } from "./.gen/providers/aws/data-aws-fis-experiment-templates";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new DataAwsFisExperimentTemplates(this, "example", {
      tags: {
        Name: "example",
        Tier: Token.asString(1),
      },
    });
    const dataAwsIamPolicyDocumentExample = new DataAwsIamPolicyDocument(
      this,
      "example_1",
      {
        statement: [
          {
            actions: ["fis:StartExperiment"],
            effect: "Allow",
            resources: [
              "arn:aws:fis:*:*:experiment-template/" +
                Token.asString(Fn.lookupNested(example.ids, ["0"])),
              "arn:aws:fis:*:*:experiment/*",
            ],
            sid: "StartFISExperiment",
          },
        ],
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsIamPolicyDocumentExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `tags` - (Optional) Map of tags, each pair of which must exactly match
  a pair on the desired experiment templates.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `ids` - List of all the experiment template ids found.

<!-- cache-key: cdktf-0.20.8 input-a513ae368cb8423036cb07c5fb9b81e5dad86afc1dfb866d0f71b299557550a9 -->