// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package DatasetIamBindingCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DatasetIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression string `pulumi:"expression"`
	Title string `pulumi:"title"`
}

type DatasetIamBindingConditionInput interface {
	pulumi.Input

	ToDatasetIamBindingConditionOutput() DatasetIamBindingConditionOutput
	ToDatasetIamBindingConditionOutputWithContext(context.Context) DatasetIamBindingConditionOutput
}

type DatasetIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression pulumi.StringInput `pulumi:"expression"`
	Title pulumi.StringInput `pulumi:"title"`
}

func (DatasetIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamBindingCondition)(nil)).Elem()
}

func (i DatasetIamBindingConditionArgs) ToDatasetIamBindingConditionOutput() DatasetIamBindingConditionOutput {
	return i.ToDatasetIamBindingConditionOutputWithContext(context.Background())
}

func (i DatasetIamBindingConditionArgs) ToDatasetIamBindingConditionOutputWithContext(ctx context.Context) DatasetIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamBindingConditionOutput)
}

func (i DatasetIamBindingConditionArgs) ToDatasetIamBindingConditionPtrOutput() DatasetIamBindingConditionPtrOutput {
	return i.ToDatasetIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i DatasetIamBindingConditionArgs) ToDatasetIamBindingConditionPtrOutputWithContext(ctx context.Context) DatasetIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamBindingConditionOutput).ToDatasetIamBindingConditionPtrOutputWithContext(ctx)
}

type DatasetIamBindingConditionPtrInput interface {
	pulumi.Input

	ToDatasetIamBindingConditionPtrOutput() DatasetIamBindingConditionPtrOutput
	ToDatasetIamBindingConditionPtrOutputWithContext(context.Context) DatasetIamBindingConditionPtrOutput
}

type datasetIamBindingConditionPtrType DatasetIamBindingConditionArgs

func DatasetIamBindingConditionPtr(v *DatasetIamBindingConditionArgs) DatasetIamBindingConditionPtrInput {	return (*datasetIamBindingConditionPtrType)(v)
}

func (*datasetIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamBindingCondition)(nil)).Elem()
}

func (i *datasetIamBindingConditionPtrType) ToDatasetIamBindingConditionPtrOutput() DatasetIamBindingConditionPtrOutput {
	return i.ToDatasetIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *datasetIamBindingConditionPtrType) ToDatasetIamBindingConditionPtrOutputWithContext(ctx context.Context) DatasetIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamBindingConditionPtrOutput)
}

type DatasetIamBindingConditionOutput struct { *pulumi.OutputState }

func (DatasetIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamBindingCondition)(nil)).Elem()
}

func (o DatasetIamBindingConditionOutput) ToDatasetIamBindingConditionOutput() DatasetIamBindingConditionOutput {
	return o
}

func (o DatasetIamBindingConditionOutput) ToDatasetIamBindingConditionOutputWithContext(ctx context.Context) DatasetIamBindingConditionOutput {
	return o
}

func (o DatasetIamBindingConditionOutput) ToDatasetIamBindingConditionPtrOutput() DatasetIamBindingConditionPtrOutput {
	return o.ToDatasetIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o DatasetIamBindingConditionOutput) ToDatasetIamBindingConditionPtrOutputWithContext(ctx context.Context) DatasetIamBindingConditionPtrOutput {
	return o.ApplyT(func(v DatasetIamBindingCondition) *DatasetIamBindingCondition {
		return &v
	}).(DatasetIamBindingConditionPtrOutput)
}
func (o DatasetIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DatasetIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatasetIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v DatasetIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o DatasetIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v DatasetIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type DatasetIamBindingConditionPtrOutput struct { *pulumi.OutputState}

func (DatasetIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamBindingCondition)(nil)).Elem()
}

func (o DatasetIamBindingConditionPtrOutput) ToDatasetIamBindingConditionPtrOutput() DatasetIamBindingConditionPtrOutput {
	return o
}

func (o DatasetIamBindingConditionPtrOutput) ToDatasetIamBindingConditionPtrOutputWithContext(ctx context.Context) DatasetIamBindingConditionPtrOutput {
	return o
}

func (o DatasetIamBindingConditionPtrOutput) Elem() DatasetIamBindingConditionOutput {
	return o.ApplyT(func (v *DatasetIamBindingCondition) DatasetIamBindingCondition { return *v }).(DatasetIamBindingConditionOutput)
}

func (o DatasetIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DatasetIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatasetIamBindingConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v DatasetIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o DatasetIamBindingConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v DatasetIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetIamBindingConditionOutput{})
	pulumi.RegisterOutputType(DatasetIamBindingConditionPtrOutput{})
}
