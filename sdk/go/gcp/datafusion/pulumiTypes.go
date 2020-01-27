// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package datafusion

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InstanceNetworkConfig struct {
	IpAllocation string `pulumi:"ipAllocation"`
	Network string `pulumi:"network"`
}

type InstanceNetworkConfigInput interface {
	pulumi.Input

	ToInstanceNetworkConfigOutput() InstanceNetworkConfigOutput
	ToInstanceNetworkConfigOutputWithContext(context.Context) InstanceNetworkConfigOutput
}

type InstanceNetworkConfigArgs struct {
	IpAllocation pulumi.StringInput `pulumi:"ipAllocation"`
	Network pulumi.StringInput `pulumi:"network"`
}

func (InstanceNetworkConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNetworkConfig)(nil)).Elem()
}

func (i InstanceNetworkConfigArgs) ToInstanceNetworkConfigOutput() InstanceNetworkConfigOutput {
	return i.ToInstanceNetworkConfigOutputWithContext(context.Background())
}

func (i InstanceNetworkConfigArgs) ToInstanceNetworkConfigOutputWithContext(ctx context.Context) InstanceNetworkConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNetworkConfigOutput)
}

func (i InstanceNetworkConfigArgs) ToInstanceNetworkConfigPtrOutput() InstanceNetworkConfigPtrOutput {
	return i.ToInstanceNetworkConfigPtrOutputWithContext(context.Background())
}

func (i InstanceNetworkConfigArgs) ToInstanceNetworkConfigPtrOutputWithContext(ctx context.Context) InstanceNetworkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNetworkConfigOutput).ToInstanceNetworkConfigPtrOutputWithContext(ctx)
}

type InstanceNetworkConfigPtrInput interface {
	pulumi.Input

	ToInstanceNetworkConfigPtrOutput() InstanceNetworkConfigPtrOutput
	ToInstanceNetworkConfigPtrOutputWithContext(context.Context) InstanceNetworkConfigPtrOutput
}

type instanceNetworkConfigPtrType InstanceNetworkConfigArgs

func InstanceNetworkConfigPtr(v *InstanceNetworkConfigArgs) InstanceNetworkConfigPtrInput {	return (*instanceNetworkConfigPtrType)(v)
}

func (*instanceNetworkConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceNetworkConfig)(nil)).Elem()
}

func (i *instanceNetworkConfigPtrType) ToInstanceNetworkConfigPtrOutput() InstanceNetworkConfigPtrOutput {
	return i.ToInstanceNetworkConfigPtrOutputWithContext(context.Background())
}

func (i *instanceNetworkConfigPtrType) ToInstanceNetworkConfigPtrOutputWithContext(ctx context.Context) InstanceNetworkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNetworkConfigPtrOutput)
}

type InstanceNetworkConfigOutput struct { *pulumi.OutputState }

func (InstanceNetworkConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNetworkConfig)(nil)).Elem()
}

func (o InstanceNetworkConfigOutput) ToInstanceNetworkConfigOutput() InstanceNetworkConfigOutput {
	return o
}

func (o InstanceNetworkConfigOutput) ToInstanceNetworkConfigOutputWithContext(ctx context.Context) InstanceNetworkConfigOutput {
	return o
}

func (o InstanceNetworkConfigOutput) ToInstanceNetworkConfigPtrOutput() InstanceNetworkConfigPtrOutput {
	return o.ToInstanceNetworkConfigPtrOutputWithContext(context.Background())
}

func (o InstanceNetworkConfigOutput) ToInstanceNetworkConfigPtrOutputWithContext(ctx context.Context) InstanceNetworkConfigPtrOutput {
	return o.ApplyT(func(v InstanceNetworkConfig) *InstanceNetworkConfig {
		return &v
	}).(InstanceNetworkConfigPtrOutput)
}
func (o InstanceNetworkConfigOutput) IpAllocation() pulumi.StringOutput {
	return o.ApplyT(func (v InstanceNetworkConfig) string { return v.IpAllocation }).(pulumi.StringOutput)
}

func (o InstanceNetworkConfigOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func (v InstanceNetworkConfig) string { return v.Network }).(pulumi.StringOutput)
}

type InstanceNetworkConfigPtrOutput struct { *pulumi.OutputState}

func (InstanceNetworkConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceNetworkConfig)(nil)).Elem()
}

func (o InstanceNetworkConfigPtrOutput) ToInstanceNetworkConfigPtrOutput() InstanceNetworkConfigPtrOutput {
	return o
}

func (o InstanceNetworkConfigPtrOutput) ToInstanceNetworkConfigPtrOutputWithContext(ctx context.Context) InstanceNetworkConfigPtrOutput {
	return o
}

func (o InstanceNetworkConfigPtrOutput) Elem() InstanceNetworkConfigOutput {
	return o.ApplyT(func (v *InstanceNetworkConfig) InstanceNetworkConfig { return *v }).(InstanceNetworkConfigOutput)
}

func (o InstanceNetworkConfigPtrOutput) IpAllocation() pulumi.StringOutput {
	return o.ApplyT(func (v InstanceNetworkConfig) string { return v.IpAllocation }).(pulumi.StringOutput)
}

func (o InstanceNetworkConfigPtrOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func (v InstanceNetworkConfig) string { return v.Network }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceNetworkConfigOutput{})
	pulumi.RegisterOutputType(InstanceNetworkConfigPtrOutput{})
}
