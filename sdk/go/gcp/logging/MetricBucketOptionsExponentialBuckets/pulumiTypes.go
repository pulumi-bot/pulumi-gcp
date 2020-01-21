// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package MetricBucketOptionsExponentialBuckets

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type MetricBucketOptionsExponentialBuckets struct {
	GrowthFactor *float64 `pulumi:"growthFactor"`
	NumFiniteBuckets *int `pulumi:"numFiniteBuckets"`
	Scale *float64 `pulumi:"scale"`
}

type MetricBucketOptionsExponentialBucketsInput interface {
	pulumi.Input

	ToMetricBucketOptionsExponentialBucketsOutput() MetricBucketOptionsExponentialBucketsOutput
	ToMetricBucketOptionsExponentialBucketsOutputWithContext(context.Context) MetricBucketOptionsExponentialBucketsOutput
}

type MetricBucketOptionsExponentialBucketsArgs struct {
	GrowthFactor pulumi.Float64PtrInput `pulumi:"growthFactor"`
	NumFiniteBuckets pulumi.IntPtrInput `pulumi:"numFiniteBuckets"`
	Scale pulumi.Float64PtrInput `pulumi:"scale"`
}

func (MetricBucketOptionsExponentialBucketsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricBucketOptionsExponentialBuckets)(nil)).Elem()
}

func (i MetricBucketOptionsExponentialBucketsArgs) ToMetricBucketOptionsExponentialBucketsOutput() MetricBucketOptionsExponentialBucketsOutput {
	return i.ToMetricBucketOptionsExponentialBucketsOutputWithContext(context.Background())
}

func (i MetricBucketOptionsExponentialBucketsArgs) ToMetricBucketOptionsExponentialBucketsOutputWithContext(ctx context.Context) MetricBucketOptionsExponentialBucketsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricBucketOptionsExponentialBucketsOutput)
}

func (i MetricBucketOptionsExponentialBucketsArgs) ToMetricBucketOptionsExponentialBucketsPtrOutput() MetricBucketOptionsExponentialBucketsPtrOutput {
	return i.ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(context.Background())
}

func (i MetricBucketOptionsExponentialBucketsArgs) ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(ctx context.Context) MetricBucketOptionsExponentialBucketsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricBucketOptionsExponentialBucketsOutput).ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(ctx)
}

type MetricBucketOptionsExponentialBucketsPtrInput interface {
	pulumi.Input

	ToMetricBucketOptionsExponentialBucketsPtrOutput() MetricBucketOptionsExponentialBucketsPtrOutput
	ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(context.Context) MetricBucketOptionsExponentialBucketsPtrOutput
}

type metricBucketOptionsExponentialBucketsPtrType MetricBucketOptionsExponentialBucketsArgs

func MetricBucketOptionsExponentialBucketsPtr(v *MetricBucketOptionsExponentialBucketsArgs) MetricBucketOptionsExponentialBucketsPtrInput {	return (*metricBucketOptionsExponentialBucketsPtrType)(v)
}

func (*metricBucketOptionsExponentialBucketsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricBucketOptionsExponentialBuckets)(nil)).Elem()
}

func (i *metricBucketOptionsExponentialBucketsPtrType) ToMetricBucketOptionsExponentialBucketsPtrOutput() MetricBucketOptionsExponentialBucketsPtrOutput {
	return i.ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(context.Background())
}

func (i *metricBucketOptionsExponentialBucketsPtrType) ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(ctx context.Context) MetricBucketOptionsExponentialBucketsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricBucketOptionsExponentialBucketsPtrOutput)
}

type MetricBucketOptionsExponentialBucketsOutput struct { *pulumi.OutputState }

func (MetricBucketOptionsExponentialBucketsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricBucketOptionsExponentialBuckets)(nil)).Elem()
}

func (o MetricBucketOptionsExponentialBucketsOutput) ToMetricBucketOptionsExponentialBucketsOutput() MetricBucketOptionsExponentialBucketsOutput {
	return o
}

func (o MetricBucketOptionsExponentialBucketsOutput) ToMetricBucketOptionsExponentialBucketsOutputWithContext(ctx context.Context) MetricBucketOptionsExponentialBucketsOutput {
	return o
}

func (o MetricBucketOptionsExponentialBucketsOutput) ToMetricBucketOptionsExponentialBucketsPtrOutput() MetricBucketOptionsExponentialBucketsPtrOutput {
	return o.ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(context.Background())
}

func (o MetricBucketOptionsExponentialBucketsOutput) ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(ctx context.Context) MetricBucketOptionsExponentialBucketsPtrOutput {
	return o.ApplyT(func(v MetricBucketOptionsExponentialBuckets) *MetricBucketOptionsExponentialBuckets {
		return &v
	}).(MetricBucketOptionsExponentialBucketsPtrOutput)
}
func (o MetricBucketOptionsExponentialBucketsOutput) GrowthFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func (v MetricBucketOptionsExponentialBuckets) *float64 { return v.GrowthFactor }).(pulumi.Float64PtrOutput)
}

func (o MetricBucketOptionsExponentialBucketsOutput) NumFiniteBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func (v MetricBucketOptionsExponentialBuckets) *int { return v.NumFiniteBuckets }).(pulumi.IntPtrOutput)
}

func (o MetricBucketOptionsExponentialBucketsOutput) Scale() pulumi.Float64PtrOutput {
	return o.ApplyT(func (v MetricBucketOptionsExponentialBuckets) *float64 { return v.Scale }).(pulumi.Float64PtrOutput)
}

type MetricBucketOptionsExponentialBucketsPtrOutput struct { *pulumi.OutputState}

func (MetricBucketOptionsExponentialBucketsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricBucketOptionsExponentialBuckets)(nil)).Elem()
}

func (o MetricBucketOptionsExponentialBucketsPtrOutput) ToMetricBucketOptionsExponentialBucketsPtrOutput() MetricBucketOptionsExponentialBucketsPtrOutput {
	return o
}

func (o MetricBucketOptionsExponentialBucketsPtrOutput) ToMetricBucketOptionsExponentialBucketsPtrOutputWithContext(ctx context.Context) MetricBucketOptionsExponentialBucketsPtrOutput {
	return o
}

func (o MetricBucketOptionsExponentialBucketsPtrOutput) Elem() MetricBucketOptionsExponentialBucketsOutput {
	return o.ApplyT(func (v *MetricBucketOptionsExponentialBuckets) MetricBucketOptionsExponentialBuckets { return *v }).(MetricBucketOptionsExponentialBucketsOutput)
}

func (o MetricBucketOptionsExponentialBucketsPtrOutput) GrowthFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func (v MetricBucketOptionsExponentialBuckets) *float64 { return v.GrowthFactor }).(pulumi.Float64PtrOutput)
}

func (o MetricBucketOptionsExponentialBucketsPtrOutput) NumFiniteBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func (v MetricBucketOptionsExponentialBuckets) *int { return v.NumFiniteBuckets }).(pulumi.IntPtrOutput)
}

func (o MetricBucketOptionsExponentialBucketsPtrOutput) Scale() pulumi.Float64PtrOutput {
	return o.ApplyT(func (v MetricBucketOptionsExponentialBuckets) *float64 { return v.Scale }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MetricBucketOptionsExponentialBucketsOutput{})
	pulumi.RegisterOutputType(MetricBucketOptionsExponentialBucketsPtrOutput{})
}
