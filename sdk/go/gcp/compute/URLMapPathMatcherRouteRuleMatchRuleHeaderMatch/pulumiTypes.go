// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package URLMapPathMatcherRouteRuleMatchRuleHeaderMatch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch"
)

type URLMapPathMatcherRouteRuleMatchRuleHeaderMatch struct {
	ExactMatch *string `pulumi:"exactMatch"`
	HeaderName string `pulumi:"headerName"`
	InvertMatch *bool `pulumi:"invertMatch"`
	PrefixMatch *string `pulumi:"prefixMatch"`
	PresentMatch *bool `pulumi:"presentMatch"`
	RangeMatch *computeURLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch `pulumi:"rangeMatch"`
	RegexMatch *string `pulumi:"regexMatch"`
	SuffixMatch *string `pulumi:"suffixMatch"`
}

type URLMapPathMatcherRouteRuleMatchRuleHeaderMatchInput interface {
	pulumi.Input

	ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput() URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput
	ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutputWithContext(context.Context) URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput
}

type URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs struct {
	ExactMatch pulumi.StringPtrInput `pulumi:"exactMatch"`
	HeaderName pulumi.StringInput `pulumi:"headerName"`
	InvertMatch pulumi.BoolPtrInput `pulumi:"invertMatch"`
	PrefixMatch pulumi.StringPtrInput `pulumi:"prefixMatch"`
	PresentMatch pulumi.BoolPtrInput `pulumi:"presentMatch"`
	RangeMatch computeURLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchPtrInput `pulumi:"rangeMatch"`
	RegexMatch pulumi.StringPtrInput `pulumi:"regexMatch"`
	SuffixMatch pulumi.StringPtrInput `pulumi:"suffixMatch"`
}

func (URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapPathMatcherRouteRuleMatchRuleHeaderMatch)(nil)).Elem()
}

func (i URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput() URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput {
	return i.ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutputWithContext(context.Background())
}

func (i URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput)
}

type URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayInput interface {
	pulumi.Input

	ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput() URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput
	ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutputWithContext(context.Context) URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput
}

type URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArray []URLMapPathMatcherRouteRuleMatchRuleHeaderMatchInput

func (URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapPathMatcherRouteRuleMatchRuleHeaderMatch)(nil)).Elem()
}

func (i URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArray) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput() URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput {
	return i.ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutputWithContext(context.Background())
}

func (i URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArray) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput)
}

type URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput struct { *pulumi.OutputState }

func (URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapPathMatcherRouteRuleMatchRuleHeaderMatch)(nil)).Elem()
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput() URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) ExactMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) *string { return v.ExactMatch }).(pulumi.StringPtrOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) InvertMatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) *bool { return v.InvertMatch }).(pulumi.BoolPtrOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) PrefixMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) *string { return v.PrefixMatch }).(pulumi.StringPtrOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) PresentMatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) *bool { return v.PresentMatch }).(pulumi.BoolPtrOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) RangeMatch() computeURLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchPtrOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) *computeURLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch { return v.RangeMatch }).(computeURLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchPtrOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) RegexMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) *string { return v.RegexMatch }).(pulumi.StringPtrOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput) SuffixMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleHeaderMatch) *string { return v.SuffixMatch }).(pulumi.StringPtrOutput)
}

type URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput struct { *pulumi.OutputState}

func (URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapPathMatcherRouteRuleMatchRuleHeaderMatch)(nil)).Elem()
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput() URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput) ToURLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput) Index(i pulumi.IntInput) URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) URLMapPathMatcherRouteRuleMatchRuleHeaderMatch {
		return vs[0].([]URLMapPathMatcherRouteRuleMatchRuleHeaderMatch)[vs[1].(int)]
	}).(URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput)
}

func init() {
	pulumi.RegisterOutputType(URLMapPathMatcherRouteRuleMatchRuleHeaderMatchOutput{})
	pulumi.RegisterOutputType(URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArrayOutput{})
}
