// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package InstanceTemplateServiceAccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InstanceTemplateServiceAccount struct {
	Email *string `pulumi:"email"`
	Scopes []string `pulumi:"scopes"`
}

type InstanceTemplateServiceAccountInput interface {
	pulumi.Input

	ToInstanceTemplateServiceAccountOutput() InstanceTemplateServiceAccountOutput
	ToInstanceTemplateServiceAccountOutputWithContext(context.Context) InstanceTemplateServiceAccountOutput
}

type InstanceTemplateServiceAccountArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Scopes pulumi.StringArrayInput `pulumi:"scopes"`
}

func (InstanceTemplateServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplateServiceAccount)(nil)).Elem()
}

func (i InstanceTemplateServiceAccountArgs) ToInstanceTemplateServiceAccountOutput() InstanceTemplateServiceAccountOutput {
	return i.ToInstanceTemplateServiceAccountOutputWithContext(context.Background())
}

func (i InstanceTemplateServiceAccountArgs) ToInstanceTemplateServiceAccountOutputWithContext(ctx context.Context) InstanceTemplateServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTemplateServiceAccountOutput)
}

func (i InstanceTemplateServiceAccountArgs) ToInstanceTemplateServiceAccountPtrOutput() InstanceTemplateServiceAccountPtrOutput {
	return i.ToInstanceTemplateServiceAccountPtrOutputWithContext(context.Background())
}

func (i InstanceTemplateServiceAccountArgs) ToInstanceTemplateServiceAccountPtrOutputWithContext(ctx context.Context) InstanceTemplateServiceAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTemplateServiceAccountOutput).ToInstanceTemplateServiceAccountPtrOutputWithContext(ctx)
}

type InstanceTemplateServiceAccountPtrInput interface {
	pulumi.Input

	ToInstanceTemplateServiceAccountPtrOutput() InstanceTemplateServiceAccountPtrOutput
	ToInstanceTemplateServiceAccountPtrOutputWithContext(context.Context) InstanceTemplateServiceAccountPtrOutput
}

type instanceTemplateServiceAccountPtrType InstanceTemplateServiceAccountArgs

func InstanceTemplateServiceAccountPtr(v *InstanceTemplateServiceAccountArgs) InstanceTemplateServiceAccountPtrInput {	return (*instanceTemplateServiceAccountPtrType)(v)
}

func (*instanceTemplateServiceAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTemplateServiceAccount)(nil)).Elem()
}

func (i *instanceTemplateServiceAccountPtrType) ToInstanceTemplateServiceAccountPtrOutput() InstanceTemplateServiceAccountPtrOutput {
	return i.ToInstanceTemplateServiceAccountPtrOutputWithContext(context.Background())
}

func (i *instanceTemplateServiceAccountPtrType) ToInstanceTemplateServiceAccountPtrOutputWithContext(ctx context.Context) InstanceTemplateServiceAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTemplateServiceAccountPtrOutput)
}

type InstanceTemplateServiceAccountOutput struct { *pulumi.OutputState }

func (InstanceTemplateServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplateServiceAccount)(nil)).Elem()
}

func (o InstanceTemplateServiceAccountOutput) ToInstanceTemplateServiceAccountOutput() InstanceTemplateServiceAccountOutput {
	return o
}

func (o InstanceTemplateServiceAccountOutput) ToInstanceTemplateServiceAccountOutputWithContext(ctx context.Context) InstanceTemplateServiceAccountOutput {
	return o
}

func (o InstanceTemplateServiceAccountOutput) ToInstanceTemplateServiceAccountPtrOutput() InstanceTemplateServiceAccountPtrOutput {
	return o.ToInstanceTemplateServiceAccountPtrOutputWithContext(context.Background())
}

func (o InstanceTemplateServiceAccountOutput) ToInstanceTemplateServiceAccountPtrOutputWithContext(ctx context.Context) InstanceTemplateServiceAccountPtrOutput {
	return o.ApplyT(func(v InstanceTemplateServiceAccount) *InstanceTemplateServiceAccount {
		return &v
	}).(InstanceTemplateServiceAccountPtrOutput)
}
func (o InstanceTemplateServiceAccountOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func (v InstanceTemplateServiceAccount) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o InstanceTemplateServiceAccountOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v InstanceTemplateServiceAccount) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type InstanceTemplateServiceAccountPtrOutput struct { *pulumi.OutputState}

func (InstanceTemplateServiceAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTemplateServiceAccount)(nil)).Elem()
}

func (o InstanceTemplateServiceAccountPtrOutput) ToInstanceTemplateServiceAccountPtrOutput() InstanceTemplateServiceAccountPtrOutput {
	return o
}

func (o InstanceTemplateServiceAccountPtrOutput) ToInstanceTemplateServiceAccountPtrOutputWithContext(ctx context.Context) InstanceTemplateServiceAccountPtrOutput {
	return o
}

func (o InstanceTemplateServiceAccountPtrOutput) Elem() InstanceTemplateServiceAccountOutput {
	return o.ApplyT(func (v *InstanceTemplateServiceAccount) InstanceTemplateServiceAccount { return *v }).(InstanceTemplateServiceAccountOutput)
}

func (o InstanceTemplateServiceAccountPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func (v InstanceTemplateServiceAccount) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o InstanceTemplateServiceAccountPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v InstanceTemplateServiceAccount) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceTemplateServiceAccountOutput{})
	pulumi.RegisterOutputType(InstanceTemplateServiceAccountPtrOutput{})
}
