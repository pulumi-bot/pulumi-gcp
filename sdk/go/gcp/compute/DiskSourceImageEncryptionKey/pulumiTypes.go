// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package DiskSourceImageEncryptionKey

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DiskSourceImageEncryptionKey struct {
	KmsKeySelfLink *string `pulumi:"kmsKeySelfLink"`
	RawKey *string `pulumi:"rawKey"`
	Sha256 *string `pulumi:"sha256"`
}

type DiskSourceImageEncryptionKeyInput interface {
	pulumi.Input

	ToDiskSourceImageEncryptionKeyOutput() DiskSourceImageEncryptionKeyOutput
	ToDiskSourceImageEncryptionKeyOutputWithContext(context.Context) DiskSourceImageEncryptionKeyOutput
}

type DiskSourceImageEncryptionKeyArgs struct {
	KmsKeySelfLink pulumi.StringPtrInput `pulumi:"kmsKeySelfLink"`
	RawKey pulumi.StringPtrInput `pulumi:"rawKey"`
	Sha256 pulumi.StringPtrInput `pulumi:"sha256"`
}

func (DiskSourceImageEncryptionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSourceImageEncryptionKey)(nil)).Elem()
}

func (i DiskSourceImageEncryptionKeyArgs) ToDiskSourceImageEncryptionKeyOutput() DiskSourceImageEncryptionKeyOutput {
	return i.ToDiskSourceImageEncryptionKeyOutputWithContext(context.Background())
}

func (i DiskSourceImageEncryptionKeyArgs) ToDiskSourceImageEncryptionKeyOutputWithContext(ctx context.Context) DiskSourceImageEncryptionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSourceImageEncryptionKeyOutput)
}

func (i DiskSourceImageEncryptionKeyArgs) ToDiskSourceImageEncryptionKeyPtrOutput() DiskSourceImageEncryptionKeyPtrOutput {
	return i.ToDiskSourceImageEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i DiskSourceImageEncryptionKeyArgs) ToDiskSourceImageEncryptionKeyPtrOutputWithContext(ctx context.Context) DiskSourceImageEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSourceImageEncryptionKeyOutput).ToDiskSourceImageEncryptionKeyPtrOutputWithContext(ctx)
}

type DiskSourceImageEncryptionKeyPtrInput interface {
	pulumi.Input

	ToDiskSourceImageEncryptionKeyPtrOutput() DiskSourceImageEncryptionKeyPtrOutput
	ToDiskSourceImageEncryptionKeyPtrOutputWithContext(context.Context) DiskSourceImageEncryptionKeyPtrOutput
}

type diskSourceImageEncryptionKeyPtrType DiskSourceImageEncryptionKeyArgs

func DiskSourceImageEncryptionKeyPtr(v *DiskSourceImageEncryptionKeyArgs) DiskSourceImageEncryptionKeyPtrInput {	return (*diskSourceImageEncryptionKeyPtrType)(v)
}

func (*diskSourceImageEncryptionKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSourceImageEncryptionKey)(nil)).Elem()
}

func (i *diskSourceImageEncryptionKeyPtrType) ToDiskSourceImageEncryptionKeyPtrOutput() DiskSourceImageEncryptionKeyPtrOutput {
	return i.ToDiskSourceImageEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i *diskSourceImageEncryptionKeyPtrType) ToDiskSourceImageEncryptionKeyPtrOutputWithContext(ctx context.Context) DiskSourceImageEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSourceImageEncryptionKeyPtrOutput)
}

type DiskSourceImageEncryptionKeyOutput struct { *pulumi.OutputState }

func (DiskSourceImageEncryptionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSourceImageEncryptionKey)(nil)).Elem()
}

func (o DiskSourceImageEncryptionKeyOutput) ToDiskSourceImageEncryptionKeyOutput() DiskSourceImageEncryptionKeyOutput {
	return o
}

func (o DiskSourceImageEncryptionKeyOutput) ToDiskSourceImageEncryptionKeyOutputWithContext(ctx context.Context) DiskSourceImageEncryptionKeyOutput {
	return o
}

func (o DiskSourceImageEncryptionKeyOutput) ToDiskSourceImageEncryptionKeyPtrOutput() DiskSourceImageEncryptionKeyPtrOutput {
	return o.ToDiskSourceImageEncryptionKeyPtrOutputWithContext(context.Background())
}

func (o DiskSourceImageEncryptionKeyOutput) ToDiskSourceImageEncryptionKeyPtrOutputWithContext(ctx context.Context) DiskSourceImageEncryptionKeyPtrOutput {
	return o.ApplyT(func(v DiskSourceImageEncryptionKey) *DiskSourceImageEncryptionKey {
		return &v
	}).(DiskSourceImageEncryptionKeyPtrOutput)
}
func (o DiskSourceImageEncryptionKeyOutput) KmsKeySelfLink() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DiskSourceImageEncryptionKey) *string { return v.KmsKeySelfLink }).(pulumi.StringPtrOutput)
}

func (o DiskSourceImageEncryptionKeyOutput) RawKey() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DiskSourceImageEncryptionKey) *string { return v.RawKey }).(pulumi.StringPtrOutput)
}

func (o DiskSourceImageEncryptionKeyOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DiskSourceImageEncryptionKey) *string { return v.Sha256 }).(pulumi.StringPtrOutput)
}

type DiskSourceImageEncryptionKeyPtrOutput struct { *pulumi.OutputState}

func (DiskSourceImageEncryptionKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSourceImageEncryptionKey)(nil)).Elem()
}

func (o DiskSourceImageEncryptionKeyPtrOutput) ToDiskSourceImageEncryptionKeyPtrOutput() DiskSourceImageEncryptionKeyPtrOutput {
	return o
}

func (o DiskSourceImageEncryptionKeyPtrOutput) ToDiskSourceImageEncryptionKeyPtrOutputWithContext(ctx context.Context) DiskSourceImageEncryptionKeyPtrOutput {
	return o
}

func (o DiskSourceImageEncryptionKeyPtrOutput) Elem() DiskSourceImageEncryptionKeyOutput {
	return o.ApplyT(func (v *DiskSourceImageEncryptionKey) DiskSourceImageEncryptionKey { return *v }).(DiskSourceImageEncryptionKeyOutput)
}

func (o DiskSourceImageEncryptionKeyPtrOutput) KmsKeySelfLink() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DiskSourceImageEncryptionKey) *string { return v.KmsKeySelfLink }).(pulumi.StringPtrOutput)
}

func (o DiskSourceImageEncryptionKeyPtrOutput) RawKey() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DiskSourceImageEncryptionKey) *string { return v.RawKey }).(pulumi.StringPtrOutput)
}

func (o DiskSourceImageEncryptionKeyPtrOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DiskSourceImageEncryptionKey) *string { return v.Sha256 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskSourceImageEncryptionKeyOutput{})
	pulumi.RegisterOutputType(DiskSourceImageEncryptionKeyPtrOutput{})
}
