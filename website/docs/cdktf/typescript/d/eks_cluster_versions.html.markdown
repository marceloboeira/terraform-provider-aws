---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "AWS: aws_eks_cluster_versions"
description: |-
  Terraform data source for managing AWS EKS (Elastic Kubernetes) Cluster Versions.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_eks_cluster_versions

Terraform data source for managing AWS EKS (Elastic Kubernetes) Cluster Versions.

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
import { DataAwsEksClusterVersions } from "./.gen/providers/aws/data-aws-eks-cluster-versions";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsEksClusterVersions(this, "example", {});
  }
}

```

### Filter by Cluster Type

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEksClusterVersions } from "./.gen/providers/aws/data-aws-eks-cluster-versions";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsEksClusterVersions(this, "example", {
      clusterType: "eks",
    });
  }
}

```

### Filter by Version Status

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEksClusterVersions } from "./.gen/providers/aws/data-aws-eks-cluster-versions";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsEksClusterVersions(this, "example", {
      versionStatus: "STANDARD_SUPPORT",
    });
  }
}

```

## Argument Reference

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `clusterType` - (Optional) Type of clusters to filter by.
Currently, the only valid value is `eks`.
* `clusterVersions` - (Optional) A list of Kubernetes versions that you can use to check if EKS supports it.
* `defaultOnly` - (Optional) Whether to show only the default versions of Kubernetes supported by EKS.
* `includeAll` - (Optional) Whether to include all kubernetes versions in the response.
* `versionStatus` - (Optional) Status of the EKS cluster versions to list.
Valid values are `STANDARD_SUPPORT` or `UNSUPPORTED` or `EXTENDED_SUPPORT`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `clusterType` - Type of cluster that the version belongs to.
* `clusterVersion` - Kubernetes version supported by EKS.
* `default_platform_version` - Default eks platform version for the cluster version.
* `defaultVersion` - Default Kubernetes version for the cluster version.
* `end_of_extended_support_date` - End of extended support date for the cluster version.
* `end_of_standard_support_date` - End of standard support date for the cluster version.
* `kubernetes_patch_version` - Kubernetes patch version for the cluster version.
* `releaseDate` - Release date of the cluster version.
* `versionStatus` - Status of the EKS cluster version.

<!-- cache-key: cdktf-0.20.8 input-bab7956d183fc4501c59ae29ab510d656d752830ccc7c38233b7715b81d555e6 -->