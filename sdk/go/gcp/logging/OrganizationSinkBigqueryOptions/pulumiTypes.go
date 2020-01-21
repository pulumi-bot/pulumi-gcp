// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package OrganizationSinkBigqueryOptions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type OrganizationSinkBigqueryOptions struct {
	UsePartitionedTables bool `pulumi:"usePartitionedTables"`
}

type OrganizationSinkBigqueryOptionsInput interface {
	pulumi.Input

	ToOrganizationSinkBigqueryOptionsOutput() OrganizationSinkBigqueryOptionsOutput
	ToOrganizationSinkBigqueryOptionsOutputWithContext(context.Context) OrganizationSinkBigqueryOptionsOutput
}

type OrganizationSinkBigqueryOptionsArgs struct {
	UsePartitionedTables pulumi.BoolInput `pulumi:"usePartitionedTables"`
}

func (OrganizationSinkBigqueryOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSinkBigqueryOptions)(nil)).Elem()
}

func (i OrganizationSinkBigqueryOptionsArgs) ToOrganizationSinkBigqueryOptionsOutput() OrganizationSinkBigqueryOptionsOutput {
	return i.ToOrganizationSinkBigqueryOptionsOutputWithContext(context.Background())
}

func (i OrganizationSinkBigqueryOptionsArgs) ToOrganizationSinkBigqueryOptionsOutputWithContext(ctx context.Context) OrganizationSinkBigqueryOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkBigqueryOptionsOutput)
}

func (i OrganizationSinkBigqueryOptionsArgs) ToOrganizationSinkBigqueryOptionsPtrOutput() OrganizationSinkBigqueryOptionsPtrOutput {
	return i.ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(context.Background())
}

func (i OrganizationSinkBigqueryOptionsArgs) ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(ctx context.Context) OrganizationSinkBigqueryOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkBigqueryOptionsOutput).ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(ctx)
}

type OrganizationSinkBigqueryOptionsPtrInput interface {
	pulumi.Input

	ToOrganizationSinkBigqueryOptionsPtrOutput() OrganizationSinkBigqueryOptionsPtrOutput
	ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(context.Context) OrganizationSinkBigqueryOptionsPtrOutput
}

type organizationSinkBigqueryOptionsPtrType OrganizationSinkBigqueryOptionsArgs

func OrganizationSinkBigqueryOptionsPtr(v *OrganizationSinkBigqueryOptionsArgs) OrganizationSinkBigqueryOptionsPtrInput {	return (*organizationSinkBigqueryOptionsPtrType)(v)
}

func (*organizationSinkBigqueryOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSinkBigqueryOptions)(nil)).Elem()
}

func (i *organizationSinkBigqueryOptionsPtrType) ToOrganizationSinkBigqueryOptionsPtrOutput() OrganizationSinkBigqueryOptionsPtrOutput {
	return i.ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(context.Background())
}

func (i *organizationSinkBigqueryOptionsPtrType) ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(ctx context.Context) OrganizationSinkBigqueryOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkBigqueryOptionsPtrOutput)
}

type OrganizationSinkBigqueryOptionsOutput struct { *pulumi.OutputState }

func (OrganizationSinkBigqueryOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSinkBigqueryOptions)(nil)).Elem()
}

func (o OrganizationSinkBigqueryOptionsOutput) ToOrganizationSinkBigqueryOptionsOutput() OrganizationSinkBigqueryOptionsOutput {
	return o
}

func (o OrganizationSinkBigqueryOptionsOutput) ToOrganizationSinkBigqueryOptionsOutputWithContext(ctx context.Context) OrganizationSinkBigqueryOptionsOutput {
	return o
}

func (o OrganizationSinkBigqueryOptionsOutput) ToOrganizationSinkBigqueryOptionsPtrOutput() OrganizationSinkBigqueryOptionsPtrOutput {
	return o.ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(context.Background())
}

func (o OrganizationSinkBigqueryOptionsOutput) ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(ctx context.Context) OrganizationSinkBigqueryOptionsPtrOutput {
	return o.ApplyT(func(v OrganizationSinkBigqueryOptions) *OrganizationSinkBigqueryOptions {
		return &v
	}).(OrganizationSinkBigqueryOptionsPtrOutput)
}
func (o OrganizationSinkBigqueryOptionsOutput) UsePartitionedTables() pulumi.BoolOutput {
	return o.ApplyT(func (v OrganizationSinkBigqueryOptions) bool { return v.UsePartitionedTables }).(pulumi.BoolOutput)
}

type OrganizationSinkBigqueryOptionsPtrOutput struct { *pulumi.OutputState}

func (OrganizationSinkBigqueryOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSinkBigqueryOptions)(nil)).Elem()
}

func (o OrganizationSinkBigqueryOptionsPtrOutput) ToOrganizationSinkBigqueryOptionsPtrOutput() OrganizationSinkBigqueryOptionsPtrOutput {
	return o
}

func (o OrganizationSinkBigqueryOptionsPtrOutput) ToOrganizationSinkBigqueryOptionsPtrOutputWithContext(ctx context.Context) OrganizationSinkBigqueryOptionsPtrOutput {
	return o
}

func (o OrganizationSinkBigqueryOptionsPtrOutput) Elem() OrganizationSinkBigqueryOptionsOutput {
	return o.ApplyT(func (v *OrganizationSinkBigqueryOptions) OrganizationSinkBigqueryOptions { return *v }).(OrganizationSinkBigqueryOptionsOutput)
}

func (o OrganizationSinkBigqueryOptionsPtrOutput) UsePartitionedTables() pulumi.BoolOutput {
	return o.ApplyT(func (v OrganizationSinkBigqueryOptions) bool { return v.UsePartitionedTables }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationSinkBigqueryOptionsOutput{})
	pulumi.RegisterOutputType(OrganizationSinkBigqueryOptionsPtrOutput{})
}
