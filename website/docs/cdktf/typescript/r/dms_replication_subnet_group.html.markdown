---
subcategory: "DMS (Database Migration)"
layout: "aws"
page_title: "AWS: aws_dms_replication_subnet_group"
description: |-
  Provides a DMS (Data Migration Service) subnet group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dms_replication_subnet_group

Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.

~> **Note:** AWS requires a special IAM role called `dms-vpc-role` when using this resource. See the example below to create it as part of your configuration.

## Example Usage

### Basic

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DmsReplicationSubnetGroup } from "./.gen/providers/aws/dms-replication-subnet-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DmsReplicationSubnetGroup(this, "example", {
      replicationSubnetGroupDescription: "Example replication subnet group",
      replicationSubnetGroupId: "example-dms-replication-subnet-group-tf",
      subnetIds: ["subnet-12345678", "subnet-12345679"],
      tags: {
        Name: "example",
      },
    });
  }
}

```

### Creating special IAM role

If your account does not already include the `dms-vpc-role` IAM role, you will need to create it to allow DMS to manage subnets in the VPC.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DmsReplicationSubnetGroup } from "./.gen/providers/aws/dms-replication-subnet-group";
import { IamRole } from "./.gen/providers/aws/iam-role";
import { IamRolePolicyAttachment } from "./.gen/providers/aws/iam-role-policy-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const dmsVpcRole = new IamRole(this, "dms-vpc-role", {
      assumeRolePolicy: Token.asString(
        Fn.jsonencode({
          Statement: [
            {
              Action: "sts:AssumeRole",
              Effect: "Allow",
              Principal: {
                Service: "dms.amazonaws.com",
              },
            },
          ],
          Version: "2012-10-17",
        })
      ),
      description: "Allows DMS to manage VPC",
      name: "dms-vpc-role",
    });
    const example = new IamRolePolicyAttachment(this, "example", {
      policyArn:
        "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
      role: dmsVpcRole.name,
    });
    const awsDmsReplicationSubnetGroupExample = new DmsReplicationSubnetGroup(
      this,
      "example_2",
      {
        dependsOn: [example],
        replicationSubnetGroupDescription: "Example",
        replicationSubnetGroupId: "example-id",
        subnetIds: ["subnet-12345678", "subnet-12345679"],
        tags: {
          Name: "example-id",
        },
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDmsReplicationSubnetGroupExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `replicationSubnetGroupDescription` - (Required) Description for the subnet group.
* `replicationSubnetGroupId` - (Required) Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
* `subnetIds` - (Required) List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `vpcId` - The ID of the VPC the subnet group is in.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `15m`)
- `update` - (Default `15m`)
- `delete` - (Default `15m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import replication subnet groups using the `replicationSubnetGroupId`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DmsReplicationSubnetGroup } from "./.gen/providers/aws/dms-replication-subnet-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DmsReplicationSubnetGroup.generateConfigForImport(
      this,
      "test",
      "test-dms-replication-subnet-group-tf"
    );
  }
}

```

Using `terraform import`, import replication subnet groups using the `replicationSubnetGroupId`. For example:

```console
% terraform import aws_dms_replication_subnet_group.test test-dms-replication-subnet-group-tf
```

<!-- cache-key: cdktf-0.20.8 input-e197ab059787bbc1ac6325081867dbb3c4420ed1a8aa7ebb5521a9f2e54f3534 -->