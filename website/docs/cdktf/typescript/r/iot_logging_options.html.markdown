---
subcategory: "IoT Core"
layout: "aws"
page_title: "AWS: aws_iot_logging_options"
description: |-
    Provides a resource to manage default logging options.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_iot_logging_options

Provides a resource to manage [default logging options](https://docs.aws.amazon.com/iot/latest/developerguide/configure-logging.html#configure-logging-console).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { IotLoggingOptions } from "./.gen/providers/aws/iot-logging-options";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new IotLoggingOptions(this, "example", {
      defaultLogLevel: "WARN",
      roleArn: Token.asString(awsIamRoleExample.arn),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `defaultLogLevel` - (Optional) The default logging level. Valid Values: `"DEBUG"`, `"INFO"`, `"ERROR"`, `"WARN"`, `"DISABLED"`.
* `disableAllLogs` - (Optional) If `true` all logs are disabled. The default is `false`.
* `roleArn` - (Required) The ARN of the role that allows IoT to write to Cloudwatch logs.

## Attribute Reference

This resource exports no additional attributes.

<!-- cache-key: cdktf-0.20.8 input-1d008a229305c1a5763a924f49f4ed0a2cf09a0dbd05e7a519d3adb9058d75f5 -->