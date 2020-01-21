// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegionDiskSourceSnapshotEncryptionKey

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegionDiskSourceSnapshotEncryptionKey struct {
	KmsKeyName *string `pulumi:"kmsKeyName"`
	RawKey *string `pulumi:"rawKey"`
	Sha256 *string `pulumi:"sha256"`
}

type RegionDiskSourceSnapshotEncryptionKeyInput interface {
	pulumi.Input

	ToRegionDiskSourceSnapshotEncryptionKeyOutput() RegionDiskSourceSnapshotEncryptionKeyOutput
	ToRegionDiskSourceSnapshotEncryptionKeyOutputWithContext(context.Context) RegionDiskSourceSnapshotEncryptionKeyOutput
}

type RegionDiskSourceSnapshotEncryptionKeyArgs struct {
	KmsKeyName pulumi.StringPtrInput `pulumi:"kmsKeyName"`
	RawKey pulumi.StringPtrInput `pulumi:"rawKey"`
	Sha256 pulumi.StringPtrInput `pulumi:"sha256"`
}

func (RegionDiskSourceSnapshotEncryptionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDiskSourceSnapshotEncryptionKey)(nil)).Elem()
}

func (i RegionDiskSourceSnapshotEncryptionKeyArgs) ToRegionDiskSourceSnapshotEncryptionKeyOutput() RegionDiskSourceSnapshotEncryptionKeyOutput {
	return i.ToRegionDiskSourceSnapshotEncryptionKeyOutputWithContext(context.Background())
}

func (i RegionDiskSourceSnapshotEncryptionKeyArgs) ToRegionDiskSourceSnapshotEncryptionKeyOutputWithContext(ctx context.Context) RegionDiskSourceSnapshotEncryptionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskSourceSnapshotEncryptionKeyOutput)
}

func (i RegionDiskSourceSnapshotEncryptionKeyArgs) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutput() RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return i.ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i RegionDiskSourceSnapshotEncryptionKeyArgs) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskSourceSnapshotEncryptionKeyOutput).ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(ctx)
}

type RegionDiskSourceSnapshotEncryptionKeyPtrInput interface {
	pulumi.Input

	ToRegionDiskSourceSnapshotEncryptionKeyPtrOutput() RegionDiskSourceSnapshotEncryptionKeyPtrOutput
	ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(context.Context) RegionDiskSourceSnapshotEncryptionKeyPtrOutput
}

type regionDiskSourceSnapshotEncryptionKeyPtrType RegionDiskSourceSnapshotEncryptionKeyArgs

func RegionDiskSourceSnapshotEncryptionKeyPtr(v *RegionDiskSourceSnapshotEncryptionKeyArgs) RegionDiskSourceSnapshotEncryptionKeyPtrInput {	return (*regionDiskSourceSnapshotEncryptionKeyPtrType)(v)
}

func (*regionDiskSourceSnapshotEncryptionKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDiskSourceSnapshotEncryptionKey)(nil)).Elem()
}

func (i *regionDiskSourceSnapshotEncryptionKeyPtrType) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutput() RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return i.ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i *regionDiskSourceSnapshotEncryptionKeyPtrType) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskSourceSnapshotEncryptionKeyPtrOutput)
}

type RegionDiskSourceSnapshotEncryptionKeyOutput struct { *pulumi.OutputState }

func (RegionDiskSourceSnapshotEncryptionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDiskSourceSnapshotEncryptionKey)(nil)).Elem()
}

func (o RegionDiskSourceSnapshotEncryptionKeyOutput) ToRegionDiskSourceSnapshotEncryptionKeyOutput() RegionDiskSourceSnapshotEncryptionKeyOutput {
	return o
}

func (o RegionDiskSourceSnapshotEncryptionKeyOutput) ToRegionDiskSourceSnapshotEncryptionKeyOutputWithContext(ctx context.Context) RegionDiskSourceSnapshotEncryptionKeyOutput {
	return o
}

func (o RegionDiskSourceSnapshotEncryptionKeyOutput) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutput() RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return o.ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(context.Background())
}

func (o RegionDiskSourceSnapshotEncryptionKeyOutput) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return o.ApplyT(func(v RegionDiskSourceSnapshotEncryptionKey) *RegionDiskSourceSnapshotEncryptionKey {
		return &v
	}).(RegionDiskSourceSnapshotEncryptionKeyPtrOutput)
}
func (o RegionDiskSourceSnapshotEncryptionKeyOutput) KmsKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionDiskSourceSnapshotEncryptionKey) *string { return v.KmsKeyName }).(pulumi.StringPtrOutput)
}

func (o RegionDiskSourceSnapshotEncryptionKeyOutput) RawKey() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionDiskSourceSnapshotEncryptionKey) *string { return v.RawKey }).(pulumi.StringPtrOutput)
}

func (o RegionDiskSourceSnapshotEncryptionKeyOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionDiskSourceSnapshotEncryptionKey) *string { return v.Sha256 }).(pulumi.StringPtrOutput)
}

type RegionDiskSourceSnapshotEncryptionKeyPtrOutput struct { *pulumi.OutputState}

func (RegionDiskSourceSnapshotEncryptionKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDiskSourceSnapshotEncryptionKey)(nil)).Elem()
}

func (o RegionDiskSourceSnapshotEncryptionKeyPtrOutput) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutput() RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return o
}

func (o RegionDiskSourceSnapshotEncryptionKeyPtrOutput) ToRegionDiskSourceSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) RegionDiskSourceSnapshotEncryptionKeyPtrOutput {
	return o
}

func (o RegionDiskSourceSnapshotEncryptionKeyPtrOutput) Elem() RegionDiskSourceSnapshotEncryptionKeyOutput {
	return o.ApplyT(func (v *RegionDiskSourceSnapshotEncryptionKey) RegionDiskSourceSnapshotEncryptionKey { return *v }).(RegionDiskSourceSnapshotEncryptionKeyOutput)
}

func (o RegionDiskSourceSnapshotEncryptionKeyPtrOutput) KmsKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionDiskSourceSnapshotEncryptionKey) *string { return v.KmsKeyName }).(pulumi.StringPtrOutput)
}

func (o RegionDiskSourceSnapshotEncryptionKeyPtrOutput) RawKey() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionDiskSourceSnapshotEncryptionKey) *string { return v.RawKey }).(pulumi.StringPtrOutput)
}

func (o RegionDiskSourceSnapshotEncryptionKeyPtrOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionDiskSourceSnapshotEncryptionKey) *string { return v.Sha256 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionDiskSourceSnapshotEncryptionKeyOutput{})
	pulumi.RegisterOutputType(RegionDiskSourceSnapshotEncryptionKeyPtrOutput{})
}
