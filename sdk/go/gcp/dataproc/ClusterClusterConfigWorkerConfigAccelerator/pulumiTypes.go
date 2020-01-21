// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterClusterConfigWorkerConfigAccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterClusterConfigWorkerConfigAccelerator struct {
	AcceleratorCount int `pulumi:"acceleratorCount"`
	AcceleratorType string `pulumi:"acceleratorType"`
}

type ClusterClusterConfigWorkerConfigAcceleratorInput interface {
	pulumi.Input

	ToClusterClusterConfigWorkerConfigAcceleratorOutput() ClusterClusterConfigWorkerConfigAcceleratorOutput
	ToClusterClusterConfigWorkerConfigAcceleratorOutputWithContext(context.Context) ClusterClusterConfigWorkerConfigAcceleratorOutput
}

type ClusterClusterConfigWorkerConfigAcceleratorArgs struct {
	AcceleratorCount pulumi.IntInput `pulumi:"acceleratorCount"`
	AcceleratorType pulumi.StringInput `pulumi:"acceleratorType"`
}

func (ClusterClusterConfigWorkerConfigAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigWorkerConfigAccelerator)(nil)).Elem()
}

func (i ClusterClusterConfigWorkerConfigAcceleratorArgs) ToClusterClusterConfigWorkerConfigAcceleratorOutput() ClusterClusterConfigWorkerConfigAcceleratorOutput {
	return i.ToClusterClusterConfigWorkerConfigAcceleratorOutputWithContext(context.Background())
}

func (i ClusterClusterConfigWorkerConfigAcceleratorArgs) ToClusterClusterConfigWorkerConfigAcceleratorOutputWithContext(ctx context.Context) ClusterClusterConfigWorkerConfigAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigWorkerConfigAcceleratorOutput)
}

type ClusterClusterConfigWorkerConfigAcceleratorArrayInput interface {
	pulumi.Input

	ToClusterClusterConfigWorkerConfigAcceleratorArrayOutput() ClusterClusterConfigWorkerConfigAcceleratorArrayOutput
	ToClusterClusterConfigWorkerConfigAcceleratorArrayOutputWithContext(context.Context) ClusterClusterConfigWorkerConfigAcceleratorArrayOutput
}

type ClusterClusterConfigWorkerConfigAcceleratorArray []ClusterClusterConfigWorkerConfigAcceleratorInput

func (ClusterClusterConfigWorkerConfigAcceleratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterClusterConfigWorkerConfigAccelerator)(nil)).Elem()
}

func (i ClusterClusterConfigWorkerConfigAcceleratorArray) ToClusterClusterConfigWorkerConfigAcceleratorArrayOutput() ClusterClusterConfigWorkerConfigAcceleratorArrayOutput {
	return i.ToClusterClusterConfigWorkerConfigAcceleratorArrayOutputWithContext(context.Background())
}

func (i ClusterClusterConfigWorkerConfigAcceleratorArray) ToClusterClusterConfigWorkerConfigAcceleratorArrayOutputWithContext(ctx context.Context) ClusterClusterConfigWorkerConfigAcceleratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigWorkerConfigAcceleratorArrayOutput)
}

type ClusterClusterConfigWorkerConfigAcceleratorOutput struct { *pulumi.OutputState }

func (ClusterClusterConfigWorkerConfigAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigWorkerConfigAccelerator)(nil)).Elem()
}

func (o ClusterClusterConfigWorkerConfigAcceleratorOutput) ToClusterClusterConfigWorkerConfigAcceleratorOutput() ClusterClusterConfigWorkerConfigAcceleratorOutput {
	return o
}

func (o ClusterClusterConfigWorkerConfigAcceleratorOutput) ToClusterClusterConfigWorkerConfigAcceleratorOutputWithContext(ctx context.Context) ClusterClusterConfigWorkerConfigAcceleratorOutput {
	return o
}

func (o ClusterClusterConfigWorkerConfigAcceleratorOutput) AcceleratorCount() pulumi.IntOutput {
	return o.ApplyT(func (v ClusterClusterConfigWorkerConfigAccelerator) int { return v.AcceleratorCount }).(pulumi.IntOutput)
}

func (o ClusterClusterConfigWorkerConfigAcceleratorOutput) AcceleratorType() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterClusterConfigWorkerConfigAccelerator) string { return v.AcceleratorType }).(pulumi.StringOutput)
}

type ClusterClusterConfigWorkerConfigAcceleratorArrayOutput struct { *pulumi.OutputState}

func (ClusterClusterConfigWorkerConfigAcceleratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterClusterConfigWorkerConfigAccelerator)(nil)).Elem()
}

func (o ClusterClusterConfigWorkerConfigAcceleratorArrayOutput) ToClusterClusterConfigWorkerConfigAcceleratorArrayOutput() ClusterClusterConfigWorkerConfigAcceleratorArrayOutput {
	return o
}

func (o ClusterClusterConfigWorkerConfigAcceleratorArrayOutput) ToClusterClusterConfigWorkerConfigAcceleratorArrayOutputWithContext(ctx context.Context) ClusterClusterConfigWorkerConfigAcceleratorArrayOutput {
	return o
}

func (o ClusterClusterConfigWorkerConfigAcceleratorArrayOutput) Index(i pulumi.IntInput) ClusterClusterConfigWorkerConfigAcceleratorOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ClusterClusterConfigWorkerConfigAccelerator {
		return vs[0].([]ClusterClusterConfigWorkerConfigAccelerator)[vs[1].(int)]
	}).(ClusterClusterConfigWorkerConfigAcceleratorOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterClusterConfigWorkerConfigAcceleratorOutput{})
	pulumi.RegisterOutputType(ClusterClusterConfigWorkerConfigAcceleratorArrayOutput{})
}
