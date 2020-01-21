// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package FunctionEventTrigger

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/cloudfunctions/FunctionEventTriggerFailurePolicy"
)

type FunctionEventTrigger struct {
	EventType string `pulumi:"eventType"`
	FailurePolicy *cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicy `pulumi:"failurePolicy"`
	Resource string `pulumi:"resource"`
}

type FunctionEventTriggerInput interface {
	pulumi.Input

	ToFunctionEventTriggerOutput() FunctionEventTriggerOutput
	ToFunctionEventTriggerOutputWithContext(context.Context) FunctionEventTriggerOutput
}

type FunctionEventTriggerArgs struct {
	EventType pulumi.StringInput `pulumi:"eventType"`
	FailurePolicy cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicyPtrInput `pulumi:"failurePolicy"`
	Resource pulumi.StringInput `pulumi:"resource"`
}

func (FunctionEventTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionEventTrigger)(nil)).Elem()
}

func (i FunctionEventTriggerArgs) ToFunctionEventTriggerOutput() FunctionEventTriggerOutput {
	return i.ToFunctionEventTriggerOutputWithContext(context.Background())
}

func (i FunctionEventTriggerArgs) ToFunctionEventTriggerOutputWithContext(ctx context.Context) FunctionEventTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionEventTriggerOutput)
}

func (i FunctionEventTriggerArgs) ToFunctionEventTriggerPtrOutput() FunctionEventTriggerPtrOutput {
	return i.ToFunctionEventTriggerPtrOutputWithContext(context.Background())
}

func (i FunctionEventTriggerArgs) ToFunctionEventTriggerPtrOutputWithContext(ctx context.Context) FunctionEventTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionEventTriggerOutput).ToFunctionEventTriggerPtrOutputWithContext(ctx)
}

type FunctionEventTriggerPtrInput interface {
	pulumi.Input

	ToFunctionEventTriggerPtrOutput() FunctionEventTriggerPtrOutput
	ToFunctionEventTriggerPtrOutputWithContext(context.Context) FunctionEventTriggerPtrOutput
}

type functionEventTriggerPtrType FunctionEventTriggerArgs

func FunctionEventTriggerPtr(v *FunctionEventTriggerArgs) FunctionEventTriggerPtrInput {	return (*functionEventTriggerPtrType)(v)
}

func (*functionEventTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionEventTrigger)(nil)).Elem()
}

func (i *functionEventTriggerPtrType) ToFunctionEventTriggerPtrOutput() FunctionEventTriggerPtrOutput {
	return i.ToFunctionEventTriggerPtrOutputWithContext(context.Background())
}

func (i *functionEventTriggerPtrType) ToFunctionEventTriggerPtrOutputWithContext(ctx context.Context) FunctionEventTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionEventTriggerPtrOutput)
}

type FunctionEventTriggerOutput struct { *pulumi.OutputState }

func (FunctionEventTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionEventTrigger)(nil)).Elem()
}

func (o FunctionEventTriggerOutput) ToFunctionEventTriggerOutput() FunctionEventTriggerOutput {
	return o
}

func (o FunctionEventTriggerOutput) ToFunctionEventTriggerOutputWithContext(ctx context.Context) FunctionEventTriggerOutput {
	return o
}

func (o FunctionEventTriggerOutput) ToFunctionEventTriggerPtrOutput() FunctionEventTriggerPtrOutput {
	return o.ToFunctionEventTriggerPtrOutputWithContext(context.Background())
}

func (o FunctionEventTriggerOutput) ToFunctionEventTriggerPtrOutputWithContext(ctx context.Context) FunctionEventTriggerPtrOutput {
	return o.ApplyT(func(v FunctionEventTrigger) *FunctionEventTrigger {
		return &v
	}).(FunctionEventTriggerPtrOutput)
}
func (o FunctionEventTriggerOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func (v FunctionEventTrigger) string { return v.EventType }).(pulumi.StringOutput)
}

func (o FunctionEventTriggerOutput) FailurePolicy() cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicyPtrOutput {
	return o.ApplyT(func (v FunctionEventTrigger) *cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicy { return v.FailurePolicy }).(cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicyPtrOutput)
}

func (o FunctionEventTriggerOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func (v FunctionEventTrigger) string { return v.Resource }).(pulumi.StringOutput)
}

type FunctionEventTriggerPtrOutput struct { *pulumi.OutputState}

func (FunctionEventTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionEventTrigger)(nil)).Elem()
}

func (o FunctionEventTriggerPtrOutput) ToFunctionEventTriggerPtrOutput() FunctionEventTriggerPtrOutput {
	return o
}

func (o FunctionEventTriggerPtrOutput) ToFunctionEventTriggerPtrOutputWithContext(ctx context.Context) FunctionEventTriggerPtrOutput {
	return o
}

func (o FunctionEventTriggerPtrOutput) Elem() FunctionEventTriggerOutput {
	return o.ApplyT(func (v *FunctionEventTrigger) FunctionEventTrigger { return *v }).(FunctionEventTriggerOutput)
}

func (o FunctionEventTriggerPtrOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func (v FunctionEventTrigger) string { return v.EventType }).(pulumi.StringOutput)
}

func (o FunctionEventTriggerPtrOutput) FailurePolicy() cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicyPtrOutput {
	return o.ApplyT(func (v FunctionEventTrigger) *cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicy { return v.FailurePolicy }).(cloudfunctionsFunctionEventTriggerFailurePolicy.FunctionEventTriggerFailurePolicyPtrOutput)
}

func (o FunctionEventTriggerPtrOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func (v FunctionEventTrigger) string { return v.Resource }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionEventTriggerOutput{})
	pulumi.RegisterOutputType(FunctionEventTriggerPtrOutput{})
}
