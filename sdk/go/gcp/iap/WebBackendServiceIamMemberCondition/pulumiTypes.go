// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package WebBackendServiceIamMemberCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type WebBackendServiceIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression string `pulumi:"expression"`
	Title string `pulumi:"title"`
}

type WebBackendServiceIamMemberConditionInput interface {
	pulumi.Input

	ToWebBackendServiceIamMemberConditionOutput() WebBackendServiceIamMemberConditionOutput
	ToWebBackendServiceIamMemberConditionOutputWithContext(context.Context) WebBackendServiceIamMemberConditionOutput
}

type WebBackendServiceIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression pulumi.StringInput `pulumi:"expression"`
	Title pulumi.StringInput `pulumi:"title"`
}

func (WebBackendServiceIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamMemberCondition)(nil)).Elem()
}

func (i WebBackendServiceIamMemberConditionArgs) ToWebBackendServiceIamMemberConditionOutput() WebBackendServiceIamMemberConditionOutput {
	return i.ToWebBackendServiceIamMemberConditionOutputWithContext(context.Background())
}

func (i WebBackendServiceIamMemberConditionArgs) ToWebBackendServiceIamMemberConditionOutputWithContext(ctx context.Context) WebBackendServiceIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberConditionOutput)
}

func (i WebBackendServiceIamMemberConditionArgs) ToWebBackendServiceIamMemberConditionPtrOutput() WebBackendServiceIamMemberConditionPtrOutput {
	return i.ToWebBackendServiceIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i WebBackendServiceIamMemberConditionArgs) ToWebBackendServiceIamMemberConditionPtrOutputWithContext(ctx context.Context) WebBackendServiceIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberConditionOutput).ToWebBackendServiceIamMemberConditionPtrOutputWithContext(ctx)
}

type WebBackendServiceIamMemberConditionPtrInput interface {
	pulumi.Input

	ToWebBackendServiceIamMemberConditionPtrOutput() WebBackendServiceIamMemberConditionPtrOutput
	ToWebBackendServiceIamMemberConditionPtrOutputWithContext(context.Context) WebBackendServiceIamMemberConditionPtrOutput
}

type webBackendServiceIamMemberConditionPtrType WebBackendServiceIamMemberConditionArgs

func WebBackendServiceIamMemberConditionPtr(v *WebBackendServiceIamMemberConditionArgs) WebBackendServiceIamMemberConditionPtrInput {	return (*webBackendServiceIamMemberConditionPtrType)(v)
}

func (*webBackendServiceIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebBackendServiceIamMemberCondition)(nil)).Elem()
}

func (i *webBackendServiceIamMemberConditionPtrType) ToWebBackendServiceIamMemberConditionPtrOutput() WebBackendServiceIamMemberConditionPtrOutput {
	return i.ToWebBackendServiceIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *webBackendServiceIamMemberConditionPtrType) ToWebBackendServiceIamMemberConditionPtrOutputWithContext(ctx context.Context) WebBackendServiceIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberConditionPtrOutput)
}

type WebBackendServiceIamMemberConditionOutput struct { *pulumi.OutputState }

func (WebBackendServiceIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamMemberCondition)(nil)).Elem()
}

func (o WebBackendServiceIamMemberConditionOutput) ToWebBackendServiceIamMemberConditionOutput() WebBackendServiceIamMemberConditionOutput {
	return o
}

func (o WebBackendServiceIamMemberConditionOutput) ToWebBackendServiceIamMemberConditionOutputWithContext(ctx context.Context) WebBackendServiceIamMemberConditionOutput {
	return o
}

func (o WebBackendServiceIamMemberConditionOutput) ToWebBackendServiceIamMemberConditionPtrOutput() WebBackendServiceIamMemberConditionPtrOutput {
	return o.ToWebBackendServiceIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o WebBackendServiceIamMemberConditionOutput) ToWebBackendServiceIamMemberConditionPtrOutputWithContext(ctx context.Context) WebBackendServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v WebBackendServiceIamMemberCondition) *WebBackendServiceIamMemberCondition {
		return &v
	}).(WebBackendServiceIamMemberConditionPtrOutput)
}
func (o WebBackendServiceIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebBackendServiceIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebBackendServiceIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v WebBackendServiceIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o WebBackendServiceIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v WebBackendServiceIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type WebBackendServiceIamMemberConditionPtrOutput struct { *pulumi.OutputState}

func (WebBackendServiceIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebBackendServiceIamMemberCondition)(nil)).Elem()
}

func (o WebBackendServiceIamMemberConditionPtrOutput) ToWebBackendServiceIamMemberConditionPtrOutput() WebBackendServiceIamMemberConditionPtrOutput {
	return o
}

func (o WebBackendServiceIamMemberConditionPtrOutput) ToWebBackendServiceIamMemberConditionPtrOutputWithContext(ctx context.Context) WebBackendServiceIamMemberConditionPtrOutput {
	return o
}

func (o WebBackendServiceIamMemberConditionPtrOutput) Elem() WebBackendServiceIamMemberConditionOutput {
	return o.ApplyT(func (v *WebBackendServiceIamMemberCondition) WebBackendServiceIamMemberCondition { return *v }).(WebBackendServiceIamMemberConditionOutput)
}

func (o WebBackendServiceIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebBackendServiceIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebBackendServiceIamMemberConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v WebBackendServiceIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o WebBackendServiceIamMemberConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v WebBackendServiceIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebBackendServiceIamMemberConditionOutput{})
	pulumi.RegisterOutputType(WebBackendServiceIamMemberConditionPtrOutput{})
}
