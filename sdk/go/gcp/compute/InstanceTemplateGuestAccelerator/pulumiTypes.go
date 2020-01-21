// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package InstanceTemplateGuestAccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InstanceTemplateGuestAccelerator struct {
	Count int `pulumi:"count"`
	Type string `pulumi:"type"`
}

type InstanceTemplateGuestAcceleratorInput interface {
	pulumi.Input

	ToInstanceTemplateGuestAcceleratorOutput() InstanceTemplateGuestAcceleratorOutput
	ToInstanceTemplateGuestAcceleratorOutputWithContext(context.Context) InstanceTemplateGuestAcceleratorOutput
}

type InstanceTemplateGuestAcceleratorArgs struct {
	Count pulumi.IntInput `pulumi:"count"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (InstanceTemplateGuestAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplateGuestAccelerator)(nil)).Elem()
}

func (i InstanceTemplateGuestAcceleratorArgs) ToInstanceTemplateGuestAcceleratorOutput() InstanceTemplateGuestAcceleratorOutput {
	return i.ToInstanceTemplateGuestAcceleratorOutputWithContext(context.Background())
}

func (i InstanceTemplateGuestAcceleratorArgs) ToInstanceTemplateGuestAcceleratorOutputWithContext(ctx context.Context) InstanceTemplateGuestAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTemplateGuestAcceleratorOutput)
}

type InstanceTemplateGuestAcceleratorArrayInput interface {
	pulumi.Input

	ToInstanceTemplateGuestAcceleratorArrayOutput() InstanceTemplateGuestAcceleratorArrayOutput
	ToInstanceTemplateGuestAcceleratorArrayOutputWithContext(context.Context) InstanceTemplateGuestAcceleratorArrayOutput
}

type InstanceTemplateGuestAcceleratorArray []InstanceTemplateGuestAcceleratorInput

func (InstanceTemplateGuestAcceleratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceTemplateGuestAccelerator)(nil)).Elem()
}

func (i InstanceTemplateGuestAcceleratorArray) ToInstanceTemplateGuestAcceleratorArrayOutput() InstanceTemplateGuestAcceleratorArrayOutput {
	return i.ToInstanceTemplateGuestAcceleratorArrayOutputWithContext(context.Background())
}

func (i InstanceTemplateGuestAcceleratorArray) ToInstanceTemplateGuestAcceleratorArrayOutputWithContext(ctx context.Context) InstanceTemplateGuestAcceleratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTemplateGuestAcceleratorArrayOutput)
}

type InstanceTemplateGuestAcceleratorOutput struct { *pulumi.OutputState }

func (InstanceTemplateGuestAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplateGuestAccelerator)(nil)).Elem()
}

func (o InstanceTemplateGuestAcceleratorOutput) ToInstanceTemplateGuestAcceleratorOutput() InstanceTemplateGuestAcceleratorOutput {
	return o
}

func (o InstanceTemplateGuestAcceleratorOutput) ToInstanceTemplateGuestAcceleratorOutputWithContext(ctx context.Context) InstanceTemplateGuestAcceleratorOutput {
	return o
}

func (o InstanceTemplateGuestAcceleratorOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func (v InstanceTemplateGuestAccelerator) int { return v.Count }).(pulumi.IntOutput)
}

func (o InstanceTemplateGuestAcceleratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v InstanceTemplateGuestAccelerator) string { return v.Type }).(pulumi.StringOutput)
}

type InstanceTemplateGuestAcceleratorArrayOutput struct { *pulumi.OutputState}

func (InstanceTemplateGuestAcceleratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceTemplateGuestAccelerator)(nil)).Elem()
}

func (o InstanceTemplateGuestAcceleratorArrayOutput) ToInstanceTemplateGuestAcceleratorArrayOutput() InstanceTemplateGuestAcceleratorArrayOutput {
	return o
}

func (o InstanceTemplateGuestAcceleratorArrayOutput) ToInstanceTemplateGuestAcceleratorArrayOutputWithContext(ctx context.Context) InstanceTemplateGuestAcceleratorArrayOutput {
	return o
}

func (o InstanceTemplateGuestAcceleratorArrayOutput) Index(i pulumi.IntInput) InstanceTemplateGuestAcceleratorOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) InstanceTemplateGuestAccelerator {
		return vs[0].([]InstanceTemplateGuestAccelerator)[vs[1].(int)]
	}).(InstanceTemplateGuestAcceleratorOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceTemplateGuestAcceleratorOutput{})
	pulumi.RegisterOutputType(InstanceTemplateGuestAcceleratorArrayOutput{})
}
