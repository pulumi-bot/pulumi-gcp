// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package JobAppEngineHttpTargetAppEngineRouting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type JobAppEngineHttpTargetAppEngineRouting struct {
	Instance *string `pulumi:"instance"`
	Service *string `pulumi:"service"`
	Version *string `pulumi:"version"`
}

type JobAppEngineHttpTargetAppEngineRoutingInput interface {
	pulumi.Input

	ToJobAppEngineHttpTargetAppEngineRoutingOutput() JobAppEngineHttpTargetAppEngineRoutingOutput
	ToJobAppEngineHttpTargetAppEngineRoutingOutputWithContext(context.Context) JobAppEngineHttpTargetAppEngineRoutingOutput
}

type JobAppEngineHttpTargetAppEngineRoutingArgs struct {
	Instance pulumi.StringPtrInput `pulumi:"instance"`
	Service pulumi.StringPtrInput `pulumi:"service"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (JobAppEngineHttpTargetAppEngineRoutingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobAppEngineHttpTargetAppEngineRouting)(nil)).Elem()
}

func (i JobAppEngineHttpTargetAppEngineRoutingArgs) ToJobAppEngineHttpTargetAppEngineRoutingOutput() JobAppEngineHttpTargetAppEngineRoutingOutput {
	return i.ToJobAppEngineHttpTargetAppEngineRoutingOutputWithContext(context.Background())
}

func (i JobAppEngineHttpTargetAppEngineRoutingArgs) ToJobAppEngineHttpTargetAppEngineRoutingOutputWithContext(ctx context.Context) JobAppEngineHttpTargetAppEngineRoutingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobAppEngineHttpTargetAppEngineRoutingOutput)
}

func (i JobAppEngineHttpTargetAppEngineRoutingArgs) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutput() JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return i.ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(context.Background())
}

func (i JobAppEngineHttpTargetAppEngineRoutingArgs) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(ctx context.Context) JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobAppEngineHttpTargetAppEngineRoutingOutput).ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(ctx)
}

type JobAppEngineHttpTargetAppEngineRoutingPtrInput interface {
	pulumi.Input

	ToJobAppEngineHttpTargetAppEngineRoutingPtrOutput() JobAppEngineHttpTargetAppEngineRoutingPtrOutput
	ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(context.Context) JobAppEngineHttpTargetAppEngineRoutingPtrOutput
}

type jobAppEngineHttpTargetAppEngineRoutingPtrType JobAppEngineHttpTargetAppEngineRoutingArgs

func JobAppEngineHttpTargetAppEngineRoutingPtr(v *JobAppEngineHttpTargetAppEngineRoutingArgs) JobAppEngineHttpTargetAppEngineRoutingPtrInput {	return (*jobAppEngineHttpTargetAppEngineRoutingPtrType)(v)
}

func (*jobAppEngineHttpTargetAppEngineRoutingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAppEngineHttpTargetAppEngineRouting)(nil)).Elem()
}

func (i *jobAppEngineHttpTargetAppEngineRoutingPtrType) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutput() JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return i.ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(context.Background())
}

func (i *jobAppEngineHttpTargetAppEngineRoutingPtrType) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(ctx context.Context) JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobAppEngineHttpTargetAppEngineRoutingPtrOutput)
}

type JobAppEngineHttpTargetAppEngineRoutingOutput struct { *pulumi.OutputState }

func (JobAppEngineHttpTargetAppEngineRoutingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobAppEngineHttpTargetAppEngineRouting)(nil)).Elem()
}

func (o JobAppEngineHttpTargetAppEngineRoutingOutput) ToJobAppEngineHttpTargetAppEngineRoutingOutput() JobAppEngineHttpTargetAppEngineRoutingOutput {
	return o
}

func (o JobAppEngineHttpTargetAppEngineRoutingOutput) ToJobAppEngineHttpTargetAppEngineRoutingOutputWithContext(ctx context.Context) JobAppEngineHttpTargetAppEngineRoutingOutput {
	return o
}

func (o JobAppEngineHttpTargetAppEngineRoutingOutput) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutput() JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return o.ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(context.Background())
}

func (o JobAppEngineHttpTargetAppEngineRoutingOutput) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(ctx context.Context) JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return o.ApplyT(func(v JobAppEngineHttpTargetAppEngineRouting) *JobAppEngineHttpTargetAppEngineRouting {
		return &v
	}).(JobAppEngineHttpTargetAppEngineRoutingPtrOutput)
}
func (o JobAppEngineHttpTargetAppEngineRoutingOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobAppEngineHttpTargetAppEngineRouting) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o JobAppEngineHttpTargetAppEngineRoutingOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobAppEngineHttpTargetAppEngineRouting) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o JobAppEngineHttpTargetAppEngineRoutingOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobAppEngineHttpTargetAppEngineRouting) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type JobAppEngineHttpTargetAppEngineRoutingPtrOutput struct { *pulumi.OutputState}

func (JobAppEngineHttpTargetAppEngineRoutingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAppEngineHttpTargetAppEngineRouting)(nil)).Elem()
}

func (o JobAppEngineHttpTargetAppEngineRoutingPtrOutput) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutput() JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return o
}

func (o JobAppEngineHttpTargetAppEngineRoutingPtrOutput) ToJobAppEngineHttpTargetAppEngineRoutingPtrOutputWithContext(ctx context.Context) JobAppEngineHttpTargetAppEngineRoutingPtrOutput {
	return o
}

func (o JobAppEngineHttpTargetAppEngineRoutingPtrOutput) Elem() JobAppEngineHttpTargetAppEngineRoutingOutput {
	return o.ApplyT(func (v *JobAppEngineHttpTargetAppEngineRouting) JobAppEngineHttpTargetAppEngineRouting { return *v }).(JobAppEngineHttpTargetAppEngineRoutingOutput)
}

func (o JobAppEngineHttpTargetAppEngineRoutingPtrOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobAppEngineHttpTargetAppEngineRouting) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o JobAppEngineHttpTargetAppEngineRoutingPtrOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobAppEngineHttpTargetAppEngineRouting) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o JobAppEngineHttpTargetAppEngineRoutingPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobAppEngineHttpTargetAppEngineRouting) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(JobAppEngineHttpTargetAppEngineRoutingOutput{})
	pulumi.RegisterOutputType(JobAppEngineHttpTargetAppEngineRoutingPtrOutput{})
}
