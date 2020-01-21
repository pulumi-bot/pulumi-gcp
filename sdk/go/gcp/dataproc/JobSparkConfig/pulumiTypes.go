// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package JobSparkConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/JobSparkConfigLoggingConfig"
)

type JobSparkConfig struct {
	ArchiveUris []string `pulumi:"archiveUris"`
	Args []string `pulumi:"args"`
	FileUris []string `pulumi:"fileUris"`
	JarFileUris []string `pulumi:"jarFileUris"`
	LoggingConfig *dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfig `pulumi:"loggingConfig"`
	MainClass *string `pulumi:"mainClass"`
	MainJarFileUri *string `pulumi:"mainJarFileUri"`
	Properties map[string]string `pulumi:"properties"`
}

type JobSparkConfigInput interface {
	pulumi.Input

	ToJobSparkConfigOutput() JobSparkConfigOutput
	ToJobSparkConfigOutputWithContext(context.Context) JobSparkConfigOutput
}

type JobSparkConfigArgs struct {
	ArchiveUris pulumi.StringArrayInput `pulumi:"archiveUris"`
	Args pulumi.StringArrayInput `pulumi:"args"`
	FileUris pulumi.StringArrayInput `pulumi:"fileUris"`
	JarFileUris pulumi.StringArrayInput `pulumi:"jarFileUris"`
	LoggingConfig dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfigPtrInput `pulumi:"loggingConfig"`
	MainClass pulumi.StringPtrInput `pulumi:"mainClass"`
	MainJarFileUri pulumi.StringPtrInput `pulumi:"mainJarFileUri"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
}

func (JobSparkConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSparkConfig)(nil)).Elem()
}

func (i JobSparkConfigArgs) ToJobSparkConfigOutput() JobSparkConfigOutput {
	return i.ToJobSparkConfigOutputWithContext(context.Background())
}

func (i JobSparkConfigArgs) ToJobSparkConfigOutputWithContext(ctx context.Context) JobSparkConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSparkConfigOutput)
}

func (i JobSparkConfigArgs) ToJobSparkConfigPtrOutput() JobSparkConfigPtrOutput {
	return i.ToJobSparkConfigPtrOutputWithContext(context.Background())
}

func (i JobSparkConfigArgs) ToJobSparkConfigPtrOutputWithContext(ctx context.Context) JobSparkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSparkConfigOutput).ToJobSparkConfigPtrOutputWithContext(ctx)
}

type JobSparkConfigPtrInput interface {
	pulumi.Input

	ToJobSparkConfigPtrOutput() JobSparkConfigPtrOutput
	ToJobSparkConfigPtrOutputWithContext(context.Context) JobSparkConfigPtrOutput
}

type jobSparkConfigPtrType JobSparkConfigArgs

func JobSparkConfigPtr(v *JobSparkConfigArgs) JobSparkConfigPtrInput {	return (*jobSparkConfigPtrType)(v)
}

func (*jobSparkConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSparkConfig)(nil)).Elem()
}

func (i *jobSparkConfigPtrType) ToJobSparkConfigPtrOutput() JobSparkConfigPtrOutput {
	return i.ToJobSparkConfigPtrOutputWithContext(context.Background())
}

func (i *jobSparkConfigPtrType) ToJobSparkConfigPtrOutputWithContext(ctx context.Context) JobSparkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSparkConfigPtrOutput)
}

type JobSparkConfigOutput struct { *pulumi.OutputState }

func (JobSparkConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSparkConfig)(nil)).Elem()
}

func (o JobSparkConfigOutput) ToJobSparkConfigOutput() JobSparkConfigOutput {
	return o
}

func (o JobSparkConfigOutput) ToJobSparkConfigOutputWithContext(ctx context.Context) JobSparkConfigOutput {
	return o
}

func (o JobSparkConfigOutput) ToJobSparkConfigPtrOutput() JobSparkConfigPtrOutput {
	return o.ToJobSparkConfigPtrOutputWithContext(context.Background())
}

func (o JobSparkConfigOutput) ToJobSparkConfigPtrOutputWithContext(ctx context.Context) JobSparkConfigPtrOutput {
	return o.ApplyT(func(v JobSparkConfig) *JobSparkConfig {
		return &v
	}).(JobSparkConfigPtrOutput)
}
func (o JobSparkConfigOutput) ArchiveUris() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.ArchiveUris }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigOutput) FileUris() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.FileUris }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigOutput) JarFileUris() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.JarFileUris }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigOutput) LoggingConfig() dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfigPtrOutput {
	return o.ApplyT(func (v JobSparkConfig) *dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfig { return v.LoggingConfig }).(dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfigPtrOutput)
}

func (o JobSparkConfigOutput) MainClass() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobSparkConfig) *string { return v.MainClass }).(pulumi.StringPtrOutput)
}

func (o JobSparkConfigOutput) MainJarFileUri() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobSparkConfig) *string { return v.MainJarFileUri }).(pulumi.StringPtrOutput)
}

func (o JobSparkConfigOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func (v JobSparkConfig) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

type JobSparkConfigPtrOutput struct { *pulumi.OutputState}

func (JobSparkConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSparkConfig)(nil)).Elem()
}

func (o JobSparkConfigPtrOutput) ToJobSparkConfigPtrOutput() JobSparkConfigPtrOutput {
	return o
}

func (o JobSparkConfigPtrOutput) ToJobSparkConfigPtrOutputWithContext(ctx context.Context) JobSparkConfigPtrOutput {
	return o
}

func (o JobSparkConfigPtrOutput) Elem() JobSparkConfigOutput {
	return o.ApplyT(func (v *JobSparkConfig) JobSparkConfig { return *v }).(JobSparkConfigOutput)
}

func (o JobSparkConfigPtrOutput) ArchiveUris() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.ArchiveUris }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigPtrOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigPtrOutput) FileUris() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.FileUris }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigPtrOutput) JarFileUris() pulumi.StringArrayOutput {
	return o.ApplyT(func (v JobSparkConfig) []string { return v.JarFileUris }).(pulumi.StringArrayOutput)
}

func (o JobSparkConfigPtrOutput) LoggingConfig() dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfigPtrOutput {
	return o.ApplyT(func (v JobSparkConfig) *dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfig { return v.LoggingConfig }).(dataprocJobSparkConfigLoggingConfig.JobSparkConfigLoggingConfigPtrOutput)
}

func (o JobSparkConfigPtrOutput) MainClass() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobSparkConfig) *string { return v.MainClass }).(pulumi.StringPtrOutput)
}

func (o JobSparkConfigPtrOutput) MainJarFileUri() pulumi.StringPtrOutput {
	return o.ApplyT(func (v JobSparkConfig) *string { return v.MainJarFileUri }).(pulumi.StringPtrOutput)
}

func (o JobSparkConfigPtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func (v JobSparkConfig) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(JobSparkConfigOutput{})
	pulumi.RegisterOutputType(JobSparkConfigPtrOutput{})
}
