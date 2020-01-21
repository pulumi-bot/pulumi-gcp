// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package AttestorIamBindingCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AttestorIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression string `pulumi:"expression"`
	Title string `pulumi:"title"`
}

type AttestorIamBindingConditionInput interface {
	pulumi.Input

	ToAttestorIamBindingConditionOutput() AttestorIamBindingConditionOutput
	ToAttestorIamBindingConditionOutputWithContext(context.Context) AttestorIamBindingConditionOutput
}

type AttestorIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression pulumi.StringInput `pulumi:"expression"`
	Title pulumi.StringInput `pulumi:"title"`
}

func (AttestorIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamBindingCondition)(nil)).Elem()
}

func (i AttestorIamBindingConditionArgs) ToAttestorIamBindingConditionOutput() AttestorIamBindingConditionOutput {
	return i.ToAttestorIamBindingConditionOutputWithContext(context.Background())
}

func (i AttestorIamBindingConditionArgs) ToAttestorIamBindingConditionOutputWithContext(ctx context.Context) AttestorIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamBindingConditionOutput)
}

func (i AttestorIamBindingConditionArgs) ToAttestorIamBindingConditionPtrOutput() AttestorIamBindingConditionPtrOutput {
	return i.ToAttestorIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i AttestorIamBindingConditionArgs) ToAttestorIamBindingConditionPtrOutputWithContext(ctx context.Context) AttestorIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamBindingConditionOutput).ToAttestorIamBindingConditionPtrOutputWithContext(ctx)
}

type AttestorIamBindingConditionPtrInput interface {
	pulumi.Input

	ToAttestorIamBindingConditionPtrOutput() AttestorIamBindingConditionPtrOutput
	ToAttestorIamBindingConditionPtrOutputWithContext(context.Context) AttestorIamBindingConditionPtrOutput
}

type attestorIamBindingConditionPtrType AttestorIamBindingConditionArgs

func AttestorIamBindingConditionPtr(v *AttestorIamBindingConditionArgs) AttestorIamBindingConditionPtrInput {	return (*attestorIamBindingConditionPtrType)(v)
}

func (*attestorIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestorIamBindingCondition)(nil)).Elem()
}

func (i *attestorIamBindingConditionPtrType) ToAttestorIamBindingConditionPtrOutput() AttestorIamBindingConditionPtrOutput {
	return i.ToAttestorIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *attestorIamBindingConditionPtrType) ToAttestorIamBindingConditionPtrOutputWithContext(ctx context.Context) AttestorIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamBindingConditionPtrOutput)
}

type AttestorIamBindingConditionOutput struct { *pulumi.OutputState }

func (AttestorIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamBindingCondition)(nil)).Elem()
}

func (o AttestorIamBindingConditionOutput) ToAttestorIamBindingConditionOutput() AttestorIamBindingConditionOutput {
	return o
}

func (o AttestorIamBindingConditionOutput) ToAttestorIamBindingConditionOutputWithContext(ctx context.Context) AttestorIamBindingConditionOutput {
	return o
}

func (o AttestorIamBindingConditionOutput) ToAttestorIamBindingConditionPtrOutput() AttestorIamBindingConditionPtrOutput {
	return o.ToAttestorIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o AttestorIamBindingConditionOutput) ToAttestorIamBindingConditionPtrOutputWithContext(ctx context.Context) AttestorIamBindingConditionPtrOutput {
	return o.ApplyT(func(v AttestorIamBindingCondition) *AttestorIamBindingCondition {
		return &v
	}).(AttestorIamBindingConditionPtrOutput)
}
func (o AttestorIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v AttestorIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AttestorIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o AttestorIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type AttestorIamBindingConditionPtrOutput struct { *pulumi.OutputState}

func (AttestorIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestorIamBindingCondition)(nil)).Elem()
}

func (o AttestorIamBindingConditionPtrOutput) ToAttestorIamBindingConditionPtrOutput() AttestorIamBindingConditionPtrOutput {
	return o
}

func (o AttestorIamBindingConditionPtrOutput) ToAttestorIamBindingConditionPtrOutputWithContext(ctx context.Context) AttestorIamBindingConditionPtrOutput {
	return o
}

func (o AttestorIamBindingConditionPtrOutput) Elem() AttestorIamBindingConditionOutput {
	return o.ApplyT(func (v *AttestorIamBindingCondition) AttestorIamBindingCondition { return *v }).(AttestorIamBindingConditionOutput)
}

func (o AttestorIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v AttestorIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AttestorIamBindingConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o AttestorIamBindingConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestorIamBindingConditionOutput{})
	pulumi.RegisterOutputType(AttestorIamBindingConditionPtrOutput{})
}
