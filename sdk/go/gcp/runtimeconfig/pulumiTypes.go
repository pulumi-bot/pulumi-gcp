// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ConfigIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// ConfigIamBindingConditionInput is an input type that accepts ConfigIamBindingConditionArgs and ConfigIamBindingConditionOutput values.
// You can construct a concrete instance of `ConfigIamBindingConditionInput` via:
//
//          ConfigIamBindingConditionArgs{...}
type ConfigIamBindingConditionInput interface {
	pulumi.Input

	ToConfigIamBindingConditionOutput() ConfigIamBindingConditionOutput
	ToConfigIamBindingConditionOutputWithContext(context.Context) ConfigIamBindingConditionOutput
}

type ConfigIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (ConfigIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigIamBindingCondition)(nil)).Elem()
}

func (i ConfigIamBindingConditionArgs) ToConfigIamBindingConditionOutput() ConfigIamBindingConditionOutput {
	return i.ToConfigIamBindingConditionOutputWithContext(context.Background())
}

func (i ConfigIamBindingConditionArgs) ToConfigIamBindingConditionOutputWithContext(ctx context.Context) ConfigIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamBindingConditionOutput)
}

func (i ConfigIamBindingConditionArgs) ToConfigIamBindingConditionPtrOutput() ConfigIamBindingConditionPtrOutput {
	return i.ToConfigIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i ConfigIamBindingConditionArgs) ToConfigIamBindingConditionPtrOutputWithContext(ctx context.Context) ConfigIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamBindingConditionOutput).ToConfigIamBindingConditionPtrOutput()
}

// ConfigIamBindingConditionPtrInput is an input type that accepts ConfigIamBindingConditionArgs, ConfigIamBindingConditionPtr and ConfigIamBindingConditionPtrOutput values.
// You can construct a concrete instance of `ConfigIamBindingConditionPtrInput` via:
//
//          ConfigIamBindingConditionArgs{...}
//
//  or:
//
//          nil
type ConfigIamBindingConditionPtrInput interface {
	pulumi.Input

	ToConfigIamBindingConditionPtrOutput() ConfigIamBindingConditionPtrOutput
	ToConfigIamBindingConditionPtrOutputWithContext(context.Context) ConfigIamBindingConditionPtrOutput
}

type configIamBindingConditionPtrType ConfigIamBindingConditionArgs

func ConfigIamBindingConditionPtr(v *ConfigIamBindingConditionArgs) ConfigIamBindingConditionPtrInput {
	return (*configIamBindingConditionPtrType)(v)
}

func (*configIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigIamBindingCondition)(nil)).Elem()
}

func (i *configIamBindingConditionPtrType) ToConfigIamBindingConditionPtrOutput() ConfigIamBindingConditionPtrOutput {
	return i.ToConfigIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *configIamBindingConditionPtrType) ToConfigIamBindingConditionPtrOutputWithContext(ctx context.Context) ConfigIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamBindingConditionOutput).ToConfigIamBindingConditionPtrOutput()
}

type ConfigIamBindingConditionOutput struct{ *pulumi.OutputState }

func (ConfigIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigIamBindingCondition)(nil)).Elem()
}

func (o ConfigIamBindingConditionOutput) ToConfigIamBindingConditionOutput() ConfigIamBindingConditionOutput {
	return o
}

func (o ConfigIamBindingConditionOutput) ToConfigIamBindingConditionOutputWithContext(ctx context.Context) ConfigIamBindingConditionOutput {
	return o
}

func (o ConfigIamBindingConditionOutput) ToConfigIamBindingConditionPtrOutput() ConfigIamBindingConditionPtrOutput {
	return o.ToConfigIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o ConfigIamBindingConditionOutput) ToConfigIamBindingConditionPtrOutputWithContext(ctx context.Context) ConfigIamBindingConditionPtrOutput {
	return o.ApplyT(func(v ConfigIamBindingCondition) *ConfigIamBindingCondition {
		return &v
	}).(ConfigIamBindingConditionPtrOutput)
}
func (o ConfigIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o ConfigIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type ConfigIamBindingConditionPtrOutput struct{ *pulumi.OutputState }

func (ConfigIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigIamBindingCondition)(nil)).Elem()
}

func (o ConfigIamBindingConditionPtrOutput) ToConfigIamBindingConditionPtrOutput() ConfigIamBindingConditionPtrOutput {
	return o
}

func (o ConfigIamBindingConditionPtrOutput) ToConfigIamBindingConditionPtrOutputWithContext(ctx context.Context) ConfigIamBindingConditionPtrOutput {
	return o
}

func (o ConfigIamBindingConditionPtrOutput) Elem() ConfigIamBindingConditionOutput {
	return o.ApplyT(func(v *ConfigIamBindingCondition) ConfigIamBindingCondition { return *v }).(ConfigIamBindingConditionOutput)
}

func (o ConfigIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConfigIamBindingConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o ConfigIamBindingConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type ConfigIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// ConfigIamMemberConditionInput is an input type that accepts ConfigIamMemberConditionArgs and ConfigIamMemberConditionOutput values.
// You can construct a concrete instance of `ConfigIamMemberConditionInput` via:
//
//          ConfigIamMemberConditionArgs{...}
type ConfigIamMemberConditionInput interface {
	pulumi.Input

	ToConfigIamMemberConditionOutput() ConfigIamMemberConditionOutput
	ToConfigIamMemberConditionOutputWithContext(context.Context) ConfigIamMemberConditionOutput
}

type ConfigIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
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
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamMemberConditionOutput).ToConfigIamMemberConditionPtrOutput()
}

// ConfigIamMemberConditionPtrInput is an input type that accepts ConfigIamMemberConditionArgs, ConfigIamMemberConditionPtr and ConfigIamMemberConditionPtrOutput values.
// You can construct a concrete instance of `ConfigIamMemberConditionPtrInput` via:
//
//          ConfigIamMemberConditionArgs{...}
//
//  or:
//
//          nil
type ConfigIamMemberConditionPtrInput interface {
	pulumi.Input

	ToConfigIamMemberConditionPtrOutput() ConfigIamMemberConditionPtrOutput
	ToConfigIamMemberConditionPtrOutputWithContext(context.Context) ConfigIamMemberConditionPtrOutput
}

type configIamMemberConditionPtrType ConfigIamMemberConditionArgs

func ConfigIamMemberConditionPtr(v *ConfigIamMemberConditionArgs) ConfigIamMemberConditionPtrInput {
	return (*configIamMemberConditionPtrType)(v)
}

func (*configIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigIamMemberCondition)(nil)).Elem()
}

func (i *configIamMemberConditionPtrType) ToConfigIamMemberConditionPtrOutput() ConfigIamMemberConditionPtrOutput {
	return i.ToConfigIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *configIamMemberConditionPtrType) ToConfigIamMemberConditionPtrOutputWithContext(ctx context.Context) ConfigIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamMemberConditionOutput).ToConfigIamMemberConditionPtrOutput()
}

type ConfigIamMemberConditionOutput struct{ *pulumi.OutputState }

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
	return o.ApplyT(func(v ConfigIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o ConfigIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type ConfigIamMemberConditionPtrOutput struct{ *pulumi.OutputState }

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
	return o.ApplyT(func(v *ConfigIamMemberCondition) ConfigIamMemberCondition { return *v }).(ConfigIamMemberConditionOutput)
}

func (o ConfigIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConfigIamMemberConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o ConfigIamMemberConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigIamBindingConditionOutput{})
	pulumi.RegisterOutputType(ConfigIamBindingConditionPtrOutput{})
	pulumi.RegisterOutputType(ConfigIamMemberConditionOutput{})
	pulumi.RegisterOutputType(ConfigIamMemberConditionPtrOutput{})
}
