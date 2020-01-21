// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel struct {
	Name string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelInput interface {
	pulumi.Input

	ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput() URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput
	ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutputWithContext(context.Context) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput
}

type URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)(nil)).Elem()
}

func (i URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput() URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput {
	return i.ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutputWithContext(context.Background())
}

func (i URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput)
}

type URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayInput interface {
	pulumi.Input

	ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput() URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput
	ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutputWithContext(context.Context) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput
}

type URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray []URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelInput

func (URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)(nil)).Elem()
}

func (i URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput() URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput {
	return i.ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutputWithContext(context.Background())
}

func (i URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput)
}

type URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput struct { *pulumi.OutputState }

func (URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)(nil)).Elem()
}

func (o URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput() URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) string { return v.Name }).(pulumi.StringOutput)
}

func (o URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) string { return v.Value }).(pulumi.StringOutput)
}

type URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput struct { *pulumi.OutputState}

func (URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)(nil)).Elem()
}

func (o URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput() URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput) ToURLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutputWithContext(ctx context.Context) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput {
	return o
}

func (o URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput) Index(i pulumi.IntInput) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
		return vs[0].([]URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel)[vs[1].(int)]
	}).(URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput)
}

func init() {
	pulumi.RegisterOutputType(URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelOutput{})
	pulumi.RegisterOutputType(URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput{})
}
