---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_core_network"
description: |-
  Manages a Network Manager Core Network.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_networkmanager_core_network

Manages a Network Manager Core Network.

Use this resource to create and manage a core network within a global network.

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
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkmanagerCoreNetwork(self, "example",
            global_network_id=Token.as_string(aws_networkmanager_global_network_example.id)
        )
```

### With description

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkmanagerCoreNetwork(self, "example",
            description="example",
            global_network_id=Token.as_string(aws_networkmanager_global_network_example.id)
        )
```

### With tags

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkmanagerCoreNetwork(self, "example",
            global_network_id=Token.as_string(aws_networkmanager_global_network_example.id),
            tags={
                "hello": "world"
            }
        )
```

### With VPC Attachment (Single Region)

The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `create_base_policy` argument to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `terraform apply` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `create_base_policy` argument. There are 2 options to implement this:

- Option 1: Use the `base_policy_document` argument that allows the most customizations to a base policy. Use this to customize the `edge_locations` `asn`. In the example below, `us-west-2` and ASN `65500` are used in the base policy.
- Option 2: Use the `create_base_policy` argument only. This creates a base policy in the region specified in the `provider` block.

#### Option 1 - using base_policy_document

If you require a custom ASN for the edge location, please use the `base_policy_document` argument to pass a specific ASN. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_networkmanager_core_network_policy_document import DataAwsNetworkmanagerCoreNetworkPolicyDocument
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
from imports.aws.networkmanager_core_network_policy_attachment import NetworkmanagerCoreNetworkPolicyAttachment
from imports.aws.networkmanager_global_network import NetworkmanagerGlobalNetwork
from imports.aws.networkmanager_vpc_attachment import NetworkmanagerVpcAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = NetworkmanagerGlobalNetwork(self, "example")
        base = DataAwsNetworkmanagerCoreNetworkPolicyDocument(self, "base",
            core_network_configuration=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfiguration(
                asn_ranges=["65022-65534"],
                edge_locations=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    asn="65500",
                    location="us-west-2"
                )
                ]
            )
            ],
            segments=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment"
            )
            ]
        )
        aws_networkmanager_core_network_example = NetworkmanagerCoreNetwork(self, "example_2",
            base_policy_document=Token.as_string(base.json),
            create_base_policy=True,
            global_network_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_example.override_logical_id("example")
        aws_networkmanager_vpc_attachment_example =
        NetworkmanagerVpcAttachment(self, "example_3",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            subnet_arns=Token.as_list(
                Fn.lookup_nested(aws_subnet_example, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example.arn)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_vpc_attachment_example.override_logical_id("example")
        data_aws_networkmanager_core_network_policy_document_example =
        DataAwsNetworkmanagerCoreNetworkPolicyDocument(self, "example_4",
            core_network_configuration=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfiguration(
                asn_ranges=["65022-65534"],
                edge_locations=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    asn="65500",
                    location="us-west-2"
                )
                ]
            )
            ],
            segment_actions=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActions(
                action="create-route",
                destination_cidr_blocks=["0.0.0.0/0"],
                destinations=[
                    Token.as_string(aws_networkmanager_vpc_attachment_example.id)
                ],
                segment="segment"
            )
            ],
            segments=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment"
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_networkmanager_core_network_policy_document_example.override_logical_id("example")
        aws_networkmanager_core_network_policy_attachment_example =
        NetworkmanagerCoreNetworkPolicyAttachment(self, "example_5",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            policy_document=Token.as_string(data_aws_networkmanager_core_network_policy_document_example.json)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_policy_attachment_example.override_logical_id("example")
```

#### Option 2 - create_base_policy only

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_networkmanager_core_network_policy_document import DataAwsNetworkmanagerCoreNetworkPolicyDocument
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
from imports.aws.networkmanager_core_network_policy_attachment import NetworkmanagerCoreNetworkPolicyAttachment
from imports.aws.networkmanager_global_network import NetworkmanagerGlobalNetwork
from imports.aws.networkmanager_vpc_attachment import NetworkmanagerVpcAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = NetworkmanagerGlobalNetwork(self, "example")
        aws_networkmanager_core_network_example = NetworkmanagerCoreNetwork(self, "example_1",
            create_base_policy=True,
            global_network_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_example.override_logical_id("example")
        aws_networkmanager_vpc_attachment_example =
        NetworkmanagerVpcAttachment(self, "example_2",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            subnet_arns=Token.as_list(
                Fn.lookup_nested(aws_subnet_example, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example.arn)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_vpc_attachment_example.override_logical_id("example")
        data_aws_networkmanager_core_network_policy_document_example =
        DataAwsNetworkmanagerCoreNetworkPolicyDocument(self, "example_3",
            core_network_configuration=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfiguration(
                asn_ranges=["65022-65534"],
                edge_locations=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    location="us-west-2"
                )
                ]
            )
            ],
            segment_actions=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActions(
                action="create-route",
                destination_cidr_blocks=["0.0.0.0/0"],
                destinations=[
                    Token.as_string(aws_networkmanager_vpc_attachment_example.id)
                ],
                segment="segment"
            )
            ],
            segments=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment"
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_networkmanager_core_network_policy_document_example.override_logical_id("example")
        aws_networkmanager_core_network_policy_attachment_example =
        NetworkmanagerCoreNetworkPolicyAttachment(self, "example_4",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            policy_document=Token.as_string(data_aws_networkmanager_core_network_policy_document_example.json)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_policy_attachment_example.override_logical_id("example")
```

### With VPC Attachment (Multi-Region)

The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `create_base_policy` argument of the [`aws_networkmanager_core_network` resource](/docs/providers/aws/r/networkmanager_core_network.html) to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `terraform apply` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `create_base_policy` argument. For multi-region in a core network that does not yet have a `LIVE` policy, there are 2 options:

- Option 1: Use the `base_policy_document` argument that allows the most customizations to a base policy. Use this to customize the `edge_locations` `asn`. In the example below, `us-west-2`, `us-east-1` and specific ASNs are used in the base policy.
- Option 2: Pass a list of regions to the `aws_networkmanager_core_network` `base_policy_regions` argument. In the example below, `us-west-2` and `us-east-1` are specified in the base policy.

#### Option 1 - using base_policy_document

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_networkmanager_core_network_policy_document import DataAwsNetworkmanagerCoreNetworkPolicyDocument
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
from imports.aws.networkmanager_core_network_policy_attachment import NetworkmanagerCoreNetworkPolicyAttachment
from imports.aws.networkmanager_global_network import NetworkmanagerGlobalNetwork
from imports.aws.networkmanager_vpc_attachment import NetworkmanagerVpcAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = NetworkmanagerGlobalNetwork(self, "example")
        base = DataAwsNetworkmanagerCoreNetworkPolicyDocument(self, "base",
            core_network_configuration=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfiguration(
                asn_ranges=["65022-65534"],
                edge_locations=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    asn="65500",
                    location="us-west-2"
                ), DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    asn="65501",
                    location="us-east-1"
                )
                ]
            )
            ],
            segments=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment"
            )
            ]
        )
        aws_networkmanager_core_network_example = NetworkmanagerCoreNetwork(self, "example_2",
            base_policy_document=Token.as_string(base.json),
            create_base_policy=True,
            global_network_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_example.override_logical_id("example")
        example_us_east1 = NetworkmanagerVpcAttachment(self, "example_us_east_1",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            provider="alternate",
            subnet_arns=Token.as_list(
                Fn.lookup_nested(aws_subnet_example_us_east1, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example_us_east1.arn)
        )
        example_us_west2 = NetworkmanagerVpcAttachment(self, "example_us_west_2",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            subnet_arns=Token.as_list(
                Fn.lookup_nested(aws_subnet_example_us_west2, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example_us_west2.arn)
        )
        data_aws_networkmanager_core_network_policy_document_example =
        DataAwsNetworkmanagerCoreNetworkPolicyDocument(self, "example_5",
            core_network_configuration=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfiguration(
                asn_ranges=["65022-65534"],
                edge_locations=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    asn="65500",
                    location="us-west-2"
                ), DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    asn="65501",
                    location="us-east-1"
                )
                ]
            )
            ],
            segment_actions=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActions(
                action="create-route",
                destination_cidr_blocks=["10.0.0.0/16"],
                destinations=[example_us_west2.id],
                segment="segment"
            ), DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActions(
                action="create-route",
                destination_cidr_blocks=["10.1.0.0/16"],
                destinations=[example_us_east1.id],
                segment="segment"
            )
            ],
            segments=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment"
            ), DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment2"
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_networkmanager_core_network_policy_document_example.override_logical_id("example")
        aws_networkmanager_core_network_policy_attachment_example =
        NetworkmanagerCoreNetworkPolicyAttachment(self, "example_6",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            policy_document=Token.as_string(data_aws_networkmanager_core_network_policy_document_example.json)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_policy_attachment_example.override_logical_id("example")
```

#### Option 2 - using base_policy_regions

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_networkmanager_core_network_policy_document import DataAwsNetworkmanagerCoreNetworkPolicyDocument
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
from imports.aws.networkmanager_core_network_policy_attachment import NetworkmanagerCoreNetworkPolicyAttachment
from imports.aws.networkmanager_global_network import NetworkmanagerGlobalNetwork
from imports.aws.networkmanager_vpc_attachment import NetworkmanagerVpcAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = NetworkmanagerGlobalNetwork(self, "example")
        aws_networkmanager_core_network_example = NetworkmanagerCoreNetwork(self, "example_1",
            base_policy_regions=["us-west-2", "us-east-1"],
            create_base_policy=True,
            global_network_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_example.override_logical_id("example")
        example_us_east1 = NetworkmanagerVpcAttachment(self, "example_us_east_1",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            provider="alternate",
            subnet_arns=Token.as_list(
                Fn.lookup_nested(aws_subnet_example_us_east1, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example_us_east1.arn)
        )
        example_us_west2 = NetworkmanagerVpcAttachment(self, "example_us_west_2",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            subnet_arns=Token.as_list(
                Fn.lookup_nested(aws_subnet_example_us_west2, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example_us_west2.arn)
        )
        data_aws_networkmanager_core_network_policy_document_example =
        DataAwsNetworkmanagerCoreNetworkPolicyDocument(self, "example_4",
            core_network_configuration=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfiguration(
                asn_ranges=["65022-65534"],
                edge_locations=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    location="us-west-2"
                ), DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations(
                    location="us-east-1"
                )
                ]
            )
            ],
            segment_actions=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActions(
                action="create-route",
                destination_cidr_blocks=["10.0.0.0/16"],
                destinations=[example_us_west2.id],
                segment="segment"
            ), DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActions(
                action="create-route",
                destination_cidr_blocks=["10.1.0.0/16"],
                destinations=[example_us_east1.id],
                segment="segment"
            )
            ],
            segments=[DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment"
            ), DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegments(
                name="segment2"
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_networkmanager_core_network_policy_document_example.override_logical_id("example")
        aws_networkmanager_core_network_policy_attachment_example =
        NetworkmanagerCoreNetworkPolicyAttachment(self, "example_5",
            core_network_id=Token.as_string(aws_networkmanager_core_network_example.id),
            policy_document=Token.as_string(data_aws_networkmanager_core_network_policy_document_example.json)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_core_network_policy_attachment_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `global_network_id` - (Required) ID of the global network that a core network will be a part of.

The following arguments are optional:

* `base_policy_document` - (Optional, conflicts with `base_policy_regions`) Sets the base policy document for the core network. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
* `base_policy_regions` - (Optional, conflicts with `base_policy_document`) List of regions to add to the base policy. The base policy created by setting the `create_base_policy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `base_policy_regions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
* `create_base_policy` - (Optional) Whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the [`aws_networkmanager_core_network_policy_attachment` resource](/docs/providers/aws/r/networkmanager_core_network_policy_attachment.html). This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Terraform snippet can be found above [for VPC Attachment in a single region](#with-vpc-attachment-single-region) and [for VPC Attachment multi-region](#with-vpc-attachment-multi-region). An example base policy is shown below. This base policy is overridden with the policy that you specify in the [`aws_networkmanager_core_network_policy_attachment` resource](/docs/providers/aws/r/networkmanager_core_network_policy_attachment.html).

```json
{
  "version": "2021.12",
  "core-network-configuration": {
    "asn-ranges": [
      "64512-65534"
    ],
    "vpn-ecmp-support": false,
    "edge-locations": [
      {
        "location": "us-east-1"
      }
    ]
  },
  "segments": [
    {
      "name": "segment",
      "description": "base-policy",
      "isolate-attachments": false,
      "require-attachment-acceptance": false
    }
  ]
}
```

* `description` - (Optional) Description of the Core Network.
* `tags` - (Optional) Key-value tags for the Core Network. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Core Network ARN.
* `created_at` - Timestamp when a core network was created.
* `edges` - One or more blocks detailing the edges within a core network. [Detailed below](#edges).
* `id` - Core Network ID.
* `segments` - One or more blocks detailing the segments within a core network. [Detailed below](#segments).
* `state` - Current state of a core network.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

### `edges`

The `edges` configuration block supports the following arguments:

* `asn` - ASN of a core network edge.
* `edge_location` - Region where a core network edge is located.
* `inside_cidr_blocks` - Inside IP addresses used for core network edges.

### `segments`

The `segments` configuration block supports the following arguments:

* `edge_locations` - Regions where the edges are located.
* `name` - Name of a core network segment.
* `shared_segments` - Shared segments of a core network.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `delete` - (Default `30m`)
* `update` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_networkmanager_core_network` using the core network ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_core_network import NetworkmanagerCoreNetwork
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkmanagerCoreNetwork.generate_config_for_import(self, "example", "core-network-0d47f6t230mz46dy4")
```

Using `terraform import`, import `aws_networkmanager_core_network` using the core network ID. For example:

```console
% terraform import aws_networkmanager_core_network.example core-network-0d47f6t230mz46dy4
```

<!-- cache-key: cdktf-0.20.8 input-a8f63cb5ad0be1f33999c52038490ce409174d1bad497bbe3c33f417d8bd250a -->