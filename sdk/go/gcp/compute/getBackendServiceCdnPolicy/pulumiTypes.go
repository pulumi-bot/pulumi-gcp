// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getBackendServiceCdnPolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/getBackendServiceCdnPolicyCacheKeyPolicy"
)

type GetBackendServiceCdnPolicy struct {
	CacheKeyPolicies []computegetBackendServiceCdnPolicyCacheKeyPolicy.GetBackendServiceCdnPolicyCacheKeyPolicy `pulumi:"cacheKeyPolicies"`
	SignedUrlCacheMaxAgeSec int `pulumi:"signedUrlCacheMaxAgeSec"`
}

type GetBackendServiceCdnPolicyInput interface {
	pulumi.Input

	ToGetBackendServiceCdnPolicyOutput() GetBackendServiceCdnPolicyOutput
	ToGetBackendServiceCdnPolicyOutputWithContext(context.Context) GetBackendServiceCdnPolicyOutput
}

type GetBackendServiceCdnPolicyArgs struct {
	CacheKeyPolicies computegetBackendServiceCdnPolicyCacheKeyPolicy.GetBackendServiceCdnPolicyCacheKeyPolicyArrayInput `pulumi:"cacheKeyPolicies"`
	SignedUrlCacheMaxAgeSec pulumi.IntInput `pulumi:"signedUrlCacheMaxAgeSec"`
}

func (GetBackendServiceCdnPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendServiceCdnPolicy)(nil)).Elem()
}

func (i GetBackendServiceCdnPolicyArgs) ToGetBackendServiceCdnPolicyOutput() GetBackendServiceCdnPolicyOutput {
	return i.ToGetBackendServiceCdnPolicyOutputWithContext(context.Background())
}

func (i GetBackendServiceCdnPolicyArgs) ToGetBackendServiceCdnPolicyOutputWithContext(ctx context.Context) GetBackendServiceCdnPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackendServiceCdnPolicyOutput)
}

type GetBackendServiceCdnPolicyArrayInput interface {
	pulumi.Input

	ToGetBackendServiceCdnPolicyArrayOutput() GetBackendServiceCdnPolicyArrayOutput
	ToGetBackendServiceCdnPolicyArrayOutputWithContext(context.Context) GetBackendServiceCdnPolicyArrayOutput
}

type GetBackendServiceCdnPolicyArray []GetBackendServiceCdnPolicyInput

func (GetBackendServiceCdnPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackendServiceCdnPolicy)(nil)).Elem()
}

func (i GetBackendServiceCdnPolicyArray) ToGetBackendServiceCdnPolicyArrayOutput() GetBackendServiceCdnPolicyArrayOutput {
	return i.ToGetBackendServiceCdnPolicyArrayOutputWithContext(context.Background())
}

func (i GetBackendServiceCdnPolicyArray) ToGetBackendServiceCdnPolicyArrayOutputWithContext(ctx context.Context) GetBackendServiceCdnPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackendServiceCdnPolicyArrayOutput)
}

type GetBackendServiceCdnPolicyOutput struct { *pulumi.OutputState }

func (GetBackendServiceCdnPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendServiceCdnPolicy)(nil)).Elem()
}

func (o GetBackendServiceCdnPolicyOutput) ToGetBackendServiceCdnPolicyOutput() GetBackendServiceCdnPolicyOutput {
	return o
}

func (o GetBackendServiceCdnPolicyOutput) ToGetBackendServiceCdnPolicyOutputWithContext(ctx context.Context) GetBackendServiceCdnPolicyOutput {
	return o
}

func (o GetBackendServiceCdnPolicyOutput) CacheKeyPolicies() computegetBackendServiceCdnPolicyCacheKeyPolicy.GetBackendServiceCdnPolicyCacheKeyPolicyArrayOutput {
	return o.ApplyT(func (v GetBackendServiceCdnPolicy) []computegetBackendServiceCdnPolicyCacheKeyPolicy.GetBackendServiceCdnPolicyCacheKeyPolicy { return v.CacheKeyPolicies }).(computegetBackendServiceCdnPolicyCacheKeyPolicy.GetBackendServiceCdnPolicyCacheKeyPolicyArrayOutput)
}

func (o GetBackendServiceCdnPolicyOutput) SignedUrlCacheMaxAgeSec() pulumi.IntOutput {
	return o.ApplyT(func (v GetBackendServiceCdnPolicy) int { return v.SignedUrlCacheMaxAgeSec }).(pulumi.IntOutput)
}

type GetBackendServiceCdnPolicyArrayOutput struct { *pulumi.OutputState}

func (GetBackendServiceCdnPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackendServiceCdnPolicy)(nil)).Elem()
}

func (o GetBackendServiceCdnPolicyArrayOutput) ToGetBackendServiceCdnPolicyArrayOutput() GetBackendServiceCdnPolicyArrayOutput {
	return o
}

func (o GetBackendServiceCdnPolicyArrayOutput) ToGetBackendServiceCdnPolicyArrayOutputWithContext(ctx context.Context) GetBackendServiceCdnPolicyArrayOutput {
	return o
}

func (o GetBackendServiceCdnPolicyArrayOutput) Index(i pulumi.IntInput) GetBackendServiceCdnPolicyOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetBackendServiceCdnPolicy {
		return vs[0].([]GetBackendServiceCdnPolicy)[vs[1].(int)]
	}).(GetBackendServiceCdnPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackendServiceCdnPolicyOutput{})
	pulumi.RegisterOutputType(GetBackendServiceCdnPolicyArrayOutput{})
}
