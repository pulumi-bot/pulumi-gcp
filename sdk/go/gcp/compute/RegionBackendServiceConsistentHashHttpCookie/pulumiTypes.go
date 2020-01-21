// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegionBackendServiceConsistentHashHttpCookie

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/RegionBackendServiceConsistentHashHttpCookieTtl"
)

type RegionBackendServiceConsistentHashHttpCookie struct {
	Name *string `pulumi:"name"`
	Path *string `pulumi:"path"`
	Ttl *computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtl `pulumi:"ttl"`
}

type RegionBackendServiceConsistentHashHttpCookieInput interface {
	pulumi.Input

	ToRegionBackendServiceConsistentHashHttpCookieOutput() RegionBackendServiceConsistentHashHttpCookieOutput
	ToRegionBackendServiceConsistentHashHttpCookieOutputWithContext(context.Context) RegionBackendServiceConsistentHashHttpCookieOutput
}

type RegionBackendServiceConsistentHashHttpCookieArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Path pulumi.StringPtrInput `pulumi:"path"`
	Ttl computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtlPtrInput `pulumi:"ttl"`
}

func (RegionBackendServiceConsistentHashHttpCookieArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionBackendServiceConsistentHashHttpCookie)(nil)).Elem()
}

func (i RegionBackendServiceConsistentHashHttpCookieArgs) ToRegionBackendServiceConsistentHashHttpCookieOutput() RegionBackendServiceConsistentHashHttpCookieOutput {
	return i.ToRegionBackendServiceConsistentHashHttpCookieOutputWithContext(context.Background())
}

func (i RegionBackendServiceConsistentHashHttpCookieArgs) ToRegionBackendServiceConsistentHashHttpCookieOutputWithContext(ctx context.Context) RegionBackendServiceConsistentHashHttpCookieOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionBackendServiceConsistentHashHttpCookieOutput)
}

func (i RegionBackendServiceConsistentHashHttpCookieArgs) ToRegionBackendServiceConsistentHashHttpCookiePtrOutput() RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return i.ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(context.Background())
}

func (i RegionBackendServiceConsistentHashHttpCookieArgs) ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(ctx context.Context) RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionBackendServiceConsistentHashHttpCookieOutput).ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(ctx)
}

type RegionBackendServiceConsistentHashHttpCookiePtrInput interface {
	pulumi.Input

	ToRegionBackendServiceConsistentHashHttpCookiePtrOutput() RegionBackendServiceConsistentHashHttpCookiePtrOutput
	ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(context.Context) RegionBackendServiceConsistentHashHttpCookiePtrOutput
}

type regionBackendServiceConsistentHashHttpCookiePtrType RegionBackendServiceConsistentHashHttpCookieArgs

func RegionBackendServiceConsistentHashHttpCookiePtr(v *RegionBackendServiceConsistentHashHttpCookieArgs) RegionBackendServiceConsistentHashHttpCookiePtrInput {	return (*regionBackendServiceConsistentHashHttpCookiePtrType)(v)
}

func (*regionBackendServiceConsistentHashHttpCookiePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionBackendServiceConsistentHashHttpCookie)(nil)).Elem()
}

func (i *regionBackendServiceConsistentHashHttpCookiePtrType) ToRegionBackendServiceConsistentHashHttpCookiePtrOutput() RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return i.ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(context.Background())
}

func (i *regionBackendServiceConsistentHashHttpCookiePtrType) ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(ctx context.Context) RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionBackendServiceConsistentHashHttpCookiePtrOutput)
}

type RegionBackendServiceConsistentHashHttpCookieOutput struct { *pulumi.OutputState }

func (RegionBackendServiceConsistentHashHttpCookieOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionBackendServiceConsistentHashHttpCookie)(nil)).Elem()
}

func (o RegionBackendServiceConsistentHashHttpCookieOutput) ToRegionBackendServiceConsistentHashHttpCookieOutput() RegionBackendServiceConsistentHashHttpCookieOutput {
	return o
}

func (o RegionBackendServiceConsistentHashHttpCookieOutput) ToRegionBackendServiceConsistentHashHttpCookieOutputWithContext(ctx context.Context) RegionBackendServiceConsistentHashHttpCookieOutput {
	return o
}

func (o RegionBackendServiceConsistentHashHttpCookieOutput) ToRegionBackendServiceConsistentHashHttpCookiePtrOutput() RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return o.ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(context.Background())
}

func (o RegionBackendServiceConsistentHashHttpCookieOutput) ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(ctx context.Context) RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return o.ApplyT(func(v RegionBackendServiceConsistentHashHttpCookie) *RegionBackendServiceConsistentHashHttpCookie {
		return &v
	}).(RegionBackendServiceConsistentHashHttpCookiePtrOutput)
}
func (o RegionBackendServiceConsistentHashHttpCookieOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionBackendServiceConsistentHashHttpCookie) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RegionBackendServiceConsistentHashHttpCookieOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionBackendServiceConsistentHashHttpCookie) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o RegionBackendServiceConsistentHashHttpCookieOutput) Ttl() computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return o.ApplyT(func (v RegionBackendServiceConsistentHashHttpCookie) *computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtl { return v.Ttl }).(computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtlPtrOutput)
}

type RegionBackendServiceConsistentHashHttpCookiePtrOutput struct { *pulumi.OutputState}

func (RegionBackendServiceConsistentHashHttpCookiePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionBackendServiceConsistentHashHttpCookie)(nil)).Elem()
}

func (o RegionBackendServiceConsistentHashHttpCookiePtrOutput) ToRegionBackendServiceConsistentHashHttpCookiePtrOutput() RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return o
}

func (o RegionBackendServiceConsistentHashHttpCookiePtrOutput) ToRegionBackendServiceConsistentHashHttpCookiePtrOutputWithContext(ctx context.Context) RegionBackendServiceConsistentHashHttpCookiePtrOutput {
	return o
}

func (o RegionBackendServiceConsistentHashHttpCookiePtrOutput) Elem() RegionBackendServiceConsistentHashHttpCookieOutput {
	return o.ApplyT(func (v *RegionBackendServiceConsistentHashHttpCookie) RegionBackendServiceConsistentHashHttpCookie { return *v }).(RegionBackendServiceConsistentHashHttpCookieOutput)
}

func (o RegionBackendServiceConsistentHashHttpCookiePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionBackendServiceConsistentHashHttpCookie) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RegionBackendServiceConsistentHashHttpCookiePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionBackendServiceConsistentHashHttpCookie) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o RegionBackendServiceConsistentHashHttpCookiePtrOutput) Ttl() computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtlPtrOutput {
	return o.ApplyT(func (v RegionBackendServiceConsistentHashHttpCookie) *computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtl { return v.Ttl }).(computeRegionBackendServiceConsistentHashHttpCookieTtl.RegionBackendServiceConsistentHashHttpCookieTtlPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionBackendServiceConsistentHashHttpCookieOutput{})
	pulumi.RegisterOutputType(RegionBackendServiceConsistentHashHttpCookiePtrOutput{})
}
