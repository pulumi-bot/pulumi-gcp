// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package filestore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InstanceFileShares struct {
	CapacityGb int    `pulumi:"capacityGb"`
	Name       string `pulumi:"name"`
}

type InstanceFileSharesInput interface {
	pulumi.Input

	ToInstanceFileSharesOutput() InstanceFileSharesOutput
	ToInstanceFileSharesOutputWithContext(context.Context) InstanceFileSharesOutput
}

type InstanceFileSharesArgs struct {
	CapacityGb pulumi.IntInput    `pulumi:"capacityGb"`
	Name       pulumi.StringInput `pulumi:"name"`
}

func (InstanceFileSharesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFileShares)(nil)).Elem()
}

func (i InstanceFileSharesArgs) ToInstanceFileSharesOutput() InstanceFileSharesOutput {
	return i.ToInstanceFileSharesOutputWithContext(context.Background())
}

func (i InstanceFileSharesArgs) ToInstanceFileSharesOutputWithContext(ctx context.Context) InstanceFileSharesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFileSharesOutput)
}

func (i InstanceFileSharesArgs) ToInstanceFileSharesPtrOutput() InstanceFileSharesPtrOutput {
	return i.ToInstanceFileSharesPtrOutputWithContext(context.Background())
}

func (i InstanceFileSharesArgs) ToInstanceFileSharesPtrOutputWithContext(ctx context.Context) InstanceFileSharesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFileSharesOutput).ToInstanceFileSharesPtrOutputWithContext(ctx)
}

type InstanceFileSharesPtrInput interface {
	pulumi.Input

	ToInstanceFileSharesPtrOutput() InstanceFileSharesPtrOutput
	ToInstanceFileSharesPtrOutputWithContext(context.Context) InstanceFileSharesPtrOutput
}

type instanceFileSharesPtrType InstanceFileSharesArgs

func InstanceFileSharesPtr(v *InstanceFileSharesArgs) InstanceFileSharesPtrInput {
	return (*instanceFileSharesPtrType)(v)
}

func (*instanceFileSharesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFileShares)(nil)).Elem()
}

func (i *instanceFileSharesPtrType) ToInstanceFileSharesPtrOutput() InstanceFileSharesPtrOutput {
	return i.ToInstanceFileSharesPtrOutputWithContext(context.Background())
}

func (i *instanceFileSharesPtrType) ToInstanceFileSharesPtrOutputWithContext(ctx context.Context) InstanceFileSharesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFileSharesPtrOutput)
}

type InstanceFileSharesOutput struct{ *pulumi.OutputState }

func (InstanceFileSharesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFileShares)(nil)).Elem()
}

func (o InstanceFileSharesOutput) ToInstanceFileSharesOutput() InstanceFileSharesOutput {
	return o
}

func (o InstanceFileSharesOutput) ToInstanceFileSharesOutputWithContext(ctx context.Context) InstanceFileSharesOutput {
	return o
}

func (o InstanceFileSharesOutput) ToInstanceFileSharesPtrOutput() InstanceFileSharesPtrOutput {
	return o.ToInstanceFileSharesPtrOutputWithContext(context.Background())
}

func (o InstanceFileSharesOutput) ToInstanceFileSharesPtrOutputWithContext(ctx context.Context) InstanceFileSharesPtrOutput {
	return o.ApplyT(func(v InstanceFileShares) *InstanceFileShares {
		return &v
	}).(InstanceFileSharesPtrOutput)
}
func (o InstanceFileSharesOutput) CapacityGb() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceFileShares) int { return v.CapacityGb }).(pulumi.IntOutput)
}

func (o InstanceFileSharesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceFileShares) string { return v.Name }).(pulumi.StringOutput)
}

type InstanceFileSharesPtrOutput struct{ *pulumi.OutputState }

func (InstanceFileSharesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFileShares)(nil)).Elem()
}

func (o InstanceFileSharesPtrOutput) ToInstanceFileSharesPtrOutput() InstanceFileSharesPtrOutput {
	return o
}

func (o InstanceFileSharesPtrOutput) ToInstanceFileSharesPtrOutputWithContext(ctx context.Context) InstanceFileSharesPtrOutput {
	return o
}

func (o InstanceFileSharesPtrOutput) Elem() InstanceFileSharesOutput {
	return o.ApplyT(func(v *InstanceFileShares) InstanceFileShares { return *v }).(InstanceFileSharesOutput)
}

func (o InstanceFileSharesPtrOutput) CapacityGb() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceFileShares) int { return v.CapacityGb }).(pulumi.IntOutput)
}

func (o InstanceFileSharesPtrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceFileShares) string { return v.Name }).(pulumi.StringOutput)
}

type InstanceNetwork struct {
	IpAddresses     []string `pulumi:"ipAddresses"`
	Modes           []string `pulumi:"modes"`
	Network         string   `pulumi:"network"`
	ReservedIpRange *string  `pulumi:"reservedIpRange"`
}

type InstanceNetworkInput interface {
	pulumi.Input

	ToInstanceNetworkOutput() InstanceNetworkOutput
	ToInstanceNetworkOutputWithContext(context.Context) InstanceNetworkOutput
}

type InstanceNetworkArgs struct {
	IpAddresses     pulumi.StringArrayInput `pulumi:"ipAddresses"`
	Modes           pulumi.StringArrayInput `pulumi:"modes"`
	Network         pulumi.StringInput      `pulumi:"network"`
	ReservedIpRange pulumi.StringPtrInput   `pulumi:"reservedIpRange"`
}

func (InstanceNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNetwork)(nil)).Elem()
}

func (i InstanceNetworkArgs) ToInstanceNetworkOutput() InstanceNetworkOutput {
	return i.ToInstanceNetworkOutputWithContext(context.Background())
}

func (i InstanceNetworkArgs) ToInstanceNetworkOutputWithContext(ctx context.Context) InstanceNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNetworkOutput)
}

type InstanceNetworkArrayInput interface {
	pulumi.Input

	ToInstanceNetworkArrayOutput() InstanceNetworkArrayOutput
	ToInstanceNetworkArrayOutputWithContext(context.Context) InstanceNetworkArrayOutput
}

type InstanceNetworkArray []InstanceNetworkInput

func (InstanceNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceNetwork)(nil)).Elem()
}

func (i InstanceNetworkArray) ToInstanceNetworkArrayOutput() InstanceNetworkArrayOutput {
	return i.ToInstanceNetworkArrayOutputWithContext(context.Background())
}

func (i InstanceNetworkArray) ToInstanceNetworkArrayOutputWithContext(ctx context.Context) InstanceNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNetworkArrayOutput)
}

type InstanceNetworkOutput struct{ *pulumi.OutputState }

func (InstanceNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNetwork)(nil)).Elem()
}

func (o InstanceNetworkOutput) ToInstanceNetworkOutput() InstanceNetworkOutput {
	return o
}

func (o InstanceNetworkOutput) ToInstanceNetworkOutputWithContext(ctx context.Context) InstanceNetworkOutput {
	return o
}

func (o InstanceNetworkOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceNetwork) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o InstanceNetworkOutput) Modes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceNetwork) []string { return v.Modes }).(pulumi.StringArrayOutput)
}

func (o InstanceNetworkOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceNetwork) string { return v.Network }).(pulumi.StringOutput)
}

func (o InstanceNetworkOutput) ReservedIpRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceNetwork) *string { return v.ReservedIpRange }).(pulumi.StringPtrOutput)
}

type InstanceNetworkArrayOutput struct{ *pulumi.OutputState }

func (InstanceNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceNetwork)(nil)).Elem()
}

func (o InstanceNetworkArrayOutput) ToInstanceNetworkArrayOutput() InstanceNetworkArrayOutput {
	return o
}

func (o InstanceNetworkArrayOutput) ToInstanceNetworkArrayOutputWithContext(ctx context.Context) InstanceNetworkArrayOutput {
	return o
}

func (o InstanceNetworkArrayOutput) Index(i pulumi.IntInput) InstanceNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceNetwork {
		return vs[0].([]InstanceNetwork)[vs[1].(int)]
	}).(InstanceNetworkOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceFileSharesOutput{})
	pulumi.RegisterOutputType(InstanceFileSharesPtrOutput{})
	pulumi.RegisterOutputType(InstanceNetworkOutput{})
	pulumi.RegisterOutputType(InstanceNetworkArrayOutput{})
}
