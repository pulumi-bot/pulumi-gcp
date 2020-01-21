// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package NodePoolNodeConfigTaint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type NodePoolNodeConfigTaint struct {
	Effect string `pulumi:"effect"`
	Key string `pulumi:"key"`
	Value string `pulumi:"value"`
}

type NodePoolNodeConfigTaintInput interface {
	pulumi.Input

	ToNodePoolNodeConfigTaintOutput() NodePoolNodeConfigTaintOutput
	ToNodePoolNodeConfigTaintOutputWithContext(context.Context) NodePoolNodeConfigTaintOutput
}

type NodePoolNodeConfigTaintArgs struct {
	Effect pulumi.StringInput `pulumi:"effect"`
	Key pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (NodePoolNodeConfigTaintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePoolNodeConfigTaint)(nil)).Elem()
}

func (i NodePoolNodeConfigTaintArgs) ToNodePoolNodeConfigTaintOutput() NodePoolNodeConfigTaintOutput {
	return i.ToNodePoolNodeConfigTaintOutputWithContext(context.Background())
}

func (i NodePoolNodeConfigTaintArgs) ToNodePoolNodeConfigTaintOutputWithContext(ctx context.Context) NodePoolNodeConfigTaintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolNodeConfigTaintOutput)
}

type NodePoolNodeConfigTaintArrayInput interface {
	pulumi.Input

	ToNodePoolNodeConfigTaintArrayOutput() NodePoolNodeConfigTaintArrayOutput
	ToNodePoolNodeConfigTaintArrayOutputWithContext(context.Context) NodePoolNodeConfigTaintArrayOutput
}

type NodePoolNodeConfigTaintArray []NodePoolNodeConfigTaintInput

func (NodePoolNodeConfigTaintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodePoolNodeConfigTaint)(nil)).Elem()
}

func (i NodePoolNodeConfigTaintArray) ToNodePoolNodeConfigTaintArrayOutput() NodePoolNodeConfigTaintArrayOutput {
	return i.ToNodePoolNodeConfigTaintArrayOutputWithContext(context.Background())
}

func (i NodePoolNodeConfigTaintArray) ToNodePoolNodeConfigTaintArrayOutputWithContext(ctx context.Context) NodePoolNodeConfigTaintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolNodeConfigTaintArrayOutput)
}

type NodePoolNodeConfigTaintOutput struct { *pulumi.OutputState }

func (NodePoolNodeConfigTaintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePoolNodeConfigTaint)(nil)).Elem()
}

func (o NodePoolNodeConfigTaintOutput) ToNodePoolNodeConfigTaintOutput() NodePoolNodeConfigTaintOutput {
	return o
}

func (o NodePoolNodeConfigTaintOutput) ToNodePoolNodeConfigTaintOutputWithContext(ctx context.Context) NodePoolNodeConfigTaintOutput {
	return o
}

func (o NodePoolNodeConfigTaintOutput) Effect() pulumi.StringOutput {
	return o.ApplyT(func (v NodePoolNodeConfigTaint) string { return v.Effect }).(pulumi.StringOutput)
}

func (o NodePoolNodeConfigTaintOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func (v NodePoolNodeConfigTaint) string { return v.Key }).(pulumi.StringOutput)
}

func (o NodePoolNodeConfigTaintOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v NodePoolNodeConfigTaint) string { return v.Value }).(pulumi.StringOutput)
}

type NodePoolNodeConfigTaintArrayOutput struct { *pulumi.OutputState}

func (NodePoolNodeConfigTaintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodePoolNodeConfigTaint)(nil)).Elem()
}

func (o NodePoolNodeConfigTaintArrayOutput) ToNodePoolNodeConfigTaintArrayOutput() NodePoolNodeConfigTaintArrayOutput {
	return o
}

func (o NodePoolNodeConfigTaintArrayOutput) ToNodePoolNodeConfigTaintArrayOutputWithContext(ctx context.Context) NodePoolNodeConfigTaintArrayOutput {
	return o
}

func (o NodePoolNodeConfigTaintArrayOutput) Index(i pulumi.IntInput) NodePoolNodeConfigTaintOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) NodePoolNodeConfigTaint {
		return vs[0].([]NodePoolNodeConfigTaint)[vs[1].(int)]
	}).(NodePoolNodeConfigTaintOutput)
}

func init() {
	pulumi.RegisterOutputType(NodePoolNodeConfigTaintOutput{})
	pulumi.RegisterOutputType(NodePoolNodeConfigTaintArrayOutput{})
}
