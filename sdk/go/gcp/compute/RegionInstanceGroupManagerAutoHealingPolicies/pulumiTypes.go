// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegionInstanceGroupManagerAutoHealingPolicies

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegionInstanceGroupManagerAutoHealingPolicies struct {
	HealthCheck string `pulumi:"healthCheck"`
	InitialDelaySec int `pulumi:"initialDelaySec"`
}

type RegionInstanceGroupManagerAutoHealingPoliciesInput interface {
	pulumi.Input

	ToRegionInstanceGroupManagerAutoHealingPoliciesOutput() RegionInstanceGroupManagerAutoHealingPoliciesOutput
	ToRegionInstanceGroupManagerAutoHealingPoliciesOutputWithContext(context.Context) RegionInstanceGroupManagerAutoHealingPoliciesOutput
}

type RegionInstanceGroupManagerAutoHealingPoliciesArgs struct {
	HealthCheck pulumi.StringInput `pulumi:"healthCheck"`
	InitialDelaySec pulumi.IntInput `pulumi:"initialDelaySec"`
}

func (RegionInstanceGroupManagerAutoHealingPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionInstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (i RegionInstanceGroupManagerAutoHealingPoliciesArgs) ToRegionInstanceGroupManagerAutoHealingPoliciesOutput() RegionInstanceGroupManagerAutoHealingPoliciesOutput {
	return i.ToRegionInstanceGroupManagerAutoHealingPoliciesOutputWithContext(context.Background())
}

func (i RegionInstanceGroupManagerAutoHealingPoliciesArgs) ToRegionInstanceGroupManagerAutoHealingPoliciesOutputWithContext(ctx context.Context) RegionInstanceGroupManagerAutoHealingPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerAutoHealingPoliciesOutput)
}

func (i RegionInstanceGroupManagerAutoHealingPoliciesArgs) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutput() RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return i.ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Background())
}

func (i RegionInstanceGroupManagerAutoHealingPoliciesArgs) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerAutoHealingPoliciesOutput).ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx)
}

type RegionInstanceGroupManagerAutoHealingPoliciesPtrInput interface {
	pulumi.Input

	ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutput() RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput
	ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Context) RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput
}

type regionInstanceGroupManagerAutoHealingPoliciesPtrType RegionInstanceGroupManagerAutoHealingPoliciesArgs

func RegionInstanceGroupManagerAutoHealingPoliciesPtr(v *RegionInstanceGroupManagerAutoHealingPoliciesArgs) RegionInstanceGroupManagerAutoHealingPoliciesPtrInput {	return (*regionInstanceGroupManagerAutoHealingPoliciesPtrType)(v)
}

func (*regionInstanceGroupManagerAutoHealingPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionInstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (i *regionInstanceGroupManagerAutoHealingPoliciesPtrType) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutput() RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return i.ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Background())
}

func (i *regionInstanceGroupManagerAutoHealingPoliciesPtrType) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput)
}

type RegionInstanceGroupManagerAutoHealingPoliciesOutput struct { *pulumi.OutputState }

func (RegionInstanceGroupManagerAutoHealingPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionInstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesOutput) ToRegionInstanceGroupManagerAutoHealingPoliciesOutput() RegionInstanceGroupManagerAutoHealingPoliciesOutput {
	return o
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesOutput) ToRegionInstanceGroupManagerAutoHealingPoliciesOutputWithContext(ctx context.Context) RegionInstanceGroupManagerAutoHealingPoliciesOutput {
	return o
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesOutput) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutput() RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o.ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(context.Background())
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesOutput) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o.ApplyT(func(v RegionInstanceGroupManagerAutoHealingPolicies) *RegionInstanceGroupManagerAutoHealingPolicies {
		return &v
	}).(RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput)
}
func (o RegionInstanceGroupManagerAutoHealingPoliciesOutput) HealthCheck() pulumi.StringOutput {
	return o.ApplyT(func (v RegionInstanceGroupManagerAutoHealingPolicies) string { return v.HealthCheck }).(pulumi.StringOutput)
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesOutput) InitialDelaySec() pulumi.IntOutput {
	return o.ApplyT(func (v RegionInstanceGroupManagerAutoHealingPolicies) int { return v.InitialDelaySec }).(pulumi.IntOutput)
}

type RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput struct { *pulumi.OutputState}

func (RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionInstanceGroupManagerAutoHealingPolicies)(nil)).Elem()
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutput() RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput) ToRegionInstanceGroupManagerAutoHealingPoliciesPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput) Elem() RegionInstanceGroupManagerAutoHealingPoliciesOutput {
	return o.ApplyT(func (v *RegionInstanceGroupManagerAutoHealingPolicies) RegionInstanceGroupManagerAutoHealingPolicies { return *v }).(RegionInstanceGroupManagerAutoHealingPoliciesOutput)
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput) HealthCheck() pulumi.StringOutput {
	return o.ApplyT(func (v RegionInstanceGroupManagerAutoHealingPolicies) string { return v.HealthCheck }).(pulumi.StringOutput)
}

func (o RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput) InitialDelaySec() pulumi.IntOutput {
	return o.ApplyT(func (v RegionInstanceGroupManagerAutoHealingPolicies) int { return v.InitialDelaySec }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionInstanceGroupManagerAutoHealingPoliciesOutput{})
	pulumi.RegisterOutputType(RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput{})
}
