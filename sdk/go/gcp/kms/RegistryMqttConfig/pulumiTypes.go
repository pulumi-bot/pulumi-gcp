// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegistryMqttConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegistryMqttConfig struct {
	MqttEnabledState string `pulumi:"mqttEnabledState"`
}

type RegistryMqttConfigInput interface {
	pulumi.Input

	ToRegistryMqttConfigOutput() RegistryMqttConfigOutput
	ToRegistryMqttConfigOutputWithContext(context.Context) RegistryMqttConfigOutput
}

type RegistryMqttConfigArgs struct {
	MqttEnabledState pulumi.StringInput `pulumi:"mqttEnabledState"`
}

func (RegistryMqttConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryMqttConfig)(nil)).Elem()
}

func (i RegistryMqttConfigArgs) ToRegistryMqttConfigOutput() RegistryMqttConfigOutput {
	return i.ToRegistryMqttConfigOutputWithContext(context.Background())
}

func (i RegistryMqttConfigArgs) ToRegistryMqttConfigOutputWithContext(ctx context.Context) RegistryMqttConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryMqttConfigOutput)
}

func (i RegistryMqttConfigArgs) ToRegistryMqttConfigPtrOutput() RegistryMqttConfigPtrOutput {
	return i.ToRegistryMqttConfigPtrOutputWithContext(context.Background())
}

func (i RegistryMqttConfigArgs) ToRegistryMqttConfigPtrOutputWithContext(ctx context.Context) RegistryMqttConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryMqttConfigOutput).ToRegistryMqttConfigPtrOutputWithContext(ctx)
}

type RegistryMqttConfigPtrInput interface {
	pulumi.Input

	ToRegistryMqttConfigPtrOutput() RegistryMqttConfigPtrOutput
	ToRegistryMqttConfigPtrOutputWithContext(context.Context) RegistryMqttConfigPtrOutput
}

type registryMqttConfigPtrType RegistryMqttConfigArgs

func RegistryMqttConfigPtr(v *RegistryMqttConfigArgs) RegistryMqttConfigPtrInput {	return (*registryMqttConfigPtrType)(v)
}

func (*registryMqttConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryMqttConfig)(nil)).Elem()
}

func (i *registryMqttConfigPtrType) ToRegistryMqttConfigPtrOutput() RegistryMqttConfigPtrOutput {
	return i.ToRegistryMqttConfigPtrOutputWithContext(context.Background())
}

func (i *registryMqttConfigPtrType) ToRegistryMqttConfigPtrOutputWithContext(ctx context.Context) RegistryMqttConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryMqttConfigPtrOutput)
}

type RegistryMqttConfigOutput struct { *pulumi.OutputState }

func (RegistryMqttConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryMqttConfig)(nil)).Elem()
}

func (o RegistryMqttConfigOutput) ToRegistryMqttConfigOutput() RegistryMqttConfigOutput {
	return o
}

func (o RegistryMqttConfigOutput) ToRegistryMqttConfigOutputWithContext(ctx context.Context) RegistryMqttConfigOutput {
	return o
}

func (o RegistryMqttConfigOutput) ToRegistryMqttConfigPtrOutput() RegistryMqttConfigPtrOutput {
	return o.ToRegistryMqttConfigPtrOutputWithContext(context.Background())
}

func (o RegistryMqttConfigOutput) ToRegistryMqttConfigPtrOutputWithContext(ctx context.Context) RegistryMqttConfigPtrOutput {
	return o.ApplyT(func(v RegistryMqttConfig) *RegistryMqttConfig {
		return &v
	}).(RegistryMqttConfigPtrOutput)
}
func (o RegistryMqttConfigOutput) MqttEnabledState() pulumi.StringOutput {
	return o.ApplyT(func (v RegistryMqttConfig) string { return v.MqttEnabledState }).(pulumi.StringOutput)
}

type RegistryMqttConfigPtrOutput struct { *pulumi.OutputState}

func (RegistryMqttConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryMqttConfig)(nil)).Elem()
}

func (o RegistryMqttConfigPtrOutput) ToRegistryMqttConfigPtrOutput() RegistryMqttConfigPtrOutput {
	return o
}

func (o RegistryMqttConfigPtrOutput) ToRegistryMqttConfigPtrOutputWithContext(ctx context.Context) RegistryMqttConfigPtrOutput {
	return o
}

func (o RegistryMqttConfigPtrOutput) Elem() RegistryMqttConfigOutput {
	return o.ApplyT(func (v *RegistryMqttConfig) RegistryMqttConfig { return *v }).(RegistryMqttConfigOutput)
}

func (o RegistryMqttConfigPtrOutput) MqttEnabledState() pulumi.StringOutput {
	return o.ApplyT(func (v RegistryMqttConfig) string { return v.MqttEnabledState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryMqttConfigOutput{})
	pulumi.RegisterOutputType(RegistryMqttConfigPtrOutput{})
}
