// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BackendServiceConsistentHashHttpCookieTtl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BackendServiceConsistentHashHttpCookieTtl struct {
	Nanos *int `pulumi:"nanos"`
	Seconds int `pulumi:"seconds"`
}

type BackendServiceConsistentHashHttpCookieTtlInput interface {
	pulumi.Input

	ToBackendServiceConsistentHashHttpCookieTtlOutput() BackendServiceConsistentHashHttpCookieTtlOutput
	ToBackendServiceConsistentHashHttpCookieTtlOutputWithContext(context.Context) BackendServiceConsistentHashHttpCookieTtlOutput
}

type BackendServiceConsistentHashHttpCookieTtlArgs struct {
	Nanos pulumi.IntPtrInput `pulumi:"nanos"`
	Seconds pulumi.IntInput `pulumi:"seconds"`
}

func (BackendServiceConsistentHashHttpCookieTtlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceConsistentHashHttpCookieTtl)(nil)).Elem()
}

func (i BackendServiceConsistentHashHttpCookieTtlArgs) ToBackendServiceConsistentHashHttpCookieTtlOutput() BackendServiceConsistentHashHttpCookieTtlOutput {
	return i.ToBackendServiceConsistentHashHttpCookieTtlOutputWithContext(context.Background())
}

func (i BackendServiceConsistentHashHttpCookieTtlArgs) ToBackendServiceConsistentHashHttpCookieTtlOutputWithContext(ctx context.Context) BackendServiceConsistentHashHttpCookieTtlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceConsistentHashHttpCookieTtlOutput)
}

func (i BackendServiceConsistentHashHttpCookieTtlArgs) ToBackendServiceConsistentHashHttpCookieTtlPtrOutput() BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return i.ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(context.Background())
}

func (i BackendServiceConsistentHashHttpCookieTtlArgs) ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(ctx context.Context) BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceConsistentHashHttpCookieTtlOutput).ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(ctx)
}

type BackendServiceConsistentHashHttpCookieTtlPtrInput interface {
	pulumi.Input

	ToBackendServiceConsistentHashHttpCookieTtlPtrOutput() BackendServiceConsistentHashHttpCookieTtlPtrOutput
	ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(context.Context) BackendServiceConsistentHashHttpCookieTtlPtrOutput
}

type backendServiceConsistentHashHttpCookieTtlPtrType BackendServiceConsistentHashHttpCookieTtlArgs

func BackendServiceConsistentHashHttpCookieTtlPtr(v *BackendServiceConsistentHashHttpCookieTtlArgs) BackendServiceConsistentHashHttpCookieTtlPtrInput {	return (*backendServiceConsistentHashHttpCookieTtlPtrType)(v)
}

func (*backendServiceConsistentHashHttpCookieTtlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceConsistentHashHttpCookieTtl)(nil)).Elem()
}

func (i *backendServiceConsistentHashHttpCookieTtlPtrType) ToBackendServiceConsistentHashHttpCookieTtlPtrOutput() BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return i.ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(context.Background())
}

func (i *backendServiceConsistentHashHttpCookieTtlPtrType) ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(ctx context.Context) BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceConsistentHashHttpCookieTtlPtrOutput)
}

type BackendServiceConsistentHashHttpCookieTtlOutput struct { *pulumi.OutputState }

func (BackendServiceConsistentHashHttpCookieTtlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceConsistentHashHttpCookieTtl)(nil)).Elem()
}

func (o BackendServiceConsistentHashHttpCookieTtlOutput) ToBackendServiceConsistentHashHttpCookieTtlOutput() BackendServiceConsistentHashHttpCookieTtlOutput {
	return o
}

func (o BackendServiceConsistentHashHttpCookieTtlOutput) ToBackendServiceConsistentHashHttpCookieTtlOutputWithContext(ctx context.Context) BackendServiceConsistentHashHttpCookieTtlOutput {
	return o
}

func (o BackendServiceConsistentHashHttpCookieTtlOutput) ToBackendServiceConsistentHashHttpCookieTtlPtrOutput() BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return o.ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(context.Background())
}

func (o BackendServiceConsistentHashHttpCookieTtlOutput) ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(ctx context.Context) BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return o.ApplyT(func(v BackendServiceConsistentHashHttpCookieTtl) *BackendServiceConsistentHashHttpCookieTtl {
		return &v
	}).(BackendServiceConsistentHashHttpCookieTtlPtrOutput)
}
func (o BackendServiceConsistentHashHttpCookieTtlOutput) Nanos() pulumi.IntPtrOutput {
	return o.ApplyT(func (v BackendServiceConsistentHashHttpCookieTtl) *int { return v.Nanos }).(pulumi.IntPtrOutput)
}

func (o BackendServiceConsistentHashHttpCookieTtlOutput) Seconds() pulumi.IntOutput {
	return o.ApplyT(func (v BackendServiceConsistentHashHttpCookieTtl) int { return v.Seconds }).(pulumi.IntOutput)
}

type BackendServiceConsistentHashHttpCookieTtlPtrOutput struct { *pulumi.OutputState}

func (BackendServiceConsistentHashHttpCookieTtlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceConsistentHashHttpCookieTtl)(nil)).Elem()
}

func (o BackendServiceConsistentHashHttpCookieTtlPtrOutput) ToBackendServiceConsistentHashHttpCookieTtlPtrOutput() BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return o
}

func (o BackendServiceConsistentHashHttpCookieTtlPtrOutput) ToBackendServiceConsistentHashHttpCookieTtlPtrOutputWithContext(ctx context.Context) BackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return o
}

func (o BackendServiceConsistentHashHttpCookieTtlPtrOutput) Elem() BackendServiceConsistentHashHttpCookieTtlOutput {
	return o.ApplyT(func (v *BackendServiceConsistentHashHttpCookieTtl) BackendServiceConsistentHashHttpCookieTtl { return *v }).(BackendServiceConsistentHashHttpCookieTtlOutput)
}

func (o BackendServiceConsistentHashHttpCookieTtlPtrOutput) Nanos() pulumi.IntPtrOutput {
	return o.ApplyT(func (v BackendServiceConsistentHashHttpCookieTtl) *int { return v.Nanos }).(pulumi.IntPtrOutput)
}

func (o BackendServiceConsistentHashHttpCookieTtlPtrOutput) Seconds() pulumi.IntOutput {
	return o.ApplyT(func (v BackendServiceConsistentHashHttpCookieTtl) int { return v.Seconds }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendServiceConsistentHashHttpCookieTtlOutput{})
	pulumi.RegisterOutputType(BackendServiceConsistentHashHttpCookieTtlPtrOutput{})
}
