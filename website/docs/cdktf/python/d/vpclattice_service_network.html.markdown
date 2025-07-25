---
subcategory: "VPC Lattice"
layout: "aws"
page_title: "AWS: aws_vpclattice_service_network"
description: |-
  Terraform data source for managing an AWS VPC Lattice Service Network.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_vpclattice_service_network

Terraform data source for managing an AWS VPC Lattice Service Network.

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
from imports.aws.data_aws_vpclattice_service_network import DataAwsVpclatticeServiceNetwork
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsVpclatticeServiceNetwork(self, "example",
            service_network_identifier="snsa-01112223334445556"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `service_network_identifier` - (Required) Identifier of the service network.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Service Network.
* `auth_type` - Authentication type for the service network. Either `NONE` or `AWS_IAM`.
* `created_at` - Date and time the service network was created.
* `id` - ID of the service network.
* `last_updated_at` - Date and time the service network was last updated.
* `name` - Name of the service network.
* `number_of_associated_services` - Number of services associated with this service network.
* `number_of_associated_vpcs` - Number of VPCs associated with this service network.

<!-- cache-key: cdktf-0.20.8 input-2c51ddaae628d2b7c8c6f40ed62775719fe1cd2f4eb9e1f1b9fca02f7c958b10 -->