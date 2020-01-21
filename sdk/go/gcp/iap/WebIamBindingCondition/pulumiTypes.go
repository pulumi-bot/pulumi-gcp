// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package WebIamBindingCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type WebIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression string `pulumi:"expression"`
	Title string `pulumi:"title"`
}

type WebIamBindingConditionInput interface {
	pulumi.Input

	ToWebIamBindingConditionOutput() WebIamBindingConditionOutput
	ToWebIamBindingConditionOutputWithContext(context.Context) WebIamBindingConditionOutput
}

type WebIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression pulumi.StringInput `pulumi:"expression"`
	Title pulumi.StringInput `pulumi:"title"`
}

func (WebIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamBindingCondition)(nil)).Elem()
}

func (i WebIamBindingConditionArgs) ToWebIamBindingConditionOutput() WebIamBindingConditionOutput {
	return i.ToWebIamBindingConditionOutputWithContext(context.Background())
}

func (i WebIamBindingConditionArgs) ToWebIamBindingConditionOutputWithContext(ctx context.Context) WebIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingConditionOutput)
}

func (i WebIamBindingConditionArgs) ToWebIamBindingConditionPtrOutput() WebIamBindingConditionPtrOutput {
	return i.ToWebIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i WebIamBindingConditionArgs) ToWebIamBindingConditionPtrOutputWithContext(ctx context.Context) WebIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingConditionOutput).ToWebIamBindingConditionPtrOutputWithContext(ctx)
}

type WebIamBindingConditionPtrInput interface {
	pulumi.Input

	ToWebIamBindingConditionPtrOutput() WebIamBindingConditionPtrOutput
	ToWebIamBindingConditionPtrOutputWithContext(context.Context) WebIamBindingConditionPtrOutput
}

type webIamBindingConditionPtrType WebIamBindingConditionArgs

func WebIamBindingConditionPtr(v *WebIamBindingConditionArgs) WebIamBindingConditionPtrInput {	return (*webIamBindingConditionPtrType)(v)
}

func (*webIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebIamBindingCondition)(nil)).Elem()
}

func (i *webIamBindingConditionPtrType) ToWebIamBindingConditionPtrOutput() WebIamBindingConditionPtrOutput {
	return i.ToWebIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *webIamBindingConditionPtrType) ToWebIamBindingConditionPtrOutputWithContext(ctx context.Context) WebIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingConditionPtrOutput)
}

type WebIamBindingConditionOutput struct { *pulumi.OutputState }

func (WebIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamBindingCondition)(nil)).Elem()
}

func (o WebIamBindingConditionOutput) ToWebIamBindingConditionOutput() WebIamBindingConditionOutput {
	return o
}

func (o WebIamBindingConditionOutput) ToWebIamBindingConditionOutputWithContext(ctx context.Context) WebIamBindingConditionOutput {
	return o
}

func (o WebIamBindingConditionOutput) ToWebIamBindingConditionPtrOutput() WebIamBindingConditionPtrOutput {
	return o.ToWebIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o WebIamBindingConditionOutput) ToWebIamBindingConditionPtrOutputWithContext(ctx context.Context) WebIamBindingConditionPtrOutput {
	return o.ApplyT(func(v WebIamBindingCondition) *WebIamBindingCondition {
		return &v
	}).(WebIamBindingConditionPtrOutput)
}
func (o WebIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v WebIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o WebIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v WebIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type WebIamBindingConditionPtrOutput struct { *pulumi.OutputState}

func (WebIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebIamBindingCondition)(nil)).Elem()
}

func (o WebIamBindingConditionPtrOutput) ToWebIamBindingConditionPtrOutput() WebIamBindingConditionPtrOutput {
	return o
}

func (o WebIamBindingConditionPtrOutput) ToWebIamBindingConditionPtrOutputWithContext(ctx context.Context) WebIamBindingConditionPtrOutput {
	return o
}

func (o WebIamBindingConditionPtrOutput) Elem() WebIamBindingConditionOutput {
	return o.ApplyT(func (v *WebIamBindingCondition) WebIamBindingCondition { return *v }).(WebIamBindingConditionOutput)
}

func (o WebIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebIamBindingConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v WebIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o WebIamBindingConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v WebIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebIamBindingConditionOutput{})
	pulumi.RegisterOutputType(WebIamBindingConditionPtrOutput{})
}
