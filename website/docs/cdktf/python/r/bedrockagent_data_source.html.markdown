---
subcategory: "Bedrock Agents"
layout: "aws"
page_title: "AWS: aws_bedrockagent_data_source"
description: |-
  Terraform resource for managing an AWS Agents for Amazon Bedrock Data Source.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_bedrockagent_data_source

Terraform resource for managing an AWS Agents for Amazon Bedrock Data Source.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.bedrockagent_data_source import BedrockagentDataSource
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BedrockagentDataSource(self, "example",
            data_source_configuration=[BedrockagentDataSourceDataSourceConfiguration(
                s3_configuration=[BedrockagentDataSourceDataSourceConfigurationS3Configuration(
                    bucket_arn="arn:aws:s3:::example-bucket"
                )
                ],
                type="S3"
            )
            ],
            knowledge_base_id="EMDPPAYPZI",
            name="example"
        )
```

## Argument Reference

The following arguments are required:

* `data_source_configuration` - (Required) Details about how the data source is stored. See [`data_source_configuration` block](#data_source_configuration-block) for details.
* `knowledge_base_id` - (Required) Unique identifier of the knowledge base to which the data source belongs.
* `name` - (Required, Forces new resource) Name of the data source.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `data_deletion_policy` - (Optional) Data deletion policy for a data source. Valid values: `RETAIN`, `DELETE`.
* `description` - (Optional) Description of the data source.
* `server_side_encryption_configuration` - (Optional) Details about the configuration of the server-side encryption. See [`server_side_encryption_configuration` block](#server_side_encryption_configuration-block) for details.
* `vector_ingestion_configuration` - (Optional, Forces new resource) Details about the configuration of the server-side encryption. See [`vector_ingestion_configuration` block](#vector_ingestion_configuration-block) for details.

### `data_source_configuration` block

The `data_source_configuration` configuration block supports the following arguments:

* `type` - (Required) Type of storage for the data source. Valid values: `S3`, `WEB`, `CONFLUENCE`, `SALESFORCE`, `SHAREPOINT`, `CUSTOM`, `REDSHIFT_METADATA`.
* `confluence_configuration` - (Optional) Details about the configuration of the Confluence data source. See [`confluence_data_source_configuration` block](#confluence_data_source_configuration-block) for details.
* `s3_configuration` - (Optional) Details about the configuration of the S3 object containing the data source. See [`s3_data_source_configuration` block](#s3_data_source_configuration-block) for details.
* `salesforce_configuration` - (Optional) Details about the configuration of the Salesforce data source. See [`salesforce_data_source_configuration` block](#salesforce_data_source_configuration-block) for details.
* `share_point_configuration` - (Optional) Details about the configuration of the SharePoint data source. See [`share_point_data_source_configuration` block](#share_point_data_source_configuration-block) for details.
* `web_configuration` - (Optional) Details about the configuration of the web data source. See [`web_data_source_configuration` block](#web_data_source_configuration-block) for details.

### `confluence_data_source_configuration` block

The `confluence_data_source_configuration` configuration block supports the following arguments:

* `source_configuration` - (Required) The endpoint information to connect to your Confluence data source. See [`source_configuration` block](#confluence-source_configuration-block) for details.
* `crawler_configuration` - (Optional) Configuration for Confluence content. See [`crawler_configuration` block](#crawler_configuration-block) for details.

For more details, see the [Amazon BedrockAgent Confluence documentation][1].

### Confluence `source_configuration` block

The `source_configuration` configuration block supports the following arguments:

* `auth_type` - (Required) The supported authentication type to authenticate and connect to your Confluence instance. Valid values: `BASIC`, `OAUTH2_CLIENT_CREDENTIALS`.
* `credentials_secret_arn` - (Required) The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Confluence instance URL. For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see Confluence connection configuration. Pattern: ^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$.
* `host_type` - (Required) The supported host type, whether online/cloud or server/on-premises. Valid values: `SAAS`.
* `host_url` - (Required) The Confluence host URL or instance URL. Pattern: `^https://[A-Za-z0-9][^\s]*$`.

### `s3_data_source_configuration` block

The `s3_data_source_configuration` configuration block supports the following arguments:

* `bucket_arn` - (Required) ARN of the bucket that contains the data source.
* `bucket_owner_account_id` - (Optional) Bucket account owner ID for the S3 bucket.
* `inclusion_prefixes` - (Optional) List of S3 prefixes that define the object containing the data sources. For more information, see [Organizing objects using prefixes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-prefixes.html).

### `salesforce_data_source_configuration` block

The `salesforce_data_source_configuration` configuration block supports the following arguments:

* `source_configuration` - (Required) The endpoint information to connect to your Salesforce data source. See [`source_configuration` block](#salesforce-source_configuration-block) for details.
* `crawler_configuration` - (Optional) Configuration for Salesforce content. See [`crawler_configuration` block](#crawler_configuration-block) for details.

For more details, see the [Amazon BedrockAgent Salesforce documentation][2].

### Salesforce `source_configuration` block

The `source_configuration` configuration block supports the following arguments:

* `auth_type` - (Required) The supported authentication type to authenticate and connect to your Salesforce instance. Valid values: OAUTH2_CLIENT_CREDENTIALS.
* `credentials_secret_arn` - (Required) The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Salesforce instance URL. For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see Salesforce connection configuration. Pattern: ^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$.
* `host_url` - (Required) The Salesforce host URL or instance URL. Pattern: `^https://[A-Za-z0-9][^\s]*$`.

### `crawler_configuration` block

The `crawler_configuration` configuration block supports the following arguments:

* `filter_configuration` - (Optional) The Salesforce standard object configuration. See [`filter_configuration` block](#filter_configuration-block) for details.

### `filter_configuration` block

The `filter_configuration` configuration block supports the following arguments:

* `type` - (Required) The type of filtering that you want to apply to certain objects or content of the data source. For example, the PATTERN type is regular expression patterns you can apply to filter your content.
* `pattern_object_filter` - (Optional) The configuration of filtering certain objects or content types of the data source. See [`pattern_object_filter` block](#pattern_object_filter-block) for details.

### `pattern_object_filter` block

The `pattern_object_filter` configuration block supports the following arguments:

* `filters` - (Required) The configuration of specific filters applied to your data source content. Minimum of 1 filter and maximum of 25 filters.

Each filter object should contain the following configuration:

* `object_type` - (Required) The supported object type or content type of the data source.
* `exclusion_filters` - (Optional) A list of one or more exclusion regular expression patterns to exclude certain object types that adhere to the pattern.
* `inclusion_filters` - (Optional) A list of one or more inclusion regular expression patterns to include certain object types that adhere to the pattern.

### `share_point_data_source_configuration` block

The `share_point_data_source_configuration` configuration block supports the following arguments:

* `source_configuration` - (Required) The endpoint information to connect to your SharePoint data source. See [`source_configuration` block](#sharepoint-source_configuration-block) for details.
* `crawler_configuration` - (Optional) Configuration for SharePoint content. See [`crawler_configuration` block](#crawler_configuration-block) for details.

For more details, see the [Amazon BedrockAgent SharePoint documentation][3].

### SharePoint `source_configuration` block

The `source_configuration` configuration block supports the following arguments:

* `auth_type` - (Required) The supported authentication type to authenticate and connect to your SharePoint site. Valid values: `OAUTH2_CLIENT_CREDENTIALS`, `OAUTH2_SHAREPOINT_APP_ONLY_CLIENT_CREDENTIALS`.
* `credentials_secret_arn` - (Required) The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your SharePoint site. For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see SharePoint connection configuration. Pattern: ^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$.
* `domain` - (Required) The domain of your SharePoint instance or site URL/URLs.
* `host_type` - (Required) The supported host type, whether online/cloud or server/on-premises. Valid values: `ONLINE`.
* `site_urls` - (Required) A list of one or more SharePoint site URLs.
* `tenant_id` - (Optional) The identifier of your Microsoft 365 tenant.

### `web_data_source_configuration` block

The `web_data_source_configuration` configuration block supports the following arguments:

* `source_configuration` - (Required) Endpoint information to connect to your web data source. See [`source_configuration` block](#web-source_configuration-block) for details.
* `crawler_configuration` - (Optional) Configuration for web content. See [`crawler_configuration` block](#web-crawler_configuration-block) for details.

### Web `source_configuration` block

The `source_configuration` configuration block supports the following arguments:

* `url_configuration` - (Required) The URL configuration of your web data source. See [`url_configuration` block](#url_configuration-block) for details.

### `url_configuration` block

The `url_configuration` configuration block supports the following arguments:

* `seed_urls` - (Optional) List of one or more seed URLs to crawl. See [`seed_urls` block](#seed_urls-block) for details.

### `seed_urls` block

The `seed_urls` configuration block supports the following arguments:

* `url` - (Optional) Seed or starting point URL. Must match the pattern `^https?://[A-Za-z0-9][^\s]*$`.

### Web `crawler_configuration` block

The `crawler_configuration` configuration block supports the following arguments:

* `exclusion_filters` - (Optional) List of one or more exclusion regular expression patterns to exclude certain object types that adhere to the pattern.
* `inclusion_filters` - (Optional) List of one or more inclusion regular expression patterns to include certain object types that adhere to the pattern.
* `scope` - (Optional) Scope of what is crawled for your URLs.
* `user_agent` - (Optional) String used for identifying the crawler or a bot when it accesses a web server. Default value is `bedrockbot_UUID`.
* `crawler_limits` - (Optional) Configuration of crawl limits for the web URLs. See [`crawler_limits` block](#crawler_limits-block) for details.

### `crawler_limits` block

The `crawler_limits` configuration block supports the following arguments:

* `max_pages` - (Optional) Max number of web pages crawled from your source URLs, up to 25,000 pages.
* `rate_limit` - (Optional) Max rate at which pages are crawled, up to 300 per minute per host.

### `server_side_encryption_configuration` block

The `server_side_encryption_configuration` configuration block supports the following arguments:

* `kms_key_arn` - (Optional) ARN of the AWS KMS key used to encrypt the resource.

### `vector_ingestion_configuration` block

The `vector_ingestion_configuration` configuration block supports the following arguments:

* `chunking_configuration` - (Optional, Forces new resource) Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. See [`chunking_configuration` block](#chunking_configuration-block) for details.
* `custom_transformation_configuration`- (Optional, Forces new resource) Configuration for custom transformation of data source documents.
* `parsing_configuration` - (Optional, Forces new resource) Configuration for custom parsing of data source documents. See [`parsing_configuration` block](#parsing_configuration-block) for details.

### `chunking_configuration` block

 The `chunking_configuration` configuration block supports the following arguments:

* `chunking_strategy` - (Required, Forces new resource) Option for chunking your source data, either in fixed-sized chunks or as one chunk. Valid values: `FIXED_SIZE`, `HIERARCHICAL`, `SEMANTIC`, `NONE`.
* `fixed_size_chunking_configuration` - (Optional, Forces new resource) Configurations for when you choose fixed-size chunking. Requires chunking_strategy as `FIXED_SIZE`. See [`fixed_size_chunking_configuration`](#fixed_size_chunking_configuration-block) for details.
* `hierarchical_chunking_configuration` - (Optional, Forces new resource) Configurations for when you choose hierarchical chunking. Requires chunking_strategy as `HIERARCHICAL`. See [`hierarchical_chunking_configuration`](#hierarchical_chunking_configuration-block) for details.
* `semantic_chunking_configuration` - (Optional, Forces new resource) Configurations for when you choose semantic chunking. Requires chunking_strategy as `SEMANTIC`. See [`semantic_chunking_configuration`](#semantic_chunking_configuration-block) for details.

### `fixed_size_chunking_configuration` block

The `fixed_size_chunking_configuration` block supports the following arguments:

* `max_tokens` - (Required, Forces new resource) Maximum number of tokens to include in a chunk.
* `overlap_percentage` - (Optional, Forces new resource) Percentage of overlap between adjacent chunks of a data source.

### `hierarchical_chunking_configuration` block

The `hierarchical_chunking_configuration` block supports the following arguments:

* `level_configuration` - (Required, Forces new resource) Maximum number of tokens to include in a chunk. Must contain two `level_configurations`. See [`level_configurations`](#level_configuration-block) for details.
* `overlap_tokens` - (Required, Forces new resource) The number of tokens to repeat across chunks in the same layer.

### `level_configuration` block

The `level_configuration` block supports the following arguments:

* `max_tokens` - (Required) The maximum number of tokens that a chunk can contain in this layer.

### `semantic_chunking_configuration` block

The `semantic_chunking_configuration` block supports the following arguments:

* `breakpoint_percentile_threshold` - (Required, Forces new resource) The dissimilarity threshold for splitting chunks.
* `buffer_size` - (Required, Forces new resource) The buffer size.
* `max_token` - (Required, Forces new resource) The maximum number of tokens a chunk can contain.

### `custom_transformation_configuration` block

The `custom_transformation_configuration` block supports the following arguments:

* `intermediate_storage` - (Required, Forces new resource) The intermediate storage for custom transformation.
* `transformation` - (Required) A custom processing step for documents moving through the data source ingestion pipeline.

### `intermediate_storage` block

The `intermediate_storage` block supports the following arguments:

* `s3_location` - (Required, Forces new resource) Configuration block for intermedia S3 storage.

### `s3_location` block

The `s3_location` block supports the following arguments:

* `uri` - (Required, Forces new resource) S3 URI for intermediate storage.

### `transformation` block

The `transformation` block supports the following arguments:

* `step_to_apply` - (Required, Forces new resource) When the service applies the transformation. Currently only `POST_CHUNKING` is supported.
* `transformation_function` - (Required) The lambda function that processes documents.

### `transformation_function` block

The `transformation_function` block supports the following arguments:

* `transformation_lambda_configuration` - (Required, Forces new resource) The configuration of the lambda function.

### `transformation_lambda_configuration` block

The `transformation_lambda_configuration` block supports the following arguments:

* `lambda_arn` - (Required, Forces new resource) The ARN of the lambda to use for custom transformation.

### `parsing_configuration` block

The `parsing_configuration` configuration block supports the following arguments:

* `parsing_strategy` - (Required) Currently only `BEDROCK_FOUNDATION_MODEL` is supported
* `bedrock_foundation_model_configuration` - (Optional) Settings for a foundation model used to parse documents in a data source. See [`bedrock_foundation_model_configuration` block](#bedrock_foundation_model_configuration-block) for details.

### `bedrock_foundation_model_configuration` block

The `bedrock_foundation_model_configuration` configuration block supports the following arguments:

* `model_arn` - (Required) The ARN of the model used to parse documents
* `parsing_prompt` - (Optional) Instructions for interpreting the contents of the document. See [`parsing_prompt` block](#parsing_prompt-block) for details.

### `parsing_prompt` block

The `parsing_prompt` configuration block supports the following arguments:

* `parsing_prompt_string` - (Required) Instructions for interpreting the contents of the document.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `data_source_id` -  Unique identifier of the data source.
* `id` -  Identifier of the data source which consists of the data source ID and the knowledge base ID.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Agents for Amazon Bedrock Data Source using the data source ID and the knowledge base ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.bedrockagent_data_source import BedrockagentDataSource
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BedrockagentDataSource.generate_config_for_import(self, "example", "GWCMFMQF6T,EMDPPAYPZI")
```

Using `terraform import`, import Agents for Amazon Bedrock Data Source using the data source ID and the knowledge base ID. For example:

```console
% terraform import aws_bedrockagent_data_source.example GWCMFMQF6T,EMDPPAYPZI
```

[1]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_ConfluenceDataSourceConfiguration.html
[2]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_SalesforceDataSourceConfiguration.html
[3]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_SharePointDataSourceConfiguration.html
[4]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_WebDataSourceConfiguration.html

<!-- cache-key: cdktf-0.20.8 input-6c5d65619f26f4664cb0988d2b1b6935663c952fb48ce89395b8908c00666e96 -->