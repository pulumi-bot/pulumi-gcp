// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization struct {
	Target float64 `pulumi:"target"`
}

type RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationInput interface {
	pulumi.Input

	ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput
	ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputWithContext(context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput
}

type RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs struct {
	Target pulumi.Float64Input `pulumi:"target"`
}

func (RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization)(nil)).Elem()
}

func (i RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput {
	return i.ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputWithContext(context.Background())
}

func (i RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputWithContext(ctx context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput)
}

func (i RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return i.ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(context.Background())
}

func (i RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(ctx context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput).ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(ctx)
}

type RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrInput interface {
	pulumi.Input

	ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput
	ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput
}

type regionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrType RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs

func RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtr(v *RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationArgs) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrInput {	return (*regionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrType)(v)
}

func (*regionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization)(nil)).Elem()
}

func (i *regionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrType) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return i.ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(context.Background())
}

func (i *regionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrType) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(ctx context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput)
}

type RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput struct { *pulumi.OutputState }

func (RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization)(nil)).Elem()
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput {
	return o
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputWithContext(ctx context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput {
	return o
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return o.ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(context.Background())
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(ctx context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return o.ApplyT(func(v RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization) *RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization {
		return &v
	}).(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput)
}
func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput) Target() pulumi.Float64Output {
	return o.ApplyT(func (v RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization) float64 { return v.Target }).(pulumi.Float64Output)
}

type RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput struct { *pulumi.OutputState}

func (RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization)(nil)).Elem()
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return o
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput) ToRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutputWithContext(ctx context.Context) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput {
	return o
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput) Elem() RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput {
	return o.ApplyT(func (v *RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization) RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization { return *v }).(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput)
}

func (o RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput) Target() pulumi.Float64Output {
	return o.ApplyT(func (v RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization) float64 { return v.Target }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutput{})
	pulumi.RegisterOutputType(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationPtrOutput{})
}
