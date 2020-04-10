// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtasks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type QueueAppEngineRoutingOverride struct {
	Host     *string `pulumi:"host"`
	Instance *string `pulumi:"instance"`
	Service  *string `pulumi:"service"`
	Version  *string `pulumi:"version"`
}

type QueueAppEngineRoutingOverrideInput interface {
	pulumi.Input

	ToQueueAppEngineRoutingOverrideOutput() QueueAppEngineRoutingOverrideOutput
	ToQueueAppEngineRoutingOverrideOutputWithContext(context.Context) QueueAppEngineRoutingOverrideOutput
}

type QueueAppEngineRoutingOverrideArgs struct {
	Host     pulumi.StringPtrInput `pulumi:"host"`
	Instance pulumi.StringPtrInput `pulumi:"instance"`
	Service  pulumi.StringPtrInput `pulumi:"service"`
	Version  pulumi.StringPtrInput `pulumi:"version"`
}

func (QueueAppEngineRoutingOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueAppEngineRoutingOverride)(nil)).Elem()
}

func (i QueueAppEngineRoutingOverrideArgs) ToQueueAppEngineRoutingOverrideOutput() QueueAppEngineRoutingOverrideOutput {
	return i.ToQueueAppEngineRoutingOverrideOutputWithContext(context.Background())
}

func (i QueueAppEngineRoutingOverrideArgs) ToQueueAppEngineRoutingOverrideOutputWithContext(ctx context.Context) QueueAppEngineRoutingOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAppEngineRoutingOverrideOutput)
}

func (i QueueAppEngineRoutingOverrideArgs) ToQueueAppEngineRoutingOverridePtrOutput() QueueAppEngineRoutingOverridePtrOutput {
	return i.ToQueueAppEngineRoutingOverridePtrOutputWithContext(context.Background())
}

func (i QueueAppEngineRoutingOverrideArgs) ToQueueAppEngineRoutingOverridePtrOutputWithContext(ctx context.Context) QueueAppEngineRoutingOverridePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAppEngineRoutingOverrideOutput).ToQueueAppEngineRoutingOverridePtrOutputWithContext(ctx)
}

type QueueAppEngineRoutingOverridePtrInput interface {
	pulumi.Input

	ToQueueAppEngineRoutingOverridePtrOutput() QueueAppEngineRoutingOverridePtrOutput
	ToQueueAppEngineRoutingOverridePtrOutputWithContext(context.Context) QueueAppEngineRoutingOverridePtrOutput
}

type queueAppEngineRoutingOverridePtrType QueueAppEngineRoutingOverrideArgs

func QueueAppEngineRoutingOverridePtr(v *QueueAppEngineRoutingOverrideArgs) QueueAppEngineRoutingOverridePtrInput {
	return (*queueAppEngineRoutingOverridePtrType)(v)
}

func (*queueAppEngineRoutingOverridePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAppEngineRoutingOverride)(nil)).Elem()
}

func (i *queueAppEngineRoutingOverridePtrType) ToQueueAppEngineRoutingOverridePtrOutput() QueueAppEngineRoutingOverridePtrOutput {
	return i.ToQueueAppEngineRoutingOverridePtrOutputWithContext(context.Background())
}

func (i *queueAppEngineRoutingOverridePtrType) ToQueueAppEngineRoutingOverridePtrOutputWithContext(ctx context.Context) QueueAppEngineRoutingOverridePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAppEngineRoutingOverridePtrOutput)
}

type QueueAppEngineRoutingOverrideOutput struct{ *pulumi.OutputState }

func (QueueAppEngineRoutingOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueAppEngineRoutingOverride)(nil)).Elem()
}

func (o QueueAppEngineRoutingOverrideOutput) ToQueueAppEngineRoutingOverrideOutput() QueueAppEngineRoutingOverrideOutput {
	return o
}

func (o QueueAppEngineRoutingOverrideOutput) ToQueueAppEngineRoutingOverrideOutputWithContext(ctx context.Context) QueueAppEngineRoutingOverrideOutput {
	return o
}

func (o QueueAppEngineRoutingOverrideOutput) ToQueueAppEngineRoutingOverridePtrOutput() QueueAppEngineRoutingOverridePtrOutput {
	return o.ToQueueAppEngineRoutingOverridePtrOutputWithContext(context.Background())
}

func (o QueueAppEngineRoutingOverrideOutput) ToQueueAppEngineRoutingOverridePtrOutputWithContext(ctx context.Context) QueueAppEngineRoutingOverridePtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *QueueAppEngineRoutingOverride {
		return &v
	}).(QueueAppEngineRoutingOverridePtrOutput)
}
func (o QueueAppEngineRoutingOverrideOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o QueueAppEngineRoutingOverrideOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o QueueAppEngineRoutingOverrideOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o QueueAppEngineRoutingOverrideOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type QueueAppEngineRoutingOverridePtrOutput struct{ *pulumi.OutputState }

func (QueueAppEngineRoutingOverridePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAppEngineRoutingOverride)(nil)).Elem()
}

func (o QueueAppEngineRoutingOverridePtrOutput) ToQueueAppEngineRoutingOverridePtrOutput() QueueAppEngineRoutingOverridePtrOutput {
	return o
}

func (o QueueAppEngineRoutingOverridePtrOutput) ToQueueAppEngineRoutingOverridePtrOutputWithContext(ctx context.Context) QueueAppEngineRoutingOverridePtrOutput {
	return o
}

func (o QueueAppEngineRoutingOverridePtrOutput) Elem() QueueAppEngineRoutingOverrideOutput {
	return o.ApplyT(func(v *QueueAppEngineRoutingOverride) QueueAppEngineRoutingOverride { return *v }).(QueueAppEngineRoutingOverrideOutput)
}

func (o QueueAppEngineRoutingOverridePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o QueueAppEngineRoutingOverridePtrOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o QueueAppEngineRoutingOverridePtrOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o QueueAppEngineRoutingOverridePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type QueueRateLimits struct {
	MaxBurstSize            *int     `pulumi:"maxBurstSize"`
	MaxConcurrentDispatches *int     `pulumi:"maxConcurrentDispatches"`
	MaxDispatchesPerSecond  *float64 `pulumi:"maxDispatchesPerSecond"`
}

type QueueRateLimitsInput interface {
	pulumi.Input

	ToQueueRateLimitsOutput() QueueRateLimitsOutput
	ToQueueRateLimitsOutputWithContext(context.Context) QueueRateLimitsOutput
}

type QueueRateLimitsArgs struct {
	MaxBurstSize            pulumi.IntPtrInput     `pulumi:"maxBurstSize"`
	MaxConcurrentDispatches pulumi.IntPtrInput     `pulumi:"maxConcurrentDispatches"`
	MaxDispatchesPerSecond  pulumi.Float64PtrInput `pulumi:"maxDispatchesPerSecond"`
}

func (QueueRateLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueRateLimits)(nil)).Elem()
}

func (i QueueRateLimitsArgs) ToQueueRateLimitsOutput() QueueRateLimitsOutput {
	return i.ToQueueRateLimitsOutputWithContext(context.Background())
}

func (i QueueRateLimitsArgs) ToQueueRateLimitsOutputWithContext(ctx context.Context) QueueRateLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueRateLimitsOutput)
}

func (i QueueRateLimitsArgs) ToQueueRateLimitsPtrOutput() QueueRateLimitsPtrOutput {
	return i.ToQueueRateLimitsPtrOutputWithContext(context.Background())
}

func (i QueueRateLimitsArgs) ToQueueRateLimitsPtrOutputWithContext(ctx context.Context) QueueRateLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueRateLimitsOutput).ToQueueRateLimitsPtrOutputWithContext(ctx)
}

type QueueRateLimitsPtrInput interface {
	pulumi.Input

	ToQueueRateLimitsPtrOutput() QueueRateLimitsPtrOutput
	ToQueueRateLimitsPtrOutputWithContext(context.Context) QueueRateLimitsPtrOutput
}

type queueRateLimitsPtrType QueueRateLimitsArgs

func QueueRateLimitsPtr(v *QueueRateLimitsArgs) QueueRateLimitsPtrInput {
	return (*queueRateLimitsPtrType)(v)
}

func (*queueRateLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueRateLimits)(nil)).Elem()
}

func (i *queueRateLimitsPtrType) ToQueueRateLimitsPtrOutput() QueueRateLimitsPtrOutput {
	return i.ToQueueRateLimitsPtrOutputWithContext(context.Background())
}

func (i *queueRateLimitsPtrType) ToQueueRateLimitsPtrOutputWithContext(ctx context.Context) QueueRateLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueRateLimitsPtrOutput)
}

type QueueRateLimitsOutput struct{ *pulumi.OutputState }

func (QueueRateLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueRateLimits)(nil)).Elem()
}

func (o QueueRateLimitsOutput) ToQueueRateLimitsOutput() QueueRateLimitsOutput {
	return o
}

func (o QueueRateLimitsOutput) ToQueueRateLimitsOutputWithContext(ctx context.Context) QueueRateLimitsOutput {
	return o
}

func (o QueueRateLimitsOutput) ToQueueRateLimitsPtrOutput() QueueRateLimitsPtrOutput {
	return o.ToQueueRateLimitsPtrOutputWithContext(context.Background())
}

func (o QueueRateLimitsOutput) ToQueueRateLimitsPtrOutputWithContext(ctx context.Context) QueueRateLimitsPtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *QueueRateLimits {
		return &v
	}).(QueueRateLimitsPtrOutput)
}
func (o QueueRateLimitsOutput) MaxBurstSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *int { return v.MaxBurstSize }).(pulumi.IntPtrOutput)
}

func (o QueueRateLimitsOutput) MaxConcurrentDispatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *int { return v.MaxConcurrentDispatches }).(pulumi.IntPtrOutput)
}

func (o QueueRateLimitsOutput) MaxDispatchesPerSecond() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *float64 { return v.MaxDispatchesPerSecond }).(pulumi.Float64PtrOutput)
}

type QueueRateLimitsPtrOutput struct{ *pulumi.OutputState }

func (QueueRateLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueRateLimits)(nil)).Elem()
}

func (o QueueRateLimitsPtrOutput) ToQueueRateLimitsPtrOutput() QueueRateLimitsPtrOutput {
	return o
}

func (o QueueRateLimitsPtrOutput) ToQueueRateLimitsPtrOutputWithContext(ctx context.Context) QueueRateLimitsPtrOutput {
	return o
}

func (o QueueRateLimitsPtrOutput) Elem() QueueRateLimitsOutput {
	return o.ApplyT(func(v *QueueRateLimits) QueueRateLimits { return *v }).(QueueRateLimitsOutput)
}

func (o QueueRateLimitsPtrOutput) MaxBurstSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *int { return v.MaxBurstSize }).(pulumi.IntPtrOutput)
}

func (o QueueRateLimitsPtrOutput) MaxConcurrentDispatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *int { return v.MaxConcurrentDispatches }).(pulumi.IntPtrOutput)
}

func (o QueueRateLimitsPtrOutput) MaxDispatchesPerSecond() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *float64 { return v.MaxDispatchesPerSecond }).(pulumi.Float64PtrOutput)
}

type QueueRetryConfig struct {
	MaxAttempts      *int    `pulumi:"maxAttempts"`
	MaxBackoff       *string `pulumi:"maxBackoff"`
	MaxDoublings     *int    `pulumi:"maxDoublings"`
	MaxRetryDuration *string `pulumi:"maxRetryDuration"`
	MinBackoff       *string `pulumi:"minBackoff"`
}

type QueueRetryConfigInput interface {
	pulumi.Input

	ToQueueRetryConfigOutput() QueueRetryConfigOutput
	ToQueueRetryConfigOutputWithContext(context.Context) QueueRetryConfigOutput
}

type QueueRetryConfigArgs struct {
	MaxAttempts      pulumi.IntPtrInput    `pulumi:"maxAttempts"`
	MaxBackoff       pulumi.StringPtrInput `pulumi:"maxBackoff"`
	MaxDoublings     pulumi.IntPtrInput    `pulumi:"maxDoublings"`
	MaxRetryDuration pulumi.StringPtrInput `pulumi:"maxRetryDuration"`
	MinBackoff       pulumi.StringPtrInput `pulumi:"minBackoff"`
}

func (QueueRetryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueRetryConfig)(nil)).Elem()
}

func (i QueueRetryConfigArgs) ToQueueRetryConfigOutput() QueueRetryConfigOutput {
	return i.ToQueueRetryConfigOutputWithContext(context.Background())
}

func (i QueueRetryConfigArgs) ToQueueRetryConfigOutputWithContext(ctx context.Context) QueueRetryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueRetryConfigOutput)
}

func (i QueueRetryConfigArgs) ToQueueRetryConfigPtrOutput() QueueRetryConfigPtrOutput {
	return i.ToQueueRetryConfigPtrOutputWithContext(context.Background())
}

func (i QueueRetryConfigArgs) ToQueueRetryConfigPtrOutputWithContext(ctx context.Context) QueueRetryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueRetryConfigOutput).ToQueueRetryConfigPtrOutputWithContext(ctx)
}

type QueueRetryConfigPtrInput interface {
	pulumi.Input

	ToQueueRetryConfigPtrOutput() QueueRetryConfigPtrOutput
	ToQueueRetryConfigPtrOutputWithContext(context.Context) QueueRetryConfigPtrOutput
}

type queueRetryConfigPtrType QueueRetryConfigArgs

func QueueRetryConfigPtr(v *QueueRetryConfigArgs) QueueRetryConfigPtrInput {
	return (*queueRetryConfigPtrType)(v)
}

func (*queueRetryConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueRetryConfig)(nil)).Elem()
}

func (i *queueRetryConfigPtrType) ToQueueRetryConfigPtrOutput() QueueRetryConfigPtrOutput {
	return i.ToQueueRetryConfigPtrOutputWithContext(context.Background())
}

func (i *queueRetryConfigPtrType) ToQueueRetryConfigPtrOutputWithContext(ctx context.Context) QueueRetryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueRetryConfigPtrOutput)
}

type QueueRetryConfigOutput struct{ *pulumi.OutputState }

func (QueueRetryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueRetryConfig)(nil)).Elem()
}

func (o QueueRetryConfigOutput) ToQueueRetryConfigOutput() QueueRetryConfigOutput {
	return o
}

func (o QueueRetryConfigOutput) ToQueueRetryConfigOutputWithContext(ctx context.Context) QueueRetryConfigOutput {
	return o
}

func (o QueueRetryConfigOutput) ToQueueRetryConfigPtrOutput() QueueRetryConfigPtrOutput {
	return o.ToQueueRetryConfigPtrOutputWithContext(context.Background())
}

func (o QueueRetryConfigOutput) ToQueueRetryConfigPtrOutputWithContext(ctx context.Context) QueueRetryConfigPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *QueueRetryConfig {
		return &v
	}).(QueueRetryConfigPtrOutput)
}
func (o QueueRetryConfigOutput) MaxAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *int { return v.MaxAttempts }).(pulumi.IntPtrOutput)
}

func (o QueueRetryConfigOutput) MaxBackoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MaxBackoff }).(pulumi.StringPtrOutput)
}

func (o QueueRetryConfigOutput) MaxDoublings() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *int { return v.MaxDoublings }).(pulumi.IntPtrOutput)
}

func (o QueueRetryConfigOutput) MaxRetryDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MaxRetryDuration }).(pulumi.StringPtrOutput)
}

func (o QueueRetryConfigOutput) MinBackoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MinBackoff }).(pulumi.StringPtrOutput)
}

type QueueRetryConfigPtrOutput struct{ *pulumi.OutputState }

func (QueueRetryConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueRetryConfig)(nil)).Elem()
}

func (o QueueRetryConfigPtrOutput) ToQueueRetryConfigPtrOutput() QueueRetryConfigPtrOutput {
	return o
}

func (o QueueRetryConfigPtrOutput) ToQueueRetryConfigPtrOutputWithContext(ctx context.Context) QueueRetryConfigPtrOutput {
	return o
}

func (o QueueRetryConfigPtrOutput) Elem() QueueRetryConfigOutput {
	return o.ApplyT(func(v *QueueRetryConfig) QueueRetryConfig { return *v }).(QueueRetryConfigOutput)
}

func (o QueueRetryConfigPtrOutput) MaxAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *int { return v.MaxAttempts }).(pulumi.IntPtrOutput)
}

func (o QueueRetryConfigPtrOutput) MaxBackoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MaxBackoff }).(pulumi.StringPtrOutput)
}

func (o QueueRetryConfigPtrOutput) MaxDoublings() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *int { return v.MaxDoublings }).(pulumi.IntPtrOutput)
}

func (o QueueRetryConfigPtrOutput) MaxRetryDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MaxRetryDuration }).(pulumi.StringPtrOutput)
}

func (o QueueRetryConfigPtrOutput) MinBackoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MinBackoff }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueAppEngineRoutingOverrideOutput{})
	pulumi.RegisterOutputType(QueueAppEngineRoutingOverridePtrOutput{})
	pulumi.RegisterOutputType(QueueRateLimitsOutput{})
	pulumi.RegisterOutputType(QueueRateLimitsPtrOutput{})
	pulumi.RegisterOutputType(QueueRetryConfigOutput{})
	pulumi.RegisterOutputType(QueueRetryConfigPtrOutput{})
}
