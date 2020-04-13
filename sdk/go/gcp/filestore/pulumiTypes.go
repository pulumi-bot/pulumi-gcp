// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package filestore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InstanceFileShares struct {
	// -
	// (Required)
	// File share capacity in GiB. This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	CapacityGb int `pulumi:"capacityGb"`
	// -
	// (Required)
	// The name of the fileshare (16 characters or less)
	Name string `pulumi:"name"`
}

type InstanceFileSharesInput interface {
	pulumi.Input

	ToInstanceFileSharesOutput() InstanceFileSharesOutput
	ToInstanceFileSharesOutputWithContext(context.Context) InstanceFileSharesOutput
}

type InstanceFileSharesArgs struct {
	// -
	// (Required)
	// File share capacity in GiB. This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	CapacityGb pulumi.IntInput `pulumi:"capacityGb"`
	// -
	// (Required)
	// The name of the fileshare (16 characters or less)
	Name pulumi.StringInput `pulumi:"name"`
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

// -
// (Required)
// File share capacity in GiB. This must be at least 1024 GiB
// for the standard tier, or 2560 GiB for the premium tier.
func (o InstanceFileSharesOutput) CapacityGb() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceFileShares) int { return v.CapacityGb }).(pulumi.IntOutput)
}

// -
// (Required)
// The name of the fileshare (16 characters or less)
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

// -
// (Required)
// File share capacity in GiB. This must be at least 1024 GiB
// for the standard tier, or 2560 GiB for the premium tier.
func (o InstanceFileSharesPtrOutput) CapacityGb() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceFileShares) int { return v.CapacityGb }).(pulumi.IntOutput)
}

// -
// (Required)
// The name of the fileshare (16 characters or less)
func (o InstanceFileSharesPtrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceFileShares) string { return v.Name }).(pulumi.StringOutput)
}

type InstanceNetwork struct {
	// -
	// A list of IPv4 or IPv6 addresses.
	IpAddresses []string `pulumi:"ipAddresses"`
	// -
	// (Required)
	// IP versions for which the instance has
	// IP addresses assigned.
	Modes []string `pulumi:"modes"`
	// -
	// (Required)
	// The name of the GCE VPC network to which the
	// instance is connected.
	Network string `pulumi:"network"`
	// -
	// (Optional)
	// A /29 CIDR block that identifies the range of IP
	// addresses reserved for this instance.
	ReservedIpRange *string `pulumi:"reservedIpRange"`
}

type InstanceNetworkInput interface {
	pulumi.Input

	ToInstanceNetworkOutput() InstanceNetworkOutput
	ToInstanceNetworkOutputWithContext(context.Context) InstanceNetworkOutput
}

type InstanceNetworkArgs struct {
	// -
	// A list of IPv4 or IPv6 addresses.
	IpAddresses pulumi.StringArrayInput `pulumi:"ipAddresses"`
	// -
	// (Required)
	// IP versions for which the instance has
	// IP addresses assigned.
	Modes pulumi.StringArrayInput `pulumi:"modes"`
	// -
	// (Required)
	// The name of the GCE VPC network to which the
	// instance is connected.
	Network pulumi.StringInput `pulumi:"network"`
	// -
	// (Optional)
	// A /29 CIDR block that identifies the range of IP
	// addresses reserved for this instance.
	ReservedIpRange pulumi.StringPtrInput `pulumi:"reservedIpRange"`
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

// -
// A list of IPv4 or IPv6 addresses.
func (o InstanceNetworkOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceNetwork) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// -
// (Required)
// IP versions for which the instance has
// IP addresses assigned.
func (o InstanceNetworkOutput) Modes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceNetwork) []string { return v.Modes }).(pulumi.StringArrayOutput)
}

// -
// (Required)
// The name of the GCE VPC network to which the
// instance is connected.
func (o InstanceNetworkOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceNetwork) string { return v.Network }).(pulumi.StringOutput)
}

// -
// (Optional)
// A /29 CIDR block that identifies the range of IP
// addresses reserved for this instance.
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
