// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RouterBgp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/RouterBgpAdvertisedIpRange"
)

type RouterBgp struct {
	AdvertiseMode *string `pulumi:"advertiseMode"`
	AdvertisedGroups []string `pulumi:"advertisedGroups"`
	AdvertisedIpRanges []computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRange `pulumi:"advertisedIpRanges"`
	Asn int `pulumi:"asn"`
}

type RouterBgpInput interface {
	pulumi.Input

	ToRouterBgpOutput() RouterBgpOutput
	ToRouterBgpOutputWithContext(context.Context) RouterBgpOutput
}

type RouterBgpArgs struct {
	AdvertiseMode pulumi.StringPtrInput `pulumi:"advertiseMode"`
	AdvertisedGroups pulumi.StringArrayInput `pulumi:"advertisedGroups"`
	AdvertisedIpRanges computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRangeArrayInput `pulumi:"advertisedIpRanges"`
	Asn pulumi.IntInput `pulumi:"asn"`
}

func (RouterBgpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterBgp)(nil)).Elem()
}

func (i RouterBgpArgs) ToRouterBgpOutput() RouterBgpOutput {
	return i.ToRouterBgpOutputWithContext(context.Background())
}

func (i RouterBgpArgs) ToRouterBgpOutputWithContext(ctx context.Context) RouterBgpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBgpOutput)
}

func (i RouterBgpArgs) ToRouterBgpPtrOutput() RouterBgpPtrOutput {
	return i.ToRouterBgpPtrOutputWithContext(context.Background())
}

func (i RouterBgpArgs) ToRouterBgpPtrOutputWithContext(ctx context.Context) RouterBgpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBgpOutput).ToRouterBgpPtrOutputWithContext(ctx)
}

type RouterBgpPtrInput interface {
	pulumi.Input

	ToRouterBgpPtrOutput() RouterBgpPtrOutput
	ToRouterBgpPtrOutputWithContext(context.Context) RouterBgpPtrOutput
}

type routerBgpPtrType RouterBgpArgs

func RouterBgpPtr(v *RouterBgpArgs) RouterBgpPtrInput {	return (*routerBgpPtrType)(v)
}

func (*routerBgpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBgp)(nil)).Elem()
}

func (i *routerBgpPtrType) ToRouterBgpPtrOutput() RouterBgpPtrOutput {
	return i.ToRouterBgpPtrOutputWithContext(context.Background())
}

func (i *routerBgpPtrType) ToRouterBgpPtrOutputWithContext(ctx context.Context) RouterBgpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBgpPtrOutput)
}

type RouterBgpOutput struct { *pulumi.OutputState }

func (RouterBgpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterBgp)(nil)).Elem()
}

func (o RouterBgpOutput) ToRouterBgpOutput() RouterBgpOutput {
	return o
}

func (o RouterBgpOutput) ToRouterBgpOutputWithContext(ctx context.Context) RouterBgpOutput {
	return o
}

func (o RouterBgpOutput) ToRouterBgpPtrOutput() RouterBgpPtrOutput {
	return o.ToRouterBgpPtrOutputWithContext(context.Background())
}

func (o RouterBgpOutput) ToRouterBgpPtrOutputWithContext(ctx context.Context) RouterBgpPtrOutput {
	return o.ApplyT(func(v RouterBgp) *RouterBgp {
		return &v
	}).(RouterBgpPtrOutput)
}
func (o RouterBgpOutput) AdvertiseMode() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RouterBgp) *string { return v.AdvertiseMode }).(pulumi.StringPtrOutput)
}

func (o RouterBgpOutput) AdvertisedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RouterBgp) []string { return v.AdvertisedGroups }).(pulumi.StringArrayOutput)
}

func (o RouterBgpOutput) AdvertisedIpRanges() computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRangeArrayOutput {
	return o.ApplyT(func (v RouterBgp) []computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRange { return v.AdvertisedIpRanges }).(computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRangeArrayOutput)
}

func (o RouterBgpOutput) Asn() pulumi.IntOutput {
	return o.ApplyT(func (v RouterBgp) int { return v.Asn }).(pulumi.IntOutput)
}

type RouterBgpPtrOutput struct { *pulumi.OutputState}

func (RouterBgpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBgp)(nil)).Elem()
}

func (o RouterBgpPtrOutput) ToRouterBgpPtrOutput() RouterBgpPtrOutput {
	return o
}

func (o RouterBgpPtrOutput) ToRouterBgpPtrOutputWithContext(ctx context.Context) RouterBgpPtrOutput {
	return o
}

func (o RouterBgpPtrOutput) Elem() RouterBgpOutput {
	return o.ApplyT(func (v *RouterBgp) RouterBgp { return *v }).(RouterBgpOutput)
}

func (o RouterBgpPtrOutput) AdvertiseMode() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RouterBgp) *string { return v.AdvertiseMode }).(pulumi.StringPtrOutput)
}

func (o RouterBgpPtrOutput) AdvertisedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RouterBgp) []string { return v.AdvertisedGroups }).(pulumi.StringArrayOutput)
}

func (o RouterBgpPtrOutput) AdvertisedIpRanges() computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRangeArrayOutput {
	return o.ApplyT(func (v RouterBgp) []computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRange { return v.AdvertisedIpRanges }).(computeRouterBgpAdvertisedIpRange.RouterBgpAdvertisedIpRangeArrayOutput)
}

func (o RouterBgpPtrOutput) Asn() pulumi.IntOutput {
	return o.ApplyT(func (v RouterBgp) int { return v.Asn }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(RouterBgpOutput{})
	pulumi.RegisterOutputType(RouterBgpPtrOutput{})
}
