// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package FunctionSourceRepository

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type FunctionSourceRepository struct {
	DeployedUrl *string `pulumi:"deployedUrl"`
	Url string `pulumi:"url"`
}

type FunctionSourceRepositoryInput interface {
	pulumi.Input

	ToFunctionSourceRepositoryOutput() FunctionSourceRepositoryOutput
	ToFunctionSourceRepositoryOutputWithContext(context.Context) FunctionSourceRepositoryOutput
}

type FunctionSourceRepositoryArgs struct {
	DeployedUrl pulumi.StringPtrInput `pulumi:"deployedUrl"`
	Url pulumi.StringInput `pulumi:"url"`
}

func (FunctionSourceRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionSourceRepository)(nil)).Elem()
}

func (i FunctionSourceRepositoryArgs) ToFunctionSourceRepositoryOutput() FunctionSourceRepositoryOutput {
	return i.ToFunctionSourceRepositoryOutputWithContext(context.Background())
}

func (i FunctionSourceRepositoryArgs) ToFunctionSourceRepositoryOutputWithContext(ctx context.Context) FunctionSourceRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionSourceRepositoryOutput)
}

func (i FunctionSourceRepositoryArgs) ToFunctionSourceRepositoryPtrOutput() FunctionSourceRepositoryPtrOutput {
	return i.ToFunctionSourceRepositoryPtrOutputWithContext(context.Background())
}

func (i FunctionSourceRepositoryArgs) ToFunctionSourceRepositoryPtrOutputWithContext(ctx context.Context) FunctionSourceRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionSourceRepositoryOutput).ToFunctionSourceRepositoryPtrOutputWithContext(ctx)
}

type FunctionSourceRepositoryPtrInput interface {
	pulumi.Input

	ToFunctionSourceRepositoryPtrOutput() FunctionSourceRepositoryPtrOutput
	ToFunctionSourceRepositoryPtrOutputWithContext(context.Context) FunctionSourceRepositoryPtrOutput
}

type functionSourceRepositoryPtrType FunctionSourceRepositoryArgs

func FunctionSourceRepositoryPtr(v *FunctionSourceRepositoryArgs) FunctionSourceRepositoryPtrInput {	return (*functionSourceRepositoryPtrType)(v)
}

func (*functionSourceRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionSourceRepository)(nil)).Elem()
}

func (i *functionSourceRepositoryPtrType) ToFunctionSourceRepositoryPtrOutput() FunctionSourceRepositoryPtrOutput {
	return i.ToFunctionSourceRepositoryPtrOutputWithContext(context.Background())
}

func (i *functionSourceRepositoryPtrType) ToFunctionSourceRepositoryPtrOutputWithContext(ctx context.Context) FunctionSourceRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionSourceRepositoryPtrOutput)
}

type FunctionSourceRepositoryOutput struct { *pulumi.OutputState }

func (FunctionSourceRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionSourceRepository)(nil)).Elem()
}

func (o FunctionSourceRepositoryOutput) ToFunctionSourceRepositoryOutput() FunctionSourceRepositoryOutput {
	return o
}

func (o FunctionSourceRepositoryOutput) ToFunctionSourceRepositoryOutputWithContext(ctx context.Context) FunctionSourceRepositoryOutput {
	return o
}

func (o FunctionSourceRepositoryOutput) ToFunctionSourceRepositoryPtrOutput() FunctionSourceRepositoryPtrOutput {
	return o.ToFunctionSourceRepositoryPtrOutputWithContext(context.Background())
}

func (o FunctionSourceRepositoryOutput) ToFunctionSourceRepositoryPtrOutputWithContext(ctx context.Context) FunctionSourceRepositoryPtrOutput {
	return o.ApplyT(func(v FunctionSourceRepository) *FunctionSourceRepository {
		return &v
	}).(FunctionSourceRepositoryPtrOutput)
}
func (o FunctionSourceRepositoryOutput) DeployedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func (v FunctionSourceRepository) *string { return v.DeployedUrl }).(pulumi.StringPtrOutput)
}

func (o FunctionSourceRepositoryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func (v FunctionSourceRepository) string { return v.Url }).(pulumi.StringOutput)
}

type FunctionSourceRepositoryPtrOutput struct { *pulumi.OutputState}

func (FunctionSourceRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionSourceRepository)(nil)).Elem()
}

func (o FunctionSourceRepositoryPtrOutput) ToFunctionSourceRepositoryPtrOutput() FunctionSourceRepositoryPtrOutput {
	return o
}

func (o FunctionSourceRepositoryPtrOutput) ToFunctionSourceRepositoryPtrOutputWithContext(ctx context.Context) FunctionSourceRepositoryPtrOutput {
	return o
}

func (o FunctionSourceRepositoryPtrOutput) Elem() FunctionSourceRepositoryOutput {
	return o.ApplyT(func (v *FunctionSourceRepository) FunctionSourceRepository { return *v }).(FunctionSourceRepositoryOutput)
}

func (o FunctionSourceRepositoryPtrOutput) DeployedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func (v FunctionSourceRepository) *string { return v.DeployedUrl }).(pulumi.StringPtrOutput)
}

func (o FunctionSourceRepositoryPtrOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func (v FunctionSourceRepository) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionSourceRepositoryOutput{})
	pulumi.RegisterOutputType(FunctionSourceRepositoryPtrOutput{})
}
