---
subcategory: "Application Auto Scaling"
layout: "aws"
page_title: "AWS: aws_appautoscaling_policy"
description: |-
  Provides an Application AutoScaling Policy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_appautoscaling_policy

Provides an Application AutoScaling Policy resource.

## Example Usage

### DynamoDB Table Autoscaling

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppautoscalingPolicy } from "./.gen/providers/aws/appautoscaling-policy";
import { AppautoscalingTarget } from "./.gen/providers/aws/appautoscaling-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const dynamodbTableReadTarget = new AppautoscalingTarget(
      this,
      "dynamodb_table_read_target",
      {
        maxCapacity: 100,
        minCapacity: 5,
        resourceId: "table/tableName",
        scalableDimension: "dynamodb:table:ReadCapacityUnits",
        serviceNamespace: "dynamodb",
      }
    );
    new AppautoscalingPolicy(this, "dynamodb_table_read_policy", {
      name:
        "DynamoDBReadCapacityUtilization:${" +
        dynamodbTableReadTarget.resourceId +
        "}",
      policyType: "TargetTrackingScaling",
      resourceId: dynamodbTableReadTarget.resourceId,
      scalableDimension: dynamodbTableReadTarget.scalableDimension,
      serviceNamespace: dynamodbTableReadTarget.serviceNamespace,
      targetTrackingScalingPolicyConfiguration: {
        predefinedMetricSpecification: {
          predefinedMetricType: "DynamoDBReadCapacityUtilization",
        },
        targetValue: 70,
      },
    });
  }
}

```

### ECS Service Autoscaling

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Op, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppautoscalingPolicy } from "./.gen/providers/aws/appautoscaling-policy";
import { AppautoscalingTarget } from "./.gen/providers/aws/appautoscaling-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const ecsTarget = new AppautoscalingTarget(this, "ecs_target", {
      maxCapacity: 4,
      minCapacity: 1,
      resourceId: "service/clusterName/serviceName",
      scalableDimension: "ecs:service:DesiredCount",
      serviceNamespace: "ecs",
    });
    new AppautoscalingPolicy(this, "ecs_policy", {
      name: "scale-down",
      policyType: "StepScaling",
      resourceId: ecsTarget.resourceId,
      scalableDimension: ecsTarget.scalableDimension,
      serviceNamespace: ecsTarget.serviceNamespace,
      stepScalingPolicyConfiguration: {
        adjustmentType: "ChangeInCapacity",
        cooldown: 60,
        metricAggregationType: "Maximum",
        stepAdjustment: [
          {
            metricIntervalUpperBound: Token.asString(0),
            scalingAdjustment: Token.asNumber(Op.negate(1)),
          },
        ],
      },
    });
  }
}

```

### Preserve desired count when updating an autoscaled ECS Service

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcsService(this, "ecs_service", {
      cluster: "clusterName",
      desiredCount: 2,
      lifecycle: {
        ignoreChanges: [desiredCount],
      },
      name: "serviceName",
      taskDefinition: "taskDefinitionFamily:1",
    });
  }
}

```

### Aurora Read Replica Autoscaling

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppautoscalingPolicy } from "./.gen/providers/aws/appautoscaling-policy";
import { AppautoscalingTarget } from "./.gen/providers/aws/appautoscaling-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const replicas = new AppautoscalingTarget(this, "replicas", {
      maxCapacity: 15,
      minCapacity: 1,
      resourceId: "cluster:${" + example.id + "}",
      scalableDimension: "rds:cluster:ReadReplicaCount",
      serviceNamespace: "rds",
    });
    const awsAppautoscalingPolicyReplicas = new AppautoscalingPolicy(
      this,
      "replicas_1",
      {
        name: "cpu-auto-scaling",
        policyType: "TargetTrackingScaling",
        resourceId: replicas.resourceId,
        scalableDimension: replicas.scalableDimension,
        serviceNamespace: replicas.serviceNamespace,
        targetTrackingScalingPolicyConfiguration: {
          predefinedMetricSpecification: {
            predefinedMetricType: "RDSReaderAverageCPUUtilization",
          },
          scaleInCooldown: 300,
          scaleOutCooldown: 300,
          targetValue: 75,
        },
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsAppautoscalingPolicyReplicas.overrideLogicalId("replicas");
  }
}

```

### Create target tracking scaling policy using metric math

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppautoscalingPolicy } from "./.gen/providers/aws/appautoscaling-policy";
import { AppautoscalingTarget } from "./.gen/providers/aws/appautoscaling-target";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const ecsTarget = new AppautoscalingTarget(this, "ecs_target", {
      maxCapacity: 4,
      minCapacity: 1,
      resourceId: "service/clusterName/serviceName",
      scalableDimension: "ecs:service:DesiredCount",
      serviceNamespace: "ecs",
    });
    new AppautoscalingPolicy(this, "example", {
      name: "foo",
      policyType: "TargetTrackingScaling",
      resourceId: ecsTarget.resourceId,
      scalableDimension: ecsTarget.scalableDimension,
      serviceNamespace: ecsTarget.serviceNamespace,
      targetTrackingScalingPolicyConfiguration: {
        customizedMetricSpecification: {
          metrics: [
            {
              id: "m1",
              label:
                "Get the queue size (the number of messages waiting to be processed)",
              metricStat: {
                metric: {
                  dimensions: [
                    {
                      name: "QueueName",
                      value: "my-queue",
                    },
                  ],
                  metricName: "ApproximateNumberOfMessagesVisible",
                  namespace: "AWS/SQS",
                },
                stat: "Sum",
              },
              returnData: false,
            },
            {
              id: "m2",
              label:
                "Get the ECS running task count (the number of currently running tasks)",
              metricStat: {
                metric: {
                  dimensions: [
                    {
                      name: "ClusterName",
                      value: "default",
                    },
                    {
                      name: "ServiceName",
                      value: "web-app",
                    },
                  ],
                  metricName: "RunningTaskCount",
                  namespace: "ECS/ContainerInsights",
                },
                stat: "Average",
              },
              returnData: false,
            },
            {
              expression: "m1 / m2",
              id: "e1",
              label: "Calculate the backlog per instance",
              returnData: true,
            },
          ],
        },
        targetValue: 100,
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the policy. Must be between 1 and 255 characters in length.
* `policyType` - (Optional) Policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
* `resourceId` - (Required) Resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html)
* `scalableDimension` - (Required) Scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html)
* `serviceNamespace` - (Required) AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html)
* `stepScalingPolicyConfiguration` - (Optional) Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
* `targetTrackingScalingPolicyConfiguration` - (Optional) Target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.

### step_scaling_policy_configuration

The `stepScalingPolicyConfiguration` configuration block supports the following arguments:

* `adjustmentType` - (Required) Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
* `cooldown` - (Required) Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
* `metricAggregationType` - (Optional) Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
* `minAdjustmentMagnitude` - (Optional) Minimum number to adjust your scalable dimension as a result of a scaling activity. If the adjustment type is PercentChangeInCapacity, the scaling policy changes the scalable dimension of the scalable target by this amount.
* `stepAdjustment` - (Optional) Set of adjustments that manage scaling. These have the following structure:

 ```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Op, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppautoscalingPolicy } from "./.gen/providers/aws/appautoscaling-policy";
interface MyConfig {
  name: any;
  resourceId: any;
  scalableDimension: any;
  serviceNamespace: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new AppautoscalingPolicy(this, "ecs_policy", {
      stepScalingPolicyConfiguration: {
        stepAdjustment: [
          {
            metricIntervalLowerBound: Token.asString(1),
            metricIntervalUpperBound: Token.asString(2),
            scalingAdjustment: Token.asNumber(Op.negate(1)),
          },
          {
            metricIntervalLowerBound: Token.asString(2),
            metricIntervalUpperBound: Token.asString(3),
            scalingAdjustment: 1,
          },
        ],
      },
      name: config.name,
      resourceId: config.resourceId,
      scalableDimension: config.scalableDimension,
      serviceNamespace: config.serviceNamespace,
    });
  }
}

```

* `metricIntervalLowerBound` - (Optional) Lower bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as negative infinity.
* `metricIntervalUpperBound` - (Optional) Upper bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as infinity. The upper bound must be greater than the lower bound.
* `scalingAdjustment` - (Required) Number of members by which to scale, when the adjustment bounds are breached. A positive value scales up. A negative value scales down.

### target_tracking_scaling_policy_configuration

The `targetTrackingScalingPolicyConfiguration` configuration block supports the following arguments:

* `targetValue` - (Required) Target value for the metric.
* `disableScaleIn` - (Optional) Whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is `false`.
* `scaleInCooldown` - (Optional) Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
* `scaleOutCooldown` - (Optional) Amount of time, in seconds, after a scale out activity completes before another scale out activity can start.
* `customizedMetricSpecification` - (Optional) Custom CloudWatch metric. Documentation can be found  at: [AWS Customized Metric Specification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CustomizedMetricSpecification.html). See supported fields below.
* `predefinedMetricSpecification` - (Optional) Predefined metric. See supported fields below.

### target_tracking_scaling_policy_configuration customized_metric_specification

Example usage:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppautoscalingPolicy } from "./.gen/providers/aws/appautoscaling-policy";
interface MyConfig {
  name: any;
  resourceId: any;
  scalableDimension: any;
  serviceNamespace: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new AppautoscalingPolicy(this, "example", {
      policyType: "TargetTrackingScaling",
      targetTrackingScalingPolicyConfiguration: {
        customizedMetricSpecification: {
          dimensions: [
            {
              name: "MyOptionalMetricDimensionName",
              value: "MyOptionalMetricDimensionValue",
            },
          ],
          metricName: "MyUtilizationMetric",
          namespace: "MyNamespace",
          statistic: "Average",
          unit: "Percent",
        },
        targetValue: 40,
      },
      name: config.name,
      resourceId: config.resourceId,
      scalableDimension: config.scalableDimension,
      serviceNamespace: config.serviceNamespace,
    });
  }
}

```

The `targetTrackingScalingPolicyConfiguration` `customizedMetricSpecification` configuration block supports the following arguments:

* `dimensions` - (Optional) Configuration block(s) with the dimensions of the metric if the metric was published with dimensions. Detailed below.
* `metricName` - (Optional) Name of the metric.
* `namespace` - (Optional) Namespace of the metric.
* `statistic` - (Optional) Statistic of the metric. Valid values: `Average`, `Minimum`, `Maximum`, `SampleCount`, and `Sum`.
* `unit` - (Optional) Unit of the metric.
* `metrics` - (Optional) Metrics to include, as a metric data query.

### target_tracking_scaling_policy_configuration customized_metric_specification dimensions

The `targetTrackingScalingPolicyConfiguration` `customizedMetricSpecification` `dimensions` configuration block supports the following arguments:

* `name` - (Required) Name of the dimension.
* `value` - (Required) Value of the dimension.

### target_tracking_scaling_policy_configuration customized_metric_specification metrics

The `targetTrackingScalingPolicyConfiguration` `customizedMetricSpecification` `metrics` configuration block supports the following arguments:

* `expression` - (Optional) Math expression used on the returned metric. You must specify either `expression` or `metricStat`, but not both.
* `id` - (Required) Short name for the metric used in target tracking scaling policy.
* `label` - (Optional) Human-readable label for this metric or expression.
* `metricStat` - (Optional) Structure that defines CloudWatch metric to be used in target tracking scaling policy. You must specify either `expression` or `metricStat`, but not both.
* `returnData` - (Optional) Boolean that indicates whether to return the timestamps and raw data values of this metric, the default is true

### target_tracking_scaling_policy_configuration customized_metric_specification metrics metric_stat

The `targetTrackingScalingPolicyConfiguration` `customizedMetricSpecification` `metrics` `metricStat` configuration block supports the following arguments:

* `metric` - (Required) Structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
* `stat` - (Required) Statistic of the metrics to return.
* `unit` - (Optional) Unit of the metrics to return.

### target_tracking_scaling_policy_configuration customized_metric_specification metrics metric

The `targetTrackingScalingPolicyConfiguration` `customizedMetricSpecification` `metrics` `metric` configuration block supports the following arguments:

* `dimensions` - (Optional) Dimensions of the metric.
* `metricName` - (Required) Name of the metric.
* `namespace` - (Required) Namespace of the metric.

### target_tracking_scaling_policy_configuration customized_metric_specification metrics dimensions

The `targetTrackingScalingPolicyConfiguration` `customizedMetricSpecification` `metrics` `dimensions` configuration block supports the following arguments:

* `name` - (Required) Name of the dimension.
* `value` - (Required) Value of the dimension.

### target_tracking_scaling_policy_configuration predefined_metric_specification

The `targetTrackingScalingPolicyConfiguration` `predefinedMetricSpecification` configuration block supports the following arguments:

* `predefinedMetricType` - (Required) Metric type.
* `resourceLabel` - (Optional) Reserved for future use if the `predefinedMetricType` is not `ALBRequestCountPerTarget`. If the `predefinedMetricType` is `ALBRequestCountPerTarget`, you must specify this argument. Documentation can be found at: [AWS Predefined Scaling Metric Specification](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_PredefinedScalingMetricSpecification.html). Must be less than or equal to 1023 characters in length.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `alarmArns` - List of CloudWatch alarm ARNs associated with the scaling policy.
* `arn` - ARN assigned by AWS to the scaling policy.
* `name` - Scaling policy's name.
* `policyType` - Scaling policy's type.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Application AutoScaling Policy using the `service-namespace` , `resource-id`, `scalable-dimension` and `policy-name` separated by `/`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppautoscalingPolicy } from "./.gen/providers/aws/appautoscaling-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    AppautoscalingPolicy.generateConfigForImport(
      this,
      "testPolicy",
      "service-namespace/resource-id/scalable-dimension/policy-name"
    );
  }
}

```

Using `terraform import`, import Application AutoScaling Policy using the `service-namespace` , `resource-id`, `scalable-dimension` and `policy-name` separated by `/`. For example:

```console
% terraform import aws_appautoscaling_policy.test-policy service-namespace/resource-id/scalable-dimension/policy-name
```

<!-- cache-key: cdktf-0.20.8 input-2743c6c027c30780d655bad04c00aa7e64161c9de89bbec5f6090657d7402e8c -->