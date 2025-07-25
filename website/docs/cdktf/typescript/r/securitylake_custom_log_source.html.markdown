---
subcategory: "Security Lake"
layout: "aws"
page_title: "AWS: aws_securitylake_custom_log_source"
description: |-
  Terraform resource for managing an AWS Security Lake Custom Log Source.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_securitylake_custom_log_source

Terraform resource for managing an AWS Security Lake Custom Log Source.

~> **NOTE:** The underlying `aws_securitylake_data_lake` must be configured before creating the `aws_securitylake_custom_log_source`. Use a `dependsOn` statement.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecuritylakeCustomLogSource } from "./.gen/providers/aws/securitylake-custom-log-source";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SecuritylakeCustomLogSource(this, "example", {
      configuration: [
        {
          crawlerConfiguration: [
            {
              roleArn: customLog.arn,
            },
          ],
          providerIdentity: [
            {
              externalId: "example-id",
              principal: "123456789012",
            },
          ],
        },
      ],
      dependsOn: [awsSecuritylakeDataLakeExample],
      eventClasses: ["FILE_ACTIVITY"],
      sourceName: "example-name",
      sourceVersion: "1.0",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `configuration` - (Required) The configuration for the third-party custom source.
    * `crawlerConfiguration` - (Required) The configuration for the Glue Crawler for the third-party custom source.
        * `roleArn` - (Required) The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role to be used by the AWS Glue crawler.
    * `providerIdentity` - (Required) The identity of the log provider for the third-party custom source.
        * `externalId` - (Required) The external ID used to estalish trust relationship with the AWS identity.
        * `principal` - (Required) The AWS identity principal.
* `eventClasses` - (Optional) The Open Cybersecurity Schema Framework (OCSF) event classes which describes the type of data that the custom source will send to Security Lake.
* `sourceName` - (Required) Specify the name for a third-party custom source.
  This must be a Regionally unique value.
  Has a maximum length of 20.
* `sourceVersion` - (Optional) Specify the source version for the third-party custom source, to limit log collection to a specific version of custom data source.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `attributes` - The attributes of a third-party custom source.
    * `crawler_arn` - The ARN of the AWS Glue crawler.
    * `database_arn` - The ARN of the AWS Glue database where results are written.
    * `tableArn` - The ARN of the AWS Glue table.
* `providerDetails` - The details of the log provider for a third-party custom source.
    * `location` - The location of the partition in the Amazon S3 bucket for Security Lake.
    * `roleArn` - The ARN of the IAM role to be used by the entity putting logs into your custom source partition.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AWS log sources using the source name. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecuritylakeCustomLogSource } from "./.gen/providers/aws/securitylake-custom-log-source";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecuritylakeCustomLogSource.generateConfigForImport(
      this,
      "example",
      "example-name"
    );
  }
}

```

Using `terraform import`, import Custom log sources using the source name. For example:

```console
% terraform import aws_securitylake_custom_log_source.example example-name
```

<!-- cache-key: cdktf-0.20.8 input-9c50acd5144c36e9f75a2a9244059b6a5445fddf18b179c57a4aa62e47427f2c -->