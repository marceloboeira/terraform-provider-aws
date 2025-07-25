---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_maintenance_window_target"
description: |-
  Provides an SSM Maintenance Window Target resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssm_maintenance_window_target

Provides an SSM Maintenance Window Target resource

## Example Usage

### Instance Target

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindow } from "./.gen/providers/aws/ssm-maintenance-window";
import { SsmMaintenanceWindowTarget } from "./.gen/providers/aws/ssm-maintenance-window-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const window = new SsmMaintenanceWindow(this, "window", {
      cutoff: 1,
      duration: 3,
      name: "maintenance-window-webapp",
      schedule: "cron(0 16 ? * TUE *)",
    });
    new SsmMaintenanceWindowTarget(this, "target1", {
      description: "This is a maintenance window target",
      name: "maintenance-window-target",
      resourceType: "INSTANCE",
      targets: [
        {
          key: "tag:Name",
          values: ["acceptance_test"],
        },
      ],
      windowId: window.id,
    });
  }
}

```

### Resource Group Target

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindow } from "./.gen/providers/aws/ssm-maintenance-window";
import { SsmMaintenanceWindowTarget } from "./.gen/providers/aws/ssm-maintenance-window-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const window = new SsmMaintenanceWindow(this, "window", {
      cutoff: 1,
      duration: 3,
      name: "maintenance-window-webapp",
      schedule: "cron(0 16 ? * TUE *)",
    });
    new SsmMaintenanceWindowTarget(this, "target1", {
      description: "This is a maintenance window target",
      name: "maintenance-window-target",
      resourceType: "RESOURCE_GROUP",
      targets: [
        {
          key: "resource-groups:ResourceTypeFilters",
          values: ["AWS::EC2::Instance"],
        },
      ],
      windowId: window.id,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `windowId` - (Required) The Id of the maintenance window to register the target with.
* `name` - (Optional) The name of the maintenance window target.
* `description` - (Optional) The description of the maintenance window target.
* `resourceType` - (Required) The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
* `targets` - (Required) The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
 (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
* `ownerInformation` - (Optional) User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the maintenance window target.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SSM Maintenance Window targets using `WINDOW_ID/WINDOW_TARGET_ID`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindowTarget } from "./.gen/providers/aws/ssm-maintenance-window-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SsmMaintenanceWindowTarget.generateConfigForImport(
      this,
      "example",
      "mw-0c50858d01EXAMPLE/23639a0b-ddbc-4bca-9e72-78d96EXAMPLE"
    );
  }
}

```

Using `terraform import`, import SSM Maintenance Window targets using `WINDOW_ID/WINDOW_TARGET_ID`. For example:

```console
% terraform import aws_ssm_maintenance_window_target.example mw-0c50858d01EXAMPLE/23639a0b-ddbc-4bca-9e72-78d96EXAMPLE
```

<!-- cache-key: cdktf-0.20.8 input-8716dd3aed621b76012dead6aa22e71c66c1000751f5bcc3d304bf6e3c8d7174 -->