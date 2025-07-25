---
subcategory: "AppStream 2.0"
layout: "aws"
page_title: "AWS: aws_appstream_fleet"
description: |-
  Provides an AppStream fleet
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_appstream_fleet

Provides an AppStream fleet.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppstreamFleet } from "./.gen/providers/aws/appstream-fleet";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AppstreamFleet(this, "test_fleet", {
      computeCapacity: {
        desiredInstances: 1,
      },
      description: "test fleet",
      displayName: "test-fleet",
      enableDefaultInternetAccess: false,
      fleetType: "ON_DEMAND",
      idleDisconnectTimeoutInSeconds: 60,
      imageName: "Amazon-AppStream2-Sample-Image-03-11-2023",
      instanceType: "stream.standard.large",
      maxUserDurationInSeconds: 600,
      name: "test-fleet",
      tags: {
        TagName: "tag-value",
      },
      vpcConfig: {
        subnetIds: ["subnet-06e9b13400c225127"],
      },
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `computeCapacity` - (Required) Configuration block for the desired capacity of the fleet. See below.
* `instanceType` - (Required) Instance type to use when launching fleet instances.
* `name` - (Required) Unique name for the fleet.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) Description to display.
* `disconnectTimeoutInSeconds` - (Optional) Amount of time that a streaming session remains active after users disconnect.
* `displayName` - (Optional) Human-readable friendly name for the AppStream fleet.
* `domainJoinInfo` - (Optional) Configuration block for the name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain. See below.
* `enableDefaultInternetAccess` - (Optional) Enables or disables default internet access for the fleet.
* `fleetType` - (Optional) Fleet type. Valid values are: `ON_DEMAND`, `ALWAYS_ON`
* `iamRoleArn` - (Optional) ARN of the IAM role to apply to the fleet.
* `idleDisconnectTimeoutInSeconds` - (Optional) Amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the `disconnectTimeoutInSeconds` time interval begins. Defaults to `0`. Valid value is between `60` and `3600 `seconds.
* `imageName` - (Optional) Name of the image used to create the fleet.
* `imageArn` - (Optional) ARN of the public, private, or shared image to use.
* `streamView` - (Optional) AppStream 2.0 view that is displayed to your users when they stream from the fleet. When `APP` is specified, only the windows of applications opened by users display. When `DESKTOP` is specified, the standard desktop that is provided by the operating system displays. If not specified, defaults to `APP`.
* `maxSessionsPerInstance` - (Optional) The maximum number of user sessions on an instance. This only applies to multi-session fleets.
* `maxUserDurationInSeconds` - (Optional) Maximum amount of time that a streaming session can remain active, in seconds.
* `vpcConfig` - (Optional) Configuration block for the VPC configuration for the image builder. See below.
* `tags` - (Optional) Map of tags to attach to AppStream instances.

### `computeCapacity`

Exactly one of `desiredInstances` or `desiredSessions` must be set, based on the type of fleet being created.

* `desiredInstances` - (Optional) Desired number of streaming instances.
* `desiredSessions` - (Optional) Desired number of user sessions for a multi-session fleet. This is not allowed for single-session fleets.

### `domainJoinInfo`

* `directoryName` - (Optional) Fully qualified name of the directory (for example, corp.example.com).
* `organizationalUnitDistinguishedName` - (Optional) Distinguished name of the organizational unit for computer accounts.

### `vpcConfig`

* `securityGroupIds` - Identifiers of the security groups for the fleet or image builder.
* `subnetIds` - Identifiers of the subnets to which a network interface is attached from the fleet instance or image builder instance.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Unique identifier (ID) of the appstream fleet.
* `arn` - ARN of the appstream fleet.
* `state` - State of the fleet. Can be `STARTING`, `RUNNING`, `STOPPING` or `STOPPED`
* `createdTime` -  Date and time, in UTC and extended RFC 3339 format, when the fleet was created.
* `computeCapacity` - Describes the capacity status for a fleet.

### `computeCapacity`

* `available` - Number of currently available instances that can be used to stream sessions.
* `inUse` - Number of instances in use for streaming.
* `running` - Total number of simultaneous streaming instances that are running.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_appstream_fleet` using the id. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppstreamFleet } from "./.gen/providers/aws/appstream-fleet";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    AppstreamFleet.generateConfigForImport(this, "example", "fleetNameExample");
  }
}

```

Using `terraform import`, import `aws_appstream_fleet` using the id. For example:

```console
% terraform import aws_appstream_fleet.example fleetNameExample
```

<!-- cache-key: cdktf-0.20.8 input-4bf23b179acf43136f029e145655ce0ea59494badf16f0de0825bb64c13edf31 -->