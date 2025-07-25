---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_subnets"
description: |-
    Get information about a set of subnets.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_subnets

This resource can be useful for getting back a set of subnet IDs.

## Example Usage

The following shows outputting all CIDR blocks for every subnet ID in a VPC.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import {
  Fn,
  Token,
  TerraformIterator,
  TerraformOutput,
  TerraformStack,
} from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsSubnet } from "./.gen/providers/aws/data-aws-subnet";
import { DataAwsSubnets } from "./.gen/providers/aws/data-aws-subnets";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new DataAwsSubnets(this, "example", {
      filter: [
        {
          name: "vpc-id",
          values: [vpcId.stringValue],
        },
      ],
    });
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const exampleForEachIterator = TerraformIterator.fromList(
      Token.asAny(Fn.toset(example.ids))
    );
    const dataAwsSubnetExample = new DataAwsSubnet(this, "example_1", {
      id: Token.asString(exampleForEachIterator.value),
      forEach: exampleForEachIterator,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsSubnetExample.overrideLogicalId("example");
    new TerraformOutput(this, "subnet_cidr_blocks", {
      value:
        "${[ for s in ${" + dataAwsSubnetExample.fqn + "} : s.cidr_block]}",
    });
  }
}

```

The following example retrieves a set of all subnets in a VPC with a custom
tag of `Tier` set to a value of "Private" so that the `aws_instance` resource
can loop through the subnets, putting instances across availability zones.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformIterator, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsSubnets } from "./.gen/providers/aws/data-aws-subnets";
import { Instance } from "./.gen/providers/aws/instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const private = new DataAwsSubnets(this, "private", {
      filter: [
        {
          name: "vpc-id",
          values: [vpcId.stringValue],
        },
      ],
      tags: {
        Tier: "Private",
      },
    });
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const appForEachIterator = TerraformIterator.fromList(
      Token.asAny(Fn.toset(private.ids))
    );
    new Instance(this, "app", {
      ami: ami.stringValue,
      instanceType: "t2.micro",
      subnetId: Token.asString(appForEachIterator.value),
      forEach: appForEachIterator,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `filter` - (Optional) Custom filter block as described below.
* `tags` - (Optional) Map of tags, each pair of which must exactly match
  a pair on the desired subnets.

### `filter`

More complex filters can be expressed using one or more `filter` sub-blocks, which take the following arguments:

* `name` - (Required) Name of the field to filter by, as defined by
  [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html).
  For example, if matching against tag `Name`, use:
* `values` - (Required) Set of values that are accepted for the given field.
  A Subnet will be selected if any one of the given values matches.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsSubnets } from "./.gen/providers/aws/data-aws-subnets";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsSubnets(this, "selected", {
      filter: [
        {
          name: "tag:Name",
          values: [""],
        },
      ],
    });
  }
}

```

* `values` - (Required) Set of values that are accepted for the given field.
  Subnet IDs will be selected if any one of the given values match.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `ids` - List of all the subnet ids found.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-a516d3773d3fb6fb5eed16680984b7fb1ea878fc66151227312adb40efabd5dc -->