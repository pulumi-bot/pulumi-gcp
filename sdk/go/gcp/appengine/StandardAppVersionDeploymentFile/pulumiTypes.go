// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package StandardAppVersionDeploymentFile

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type StandardAppVersionDeploymentFile struct {
	Name string `pulumi:"name"`
	Sha1Sum *string `pulumi:"sha1Sum"`
	SourceUrl string `pulumi:"sourceUrl"`
}

type StandardAppVersionDeploymentFileInput interface {
	pulumi.Input

	ToStandardAppVersionDeploymentFileOutput() StandardAppVersionDeploymentFileOutput
	ToStandardAppVersionDeploymentFileOutputWithContext(context.Context) StandardAppVersionDeploymentFileOutput
}

type StandardAppVersionDeploymentFileArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Sha1Sum pulumi.StringPtrInput `pulumi:"sha1Sum"`
	SourceUrl pulumi.StringInput `pulumi:"sourceUrl"`
}

func (StandardAppVersionDeploymentFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersionDeploymentFile)(nil)).Elem()
}

func (i StandardAppVersionDeploymentFileArgs) ToStandardAppVersionDeploymentFileOutput() StandardAppVersionDeploymentFileOutput {
	return i.ToStandardAppVersionDeploymentFileOutputWithContext(context.Background())
}

func (i StandardAppVersionDeploymentFileArgs) ToStandardAppVersionDeploymentFileOutputWithContext(ctx context.Context) StandardAppVersionDeploymentFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionDeploymentFileOutput)
}

type StandardAppVersionDeploymentFileArrayInput interface {
	pulumi.Input

	ToStandardAppVersionDeploymentFileArrayOutput() StandardAppVersionDeploymentFileArrayOutput
	ToStandardAppVersionDeploymentFileArrayOutputWithContext(context.Context) StandardAppVersionDeploymentFileArrayOutput
}

type StandardAppVersionDeploymentFileArray []StandardAppVersionDeploymentFileInput

func (StandardAppVersionDeploymentFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardAppVersionDeploymentFile)(nil)).Elem()
}

func (i StandardAppVersionDeploymentFileArray) ToStandardAppVersionDeploymentFileArrayOutput() StandardAppVersionDeploymentFileArrayOutput {
	return i.ToStandardAppVersionDeploymentFileArrayOutputWithContext(context.Background())
}

func (i StandardAppVersionDeploymentFileArray) ToStandardAppVersionDeploymentFileArrayOutputWithContext(ctx context.Context) StandardAppVersionDeploymentFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionDeploymentFileArrayOutput)
}

type StandardAppVersionDeploymentFileOutput struct { *pulumi.OutputState }

func (StandardAppVersionDeploymentFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersionDeploymentFile)(nil)).Elem()
}

func (o StandardAppVersionDeploymentFileOutput) ToStandardAppVersionDeploymentFileOutput() StandardAppVersionDeploymentFileOutput {
	return o
}

func (o StandardAppVersionDeploymentFileOutput) ToStandardAppVersionDeploymentFileOutputWithContext(ctx context.Context) StandardAppVersionDeploymentFileOutput {
	return o
}

func (o StandardAppVersionDeploymentFileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v StandardAppVersionDeploymentFile) string { return v.Name }).(pulumi.StringOutput)
}

func (o StandardAppVersionDeploymentFileOutput) Sha1Sum() pulumi.StringPtrOutput {
	return o.ApplyT(func (v StandardAppVersionDeploymentFile) *string { return v.Sha1Sum }).(pulumi.StringPtrOutput)
}

func (o StandardAppVersionDeploymentFileOutput) SourceUrl() pulumi.StringOutput {
	return o.ApplyT(func (v StandardAppVersionDeploymentFile) string { return v.SourceUrl }).(pulumi.StringOutput)
}

type StandardAppVersionDeploymentFileArrayOutput struct { *pulumi.OutputState}

func (StandardAppVersionDeploymentFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardAppVersionDeploymentFile)(nil)).Elem()
}

func (o StandardAppVersionDeploymentFileArrayOutput) ToStandardAppVersionDeploymentFileArrayOutput() StandardAppVersionDeploymentFileArrayOutput {
	return o
}

func (o StandardAppVersionDeploymentFileArrayOutput) ToStandardAppVersionDeploymentFileArrayOutputWithContext(ctx context.Context) StandardAppVersionDeploymentFileArrayOutput {
	return o
}

func (o StandardAppVersionDeploymentFileArrayOutput) Index(i pulumi.IntInput) StandardAppVersionDeploymentFileOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) StandardAppVersionDeploymentFile {
		return vs[0].([]StandardAppVersionDeploymentFile)[vs[1].(int)]
	}).(StandardAppVersionDeploymentFileOutput)
}

func init() {
	pulumi.RegisterOutputType(StandardAppVersionDeploymentFileOutput{})
	pulumi.RegisterOutputType(StandardAppVersionDeploymentFileArrayOutput{})
}
