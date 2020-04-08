// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ml

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type EngineModelDefaultVersion struct {
	Name string `pulumi:"name"`
}

type EngineModelDefaultVersionInput interface {
	pulumi.Input

	ToEngineModelDefaultVersionOutput() EngineModelDefaultVersionOutput
	ToEngineModelDefaultVersionOutputWithContext(context.Context) EngineModelDefaultVersionOutput
}

type EngineModelDefaultVersionArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (EngineModelDefaultVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineModelDefaultVersion)(nil)).Elem()
}

func (i EngineModelDefaultVersionArgs) ToEngineModelDefaultVersionOutput() EngineModelDefaultVersionOutput {
	return i.ToEngineModelDefaultVersionOutputWithContext(context.Background())
}

func (i EngineModelDefaultVersionArgs) ToEngineModelDefaultVersionOutputWithContext(ctx context.Context) EngineModelDefaultVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineModelDefaultVersionOutput)
}

func (i EngineModelDefaultVersionArgs) ToEngineModelDefaultVersionPtrOutput() EngineModelDefaultVersionPtrOutput {
	return i.ToEngineModelDefaultVersionPtrOutputWithContext(context.Background())
}

func (i EngineModelDefaultVersionArgs) ToEngineModelDefaultVersionPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineModelDefaultVersionOutput).ToEngineModelDefaultVersionPtrOutputWithContext(ctx)
}

type EngineModelDefaultVersionPtrInput interface {
	pulumi.Input

	ToEngineModelDefaultVersionPtrOutput() EngineModelDefaultVersionPtrOutput
	ToEngineModelDefaultVersionPtrOutputWithContext(context.Context) EngineModelDefaultVersionPtrOutput
}

type engineModelDefaultVersionPtrType EngineModelDefaultVersionArgs

func EngineModelDefaultVersionPtr(v *EngineModelDefaultVersionArgs) EngineModelDefaultVersionPtrInput {
	return (*engineModelDefaultVersionPtrType)(v)
}

func (*engineModelDefaultVersionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineModelDefaultVersion)(nil)).Elem()
}

func (i *engineModelDefaultVersionPtrType) ToEngineModelDefaultVersionPtrOutput() EngineModelDefaultVersionPtrOutput {
	return i.ToEngineModelDefaultVersionPtrOutputWithContext(context.Background())
}

func (i *engineModelDefaultVersionPtrType) ToEngineModelDefaultVersionPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineModelDefaultVersionPtrOutput)
}

type EngineModelDefaultVersionOutput struct{ *pulumi.OutputState }

func (EngineModelDefaultVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineModelDefaultVersion)(nil)).Elem()
}

func (o EngineModelDefaultVersionOutput) ToEngineModelDefaultVersionOutput() EngineModelDefaultVersionOutput {
	return o
}

func (o EngineModelDefaultVersionOutput) ToEngineModelDefaultVersionOutputWithContext(ctx context.Context) EngineModelDefaultVersionOutput {
	return o
}

func (o EngineModelDefaultVersionOutput) ToEngineModelDefaultVersionPtrOutput() EngineModelDefaultVersionPtrOutput {
	return o.ToEngineModelDefaultVersionPtrOutputWithContext(context.Background())
}

func (o EngineModelDefaultVersionOutput) ToEngineModelDefaultVersionPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionPtrOutput {
	return o.ApplyT(func(v EngineModelDefaultVersion) *EngineModelDefaultVersion {
		return &v
	}).(EngineModelDefaultVersionPtrOutput)
}
func (o EngineModelDefaultVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EngineModelDefaultVersion) string { return v.Name }).(pulumi.StringOutput)
}

type EngineModelDefaultVersionPtrOutput struct{ *pulumi.OutputState }

func (EngineModelDefaultVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineModelDefaultVersion)(nil)).Elem()
}

func (o EngineModelDefaultVersionPtrOutput) ToEngineModelDefaultVersionPtrOutput() EngineModelDefaultVersionPtrOutput {
	return o
}

func (o EngineModelDefaultVersionPtrOutput) ToEngineModelDefaultVersionPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionPtrOutput {
	return o
}

func (o EngineModelDefaultVersionPtrOutput) Elem() EngineModelDefaultVersionOutput {
	return o.ApplyT(func(v *EngineModelDefaultVersion) EngineModelDefaultVersion { return *v }).(EngineModelDefaultVersionOutput)
}

func (o EngineModelDefaultVersionPtrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EngineModelDefaultVersion) string { return v.Name }).(pulumi.StringOutput)
}

type EngineModelDefaultVersionArgs struct {
	Name string `pulumi:"name"`
}

type EngineModelDefaultVersionArgsInput interface {
	pulumi.Input

	ToEngineModelDefaultVersionArgsOutput() EngineModelDefaultVersionArgsOutput
	ToEngineModelDefaultVersionArgsOutputWithContext(context.Context) EngineModelDefaultVersionArgsOutput
}

type EngineModelDefaultVersionArgsArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (EngineModelDefaultVersionArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineModelDefaultVersionArgs)(nil)).Elem()
}

func (i EngineModelDefaultVersionArgsArgs) ToEngineModelDefaultVersionArgsOutput() EngineModelDefaultVersionArgsOutput {
	return i.ToEngineModelDefaultVersionArgsOutputWithContext(context.Background())
}

func (i EngineModelDefaultVersionArgsArgs) ToEngineModelDefaultVersionArgsOutputWithContext(ctx context.Context) EngineModelDefaultVersionArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineModelDefaultVersionArgsOutput)
}

func (i EngineModelDefaultVersionArgsArgs) ToEngineModelDefaultVersionArgsPtrOutput() EngineModelDefaultVersionArgsPtrOutput {
	return i.ToEngineModelDefaultVersionArgsPtrOutputWithContext(context.Background())
}

func (i EngineModelDefaultVersionArgsArgs) ToEngineModelDefaultVersionArgsPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineModelDefaultVersionArgsOutput).ToEngineModelDefaultVersionArgsPtrOutputWithContext(ctx)
}

type EngineModelDefaultVersionArgsPtrInput interface {
	pulumi.Input

	ToEngineModelDefaultVersionArgsPtrOutput() EngineModelDefaultVersionArgsPtrOutput
	ToEngineModelDefaultVersionArgsPtrOutputWithContext(context.Context) EngineModelDefaultVersionArgsPtrOutput
}

type engineModelDefaultVersionArgsPtrType EngineModelDefaultVersionArgsArgs

func EngineModelDefaultVersionArgsPtr(v *EngineModelDefaultVersionArgsArgs) EngineModelDefaultVersionArgsPtrInput {
	return (*engineModelDefaultVersionArgsPtrType)(v)
}

func (*engineModelDefaultVersionArgsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineModelDefaultVersionArgs)(nil)).Elem()
}

func (i *engineModelDefaultVersionArgsPtrType) ToEngineModelDefaultVersionArgsPtrOutput() EngineModelDefaultVersionArgsPtrOutput {
	return i.ToEngineModelDefaultVersionArgsPtrOutputWithContext(context.Background())
}

func (i *engineModelDefaultVersionArgsPtrType) ToEngineModelDefaultVersionArgsPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineModelDefaultVersionArgsPtrOutput)
}

type EngineModelDefaultVersionArgsOutput struct{ *pulumi.OutputState }

func (EngineModelDefaultVersionArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineModelDefaultVersionArgs)(nil)).Elem()
}

func (o EngineModelDefaultVersionArgsOutput) ToEngineModelDefaultVersionArgsOutput() EngineModelDefaultVersionArgsOutput {
	return o
}

func (o EngineModelDefaultVersionArgsOutput) ToEngineModelDefaultVersionArgsOutputWithContext(ctx context.Context) EngineModelDefaultVersionArgsOutput {
	return o
}

func (o EngineModelDefaultVersionArgsOutput) ToEngineModelDefaultVersionArgsPtrOutput() EngineModelDefaultVersionArgsPtrOutput {
	return o.ToEngineModelDefaultVersionArgsPtrOutputWithContext(context.Background())
}

func (o EngineModelDefaultVersionArgsOutput) ToEngineModelDefaultVersionArgsPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionArgsPtrOutput {
	return o.ApplyT(func(v EngineModelDefaultVersionArgs) *EngineModelDefaultVersionArgs {
		return &v
	}).(EngineModelDefaultVersionArgsPtrOutput)
}
func (o EngineModelDefaultVersionArgsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EngineModelDefaultVersionArgs) string { return v.Name }).(pulumi.StringOutput)
}

type EngineModelDefaultVersionArgsPtrOutput struct{ *pulumi.OutputState }

func (EngineModelDefaultVersionArgsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineModelDefaultVersionArgs)(nil)).Elem()
}

func (o EngineModelDefaultVersionArgsPtrOutput) ToEngineModelDefaultVersionArgsPtrOutput() EngineModelDefaultVersionArgsPtrOutput {
	return o
}

func (o EngineModelDefaultVersionArgsPtrOutput) ToEngineModelDefaultVersionArgsPtrOutputWithContext(ctx context.Context) EngineModelDefaultVersionArgsPtrOutput {
	return o
}

func (o EngineModelDefaultVersionArgsPtrOutput) Elem() EngineModelDefaultVersionArgsOutput {
	return o.ApplyT(func(v *EngineModelDefaultVersionArgs) EngineModelDefaultVersionArgs { return *v }).(EngineModelDefaultVersionArgsOutput)
}

func (o EngineModelDefaultVersionArgsPtrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EngineModelDefaultVersionArgs) string { return v.Name }).(pulumi.StringOutput)
}

type EngineModelDefaultVersionState struct {
	Name string `pulumi:"name"`
}

type EngineModelDefaultVersionStateInput interface {
	pulumi.Input

	ToEngineModelDefaultVersionStateOutput() EngineModelDefaultVersionStateOutput
	ToEngineModelDefaultVersionStateOutputWithContext(context.Context) EngineModelDefaultVersionStateOutput
}

type EngineModelDefaultVersionStateArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (EngineModelDefaultVersionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineModelDefaultVersionState)(nil)).Elem()
}

func (i EngineModelDefaultVersionStateArgs) ToEngineModelDefaultVersionStateOutput() EngineModelDefaultVersionStateOutput {
	return i.ToEngineModelDefaultVersionStateOutputWithContext(context.Background())
}

func (i EngineModelDefaultVersionStateArgs) ToEngineModelDefaultVersionStateOutputWithContext(ctx context.Context) EngineModelDefaultVersionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineModelDefaultVersionStateOutput)
}

type EngineModelDefaultVersionStateOutput struct{ *pulumi.OutputState }

func (EngineModelDefaultVersionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineModelDefaultVersionState)(nil)).Elem()
}

func (o EngineModelDefaultVersionStateOutput) ToEngineModelDefaultVersionStateOutput() EngineModelDefaultVersionStateOutput {
	return o
}

func (o EngineModelDefaultVersionStateOutput) ToEngineModelDefaultVersionStateOutputWithContext(ctx context.Context) EngineModelDefaultVersionStateOutput {
	return o
}

func (o EngineModelDefaultVersionStateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EngineModelDefaultVersionState) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EngineModelDefaultVersionOutput{})
	pulumi.RegisterOutputType(EngineModelDefaultVersionPtrOutput{})
	pulumi.RegisterOutputType(EngineModelDefaultVersionArgsOutput{})
	pulumi.RegisterOutputType(EngineModelDefaultVersionArgsPtrOutput{})
	pulumi.RegisterOutputType(EngineModelDefaultVersionStateOutput{})
}
