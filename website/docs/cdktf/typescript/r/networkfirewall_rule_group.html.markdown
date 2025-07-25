---
subcategory: "Network Firewall"
layout: "aws"
page_title: "AWS: aws_networkfirewall_rule_group"
description: |-
  Provides an AWS Network Firewall Rule Group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_networkfirewall_rule_group

Provides an AWS Network Firewall Rule Group Resource

## Example Usage

### Stateful Inspection for denying access to a domain

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkfirewallRuleGroup(this, "example", {
      capacity: 100,
      name: "example",
      ruleGroup: {
        rulesSource: {
          rulesSourceList: {
            generatedRulesType: "DENYLIST",
            targetTypes: ["HTTP_HOST"],
            targets: ["test.example.com"],
          },
        },
      },
      tags: {
        Tag1: "Value1",
        Tag2: "Value2",
      },
      type: "STATEFUL",
    });
  }
}

```

### Stateful Inspection for permitting packets from a source IP address

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformIterator, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const ips = ["1.1.1.1/32", "1.0.0.1/32"];
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const exampleDynamicIterator0 = TerraformIterator.fromList(
      Token.asAny(ips)
    );
    new NetworkfirewallRuleGroup(this, "example", {
      capacity: 50,
      description: "Permits http traffic from source",
      name: "example",
      ruleGroup: {
        rulesSource: {
          statefulRule: exampleDynamicIterator0.dynamic({
            action: "PASS",
            header: [
              {
                destination: "ANY",
                destination_port: "ANY",
                direction: "ANY",
                protocol: "HTTP",
                source: exampleDynamicIterator0.value,
                source_port: "ANY",
              },
            ],
            rule_option: [
              {
                keyword: "sid",
                settings: ["1"],
              },
            ],
          }),
        },
      },
      tags: {
        Name: "permit HTTP from source",
      },
      type: "STATEFUL",
    });
  }
}

```

### Stateful Inspection for blocking packets from going to an intended destination

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkfirewallRuleGroup(this, "example", {
      capacity: 100,
      name: "example",
      ruleGroup: {
        rulesSource: {
          statefulRule: [
            {
              action: "DROP",
              header: {
                destination: "124.1.1.24/32",
                destinationPort: Token.asString(53),
                direction: "ANY",
                protocol: "TCP",
                source: "1.2.3.4/32",
                sourcePort: Token.asString(53),
              },
              ruleOption: [
                {
                  keyword: "sid",
                  settings: ["1"],
                },
              ],
            },
          ],
        },
      },
      tags: {
        Tag1: "Value1",
        Tag2: "Value2",
      },
      type: "STATEFUL",
    });
  }
}

```

### Stateful Inspection from rules specifications defined in Suricata flat format

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkfirewallRuleGroup(this, "example", {
      capacity: 100,
      name: "example",
      rules: Token.asString(Fn.file("example.rules")),
      tags: {
        Tag1: "Value1",
        Tag2: "Value2",
      },
      type: "STATEFUL",
    });
  }
}

```

### Stateful Inspection from rule group specifications using rule variables and Suricata format rules

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkfirewallRuleGroup(this, "example", {
      capacity: 100,
      name: "example",
      ruleGroup: {
        ruleVariables: {
          ipSets: [
            {
              ipSet: {
                definition: ["10.0.0.0/16", "10.0.1.0/24", "192.168.0.0/16"],
              },
              key: "WEBSERVERS_HOSTS",
            },
            {
              ipSet: {
                definition: ["1.2.3.4/32"],
              },
              key: "EXTERNAL_HOST",
            },
          ],
          portSets: [
            {
              key: "HTTP_PORTS",
              portSet: {
                definition: ["443", "80"],
              },
            },
          ],
        },
        rulesSource: {
          rulesString: Token.asString(Fn.file("suricata_rules_file")),
        },
      },
      tags: {
        Tag1: "Value1",
        Tag2: "Value2",
      },
      type: "STATEFUL",
    });
  }
}

```

### Stateless Inspection with a Custom Action

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkfirewallRuleGroup(this, "example", {
      capacity: 100,
      description: "Stateless Rate Limiting Rule",
      name: "example",
      ruleGroup: {
        rulesSource: {
          statelessRulesAndCustomActions: {
            customAction: [
              {
                actionDefinition: {
                  publishMetricAction: {
                    dimension: [
                      {
                        value: "2",
                      },
                    ],
                  },
                },
                actionName: "ExampleMetricsAction",
              },
            ],
            statelessRule: [
              {
                priority: 1,
                ruleDefinition: {
                  actions: ["aws:pass", "ExampleMetricsAction"],
                  matchAttributes: {
                    destination: [
                      {
                        addressDefinition: "124.1.1.5/32",
                      },
                    ],
                    destinationPort: [
                      {
                        fromPort: 443,
                        toPort: 443,
                      },
                    ],
                    protocols: [6],
                    source: [
                      {
                        addressDefinition: "1.2.3.4/32",
                      },
                    ],
                    sourcePort: [
                      {
                        fromPort: 443,
                        toPort: 443,
                      },
                    ],
                    tcpFlag: [
                      {
                        flags: ["SYN"],
                        masks: ["SYN", "ACK"],
                      },
                    ],
                  },
                },
              },
            ],
          },
        },
      },
      tags: {
        Tag1: "Value1",
        Tag2: "Value2",
      },
      type: "STATELESS",
    });
  }
}

```

### IP Set References to the Rule Group

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkfirewallRuleGroup(this, "example", {
      capacity: 100,
      name: "example",
      ruleGroup: {
        referenceSets: {
          ipSetReferences: [
            {
              ipSetReference: [
                {
                  referenceArn: thisVar.arn,
                },
              ],
              key: "example",
            },
          ],
        },
        rulesSource: {
          rulesSourceList: {
            generatedRulesType: "DENYLIST",
            targetTypes: ["HTTP_HOST"],
            targets: ["test.example.com"],
          },
        },
      },
      tags: {
        Tag1: "Value1",
        Tag2: "Value2",
      },
      type: "STATEFUL",
    });
  }
}

```

### Example with S3 as source for the suricata rules

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsS3Object } from "./.gen/providers/aws/data-aws-s3-object";
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const suricataRules = new DataAwsS3Object(this, "suricata_rules", {
      bucket: Token.asString(awsS3BucketSuricataRules.id),
      key: "rules/custom.rules",
    });
    new NetworkfirewallRuleGroup(this, "s3_rules_example", {
      capacity: 1000,
      name: "my-terraform-s3-rules",
      ruleGroup: {
        ruleVariables: {
          ipSets: [
            {
              ipSet: {
                definition: ["10.0.0.0/16", "192.168.0.0/16", "172.16.0.0/12"],
              },
              key: "HOME_NET",
            },
          ],
          portSets: [
            {
              key: "HTTP_PORTS",
              portSet: {
                definition: ["443", "80"],
              },
            },
          ],
        },
        rulesSource: {
          rulesString: Token.asString(suricataRules.body),
        },
      },
      tags: {
        ManagedBy: "terraform",
      },
      type: "STATEFUL",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `capacity` - (Required, Forces new resource) The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
* `description` - (Optional) A friendly description of the rule group.
* `encryptionConfiguration` - (Optional) KMS encryption configuration settings. See [Encryption Configuration](#encryption-configuration) below for details.
* `name` - (Required, Forces new resource) A friendly name of the rule group.
* `ruleGroup` - (Optional) A configuration block that defines the rule group rules. Required unless `rules` is specified. See [Rule Group](#rule-group) below for details.
* `rules` - (Optional) The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
* `tags` - (Optional) A map of key:value pairs to associate with the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `type` - (Required) Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.

### Encryption Configuration

`encryptionConfiguration` settings for customer managed KMS keys. Remove this block to use the default AWS-managed KMS encryption (rather than setting `type` to `AWS_OWNED_KMS_KEY`).

* `keyId` - (Optional) The ID of the customer managed key. You can use any of the [key identifiers](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
* `type` - (Required) The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are `CUSTOMER_KMS` and `AWS_OWNED_KMS_KEY`.

### Rule Group

The `ruleGroup` block supports the following argument:

* `referenceSets` - (Optional) A configuration block that defines the IP Set References for the rule group. See [Reference Sets](#reference-sets) below for details. Please notes that there can only be a maximum of 5 `referenceSets` in a `ruleGroup`. See the [AWS documentation](https://docs.aws.amazon.com/network-firewall/latest/developerguide/rule-groups-ip-set-references.html#rule-groups-ip-set-reference-limits) for details.

* `ruleVariables` - (Optional) A configuration block that defines additional settings available to use in the rules defined in the rule group. Can only be specified for **stateful** rule groups. See [Rule Variables](#rule-variables) below for details.

* `rulesSource` - (Required) A configuration block that defines the stateful or stateless rules for the rule group. See [Rules Source](#rules-source) below for details.

* `statefulRuleOptions` - (Optional) A configuration block that defines stateful rule options for the rule group. See [Stateful Rule Options](#stateful-rule-options) below for details.

### Reference Sets

The `referenceSets` block supports the following arguments:

* `ipSetReference` - (Optional) Set of configuration blocks that define the IP Reference information. See [IP Set Reference](#ip-set-reference) below for details.

### Rule Variables

The `ruleVariables` block supports the following arguments:

* `ipSets` - (Optional) Set of configuration blocks that define IP address information. See [IP Sets](#ip-sets) below for details.

* `portSets` - (Optional) Set of configuration blocks that define port range information. See [Port Sets](#port-sets) below for details.

### IP Sets

The `ipSets` block supports the following arguments:

* `key` - (Required) A unique alphanumeric string to identify the `ipSet`.

* `ipSet` - (Required) A configuration block that defines a set of IP addresses. See [IP Set](#ip-set) below for details.

### IP Set

The `ipSet` configuration block supports the following argument:

* `definition` - (Required) Set of IP addresses and address ranges, in CIDR notation.

### IP Set Reference

The `ipSetReference` configuration block supports the following argument:

* `key` - (Required) A unique alphanumeric string to identify the `ipSet`.

* `referenceArn` - (Required) Set of Managed Prefix IP ARN(s)

### Port Sets

The `portSets` block supports the following arguments:

* `key` - (Required) An unique alphanumeric string to identify the `portSet`.

* `portSet` - (Required) A configuration block that defines a set of port ranges. See [Port Set](#port-set) below for details.

### Port Set

The `portSet` configuration block suppports the following argument:

* `definition` - (Required) Set of port ranges.

### Rules Source

The `rulesSource` block supports the following arguments:

~> **NOTE:** Only one of `rulesSourceList`, `rulesString`, `statefulRule`, or `statelessRulesAndCustomActions` must be specified.

* `rulesSourceList` - (Optional) A configuration block containing **stateful** inspection criteria for a domain list rule group. See [Rules Source List](#rules-source-list) below for details.

* `rulesString` - (Optional) Stateful inspection criteria, provided in Suricata compatible rules. These rules contain the inspection criteria and the action to take for traffic that matches the criteria, so this type of rule group doesn’t have a separate action setting.

* `statefulRule` - (Optional) Set of configuration blocks containing **stateful** inspection criteria for 5-tuple rules to be used together in a rule group. See [Stateful Rule](#stateful-rule) below for details.

* `statelessRulesAndCustomActions` - (Optional) A configuration block containing **stateless** inspection criteria for a stateless rule group. See [Stateless Rules and Custom Actions](#stateless-rules-and-custom-actions) below for details.

### Stateful Rule Options

The `statefulRuleOptions` block supports the following argument:

~> **NOTE:** If the `STRICT_ORDER` rule order is specified, this rule group can only be referenced in firewall policies that also utilize `STRICT_ORDER` for the stateful engine. `STRICT_ORDER` can only be specified when using a `rulesSource` of `rulesString` or `statefulRule`.

* `ruleOrder` - (Required) Indicates how to manage the order of the rule evaluation for the rule group. Default value: `DEFAULT_ACTION_ORDER`. Valid values: `DEFAULT_ACTION_ORDER`, `STRICT_ORDER`.

### Rules Source List

The `rulesSourceList` block supports the following arguments:

* `generatedRulesType` - (Required) String value to specify whether domains in the target list are allowed or denied access. Valid values: `ALLOWLIST`, `DENYLIST`.

* `targetTypes` - (Required) Set of types of domain specifications that are provided in the `targets` argument. Valid values: `HTTP_HOST`, `TLS_SNI`.

* `targets` - (Required) Set of domains that you want to inspect for in your traffic flows.

### Stateful Rule

The `statefulRule` block supports the following arguments:

* `action` - (Required) Action to take with packets in a traffic flow when the flow matches the stateful rule criteria. For all actions, AWS Network Firewall performs the specified action and discontinues stateful inspection of the traffic flow. Valid values: `ALERT`, `DROP`, `PASS`, or `REJECT`.

* `header` - (Required) A configuration block containing the stateful 5-tuple inspection criteria for the rule, used to inspect traffic flows. See [Header](#header) below for details.

* `ruleOption` - (Required) Set of configuration blocks containing additional settings for a stateful rule. See [Rule Option](#rule-option) below for details.

### Stateless Rules and Custom Actions

The `statelessRulesAndCustomActions` block supports the following arguments:

* `customAction` - (Optional) Set of configuration blocks containing custom action definitions that are available for use by the set of `stateless rule`. See [Custom Action](#custom-action) below for details.

* `statelessRule` - (Required) Set of configuration blocks containing the stateless rules for use in the stateless rule group. See [Stateless Rule](#stateless-rule) below for details.

### Header

The `header` block supports the following arguments:

* `destination` - (Required) The destination IP address or address range to inspect for, in CIDR notation. To match with any address, specify `ANY`.

* `destinationPort` - (Required) The destination port to inspect for. To match with any address, specify `ANY`.

* `direction` - (Required) The direction of traffic flow to inspect. Valid values: `ANY` or `FORWARD`.

* `protocol` - (Required) The protocol to inspect. Valid values: `IP`, `TCP`, `UDP`, `ICMP`, `HTTP`, `FTP`, `TLS`, `SMB`, `DNS`, `DCERPC`, `SSH`, `SMTP`, `IMAP`, `MSN`, `KRB5`, `IKEV2`, `TFTP`, `NTP`, `DHCP`.

* `source` - (Required) The source IP address or address range for, in CIDR notation. To match with any address, specify `ANY`.

* `sourcePort` - (Required) The source port to inspect for. To match with any address, specify `ANY`.

### Rule Option

The `ruleOption` block supports the following arguments:

* `keyword` - (Required) Keyword defined by open source detection systems like Snort or Suricata for stateful rule inspection.
See [Snort General Rule Options](http://manual-snort-org.s3-website-us-east-1.amazonaws.com/node31.html) or [Suricata Rule Options](https://suricata.readthedocs.io/en/suricata-5.0.1/rules/intro.html#rule-options) for more details.
* `settings` - (Optional) Set of strings for additional settings to use in stateful rule inspection.

### Custom Action

The `customAction` block supports the following arguments:

* `actionDefinition` - (Required) A configuration block describing the custom action associated with the `actionName`. See [Action Definition](#action-definition) below for details.

* `actionName` - (Required, Forces new resource) A friendly name of the custom action.

### Stateless Rule

The `statelessRule` block supports the following arguments:

* `priority` - (Required) A setting that indicates the order in which to run this rule relative to all of the rules that are defined for a stateless rule group. AWS Network Firewall evaluates the rules in a rule group starting with the lowest priority setting.

* `ruleDefinition` - (Required) A configuration block defining the stateless 5-tuple packet inspection criteria and the action to take on a packet that matches the criteria. See [Rule Definition](#rule-definition) below for details.

### Rule Definition

The `ruleDefinition` block supports the following arguments:

* `actions` - (Required) Set of actions to take on a packet that matches one of the stateless rule definition's `matchAttributes`. For every rule you must specify 1 standard action, and you can add custom actions. Standard actions include: `aws:pass`, `aws:drop`, `aws:forward_to_sfe`.

* `matchAttributes` - (Required) A configuration block containing criteria for AWS Network Firewall to use to inspect an individual packet in stateless rule inspection. See [Match Attributes](#match-attributes) below for details.

### Match Attributes

The `matchAttributes` block supports the following arguments:

* `destination` - (Optional) Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See [Destination](#destination) below for details.

* `destinationPort` - (Optional) Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See [Destination Port](#destination-port) below for details.

* `protocols` - (Optional) Set of protocols to inspect for, specified using the protocol's assigned internet protocol number (IANA). If not specified, this matches with any protocol.

* `source` - (Optional) Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See [Source](#source) below for details.

* `sourcePort` - (Optional) Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See [Source Port](#source-port) below for details.

* `tcpFlag` - (Optional) Set of configuration blocks containing the TCP flags and masks to inspect for. If not specified, this matches with any settings.

### Action Definition

The `actionDefinition` block supports the following argument:

* `publishMetricAction` - (Required) A configuration block describing the stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. You can pair this custom action with any of the standard stateless rule actions. See [Publish Metric Action](#publish-metric-action) below for details.

### Publish Metric Action

The `publishMetricAction` block supports the following argument:

* `dimension` - (Required) Set of configuration blocks containing the dimension settings to use for Amazon CloudWatch custom metrics. See [Dimension](#dimension) below for details.

### Dimension

The `dimension` block supports the following argument:

* `value` - (Required) The value to use in the custom metric dimension.

### Destination

The `destination` block supports the following argument:

* `addressDefinition` - (Required)  An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.

### Destination Port

The `destinationPort` block supports the following arguments:

* `fromPort` - (Required) The lower limit of the port range. This must be less than or equal to the `toPort`.

* `toPort` - (Optional) The upper limit of the port range. This must be greater than or equal to the `fromPort`.

### Source

The `source` block supports the following argument:

* `addressDefinition` - (Required)  An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.

### Source Port

The `sourcePort` block supports the following arguments:

* `fromPort` - (Required) The lower limit of the port range. This must be less than or equal to the `toPort`.

* `toPort` - (Optional) The upper limit of the port range. This must be greater than or equal to the `fromPort`.

### TCP Flag

The `tcpFlag` block supports the following arguments:

* `flags` - (Required) Set of flags to look for in a packet. This setting can only specify values that are also specified in `masks`.
Valid values: `FIN`, `SYN`, `RST`, `PSH`, `ACK`, `URG`, `ECE`, `CWR`.

* `masks` - (Optional) Set of flags to consider in the inspection. To inspect all flags, leave this empty.
Valid values: `FIN`, `SYN`, `RST`, `PSH`, `ACK`, `URG`, `ECE`, `CWR`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Amazon Resource Name (ARN) that identifies the rule group.

* `arn` - The Amazon Resource Name (ARN) that identifies the rule group.

* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

* `updateToken` - A string token used when updating the rule group.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Network Firewall Rule Groups using their `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkfirewallRuleGroup } from "./.gen/providers/aws/networkfirewall-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    NetworkfirewallRuleGroup.generateConfigForImport(
      this,
      "example",
      "arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example"
    );
  }
}

```

Using `terraform import`, import Network Firewall Rule Groups using their `arn`. For example:

```console
% terraform import aws_networkfirewall_rule_group.example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
```

<!-- cache-key: cdktf-0.20.8 input-8af1aae5893c42f1d45aad2d4399c0ee56e43ef2d30266db158f9832bb12a2ca -->