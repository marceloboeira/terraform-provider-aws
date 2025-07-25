---
subcategory: "VPN (Client)"
layout: "aws"
page_title: "AWS: aws_ec2_client_vpn_network_association"
description: |-
  Provides network associations for AWS Client VPN endpoints.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ec2_client_vpn_network_association

Provides network associations for AWS Client VPN endpoints. For more information on usage, please see the
[AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Ec2ClientVpnNetworkAssociation } from "./.gen/providers/aws/ec2-client-vpn-network-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Ec2ClientVpnNetworkAssociation(this, "example", {
      clientVpnEndpointId: Token.asString(awsEc2ClientVpnEndpointExample.id),
      subnetId: Token.asString(awsSubnetExample.id),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `clientVpnEndpointId` - (Required) The ID of the Client VPN endpoint.
* `subnetId` - (Required) The ID of the subnet to associate with the Client VPN endpoint.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The unique ID of the target network association.
* `associationId` - The unique ID of the target network association.
* `vpcId` - The ID of the VPC in which the target subnet is located.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `30m`)
- `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AWS Client VPN network associations using the endpoint ID and the association ID. Values are separated by a `,`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Ec2ClientVpnNetworkAssociation } from "./.gen/providers/aws/ec2-client-vpn-network-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Ec2ClientVpnNetworkAssociation.generateConfigForImport(
      this,
      "example",
      "cvpn-endpoint-0ac3a1abbccddd666,cvpn-assoc-0b8db902465d069ad"
    );
  }
}

```

Using `terraform import`, import AWS Client VPN network associations using the endpoint ID and the association ID. Values are separated by a `,`. For example:

```console
% terraform import aws_ec2_client_vpn_network_association.example cvpn-endpoint-0ac3a1abbccddd666,cvpn-assoc-0b8db902465d069ad
```

<!-- cache-key: cdktf-0.20.8 input-100690b48718cc3344c93c80052a9aa81188ecae560d356b4144aeb6b5ca4788 -->