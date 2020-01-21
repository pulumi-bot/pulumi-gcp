// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy struct {
	AllowCredentials *bool `pulumi:"allowCredentials"`
	AllowHeaders []string `pulumi:"allowHeaders"`
	AllowMethods []string `pulumi:"allowMethods"`
	AllowOriginRegexes []string `pulumi:"allowOriginRegexes"`
	AllowOrigins []string `pulumi:"allowOrigins"`
	Disabled bool `pulumi:"disabled"`
	ExposeHeaders []string `pulumi:"exposeHeaders"`
	MaxAge *int `pulumi:"maxAge"`
}

type RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyInput interface {
	pulumi.Input

	ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput
	ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputWithContext(context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput
}

type RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs struct {
	AllowCredentials pulumi.BoolPtrInput `pulumi:"allowCredentials"`
	AllowHeaders pulumi.StringArrayInput `pulumi:"allowHeaders"`
	AllowMethods pulumi.StringArrayInput `pulumi:"allowMethods"`
	AllowOriginRegexes pulumi.StringArrayInput `pulumi:"allowOriginRegexes"`
	AllowOrigins pulumi.StringArrayInput `pulumi:"allowOrigins"`
	Disabled pulumi.BoolInput `pulumi:"disabled"`
	ExposeHeaders pulumi.StringArrayInput `pulumi:"exposeHeaders"`
	MaxAge pulumi.IntPtrInput `pulumi:"maxAge"`
}

func (RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy)(nil)).Elem()
}

func (i RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput {
	return i.ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputWithContext(context.Background())
}

func (i RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput)
}

func (i RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return i.ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(context.Background())
}

func (i RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput).ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(ctx)
}

type RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrInput interface {
	pulumi.Input

	ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput
	ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput
}

type regionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrType RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs

func RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtr(v *RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrInput {	return (*regionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrType)(v)
}

func (*regionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy)(nil)).Elem()
}

func (i *regionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrType) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return i.ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(context.Background())
}

func (i *regionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrType) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput)
}

type RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput struct { *pulumi.OutputState }

func (RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy)(nil)).Elem()
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput {
	return o
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput {
	return o
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return o.ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(context.Background())
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return o.ApplyT(func(v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) *RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy {
		return &v
	}).(RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput)
}
func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) AllowHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowHeaders }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) AllowMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowMethods }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) AllowOriginRegexes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowOriginRegexes }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) AllowOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowOrigins }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) ExposeHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.ExposeHeaders }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

type RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput struct { *pulumi.OutputState}

func (RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy)(nil)).Elem()
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return o
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) ToRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput {
	return o
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) Elem() RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput {
	return o.ApplyT(func (v *RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy { return *v }).(RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) AllowHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowHeaders }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) AllowMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowMethods }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) AllowOriginRegexes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowOriginRegexes }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) AllowOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.AllowOrigins }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) ExposeHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) []string { return v.ExposeHeaders }).(pulumi.StringArrayOutput)
}

func (o RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutput{})
	pulumi.RegisterOutputType(RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyPtrOutput{})
}
