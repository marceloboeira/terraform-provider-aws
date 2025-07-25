// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
)

func TestListNestedObjectTypeOfEqual(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	testCases := map[string]struct {
		other attr.Type
		want  bool
	}{
		"string type": {
			other: types.StringType,
		},
		"equal type": {
			other: fwtypes.NewListNestedObjectTypeOf[ObjectA](ctx),
			want:  true,
		},
		"other struct type": {
			other: fwtypes.NewListNestedObjectTypeOf[ObjectB](ctx),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fwtypes.NewListNestedObjectTypeOf[ObjectA](ctx).Equal(testCase.other)

			if got != testCase.want {
				t.Errorf("got = %v, want = %v", got, testCase.want)
			}
		})
	}
}

func TestListNestedObjectTypeOfValueFromTerraform(t *testing.T) {
	t.Parallel()

	objectA := ObjectA{
		Name: types.StringValue("test"),
	}
	objectAType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"name": tftypes.String,
		},
	}
	objectAListType := tftypes.List{ElementType: objectAType}
	objectAValue := tftypes.NewValue(objectAType, map[string]tftypes.Value{
		"name": tftypes.NewValue(tftypes.String, "test"),
	})
	objectAListValue := tftypes.NewValue(tftypes.List{ElementType: objectAType}, []tftypes.Value{objectAValue})
	objectBType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"length": tftypes.Number,
		},
	}
	objectBValue := tftypes.NewValue(objectBType, map[string]tftypes.Value{
		"length": tftypes.NewValue(tftypes.Number, 42),
	})
	objectBListValue := tftypes.NewValue(tftypes.List{ElementType: objectBType}, []tftypes.Value{objectBValue})

	ctx := context.Background()
	testCases := map[string]struct {
		tfVal   tftypes.Value
		wantVal attr.Value
		wantErr bool
	}{
		"null value": {
			tfVal:   tftypes.NewValue(objectAListType, nil),
			wantVal: fwtypes.NewListNestedObjectValueOfNull[ObjectA](ctx),
		},
		"unknown value": {
			tfVal:   tftypes.NewValue(objectAListType, tftypes.UnknownValue),
			wantVal: fwtypes.NewListNestedObjectValueOfUnknown[ObjectA](ctx),
		},
		"valid value": {
			tfVal:   objectAListValue,
			wantVal: fwtypes.NewListNestedObjectValueOfPtrMust[ObjectA](ctx, &objectA),
		},
		"invalid Terraform value": {
			tfVal:   objectBListValue,
			wantVal: fwtypes.NewListNestedObjectValueOfPtrMust[ObjectA](ctx, &objectA),
			wantErr: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotVal, err := fwtypes.NewListNestedObjectTypeOf[ObjectA](ctx).ValueFromTerraform(ctx, testCase.tfVal)
			gotErr := err != nil

			if gotErr != testCase.wantErr {
				t.Errorf("gotErr = %v, wantErr = %v", gotErr, testCase.wantErr)
			}

			if gotErr {
				if !testCase.wantErr {
					t.Errorf("err = %q", err)
				}
			} else if diff := cmp.Diff(gotVal, testCase.wantVal); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestListNestedObjectValueOfEqual(t *testing.T) {
	t.Parallel()

	objectA := ObjectA{
		Name: types.StringValue("test"),
	}
	objectB := ObjectB{
		Length: types.Int64Value(42),
	}
	objectA2 := ObjectA{
		Name: types.StringValue("test2"),
	}

	ctx := context.Background()
	testCases := map[string]struct {
		other attr.Value
		want  bool
	}{
		"string value": {
			other: types.StringValue("test"),
		},
		"equal value": {
			other: fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectA),
			want:  true,
		},
		"struct not equal value": {
			other: fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectA2),
		},
		"other struct value": {
			other: fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectB),
		},
		"null value": {
			other: fwtypes.NewListNestedObjectValueOfNull[ObjectA](ctx),
		},
		"unknown value": {
			other: fwtypes.NewListNestedObjectValueOfUnknown[ObjectA](ctx),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectA).Equal(testCase.other)

			if got != testCase.want {
				t.Errorf("got = %v, want = %v", got, testCase.want)
			}
		})
	}
}

func TestListNestedObjectValueOfListSemanticEquals(t *testing.T) {
	t.Parallel()

	semanticallyEqual := func(ctx context.Context, a, b fwtypes.NestedCollectionValue[ObjectA]) (bool, diag.Diagnostics) {
		var diags diag.Diagnostics

		if a.Equal(b) {
			return true, diags
		}

		aSlice, d := b.ToSlice(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return false, diags
		}

		bSlice, d := b.ToSlice(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return false, diags
		}

		if a.IsNull() && len(bSlice) == 0 {
			return true, diags
		}

		if b.IsNull() && len(aSlice) == 0 {
			return true, diags
		}

		return false, diags
	}

	// test artifacts
	objectA := ObjectA{
		Name: types.StringValue("test"),
	}
	objectB := ObjectA{
		Name: types.StringValue("test2"),
	}

	emptySlice := make([]*ObjectA, 0)

	ctx := context.Background()
	testCases := map[string]struct {
		current fwtypes.ListNestedObjectValueOf[ObjectA]
		other   basetypes.ListValuable
		want    bool
	}{
		"equal value": {
			current: fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectA, fwtypes.WithSemanticEqualityFunc(semanticallyEqual)),
			other:   fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectA),
			want:    true,
		},
		"equal nil current and empty slice": {
			current: fwtypes.NewListNestedObjectValueOfNull(ctx, fwtypes.WithSemanticEqualityFunc(semanticallyEqual)),
			other:   fwtypes.NewListNestedObjectValueOfSliceMust(ctx, emptySlice),
			want:    true,
		},
		"equal empty current and nil slice": {
			current: fwtypes.NewListNestedObjectValueOfSliceMust(ctx, emptySlice, fwtypes.WithSemanticEqualityFunc(semanticallyEqual)),
			other:   fwtypes.NewListNestedObjectValueOfNull[ObjectA](ctx),
			want:    true,
		},
		"not equal": {
			current: fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectA, fwtypes.WithSemanticEqualityFunc(semanticallyEqual)),
			other:   fwtypes.NewListNestedObjectValueOfPtrMust(ctx, &objectB),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, _ := testCase.current.ListSemanticEquals(ctx, testCase.other)

			if got != testCase.want {
				t.Errorf("got = %v, want = %v", got, testCase.want)
			}
		})
	}
}
