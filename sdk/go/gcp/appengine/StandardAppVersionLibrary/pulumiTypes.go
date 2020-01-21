// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package StandardAppVersionLibrary

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type StandardAppVersionLibrary struct {
	Name *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}

type StandardAppVersionLibraryInput interface {
	pulumi.Input

	ToStandardAppVersionLibraryOutput() StandardAppVersionLibraryOutput
	ToStandardAppVersionLibraryOutputWithContext(context.Context) StandardAppVersionLibraryOutput
}

type StandardAppVersionLibraryArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (StandardAppVersionLibraryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersionLibrary)(nil)).Elem()
}

func (i StandardAppVersionLibraryArgs) ToStandardAppVersionLibraryOutput() StandardAppVersionLibraryOutput {
	return i.ToStandardAppVersionLibraryOutputWithContext(context.Background())
}

func (i StandardAppVersionLibraryArgs) ToStandardAppVersionLibraryOutputWithContext(ctx context.Context) StandardAppVersionLibraryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionLibraryOutput)
}

type StandardAppVersionLibraryArrayInput interface {
	pulumi.Input

	ToStandardAppVersionLibraryArrayOutput() StandardAppVersionLibraryArrayOutput
	ToStandardAppVersionLibraryArrayOutputWithContext(context.Context) StandardAppVersionLibraryArrayOutput
}

type StandardAppVersionLibraryArray []StandardAppVersionLibraryInput

func (StandardAppVersionLibraryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardAppVersionLibrary)(nil)).Elem()
}

func (i StandardAppVersionLibraryArray) ToStandardAppVersionLibraryArrayOutput() StandardAppVersionLibraryArrayOutput {
	return i.ToStandardAppVersionLibraryArrayOutputWithContext(context.Background())
}

func (i StandardAppVersionLibraryArray) ToStandardAppVersionLibraryArrayOutputWithContext(ctx context.Context) StandardAppVersionLibraryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionLibraryArrayOutput)
}

type StandardAppVersionLibraryOutput struct { *pulumi.OutputState }

func (StandardAppVersionLibraryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersionLibrary)(nil)).Elem()
}

func (o StandardAppVersionLibraryOutput) ToStandardAppVersionLibraryOutput() StandardAppVersionLibraryOutput {
	return o
}

func (o StandardAppVersionLibraryOutput) ToStandardAppVersionLibraryOutputWithContext(ctx context.Context) StandardAppVersionLibraryOutput {
	return o
}

func (o StandardAppVersionLibraryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v StandardAppVersionLibrary) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StandardAppVersionLibraryOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v StandardAppVersionLibrary) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type StandardAppVersionLibraryArrayOutput struct { *pulumi.OutputState}

func (StandardAppVersionLibraryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardAppVersionLibrary)(nil)).Elem()
}

func (o StandardAppVersionLibraryArrayOutput) ToStandardAppVersionLibraryArrayOutput() StandardAppVersionLibraryArrayOutput {
	return o
}

func (o StandardAppVersionLibraryArrayOutput) ToStandardAppVersionLibraryArrayOutputWithContext(ctx context.Context) StandardAppVersionLibraryArrayOutput {
	return o
}

func (o StandardAppVersionLibraryArrayOutput) Index(i pulumi.IntInput) StandardAppVersionLibraryOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) StandardAppVersionLibrary {
		return vs[0].([]StandardAppVersionLibrary)[vs[1].(int)]
	}).(StandardAppVersionLibraryOutput)
}

func init() {
	pulumi.RegisterOutputType(StandardAppVersionLibraryOutput{})
	pulumi.RegisterOutputType(StandardAppVersionLibraryArrayOutput{})
}
