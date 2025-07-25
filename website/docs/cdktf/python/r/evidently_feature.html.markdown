---
subcategory: "CloudWatch Evidently"
layout: "aws"
page_title: "AWS: aws_evidently_feature"
description: |-
  Provides a CloudWatch Evidently Feature resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_evidently_feature

Provides a CloudWatch Evidently Feature resource.

~> **Warning:** This resource is deprecated. Use [AWS AppConfig feature flags](https://aws.amazon.com/blogs/mt/using-aws-appconfig-feature-flags/) instead.

## Example Usage

### Basic

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.evidently_feature import EvidentlyFeature
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EvidentlyFeature(self, "example",
            description="example description",
            name="example",
            project=Token.as_string(aws_evidently_project_example.name),
            tags={
                "Key1": "example Feature"
            },
            variations=[EvidentlyFeatureVariations(
                name="Variation1",
                value=EvidentlyFeatureVariationsValue(
                    string_value="example"
                )
            )
            ]
        )
```

### With default variation

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.evidently_feature import EvidentlyFeature
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EvidentlyFeature(self, "example",
            default_variation="Variation2",
            name="example",
            project=Token.as_string(aws_evidently_project_example.name),
            variations=[EvidentlyFeatureVariations(
                name="Variation1",
                value=EvidentlyFeatureVariationsValue(
                    string_value="exampleval1"
                )
            ), EvidentlyFeatureVariations(
                name="Variation2",
                value=EvidentlyFeatureVariationsValue(
                    string_value="exampleval2"
                )
            )
            ]
        )
```

### With entity overrides

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.evidently_feature import EvidentlyFeature
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EvidentlyFeature(self, "example",
            entity_overrides={
                "test1": "Variation1"
            },
            name="example",
            project=Token.as_string(aws_evidently_project_example.name),
            variations=[EvidentlyFeatureVariations(
                name="Variation1",
                value=EvidentlyFeatureVariationsValue(
                    string_value="exampleval1"
                )
            ), EvidentlyFeatureVariations(
                name="Variation2",
                value=EvidentlyFeatureVariationsValue(
                    string_value="exampleval2"
                )
            )
            ]
        )
```

### With evaluation strategy

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.evidently_feature import EvidentlyFeature
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EvidentlyFeature(self, "example",
            entity_overrides={
                "test1": "Variation1"
            },
            evaluation_strategy="ALL_RULES",
            name="example",
            project=Token.as_string(aws_evidently_project_example.name),
            variations=[EvidentlyFeatureVariations(
                name="Variation1",
                value=EvidentlyFeatureVariationsValue(
                    string_value="exampleval1"
                )
            )
            ]
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `default_variation` - (Optional) The name of the variation to use as the default variation. The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature. This variation must also be listed in the `variations` structure. If you omit `default_variation`, the first variation listed in the `variations` structure is used as the default variation.
* `description` - (Optional) Specifies the description of the feature.
* `entity_overrides` - (Optional) Specify users that should always be served a specific variation of a feature. Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
* `evaluation_strategy` - (Optional) Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments. Specify `DEFAULT_VARIATION` to serve the default variation to all users instead.
* `name` - (Required) The name for the new feature. Minimum length of `1`. Maximum length of `127`.
* `project` - (Required) The name or ARN of the project that is to contain the new feature.
* `tags` - (Optional) Tags to apply to the feature. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `variations` - (Required) One or more blocks that contain the configuration of the feature's different variations. [Detailed below](#variations)

### `variations`

The `variations` block supports the following arguments:

* `name` - (Required) The name of the variation. Minimum length of `1`. Maximum length of `127`.
* `value` - (Required) A block that specifies the value assigned to this variation. [Detailed below](#value)

#### `value`

The `value` block supports the following arguments:

~> **NOTE:** You must specify exactly one of `bool_value`, `double_value`, `long_value`, `string_value`.

* `bool_value` - (Optional) If this feature uses the Boolean variation type, this field contains the Boolean value of this variation.
* `double_value` - (Optional) If this feature uses the double integer variation type, this field contains the double integer value of this variation.
* `long_value` - (Optional) If this feature uses the long variation type, this field contains the long value of this variation. Minimum value of `-9007199254740991`. Maximum value of `9007199254740991`.
* `string_value` - (Optional) If this feature uses the string variation type, this field contains the string value of this variation. Minimum length of `0`. Maximum length of `512`.

## Timeouts

[Configuration options](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts):

* `create` - (Default `2m`)
* `delete` - (Default `2m`)
* `update` - (Default `2m`)

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the feature.
* `created_time` - The date and time that the feature is created.
* `evaluation_rules` - One or more blocks that define the evaluation rules for the feature. [Detailed below](#evaluation_rules)
* `id` - The feature `name` and the project `name` or `arn` separated by a colon (`:`).
* `last_updated_time` - The date and time that the feature was most recently updated.
* `status` - The current state of the feature. Valid values are `AVAILABLE` and `UPDATING`.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).
* `value_type` - Defines the type of value used to define the different feature variations. Valid Values: `STRING`, `LONG`, `DOUBLE`, `BOOLEAN`.

### `evaluation_rules`

The `evaluation_rules` block supports the following attributes:

* `name` - The name of the experiment or launch.
* `type` - This value is `aws.evidently.splits` if this is an evaluation rule for a launch, and it is `aws.evidently.onlineab` if this is an evaluation rule for an experiment.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CloudWatch Evidently Feature using the feature `name` and `name` or `arn` of the hosting CloudWatch Evidently Project separated by a `:`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.evidently_feature import EvidentlyFeature
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EvidentlyFeature.generate_config_for_import(self, "example", "exampleFeatureName:arn:aws:evidently:us-east-1:123456789012:project/example")
```

Using `terraform import`, import CloudWatch Evidently Feature using the feature `name` and `name` or `arn` of the hosting CloudWatch Evidently Project separated by a `:`. For example:

```console
% terraform import aws_evidently_feature.example exampleFeatureName:arn:aws:evidently:us-east-1:123456789012:project/example
```

<!-- cache-key: cdktf-0.20.8 input-1504568c14ad2b3aeed9a47109d8576aa4fb7313f08add446227243aa56e4c4b -->