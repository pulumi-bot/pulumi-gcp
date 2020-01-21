// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package JobPubsubTarget

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type JobPubsubTarget struct {
	Attributes map[string]string `pulumi:"attributes"`
	Data *string `pulumi:"data"`
	TopicName string `pulumi:"topicName"`
}

type JobPubsubTargetInput interface {
	pulumi.Input

	ToJobPubsubTargetOutput() JobPubsubTargetOutput
	ToJobPubsubTargetOutputWithContext(context.Context) JobPubsubTargetOutput
}

type JobPubsubTargetArgs struct {
	Attributes pulumi.StringMapInput `pulumi:"attributes"`
	Data pulumi.StringPtrInput `pulumi:"data"`
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (JobPubsubTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobPubsubTarget)(nil)).Elem()
}

func (i JobPubsubTargetArgs) ToJobPubsubTargetOutput() JobPubsubTargetOutput {
	return i.ToJobPubsubTargetOutputWithContext(context.Background())
}

func (i JobPubsubTargetArgs) ToJobPubsubTargetOutputWithContext(ctx context.Context) JobPubsubTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPubsubTargetOutput)
}

func (i JobPubsubTargetArgs) ToJobPubsubTargetPtrOutput() JobPubsubTargetPtrOutput {
	return i.ToJobPubsubTargetPtrOutputWithContext(context.Background())
}

func (i JobPubsubTargetArgs) ToJobPubsubTargetPtrOutputWithContext(ctx context.Context) JobPubsubTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPubsubTargetOutput).ToJobPubsubTargetPtrOutputWithContext(ctx)
}

type JobPubsubTargetPtrInput interface {
	pulumi.Input

	ToJobPubsubTargetPtrOutput() JobPubsubTargetPtrOutput
	ToJobPubsubTargetPtrOutputWithContext(context.Context) JobPubsubTargetPtrOutput
}

type jobPubsubTargetPtrType JobPubsubTargetArgs

func JobPubsubTargetPtr(v *JobPubsubTargetArgs) JobPubsubTargetPtrInput {	return (*jobPubsubTargetPtrType)(v)
}

func (*jobPubsubTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobPubsubTarget)(nil)).Elem()
}

func (i *jobPubsubTargetPtrType) ToJobPubsubTargetPtrOutput() JobPubsubTargetPtrOutput {
	return i.ToJobPubsubTargetPtrOutputWithContext(context.Background())
}

func (i *jobPubsubTargetPtrType) ToJobPubsubTargetPtrOutputWithContext(ctx context.Context) JobPubsubTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPubsubTargetPtrOutput)
}

type JobPubsubTargetOutput struct { *pulumi.OutputState }

func (JobPubsubTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobPubsubTarget)(nil)).Elem()
}

func (o JobPubsubTargetOutput) ToJobPubsubTargetOutput() JobPubsubTargetOutput {
	return o
}

func (o JobPubsubTargetOutput) ToJobPubsubTargetOutputWithContext(ctx context.Context) JobPubsubTargetOutput {
	return o
}

func (o JobPubsubTargetOutput) ToJobPubsubTargetPtrOutput() JobPubsubTargetPtrOutput {
	return o.ToJobPubsubTargetPtrOutputWithContext(context.Background())
}

func (o JobPubsubTargetOutput) ToJobPubsubTargetPtrOutputWithContext(ctx context.Context) JobPubsubTargetPtrOutput {
	return o.ApplyT(func(v JobPubsubTarget) *JobPubsubTarget {
		return &v
	}).(JobPubsubTargetPtrOutput)
}
func (o JobPubsubTargetOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func (v JobPubsubTarget) map[string]string { return v.Attributes }).(pulumi.StringMapOutput)
}

func (o JobPubsubTargetOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobPubsubTarget) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o JobPubsubTargetOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func (v JobPubsubTarget) string { return v.TopicName }).(pulumi.StringOutput)
}

type JobPubsubTargetPtrOutput struct { *pulumi.OutputState}

func (JobPubsubTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobPubsubTarget)(nil)).Elem()
}

func (o JobPubsubTargetPtrOutput) ToJobPubsubTargetPtrOutput() JobPubsubTargetPtrOutput {
	return o
}

func (o JobPubsubTargetPtrOutput) ToJobPubsubTargetPtrOutputWithContext(ctx context.Context) JobPubsubTargetPtrOutput {
	return o
}

func (o JobPubsubTargetPtrOutput) Elem() JobPubsubTargetOutput {
	return o.ApplyT(func (v *JobPubsubTarget) JobPubsubTarget { return *v }).(JobPubsubTargetOutput)
}

func (o JobPubsubTargetPtrOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func (v JobPubsubTarget) map[string]string { return v.Attributes }).(pulumi.StringMapOutput)
}

func (o JobPubsubTargetPtrOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobPubsubTarget) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o JobPubsubTargetPtrOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func (v JobPubsubTarget) string { return v.TopicName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobPubsubTargetOutput{})
	pulumi.RegisterOutputType(JobPubsubTargetPtrOutput{})
}
