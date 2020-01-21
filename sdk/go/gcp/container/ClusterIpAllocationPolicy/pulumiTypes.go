// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterIpAllocationPolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterIpAllocationPolicy struct {
	ClusterIpv4CidrBlock *string `pulumi:"clusterIpv4CidrBlock"`
	ClusterSecondaryRangeName *string `pulumi:"clusterSecondaryRangeName"`
	ServicesIpv4CidrBlock *string `pulumi:"servicesIpv4CidrBlock"`
	ServicesSecondaryRangeName *string `pulumi:"servicesSecondaryRangeName"`
}

type ClusterIpAllocationPolicyInput interface {
	pulumi.Input

	ToClusterIpAllocationPolicyOutput() ClusterIpAllocationPolicyOutput
	ToClusterIpAllocationPolicyOutputWithContext(context.Context) ClusterIpAllocationPolicyOutput
}

type ClusterIpAllocationPolicyArgs struct {
	ClusterIpv4CidrBlock pulumi.StringPtrInput `pulumi:"clusterIpv4CidrBlock"`
	ClusterSecondaryRangeName pulumi.StringPtrInput `pulumi:"clusterSecondaryRangeName"`
	ServicesIpv4CidrBlock pulumi.StringPtrInput `pulumi:"servicesIpv4CidrBlock"`
	ServicesSecondaryRangeName pulumi.StringPtrInput `pulumi:"servicesSecondaryRangeName"`
}

func (ClusterIpAllocationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIpAllocationPolicy)(nil)).Elem()
}

func (i ClusterIpAllocationPolicyArgs) ToClusterIpAllocationPolicyOutput() ClusterIpAllocationPolicyOutput {
	return i.ToClusterIpAllocationPolicyOutputWithContext(context.Background())
}

func (i ClusterIpAllocationPolicyArgs) ToClusterIpAllocationPolicyOutputWithContext(ctx context.Context) ClusterIpAllocationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIpAllocationPolicyOutput)
}

func (i ClusterIpAllocationPolicyArgs) ToClusterIpAllocationPolicyPtrOutput() ClusterIpAllocationPolicyPtrOutput {
	return i.ToClusterIpAllocationPolicyPtrOutputWithContext(context.Background())
}

func (i ClusterIpAllocationPolicyArgs) ToClusterIpAllocationPolicyPtrOutputWithContext(ctx context.Context) ClusterIpAllocationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIpAllocationPolicyOutput).ToClusterIpAllocationPolicyPtrOutputWithContext(ctx)
}

type ClusterIpAllocationPolicyPtrInput interface {
	pulumi.Input

	ToClusterIpAllocationPolicyPtrOutput() ClusterIpAllocationPolicyPtrOutput
	ToClusterIpAllocationPolicyPtrOutputWithContext(context.Context) ClusterIpAllocationPolicyPtrOutput
}

type clusterIpAllocationPolicyPtrType ClusterIpAllocationPolicyArgs

func ClusterIpAllocationPolicyPtr(v *ClusterIpAllocationPolicyArgs) ClusterIpAllocationPolicyPtrInput {	return (*clusterIpAllocationPolicyPtrType)(v)
}

func (*clusterIpAllocationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIpAllocationPolicy)(nil)).Elem()
}

func (i *clusterIpAllocationPolicyPtrType) ToClusterIpAllocationPolicyPtrOutput() ClusterIpAllocationPolicyPtrOutput {
	return i.ToClusterIpAllocationPolicyPtrOutputWithContext(context.Background())
}

func (i *clusterIpAllocationPolicyPtrType) ToClusterIpAllocationPolicyPtrOutputWithContext(ctx context.Context) ClusterIpAllocationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIpAllocationPolicyPtrOutput)
}

type ClusterIpAllocationPolicyOutput struct { *pulumi.OutputState }

func (ClusterIpAllocationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIpAllocationPolicy)(nil)).Elem()
}

func (o ClusterIpAllocationPolicyOutput) ToClusterIpAllocationPolicyOutput() ClusterIpAllocationPolicyOutput {
	return o
}

func (o ClusterIpAllocationPolicyOutput) ToClusterIpAllocationPolicyOutputWithContext(ctx context.Context) ClusterIpAllocationPolicyOutput {
	return o
}

func (o ClusterIpAllocationPolicyOutput) ToClusterIpAllocationPolicyPtrOutput() ClusterIpAllocationPolicyPtrOutput {
	return o.ToClusterIpAllocationPolicyPtrOutputWithContext(context.Background())
}

func (o ClusterIpAllocationPolicyOutput) ToClusterIpAllocationPolicyPtrOutputWithContext(ctx context.Context) ClusterIpAllocationPolicyPtrOutput {
	return o.ApplyT(func(v ClusterIpAllocationPolicy) *ClusterIpAllocationPolicy {
		return &v
	}).(ClusterIpAllocationPolicyPtrOutput)
}
func (o ClusterIpAllocationPolicyOutput) ClusterIpv4CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ClusterIpv4CidrBlock }).(pulumi.StringPtrOutput)
}

func (o ClusterIpAllocationPolicyOutput) ClusterSecondaryRangeName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ClusterSecondaryRangeName }).(pulumi.StringPtrOutput)
}

func (o ClusterIpAllocationPolicyOutput) ServicesIpv4CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ServicesIpv4CidrBlock }).(pulumi.StringPtrOutput)
}

func (o ClusterIpAllocationPolicyOutput) ServicesSecondaryRangeName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ServicesSecondaryRangeName }).(pulumi.StringPtrOutput)
}

type ClusterIpAllocationPolicyPtrOutput struct { *pulumi.OutputState}

func (ClusterIpAllocationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIpAllocationPolicy)(nil)).Elem()
}

func (o ClusterIpAllocationPolicyPtrOutput) ToClusterIpAllocationPolicyPtrOutput() ClusterIpAllocationPolicyPtrOutput {
	return o
}

func (o ClusterIpAllocationPolicyPtrOutput) ToClusterIpAllocationPolicyPtrOutputWithContext(ctx context.Context) ClusterIpAllocationPolicyPtrOutput {
	return o
}

func (o ClusterIpAllocationPolicyPtrOutput) Elem() ClusterIpAllocationPolicyOutput {
	return o.ApplyT(func (v *ClusterIpAllocationPolicy) ClusterIpAllocationPolicy { return *v }).(ClusterIpAllocationPolicyOutput)
}

func (o ClusterIpAllocationPolicyPtrOutput) ClusterIpv4CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ClusterIpv4CidrBlock }).(pulumi.StringPtrOutput)
}

func (o ClusterIpAllocationPolicyPtrOutput) ClusterSecondaryRangeName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ClusterSecondaryRangeName }).(pulumi.StringPtrOutput)
}

func (o ClusterIpAllocationPolicyPtrOutput) ServicesIpv4CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ServicesIpv4CidrBlock }).(pulumi.StringPtrOutput)
}

func (o ClusterIpAllocationPolicyPtrOutput) ServicesSecondaryRangeName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterIpAllocationPolicy) *string { return v.ServicesSecondaryRangeName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterIpAllocationPolicyOutput{})
	pulumi.RegisterOutputType(ClusterIpAllocationPolicyPtrOutput{})
}
