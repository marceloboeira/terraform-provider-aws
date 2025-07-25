// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambda

// Exports for use in tests only.
var (
	ResourceAlias                        = resourceAlias
	ResourceCodeSigningConfig            = resourceCodeSigningConfig
	ResourceEventSourceMapping           = resourceEventSourceMapping
	ResourceFunction                     = resourceFunction
	ResourceFunctionEventInvokeConfig    = resourceFunctionEventInvokeConfig
	ResourceFunctionURL                  = resourceFunctionURL
	ResourceInvocation                   = resourceInvocation
	ResourceLayerVersion                 = resourceLayerVersion
	ResourceLayerVersionPermission       = resourceLayerVersionPermission
	ResourcePermission                   = resourcePermission
	ResourceProvisionedConcurrencyConfig = resourceProvisionedConcurrencyConfig

	FindAliasByTwoPartKey                        = findAliasByTwoPartKey
	FindCodeSigningConfigByARN                   = findCodeSigningConfigByARN
	FindEventSourceMappingByID                   = findEventSourceMappingByID
	FindFunctionByName                           = findFunctionByName
	FindFunctionEventInvokeConfigByTwoPartKey    = findFunctionEventInvokeConfigByTwoPartKey
	FindFunctionRecursionConfigByName            = findFunctionRecursionConfigByName
	FindFunctionURLByTwoPartKey                  = findFunctionURLByTwoPartKey
	FindLayerVersionByTwoPartKey                 = findLayerVersionByTwoPartKey
	FindLayerVersionPolicyByTwoPartKey           = findLayerVersionPolicyByTwoPartKey
	FindPolicyStatementByTwoPartKey              = findPolicyStatementByTwoPartKey
	FindProvisionedConcurrencyConfigByTwoPartKey = findProvisionedConcurrencyConfigByTwoPartKey
	FindRuntimeManagementConfigByTwoPartKey      = findRuntimeManagementConfigByTwoPartKey
	FunctionEventInvokeConfigParseResourceID     = functionEventInvokeConfigParseResourceID
	GetFunctionNameFromARN                       = getFunctionNameFromARN
	GetQualifierFromAliasOrVersionARN            = getQualifierFromAliasOrVersionARN
	LayerVersionParseResourceID                  = layerVersionParseResourceID
	LayerVersionPermissionParseResourceID        = layerVersionPermissionParseResourceID
	SignerServiceIsAvailable                     = signerServiceIsAvailable

	ValidFunctionName               = validFunctionName
	ValidPermissionAction           = validPermissionAction
	ValidPermissionEventSourceToken = validPermissionEventSourceToken
	ValidQualifier                  = validQualifier
	ValidPolicyStatementID          = validPolicyStatementID
)

type (
	Policy          = policy
	PolicyStatement = policyStatement
)
