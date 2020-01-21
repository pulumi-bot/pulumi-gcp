// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getClusterAddonsConfigNetworkPolicyConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetClusterAddonsConfigNetworkPolicyConfig struct {
	Disabled bool `pulumi:"disabled"`
}

type GetClusterAddonsConfigNetworkPolicyConfigInput interface {
	pulumi.Input

	ToGetClusterAddonsConfigNetworkPolicyConfigOutput() GetClusterAddonsConfigNetworkPolicyConfigOutput
	ToGetClusterAddonsConfigNetworkPolicyConfigOutputWithContext(context.Context) GetClusterAddonsConfigNetworkPolicyConfigOutput
}

type GetClusterAddonsConfigNetworkPolicyConfigArgs struct {
	Disabled pulumi.BoolInput `pulumi:"disabled"`
}

func (GetClusterAddonsConfigNetworkPolicyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterAddonsConfigNetworkPolicyConfig)(nil)).Elem()
}

func (i GetClusterAddonsConfigNetworkPolicyConfigArgs) ToGetClusterAddonsConfigNetworkPolicyConfigOutput() GetClusterAddonsConfigNetworkPolicyConfigOutput {
	return i.ToGetClusterAddonsConfigNetworkPolicyConfigOutputWithContext(context.Background())
}

func (i GetClusterAddonsConfigNetworkPolicyConfigArgs) ToGetClusterAddonsConfigNetworkPolicyConfigOutputWithContext(ctx context.Context) GetClusterAddonsConfigNetworkPolicyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterAddonsConfigNetworkPolicyConfigOutput)
}

type GetClusterAddonsConfigNetworkPolicyConfigArrayInput interface {
	pulumi.Input

	ToGetClusterAddonsConfigNetworkPolicyConfigArrayOutput() GetClusterAddonsConfigNetworkPolicyConfigArrayOutput
	ToGetClusterAddonsConfigNetworkPolicyConfigArrayOutputWithContext(context.Context) GetClusterAddonsConfigNetworkPolicyConfigArrayOutput
}

type GetClusterAddonsConfigNetworkPolicyConfigArray []GetClusterAddonsConfigNetworkPolicyConfigInput

func (GetClusterAddonsConfigNetworkPolicyConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterAddonsConfigNetworkPolicyConfig)(nil)).Elem()
}

func (i GetClusterAddonsConfigNetworkPolicyConfigArray) ToGetClusterAddonsConfigNetworkPolicyConfigArrayOutput() GetClusterAddonsConfigNetworkPolicyConfigArrayOutput {
	return i.ToGetClusterAddonsConfigNetworkPolicyConfigArrayOutputWithContext(context.Background())
}

func (i GetClusterAddonsConfigNetworkPolicyConfigArray) ToGetClusterAddonsConfigNetworkPolicyConfigArrayOutputWithContext(ctx context.Context) GetClusterAddonsConfigNetworkPolicyConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterAddonsConfigNetworkPolicyConfigArrayOutput)
}

type GetClusterAddonsConfigNetworkPolicyConfigOutput struct { *pulumi.OutputState }

func (GetClusterAddonsConfigNetworkPolicyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterAddonsConfigNetworkPolicyConfig)(nil)).Elem()
}

func (o GetClusterAddonsConfigNetworkPolicyConfigOutput) ToGetClusterAddonsConfigNetworkPolicyConfigOutput() GetClusterAddonsConfigNetworkPolicyConfigOutput {
	return o
}

func (o GetClusterAddonsConfigNetworkPolicyConfigOutput) ToGetClusterAddonsConfigNetworkPolicyConfigOutputWithContext(ctx context.Context) GetClusterAddonsConfigNetworkPolicyConfigOutput {
	return o
}

func (o GetClusterAddonsConfigNetworkPolicyConfigOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func (v GetClusterAddonsConfigNetworkPolicyConfig) bool { return v.Disabled }).(pulumi.BoolOutput)
}

type GetClusterAddonsConfigNetworkPolicyConfigArrayOutput struct { *pulumi.OutputState}

func (GetClusterAddonsConfigNetworkPolicyConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterAddonsConfigNetworkPolicyConfig)(nil)).Elem()
}

func (o GetClusterAddonsConfigNetworkPolicyConfigArrayOutput) ToGetClusterAddonsConfigNetworkPolicyConfigArrayOutput() GetClusterAddonsConfigNetworkPolicyConfigArrayOutput {
	return o
}

func (o GetClusterAddonsConfigNetworkPolicyConfigArrayOutput) ToGetClusterAddonsConfigNetworkPolicyConfigArrayOutputWithContext(ctx context.Context) GetClusterAddonsConfigNetworkPolicyConfigArrayOutput {
	return o
}

func (o GetClusterAddonsConfigNetworkPolicyConfigArrayOutput) Index(i pulumi.IntInput) GetClusterAddonsConfigNetworkPolicyConfigOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetClusterAddonsConfigNetworkPolicyConfig {
		return vs[0].([]GetClusterAddonsConfigNetworkPolicyConfig)[vs[1].(int)]
	}).(GetClusterAddonsConfigNetworkPolicyConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterAddonsConfigNetworkPolicyConfigOutput{})
	pulumi.RegisterOutputType(GetClusterAddonsConfigNetworkPolicyConfigArrayOutput{})
}
