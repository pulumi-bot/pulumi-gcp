// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ConfigIamMemberCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ConfigIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression string `pulumi:"expression"`
	Title string `pulumi:"title"`
}

type ConfigIamMemberConditionInput interface {
	pulumi.Input

	ToConfigIamMemberConditionOutput() ConfigIamMemberConditionOutput
	ToConfigIamMemberConditionOutputWithContext(context.Context) ConfigIamMemberConditionOutput
}

type ConfigIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression pulumi.StringInput `pulumi:"expression"`
	Title pulumi.StringInput `pulumi:"title"`
}

func (ConfigIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigIamMemberCondition)(nil)).Elem()
}

func (i ConfigIamMemberConditionArgs) ToConfigIamMemberConditionOutput() ConfigIamMemberConditionOutput {
	return i.ToConfigIamMemberConditionOutputWithContext(context.Background())
}

func (i ConfigIamMemberConditionArgs) ToConfigIamMemberConditionOutputWithContext(ctx context.Context) ConfigIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamMemberConditionOutput)
}

func (i ConfigIamMemberConditionArgs) ToConfigIamMemberConditionPtrOutput() ConfigIamMemberConditionPtrOutput {
	return i.ToConfigIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i ConfigIamMemberConditionArgs) ToConfigIamMemberConditionPtrOutputWithContext(ctx context.Context) ConfigIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamMemberConditionOutput).ToConfigIamMemberConditionPtrOutputWithContext(ctx)
}

type ConfigIamMemberConditionPtrInput interface {
	pulumi.Input

	ToConfigIamMemberConditionPtrOutput() ConfigIamMemberConditionPtrOutput
	ToConfigIamMemberConditionPtrOutputWithContext(context.Context) ConfigIamMemberConditionPtrOutput
}

type configIamMemberConditionPtrType ConfigIamMemberConditionArgs

func ConfigIamMemberConditionPtr(v *ConfigIamMemberConditionArgs) ConfigIamMemberConditionPtrInput {	return (*configIamMemberConditionPtrType)(v)
}

func (*configIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigIamMemberCondition)(nil)).Elem()
}

func (i *configIamMemberConditionPtrType) ToConfigIamMemberConditionPtrOutput() ConfigIamMemberConditionPtrOutput {
	return i.ToConfigIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *configIamMemberConditionPtrType) ToConfigIamMemberConditionPtrOutputWithContext(ctx context.Context) ConfigIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamMemberConditionPtrOutput)
}

type ConfigIamMemberConditionOutput struct { *pulumi.OutputState }

func (ConfigIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigIamMemberCondition)(nil)).Elem()
}

func (o ConfigIamMemberConditionOutput) ToConfigIamMemberConditionOutput() ConfigIamMemberConditionOutput {
	return o
}

func (o ConfigIamMemberConditionOutput) ToConfigIamMemberConditionOutputWithContext(ctx context.Context) ConfigIamMemberConditionOutput {
	return o
}

func (o ConfigIamMemberConditionOutput) ToConfigIamMemberConditionPtrOutput() ConfigIamMemberConditionPtrOutput {
	return o.ToConfigIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o ConfigIamMemberConditionOutput) ToConfigIamMemberConditionPtrOutputWithContext(ctx context.Context) ConfigIamMemberConditionPtrOutput {
	return o.ApplyT(func(v ConfigIamMemberCondition) *ConfigIamMemberCondition {
		return &v
	}).(ConfigIamMemberConditionPtrOutput)
}
func (o ConfigIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v ConfigIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o ConfigIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v ConfigIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type ConfigIamMemberConditionPtrOutput struct { *pulumi.OutputState}

func (ConfigIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigIamMemberCondition)(nil)).Elem()
}

func (o ConfigIamMemberConditionPtrOutput) ToConfigIamMemberConditionPtrOutput() ConfigIamMemberConditionPtrOutput {
	return o
}

func (o ConfigIamMemberConditionPtrOutput) ToConfigIamMemberConditionPtrOutputWithContext(ctx context.Context) ConfigIamMemberConditionPtrOutput {
	return o
}

func (o ConfigIamMemberConditionPtrOutput) Elem() ConfigIamMemberConditionOutput {
	return o.ApplyT(func (v *ConfigIamMemberCondition) ConfigIamMemberCondition { return *v }).(ConfigIamMemberConditionOutput)
}

func (o ConfigIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigIamMemberConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func (v ConfigIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o ConfigIamMemberConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func (v ConfigIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigIamMemberConditionOutput{})
	pulumi.RegisterOutputType(ConfigIamMemberConditionPtrOutput{})
}
