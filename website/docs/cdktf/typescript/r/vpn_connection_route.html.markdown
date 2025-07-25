---
subcategory: "VPN (Site-to-Site)"
layout: "aws"
page_title: "AWS: aws_vpn_connection_route"
description: |-
  Provides a static route between a VPN connection and a customer gateway.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpn_connection_route

Provides a static route between a VPN connection and a customer gateway.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CustomerGateway } from "./.gen/providers/aws/customer-gateway";
import { Vpc } from "./.gen/providers/aws/vpc";
import { VpnConnection } from "./.gen/providers/aws/vpn-connection";
import { VpnConnectionRoute } from "./.gen/providers/aws/vpn-connection-route";
import { VpnGateway } from "./.gen/providers/aws/vpn-gateway";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const customerGateway = new CustomerGateway(this, "customer_gateway", {
      bgpAsn: Token.asString(65000),
      ipAddress: "172.0.0.1",
      type: "ipsec.1",
    });
    const vpc = new Vpc(this, "vpc", {
      cidrBlock: "10.0.0.0/16",
    });
    const vpnGateway = new VpnGateway(this, "vpn_gateway", {
      vpcId: vpc.id,
    });
    const main = new VpnConnection(this, "main", {
      customerGatewayId: customerGateway.id,
      staticRoutesOnly: true,
      type: "ipsec.1",
      vpnGatewayId: vpnGateway.id,
    });
    new VpnConnectionRoute(this, "office", {
      destinationCidrBlock: "192.168.10.0/24",
      vpnConnectionId: main.id,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `destinationCidrBlock` - (Required) The CIDR block associated with the local subnet of the customer network.
* `vpnConnectionId` - (Required) The ID of the VPN connection.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `destinationCidrBlock` - The CIDR block associated with the local subnet of the customer network.
* `vpnConnectionId` - The ID of the VPN connection.

<!-- cache-key: cdktf-0.20.8 input-8f8d512d7c4f5cec27ea7153220a06d9aac35bc3fba6d1e950beacddc5f62a6f -->