---
subcategory: "Service Catalog"
layout: "aws"
page_title: "AWS: aws_servicecatalog_tag_option"
description: |-
  Manages a Service Catalog Tag Option
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_servicecatalog_tag_option

Manages a Service Catalog Tag Option.

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
import { ServicecatalogTagOption } from "./.gen/providers/aws/servicecatalog-tag-option";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ServicecatalogTagOption(this, "example", {
      key: "nyckel",
      value: "v\xE4rde",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `key` - (Required) Tag option key.
* `value` - (Required) Tag option value.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `active` - (Optional) Whether tag option is active. Default is `true`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Identifier (e.g., `tag-pjtvagohlyo3m`).
* `ownerId` - AWS account ID of the owner account that created the tag option.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `3m`)
- `read` - (Default `10m`)
- `update` - (Default `3m`)
- `delete` - (Default `3m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_servicecatalog_tag_option` using the tag option ID. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ServicecatalogTagOption } from "./.gen/providers/aws/servicecatalog-tag-option";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ServicecatalogTagOption.generateConfigForImport(
      this,
      "example",
      "tag-pjtvagohlyo3m"
    );
  }
}

```

Using `terraform import`, import `aws_servicecatalog_tag_option` using the tag option ID. For example:

```console
% terraform import aws_servicecatalog_tag_option.example tag-pjtvagohlyo3m
```

<!-- cache-key: cdktf-0.20.8 input-c0a2d3c7b00312457537db89b9cfb0cd5e52a148b85874b2a0846d6097f5d069 -->