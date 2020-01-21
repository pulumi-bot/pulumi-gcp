// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterClusterConfigGceClusterConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterClusterConfigGceClusterConfig struct {
	InternalIpOnly *bool `pulumi:"internalIpOnly"`
	Metadata map[string]string `pulumi:"metadata"`
	Network *string `pulumi:"network"`
	ServiceAccount *string `pulumi:"serviceAccount"`
	ServiceAccountScopes []string `pulumi:"serviceAccountScopes"`
	Subnetwork *string `pulumi:"subnetwork"`
	Tags []string `pulumi:"tags"`
	Zone *string `pulumi:"zone"`
}

type ClusterClusterConfigGceClusterConfigInput interface {
	pulumi.Input

	ToClusterClusterConfigGceClusterConfigOutput() ClusterClusterConfigGceClusterConfigOutput
	ToClusterClusterConfigGceClusterConfigOutputWithContext(context.Context) ClusterClusterConfigGceClusterConfigOutput
}

type ClusterClusterConfigGceClusterConfigArgs struct {
	InternalIpOnly pulumi.BoolPtrInput `pulumi:"internalIpOnly"`
	Metadata pulumi.StringMapInput `pulumi:"metadata"`
	Network pulumi.StringPtrInput `pulumi:"network"`
	ServiceAccount pulumi.StringPtrInput `pulumi:"serviceAccount"`
	ServiceAccountScopes pulumi.StringArrayInput `pulumi:"serviceAccountScopes"`
	Subnetwork pulumi.StringPtrInput `pulumi:"subnetwork"`
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (ClusterClusterConfigGceClusterConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigGceClusterConfig)(nil)).Elem()
}

func (i ClusterClusterConfigGceClusterConfigArgs) ToClusterClusterConfigGceClusterConfigOutput() ClusterClusterConfigGceClusterConfigOutput {
	return i.ToClusterClusterConfigGceClusterConfigOutputWithContext(context.Background())
}

func (i ClusterClusterConfigGceClusterConfigArgs) ToClusterClusterConfigGceClusterConfigOutputWithContext(ctx context.Context) ClusterClusterConfigGceClusterConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigGceClusterConfigOutput)
}

func (i ClusterClusterConfigGceClusterConfigArgs) ToClusterClusterConfigGceClusterConfigPtrOutput() ClusterClusterConfigGceClusterConfigPtrOutput {
	return i.ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(context.Background())
}

func (i ClusterClusterConfigGceClusterConfigArgs) ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigGceClusterConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigGceClusterConfigOutput).ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(ctx)
}

type ClusterClusterConfigGceClusterConfigPtrInput interface {
	pulumi.Input

	ToClusterClusterConfigGceClusterConfigPtrOutput() ClusterClusterConfigGceClusterConfigPtrOutput
	ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(context.Context) ClusterClusterConfigGceClusterConfigPtrOutput
}

type clusterClusterConfigGceClusterConfigPtrType ClusterClusterConfigGceClusterConfigArgs

func ClusterClusterConfigGceClusterConfigPtr(v *ClusterClusterConfigGceClusterConfigArgs) ClusterClusterConfigGceClusterConfigPtrInput {	return (*clusterClusterConfigGceClusterConfigPtrType)(v)
}

func (*clusterClusterConfigGceClusterConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClusterConfigGceClusterConfig)(nil)).Elem()
}

func (i *clusterClusterConfigGceClusterConfigPtrType) ToClusterClusterConfigGceClusterConfigPtrOutput() ClusterClusterConfigGceClusterConfigPtrOutput {
	return i.ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(context.Background())
}

func (i *clusterClusterConfigGceClusterConfigPtrType) ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigGceClusterConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClusterConfigGceClusterConfigPtrOutput)
}

type ClusterClusterConfigGceClusterConfigOutput struct { *pulumi.OutputState }

func (ClusterClusterConfigGceClusterConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClusterConfigGceClusterConfig)(nil)).Elem()
}

func (o ClusterClusterConfigGceClusterConfigOutput) ToClusterClusterConfigGceClusterConfigOutput() ClusterClusterConfigGceClusterConfigOutput {
	return o
}

func (o ClusterClusterConfigGceClusterConfigOutput) ToClusterClusterConfigGceClusterConfigOutputWithContext(ctx context.Context) ClusterClusterConfigGceClusterConfigOutput {
	return o
}

func (o ClusterClusterConfigGceClusterConfigOutput) ToClusterClusterConfigGceClusterConfigPtrOutput() ClusterClusterConfigGceClusterConfigPtrOutput {
	return o.ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(context.Background())
}

func (o ClusterClusterConfigGceClusterConfigOutput) ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigGceClusterConfigPtrOutput {
	return o.ApplyT(func(v ClusterClusterConfigGceClusterConfig) *ClusterClusterConfigGceClusterConfig {
		return &v
	}).(ClusterClusterConfigGceClusterConfigPtrOutput)
}
func (o ClusterClusterConfigGceClusterConfigOutput) InternalIpOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *bool { return v.InternalIpOnly }).(pulumi.BoolPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o ClusterClusterConfigGceClusterConfigOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.Network }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigOutput) ServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.ServiceAccount }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigOutput) ServiceAccountScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) []string { return v.ServiceAccountScopes }).(pulumi.StringArrayOutput)
}

func (o ClusterClusterConfigGceClusterConfigOutput) Subnetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.Subnetwork }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o ClusterClusterConfigGceClusterConfigOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type ClusterClusterConfigGceClusterConfigPtrOutput struct { *pulumi.OutputState}

func (ClusterClusterConfigGceClusterConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClusterConfigGceClusterConfig)(nil)).Elem()
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) ToClusterClusterConfigGceClusterConfigPtrOutput() ClusterClusterConfigGceClusterConfigPtrOutput {
	return o
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) ToClusterClusterConfigGceClusterConfigPtrOutputWithContext(ctx context.Context) ClusterClusterConfigGceClusterConfigPtrOutput {
	return o
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) Elem() ClusterClusterConfigGceClusterConfigOutput {
	return o.ApplyT(func (v *ClusterClusterConfigGceClusterConfig) ClusterClusterConfigGceClusterConfig { return *v }).(ClusterClusterConfigGceClusterConfigOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) InternalIpOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *bool { return v.InternalIpOnly }).(pulumi.BoolPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.Network }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) ServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.ServiceAccount }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) ServiceAccountScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) []string { return v.ServiceAccountScopes }).(pulumi.StringArrayOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) Subnetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.Subnetwork }).(pulumi.StringPtrOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o ClusterClusterConfigGceClusterConfigPtrOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterClusterConfigGceClusterConfig) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterClusterConfigGceClusterConfigOutput{})
	pulumi.RegisterOutputType(ClusterClusterConfigGceClusterConfigPtrOutput{})
}
