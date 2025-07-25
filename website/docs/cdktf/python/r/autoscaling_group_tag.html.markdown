---
subcategory: "Auto Scaling"
layout: "aws"
page_title: "AWS: aws_autoscaling_group_tag"
description: |-
  Manages an individual Autoscaling Group tag
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_autoscaling_group_tag

Manages an individual Autoscaling Group (ASG) tag. This resource should only be used in cases where ASGs are created outside Terraform (e.g., ASGs implicitly created by EKS Node Groups).

~> **NOTE:** This tagging resource should not be combined with the Terraform resource for managing the parent resource. For example, using `aws_autoscaling_group` and `aws_autoscaling_group_tag` to manage tags of the same ASG will cause a perpetual difference where the `aws_autoscaling_group` resource will try to remove the tag being added by the `aws_autoscaling_group_tag` resource.

~> **NOTE:** This tagging resource does not use the [provider `ignore_tags` configuration](/docs/providers/aws/index.html#ignore_tags).

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformIterator, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.autoscaling_group_tag import AutoscalingGroupTagA
from imports.aws.eks_node_group import EksNodeGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, nodeRoleArn, scalingConfig, subnetIds):
        super().__init__(scope, name)
        # In most cases loops should be handled in the programming language context and
        #     not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
        #     you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
        #     you need to keep this like it is.
        example_for_each_iterator = TerraformIterator.from_list(
            Token.as_any(
                Fn.toset("${[ for asg in ${" +
                    Fn.flatten("${[ for resources in ${" + aws_eks_node_group_example.resources + "} : resources.autoscaling_groups]}") + "} : asg.name]}")))
        AutoscalingGroupTagA(self, "example",
            autoscaling_group_name=Token.as_string(example_for_each_iterator.value),
            tag=AutoscalingGroupTagTag(
                key="k8s.io/cluster-autoscaler/node-template/label/eks.amazonaws.com/capacityType",
                propagate_at_launch=False,
                value="SPOT"
            ),
            for_each=example_for_each_iterator
        )
        aws_eks_node_group_example = EksNodeGroup(self, "example_1",
            cluster_name="example",
            node_group_name="example",
            node_role_arn=node_role_arn,
            scaling_config=scaling_config,
            subnet_ids=subnet_ids
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_eks_node_group_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `autoscaling_group_name` - (Required) Name of the Autoscaling Group to apply the tag to.
* `tag` - (Required) Tag to create. The `tag` block is documented below.

The `tag` block supports the following arguments:

* `key` - (Required) Tag name.
* `value` - (Required) Tag value.
* `propagate_at_launch` - (Required) Whether to propagate the tags to instances launched by the ASG.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ASG name and key, separated by a comma (`,`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_autoscaling_group_tag` using the ASG name and key, separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.autoscaling_group_tag import AutoscalingGroupTagA
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AutoscalingGroupTagA.generate_config_for_import(self, "example", "asg-example,k8s.io/cluster-autoscaler/node-template/label/eks.amazonaws.com/capacityType")
```

Using `terraform import`, import `aws_autoscaling_group_tag` using the ASG name and key, separated by a comma (`,`). For example:

```console
% terraform import aws_autoscaling_group_tag.example asg-example,k8s.io/cluster-autoscaler/node-template/label/eks.amazonaws.com/capacityType
```

<!-- cache-key: cdktf-0.20.8 input-1c08bdd424cf14c34ad4e8bfd1e8a6cb13c02607e0f15f68371b532bee4057c7 -->