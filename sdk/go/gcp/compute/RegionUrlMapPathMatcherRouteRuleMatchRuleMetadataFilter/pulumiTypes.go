// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel"
)

type RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter struct {
	FilterLabels []computeRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel `pulumi:"filterLabels"`
	FilterMatchCriteria string `pulumi:"filterMatchCriteria"`
}

type RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterInput interface {
	pulumi.Input

	ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput() RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput
	ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutputWithContext(context.Context) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput
}

type RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs struct {
	FilterLabels computeRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayInput `pulumi:"filterLabels"`
	FilterMatchCriteria pulumi.StringInput `pulumi:"filterMatchCriteria"`
}

func (RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)(nil)).Elem()
}

func (i RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput() RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput {
	return i.ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutputWithContext(context.Background())
}

func (i RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput)
}

type RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayInput interface {
	pulumi.Input

	ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput() RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput
	ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutputWithContext(context.Context) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput
}

type RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray []RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterInput

func (RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)(nil)).Elem()
}

func (i RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput() RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput {
	return i.ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutputWithContext(context.Background())
}

func (i RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput)
}

type RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput struct { *pulumi.OutputState }

func (RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)(nil)).Elem()
}

func (o RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput() RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput {
	return o
}

func (o RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput {
	return o
}

func (o RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput) FilterLabels() computeRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) []computeRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel { return v.FilterLabels }).(computeRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel.RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArrayOutput)
}

func (o RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput) FilterMatchCriteria() pulumi.StringOutput {
	return o.ApplyT(func (v RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) string { return v.FilterMatchCriteria }).(pulumi.StringOutput)
}

type RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput struct { *pulumi.OutputState}

func (RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)(nil)).Elem()
}

func (o RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput() RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput {
	return o
}

func (o RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput) ToRegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutputWithContext(ctx context.Context) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput {
	return o
}

func (o RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput) Index(i pulumi.IntInput) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
		return vs[0].([]RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter)[vs[1].(int)]
	}).(RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterOutput{})
	pulumi.RegisterOutputType(RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArrayOutput{})
}
