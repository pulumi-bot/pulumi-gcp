// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterClusterConfigMasterConfigAccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterClusterConfigMasterConfigAccelerator struct {
	AcceleratorCount int `pulumi:"acceleratorCount"`
	AcceleratorType string `pulumi:"acceleratorType"`
}

type ClusterClusterConfigMasterConfigAcceleratorInput interface {
	pulumi.Input

	ToClusterClusterConfigMasterConfigAcceleratorOutput() ClusterClusterConfigMasterConfigAcceleratorOutput
	ToClusterClusterConfigMasterConfigAcceleratorOutputWithContext(context.Context) ClusterClusterConfigMasterConfigAcceleratorOutput
}

type ClusterClusterConfigMasterConfigAcceleratorArgs struct {
	AcceleratorCount pulumi.IntInput `pulumi:"acceleratorCount"`
	AcceleratorType pulumi.StringInput `pulumi:"acceleratorType"`
}

func (ClusterClusterConfigMasterConfigAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigMasterConfigAccelerator)(nil)).Elem()
}

func (i ClusterClusterConfigMasterConfigAcceleratorArgs) ToClusterClusterConfigMasterConfigAcceleratorOutput() ClusterClusterConfigMasterConfigAcceleratorOutput {
	return i.ToClusterClusterConfigMasterConfigAcceleratorOutputWithContext(context.Background())
}

func (i ClusterClusterConfigMasterConfigAcceleratorArgs) ToClusterClusterConfigMasterConfigAcceleratorOutputWithContext(ctx context.Context) ClusterClusterConfigMasterConfigAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigMasterConfigAcceleratorOutput)
}

type ClusterClusterConfigMasterConfigAcceleratorArrayInput interface {
	pulumi.Input

	ToClusterClusterConfigMasterConfigAcceleratorArrayOutput() ClusterClusterConfigMasterConfigAcceleratorArrayOutput
	ToClusterClusterConfigMasterConfigAcceleratorArrayOutputWithContext(context.Context) ClusterClusterConfigMasterConfigAcceleratorArrayOutput
}

type ClusterClusterConfigMasterConfigAcceleratorArray []ClusterClusterConfigMasterConfigAcceleratorInput

func (ClusterClusterConfigMasterConfigAcceleratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterClusterConfigMasterConfigAccelerator)(nil)).Elem()
}

func (i ClusterClusterConfigMasterConfigAcceleratorArray) ToClusterClusterConfigMasterConfigAcceleratorArrayOutput() ClusterClusterConfigMasterConfigAcceleratorArrayOutput {
	return i.ToClusterClusterConfigMasterConfigAcceleratorArrayOutputWithContext(context.Background())
}

func (i ClusterClusterConfigMasterConfigAcceleratorArray) ToClusterClusterConfigMasterConfigAcceleratorArrayOutputWithContext(ctx context.Context) ClusterClusterConfigMasterConfigAcceleratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigMasterConfigAcceleratorArrayOutput)
}

type ClusterClusterConfigMasterConfigAcceleratorOutput struct { *pulumi.OutputState }

func (ClusterClusterConfigMasterConfigAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigMasterConfigAccelerator)(nil)).Elem()
}

func (o ClusterClusterConfigMasterConfigAcceleratorOutput) ToClusterClusterConfigMasterConfigAcceleratorOutput() ClusterClusterConfigMasterConfigAcceleratorOutput {
	return o
}

func (o ClusterClusterConfigMasterConfigAcceleratorOutput) ToClusterClusterConfigMasterConfigAcceleratorOutputWithContext(ctx context.Context) ClusterClusterConfigMasterConfigAcceleratorOutput {
	return o
}

func (o ClusterClusterConfigMasterConfigAcceleratorOutput) AcceleratorCount() pulumi.IntOutput {
	return o.ApplyT(func (v ClusterClusterConfigMasterConfigAccelerator) int { return v.AcceleratorCount }).(pulumi.IntOutput)
}

func (o ClusterClusterConfigMasterConfigAcceleratorOutput) AcceleratorType() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterClusterConfigMasterConfigAccelerator) string { return v.AcceleratorType }).(pulumi.StringOutput)
}

type ClusterClusterConfigMasterConfigAcceleratorArrayOutput struct { *pulumi.OutputState}

func (ClusterClusterConfigMasterConfigAcceleratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterClusterConfigMasterConfigAccelerator)(nil)).Elem()
}

func (o ClusterClusterConfigMasterConfigAcceleratorArrayOutput) ToClusterClusterConfigMasterConfigAcceleratorArrayOutput() ClusterClusterConfigMasterConfigAcceleratorArrayOutput {
	return o
}

func (o ClusterClusterConfigMasterConfigAcceleratorArrayOutput) ToClusterClusterConfigMasterConfigAcceleratorArrayOutputWithContext(ctx context.Context) ClusterClusterConfigMasterConfigAcceleratorArrayOutput {
	return o
}

func (o ClusterClusterConfigMasterConfigAcceleratorArrayOutput) Index(i pulumi.IntInput) ClusterClusterConfigMasterConfigAcceleratorOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ClusterClusterConfigMasterConfigAccelerator {
		return vs[0].([]ClusterClusterConfigMasterConfigAccelerator)[vs[1].(int)]
	}).(ClusterClusterConfigMasterConfigAcceleratorOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterClusterConfigMasterConfigAcceleratorOutput{})
	pulumi.RegisterOutputType(ClusterClusterConfigMasterConfigAcceleratorArrayOutput{})
}
