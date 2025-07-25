// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_verifiedaccess_trust_provider", name="Verified Access Trust Provider")
// @Tags(identifierAttribute="id")
// @Testing(tagsTest=false)
func resourceVerifiedAccessTrustProvider() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceVerifiedAccessTrustProviderCreate,
		ReadWithoutTimeout:   resourceVerifiedAccessTrustProviderRead,
		UpdateWithoutTimeout: resourceVerifiedAccessTrustProviderUpdate,
		DeleteWithoutTimeout: resourceVerifiedAccessTrustProviderDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrDescription: {
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_options": {
				Type:     schema.TypeList,
				ForceNew: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tenant_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"device_trust_provider_type": {
				Type:             schema.TypeString,
				ForceNew:         true,
				Optional:         true,
				ValidateDiagFunc: enum.Validate[awstypes.DeviceTrustProviderType](),
			},
			"native_application_oidc_options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authorization_endpoint": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
						names.AttrClientID: {
							Type:     schema.TypeString,
							ForceNew: true,
							Optional: true,
						},
						names.AttrClientSecret: {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						names.AttrIssuer: {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
						"public_signing_key_endpoint": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
						names.AttrScope: {
							Type:     schema.TypeString,
							Optional: true,
						},
						"token_endpoint": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
						"user_info_endpoint": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
					},
				},
			},
			"oidc_options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authorization_endpoint": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
						names.AttrClientID: {
							Type:     schema.TypeString,
							ForceNew: true,
							Optional: true,
						},
						names.AttrClientSecret: {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						names.AttrIssuer: {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
						names.AttrScope: {
							Type:     schema.TypeString,
							Optional: true,
						},
						"token_endpoint": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
						"user_info_endpoint": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPS,
						},
					},
				},
			},
			"policy_reference_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"sse_specification": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"customer_managed_key_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						names.AttrKMSKeyARN: {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidARN,
						},
					},
				},
			},
			"trust_provider_type": {
				Type:             schema.TypeString,
				ForceNew:         true,
				Required:         true,
				ValidateDiagFunc: enum.Validate[awstypes.TrustProviderType](),
			},
			"user_trust_provider_type": {
				Type:             schema.TypeString,
				ForceNew:         true,
				Optional:         true,
				ValidateDiagFunc: enum.Validate[awstypes.UserTrustProviderType](),
			},
		},
	}
}

func resourceVerifiedAccessTrustProviderCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	input := ec2.CreateVerifiedAccessTrustProviderInput{
		ClientToken:         aws.String(sdkid.UniqueId()),
		PolicyReferenceName: aws.String(d.Get("policy_reference_name").(string)),
		TagSpecifications:   getTagSpecificationsIn(ctx, awstypes.ResourceTypeVerifiedAccessTrustProvider),
		TrustProviderType:   awstypes.TrustProviderType(d.Get("trust_provider_type").(string)),
	}

	if v, ok := d.GetOk(names.AttrDescription); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("device_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.DeviceOptions = expandCreateVerifiedAccessTrustProviderDeviceOptions(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk("device_trust_provider_type"); ok {
		input.DeviceTrustProviderType = awstypes.DeviceTrustProviderType(v.(string))
	}

	if v, ok := d.GetOk("native_application_oidc_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.NativeApplicationOidcOptions = expandCreateVerifiedAccessTrustProviderNativeApplicationOIDCOptions(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk("oidc_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.OidcOptions = expandCreateVerifiedAccessTrustProviderOIDCOptions(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk("sse_specification"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.SseSpecification = expandVerifiedAccessSSESpecificationRequest(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk("user_trust_provider_type"); ok {
		input.UserTrustProviderType = awstypes.UserTrustProviderType(v.(string))
	}

	output, err := conn.CreateVerifiedAccessTrustProvider(ctx, &input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating Verified Access Trust Provider: %s", err)
	}

	d.SetId(aws.ToString(output.VerifiedAccessTrustProvider.VerifiedAccessTrustProviderId))

	return append(diags, resourceVerifiedAccessTrustProviderRead(ctx, d, meta)...)
}

func resourceVerifiedAccessTrustProviderRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	output, err := findVerifiedAccessTrustProviderByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EC2 Verified Access Trust Provider (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading Verified Access Trust Provider (%s): %s", d.Id(), err)
	}

	d.Set(names.AttrDescription, output.Description)
	if v := output.DeviceOptions; v != nil {
		if err := d.Set("device_options", flattenDeviceOptions(v)); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting device_options: %s", err)
		}
	} else {
		d.Set("device_options", nil)
	}
	d.Set("device_trust_provider_type", output.DeviceTrustProviderType)
	if v := output.OidcOptions; v != nil {
		if err := d.Set("oidc_options", flattenOIDCOptions(v, d.Get("oidc_options.0.client_secret").(string))); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting oidc_options: %s", err)
		}
	} else {
		d.Set("oidc_options", nil)
	}
	if v := output.NativeApplicationOidcOptions; v != nil {
		if err := d.Set("native_application_oidc_options", flattenNativeApplicationOIDCOptions(v, d.Get("native_application_oidc_options.0.client_secret").(string))); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting native_application_oidc_options: %s", err)
		}
	} else {
		d.Set("native_application_oidc_options", nil)
	}

	d.Set("policy_reference_name", output.PolicyReferenceName)
	d.Set("trust_provider_type", output.TrustProviderType)
	if err := d.Set("sse_specification", flattenVerifiedAccessSSESpecificationResponse(output.SseSpecification)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting sse_specification: %s", err)
	}
	d.Set("user_trust_provider_type", output.UserTrustProviderType)

	setTagsOut(ctx, output.Tags)

	return diags
}

func resourceVerifiedAccessTrustProviderUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	if d.HasChangesExcept(names.AttrTags, names.AttrTagsAll) {
		input := ec2.ModifyVerifiedAccessTrustProviderInput{
			ClientToken:                   aws.String(sdkid.UniqueId()),
			VerifiedAccessTrustProviderId: aws.String(d.Id()),
		}

		if d.HasChange(names.AttrDescription) {
			input.Description = aws.String(d.Get(names.AttrDescription).(string))
		}

		if d.HasChange("oidc_options") {
			if v, ok := d.GetOk("oidc_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
				input.OidcOptions = expandModifyVerifiedAccessTrustProviderOIDCOptions(v.([]any)[0].(map[string]any))
			}
		}

		if d.HasChange("native_application_oidc_options") {
			if v, ok := d.GetOk("native_application_oidc_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
				input.NativeApplicationOidcOptions = expandModifyVerifiedAccessTrustProviderNativeApplicationOIDCOptions(v.([]any)[0].(map[string]any))
			}
		}

		_, err := conn.ModifyVerifiedAccessTrustProvider(ctx, &input)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating Verified Access Trust Provider (%s): %s", d.Id(), err)
		}
	}

	return append(diags, resourceVerifiedAccessTrustProviderRead(ctx, d, meta)...)
}

func resourceVerifiedAccessTrustProviderDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	log.Printf("[INFO] Deleting Verified Access Trust Provider: %s", d.Id())
	input := ec2.DeleteVerifiedAccessTrustProviderInput{
		ClientToken:                   aws.String(sdkid.UniqueId()),
		VerifiedAccessTrustProviderId: aws.String(d.Id()),
	}
	_, err := conn.DeleteVerifiedAccessTrustProvider(ctx, &input)

	if tfawserr.ErrCodeEquals(err, errCodeInvalidVerifiedAccessTrustProviderIdNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting Verified Access Trust Provider (%s): %s", d.Id(), err)
	}

	return diags
}

func flattenDeviceOptions(apiObject *awstypes.DeviceOptions) []any {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]any{}

	if v := apiObject.TenantId; v != nil {
		tfMap["tenant_id"] = aws.ToString(v)
	}

	return []any{tfMap}
}

func flattenNativeApplicationOIDCOptions(apiObject *awstypes.NativeApplicationOidcOptions, clientSecret string) []any {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]any{
		names.AttrClientSecret: clientSecret,
	}

	if v := apiObject.AuthorizationEndpoint; v != nil {
		tfMap["authorization_endpoint"] = aws.ToString(v)
	}

	if v := apiObject.ClientId; v != nil {
		tfMap[names.AttrClientID] = aws.ToString(v)
	}

	if v := apiObject.Issuer; v != nil {
		tfMap[names.AttrIssuer] = aws.ToString(v)
	}

	if v := apiObject.PublicSigningKeyEndpoint; v != nil {
		tfMap["public_signing_key_endpoint"] = aws.ToString(v)
	}

	if v := apiObject.Scope; v != nil {
		tfMap[names.AttrScope] = aws.ToString(v)
	}

	if v := apiObject.TokenEndpoint; v != nil {
		tfMap["token_endpoint"] = aws.ToString(v)
	}

	if v := apiObject.UserInfoEndpoint; v != nil {
		tfMap["user_info_endpoint"] = aws.ToString(v)
	}

	return []any{tfMap}
}

func flattenOIDCOptions(apiObject *awstypes.OidcOptions, clientSecret string) []any {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]any{
		names.AttrClientSecret: clientSecret,
	}

	if v := apiObject.AuthorizationEndpoint; v != nil {
		tfMap["authorization_endpoint"] = aws.ToString(v)
	}

	if v := apiObject.ClientId; v != nil {
		tfMap[names.AttrClientID] = aws.ToString(v)
	}

	if v := apiObject.Issuer; v != nil {
		tfMap[names.AttrIssuer] = aws.ToString(v)
	}

	if v := apiObject.Scope; v != nil {
		tfMap[names.AttrScope] = aws.ToString(v)
	}

	if v := apiObject.TokenEndpoint; v != nil {
		tfMap["token_endpoint"] = aws.ToString(v)
	}

	if v := apiObject.UserInfoEndpoint; v != nil {
		tfMap["user_info_endpoint"] = aws.ToString(v)
	}

	return []any{tfMap}
}

func expandCreateVerifiedAccessTrustProviderDeviceOptions(tfMap map[string]any) *awstypes.CreateVerifiedAccessTrustProviderDeviceOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.CreateVerifiedAccessTrustProviderDeviceOptions{}

	if v, ok := tfMap["tenant_id"].(string); ok && v != "" {
		apiObject.TenantId = aws.String(v)
	}

	return apiObject
}

func expandCreateVerifiedAccessTrustProviderOIDCOptions(tfMap map[string]any) *awstypes.CreateVerifiedAccessTrustProviderOidcOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.CreateVerifiedAccessTrustProviderOidcOptions{}

	if v, ok := tfMap["authorization_endpoint"].(string); ok && v != "" {
		apiObject.AuthorizationEndpoint = aws.String(v)
	}

	if v, ok := tfMap[names.AttrClientID].(string); ok && v != "" {
		apiObject.ClientId = aws.String(v)
	}

	if v, ok := tfMap[names.AttrClientSecret].(string); ok && v != "" {
		apiObject.ClientSecret = aws.String(v)
	}

	if v, ok := tfMap[names.AttrIssuer].(string); ok && v != "" {
		apiObject.Issuer = aws.String(v)
	}

	if v, ok := tfMap[names.AttrScope].(string); ok && v != "" {
		apiObject.Scope = aws.String(v)
	}

	if v, ok := tfMap["token_endpoint"].(string); ok && v != "" {
		apiObject.TokenEndpoint = aws.String(v)
	}

	if v, ok := tfMap["user_info_endpoint"].(string); ok && v != "" {
		apiObject.UserInfoEndpoint = aws.String(v)
	}

	return apiObject
}

func expandCreateVerifiedAccessTrustProviderNativeApplicationOIDCOptions(tfMap map[string]any) *awstypes.CreateVerifiedAccessNativeApplicationOidcOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.CreateVerifiedAccessNativeApplicationOidcOptions{}

	if v, ok := tfMap["authorization_endpoint"].(string); ok && v != "" {
		apiObject.AuthorizationEndpoint = aws.String(v)
	}

	if v, ok := tfMap[names.AttrClientID].(string); ok && v != "" {
		apiObject.ClientId = aws.String(v)
	}

	if v, ok := tfMap[names.AttrClientSecret].(string); ok && v != "" {
		apiObject.ClientSecret = aws.String(v)
	}

	if v, ok := tfMap["public_signing_key_endpoint"].(string); ok && v != "" {
		apiObject.PublicSigningKeyEndpoint = aws.String(v)
	}

	if v, ok := tfMap[names.AttrIssuer].(string); ok && v != "" {
		apiObject.Issuer = aws.String(v)
	}

	if v, ok := tfMap[names.AttrScope].(string); ok && v != "" {
		apiObject.Scope = aws.String(v)
	}

	if v, ok := tfMap["token_endpoint"].(string); ok && v != "" {
		apiObject.TokenEndpoint = aws.String(v)
	}

	if v, ok := tfMap["user_info_endpoint"].(string); ok && v != "" {
		apiObject.UserInfoEndpoint = aws.String(v)
	}

	return apiObject
}

func expandModifyVerifiedAccessTrustProviderNativeApplicationOIDCOptions(tfMap map[string]any) *awstypes.ModifyVerifiedAccessNativeApplicationOidcOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.ModifyVerifiedAccessNativeApplicationOidcOptions{}

	if v, ok := tfMap[names.AttrScope].(string); ok && v != "" {
		apiObject.Scope = aws.String(v)
	}

	return apiObject
}

func expandModifyVerifiedAccessTrustProviderOIDCOptions(tfMap map[string]any) *awstypes.ModifyVerifiedAccessTrustProviderOidcOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.ModifyVerifiedAccessTrustProviderOidcOptions{}

	if v, ok := tfMap[names.AttrScope].(string); ok && v != "" {
		apiObject.Scope = aws.String(v)
	}

	return apiObject
}
