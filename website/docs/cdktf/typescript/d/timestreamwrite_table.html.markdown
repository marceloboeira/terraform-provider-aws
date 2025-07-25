---
subcategory: "Timestream Write"
layout: "aws"
page_title: "AWS: aws_timestreamwrite_table"
description: |-
  Terraform data source for managing an AWS Timestream Write Table.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_timestreamwrite_table

Terraform data source for managing an AWS Timestream Write Table.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsTimestreamwriteTable } from "./.gen/providers/aws/data-aws-timestreamwrite-table";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsTimestreamwriteTable(this, "test", {
      databaseName: Token.asString(awsTimestreamwriteDatabaseTest.databaseName),
      name: Token.asString(awsTimestreamwriteTableTest.tableName),
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `databaseName` - (Required) Name of the Timestream database.
* `name` - (Required) Name of the Timestream table.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN that uniquely identifies the table.
* `creationTime` - Time that table was created.
* `databaseName` - Name of database.
* `lastUpdatedTime` - Last time table was updated.
* `magneticStoreWriteProperties` - Object containing the following attributes to desribe magnetic store writes.
    * `enableMagneticStoreWrites` - Flag that is set based on if magnetic store writes are enabled.
    * `magneticStoreRejectedDataLocation` - Object containing the following attributes to describe error reports for records rejected during magnetic store writes.
        * `s3Configuration` - Object containing the following attributes to describe the configuration of an s3 location to write error reports for records rejected.
            * `bucketName` - Name of S3 bucket.
            * `encryption_object` - Encryption option for  S3 location.
            * `kmsKeyId` - AWS KMS key ID for S3 location with AWS maanged key.
            * `objectKeyPrefix` -  Object key preview for S3 location.
* `retentionProperties` -  Object containing the following attributes to describe the retention duration for the memory and magnetic stores.
    * `magneticStoreRetentionPeriodInDays` - Duration in days in which the data must be stored in magnetic store.
    * `memoryStoreRetentionPeriodInHours` - Duration in hours in which the data must be stored in memory store.
* `schema` -  Object containing the following attributes to describe the schema of the table.
    * `type` - Type of partition key.
    * `partitionKey` - Level of enforcement for the specification of a dimension key in ingested records.
    * `name` - Name of the timestream attribute used for a dimension key.
* `name` - Name of the table.
* `tableStatus` - Current state of table.

<!-- cache-key: cdktf-0.20.8 input-b0937170355b75a46b1d7a1819d28a5088d7690e859dcbec324ce944e5b08438 -->