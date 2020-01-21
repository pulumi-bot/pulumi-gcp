// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package SnapshotSnapshotEncryptionKey

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SnapshotSnapshotEncryptionKey struct {
	RawKey string `pulumi:"rawKey"`
	Sha256 *string `pulumi:"sha256"`
}

type SnapshotSnapshotEncryptionKeyInput interface {
	pulumi.Input

	ToSnapshotSnapshotEncryptionKeyOutput() SnapshotSnapshotEncryptionKeyOutput
	ToSnapshotSnapshotEncryptionKeyOutputWithContext(context.Context) SnapshotSnapshotEncryptionKeyOutput
}

type SnapshotSnapshotEncryptionKeyArgs struct {
	RawKey pulumi.StringInput `pulumi:"rawKey"`
	Sha256 pulumi.StringPtrInput `pulumi:"sha256"`
}

func (SnapshotSnapshotEncryptionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSnapshotEncryptionKey)(nil)).Elem()
}

func (i SnapshotSnapshotEncryptionKeyArgs) ToSnapshotSnapshotEncryptionKeyOutput() SnapshotSnapshotEncryptionKeyOutput {
	return i.ToSnapshotSnapshotEncryptionKeyOutputWithContext(context.Background())
}

func (i SnapshotSnapshotEncryptionKeyArgs) ToSnapshotSnapshotEncryptionKeyOutputWithContext(ctx context.Context) SnapshotSnapshotEncryptionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSnapshotEncryptionKeyOutput)
}

func (i SnapshotSnapshotEncryptionKeyArgs) ToSnapshotSnapshotEncryptionKeyPtrOutput() SnapshotSnapshotEncryptionKeyPtrOutput {
	return i.ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i SnapshotSnapshotEncryptionKeyArgs) ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) SnapshotSnapshotEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSnapshotEncryptionKeyOutput).ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(ctx)
}

type SnapshotSnapshotEncryptionKeyPtrInput interface {
	pulumi.Input

	ToSnapshotSnapshotEncryptionKeyPtrOutput() SnapshotSnapshotEncryptionKeyPtrOutput
	ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(context.Context) SnapshotSnapshotEncryptionKeyPtrOutput
}

type snapshotSnapshotEncryptionKeyPtrType SnapshotSnapshotEncryptionKeyArgs

func SnapshotSnapshotEncryptionKeyPtr(v *SnapshotSnapshotEncryptionKeyArgs) SnapshotSnapshotEncryptionKeyPtrInput {	return (*snapshotSnapshotEncryptionKeyPtrType)(v)
}

func (*snapshotSnapshotEncryptionKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSnapshotEncryptionKey)(nil)).Elem()
}

func (i *snapshotSnapshotEncryptionKeyPtrType) ToSnapshotSnapshotEncryptionKeyPtrOutput() SnapshotSnapshotEncryptionKeyPtrOutput {
	return i.ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i *snapshotSnapshotEncryptionKeyPtrType) ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) SnapshotSnapshotEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSnapshotEncryptionKeyPtrOutput)
}

type SnapshotSnapshotEncryptionKeyOutput struct { *pulumi.OutputState }

func (SnapshotSnapshotEncryptionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSnapshotEncryptionKey)(nil)).Elem()
}

func (o SnapshotSnapshotEncryptionKeyOutput) ToSnapshotSnapshotEncryptionKeyOutput() SnapshotSnapshotEncryptionKeyOutput {
	return o
}

func (o SnapshotSnapshotEncryptionKeyOutput) ToSnapshotSnapshotEncryptionKeyOutputWithContext(ctx context.Context) SnapshotSnapshotEncryptionKeyOutput {
	return o
}

func (o SnapshotSnapshotEncryptionKeyOutput) ToSnapshotSnapshotEncryptionKeyPtrOutput() SnapshotSnapshotEncryptionKeyPtrOutput {
	return o.ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(context.Background())
}

func (o SnapshotSnapshotEncryptionKeyOutput) ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) SnapshotSnapshotEncryptionKeyPtrOutput {
	return o.ApplyT(func(v SnapshotSnapshotEncryptionKey) *SnapshotSnapshotEncryptionKey {
		return &v
	}).(SnapshotSnapshotEncryptionKeyPtrOutput)
}
func (o SnapshotSnapshotEncryptionKeyOutput) RawKey() pulumi.StringOutput {
	return o.ApplyT(func (v SnapshotSnapshotEncryptionKey) string { return v.RawKey }).(pulumi.StringOutput)
}

func (o SnapshotSnapshotEncryptionKeyOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SnapshotSnapshotEncryptionKey) *string { return v.Sha256 }).(pulumi.StringPtrOutput)
}

type SnapshotSnapshotEncryptionKeyPtrOutput struct { *pulumi.OutputState}

func (SnapshotSnapshotEncryptionKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSnapshotEncryptionKey)(nil)).Elem()
}

func (o SnapshotSnapshotEncryptionKeyPtrOutput) ToSnapshotSnapshotEncryptionKeyPtrOutput() SnapshotSnapshotEncryptionKeyPtrOutput {
	return o
}

func (o SnapshotSnapshotEncryptionKeyPtrOutput) ToSnapshotSnapshotEncryptionKeyPtrOutputWithContext(ctx context.Context) SnapshotSnapshotEncryptionKeyPtrOutput {
	return o
}

func (o SnapshotSnapshotEncryptionKeyPtrOutput) Elem() SnapshotSnapshotEncryptionKeyOutput {
	return o.ApplyT(func (v *SnapshotSnapshotEncryptionKey) SnapshotSnapshotEncryptionKey { return *v }).(SnapshotSnapshotEncryptionKeyOutput)
}

func (o SnapshotSnapshotEncryptionKeyPtrOutput) RawKey() pulumi.StringOutput {
	return o.ApplyT(func (v SnapshotSnapshotEncryptionKey) string { return v.RawKey }).(pulumi.StringOutput)
}

func (o SnapshotSnapshotEncryptionKeyPtrOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SnapshotSnapshotEncryptionKey) *string { return v.Sha256 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotSnapshotEncryptionKeyOutput{})
	pulumi.RegisterOutputType(SnapshotSnapshotEncryptionKeyPtrOutput{})
}
