// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	awstypes "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	"github.com/hashicorp/terraform-provider-aws/internal/retry"
	tfslices "github.com/hashicorp/terraform-provider-aws/internal/slices"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_cloudwatch_query_definition", name="Query Definition")
func resourceQueryDefinition() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceQueryDefinitionPut,
		ReadWithoutTimeout:   resourceQueryDefinitionRead,
		UpdateWithoutTimeout: resourceQueryDefinitionPut,
		DeleteWithoutTimeout: resourceQueryDefinitionDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceQueryDefinitionImport,
		},

		Schema: map[string]*schema.Schema{
			names.AttrName: {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 255),
					validation.StringMatch(regexache.MustCompile(`^([^:*\/]+\/?)*[^:*\/]+$`), "cannot contain a colon or asterisk and cannot start or end with a slash"),
				),
			},
			"log_group_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.Any(
						validLogGroupName,
						verify.ValidARN,
					),
				},
			},
			"query_definition_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"query_string": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceQueryDefinitionPut(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).LogsClient(ctx)

	name := d.Get(names.AttrName).(string)
	input := &cloudwatchlogs.PutQueryDefinitionInput{
		Name:        aws.String(name),
		QueryString: aws.String(d.Get("query_string").(string)),
	}

	if v, ok := d.GetOk("log_group_names"); ok && len(v.([]any)) > 0 {
		input.LogGroupNames = flex.ExpandStringValueList(v.([]any))
	}

	if !d.IsNewResource() {
		input.QueryDefinitionId = aws.String(d.Id())
	}

	output, err := conn.PutQueryDefinition(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "putting CloudWatch Logs Query Definition (%s): %s", name, err)
	}

	if d.IsNewResource() {
		d.SetId(aws.ToString(output.QueryDefinitionId))
	}

	return append(diags, resourceQueryDefinitionRead(ctx, d, meta)...)
}

func resourceQueryDefinitionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).LogsClient(ctx)

	result, err := findQueryDefinitionByTwoPartKey(ctx, conn, d.Get(names.AttrName).(string), d.Id())

	if !d.IsNewResource() && retry.NotFound(err) {
		log.Printf("[WARN] CloudWatch Logs Query Definition (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading CloudWatch Logs Query Definition (%s): %s", d.Id(), err)
	}

	d.Set("log_group_names", result.LogGroupNames)
	d.Set(names.AttrName, result.Name)
	d.Set("query_definition_id", result.QueryDefinitionId)
	d.Set("query_string", result.QueryString)

	return diags
}

func resourceQueryDefinitionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).LogsClient(ctx)

	log.Printf("[INFO] Deleting CloudWatch Logs Query Definition: %s", d.Id())
	_, err := conn.DeleteQueryDefinition(ctx, &cloudwatchlogs.DeleteQueryDefinitionInput{
		QueryDefinitionId: aws.String(d.Id()),
	})

	if errs.IsA[*awstypes.ResourceNotFoundException](err) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting CloudWatch Logs Query Definition (%s): %s", d.Id(), err)
	}

	return diags
}

func resourceQueryDefinitionImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	arn, err := arn.Parse(d.Id())
	if err != nil {
		return nil, fmt.Errorf("unexpected format for ID (%s), expected a CloudWatch query definition ARN", d.Id())
	}

	if arn.Service != "logs" {
		return nil, fmt.Errorf("unexpected format for ID (%s), expected a CloudWatch query definition ARN", d.Id())
	}

	matcher := regexache.MustCompile("^query-definition:(" + verify.UUIDRegexPattern + ")$")
	matches := matcher.FindStringSubmatch(arn.Resource)
	if len(matches) != 2 {
		return nil, fmt.Errorf("unexpected format for ID (%s), expected a CloudWatch query definition ARN", d.Id())
	}

	d.SetId(matches[1])

	return []*schema.ResourceData{d}, nil
}

func findQueryDefinitionByTwoPartKey(ctx context.Context, conn *cloudwatchlogs.Client, name, queryDefinitionID string) (*awstypes.QueryDefinition, error) {
	input := cloudwatchlogs.DescribeQueryDefinitionsInput{}
	if name != "" {
		input.QueryDefinitionNamePrefix = aws.String(name)
	}

	return findQueryDefinition(ctx, conn, &input, func(v *awstypes.QueryDefinition) bool {
		return aws.ToString(v.QueryDefinitionId) == queryDefinitionID
	})
}

func findQueryDefinition(ctx context.Context, conn *cloudwatchlogs.Client, input *cloudwatchlogs.DescribeQueryDefinitionsInput, filter tfslices.Predicate[*awstypes.QueryDefinition]) (*awstypes.QueryDefinition, error) {
	output, err := findQueryDefinitions(ctx, conn, input, filter, tfslices.WithReturnFirstMatch)

	if err != nil {
		return nil, err
	}

	return tfresource.AssertSingleValueResult(output)
}

func findQueryDefinitions(ctx context.Context, conn *cloudwatchlogs.Client, input *cloudwatchlogs.DescribeQueryDefinitionsInput, filter tfslices.Predicate[*awstypes.QueryDefinition], optFns ...tfslices.FinderOptionsFunc) ([]awstypes.QueryDefinition, error) {
	var output []awstypes.QueryDefinition
	opts := tfslices.NewFinderOptions(optFns)

	err := describeQueryDefinitionsPages(ctx, conn, input, func(page *cloudwatchlogs.DescribeQueryDefinitionsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, v := range page.QueryDefinitions {
			if filter(&v) {
				output = append(output, v)
				if opts.ReturnFirstMatch() {
					return false
				}
			}
		}

		return !lastPage
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}
