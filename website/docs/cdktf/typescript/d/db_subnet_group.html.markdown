---
subcategory: "RDS (Relational Database)"
layout: "aws"
page_title: "AWS: aws_db_subnet_group"
description: |-
  Get information on an RDS Database Subnet Group.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_db_subnet_group

Use this data source to get information about an RDS subnet group.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsDbSubnetGroup } from "./.gen/providers/aws/data-aws-db-subnet-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsDbSubnetGroup(this, "database", {
      name: "my-test-database-subnet-group",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the RDS database subnet group.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN for the DB subnet group.
* `description` - Provides the description of the DB subnet group.
* `status` - Provides the status of the DB subnet group.
* `subnetIds` - Contains a list of subnet identifiers.
* `supportedNetworkTypes` - The network type of the DB subnet group.
* `vpcId` - Provides the VPC ID of the DB subnet group.

<!-- cache-key: cdktf-0.20.8 input-006d32fd18271d3ec26252fbfb306e5eccbee743f5b9ec67a261eca5f2485207 -->