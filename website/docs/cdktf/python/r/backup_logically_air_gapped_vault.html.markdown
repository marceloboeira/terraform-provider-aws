---
subcategory: "Backup"
layout: "aws"
page_title: "AWS: aws_backup_logically_air_gapped_vault"
description: |-
  Terraform resource for managing an AWS Backup Logically Air Gapped Vault.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_backup_logically_air_gapped_vault

Terraform resource for managing an AWS Backup Logically Air Gapped Vault.

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
from imports.aws.backup_logically_air_gapped_vault import BackupLogicallyAirGappedVault
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BackupLogicallyAirGappedVault(self, "example",
            max_retention_days=7,
            min_retention_days=7,
            name="lag-example-vault"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the Logically Air Gapped Backup Vault to create.
* `max_retention_days` - (Required) Maximum retention period that the Logically Air Gapped Backup Vault retains recovery points.
* `min_retention_days` - (Required) Minimum retention period that the Logically Air Gapped Backup Vault retains recovery points.
* `tags` - (Optional) Metadata that you can assign to help organize the resources that you create. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The name of the Logically Air Gapped Backup Vault.
* `arn` - The ARN of the Logically Air Gapped Backup Vault.
* `recovery_points` - The number of recovery points that are stored in a Logically Air Gapped Backup Vault.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Backup Logically Air Gapped Vault using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.backup_logically_air_gapped_vault import BackupLogicallyAirGappedVault
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BackupLogicallyAirGappedVault.generate_config_for_import(self, "example", "lag-example-vault")
```

Using `terraform import`, import Backup Logically Air Gapped Vault using the `id`. For example:

```console
% terraform import aws_backup_logically_air_gapped_vault.example lag-example-vault
```

<!-- cache-key: cdktf-0.20.8 input-0ff68cd87217b0d55ec07fc6702526d5d80d3687ee2656847363c5409101b77a -->