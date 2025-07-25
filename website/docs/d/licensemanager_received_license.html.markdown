---
subcategory: "License Manager"
layout: "aws"
page_title: "AWS: aws_licensemanager_received_license"
description: |-
    Get information about a set of license manager received license
---

# Data Source: aws_licensemanager_received_license

This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.

## Example Usage

The following shows getting the received license data using and ARN.

```terraform
data "aws_licensemanager_received_license" "test" {
  license_arn = "arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0"
}
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `license_arn` - (Required) The ARN of the received license you want data for.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - The received license ARN (Same as: `license_arn`).
* `beneficiary` - Granted license beneficiary. This is in the form of the ARN of the root user of the account.
* `consumption_configuration` - Configuration for consumption of the license. [Detailed below](#consumption_configuration)
* `create_time` - Creation time of the granted license in RFC 3339 format.
* `entitlements` - License entitlements. [Detailed below](#entitlements)
* `home_region` - Home Region of the granted license.
* `issuer` - Granted license issuer. [Detailed below](#issuer)
* `license_arn` - Amazon Resource Name (ARN) of the license.
* `license_metadata`- Granted license metadata. This is in the form of a set of all meta data. [Detailed below](#license_metadata)
* `license_name` - License name.
* `product_name` - Product name.
* `product_sku ` - Product SKU.
* `received_metadata` - Granted license received metadata. [Detailed below](#received_metadata)
* `status` - Granted license status.
* `validity` - Date and time range during which the granted license is valid, in ISO8601-UTC format. [Detailed below](#validity)
* `version` - Version of the granted license.

### consumption_configuration

* `borrow_configuration` - Details about a borrow configuration. [Detailed below](#borrow_configuration)
* `provisional_configuration` - Details about a provisional configuration. [Detailed below](#provisional_configuration)
* `renewal_frequency` - Renewal frequency.

#### borrow_configuration

A list with a single map.

* `allow_early_check_in` - Indicates whether early check-ins are allowed.
* `max_time_to_live_in_minutes` - Maximum time for the borrow configuration, in minutes.

#### provisional_configuration

A list with a single map.

* `max_time_to_live_in_minutes` - Maximum time for the provisional configuration, in minutes.

### entitlements

A list with a single map.

* `allow_check_in` - Indicates whether check-ins are allowed.
* `max_count` - Maximum entitlement count. Use if the unit is not None.
* `name` - Entitlement name.
* `overage` - Indicates whether overages are allowed.
* `unit` - Entitlement unit.
* `value` - Entitlement resource. Use only if the unit is None.

### issuer

A list with a single map.

* `key_fingerprint` - Issuer key fingerprint.
* `name` - Issuer name.
* `sign_key` - Asymmetric KMS key from AWS Key Management Service. The KMS key must have a key usage of sign and verify, and support the RSASSA-PSS SHA-256 signing algorithm.

### license_metadata

Each metadata item will have the following attributes.

* `name` - The key name.
* `value` - The value.

### received_metadata

A list with a single map.

* `allowed_operations` - A list of allowed operations.
* `received_status` - Received status.
* `received_status_reason` - Received status reason.

### validity

A list with a single map.

* `begin` - Start of the validity time range.
* `end` - End of the validity time range.
