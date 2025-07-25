---
subcategory: "RDS (Relational Database)"
layout: "aws"
page_title: "AWS: aws_db_event_subscription"
description: |-
  Provides a DB event subscription resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_db_event_subscription

Provides a DB event subscription resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DbEventSubscription } from "./.gen/providers/aws/db-event-subscription";
import { DbInstance } from "./.gen/providers/aws/db-instance";
import { SnsTopic } from "./.gen/providers/aws/sns-topic";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const defaultVar = new DbInstance(this, "default", {
      allocatedStorage: 10,
      dbName: "mydb",
      dbSubnetGroupName: "my_database_subnet_group",
      engine: "mysql",
      engineVersion: "5.6.17",
      instanceClass: "db.t2.micro",
      parameterGroupName: "default.mysql5.6",
      password: "bar",
      username: "foo",
    });
    const awsSnsTopicDefault = new SnsTopic(this, "default_1", {
      name: "rds-events",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsSnsTopicDefault.overrideLogicalId("default");
    const awsDbEventSubscriptionDefault = new DbEventSubscription(
      this,
      "default_2",
      {
        eventCategories: [
          "availability",
          "deletion",
          "failover",
          "failure",
          "low storage",
          "maintenance",
          "notification",
          "read replica",
          "recovery",
          "restoration",
        ],
        name: "rds-event-sub",
        snsTopic: Token.asString(awsSnsTopicDefault.arn),
        sourceIds: [defaultVar.identifier],
        sourceType: "db-instance",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDbEventSubscriptionDefault.overrideLogicalId("default");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Optional) The name of the DB event subscription. By default generated by Terraform.
* `namePrefix` - (Optional) The name of the DB event subscription. Conflicts with `name`.
* `snsTopic` - (Required) The SNS topic to send events to.
* `sourceIds` - (Optional) A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
* `sourceType` - (Optional) The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
* `eventCategories` - (Optional) A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
* `enabled` - (Optional) A boolean flag to enable/disable the subscription. Defaults to true.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The name of the RDS event notification subscription
* `arn` - The Amazon Resource Name of the RDS event notification subscription
* `customerAwsId` - The AWS customer account associated with the RDS event notification subscription
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `40m`)
- `delete` - (Default `40m`)
- `update` - (Default `40m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DB Event Subscriptions using the `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DbEventSubscription } from "./.gen/providers/aws/db-event-subscription";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DbEventSubscription.generateConfigForImport(
      this,
      "default",
      "rds-event-sub"
    );
  }
}

```

Using `terraform import`, import DB Event Subscriptions using the `name`. For example:

```console
% terraform import aws_db_event_subscription.default rds-event-sub
```

<!-- cache-key: cdktf-0.20.8 input-5515aa2f197c378d5a0b8d843018de9857d510b08618c24b71f7c214a716ca30 -->