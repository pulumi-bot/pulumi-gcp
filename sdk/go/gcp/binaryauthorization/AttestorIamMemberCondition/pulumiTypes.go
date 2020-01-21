// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package AttestorIamMemberCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AttestorIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression string `pulumi:"expression"`
	Title string `pulumi:"title"`
}

type AttestorIamMemberConditionInput interface {
	pulumi.Input

	ToAttestorIamMemberConditionOutput() AttestorIamMemberConditionOutput
	ToAttestorIamMemberConditionOutputWithContext(context.Context) AttestorIamMemberConditionOutput
}

type AttestorIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression pulumi.StringInput `pulumi:"expression"`
	Title pulumi.StringInput `pulumi:"title"`
}

func (AttestorIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamMemberCondition)(nil)).Elem()
}

func (i AttestorIamMemberConditionArgs) ToAttestorIamMemberConditionOutput() AttestorIamMemberConditionOutput {
	return i.ToAttestorIamMemberConditionOutputWithContext(context.Background())
}

func (i AttestorIamMemberConditionArgs) ToAttestorIamMemberConditionOutputWithContext(ctx context.Context) AttestorIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamMemberConditionOutput)
}

func (i AttestorIamMemberConditionArgs) ToAttestorIamMemberConditionPtrOutput() AttestorIamMemberConditionPtrOutput {
	return i.ToAttestorIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i AttestorIamMemberConditionArgs) ToAttestorIamMemberConditionPtrOutputWithContext(ctx context.Context) AttestorIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamMemberConditionOutput).ToAttestorIamMemberConditionPtrOutputWithContext(ctx)
}

type AttestorIamMemberConditionPtrInput interface {
	pulumi.Input

	ToAttestorIamMemberConditionPtrOutput() AttestorIamMemberConditionPtrOutput
	ToAttestorIamMemberConditionPtrOutputWithContext(context.Context) AttestorIamMemberConditionPtrOutput
}

type attestorIamMemberConditionPtrType AttestorIamMemberConditionArgs

func AttestorIamMemberConditionPtr(v *AttestorIamMemberConditionArgs) AttestorIamMemberConditionPtrInput {	return (*attestorIamMemberConditionPtrType)(v)
}

func (*attestorIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestorIamMemberCondition)(nil)).Elem()
}

func (i *attestorIamMemberConditionPtrType) ToAttestorIamMemberConditionPtrOutput() AttestorIamMemberConditionPtrOutput {
	return i.ToAttestorIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *attestorIamMemberConditionPtrType) ToAttestorIamMemberConditionPtrOutputWithContext(ctx context.Context) AttestorIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamMemberConditionPtrOutput)
}

type AttestorIamMemberConditionOutput struct { *pulumi.OutputState }

func (AttestorIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamMemberCondition)(nil)).Elem()
}

func (o AttestorIamMemberConditionOutput) ToAttestorIamMemberConditionOutput() AttestorIamMemberConditionOutput {
	return o
}

func (o AttestorIamMemberConditionOutput) ToAttestorIamMemberConditionOutputWithContext(ctx context.Context) AttestorIamMemberConditionOutput {
	return o
}

func (o AttestorIamMemberConditionOutput) ToAttestorIamMemberConditionPtrOutput() AttestorIamMemberConditionPtrOutput {
	return o.ToAttestorIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o AttestorIamMemberConditionOutput) ToAttestorIamMemberConditionPtrOutputWithContext(ctx context.Context) AttestorIamMemberConditionPtrOutput {
	return o.ApplyT(func(v AttestorIamMemberCondition) *AttestorIamMemberCondition {
		return &v
	}).(AttestorIamMemberConditionPtrOutput)
}
func (o AttestorIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v AttestorIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AttestorIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o AttestorIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type AttestorIamMemberConditionPtrOutput struct { *pulumi.OutputState}

func (AttestorIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestorIamMemberCondition)(nil)).Elem()
}

func (o AttestorIamMemberConditionPtrOutput) ToAttestorIamMemberConditionPtrOutput() AttestorIamMemberConditionPtrOutput {
	return o
}

func (o AttestorIamMemberConditionPtrOutput) ToAttestorIamMemberConditionPtrOutputWithContext(ctx context.Context) AttestorIamMemberConditionPtrOutput {
	return o
}

func (o AttestorIamMemberConditionPtrOutput) Elem() AttestorIamMemberConditionOutput {
	return o.ApplyT(func (v *AttestorIamMemberCondition) AttestorIamMemberCondition { return *v }).(AttestorIamMemberConditionOutput)
}

func (o AttestorIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v AttestorIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AttestorIamMemberConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o AttestorIamMemberConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v AttestorIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestorIamMemberConditionOutput{})
	pulumi.RegisterOutputType(AttestorIamMemberConditionPtrOutput{})
}
