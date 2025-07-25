---
subcategory: "ACM PCA (Certificate Manager Private Certificate Authority)"
layout: "aws"
page_title: "AWS: aws_acmpca_certificate_authority"
description: |-
  Provides a resource to manage AWS Certificate Manager Private Certificate Authorities
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_acmpca_certificate_authority

Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).

~> **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificate_signing_request` attribute. The [`aws_acmpca_certificate_authority_certificate`](/docs/providers/aws/r/acmpca_certificate_authority_certificate.html) resource can be used for this purpose.

## Example Usage

### Basic

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.acmpca_certificate_authority import AcmpcaCertificateAuthority
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AcmpcaCertificateAuthority(self, "example",
            certificate_authority_configuration=AcmpcaCertificateAuthorityCertificateAuthorityConfiguration(
                key_algorithm="RSA_4096",
                signing_algorithm="SHA512WITHRSA",
                subject=AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject(
                    common_name="example.com"
                )
            ),
            permanent_deletion_time_in_days=7
        )
```

### Short-lived certificate

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.acmpca_certificate_authority import AcmpcaCertificateAuthority
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AcmpcaCertificateAuthority(self, "example",
            certificate_authority_configuration=AcmpcaCertificateAuthorityCertificateAuthorityConfiguration(
                key_algorithm="RSA_4096",
                signing_algorithm="SHA512WITHRSA",
                subject=AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject(
                    common_name="example.com"
                )
            ),
            usage_mode="SHORT_LIVED_CERTIFICATE"
        )
```

### Enable Certificate Revocation List

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.acmpca_certificate_authority import AcmpcaCertificateAuthority
from imports.aws.data_aws_iam_policy_document import DataAwsIamPolicyDocument
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_bucket_policy import S3BucketPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = S3Bucket(self, "example",
            bucket="example",
            force_destroy=True
        )
        acmpca_bucket_access = DataAwsIamPolicyDocument(self, "acmpca_bucket_access",
            statement=[DataAwsIamPolicyDocumentStatement(
                actions=["s3:GetBucketAcl", "s3:GetBucketLocation", "s3:PutObject", "s3:PutObjectAcl"
                ],
                principals=[DataAwsIamPolicyDocumentStatementPrincipals(
                    identifiers=["acm-pca.amazonaws.com"],
                    type="Service"
                )
                ],
                resources=[example.arn, "${" + example.arn + "}/*"]
            )
            ]
        )
        aws_s3_bucket_policy_example = S3BucketPolicy(self, "example_2",
            bucket=example.id,
            policy=Token.as_string(acmpca_bucket_access.json)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_policy_example.override_logical_id("example")
        aws_acmpca_certificate_authority_example = AcmpcaCertificateAuthority(self, "example_3",
            certificate_authority_configuration=AcmpcaCertificateAuthorityCertificateAuthorityConfiguration(
                key_algorithm="RSA_4096",
                signing_algorithm="SHA512WITHRSA",
                subject=AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject(
                    common_name="example.com"
                )
            ),
            depends_on=[aws_s3_bucket_policy_example],
            revocation_configuration=AcmpcaCertificateAuthorityRevocationConfiguration(
                crl_configuration=AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration(
                    custom_cname="crl.example.com",
                    enabled=True,
                    expiration_in_days=7,
                    s3_bucket_name=example.id,
                    s3_object_acl="BUCKET_OWNER_FULL_CONTROL"
                )
            )
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_acmpca_certificate_authority_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `certificate_authority_configuration` - (Required) Nested argument containing algorithms and certificate subject information. Defined below.
* `enabled` - (Optional) Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
* `revocation_configuration` - (Optional) Nested argument containing revocation configuration. Defined below.
* `usage_mode` - (Optional) Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
* `tags` - (Optional) Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `type` - (Optional) Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
* `key_storage_security_standard` - (Optional) Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
* `permanent_deletion_time_in_days` - (Optional) Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.

### certificate_authority_configuration

* `key_algorithm` - (Required) Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
* `signing_algorithm` - (Required) Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
* `subject` - (Required) Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.

#### subject

Contains information about the certificate subject. Identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service.

* `common_name` - (Optional) Fully qualified domain name (FQDN) associated with the certificate subject. Must be less than or equal to 64 characters in length.
* `country` - (Optional) Two digit code that specifies the country in which the certificate subject located. Must be less than or equal to 2 characters in length.
* `distinguished_name_qualifier` - (Optional) Disambiguating information for the certificate subject. Must be less than or equal to 64 characters in length.
* `generation_qualifier` - (Optional) Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third. Must be less than or equal to 3 characters in length.
* `given_name` - (Optional) First name. Must be less than or equal to 16 characters in length.
* `initials` - (Optional) Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`. Must be less than or equal to 5 characters in length.
* `locality` - (Optional) Locality (such as a city or town) in which the certificate subject is located. Must be less than or equal to 128 characters in length.
* `organization` - (Optional) Legal name of the organization with which the certificate subject is affiliated. Must be less than or equal to 64 characters in length.
* `organizational_unit` - (Optional) Subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated. Must be less than or equal to 64 characters in length.
* `pseudonym` - (Optional) Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza. Must be less than or equal to 128 characters in length.
* `state` - (Optional) State in which the subject of the certificate is located. Must be less than or equal to 128 characters in length.
* `surname` - (Optional) Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first. Must be less than or equal to 40 characters in length.
* `title` - (Optional) Title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject. Must be less than or equal to 64 characters in length.

### revocation_configuration

* `crl_configuration` - (Optional) Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
* `ocsp_configuration` - (Optional) Nested argument containing configuration of
the custom OCSP responder endpoint. Defined below.

#### crl_configuration

* `custom_cname` - (Optional) Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public. Must be less than or equal to 253 characters in length.
* `enabled` - (Optional) Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
* `expiration_in_days` - (Optional, Required if `enabled` is `true`) Number of days until a certificate expires. Must be between 1 and 5000.
* `s3_bucket_name` - (Optional, Required if `enabled` is `true`) Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket. Must be between 3 and 255 characters in length.
* `s3_object_acl` - (Optional) Determines whether the CRL will be publicly readable or privately held in the CRL Amazon S3 bucket. Defaults to `PUBLIC_READ`.

#### ocsp_configuration

* `enabled` - (Required) Boolean value that specifies whether a custom OCSP responder is enabled.
* `ocsp_custom_cname` - (Optional) CNAME specifying a customized OCSP domain. Note: The value of the CNAME must not include a protocol prefix such as "http://" or "https://".

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ARN of the certificate authority.
* `arn` - ARN of the certificate authority.
* `certificate` - Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
* `certificate_chain` - Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
* `certificate_signing_request` - The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
* `not_after` - Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
* `not_before` - Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
* `serial` - Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `1m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_acmpca_certificate_authority` using the certificate authority ARN. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.acmpca_certificate_authority import AcmpcaCertificateAuthority
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AcmpcaCertificateAuthority.generate_config_for_import(self, "example", "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012")
```

Using `terraform import`, import `aws_acmpca_certificate_authority` using the certificate authority ARN. For example:

```console
% terraform import aws_acmpca_certificate_authority.example arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
```

<!-- cache-key: cdktf-0.20.8 input-449076f65ae9804f638028cc528cd156eddcb4de1fdb91f7fea07302be21940d -->