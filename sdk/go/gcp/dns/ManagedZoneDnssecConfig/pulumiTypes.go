// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ManagedZoneDnssecConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/dns/ManagedZoneDnssecConfigDefaultKeySpec"
)

type ManagedZoneDnssecConfig struct {
	DefaultKeySpecs []dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpec `pulumi:"defaultKeySpecs"`
	Kind *string `pulumi:"kind"`
	NonExistence *string `pulumi:"nonExistence"`
	State *string `pulumi:"state"`
}

type ManagedZoneDnssecConfigInput interface {
	pulumi.Input

	ToManagedZoneDnssecConfigOutput() ManagedZoneDnssecConfigOutput
	ToManagedZoneDnssecConfigOutputWithContext(context.Context) ManagedZoneDnssecConfigOutput
}

type ManagedZoneDnssecConfigArgs struct {
	DefaultKeySpecs dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpecArrayInput `pulumi:"defaultKeySpecs"`
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	NonExistence pulumi.StringPtrInput `pulumi:"nonExistence"`
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (ManagedZoneDnssecConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedZoneDnssecConfig)(nil)).Elem()
}

func (i ManagedZoneDnssecConfigArgs) ToManagedZoneDnssecConfigOutput() ManagedZoneDnssecConfigOutput {
	return i.ToManagedZoneDnssecConfigOutputWithContext(context.Background())
}

func (i ManagedZoneDnssecConfigArgs) ToManagedZoneDnssecConfigOutputWithContext(ctx context.Context) ManagedZoneDnssecConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedZoneDnssecConfigOutput)
}

func (i ManagedZoneDnssecConfigArgs) ToManagedZoneDnssecConfigPtrOutput() ManagedZoneDnssecConfigPtrOutput {
	return i.ToManagedZoneDnssecConfigPtrOutputWithContext(context.Background())
}

func (i ManagedZoneDnssecConfigArgs) ToManagedZoneDnssecConfigPtrOutputWithContext(ctx context.Context) ManagedZoneDnssecConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedZoneDnssecConfigOutput).ToManagedZoneDnssecConfigPtrOutputWithContext(ctx)
}

type ManagedZoneDnssecConfigPtrInput interface {
	pulumi.Input

	ToManagedZoneDnssecConfigPtrOutput() ManagedZoneDnssecConfigPtrOutput
	ToManagedZoneDnssecConfigPtrOutputWithContext(context.Context) ManagedZoneDnssecConfigPtrOutput
}

type managedZoneDnssecConfigPtrType ManagedZoneDnssecConfigArgs

func ManagedZoneDnssecConfigPtr(v *ManagedZoneDnssecConfigArgs) ManagedZoneDnssecConfigPtrInput {	return (*managedZoneDnssecConfigPtrType)(v)
}

func (*managedZoneDnssecConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedZoneDnssecConfig)(nil)).Elem()
}

func (i *managedZoneDnssecConfigPtrType) ToManagedZoneDnssecConfigPtrOutput() ManagedZoneDnssecConfigPtrOutput {
	return i.ToManagedZoneDnssecConfigPtrOutputWithContext(context.Background())
}

func (i *managedZoneDnssecConfigPtrType) ToManagedZoneDnssecConfigPtrOutputWithContext(ctx context.Context) ManagedZoneDnssecConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedZoneDnssecConfigPtrOutput)
}

type ManagedZoneDnssecConfigOutput struct { *pulumi.OutputState }

func (ManagedZoneDnssecConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedZoneDnssecConfig)(nil)).Elem()
}

func (o ManagedZoneDnssecConfigOutput) ToManagedZoneDnssecConfigOutput() ManagedZoneDnssecConfigOutput {
	return o
}

func (o ManagedZoneDnssecConfigOutput) ToManagedZoneDnssecConfigOutputWithContext(ctx context.Context) ManagedZoneDnssecConfigOutput {
	return o
}

func (o ManagedZoneDnssecConfigOutput) ToManagedZoneDnssecConfigPtrOutput() ManagedZoneDnssecConfigPtrOutput {
	return o.ToManagedZoneDnssecConfigPtrOutputWithContext(context.Background())
}

func (o ManagedZoneDnssecConfigOutput) ToManagedZoneDnssecConfigPtrOutputWithContext(ctx context.Context) ManagedZoneDnssecConfigPtrOutput {
	return o.ApplyT(func(v ManagedZoneDnssecConfig) *ManagedZoneDnssecConfig {
		return &v
	}).(ManagedZoneDnssecConfigPtrOutput)
}
func (o ManagedZoneDnssecConfigOutput) DefaultKeySpecs() dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpecArrayOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) []dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpec { return v.DefaultKeySpecs }).(dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpecArrayOutput)
}

func (o ManagedZoneDnssecConfigOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ManagedZoneDnssecConfigOutput) NonExistence() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) *string { return v.NonExistence }).(pulumi.StringPtrOutput)
}

func (o ManagedZoneDnssecConfigOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ManagedZoneDnssecConfigPtrOutput struct { *pulumi.OutputState}

func (ManagedZoneDnssecConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedZoneDnssecConfig)(nil)).Elem()
}

func (o ManagedZoneDnssecConfigPtrOutput) ToManagedZoneDnssecConfigPtrOutput() ManagedZoneDnssecConfigPtrOutput {
	return o
}

func (o ManagedZoneDnssecConfigPtrOutput) ToManagedZoneDnssecConfigPtrOutputWithContext(ctx context.Context) ManagedZoneDnssecConfigPtrOutput {
	return o
}

func (o ManagedZoneDnssecConfigPtrOutput) Elem() ManagedZoneDnssecConfigOutput {
	return o.ApplyT(func (v *ManagedZoneDnssecConfig) ManagedZoneDnssecConfig { return *v }).(ManagedZoneDnssecConfigOutput)
}

func (o ManagedZoneDnssecConfigPtrOutput) DefaultKeySpecs() dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpecArrayOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) []dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpec { return v.DefaultKeySpecs }).(dnsManagedZoneDnssecConfigDefaultKeySpec.ManagedZoneDnssecConfigDefaultKeySpecArrayOutput)
}

func (o ManagedZoneDnssecConfigPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ManagedZoneDnssecConfigPtrOutput) NonExistence() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) *string { return v.NonExistence }).(pulumi.StringPtrOutput)
}

func (o ManagedZoneDnssecConfigPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ManagedZoneDnssecConfig) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedZoneDnssecConfigOutput{})
	pulumi.RegisterOutputType(ManagedZoneDnssecConfigPtrOutput{})
}
