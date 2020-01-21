// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package AutoscalingPolicyWorkerConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AutoscalingPolicyWorkerConfig struct {
	MaxInstances int `pulumi:"maxInstances"`
	MinInstances *int `pulumi:"minInstances"`
	Weight *int `pulumi:"weight"`
}

type AutoscalingPolicyWorkerConfigInput interface {
	pulumi.Input

	ToAutoscalingPolicyWorkerConfigOutput() AutoscalingPolicyWorkerConfigOutput
	ToAutoscalingPolicyWorkerConfigOutputWithContext(context.Context) AutoscalingPolicyWorkerConfigOutput
}

type AutoscalingPolicyWorkerConfigArgs struct {
	MaxInstances pulumi.IntInput `pulumi:"maxInstances"`
	MinInstances pulumi.IntPtrInput `pulumi:"minInstances"`
	Weight pulumi.IntPtrInput `pulumi:"weight"`
}

func (AutoscalingPolicyWorkerConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscalingPolicyWorkerConfig)(nil)).Elem()
}

func (i AutoscalingPolicyWorkerConfigArgs) ToAutoscalingPolicyWorkerConfigOutput() AutoscalingPolicyWorkerConfigOutput {
	return i.ToAutoscalingPolicyWorkerConfigOutputWithContext(context.Background())
}

func (i AutoscalingPolicyWorkerConfigArgs) ToAutoscalingPolicyWorkerConfigOutputWithContext(ctx context.Context) AutoscalingPolicyWorkerConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyWorkerConfigOutput)
}

func (i AutoscalingPolicyWorkerConfigArgs) ToAutoscalingPolicyWorkerConfigPtrOutput() AutoscalingPolicyWorkerConfigPtrOutput {
	return i.ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(context.Background())
}

func (i AutoscalingPolicyWorkerConfigArgs) ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(ctx context.Context) AutoscalingPolicyWorkerConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyWorkerConfigOutput).ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(ctx)
}

type AutoscalingPolicyWorkerConfigPtrInput interface {
	pulumi.Input

	ToAutoscalingPolicyWorkerConfigPtrOutput() AutoscalingPolicyWorkerConfigPtrOutput
	ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(context.Context) AutoscalingPolicyWorkerConfigPtrOutput
}

type autoscalingPolicyWorkerConfigPtrType AutoscalingPolicyWorkerConfigArgs

func AutoscalingPolicyWorkerConfigPtr(v *AutoscalingPolicyWorkerConfigArgs) AutoscalingPolicyWorkerConfigPtrInput {	return (*autoscalingPolicyWorkerConfigPtrType)(v)
}

func (*autoscalingPolicyWorkerConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscalingPolicyWorkerConfig)(nil)).Elem()
}

func (i *autoscalingPolicyWorkerConfigPtrType) ToAutoscalingPolicyWorkerConfigPtrOutput() AutoscalingPolicyWorkerConfigPtrOutput {
	return i.ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(context.Background())
}

func (i *autoscalingPolicyWorkerConfigPtrType) ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(ctx context.Context) AutoscalingPolicyWorkerConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyWorkerConfigPtrOutput)
}

type AutoscalingPolicyWorkerConfigOutput struct { *pulumi.OutputState }

func (AutoscalingPolicyWorkerConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscalingPolicyWorkerConfig)(nil)).Elem()
}

func (o AutoscalingPolicyWorkerConfigOutput) ToAutoscalingPolicyWorkerConfigOutput() AutoscalingPolicyWorkerConfigOutput {
	return o
}

func (o AutoscalingPolicyWorkerConfigOutput) ToAutoscalingPolicyWorkerConfigOutputWithContext(ctx context.Context) AutoscalingPolicyWorkerConfigOutput {
	return o
}

func (o AutoscalingPolicyWorkerConfigOutput) ToAutoscalingPolicyWorkerConfigPtrOutput() AutoscalingPolicyWorkerConfigPtrOutput {
	return o.ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(context.Background())
}

func (o AutoscalingPolicyWorkerConfigOutput) ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(ctx context.Context) AutoscalingPolicyWorkerConfigPtrOutput {
	return o.ApplyT(func(v AutoscalingPolicyWorkerConfig) *AutoscalingPolicyWorkerConfig {
		return &v
	}).(AutoscalingPolicyWorkerConfigPtrOutput)
}
func (o AutoscalingPolicyWorkerConfigOutput) MaxInstances() pulumi.IntOutput {
	return o.ApplyT(func (v AutoscalingPolicyWorkerConfig) int { return v.MaxInstances }).(pulumi.IntOutput)
}

func (o AutoscalingPolicyWorkerConfigOutput) MinInstances() pulumi.IntPtrOutput {
	return o.ApplyT(func (v AutoscalingPolicyWorkerConfig) *int { return v.MinInstances }).(pulumi.IntPtrOutput)
}

func (o AutoscalingPolicyWorkerConfigOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func (v AutoscalingPolicyWorkerConfig) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type AutoscalingPolicyWorkerConfigPtrOutput struct { *pulumi.OutputState}

func (AutoscalingPolicyWorkerConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscalingPolicyWorkerConfig)(nil)).Elem()
}

func (o AutoscalingPolicyWorkerConfigPtrOutput) ToAutoscalingPolicyWorkerConfigPtrOutput() AutoscalingPolicyWorkerConfigPtrOutput {
	return o
}

func (o AutoscalingPolicyWorkerConfigPtrOutput) ToAutoscalingPolicyWorkerConfigPtrOutputWithContext(ctx context.Context) AutoscalingPolicyWorkerConfigPtrOutput {
	return o
}

func (o AutoscalingPolicyWorkerConfigPtrOutput) Elem() AutoscalingPolicyWorkerConfigOutput {
	return o.ApplyT(func (v *AutoscalingPolicyWorkerConfig) AutoscalingPolicyWorkerConfig { return *v }).(AutoscalingPolicyWorkerConfigOutput)
}

func (o AutoscalingPolicyWorkerConfigPtrOutput) MaxInstances() pulumi.IntOutput {
	return o.ApplyT(func (v AutoscalingPolicyWorkerConfig) int { return v.MaxInstances }).(pulumi.IntOutput)
}

func (o AutoscalingPolicyWorkerConfigPtrOutput) MinInstances() pulumi.IntPtrOutput {
	return o.ApplyT(func (v AutoscalingPolicyWorkerConfig) *int { return v.MinInstances }).(pulumi.IntPtrOutput)
}

func (o AutoscalingPolicyWorkerConfigPtrOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func (v AutoscalingPolicyWorkerConfig) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoscalingPolicyWorkerConfigOutput{})
	pulumi.RegisterOutputType(AutoscalingPolicyWorkerConfigPtrOutput{})
}
