// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RepositoryIamMemberCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RepositoryIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression string `pulumi:"expression"`
	Title string `pulumi:"title"`
}

type RepositoryIamMemberConditionInput interface {
	pulumi.Input

	ToRepositoryIamMemberConditionOutput() RepositoryIamMemberConditionOutput
	ToRepositoryIamMemberConditionOutputWithContext(context.Context) RepositoryIamMemberConditionOutput
}

type RepositoryIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression pulumi.StringInput `pulumi:"expression"`
	Title pulumi.StringInput `pulumi:"title"`
}

func (RepositoryIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamMemberCondition)(nil)).Elem()
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionOutput() RepositoryIamMemberConditionOutput {
	return i.ToRepositoryIamMemberConditionOutputWithContext(context.Background())
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionOutputWithContext(ctx context.Context) RepositoryIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberConditionOutput)
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return i.ToRepositoryIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberConditionOutput).ToRepositoryIamMemberConditionPtrOutputWithContext(ctx)
}

type RepositoryIamMemberConditionPtrInput interface {
	pulumi.Input

	ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput
	ToRepositoryIamMemberConditionPtrOutputWithContext(context.Context) RepositoryIamMemberConditionPtrOutput
}

type repositoryIamMemberConditionPtrType RepositoryIamMemberConditionArgs

func RepositoryIamMemberConditionPtr(v *RepositoryIamMemberConditionArgs) RepositoryIamMemberConditionPtrInput {	return (*repositoryIamMemberConditionPtrType)(v)
}

func (*repositoryIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamMemberCondition)(nil)).Elem()
}

func (i *repositoryIamMemberConditionPtrType) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return i.ToRepositoryIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *repositoryIamMemberConditionPtrType) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberConditionPtrOutput)
}

type RepositoryIamMemberConditionOutput struct { *pulumi.OutputState }

func (RepositoryIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamMemberCondition)(nil)).Elem()
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionOutput() RepositoryIamMemberConditionOutput {
	return o
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionOutputWithContext(ctx context.Context) RepositoryIamMemberConditionOutput {
	return o
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return o.ToRepositoryIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return o.ApplyT(func(v RepositoryIamMemberCondition) *RepositoryIamMemberCondition {
		return &v
	}).(RepositoryIamMemberConditionPtrOutput)
}
func (o RepositoryIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RepositoryIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RepositoryIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v RepositoryIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o RepositoryIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v RepositoryIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type RepositoryIamMemberConditionPtrOutput struct { *pulumi.OutputState}

func (RepositoryIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamMemberCondition)(nil)).Elem()
}

func (o RepositoryIamMemberConditionPtrOutput) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return o
}

func (o RepositoryIamMemberConditionPtrOutput) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return o
}

func (o RepositoryIamMemberConditionPtrOutput) Elem() RepositoryIamMemberConditionOutput {
	return o.ApplyT(func (v *RepositoryIamMemberCondition) RepositoryIamMemberCondition { return *v }).(RepositoryIamMemberConditionOutput)
}

func (o RepositoryIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RepositoryIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RepositoryIamMemberConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v RepositoryIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o RepositoryIamMemberConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v RepositoryIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryIamMemberConditionOutput{})
	pulumi.RegisterOutputType(RepositoryIamMemberConditionPtrOutput{})
}
