// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package URLMapHostRule

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type URLMapHostRule struct {
	Description *string `pulumi:"description"`
	Hosts []string `pulumi:"hosts"`
	PathMatcher string `pulumi:"pathMatcher"`
}

type URLMapHostRuleInput interface {
	pulumi.Input

	ToURLMapHostRuleOutput() URLMapHostRuleOutput
	ToURLMapHostRuleOutputWithContext(context.Context) URLMapHostRuleOutput
}

type URLMapHostRuleArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Hosts pulumi.StringArrayInput `pulumi:"hosts"`
	PathMatcher pulumi.StringInput `pulumi:"pathMatcher"`
}

func (URLMapHostRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapHostRule)(nil)).Elem()
}

func (i URLMapHostRuleArgs) ToURLMapHostRuleOutput() URLMapHostRuleOutput {
	return i.ToURLMapHostRuleOutputWithContext(context.Background())
}

func (i URLMapHostRuleArgs) ToURLMapHostRuleOutputWithContext(ctx context.Context) URLMapHostRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapHostRuleOutput)
}

type URLMapHostRuleArrayInput interface {
	pulumi.Input

	ToURLMapHostRuleArrayOutput() URLMapHostRuleArrayOutput
	ToURLMapHostRuleArrayOutputWithContext(context.Context) URLMapHostRuleArrayOutput
}

type URLMapHostRuleArray []URLMapHostRuleInput

func (URLMapHostRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapHostRule)(nil)).Elem()
}

func (i URLMapHostRuleArray) ToURLMapHostRuleArrayOutput() URLMapHostRuleArrayOutput {
	return i.ToURLMapHostRuleArrayOutputWithContext(context.Background())
}

func (i URLMapHostRuleArray) ToURLMapHostRuleArrayOutputWithContext(ctx context.Context) URLMapHostRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapHostRuleArrayOutput)
}

type URLMapHostRuleOutput struct { *pulumi.OutputState }

func (URLMapHostRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapHostRule)(nil)).Elem()
}

func (o URLMapHostRuleOutput) ToURLMapHostRuleOutput() URLMapHostRuleOutput {
	return o
}

func (o URLMapHostRuleOutput) ToURLMapHostRuleOutputWithContext(ctx context.Context) URLMapHostRuleOutput {
	return o
}

func (o URLMapHostRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v URLMapHostRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o URLMapHostRuleOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func (v URLMapHostRule) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

func (o URLMapHostRuleOutput) PathMatcher() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapHostRule) string { return v.PathMatcher }).(pulumi.StringOutput)
}

type URLMapHostRuleArrayOutput struct { *pulumi.OutputState}

func (URLMapHostRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapHostRule)(nil)).Elem()
}

func (o URLMapHostRuleArrayOutput) ToURLMapHostRuleArrayOutput() URLMapHostRuleArrayOutput {
	return o
}

func (o URLMapHostRuleArrayOutput) ToURLMapHostRuleArrayOutputWithContext(ctx context.Context) URLMapHostRuleArrayOutput {
	return o
}

func (o URLMapHostRuleArrayOutput) Index(i pulumi.IntInput) URLMapHostRuleOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) URLMapHostRule {
		return vs[0].([]URLMapHostRule)[vs[1].(int)]
	}).(URLMapHostRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(URLMapHostRuleOutput{})
	pulumi.RegisterOutputType(URLMapHostRuleArrayOutput{})
}
