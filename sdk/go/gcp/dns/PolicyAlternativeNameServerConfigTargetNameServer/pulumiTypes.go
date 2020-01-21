// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package PolicyAlternativeNameServerConfigTargetNameServer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type PolicyAlternativeNameServerConfigTargetNameServer struct {
	Ipv4Address string `pulumi:"ipv4Address"`
}

type PolicyAlternativeNameServerConfigTargetNameServerInput interface {
	pulumi.Input

	ToPolicyAlternativeNameServerConfigTargetNameServerOutput() PolicyAlternativeNameServerConfigTargetNameServerOutput
	ToPolicyAlternativeNameServerConfigTargetNameServerOutputWithContext(context.Context) PolicyAlternativeNameServerConfigTargetNameServerOutput
}

type PolicyAlternativeNameServerConfigTargetNameServerArgs struct {
	Ipv4Address pulumi.StringInput `pulumi:"ipv4Address"`
}

func (PolicyAlternativeNameServerConfigTargetNameServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAlternativeNameServerConfigTargetNameServer)(nil)).Elem()
}

func (i PolicyAlternativeNameServerConfigTargetNameServerArgs) ToPolicyAlternativeNameServerConfigTargetNameServerOutput() PolicyAlternativeNameServerConfigTargetNameServerOutput {
	return i.ToPolicyAlternativeNameServerConfigTargetNameServerOutputWithContext(context.Background())
}

func (i PolicyAlternativeNameServerConfigTargetNameServerArgs) ToPolicyAlternativeNameServerConfigTargetNameServerOutputWithContext(ctx context.Context) PolicyAlternativeNameServerConfigTargetNameServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAlternativeNameServerConfigTargetNameServerOutput)
}

type PolicyAlternativeNameServerConfigTargetNameServerArrayInput interface {
	pulumi.Input

	ToPolicyAlternativeNameServerConfigTargetNameServerArrayOutput() PolicyAlternativeNameServerConfigTargetNameServerArrayOutput
	ToPolicyAlternativeNameServerConfigTargetNameServerArrayOutputWithContext(context.Context) PolicyAlternativeNameServerConfigTargetNameServerArrayOutput
}

type PolicyAlternativeNameServerConfigTargetNameServerArray []PolicyAlternativeNameServerConfigTargetNameServerInput

func (PolicyAlternativeNameServerConfigTargetNameServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAlternativeNameServerConfigTargetNameServer)(nil)).Elem()
}

func (i PolicyAlternativeNameServerConfigTargetNameServerArray) ToPolicyAlternativeNameServerConfigTargetNameServerArrayOutput() PolicyAlternativeNameServerConfigTargetNameServerArrayOutput {
	return i.ToPolicyAlternativeNameServerConfigTargetNameServerArrayOutputWithContext(context.Background())
}

func (i PolicyAlternativeNameServerConfigTargetNameServerArray) ToPolicyAlternativeNameServerConfigTargetNameServerArrayOutputWithContext(ctx context.Context) PolicyAlternativeNameServerConfigTargetNameServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAlternativeNameServerConfigTargetNameServerArrayOutput)
}

type PolicyAlternativeNameServerConfigTargetNameServerOutput struct { *pulumi.OutputState }

func (PolicyAlternativeNameServerConfigTargetNameServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAlternativeNameServerConfigTargetNameServer)(nil)).Elem()
}

func (o PolicyAlternativeNameServerConfigTargetNameServerOutput) ToPolicyAlternativeNameServerConfigTargetNameServerOutput() PolicyAlternativeNameServerConfigTargetNameServerOutput {
	return o
}

func (o PolicyAlternativeNameServerConfigTargetNameServerOutput) ToPolicyAlternativeNameServerConfigTargetNameServerOutputWithContext(ctx context.Context) PolicyAlternativeNameServerConfigTargetNameServerOutput {
	return o
}

func (o PolicyAlternativeNameServerConfigTargetNameServerOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func (v PolicyAlternativeNameServerConfigTargetNameServer) string { return v.Ipv4Address }).(pulumi.StringOutput)
}

type PolicyAlternativeNameServerConfigTargetNameServerArrayOutput struct { *pulumi.OutputState}

func (PolicyAlternativeNameServerConfigTargetNameServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAlternativeNameServerConfigTargetNameServer)(nil)).Elem()
}

func (o PolicyAlternativeNameServerConfigTargetNameServerArrayOutput) ToPolicyAlternativeNameServerConfigTargetNameServerArrayOutput() PolicyAlternativeNameServerConfigTargetNameServerArrayOutput {
	return o
}

func (o PolicyAlternativeNameServerConfigTargetNameServerArrayOutput) ToPolicyAlternativeNameServerConfigTargetNameServerArrayOutputWithContext(ctx context.Context) PolicyAlternativeNameServerConfigTargetNameServerArrayOutput {
	return o
}

func (o PolicyAlternativeNameServerConfigTargetNameServerArrayOutput) Index(i pulumi.IntInput) PolicyAlternativeNameServerConfigTargetNameServerOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) PolicyAlternativeNameServerConfigTargetNameServer {
		return vs[0].([]PolicyAlternativeNameServerConfigTargetNameServer)[vs[1].(int)]
	}).(PolicyAlternativeNameServerConfigTargetNameServerOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyAlternativeNameServerConfigTargetNameServerOutput{})
	pulumi.RegisterOutputType(PolicyAlternativeNameServerConfigTargetNameServerArrayOutput{})
}
