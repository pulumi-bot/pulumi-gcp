// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getClusterNodePoolAutoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetClusterNodePoolAutoscaling struct {
	MaxNodeCount int `pulumi:"maxNodeCount"`
	MinNodeCount int `pulumi:"minNodeCount"`
}

type GetClusterNodePoolAutoscalingInput interface {
	pulumi.Input

	ToGetClusterNodePoolAutoscalingOutput() GetClusterNodePoolAutoscalingOutput
	ToGetClusterNodePoolAutoscalingOutputWithContext(context.Context) GetClusterNodePoolAutoscalingOutput
}

type GetClusterNodePoolAutoscalingArgs struct {
	MaxNodeCount pulumi.IntInput `pulumi:"maxNodeCount"`
	MinNodeCount pulumi.IntInput `pulumi:"minNodeCount"`
}

func (GetClusterNodePoolAutoscalingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterNodePoolAutoscaling)(nil)).Elem()
}

func (i GetClusterNodePoolAutoscalingArgs) ToGetClusterNodePoolAutoscalingOutput() GetClusterNodePoolAutoscalingOutput {
	return i.ToGetClusterNodePoolAutoscalingOutputWithContext(context.Background())
}

func (i GetClusterNodePoolAutoscalingArgs) ToGetClusterNodePoolAutoscalingOutputWithContext(ctx context.Context) GetClusterNodePoolAutoscalingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterNodePoolAutoscalingOutput)
}

type GetClusterNodePoolAutoscalingArrayInput interface {
	pulumi.Input

	ToGetClusterNodePoolAutoscalingArrayOutput() GetClusterNodePoolAutoscalingArrayOutput
	ToGetClusterNodePoolAutoscalingArrayOutputWithContext(context.Context) GetClusterNodePoolAutoscalingArrayOutput
}

type GetClusterNodePoolAutoscalingArray []GetClusterNodePoolAutoscalingInput

func (GetClusterNodePoolAutoscalingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterNodePoolAutoscaling)(nil)).Elem()
}

func (i GetClusterNodePoolAutoscalingArray) ToGetClusterNodePoolAutoscalingArrayOutput() GetClusterNodePoolAutoscalingArrayOutput {
	return i.ToGetClusterNodePoolAutoscalingArrayOutputWithContext(context.Background())
}

func (i GetClusterNodePoolAutoscalingArray) ToGetClusterNodePoolAutoscalingArrayOutputWithContext(ctx context.Context) GetClusterNodePoolAutoscalingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterNodePoolAutoscalingArrayOutput)
}

type GetClusterNodePoolAutoscalingOutput struct { *pulumi.OutputState }

func (GetClusterNodePoolAutoscalingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterNodePoolAutoscaling)(nil)).Elem()
}

func (o GetClusterNodePoolAutoscalingOutput) ToGetClusterNodePoolAutoscalingOutput() GetClusterNodePoolAutoscalingOutput {
	return o
}

func (o GetClusterNodePoolAutoscalingOutput) ToGetClusterNodePoolAutoscalingOutputWithContext(ctx context.Context) GetClusterNodePoolAutoscalingOutput {
	return o
}

func (o GetClusterNodePoolAutoscalingOutput) MaxNodeCount() pulumi.IntOutput {
	return o.ApplyT(func (v GetClusterNodePoolAutoscaling) int { return v.MaxNodeCount }).(pulumi.IntOutput)
}

func (o GetClusterNodePoolAutoscalingOutput) MinNodeCount() pulumi.IntOutput {
	return o.ApplyT(func (v GetClusterNodePoolAutoscaling) int { return v.MinNodeCount }).(pulumi.IntOutput)
}

type GetClusterNodePoolAutoscalingArrayOutput struct { *pulumi.OutputState}

func (GetClusterNodePoolAutoscalingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterNodePoolAutoscaling)(nil)).Elem()
}

func (o GetClusterNodePoolAutoscalingArrayOutput) ToGetClusterNodePoolAutoscalingArrayOutput() GetClusterNodePoolAutoscalingArrayOutput {
	return o
}

func (o GetClusterNodePoolAutoscalingArrayOutput) ToGetClusterNodePoolAutoscalingArrayOutputWithContext(ctx context.Context) GetClusterNodePoolAutoscalingArrayOutput {
	return o
}

func (o GetClusterNodePoolAutoscalingArrayOutput) Index(i pulumi.IntInput) GetClusterNodePoolAutoscalingOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetClusterNodePoolAutoscaling {
		return vs[0].([]GetClusterNodePoolAutoscaling)[vs[1].(int)]
	}).(GetClusterNodePoolAutoscalingOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterNodePoolAutoscalingOutput{})
	pulumi.RegisterOutputType(GetClusterNodePoolAutoscalingArrayOutput{})
}
