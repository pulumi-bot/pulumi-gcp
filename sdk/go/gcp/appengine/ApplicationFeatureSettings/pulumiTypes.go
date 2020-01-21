// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ApplicationFeatureSettings

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ApplicationFeatureSettings struct {
	// Set to false to use the legacy health check instead of the readiness
	// and liveness checks.
	SplitHealthChecks bool `pulumi:"splitHealthChecks"`
}

type ApplicationFeatureSettingsInput interface {
	pulumi.Input

	ToApplicationFeatureSettingsOutput() ApplicationFeatureSettingsOutput
	ToApplicationFeatureSettingsOutputWithContext(context.Context) ApplicationFeatureSettingsOutput
}

type ApplicationFeatureSettingsArgs struct {
	// Set to false to use the legacy health check instead of the readiness
	// and liveness checks.
	SplitHealthChecks pulumi.BoolInput `pulumi:"splitHealthChecks"`
}

func (ApplicationFeatureSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationFeatureSettings)(nil)).Elem()
}

func (i ApplicationFeatureSettingsArgs) ToApplicationFeatureSettingsOutput() ApplicationFeatureSettingsOutput {
	return i.ToApplicationFeatureSettingsOutputWithContext(context.Background())
}

func (i ApplicationFeatureSettingsArgs) ToApplicationFeatureSettingsOutputWithContext(ctx context.Context) ApplicationFeatureSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFeatureSettingsOutput)
}

func (i ApplicationFeatureSettingsArgs) ToApplicationFeatureSettingsPtrOutput() ApplicationFeatureSettingsPtrOutput {
	return i.ToApplicationFeatureSettingsPtrOutputWithContext(context.Background())
}

func (i ApplicationFeatureSettingsArgs) ToApplicationFeatureSettingsPtrOutputWithContext(ctx context.Context) ApplicationFeatureSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFeatureSettingsOutput).ToApplicationFeatureSettingsPtrOutputWithContext(ctx)
}

type ApplicationFeatureSettingsPtrInput interface {
	pulumi.Input

	ToApplicationFeatureSettingsPtrOutput() ApplicationFeatureSettingsPtrOutput
	ToApplicationFeatureSettingsPtrOutputWithContext(context.Context) ApplicationFeatureSettingsPtrOutput
}

type applicationFeatureSettingsPtrType ApplicationFeatureSettingsArgs

func ApplicationFeatureSettingsPtr(v *ApplicationFeatureSettingsArgs) ApplicationFeatureSettingsPtrInput {	return (*applicationFeatureSettingsPtrType)(v)
}

func (*applicationFeatureSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFeatureSettings)(nil)).Elem()
}

func (i *applicationFeatureSettingsPtrType) ToApplicationFeatureSettingsPtrOutput() ApplicationFeatureSettingsPtrOutput {
	return i.ToApplicationFeatureSettingsPtrOutputWithContext(context.Background())
}

func (i *applicationFeatureSettingsPtrType) ToApplicationFeatureSettingsPtrOutputWithContext(ctx context.Context) ApplicationFeatureSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFeatureSettingsPtrOutput)
}

type ApplicationFeatureSettingsOutput struct { *pulumi.OutputState }

func (ApplicationFeatureSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationFeatureSettings)(nil)).Elem()
}

func (o ApplicationFeatureSettingsOutput) ToApplicationFeatureSettingsOutput() ApplicationFeatureSettingsOutput {
	return o
}

func (o ApplicationFeatureSettingsOutput) ToApplicationFeatureSettingsOutputWithContext(ctx context.Context) ApplicationFeatureSettingsOutput {
	return o
}

func (o ApplicationFeatureSettingsOutput) ToApplicationFeatureSettingsPtrOutput() ApplicationFeatureSettingsPtrOutput {
	return o.ToApplicationFeatureSettingsPtrOutputWithContext(context.Background())
}

func (o ApplicationFeatureSettingsOutput) ToApplicationFeatureSettingsPtrOutputWithContext(ctx context.Context) ApplicationFeatureSettingsPtrOutput {
	return o.ApplyT(func(v ApplicationFeatureSettings) *ApplicationFeatureSettings {
		return &v
	}).(ApplicationFeatureSettingsPtrOutput)
}
// Set to false to use the legacy health check instead of the readiness
// and liveness checks.
func (o ApplicationFeatureSettingsOutput) SplitHealthChecks() pulumi.BoolOutput {
	return o.ApplyT(func (v ApplicationFeatureSettings) bool { return v.SplitHealthChecks }).(pulumi.BoolOutput)
}

type ApplicationFeatureSettingsPtrOutput struct { *pulumi.OutputState}

func (ApplicationFeatureSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFeatureSettings)(nil)).Elem()
}

func (o ApplicationFeatureSettingsPtrOutput) ToApplicationFeatureSettingsPtrOutput() ApplicationFeatureSettingsPtrOutput {
	return o
}

func (o ApplicationFeatureSettingsPtrOutput) ToApplicationFeatureSettingsPtrOutputWithContext(ctx context.Context) ApplicationFeatureSettingsPtrOutput {
	return o
}

func (o ApplicationFeatureSettingsPtrOutput) Elem() ApplicationFeatureSettingsOutput {
	return o.ApplyT(func (v *ApplicationFeatureSettings) ApplicationFeatureSettings { return *v }).(ApplicationFeatureSettingsOutput)
}

// Set to false to use the legacy health check instead of the readiness
// and liveness checks.
func (o ApplicationFeatureSettingsPtrOutput) SplitHealthChecks() pulumi.BoolOutput {
	return o.ApplyT(func (v ApplicationFeatureSettings) bool { return v.SplitHealthChecks }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationFeatureSettingsOutput{})
	pulumi.RegisterOutputType(ApplicationFeatureSettingsPtrOutput{})
}
