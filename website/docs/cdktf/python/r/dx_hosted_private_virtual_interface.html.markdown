---
subcategory: "Direct Connect"
layout: "aws"
page_title: "AWS: aws_dx_hosted_private_virtual_interface"
description: |-
  Provides a Direct Connect hosted private virtual interface resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dx_hosted_private_virtual_interface

Provides a Direct Connect hosted private virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
A hosted virtual interface is a virtual interface that is owned by another AWS account.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.dx_hosted_private_virtual_interface import DxHostedPrivateVirtualInterface
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, ownerAccountId):
        super().__init__(scope, name)
        DxHostedPrivateVirtualInterface(self, "foo",
            address_family="ipv4",
            bgp_asn=65352,
            connection_id="dxcon-zzzzzzzz",
            name="vif-foo",
            vlan=4094,
            owner_account_id=owner_account_id
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `address_family` - (Required) The address family for the BGP peer. `ipv4 ` or `ipv6`.
* `bgp_asn` - (Required) The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
* `connection_id` - (Required) The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
* `name` - (Required) The name for the virtual interface.
* `owner_account_id` - (Required) The AWS account that will own the new virtual interface.
* `vlan` - (Required) The VLAN ID.
* `amazon_address` - (Optional) The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
* `mtu` - (Optional) The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
* `bgp_auth_key` - (Optional) The authentication key for BGP configuration.
* `customer_address` - (Optional) The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the virtual interface.
* `arn` - The ARN of the virtual interface.
* `jumbo_frame_capable` - Indicates whether jumbo frames (9001 MTU) are supported.
* `aws_device` - The Direct Connect endpoint on which the virtual interface terminates.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `update` - (Default `10m`)
- `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Direct Connect hosted private virtual interfaces using the VIF `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.dx_hosted_private_virtual_interface import DxHostedPrivateVirtualInterface
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DxHostedPrivateVirtualInterface.generate_config_for_import(self, "test", "dxvif-33cc44dd")
```

Using `terraform import`, import Direct Connect hosted private virtual interfaces using the VIF `id`. For example:

```console
% terraform import aws_dx_hosted_private_virtual_interface.test dxvif-33cc44dd
```

<!-- cache-key: cdktf-0.20.8 input-507ab4001784e0faa9df00c7b461add9ebf38d5936fa6656c65e995d1c671fde -->