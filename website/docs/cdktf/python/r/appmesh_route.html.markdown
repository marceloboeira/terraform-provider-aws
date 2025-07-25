---
subcategory: "App Mesh"
layout: "aws"
page_title: "AWS: aws_appmesh_route"
description: |-
  Provides an AWS App Mesh route resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_appmesh_route

Provides an AWS App Mesh route resource.

## Example Usage

### HTTP Routing

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_route import AppmeshRoute
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshRoute(self, "serviceb",
            mesh_name=simple.id,
            name="serviceB-route",
            spec=AppmeshRouteSpec(
                http_route=AppmeshRouteSpecHttpRoute(
                    action=AppmeshRouteSpecHttpRouteAction(
                        weighted_target=[AppmeshRouteSpecHttpRouteActionWeightedTarget(
                            virtual_node=serviceb1.name,
                            weight=90
                        ), AppmeshRouteSpecHttpRouteActionWeightedTarget(
                            virtual_node=serviceb2.name,
                            weight=10
                        )
                        ]
                    ),
                    match=AppmeshRouteSpecHttpRouteMatch(
                        prefix="/"
                    )
                )
            ),
            virtual_router_name=Token.as_string(aws_appmesh_virtual_router_serviceb.name)
        )
```

### HTTP Header Routing

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_route import AppmeshRoute
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshRoute(self, "serviceb",
            mesh_name=simple.id,
            name="serviceB-route",
            spec=AppmeshRouteSpec(
                http_route=AppmeshRouteSpecHttpRoute(
                    action=AppmeshRouteSpecHttpRouteAction(
                        weighted_target=[AppmeshRouteSpecHttpRouteActionWeightedTarget(
                            virtual_node=Token.as_string(aws_appmesh_virtual_node_serviceb.name),
                            weight=100
                        )
                        ]
                    ),
                    match=AppmeshRouteSpecHttpRouteMatch(
                        header=[AppmeshRouteSpecHttpRouteMatchHeader(
                            match=AppmeshRouteSpecHttpRouteMatchHeaderMatch(
                                prefix="123"
                            ),
                            name="clientRequestId"
                        )
                        ],
                        method="POST",
                        prefix="/",
                        scheme="https"
                    )
                )
            ),
            virtual_router_name=Token.as_string(aws_appmesh_virtual_router_serviceb.name)
        )
```

### Retry Policy

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_route import AppmeshRoute
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshRoute(self, "serviceb",
            mesh_name=simple.id,
            name="serviceB-route",
            spec=AppmeshRouteSpec(
                http_route=AppmeshRouteSpecHttpRoute(
                    action=AppmeshRouteSpecHttpRouteAction(
                        weighted_target=[AppmeshRouteSpecHttpRouteActionWeightedTarget(
                            virtual_node=Token.as_string(aws_appmesh_virtual_node_serviceb.name),
                            weight=100
                        )
                        ]
                    ),
                    match=AppmeshRouteSpecHttpRouteMatch(
                        prefix="/"
                    ),
                    retry_policy=AppmeshRouteSpecHttpRouteRetryPolicy(
                        http_retry_events=["server-error"],
                        max_retries=1,
                        per_retry_timeout=AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout(
                            unit="s",
                            value=15
                        )
                    )
                )
            ),
            virtual_router_name=Token.as_string(aws_appmesh_virtual_router_serviceb.name)
        )
```

### TCP Routing

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_route import AppmeshRoute
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshRoute(self, "serviceb",
            mesh_name=simple.id,
            name="serviceB-route",
            spec=AppmeshRouteSpec(
                tcp_route=AppmeshRouteSpecTcpRoute(
                    action=AppmeshRouteSpecTcpRouteAction(
                        weighted_target=[AppmeshRouteSpecTcpRouteActionWeightedTarget(
                            virtual_node=serviceb1.name,
                            weight=100
                        )
                        ]
                    )
                )
            ),
            virtual_router_name=Token.as_string(aws_appmesh_virtual_router_serviceb.name)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name to use for the route. Must be between 1 and 255 characters in length.
* `mesh_name` - (Required) Name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
* `mesh_owner` - (Optional) AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider][1] is currently connected to.
* `virtual_router_name` - (Required) Name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
* `spec` - (Required) Route specification to apply.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

The `spec` object supports the following:

* `grpc_route` - (Optional) GRPC routing information for the route.
* `http2_route` - (Optional) HTTP/2 routing information for the route.
* `http_route` - (Optional) HTTP routing information for the route.
* `priority` - (Optional) Priority for the route, between `0` and `1000`.
Routes are matched based on the specified value, where `0` is the highest priority.
* `tcp_route` - (Optional) TCP routing information for the route.

The `grpc_route` object supports the following:

* `action` - (Required) Action to take if a match is determined.
* `match` - (Required) Criteria for determining an gRPC request match.
* `retry_policy` - (Optional) Retry policy.
* `timeout` - (Optional) Types of timeouts.

The `http2_route` and `http_route` objects supports the following:

* `action` - (Required) Action to take if a match is determined.
* `match` - (Required) Criteria for determining an HTTP request match.
* `retry_policy` - (Optional) Retry policy.
* `timeout` - (Optional) Types of timeouts.

The `tcp_route` object supports the following:

* `action` - (Required) Action to take if a match is determined.
* `timeout` - (Optional) Types of timeouts.

The `action` object supports the following:

* `weighted_target` - (Required) Targets that traffic is routed to when a request matches the route.
You can specify one or more targets and their relative weights with which to distribute traffic.

The `timeout` object supports the following:

* `idle` - (Optional) Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.

The `idle` object supports the following:

* `unit` - (Required) Unit of time. Valid values: `ms`, `s`.
* `value` - (Required) Number of time units. Minimum value of `0`.

The `grpc_route`'s `match` object supports the following:

* `metadata` - (Optional) Data to match from the gRPC request.
* `method_name` - (Optional) Method name to match from the request. If you specify a name, you must also specify a `service_name`.
* `service_name` - (Optional) Fully qualified domain name for the service to match from the request.
* `port`- (Optional) The port number to match from the request.

The `metadata` object supports the following:

* `name` - (Required) Name of the route. Must be between 1 and 50 characters in length.
* `invert` - (Optional) If `true`, the match is on the opposite of the `match` criteria. Default is `false`.
* `match` - (Optional) Data to match from the request.

The `metadata`'s `match` object supports the following:

* `exact` - (Optional) Value sent by the client must match the specified value exactly. Must be between 1 and 255 characters in length.
* `prefix` - (Optional) Value sent by the client must begin with the specified characters. Must be between 1 and 255 characters in length.
* `port`- (Optional) The port number to match from the request.
* `range`- (Optional) Object that specifies the range of numbers that the value sent by the client must be included in.
* `regex` - (Optional) Value sent by the client must include the specified characters. Must be between 1 and 255 characters in length.
* `suffix` - (Optional) Value sent by the client must end with the specified characters. Must be between 1 and 255 characters in length.

The `grpc_route`'s `retry_policy` object supports the following:

* `grpc_retry_events` - (Optional) List of gRPC retry events.
Valid values: `cancelled`, `deadline-exceeded`, `internal`, `resource-exhausted`, `unavailable`.
* `http_retry_events` - (Optional) List of HTTP retry events.
Valid values: `client-error` (HTTP status code 409), `gateway-error` (HTTP status codes 502, 503, and 504), `server-error` (HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511), `stream-error` (retry on refused stream).
* `max_retries` - (Required) Maximum number of retries.
* `per_retry_timeout` - (Required) Per-retry timeout.
* `tcp_retry_events` - (Optional) List of TCP retry events. The only valid value is `connection-error`.

The `grpc_route`'s `timeout` object supports the following:

* `idle` - (Optional) Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
* `per_request` - (Optional) Per request timeout.

The `idle` and `per_request` objects support the following:

* `unit` - (Required) Unit of time. Valid values: `ms`, `s`.
* `value` - (Required) Number of time units. Minimum value of `0`.

The `http2_route` and `http_route`'s `match` object supports the following:

* `prefix` - (Optional) Path with which to match requests.
This parameter must always start with /, which by itself matches all requests to the virtual router service name.
* `port`- (Optional) The port number to match from the request.
* `header` - (Optional) Client request headers to match on.
* `method` - (Optional) Client request header method to match on. Valid values: `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`.
* `path` - (Optional) Client request path to match on.
* `query_parameter` - (Optional) Client request query parameters to match on.
* `scheme` - (Optional) Client request header scheme to match on. Valid values: `http`, `https`.

The `match`'s `path` object supports the following:

* `exact` - (Optional) The exact path to match on.
* `regex` - (Optional) The regex used to match the path.

The `match`'s `query_parameter` object supports the following:

* `name` - (Required) Name for the query parameter that will be matched on.
* `match` - (Optional) The query parameter to match on.

The `query_parameter`'s `match` object supports the following:

* `exact` - (Optional) The exact query parameter to match on.

The `http2_route` and `http_route`'s `retry_policy` object supports the following:

* `http_retry_events` - (Optional) List of HTTP retry events.
Valid values: `client-error` (HTTP status code 409), `gateway-error` (HTTP status codes 502, 503, and 504), `server-error` (HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511), `stream-error` (retry on refused stream).
* `max_retries` - (Required) Maximum number of retries.
* `per_retry_timeout` - (Required) Per-retry timeout.
* `tcp_retry_events` - (Optional) List of TCP retry events. The only valid value is `connection-error`.

You must specify at least one value for `http_retry_events`, or at least one value for `tcp_retry_events`.

The `http2_route` and `http_route`'s `timeout` object supports the following:

* `idle` - (Optional) Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
* `per_request` - (Optional) Per request timeout.

The `idle` and `per_request` objects support the following:

* `unit` - (Required) Unit of time. Valid values: `ms`, `s`.
* `value` - (Required) Number of time units. Minimum value of `0`.

The `per_retry_timeout` object supports the following:

* `unit` - (Required) Retry unit. Valid values: `ms`, `s`.
* `value` - (Required) Retry value.

The `weighted_target` object supports the following:

* `virtual_node` - (Required) Virtual node to associate with the weighted target. Must be between 1 and 255 characters in length.
* `weight` - (Required) Relative weight of the weighted target. An integer between 0 and 100.
* `port` - (Optional) The targeted port of the weighted object.

The `header` object supports the following:

* `name` - (Required) Name for the HTTP header in the client request that will be matched on.
* `invert` - (Optional) If `true`, the match is on the opposite of the `match` method and value. Default is `false`.
* `match` - (Optional) Method and value to match the header value sent with a request. Specify one match method.

The `header`'s `match` object supports the following:

* `exact` - (Optional) Header value sent by the client must match the specified value exactly.
* `prefix` - (Optional) Header value sent by the client must begin with the specified characters.
* `port`- (Optional) The port number to match from the request.
* `range`- (Optional) Object that specifies the range of numbers that the header value sent by the client must be included in.
* `regex` - (Optional) Header value sent by the client must include the specified characters.
* `suffix` - (Optional) Header value sent by the client must end with the specified characters.

The `range` object supports the following:

* `end` - (Required) End of the range.
* `start` - (Requited) Start of the range.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the route.
* `arn` - ARN of the route.
* `created_date` - Creation date of the route.
* `last_updated_date` - Last update date of the route.
* `resource_owner` - Resource owner's AWS account ID.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import App Mesh virtual routes using `mesh_name` and `virtual_router_name` together with the route's `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_route import AppmeshRoute
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshRoute.generate_config_for_import(self, "serviceb", "simpleapp/serviceB/serviceB-route")
```

Using `terraform import`, import App Mesh virtual routes using `mesh_name` and `virtual_router_name` together with the route's `name`. For example:

```console
% terraform import aws_appmesh_route.serviceb simpleapp/serviceB/serviceB-route
```

[1]: /docs/providers/aws/index.html

<!-- cache-key: cdktf-0.20.8 input-aae20c8ef6acd6b71b7e2f1b54df83a088e5125374e89be2d56aa3a6c58297f7 -->