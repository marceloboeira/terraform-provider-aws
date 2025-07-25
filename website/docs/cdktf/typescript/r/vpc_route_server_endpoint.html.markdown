---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_route_server_endpoint"
description: |-
  Terraform resource for managing a VPC (Virtual Private Cloud) Route Server.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_route_server_endpoint

  Provides a resource for managing a VPC (Virtual Private Cloud) Route Server Endpoint.

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
import { VpcRouteServerEndpoint } from "./.gen/providers/aws/vpc-route-server-endpoint";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new VpcRouteServerEndpoint(this, "test", {
      routeServerId: example.routeServerId,
      subnetId: main.id,
      tags: {
        Name: "Endpoint A",
      },
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `routeServerId` - (Required) The ID of the route server for which to create an endpoint.
* `subnetId` - (Required) The ID of the subnet in which to create the route server endpoint.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the route server endpoint.
* `routeServerEndpointId` - The unique identifier of the route server endpoint.
* `eniId` - The ID of the Elastic network interface for the endpoint.
* `eniAddress` - The IP address of the Elastic network interface for the endpoint.
* `vpcId` - The ID of the VPC containing the endpoint.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import VPC (Virtual Private Cloud) Route Server Endpoint using the `routeServerEndpointId`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcRouteServerEndpoint } from "./.gen/providers/aws/vpc-route-server-endpoint";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcRouteServerEndpoint.generateConfigForImport(
      this,
      "example",
      "rse-12345678"
    );
  }
}

```

Using `terraform import`, import VPC (Virtual Private Cloud) Route Server Endpoint using the `routeServerEndpointId`. For example:

```console
% terraform import aws_vpc_route_server_endpoint.example rse-12345678
```

<!-- cache-key: cdktf-0.20.8 input-0d1fcc2de0cf1f168efcff6502c1aa03bc59922affb8982ca926708a5dd15fd4 -->