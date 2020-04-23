// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtasks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type QueueAppEngineRoutingOverride struct {
	// -
	// The host that the task is sent to.
	Host *string `pulumi:"host"`
	// -
	// (Optional)
	// App instance.
	// By default, the task is sent to an instance which is available when the task is attempted.
	Instance *string `pulumi:"instance"`
	// -
	// (Optional)
	// App service.
	// By default, the task is sent to the service which is the default service when the task is attempted.
	Service *string `pulumi:"service"`
	// -
	// (Optional)
	// App version.
	// By default, the task is sent to the version which is the default version when the task is attempted.
	Version *string `pulumi:"version"`
}

// QueueAppEngineRoutingOverrideInput is an input type that accepts QueueAppEngineRoutingOverrideArgs and QueueAppEngineRoutingOverrideOutput values.
// You can construct a concrete instance of `QueueAppEngineRoutingOverrideInput` via:
//
// 		 QueueAppEngineRoutingOverrideArgs{...}
//
type QueueAppEngineRoutingOverrideInput interface {
	pulumi.Input

	ToQueueAppEngineRoutingOverrideOutput() QueueAppEngineRoutingOverrideOutput
	ToQueueAppEngineRoutingOverrideOutputWithContext(context.Context) QueueAppEngineRoutingOverrideOutput
}

type QueueAppEngineRoutingOverrideArgs struct {
	// -
	// The host that the task is sent to.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// -
	// (Optional)
	// App instance.
	// By default, the task is sent to an instance which is available when the task is attempted.
	Instance pulumi.StringPtrInput `pulumi:"instance"`
	// -
	// (Optional)
	// App service.
	// By default, the task is sent to the service which is the default service when the task is attempted.
	Service pulumi.StringPtrInput `pulumi:"service"`
	// -
	// (Optional)
	// App version.
	// By default, the task is sent to the version which is the default version when the task is attempted.
	Version pulumi.StringPtrInput `pulumi:"version"`
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

// QueueAppEngineRoutingOverridePtrInput is an input type that accepts QueueAppEngineRoutingOverrideArgs, QueueAppEngineRoutingOverridePtr and QueueAppEngineRoutingOverridePtrOutput values.
// You can construct a concrete instance of `QueueAppEngineRoutingOverridePtrInput` via:
//
// 		 QueueAppEngineRoutingOverrideArgs{...}
//
//  or:
//
// 		 nil
//
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

// -
// The host that the task is sent to.
func (o QueueAppEngineRoutingOverrideOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// App instance.
// By default, the task is sent to an instance which is available when the task is attempted.
func (o QueueAppEngineRoutingOverrideOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// App service.
// By default, the task is sent to the service which is the default service when the task is attempted.
func (o QueueAppEngineRoutingOverrideOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueAppEngineRoutingOverride) *string { return v.Service }).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// App version.
// By default, the task is sent to the version which is the default version when the task is attempted.
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

// -
// The host that the task is sent to.
func (o QueueAppEngineRoutingOverridePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAppEngineRoutingOverride) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// App instance.
// By default, the task is sent to an instance which is available when the task is attempted.
func (o QueueAppEngineRoutingOverridePtrOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAppEngineRoutingOverride) *string {
		if v == nil {
			return nil
		}
		return v.Instance
	}).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// App service.
// By default, the task is sent to the service which is the default service when the task is attempted.
func (o QueueAppEngineRoutingOverridePtrOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAppEngineRoutingOverride) *string {
		if v == nil {
			return nil
		}
		return v.Service
	}).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// App version.
// By default, the task is sent to the version which is the default version when the task is attempted.
func (o QueueAppEngineRoutingOverridePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAppEngineRoutingOverride) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type QueueRateLimits struct {
	// -
	// The max burst size.
	// Max burst size limits how fast tasks in queue are processed when many tasks are
	// in the queue and the rate is high. This field allows the queue to have a high
	// rate so processing starts shortly after a task is enqueued, but still limits
	// resource usage when many tasks are enqueued in a short period of time.
	MaxBurstSize *int `pulumi:"maxBurstSize"`
	// -
	// (Optional)
	// The maximum number of concurrent tasks that Cloud Tasks allows to
	// be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	MaxConcurrentDispatches *int `pulumi:"maxConcurrentDispatches"`
	// -
	// (Optional)
	// The maximum rate at which tasks are dispatched from this queue.
	// If unspecified when the queue is created, Cloud Tasks will pick the default.
	MaxDispatchesPerSecond *float64 `pulumi:"maxDispatchesPerSecond"`
}

// QueueRateLimitsInput is an input type that accepts QueueRateLimitsArgs and QueueRateLimitsOutput values.
// You can construct a concrete instance of `QueueRateLimitsInput` via:
//
// 		 QueueRateLimitsArgs{...}
//
type QueueRateLimitsInput interface {
	pulumi.Input

	ToQueueRateLimitsOutput() QueueRateLimitsOutput
	ToQueueRateLimitsOutputWithContext(context.Context) QueueRateLimitsOutput
}

type QueueRateLimitsArgs struct {
	// -
	// The max burst size.
	// Max burst size limits how fast tasks in queue are processed when many tasks are
	// in the queue and the rate is high. This field allows the queue to have a high
	// rate so processing starts shortly after a task is enqueued, but still limits
	// resource usage when many tasks are enqueued in a short period of time.
	MaxBurstSize pulumi.IntPtrInput `pulumi:"maxBurstSize"`
	// -
	// (Optional)
	// The maximum number of concurrent tasks that Cloud Tasks allows to
	// be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	MaxConcurrentDispatches pulumi.IntPtrInput `pulumi:"maxConcurrentDispatches"`
	// -
	// (Optional)
	// The maximum rate at which tasks are dispatched from this queue.
	// If unspecified when the queue is created, Cloud Tasks will pick the default.
	MaxDispatchesPerSecond pulumi.Float64PtrInput `pulumi:"maxDispatchesPerSecond"`
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

// QueueRateLimitsPtrInput is an input type that accepts QueueRateLimitsArgs, QueueRateLimitsPtr and QueueRateLimitsPtrOutput values.
// You can construct a concrete instance of `QueueRateLimitsPtrInput` via:
//
// 		 QueueRateLimitsArgs{...}
//
//  or:
//
// 		 nil
//
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

// -
// The max burst size.
// Max burst size limits how fast tasks in queue are processed when many tasks are
// in the queue and the rate is high. This field allows the queue to have a high
// rate so processing starts shortly after a task is enqueued, but still limits
// resource usage when many tasks are enqueued in a short period of time.
func (o QueueRateLimitsOutput) MaxBurstSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *int { return v.MaxBurstSize }).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// The maximum number of concurrent tasks that Cloud Tasks allows to
// be dispatched for this queue. After this threshold has been
// reached, Cloud Tasks stops dispatching tasks until the number of
// concurrent requests decreases.
func (o QueueRateLimitsOutput) MaxConcurrentDispatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRateLimits) *int { return v.MaxConcurrentDispatches }).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// The maximum rate at which tasks are dispatched from this queue.
// If unspecified when the queue is created, Cloud Tasks will pick the default.
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

// -
// The max burst size.
// Max burst size limits how fast tasks in queue are processed when many tasks are
// in the queue and the rate is high. This field allows the queue to have a high
// rate so processing starts shortly after a task is enqueued, but still limits
// resource usage when many tasks are enqueued in a short period of time.
func (o QueueRateLimitsPtrOutput) MaxBurstSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueRateLimits) *int {
		if v == nil {
			return nil
		}
		return v.MaxBurstSize
	}).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// The maximum number of concurrent tasks that Cloud Tasks allows to
// be dispatched for this queue. After this threshold has been
// reached, Cloud Tasks stops dispatching tasks until the number of
// concurrent requests decreases.
func (o QueueRateLimitsPtrOutput) MaxConcurrentDispatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueRateLimits) *int {
		if v == nil {
			return nil
		}
		return v.MaxConcurrentDispatches
	}).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// The maximum rate at which tasks are dispatched from this queue.
// If unspecified when the queue is created, Cloud Tasks will pick the default.
func (o QueueRateLimitsPtrOutput) MaxDispatchesPerSecond() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *QueueRateLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxDispatchesPerSecond
	}).(pulumi.Float64PtrOutput)
}

type QueueRetryConfig struct {
	// -
	// (Optional)
	// Number of attempts per task.
	// Cloud Tasks will attempt the task maxAttempts times (that is, if
	// the first attempt fails, then there will be maxAttempts - 1
	// retries). Must be >= -1.
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the default.
	// -1 indicates unlimited attempts.
	MaxAttempts *int `pulumi:"maxAttempts"`
	// -
	// (Optional)
	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	MaxBackoff *string `pulumi:"maxBackoff"`
	// -
	// (Optional)
	// The time between retries will double maxDoublings times.
	// A task's retry interval starts at minBackoff, then doubles maxDoublings times,
	// then increases linearly, and finally retries retries at intervals of maxBackoff
	// up to maxAttempts times.
	MaxDoublings *int `pulumi:"maxDoublings"`
	// -
	// (Optional)
	// If positive, maxRetryDuration specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once maxRetryDuration time has passed and the task has
	// been attempted maxAttempts times, no further attempts will be
	// made and the task will be deleted.
	// If zero, then the task age is unlimited.
	MaxRetryDuration *string `pulumi:"maxRetryDuration"`
	// -
	// (Optional)
	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	MinBackoff *string `pulumi:"minBackoff"`
}

// QueueRetryConfigInput is an input type that accepts QueueRetryConfigArgs and QueueRetryConfigOutput values.
// You can construct a concrete instance of `QueueRetryConfigInput` via:
//
// 		 QueueRetryConfigArgs{...}
//
type QueueRetryConfigInput interface {
	pulumi.Input

	ToQueueRetryConfigOutput() QueueRetryConfigOutput
	ToQueueRetryConfigOutputWithContext(context.Context) QueueRetryConfigOutput
}

type QueueRetryConfigArgs struct {
	// -
	// (Optional)
	// Number of attempts per task.
	// Cloud Tasks will attempt the task maxAttempts times (that is, if
	// the first attempt fails, then there will be maxAttempts - 1
	// retries). Must be >= -1.
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the default.
	// -1 indicates unlimited attempts.
	MaxAttempts pulumi.IntPtrInput `pulumi:"maxAttempts"`
	// -
	// (Optional)
	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	MaxBackoff pulumi.StringPtrInput `pulumi:"maxBackoff"`
	// -
	// (Optional)
	// The time between retries will double maxDoublings times.
	// A task's retry interval starts at minBackoff, then doubles maxDoublings times,
	// then increases linearly, and finally retries retries at intervals of maxBackoff
	// up to maxAttempts times.
	MaxDoublings pulumi.IntPtrInput `pulumi:"maxDoublings"`
	// -
	// (Optional)
	// If positive, maxRetryDuration specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once maxRetryDuration time has passed and the task has
	// been attempted maxAttempts times, no further attempts will be
	// made and the task will be deleted.
	// If zero, then the task age is unlimited.
	MaxRetryDuration pulumi.StringPtrInput `pulumi:"maxRetryDuration"`
	// -
	// (Optional)
	// A task will be scheduled for retry between minBackoff and
	// maxBackoff duration after it fails, if the queue's RetryConfig
	// specifies that the task should be retried.
	MinBackoff pulumi.StringPtrInput `pulumi:"minBackoff"`
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

// QueueRetryConfigPtrInput is an input type that accepts QueueRetryConfigArgs, QueueRetryConfigPtr and QueueRetryConfigPtrOutput values.
// You can construct a concrete instance of `QueueRetryConfigPtrInput` via:
//
// 		 QueueRetryConfigArgs{...}
//
//  or:
//
// 		 nil
//
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

// -
// (Optional)
// Number of attempts per task.
// Cloud Tasks will attempt the task maxAttempts times (that is, if
// the first attempt fails, then there will be maxAttempts - 1
// retries). Must be >= -1.
// If unspecified when the queue is created, Cloud Tasks will pick
// the default.
// -1 indicates unlimited attempts.
func (o QueueRetryConfigOutput) MaxAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *int { return v.MaxAttempts }).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// A task will be scheduled for retry between minBackoff and
// maxBackoff duration after it fails, if the queue's RetryConfig
// specifies that the task should be retried.
func (o QueueRetryConfigOutput) MaxBackoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MaxBackoff }).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// The time between retries will double maxDoublings times.
// A task's retry interval starts at minBackoff, then doubles maxDoublings times,
// then increases linearly, and finally retries retries at intervals of maxBackoff
// up to maxAttempts times.
func (o QueueRetryConfigOutput) MaxDoublings() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *int { return v.MaxDoublings }).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// If positive, maxRetryDuration specifies the time limit for
// retrying a failed task, measured from when the task was first
// attempted. Once maxRetryDuration time has passed and the task has
// been attempted maxAttempts times, no further attempts will be
// made and the task will be deleted.
// If zero, then the task age is unlimited.
func (o QueueRetryConfigOutput) MaxRetryDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueRetryConfig) *string { return v.MaxRetryDuration }).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// A task will be scheduled for retry between minBackoff and
// maxBackoff duration after it fails, if the queue's RetryConfig
// specifies that the task should be retried.
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

// -
// (Optional)
// Number of attempts per task.
// Cloud Tasks will attempt the task maxAttempts times (that is, if
// the first attempt fails, then there will be maxAttempts - 1
// retries). Must be >= -1.
// If unspecified when the queue is created, Cloud Tasks will pick
// the default.
// -1 indicates unlimited attempts.
func (o QueueRetryConfigPtrOutput) MaxAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueRetryConfig) *int {
		if v == nil {
			return nil
		}
		return v.MaxAttempts
	}).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// A task will be scheduled for retry between minBackoff and
// maxBackoff duration after it fails, if the queue's RetryConfig
// specifies that the task should be retried.
func (o QueueRetryConfigPtrOutput) MaxBackoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueRetryConfig) *string {
		if v == nil {
			return nil
		}
		return v.MaxBackoff
	}).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// The time between retries will double maxDoublings times.
// A task's retry interval starts at minBackoff, then doubles maxDoublings times,
// then increases linearly, and finally retries retries at intervals of maxBackoff
// up to maxAttempts times.
func (o QueueRetryConfigPtrOutput) MaxDoublings() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueRetryConfig) *int {
		if v == nil {
			return nil
		}
		return v.MaxDoublings
	}).(pulumi.IntPtrOutput)
}

// -
// (Optional)
// If positive, maxRetryDuration specifies the time limit for
// retrying a failed task, measured from when the task was first
// attempted. Once maxRetryDuration time has passed and the task has
// been attempted maxAttempts times, no further attempts will be
// made and the task will be deleted.
// If zero, then the task age is unlimited.
func (o QueueRetryConfigPtrOutput) MaxRetryDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueRetryConfig) *string {
		if v == nil {
			return nil
		}
		return v.MaxRetryDuration
	}).(pulumi.StringPtrOutput)
}

// -
// (Optional)
// A task will be scheduled for retry between minBackoff and
// maxBackoff duration after it fails, if the queue's RetryConfig
// specifies that the task should be retried.
func (o QueueRetryConfigPtrOutput) MinBackoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueRetryConfig) *string {
		if v == nil {
			return nil
		}
		return v.MinBackoff
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueAppEngineRoutingOverrideOutput{})
	pulumi.RegisterOutputType(QueueAppEngineRoutingOverridePtrOutput{})
	pulumi.RegisterOutputType(QueueRateLimitsOutput{})
	pulumi.RegisterOutputType(QueueRateLimitsPtrOutput{})
	pulumi.RegisterOutputType(QueueRetryConfigOutput{})
	pulumi.RegisterOutputType(QueueRetryConfigPtrOutput{})
}
