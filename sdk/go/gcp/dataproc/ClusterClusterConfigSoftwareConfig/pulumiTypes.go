// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterClusterConfigSoftwareConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterClusterConfigSoftwareConfig struct {
	ImageVersion *string `pulumi:"imageVersion"`
	OptionalComponents []string `pulumi:"optionalComponents"`
	OverrideProperties map[string]string `pulumi:"overrideProperties"`
	Properties map[string]interface{} `pulumi:"properties"`
}

type ClusterClusterConfigSoftwareConfigInput interface {
	pulumi.Input

	ToClusterClusterConfigSoftwareConfigOutput() ClusterClusterConfigSoftwareConfigOutput
	ToClusterClusterConfigSoftwareConfigOutputWithContext(context.Context) ClusterClusterConfigSoftwareConfigOutput
}

type ClusterClusterConfigSoftwareConfigArgs struct {
	ImageVersion pulumi.StringPtrInput `pulumi:"imageVersion"`
	OptionalComponents pulumi.StringArrayInput `pulumi:"optionalComponents"`
	OverrideProperties pulumi.StringMapInput `pulumi:"overrideProperties"`
	Properties pulumi.MapInput `pulumi:"properties"`
}

func (ClusterClusterConfigSoftwareConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigSoftwareConfig)(nil)).Elem()
}

func (i ClusterClusterConfigSoftwareConfigArgs) ToClusterClusterConfigSoftwareConfigOutput() ClusterClusterConfigSoftwareConfigOutput {
	return i.ToClusterClusterConfigSoftwareConfigOutputWithContext(context.Background())
}

func (i ClusterClusterConfigSoftwareConfigArgs) ToClusterClusterConfigSoftwareConfigOutputWithContext(ctx context.Context) ClusterClusterConfigSoftwareConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigSoftwareConfigOutput)
}

func (i ClusterClusterConfigSoftwareConfigArgs) ToClusterClusterConfigSoftwareConfigPtrOutput() ClusterClusterConfigSoftwareConfigPtrOutput {
	return i.ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(context.Background())
}

func (i ClusterClusterConfigSoftwareConfigArgs) ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigSoftwareConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigSoftwareConfigOutput).ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(ctx)
}

type ClusterClusterConfigSoftwareConfigPtrInput interface {
	pulumi.Input

	ToClusterClusterConfigSoftwareConfigPtrOutput() ClusterClusterConfigSoftwareConfigPtrOutput
	ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(context.Context) ClusterClusterConfigSoftwareConfigPtrOutput
}

type clusterClusterConfigSoftwareConfigPtrType ClusterClusterConfigSoftwareConfigArgs

func ClusterClusterConfigSoftwareConfigPtr(v *ClusterClusterConfigSoftwareConfigArgs) ClusterClusterConfigSoftwareConfigPtrInput {	return (*clusterClusterConfigSoftwareConfigPtrType)(v)
}

func (*clusterClusterConfigSoftwareConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClusterConfigSoftwareConfig)(nil)).Elem()
}

func (i *clusterClusterConfigSoftwareConfigPtrType) ToClusterClusterConfigSoftwareConfigPtrOutput() ClusterClusterConfigSoftwareConfigPtrOutput {
	return i.ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(context.Background())
}

func (i *clusterClusterConfigSoftwareConfigPtrType) ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigSoftwareConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigSoftwareConfigPtrOutput)
}

type ClusterClusterConfigSoftwareConfigOutput struct { *pulumi.OutputState }

func (ClusterClusterConfigSoftwareConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigSoftwareConfig)(nil)).Elem()
}

func (o ClusterClusterConfigSoftwareConfigOutput) ToClusterClusterConfigSoftwareConfigOutput() ClusterClusterConfigSoftwareConfigOutput {
	return o
}

func (o ClusterClusterConfigSoftwareConfigOutput) ToClusterClusterConfigSoftwareConfigOutputWithContext(ctx context.Context) ClusterClusterConfigSoftwareConfigOutput {
	return o
}

func (o ClusterClusterConfigSoftwareConfigOutput) ToClusterClusterConfigSoftwareConfigPtrOutput() ClusterClusterConfigSoftwareConfigPtrOutput {
	return o.ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(context.Background())
}

func (o ClusterClusterConfigSoftwareConfigOutput) ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigSoftwareConfigPtrOutput {
	return o.ApplyT(func(v ClusterClusterConfigSoftwareConfig) *ClusterClusterConfigSoftwareConfig {
		return &v
	}).(ClusterClusterConfigSoftwareConfigPtrOutput)
}
func (o ClusterClusterConfigSoftwareConfigOutput) ImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) *string { return v.ImageVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigSoftwareConfigOutput) OptionalComponents() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) []string { return v.OptionalComponents }).(pulumi.StringArrayOutput)
}

func (o ClusterClusterConfigSoftwareConfigOutput) OverrideProperties() pulumi.StringMapOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) map[string]string { return v.OverrideProperties }).(pulumi.StringMapOutput)
}

func (o ClusterClusterConfigSoftwareConfigOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

type ClusterClusterConfigSoftwareConfigPtrOutput struct { *pulumi.OutputState}

func (ClusterClusterConfigSoftwareConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClusterConfigSoftwareConfig)(nil)).Elem()
}

func (o ClusterClusterConfigSoftwareConfigPtrOutput) ToClusterClusterConfigSoftwareConfigPtrOutput() ClusterClusterConfigSoftwareConfigPtrOutput {
	return o
}

func (o ClusterClusterConfigSoftwareConfigPtrOutput) ToClusterClusterConfigSoftwareConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigSoftwareConfigPtrOutput {
	return o
}

func (o ClusterClusterConfigSoftwareConfigPtrOutput) Elem() ClusterClusterConfigSoftwareConfigOutput {
	return o.ApplyT(func (v *ClusterClusterConfigSoftwareConfig) ClusterClusterConfigSoftwareConfig { return *v }).(ClusterClusterConfigSoftwareConfigOutput)
}

func (o ClusterClusterConfigSoftwareConfigPtrOutput) ImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) *string { return v.ImageVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigSoftwareConfigPtrOutput) OptionalComponents() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) []string { return v.OptionalComponents }).(pulumi.StringArrayOutput)
}

func (o ClusterClusterConfigSoftwareConfigPtrOutput) OverrideProperties() pulumi.StringMapOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) map[string]string { return v.OverrideProperties }).(pulumi.StringMapOutput)
}

func (o ClusterClusterConfigSoftwareConfigPtrOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func (v ClusterClusterConfigSoftwareConfig) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterClusterConfigSoftwareConfigOutput{})
	pulumi.RegisterOutputType(ClusterClusterConfigSoftwareConfigPtrOutput{})
}
