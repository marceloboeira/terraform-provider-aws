---
subcategory: "Redshift"
layout: "aws"
page_title: "AWS: aws_redshift_snapshot_schedule"
description: |-
  Provides an Redshift Snapshot Schedule resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_redshift_snapshot_schedule

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RedshiftSnapshotSchedule } from "./.gen/providers/aws/redshift-snapshot-schedule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new RedshiftSnapshotSchedule(this, "default", {
      definitions: ["rate(12 hours)"],
      identifier: "tf-redshift-snapshot-schedule",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `identifier` - (Optional, Forces new resource) The snapshot schedule identifier. If omitted, Terraform will assign a random, unique identifier.
* `identifierPrefix` - (Optional, Forces new resource) Creates a unique
identifier beginning with the specified prefix. Conflicts with `identifier`.
* `description` - (Optional) The description of the snapshot schedule.
* `definitions` - (Optional) The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
* `forceDestroy` - (Optional) Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the Redshift Snapshot Schedule.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Redshift Snapshot Schedule using the `identifier`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RedshiftSnapshotSchedule } from "./.gen/providers/aws/redshift-snapshot-schedule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    RedshiftSnapshotSchedule.generateConfigForImport(
      this,
      "default",
      "tf-redshift-snapshot-schedule"
    );
  }
}

```

Using `terraform import`, import Redshift Snapshot Schedule using the `identifier`. For example:

```console
% terraform import aws_redshift_snapshot_schedule.default tf-redshift-snapshot-schedule
```

<!-- cache-key: cdktf-0.20.8 input-cbfcbe9d650987e15d30a73204538668ccc2e494691b5c150fb75ff1253eb683 -->