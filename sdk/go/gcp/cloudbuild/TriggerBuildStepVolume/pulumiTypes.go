// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TriggerBuildStepVolume

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TriggerBuildStepVolume struct {
	Name string `pulumi:"name"`
	Path string `pulumi:"path"`
}

type TriggerBuildStepVolumeInput interface {
	pulumi.Input

	ToTriggerBuildStepVolumeOutput() TriggerBuildStepVolumeOutput
	ToTriggerBuildStepVolumeOutputWithContext(context.Context) TriggerBuildStepVolumeOutput
}

type TriggerBuildStepVolumeArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Path pulumi.StringInput `pulumi:"path"`
}

func (TriggerBuildStepVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerBuildStepVolume)(nil)).Elem()
}

func (i TriggerBuildStepVolumeArgs) ToTriggerBuildStepVolumeOutput() TriggerBuildStepVolumeOutput {
	return i.ToTriggerBuildStepVolumeOutputWithContext(context.Background())
}

func (i TriggerBuildStepVolumeArgs) ToTriggerBuildStepVolumeOutputWithContext(ctx context.Context) TriggerBuildStepVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerBuildStepVolumeOutput)
}

type TriggerBuildStepVolumeArrayInput interface {
	pulumi.Input

	ToTriggerBuildStepVolumeArrayOutput() TriggerBuildStepVolumeArrayOutput
	ToTriggerBuildStepVolumeArrayOutputWithContext(context.Context) TriggerBuildStepVolumeArrayOutput
}

type TriggerBuildStepVolumeArray []TriggerBuildStepVolumeInput

func (TriggerBuildStepVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerBuildStepVolume)(nil)).Elem()
}

func (i TriggerBuildStepVolumeArray) ToTriggerBuildStepVolumeArrayOutput() TriggerBuildStepVolumeArrayOutput {
	return i.ToTriggerBuildStepVolumeArrayOutputWithContext(context.Background())
}

func (i TriggerBuildStepVolumeArray) ToTriggerBuildStepVolumeArrayOutputWithContext(ctx context.Context) TriggerBuildStepVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerBuildStepVolumeArrayOutput)
}

type TriggerBuildStepVolumeOutput struct { *pulumi.OutputState }

func (TriggerBuildStepVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerBuildStepVolume)(nil)).Elem()
}

func (o TriggerBuildStepVolumeOutput) ToTriggerBuildStepVolumeOutput() TriggerBuildStepVolumeOutput {
	return o
}

func (o TriggerBuildStepVolumeOutput) ToTriggerBuildStepVolumeOutputWithContext(ctx context.Context) TriggerBuildStepVolumeOutput {
	return o
}

func (o TriggerBuildStepVolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v TriggerBuildStepVolume) string { return v.Name }).(pulumi.StringOutput)
}

func (o TriggerBuildStepVolumeOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func (v TriggerBuildStepVolume) string { return v.Path }).(pulumi.StringOutput)
}

type TriggerBuildStepVolumeArrayOutput struct { *pulumi.OutputState}

func (TriggerBuildStepVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerBuildStepVolume)(nil)).Elem()
}

func (o TriggerBuildStepVolumeArrayOutput) ToTriggerBuildStepVolumeArrayOutput() TriggerBuildStepVolumeArrayOutput {
	return o
}

func (o TriggerBuildStepVolumeArrayOutput) ToTriggerBuildStepVolumeArrayOutputWithContext(ctx context.Context) TriggerBuildStepVolumeArrayOutput {
	return o
}

func (o TriggerBuildStepVolumeArrayOutput) Index(i pulumi.IntInput) TriggerBuildStepVolumeOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) TriggerBuildStepVolume {
		return vs[0].([]TriggerBuildStepVolume)[vs[1].(int)]
	}).(TriggerBuildStepVolumeOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerBuildStepVolumeOutput{})
	pulumi.RegisterOutputType(TriggerBuildStepVolumeArrayOutput{})
}
