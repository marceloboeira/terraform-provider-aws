---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_maintenance_window_task"
description: |-
  Provides an SSM Maintenance Window Task resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssm_maintenance_window_task

Provides an SSM Maintenance Window Task resource

## Example Usage

### Automation Tasks

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindowTask } from "./.gen/providers/aws/ssm-maintenance-window-task";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmMaintenanceWindowTask(this, "example", {
      maxConcurrency: Token.asString(2),
      maxErrors: Token.asString(1),
      priority: 1,
      targets: [
        {
          key: "InstanceIds",
          values: [Token.asString(awsInstanceExample.id)],
        },
      ],
      taskArn: "AWS-RestartEC2Instance",
      taskInvocationParameters: {
        automationParameters: {
          documentVersion: "$LATEST",
          parameter: [
            {
              name: "InstanceId",
              values: [Token.asString(awsInstanceExample.id)],
            },
          ],
        },
      },
      taskType: "AUTOMATION",
      windowId: Token.asString(awsSsmMaintenanceWindowExample.id),
    });
  }
}

```

### Lambda Tasks

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindowTask } from "./.gen/providers/aws/ssm-maintenance-window-task";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmMaintenanceWindowTask(this, "example", {
      maxConcurrency: Token.asString(2),
      maxErrors: Token.asString(1),
      priority: 1,
      targets: [
        {
          key: "InstanceIds",
          values: [Token.asString(awsInstanceExample.id)],
        },
      ],
      taskArn: Token.asString(awsLambdaFunctionExample.arn),
      taskInvocationParameters: {
        lambdaParameters: {
          clientContext: Token.asString(
            Fn.base64encode('{\\"key1\\":\\"value1\\"}')
          ),
          payload: '{\\"key1\\":\\"value1\\"}',
        },
      },
      taskType: "LAMBDA",
      windowId: Token.asString(awsSsmMaintenanceWindowExample.id),
    });
  }
}

```

### Run Command Tasks

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindowTask } from "./.gen/providers/aws/ssm-maintenance-window-task";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmMaintenanceWindowTask(this, "example", {
      maxConcurrency: Token.asString(2),
      maxErrors: Token.asString(1),
      priority: 1,
      targets: [
        {
          key: "InstanceIds",
          values: [Token.asString(awsInstanceExample.id)],
        },
      ],
      taskArn: "AWS-RunShellScript",
      taskInvocationParameters: {
        runCommandParameters: {
          notificationConfig: {
            notificationArn: Token.asString(awsSnsTopicExample.arn),
            notificationEvents: ["All"],
            notificationType: "Command",
          },
          outputS3Bucket: Token.asString(awsS3BucketExample.id),
          outputS3KeyPrefix: "output",
          parameter: [
            {
              name: "commands",
              values: ["date"],
            },
          ],
          serviceRoleArn: Token.asString(awsIamRoleExample.arn),
          timeoutSeconds: 600,
        },
      },
      taskType: "RUN_COMMAND",
      windowId: Token.asString(awsSsmMaintenanceWindowExample.id),
    });
  }
}

```

### Step Function Tasks

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindowTask } from "./.gen/providers/aws/ssm-maintenance-window-task";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmMaintenanceWindowTask(this, "example", {
      maxConcurrency: Token.asString(2),
      maxErrors: Token.asString(1),
      priority: 1,
      targets: [
        {
          key: "InstanceIds",
          values: [Token.asString(awsInstanceExample.id)],
        },
      ],
      taskArn: Token.asString(awsSfnActivityExample.id),
      taskInvocationParameters: {
        stepFunctionsParameters: {
          input: '{\\"key1\\":\\"value1\\"}',
          name: "example",
        },
      },
      taskType: "STEP_FUNCTIONS",
      windowId: Token.asString(awsSsmMaintenanceWindowExample.id),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `windowId` - (Required) The Id of the maintenance window to register the task with.
* `maxConcurrency` - (Optional) The maximum number of targets this task can be run for in parallel.
* `maxErrors` - (Optional) The maximum number of errors allowed before this task stops being scheduled.
* `cutoffBehavior` - (Optional) Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are `CONTINUE_TASK` and `CANCEL_TASK`.
* `taskType` - (Required) The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
* `taskArn` - (Required) The ARN of the task to execute.
* `serviceRoleArn` - (Optional) The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
* `name` - (Optional) The name of the maintenance window task.
* `description` - (Optional) The description of the maintenance window task.
* `targets` - (Optional) The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
* `priority` - (Optional) The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
* `taskInvocationParameters` - (Optional) Configuration block with parameters for task execution.

`taskInvocationParameters` supports the following:

* `automationParameters` - (Optional) The parameters for an AUTOMATION task type. Documented below.
* `lambdaParameters` - (Optional) The parameters for a LAMBDA task type. Documented below.
* `runCommandParameters` - (Optional) The parameters for a RUN_COMMAND task type. Documented below.
* `stepFunctionsParameters` - (Optional) The parameters for a STEP_FUNCTIONS task type. Documented below.

`automationParameters` supports the following:

* `documentVersion` - (Optional) The version of an Automation document to use during task execution.
* `parameter` - (Optional) The parameters for the RUN_COMMAND task execution. Documented below.

`lambdaParameters` supports the following:

* `clientContext` - (Optional) Pass client-specific information to the Lambda function that you are invoking.
* `payload` - (Optional) JSON to provide to your Lambda function as input.
* `qualifier` - (Optional) Specify a Lambda function version or alias name.

`runCommandParameters` supports the following:

* `comment` - (Optional) Information about the command(s) to execute.
* `documentHash` - (Optional) The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
* `documentHashType` - (Optional) SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: `Sha256` and `Sha1`
* `notificationConfig` - (Optional) Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
* `outputS3Bucket` - (Optional) The name of the Amazon S3 bucket.
* `outputS3KeyPrefix` - (Optional) The Amazon S3 bucket subfolder.
* `parameter` - (Optional) The parameters for the RUN_COMMAND task execution. Documented below.
* `serviceRoleArn` - (Optional) The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
* `timeoutSeconds` - (Optional) If this time is reached and the command has not already started executing, it doesn't run.
* `cloudwatchConfig` - (Optional) Configuration options for sending command output to CloudWatch Logs. Documented below.

`stepFunctionsParameters` supports the following:

* `input` - (Optional) The inputs for the STEP_FUNCTION task.
* `name` - (Optional) The name of the STEP_FUNCTION task.

`notificationConfig` supports the following:

* `notificationArn` - (Optional) An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
* `notificationEvents` - (Optional) The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
* `notificationType` - (Optional) When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`

`cloudwatchConfig` supports the following:

* `cloudwatchLogGroupName` - (Optional) The name of the CloudWatch log group where you want to send command output. If you don't specify a group name, Systems Manager automatically creates a log group for you. The log group uses the following naming format: aws/ssm/SystemsManagerDocumentName.
* `cloudwatchOutputEnabled` - (Optional) Enables Systems Manager to send command output to CloudWatch Logs.

`parameter` supports the following:

* `name` - (Required) The parameter name.
* `values` - (Required) The array of strings.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the maintenance window task.
* `id` - The ID of the maintenance window task.
* `windowTaskId` - The ID of the maintenance window task.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AWS Maintenance Window Task using the `windowId` and `windowTaskId` separated by `/`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmMaintenanceWindowTask } from "./.gen/providers/aws/ssm-maintenance-window-task";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SsmMaintenanceWindowTask.generateConfigForImport(
      this,
      "task",
      "<window_id>/<window_task_id>"
    );
  }
}

```

Using `terraform import`, import AWS Maintenance Window Task using the `windowId` and `windowTaskId` separated by `/`. For example:

```console
% terraform import aws_ssm_maintenance_window_task.task <window_id>/<window_task_id>
```

<!-- cache-key: cdktf-0.20.8 input-a15093cae75d3ea85c62bc1be18d18e6c07e1c8cee1c561d74a33f9d0e518a2f -->