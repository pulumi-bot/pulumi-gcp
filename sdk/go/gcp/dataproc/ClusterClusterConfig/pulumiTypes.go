// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterClusterConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigAutoscalingConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigEncryptionConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigGceClusterConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigInitializationAction"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigMasterConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigMasterConfigAccelerator"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigMasterConfigDiskConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigPreemptibleWorkerConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigPreemptibleWorkerConfigDiskConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigSecurityConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigSecurityConfigKerberosConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigSoftwareConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigWorkerConfig"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigWorkerConfigAccelerator"
	"https:/github.com/pulumi/pulumi-gcp/dataproc/ClusterClusterConfigWorkerConfigDiskConfig"
)

type ClusterClusterConfig struct {
	AutoscalingConfig *dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfig `pulumi:"autoscalingConfig"`
	Bucket *string `pulumi:"bucket"`
	EncryptionConfig *dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfig `pulumi:"encryptionConfig"`
	GceClusterConfig *dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfig `pulumi:"gceClusterConfig"`
	InitializationActions []dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationAction `pulumi:"initializationActions"`
	MasterConfig *dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfig `pulumi:"masterConfig"`
	PreemptibleWorkerConfig *dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfig `pulumi:"preemptibleWorkerConfig"`
	SecurityConfig *dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfig `pulumi:"securityConfig"`
	SoftwareConfig *dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfig `pulumi:"softwareConfig"`
	StagingBucket *string `pulumi:"stagingBucket"`
	WorkerConfig *dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfig `pulumi:"workerConfig"`
}

type ClusterClusterConfigInput interface {
	pulumi.Input

	ToClusterClusterConfigOutput() ClusterClusterConfigOutput
	ToClusterClusterConfigOutputWithContext(context.Context) ClusterClusterConfigOutput
}

type ClusterClusterConfigArgs struct {
	AutoscalingConfig dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfigPtrInput `pulumi:"autoscalingConfig"`
	Bucket pulumi.StringPtrInput `pulumi:"bucket"`
	EncryptionConfig dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfigPtrInput `pulumi:"encryptionConfig"`
	GceClusterConfig dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfigPtrInput `pulumi:"gceClusterConfig"`
	InitializationActions dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationActionArrayInput `pulumi:"initializationActions"`
	MasterConfig dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfigPtrInput `pulumi:"masterConfig"`
	PreemptibleWorkerConfig dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfigPtrInput `pulumi:"preemptibleWorkerConfig"`
	SecurityConfig dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfigPtrInput `pulumi:"securityConfig"`
	SoftwareConfig dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfigPtrInput `pulumi:"softwareConfig"`
	StagingBucket pulumi.StringPtrInput `pulumi:"stagingBucket"`
	WorkerConfig dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfigPtrInput `pulumi:"workerConfig"`
}

func (ClusterClusterConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfig)(nil)).Elem()
}

func (i ClusterClusterConfigArgs) ToClusterClusterConfigOutput() ClusterClusterConfigOutput {
	return i.ToClusterClusterConfigOutputWithContext(context.Background())
}

func (i ClusterClusterConfigArgs) ToClusterClusterConfigOutputWithContext(ctx context.Context) ClusterClusterConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigOutput)
}

func (i ClusterClusterConfigArgs) ToClusterClusterConfigPtrOutput() ClusterClusterConfigPtrOutput {
	return i.ToClusterClusterConfigPtrOutputWithContext(context.Background())
}

func (i ClusterClusterConfigArgs) ToClusterClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigOutput).ToClusterClusterConfigPtrOutputWithContext(ctx)
}

type ClusterClusterConfigPtrInput interface {
	pulumi.Input

	ToClusterClusterConfigPtrOutput() ClusterClusterConfigPtrOutput
	ToClusterClusterConfigPtrOutputWithContext(context.Context) ClusterClusterConfigPtrOutput
}

type clusterClusterConfigPtrType ClusterClusterConfigArgs

func ClusterClusterConfigPtr(v *ClusterClusterConfigArgs) ClusterClusterConfigPtrInput {	return (*clusterClusterConfigPtrType)(v)
}

func (*clusterClusterConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClusterConfig)(nil)).Elem()
}

func (i *clusterClusterConfigPtrType) ToClusterClusterConfigPtrOutput() ClusterClusterConfigPtrOutput {
	return i.ToClusterClusterConfigPtrOutputWithContext(context.Background())
}

func (i *clusterClusterConfigPtrType) ToClusterClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigPtrOutput)
}

type ClusterClusterConfigOutput struct { *pulumi.OutputState }

func (ClusterClusterConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfig)(nil)).Elem()
}

func (o ClusterClusterConfigOutput) ToClusterClusterConfigOutput() ClusterClusterConfigOutput {
	return o
}

func (o ClusterClusterConfigOutput) ToClusterClusterConfigOutputWithContext(ctx context.Context) ClusterClusterConfigOutput {
	return o
}

func (o ClusterClusterConfigOutput) ToClusterClusterConfigPtrOutput() ClusterClusterConfigPtrOutput {
	return o.ToClusterClusterConfigPtrOutputWithContext(context.Background())
}

func (o ClusterClusterConfigOutput) ToClusterClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigPtrOutput {
	return o.ApplyT(func(v ClusterClusterConfig) *ClusterClusterConfig {
		return &v
	}).(ClusterClusterConfigPtrOutput)
}
func (o ClusterClusterConfigOutput) AutoscalingConfig() dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfig { return v.AutoscalingConfig }).(dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfigPtrOutput)
}

func (o ClusterClusterConfigOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *string { return v.Bucket }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigOutput) EncryptionConfig() dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfig { return v.EncryptionConfig }).(dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfigPtrOutput)
}

func (o ClusterClusterConfigOutput) GceClusterConfig() dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfig { return v.GceClusterConfig }).(dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfigPtrOutput)
}

func (o ClusterClusterConfigOutput) InitializationActions() dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationActionArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfig) []dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationAction { return v.InitializationActions }).(dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationActionArrayOutput)
}

func (o ClusterClusterConfigOutput) MasterConfig() dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfig { return v.MasterConfig }).(dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfigPtrOutput)
}

func (o ClusterClusterConfigOutput) PreemptibleWorkerConfig() dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfig { return v.PreemptibleWorkerConfig }).(dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfigPtrOutput)
}

func (o ClusterClusterConfigOutput) SecurityConfig() dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfig { return v.SecurityConfig }).(dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfigPtrOutput)
}

func (o ClusterClusterConfigOutput) SoftwareConfig() dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfig { return v.SoftwareConfig }).(dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfigPtrOutput)
}

func (o ClusterClusterConfigOutput) StagingBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *string { return v.StagingBucket }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigOutput) WorkerConfig() dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfig { return v.WorkerConfig }).(dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfigPtrOutput)
}

type ClusterClusterConfigPtrOutput struct { *pulumi.OutputState}

func (ClusterClusterConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClusterConfig)(nil)).Elem()
}

func (o ClusterClusterConfigPtrOutput) ToClusterClusterConfigPtrOutput() ClusterClusterConfigPtrOutput {
	return o
}

func (o ClusterClusterConfigPtrOutput) ToClusterClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigPtrOutput {
	return o
}

func (o ClusterClusterConfigPtrOutput) Elem() ClusterClusterConfigOutput {
	return o.ApplyT(func (v *ClusterClusterConfig) ClusterClusterConfig { return *v }).(ClusterClusterConfigOutput)
}

func (o ClusterClusterConfigPtrOutput) AutoscalingConfig() dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfig { return v.AutoscalingConfig }).(dataprocClusterClusterConfigAutoscalingConfig.ClusterClusterConfigAutoscalingConfigPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *string { return v.Bucket }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) EncryptionConfig() dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfig { return v.EncryptionConfig }).(dataprocClusterClusterConfigEncryptionConfig.ClusterClusterConfigEncryptionConfigPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) GceClusterConfig() dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfig { return v.GceClusterConfig }).(dataprocClusterClusterConfigGceClusterConfig.ClusterClusterConfigGceClusterConfigPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) InitializationActions() dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationActionArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfig) []dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationAction { return v.InitializationActions }).(dataprocClusterClusterConfigInitializationAction.ClusterClusterConfigInitializationActionArrayOutput)
}

func (o ClusterClusterConfigPtrOutput) MasterConfig() dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfig { return v.MasterConfig }).(dataprocClusterClusterConfigMasterConfig.ClusterClusterConfigMasterConfigPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) PreemptibleWorkerConfig() dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfig { return v.PreemptibleWorkerConfig }).(dataprocClusterClusterConfigPreemptibleWorkerConfig.ClusterClusterConfigPreemptibleWorkerConfigPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) SecurityConfig() dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfig { return v.SecurityConfig }).(dataprocClusterClusterConfigSecurityConfig.ClusterClusterConfigSecurityConfigPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) SoftwareConfig() dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfig { return v.SoftwareConfig }).(dataprocClusterClusterConfigSoftwareConfig.ClusterClusterConfigSoftwareConfigPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) StagingBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *string { return v.StagingBucket }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigPtrOutput) WorkerConfig() dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfigPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfig) *dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfig { return v.WorkerConfig }).(dataprocClusterClusterConfigWorkerConfig.ClusterClusterConfigWorkerConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterClusterConfigOutput{})
	pulumi.RegisterOutputType(ClusterClusterConfigPtrOutput{})
}
