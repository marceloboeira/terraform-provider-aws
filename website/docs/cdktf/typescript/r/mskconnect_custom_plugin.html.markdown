---
subcategory: "Managed Streaming for Kafka Connect"
layout: "aws"
page_title: "AWS: aws_mskconnect_custom_plugin"
description: |-
  Provides an Amazon MSK Connect custom plugin resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_mskconnect_custom_plugin

Provides an Amazon MSK Connect Custom Plugin Resource.

## Example Usage

### Basic configuration

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { MskconnectCustomPlugin } from "./.gen/providers/aws/mskconnect-custom-plugin";
import { S3Bucket } from "./.gen/providers/aws/s3-bucket";
import { S3Object } from "./.gen/providers/aws/s3-object";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new S3Bucket(this, "example", {
      bucket: "example",
    });
    const awsS3ObjectExample = new S3Object(this, "example_1", {
      bucket: example.id,
      key: "debezium.zip",
      source: "debezium.zip",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3ObjectExample.overrideLogicalId("example");
    const awsMskconnectCustomPluginExample = new MskconnectCustomPlugin(
      this,
      "example_2",
      {
        contentType: "ZIP",
        location: {
          s3: {
            bucketArn: example.arn,
            fileKey: Token.asString(awsS3ObjectExample.key),
          },
        },
        name: "debezium-example",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsMskconnectCustomPluginExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required, Forces new resource) The name of the custom plugin..
* `contentType` - (Required, Forces new resource) The type of the plugin file. Allowed values are `ZIP` and `JAR`.
* `description` - (Optional, Forces new resource) A summary description of the custom plugin.
* `location` - (Required, Forces new resource) Information about the location of a custom plugin. See [`location` Block](#location-block) for details.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### `location` Block

The `location` configuration block supports the following arguments:

* `s3` - (Required, Forces new resource) Information of the plugin file stored in Amazon S3. See [`s3` Block](#s3-block) for details..

### `s3` Block

The `s3` configuration Block supports the following arguments:

* `bucketArn` - (Required, Forces new resource) The Amazon Resource Name (ARN) of an S3 bucket.
* `fileKey` - (Required, Forces new resource) The file key for an object in an S3 bucket.
* `objectVersion` - (Optional, Forces new resource) The version of an object in an S3 bucket.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - the Amazon Resource Name (ARN) of the custom plugin.
* `latestRevision` - an ID of the latest successfully created revision of the custom plugin.
* `state` - the state of the custom plugin.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import MSK Connect Custom Plugin using the plugin's `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { MskconnectCustomPlugin } from "./.gen/providers/aws/mskconnect-custom-plugin";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    MskconnectCustomPlugin.generateConfigForImport(
      this,
      "example",
      "arn:aws:kafkaconnect:eu-central-1:123456789012:custom-plugin/debezium-example/abcdefgh-1234-5678-9abc-defghijklmno-4"
    );
  }
}

```

Using `terraform import`, import MSK Connect Custom Plugin using the plugin's `arn`. For example:

```console
% terraform import aws_mskconnect_custom_plugin.example 'arn:aws:kafkaconnect:eu-central-1:123456789012:custom-plugin/debezium-example/abcdefgh-1234-5678-9abc-defghijklmno-4'
```

<!-- cache-key: cdktf-0.20.8 input-0f9575ee42765d36c7f1492f28e6e1300d8185461a507593380bcdc5ef298a83 -->