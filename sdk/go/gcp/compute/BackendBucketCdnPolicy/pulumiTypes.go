// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BackendBucketCdnPolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BackendBucketCdnPolicy struct {
	SignedUrlCacheMaxAgeSec int `pulumi:"signedUrlCacheMaxAgeSec"`
}

type BackendBucketCdnPolicyInput interface {
	pulumi.Input

	ToBackendBucketCdnPolicyOutput() BackendBucketCdnPolicyOutput
	ToBackendBucketCdnPolicyOutputWithContext(context.Context) BackendBucketCdnPolicyOutput
}

type BackendBucketCdnPolicyArgs struct {
	SignedUrlCacheMaxAgeSec pulumi.IntInput `pulumi:"signedUrlCacheMaxAgeSec"`
}

func (BackendBucketCdnPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendBucketCdnPolicy)(nil)).Elem()
}

func (i BackendBucketCdnPolicyArgs) ToBackendBucketCdnPolicyOutput() BackendBucketCdnPolicyOutput {
	return i.ToBackendBucketCdnPolicyOutputWithContext(context.Background())
}

func (i BackendBucketCdnPolicyArgs) ToBackendBucketCdnPolicyOutputWithContext(ctx context.Context) BackendBucketCdnPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketCdnPolicyOutput)
}

func (i BackendBucketCdnPolicyArgs) ToBackendBucketCdnPolicyPtrOutput() BackendBucketCdnPolicyPtrOutput {
	return i.ToBackendBucketCdnPolicyPtrOutputWithContext(context.Background())
}

func (i BackendBucketCdnPolicyArgs) ToBackendBucketCdnPolicyPtrOutputWithContext(ctx context.Context) BackendBucketCdnPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketCdnPolicyOutput).ToBackendBucketCdnPolicyPtrOutputWithContext(ctx)
}

type BackendBucketCdnPolicyPtrInput interface {
	pulumi.Input

	ToBackendBucketCdnPolicyPtrOutput() BackendBucketCdnPolicyPtrOutput
	ToBackendBucketCdnPolicyPtrOutputWithContext(context.Context) BackendBucketCdnPolicyPtrOutput
}

type backendBucketCdnPolicyPtrType BackendBucketCdnPolicyArgs

func BackendBucketCdnPolicyPtr(v *BackendBucketCdnPolicyArgs) BackendBucketCdnPolicyPtrInput {	return (*backendBucketCdnPolicyPtrType)(v)
}

func (*backendBucketCdnPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketCdnPolicy)(nil)).Elem()
}

func (i *backendBucketCdnPolicyPtrType) ToBackendBucketCdnPolicyPtrOutput() BackendBucketCdnPolicyPtrOutput {
	return i.ToBackendBucketCdnPolicyPtrOutputWithContext(context.Background())
}

func (i *backendBucketCdnPolicyPtrType) ToBackendBucketCdnPolicyPtrOutputWithContext(ctx context.Context) BackendBucketCdnPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketCdnPolicyPtrOutput)
}

type BackendBucketCdnPolicyOutput struct { *pulumi.OutputState }

func (BackendBucketCdnPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendBucketCdnPolicy)(nil)).Elem()
}

func (o BackendBucketCdnPolicyOutput) ToBackendBucketCdnPolicyOutput() BackendBucketCdnPolicyOutput {
	return o
}

func (o BackendBucketCdnPolicyOutput) ToBackendBucketCdnPolicyOutputWithContext(ctx context.Context) BackendBucketCdnPolicyOutput {
	return o
}

func (o BackendBucketCdnPolicyOutput) ToBackendBucketCdnPolicyPtrOutput() BackendBucketCdnPolicyPtrOutput {
	return o.ToBackendBucketCdnPolicyPtrOutputWithContext(context.Background())
}

func (o BackendBucketCdnPolicyOutput) ToBackendBucketCdnPolicyPtrOutputWithContext(ctx context.Context) BackendBucketCdnPolicyPtrOutput {
	return o.ApplyT(func(v BackendBucketCdnPolicy) *BackendBucketCdnPolicy {
		return &v
	}).(BackendBucketCdnPolicyPtrOutput)
}
func (o BackendBucketCdnPolicyOutput) SignedUrlCacheMaxAgeSec() pulumi.IntOutput {
	return o.ApplyT(func (v BackendBucketCdnPolicy) int { return v.SignedUrlCacheMaxAgeSec }).(pulumi.IntOutput)
}

type BackendBucketCdnPolicyPtrOutput struct { *pulumi.OutputState}

func (BackendBucketCdnPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketCdnPolicy)(nil)).Elem()
}

func (o BackendBucketCdnPolicyPtrOutput) ToBackendBucketCdnPolicyPtrOutput() BackendBucketCdnPolicyPtrOutput {
	return o
}

func (o BackendBucketCdnPolicyPtrOutput) ToBackendBucketCdnPolicyPtrOutputWithContext(ctx context.Context) BackendBucketCdnPolicyPtrOutput {
	return o
}

func (o BackendBucketCdnPolicyPtrOutput) Elem() BackendBucketCdnPolicyOutput {
	return o.ApplyT(func (v *BackendBucketCdnPolicy) BackendBucketCdnPolicy { return *v }).(BackendBucketCdnPolicyOutput)
}

func (o BackendBucketCdnPolicyPtrOutput) SignedUrlCacheMaxAgeSec() pulumi.IntOutput {
	return o.ApplyT(func (v BackendBucketCdnPolicy) int { return v.SignedUrlCacheMaxAgeSec }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendBucketCdnPolicyOutput{})
	pulumi.RegisterOutputType(BackendBucketCdnPolicyPtrOutput{})
}
