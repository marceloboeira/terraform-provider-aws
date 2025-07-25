---
subcategory: "Service Quotas"
layout: "aws"
page_title: "AWS: aws_servicequotas_template"
description: |-
  Terraform resource for managing an AWS Service Quotas Template.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_servicequotas_template

Terraform resource for managing an AWS Service Quotas Template.

-> Only the management account of an organization can alter Service Quota templates, and this must be done from the `us-east-1` region.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ServicequotasTemplate } from "./.gen/providers/aws/servicequotas-template";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ServicequotasTemplate(this, "example", {
      awsRegion: "us-east-1",
      quotaCode: "L-2ACBD22F",
      serviceCode: "lambda",
      value: Token.asNumber("80"),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `awsRegion` - (Optional) AWS Region to which the template applies.
* `region` - (Optional, **Deprecated**) AWS Region to which the template applies. Use `awsRegion` instead.
* `quotaCode` - (Required) Quota identifier. To find the quota code for a specific quota, use the [aws_servicequotas_service_quota](../d/servicequotas_service_quota.html.markdown) data source.
* `serviceCode` - (Required) Service identifier. To find the service code value for an AWS service, use the [aws_servicequotas_service](../d/servicequotas_service.html.markdown) data source.
* `value` - (Required) The new, increased value for the quota.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `globalQuota` - Indicates whether the quota is global.
* `id` - Unique identifier for the resource, which is a comma-delimited string separating `region`, `quotaCode`, and `serviceCode`.
* `quotaName` - Quota name.
* `serviceName` - Service name.
* `unit` - Unit of measurement.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Service Quotas Template using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ServicequotasTemplate } from "./.gen/providers/aws/servicequotas-template";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ServicequotasTemplate.generateConfigForImport(
      this,
      "example",
      "us-east-1,L-2ACBD22F,lambda"
    );
  }
}

```

Using `terraform import`, import Service Quotas Template using the `id`. For example:

```console
% terraform import aws_servicequotas_template.example us-east-1,L-2ACBD22F,lambda
```

<!-- cache-key: cdktf-0.20.8 input-9fb04c46d57f8b66759ab5707ec68381d3c22dffcd39d2f7f2655bebb617bcbb -->