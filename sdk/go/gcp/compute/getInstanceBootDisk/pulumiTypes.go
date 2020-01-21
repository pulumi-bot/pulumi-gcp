// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getInstanceBootDisk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/getInstanceBootDiskInitializeParam"
)

type GetInstanceBootDisk struct {
	// Whether the disk will be auto-deleted when the instance is deleted.
	AutoDelete bool `pulumi:"autoDelete"`
	// Name with which the attached disk is accessible
	// under `/dev/disk/by-id/`
	DeviceName string `pulumi:"deviceName"`
	DiskEncryptionKeyRaw string `pulumi:"diskEncryptionKeyRaw"`
	DiskEncryptionKeySha256 string `pulumi:"diskEncryptionKeySha256"`
	// Parameters with which a disk was created alongside the instance.
	// Structure is documented below.
	InitializeParams []computegetInstanceBootDiskInitializeParam.GetInstanceBootDiskInitializeParam `pulumi:"initializeParams"`
	KmsKeySelfLink string `pulumi:"kmsKeySelfLink"`
	// Read/write mode for the disk. One of `"READ_ONLY"` or `"READ_WRITE"`.
	Mode string `pulumi:"mode"`
	// The name or selfLink of the disk attached to this instance.
	Source string `pulumi:"source"`
}

type GetInstanceBootDiskInput interface {
	pulumi.Input

	ToGetInstanceBootDiskOutput() GetInstanceBootDiskOutput
	ToGetInstanceBootDiskOutputWithContext(context.Context) GetInstanceBootDiskOutput
}

type GetInstanceBootDiskArgs struct {
	// Whether the disk will be auto-deleted when the instance is deleted.
	AutoDelete pulumi.BoolInput `pulumi:"autoDelete"`
	// Name with which the attached disk is accessible
	// under `/dev/disk/by-id/`
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	DiskEncryptionKeyRaw pulumi.StringInput `pulumi:"diskEncryptionKeyRaw"`
	DiskEncryptionKeySha256 pulumi.StringInput `pulumi:"diskEncryptionKeySha256"`
	// Parameters with which a disk was created alongside the instance.
	// Structure is documented below.
	InitializeParams computegetInstanceBootDiskInitializeParam.GetInstanceBootDiskInitializeParamArrayInput `pulumi:"initializeParams"`
	KmsKeySelfLink pulumi.StringInput `pulumi:"kmsKeySelfLink"`
	// Read/write mode for the disk. One of `"READ_ONLY"` or `"READ_WRITE"`.
	Mode pulumi.StringInput `pulumi:"mode"`
	// The name or selfLink of the disk attached to this instance.
	Source pulumi.StringInput `pulumi:"source"`
}

func (GetInstanceBootDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceBootDisk)(nil)).Elem()
}

func (i GetInstanceBootDiskArgs) ToGetInstanceBootDiskOutput() GetInstanceBootDiskOutput {
	return i.ToGetInstanceBootDiskOutputWithContext(context.Background())
}

func (i GetInstanceBootDiskArgs) ToGetInstanceBootDiskOutputWithContext(ctx context.Context) GetInstanceBootDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceBootDiskOutput)
}

type GetInstanceBootDiskArrayInput interface {
	pulumi.Input

	ToGetInstanceBootDiskArrayOutput() GetInstanceBootDiskArrayOutput
	ToGetInstanceBootDiskArrayOutputWithContext(context.Context) GetInstanceBootDiskArrayOutput
}

type GetInstanceBootDiskArray []GetInstanceBootDiskInput

func (GetInstanceBootDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceBootDisk)(nil)).Elem()
}

func (i GetInstanceBootDiskArray) ToGetInstanceBootDiskArrayOutput() GetInstanceBootDiskArrayOutput {
	return i.ToGetInstanceBootDiskArrayOutputWithContext(context.Background())
}

func (i GetInstanceBootDiskArray) ToGetInstanceBootDiskArrayOutputWithContext(ctx context.Context) GetInstanceBootDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceBootDiskArrayOutput)
}

type GetInstanceBootDiskOutput struct { *pulumi.OutputState }

func (GetInstanceBootDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceBootDisk)(nil)).Elem()
}

func (o GetInstanceBootDiskOutput) ToGetInstanceBootDiskOutput() GetInstanceBootDiskOutput {
	return o
}

func (o GetInstanceBootDiskOutput) ToGetInstanceBootDiskOutputWithContext(ctx context.Context) GetInstanceBootDiskOutput {
	return o
}

// Whether the disk will be auto-deleted when the instance is deleted.
func (o GetInstanceBootDiskOutput) AutoDelete() pulumi.BoolOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) bool { return v.AutoDelete }).(pulumi.BoolOutput)
}

// Name with which the attached disk is accessible
// under `/dev/disk/by-id/`
func (o GetInstanceBootDiskOutput) DeviceName() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) string { return v.DeviceName }).(pulumi.StringOutput)
}

func (o GetInstanceBootDiskOutput) DiskEncryptionKeyRaw() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) string { return v.DiskEncryptionKeyRaw }).(pulumi.StringOutput)
}

func (o GetInstanceBootDiskOutput) DiskEncryptionKeySha256() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) string { return v.DiskEncryptionKeySha256 }).(pulumi.StringOutput)
}

// Parameters with which a disk was created alongside the instance.
// Structure is documented below.
func (o GetInstanceBootDiskOutput) InitializeParams() computegetInstanceBootDiskInitializeParam.GetInstanceBootDiskInitializeParamArrayOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) []computegetInstanceBootDiskInitializeParam.GetInstanceBootDiskInitializeParam { return v.InitializeParams }).(computegetInstanceBootDiskInitializeParam.GetInstanceBootDiskInitializeParamArrayOutput)
}

func (o GetInstanceBootDiskOutput) KmsKeySelfLink() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) string { return v.KmsKeySelfLink }).(pulumi.StringOutput)
}

// Read/write mode for the disk. One of `"READ_ONLY"` or `"READ_WRITE"`.
func (o GetInstanceBootDiskOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) string { return v.Mode }).(pulumi.StringOutput)
}

// The name or selfLink of the disk attached to this instance.
func (o GetInstanceBootDiskOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceBootDisk) string { return v.Source }).(pulumi.StringOutput)
}

type GetInstanceBootDiskArrayOutput struct { *pulumi.OutputState}

func (GetInstanceBootDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceBootDisk)(nil)).Elem()
}

func (o GetInstanceBootDiskArrayOutput) ToGetInstanceBootDiskArrayOutput() GetInstanceBootDiskArrayOutput {
	return o
}

func (o GetInstanceBootDiskArrayOutput) ToGetInstanceBootDiskArrayOutputWithContext(ctx context.Context) GetInstanceBootDiskArrayOutput {
	return o
}

func (o GetInstanceBootDiskArrayOutput) Index(i pulumi.IntInput) GetInstanceBootDiskOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetInstanceBootDisk {
		return vs[0].([]GetInstanceBootDisk)[vs[1].(int)]
	}).(GetInstanceBootDiskOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceBootDiskOutput{})
	pulumi.RegisterOutputType(GetInstanceBootDiskArrayOutput{})
}
