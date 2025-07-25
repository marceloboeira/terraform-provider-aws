---
subcategory: "AppFabric"
layout: "aws"
page_title: "AWS: aws_appfabric_ingestion_destination"
description: |-
  Terraform resource for managing an AWS AppFabric Ingestion Destination.
---

# Resource: aws_appfabric_ingestion_destination

Terraform resource for managing an AWS AppFabric Ingestion Destination.

## Example Usage

### Basic Usage

```terraform
resource "aws_appfabric_ingestion_destination" "example" {
  app_bundle_arn = aws_appfabric_app_bundle.example.arn
  ingestion_arn  = aws_appfabric_ingestion.example.arn

  processing_configuration {
    audit_log {
      format = "json"
      schema = "raw"
    }
  }

  destination_configuration {
    audit_log {
      destination {
        s3_bucket {
          bucket_name = aws_s3_bucket.example.bucket
        }
      }
    }
  }
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `app_bundle_arn` - (Required) The Amazon Resource Name (ARN) of the app bundle to use for the request.
* `ingestion_arn` - (Required) The Amazon Resource Name (ARN) of the ingestion to use for the request.
* `destination_configuration` - (Required) Contains information about the destination of ingested data.
* `processing_configuration` - (Required) Contains information about how ingested data is processed.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

Destination Configuration support the following:

* `audit_log` - (Required) Contains information about an audit log destination configuration.

Audit Log Destination Configuration support the following:

* `destination` - (Required) Contains information about an audit log destination. Only one destination (Firehose Stream) or (S3 Bucket) can be specified.

Destination support the following:

* `firehose_stream` - (Optional) Contains information about an Amazon Data Firehose delivery stream.
* `s3_bucket` - (Optional) Contains information about an Amazon S3 bucket.

Firehose Stream support the following:

* `streamName` - (Required) The name of the Amazon Data Firehose delivery stream.

S3 Bucket support the following:

* `bucketName` - (Required) The name of the Amazon S3 bucket.
* `prefix` - (Optional) The object key to use.

Processing Configuration support the following:

* `audit_log` - (Required) Contains information about an audit log processing configuration.

Audit Log Processing Configuration support the following:

* `format` - (Required) The format in which the audit logs need to be formatted. Valid values: `json`, `parquet`.
* `schema` - (Required) The event schema in which the audit logs need to be formatted. Valid values: `ocsf`, `raw`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Ingestion Destination.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `5m`)
* `update` - (Default `5m`)
* `delete` - (Default `5m`)
