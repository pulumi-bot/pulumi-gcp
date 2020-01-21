// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getClusterNodePoolUpgradeSetting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetClusterNodePoolUpgradeSetting struct {
	MaxSurge int `pulumi:"maxSurge"`
	MaxUnavailable int `pulumi:"maxUnavailable"`
}

type GetClusterNodePoolUpgradeSettingInput interface {
	pulumi.Input

	ToGetClusterNodePoolUpgradeSettingOutput() GetClusterNodePoolUpgradeSettingOutput
	ToGetClusterNodePoolUpgradeSettingOutputWithContext(context.Context) GetClusterNodePoolUpgradeSettingOutput
}

type GetClusterNodePoolUpgradeSettingArgs struct {
	MaxSurge pulumi.IntInput `pulumi:"maxSurge"`
	MaxUnavailable pulumi.IntInput `pulumi:"maxUnavailable"`
}

func (GetClusterNodePoolUpgradeSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterNodePoolUpgradeSetting)(nil)).Elem()
}

func (i GetClusterNodePoolUpgradeSettingArgs) ToGetClusterNodePoolUpgradeSettingOutput() GetClusterNodePoolUpgradeSettingOutput {
	return i.ToGetClusterNodePoolUpgradeSettingOutputWithContext(context.Background())
}

func (i GetClusterNodePoolUpgradeSettingArgs) ToGetClusterNodePoolUpgradeSettingOutputWithContext(ctx context.Context) GetClusterNodePoolUpgradeSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterNodePoolUpgradeSettingOutput)
}

type GetClusterNodePoolUpgradeSettingArrayInput interface {
	pulumi.Input

	ToGetClusterNodePoolUpgradeSettingArrayOutput() GetClusterNodePoolUpgradeSettingArrayOutput
	ToGetClusterNodePoolUpgradeSettingArrayOutputWithContext(context.Context) GetClusterNodePoolUpgradeSettingArrayOutput
}

type GetClusterNodePoolUpgradeSettingArray []GetClusterNodePoolUpgradeSettingInput

func (GetClusterNodePoolUpgradeSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterNodePoolUpgradeSetting)(nil)).Elem()
}

func (i GetClusterNodePoolUpgradeSettingArray) ToGetClusterNodePoolUpgradeSettingArrayOutput() GetClusterNodePoolUpgradeSettingArrayOutput {
	return i.ToGetClusterNodePoolUpgradeSettingArrayOutputWithContext(context.Background())
}

func (i GetClusterNodePoolUpgradeSettingArray) ToGetClusterNodePoolUpgradeSettingArrayOutputWithContext(ctx context.Context) GetClusterNodePoolUpgradeSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterNodePoolUpgradeSettingArrayOutput)
}

type GetClusterNodePoolUpgradeSettingOutput struct { *pulumi.OutputState }

func (GetClusterNodePoolUpgradeSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterNodePoolUpgradeSetting)(nil)).Elem()
}

func (o GetClusterNodePoolUpgradeSettingOutput) ToGetClusterNodePoolUpgradeSettingOutput() GetClusterNodePoolUpgradeSettingOutput {
	return o
}

func (o GetClusterNodePoolUpgradeSettingOutput) ToGetClusterNodePoolUpgradeSettingOutputWithContext(ctx context.Context) GetClusterNodePoolUpgradeSettingOutput {
	return o
}

func (o GetClusterNodePoolUpgradeSettingOutput) MaxSurge() pulumi.IntOutput {
	return o.ApplyT(func (v GetClusterNodePoolUpgradeSetting) int { return v.MaxSurge }).(pulumi.IntOutput)
}

func (o GetClusterNodePoolUpgradeSettingOutput) MaxUnavailable() pulumi.IntOutput {
	return o.ApplyT(func (v GetClusterNodePoolUpgradeSetting) int { return v.MaxUnavailable }).(pulumi.IntOutput)
}

type GetClusterNodePoolUpgradeSettingArrayOutput struct { *pulumi.OutputState}

func (GetClusterNodePoolUpgradeSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterNodePoolUpgradeSetting)(nil)).Elem()
}

func (o GetClusterNodePoolUpgradeSettingArrayOutput) ToGetClusterNodePoolUpgradeSettingArrayOutput() GetClusterNodePoolUpgradeSettingArrayOutput {
	return o
}

func (o GetClusterNodePoolUpgradeSettingArrayOutput) ToGetClusterNodePoolUpgradeSettingArrayOutputWithContext(ctx context.Context) GetClusterNodePoolUpgradeSettingArrayOutput {
	return o
}

func (o GetClusterNodePoolUpgradeSettingArrayOutput) Index(i pulumi.IntInput) GetClusterNodePoolUpgradeSettingOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetClusterNodePoolUpgradeSetting {
		return vs[0].([]GetClusterNodePoolUpgradeSetting)[vs[1].(int)]
	}).(GetClusterNodePoolUpgradeSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterNodePoolUpgradeSettingOutput{})
	pulumi.RegisterOutputType(GetClusterNodePoolUpgradeSettingArrayOutput{})
}
