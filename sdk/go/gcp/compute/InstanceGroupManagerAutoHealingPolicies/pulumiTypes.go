// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package InstanceGroupManagerAutoHealingPolicies

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InstanceGroupManagerAutoHealingPolicies struct {
	HealthCheck string `pulumi:"healthCheck"`
	InitialDelaySec int `pulumi:"initialDelaySec"`
}

type InstanceGroupManagerAutoHealingPoliciesInput interface {
	pulumi.Input

	ToInstanceGroupManagerAutoHealingPoliciesOutput() InstanceGroupManagerAutoHealingPoliciesOutput
	ToInstanceGroupManagerAutoHealingPoliciesOutputWithContext(context.Context) InstanceGroupManagerAutoHealingPoliciesOutput
}

type InstanceGroupManagerAutoHealingPoliciesArgs struct {
	HealthCheck pulumi.StringInput `pulumi:"healthCheck"`
	InitialDelaySec pulumi.IntInput `pulumi:"initialDelaySec"`
}

func (InstanceGroupManagerAutoHealingPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (i InstanceGroupManagerAutoHealingPoliciesArgs) ToInstanceGroupManagerAutoHealingPoliciesOutput() InstanceGroupManagerAutoHealingPoliciesOutput {
	return i.ToInstanceGroupManagerAutoHealingPoliciesOutputWithContext(context.Background())
}

func (i InstanceGroupManagerAutoHealingPoliciesArgs) ToInstanceGroupManagerAutoHealingPoliciesOutputWithContext(ctx context.Context) InstanceGroupManagerAutoHealingPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerAutoHealingPoliciesOutput)
}

func (i InstanceGroupManagerAutoHealingPoliciesArgs) ToInstanceGroupManagerAutoHealingPoliciesPtrOutput() InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return i.ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Background())
}

func (i InstanceGroupManagerAutoHealingPoliciesArgs) ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerAutoHealingPoliciesOutput).ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx)
}

type InstanceGroupManagerAutoHealingPoliciesPtrInput interface {
	pulumi.Input

	ToInstanceGroupManagerAutoHealingPoliciesPtrOutput() InstanceGroupManagerAutoHealingPoliciesPtrOutput
	ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Context) InstanceGroupManagerAutoHealingPoliciesPtrOutput
}

type instanceGroupManagerAutoHealingPoliciesPtrType InstanceGroupManagerAutoHealingPoliciesArgs

func InstanceGroupManagerAutoHealingPoliciesPtr(v *InstanceGroupManagerAutoHealingPoliciesArgs) InstanceGroupManagerAutoHealingPoliciesPtrInput {	return (*instanceGroupManagerAutoHealingPoliciesPtrType)(v)
}

func (*instanceGroupManagerAutoHealingPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (i *instanceGroupManagerAutoHealingPoliciesPtrType) ToInstanceGroupManagerAutoHealingPoliciesPtrOutput() InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return i.ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Background())
}

func (i *instanceGroupManagerAutoHealingPoliciesPtrType) ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerAutoHealingPoliciesPtrOutput)
}

type InstanceGroupManagerAutoHealingPoliciesOutput struct { *pulumi.OutputState }

func (InstanceGroupManagerAutoHealingPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (o InstanceGroupManagerAutoHealingPoliciesOutput) ToInstanceGroupManagerAutoHealingPoliciesOutput() InstanceGroupManagerAutoHealingPoliciesOutput {
	return o
}

func (o InstanceGroupManagerAutoHealingPoliciesOutput) ToInstanceGroupManagerAutoHealingPoliciesOutputWithContext(ctx context.Context) InstanceGroupManagerAutoHealingPoliciesOutput {
	return o
}

func (o InstanceGroupManagerAutoHealingPoliciesOutput) ToInstanceGroupManagerAutoHealingPoliciesPtrOutput() InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o.ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Background())
}

func (o InstanceGroupManagerAutoHealingPoliciesOutput) ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o.ApplyT(func(v InstanceGroupManagerAutoHealingPolicies) *InstanceGroupManagerAutoHealingPolicies {
		return &v
	}).(InstanceGroupManagerAutoHealingPoliciesPtrOutput)
}
func (o InstanceGroupManagerAutoHealingPoliciesOutput) HealthCheck() pulumi.StringOutput {
	return o.ApplyT(func (v InstanceGroupManagerAutoHealingPolicies) string { return v.HealthCheck }).(pulumi.StringOutput)
}

func (o InstanceGroupManagerAutoHealingPoliciesOutput) InitialDelaySec() pulumi.IntOutput {
	return o.ApplyT(func (v InstanceGroupManagerAutoHealingPolicies) int { return v.InitialDelaySec }).(pulumi.IntOutput)
}

type InstanceGroupManagerAutoHealingPoliciesPtrOutput struct { *pulumi.OutputState}

func (InstanceGroupManagerAutoHealingPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (o InstanceGroupManagerAutoHealingPoliciesPtrOutput) ToInstanceGroupManagerAutoHealingPoliciesPtrOutput() InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o
}

func (o InstanceGroupManagerAutoHealingPoliciesPtrOutput) ToInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o
}

func (o InstanceGroupManagerAutoHealingPoliciesPtrOutput) Elem() InstanceGroupManagerAutoHealingPoliciesOutput {
	return o.ApplyT(func (v *InstanceGroupManagerAutoHealingPolicies) InstanceGroupManagerAutoHealingPolicies { return *v }).(InstanceGroupManagerAutoHealingPoliciesOutput)
}

func (o InstanceGroupManagerAutoHealingPoliciesPtrOutput) HealthCheck() pulumi.StringOutput {
	return o.ApplyT(func (v InstanceGroupManagerAutoHealingPolicies) string { return v.HealthCheck }).(pulumi.StringOutput)
}

func (o InstanceGroupManagerAutoHealingPoliciesPtrOutput) InitialDelaySec() pulumi.IntOutput {
	return o.ApplyT(func (v InstanceGroupManagerAutoHealingPolicies) int { return v.InitialDelaySec }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceGroupManagerAutoHealingPoliciesOutput{})
	pulumi.RegisterOutputType(InstanceGroupManagerAutoHealingPoliciesPtrOutput{})
}
