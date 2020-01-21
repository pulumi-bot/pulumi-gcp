// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ProviderBatching

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ProviderBatching struct {
	EnableBatching *bool `pulumi:"enableBatching"`
	SendAfter *string `pulumi:"sendAfter"`
}

type ProviderBatchingInput interface {
	pulumi.Input

	ToProviderBatchingOutput() ProviderBatchingOutput
	ToProviderBatchingOutputWithContext(context.Context) ProviderBatchingOutput
}

type ProviderBatchingArgs struct {
	EnableBatching pulumi.BoolPtrInput `pulumi:"enableBatching"`
	SendAfter pulumi.StringPtrInput `pulumi:"sendAfter"`
}

func (ProviderBatchingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderBatching)(nil)).Elem()
}

func (i ProviderBatchingArgs) ToProviderBatchingOutput() ProviderBatchingOutput {
	return i.ToProviderBatchingOutputWithContext(context.Background())
}

func (i ProviderBatchingArgs) ToProviderBatchingOutputWithContext(ctx context.Context) ProviderBatchingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderBatchingOutput)
}

type ProviderBatchingOutput struct { *pulumi.OutputState }

func (ProviderBatchingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderBatching)(nil)).Elem()
}

func (o ProviderBatchingOutput) ToProviderBatchingOutput() ProviderBatchingOutput {
	return o
}

func (o ProviderBatchingOutput) ToProviderBatchingOutputWithContext(ctx context.Context) ProviderBatchingOutput {
	return o
}

func (o ProviderBatchingOutput) EnableBatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ProviderBatching) *bool { return v.EnableBatching }).(pulumi.BoolPtrOutput)
}

func (o ProviderBatchingOutput) SendAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ProviderBatching) *string { return v.SendAfter }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProviderBatchingOutput{})
}
