// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafusion

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InstanceNetworkConfig struct {
	// -
	// (Required)
	// The IP range in CIDR notation to use for the managed Data Fusion instance
	// nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
	IpAllocation string `pulumi:"ipAllocation"`
	// -
	// (Required)
	// Name of the network in the project with which the tenant project
	// will be peered for executing pipelines. In case of shared VPC where the network resides in another host
	// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
	Network string `pulumi:"network"`
}

type InstanceNetworkConfigInput interface {
	pulumi.Input

	ToInstanceNetworkConfigOutput() InstanceNetworkConfigOutput
	ToInstanceNetworkConfigOutputWithContext(context.Context) InstanceNetworkConfigOutput
}

type InstanceNetworkConfigArgs struct {
	// -
	// (Required)
	// The IP range in CIDR notation to use for the managed Data Fusion instance
	// nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
	IpAllocation pulumi.StringInput `pulumi:"ipAllocation"`
	// -
	// (Required)
	// Name of the network in the project with which the tenant project
	// will be peered for executing pipelines. In case of shared VPC where the network resides in another host
	// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
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

func InstanceNetworkConfigPtr(v *InstanceNetworkConfigArgs) InstanceNetworkConfigPtrInput {
	return (*instanceNetworkConfigPtrType)(v)
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

type InstanceNetworkConfigOutput struct{ *pulumi.OutputState }

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

// -
// (Required)
// The IP range in CIDR notation to use for the managed Data Fusion instance
// nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
func (o InstanceNetworkConfigOutput) IpAllocation() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceNetworkConfig) string { return v.IpAllocation }).(pulumi.StringOutput)
}

// -
// (Required)
// Name of the network in the project with which the tenant project
// will be peered for executing pipelines. In case of shared VPC where the network resides in another host
// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
func (o InstanceNetworkConfigOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceNetworkConfig) string { return v.Network }).(pulumi.StringOutput)
}

type InstanceNetworkConfigPtrOutput struct{ *pulumi.OutputState }

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
	return o.ApplyT(func(v *InstanceNetworkConfig) InstanceNetworkConfig { return *v }).(InstanceNetworkConfigOutput)
}

// -
// (Required)
// The IP range in CIDR notation to use for the managed Data Fusion instance
// nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
func (o InstanceNetworkConfigPtrOutput) IpAllocation() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceNetworkConfig) string { return v.IpAllocation }).(pulumi.StringOutput)
}

// -
// (Required)
// Name of the network in the project with which the tenant project
// will be peered for executing pipelines. In case of shared VPC where the network resides in another host
// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
func (o InstanceNetworkConfigPtrOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceNetworkConfig) string { return v.Network }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceNetworkConfigOutput{})
	pulumi.RegisterOutputType(InstanceNetworkConfigPtrOutput{})
}
