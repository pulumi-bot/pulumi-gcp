// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package StandardAppVersionDeploymentZip

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type StandardAppVersionDeploymentZip struct {
	FilesCount *int `pulumi:"filesCount"`
	SourceUrl string `pulumi:"sourceUrl"`
}

type StandardAppVersionDeploymentZipInput interface {
	pulumi.Input

	ToStandardAppVersionDeploymentZipOutput() StandardAppVersionDeploymentZipOutput
	ToStandardAppVersionDeploymentZipOutputWithContext(context.Context) StandardAppVersionDeploymentZipOutput
}

type StandardAppVersionDeploymentZipArgs struct {
	FilesCount pulumi.IntPtrInput `pulumi:"filesCount"`
	SourceUrl pulumi.StringInput `pulumi:"sourceUrl"`
}

func (StandardAppVersionDeploymentZipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersionDeploymentZip)(nil)).Elem()
}

func (i StandardAppVersionDeploymentZipArgs) ToStandardAppVersionDeploymentZipOutput() StandardAppVersionDeploymentZipOutput {
	return i.ToStandardAppVersionDeploymentZipOutputWithContext(context.Background())
}

func (i StandardAppVersionDeploymentZipArgs) ToStandardAppVersionDeploymentZipOutputWithContext(ctx context.Context) StandardAppVersionDeploymentZipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionDeploymentZipOutput)
}

func (i StandardAppVersionDeploymentZipArgs) ToStandardAppVersionDeploymentZipPtrOutput() StandardAppVersionDeploymentZipPtrOutput {
	return i.ToStandardAppVersionDeploymentZipPtrOutputWithContext(context.Background())
}

func (i StandardAppVersionDeploymentZipArgs) ToStandardAppVersionDeploymentZipPtrOutputWithContext(ctx context.Context) StandardAppVersionDeploymentZipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionDeploymentZipOutput).ToStandardAppVersionDeploymentZipPtrOutputWithContext(ctx)
}

type StandardAppVersionDeploymentZipPtrInput interface {
	pulumi.Input

	ToStandardAppVersionDeploymentZipPtrOutput() StandardAppVersionDeploymentZipPtrOutput
	ToStandardAppVersionDeploymentZipPtrOutputWithContext(context.Context) StandardAppVersionDeploymentZipPtrOutput
}

type standardAppVersionDeploymentZipPtrType StandardAppVersionDeploymentZipArgs

func StandardAppVersionDeploymentZipPtr(v *StandardAppVersionDeploymentZipArgs) StandardAppVersionDeploymentZipPtrInput {	return (*standardAppVersionDeploymentZipPtrType)(v)
}

func (*standardAppVersionDeploymentZipPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StandardAppVersionDeploymentZip)(nil)).Elem()
}

func (i *standardAppVersionDeploymentZipPtrType) ToStandardAppVersionDeploymentZipPtrOutput() StandardAppVersionDeploymentZipPtrOutput {
	return i.ToStandardAppVersionDeploymentZipPtrOutputWithContext(context.Background())
}

func (i *standardAppVersionDeploymentZipPtrType) ToStandardAppVersionDeploymentZipPtrOutputWithContext(ctx context.Context) StandardAppVersionDeploymentZipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionDeploymentZipPtrOutput)
}

type StandardAppVersionDeploymentZipOutput struct { *pulumi.OutputState }

func (StandardAppVersionDeploymentZipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersionDeploymentZip)(nil)).Elem()
}

func (o StandardAppVersionDeploymentZipOutput) ToStandardAppVersionDeploymentZipOutput() StandardAppVersionDeploymentZipOutput {
	return o
}

func (o StandardAppVersionDeploymentZipOutput) ToStandardAppVersionDeploymentZipOutputWithContext(ctx context.Context) StandardAppVersionDeploymentZipOutput {
	return o
}

func (o StandardAppVersionDeploymentZipOutput) ToStandardAppVersionDeploymentZipPtrOutput() StandardAppVersionDeploymentZipPtrOutput {
	return o.ToStandardAppVersionDeploymentZipPtrOutputWithContext(context.Background())
}

func (o StandardAppVersionDeploymentZipOutput) ToStandardAppVersionDeploymentZipPtrOutputWithContext(ctx context.Context) StandardAppVersionDeploymentZipPtrOutput {
	return o.ApplyT(func(v StandardAppVersionDeploymentZip) *StandardAppVersionDeploymentZip {
		return &v
	}).(StandardAppVersionDeploymentZipPtrOutput)
}
func (o StandardAppVersionDeploymentZipOutput) FilesCount() pulumi.IntPtrOutput {
	return o.ApplyT(func (v StandardAppVersionDeploymentZip) *int { return v.FilesCount }).(pulumi.IntPtrOutput)
}

func (o StandardAppVersionDeploymentZipOutput) SourceUrl() pulumi.StringOutput {
	return o.ApplyT(func (v StandardAppVersionDeploymentZip) string { return v.SourceUrl }).(pulumi.StringOutput)
}

type StandardAppVersionDeploymentZipPtrOutput struct { *pulumi.OutputState}

func (StandardAppVersionDeploymentZipPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StandardAppVersionDeploymentZip)(nil)).Elem()
}

func (o StandardAppVersionDeploymentZipPtrOutput) ToStandardAppVersionDeploymentZipPtrOutput() StandardAppVersionDeploymentZipPtrOutput {
	return o
}

func (o StandardAppVersionDeploymentZipPtrOutput) ToStandardAppVersionDeploymentZipPtrOutputWithContext(ctx context.Context) StandardAppVersionDeploymentZipPtrOutput {
	return o
}

func (o StandardAppVersionDeploymentZipPtrOutput) Elem() StandardAppVersionDeploymentZipOutput {
	return o.ApplyT(func (v *StandardAppVersionDeploymentZip) StandardAppVersionDeploymentZip { return *v }).(StandardAppVersionDeploymentZipOutput)
}

func (o StandardAppVersionDeploymentZipPtrOutput) FilesCount() pulumi.IntPtrOutput {
	return o.ApplyT(func (v StandardAppVersionDeploymentZip) *int { return v.FilesCount }).(pulumi.IntPtrOutput)
}

func (o StandardAppVersionDeploymentZipPtrOutput) SourceUrl() pulumi.StringOutput {
	return o.ApplyT(func (v StandardAppVersionDeploymentZip) string { return v.SourceUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StandardAppVersionDeploymentZipOutput{})
	pulumi.RegisterOutputType(StandardAppVersionDeploymentZipPtrOutput{})
}
