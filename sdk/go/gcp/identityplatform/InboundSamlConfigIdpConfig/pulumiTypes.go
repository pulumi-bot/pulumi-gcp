// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package InboundSamlConfigIdpConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/identityplatform/InboundSamlConfigIdpConfigIdpCertificate"
)

type InboundSamlConfigIdpConfig struct {
	IdpCertificates []identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificate `pulumi:"idpCertificates"`
	IdpEntityId string `pulumi:"idpEntityId"`
	SignRequest *bool `pulumi:"signRequest"`
	SsoUrl string `pulumi:"ssoUrl"`
}

type InboundSamlConfigIdpConfigInput interface {
	pulumi.Input

	ToInboundSamlConfigIdpConfigOutput() InboundSamlConfigIdpConfigOutput
	ToInboundSamlConfigIdpConfigOutputWithContext(context.Context) InboundSamlConfigIdpConfigOutput
}

type InboundSamlConfigIdpConfigArgs struct {
	IdpCertificates identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificateArrayInput `pulumi:"idpCertificates"`
	IdpEntityId pulumi.StringInput `pulumi:"idpEntityId"`
	SignRequest pulumi.BoolPtrInput `pulumi:"signRequest"`
	SsoUrl pulumi.StringInput `pulumi:"ssoUrl"`
}

func (InboundSamlConfigIdpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundSamlConfigIdpConfig)(nil)).Elem()
}

func (i InboundSamlConfigIdpConfigArgs) ToInboundSamlConfigIdpConfigOutput() InboundSamlConfigIdpConfigOutput {
	return i.ToInboundSamlConfigIdpConfigOutputWithContext(context.Background())
}

func (i InboundSamlConfigIdpConfigArgs) ToInboundSamlConfigIdpConfigOutputWithContext(ctx context.Context) InboundSamlConfigIdpConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigIdpConfigOutput)
}

func (i InboundSamlConfigIdpConfigArgs) ToInboundSamlConfigIdpConfigPtrOutput() InboundSamlConfigIdpConfigPtrOutput {
	return i.ToInboundSamlConfigIdpConfigPtrOutputWithContext(context.Background())
}

func (i InboundSamlConfigIdpConfigArgs) ToInboundSamlConfigIdpConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigIdpConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigIdpConfigOutput).ToInboundSamlConfigIdpConfigPtrOutputWithContext(ctx)
}

type InboundSamlConfigIdpConfigPtrInput interface {
	pulumi.Input

	ToInboundSamlConfigIdpConfigPtrOutput() InboundSamlConfigIdpConfigPtrOutput
	ToInboundSamlConfigIdpConfigPtrOutputWithContext(context.Context) InboundSamlConfigIdpConfigPtrOutput
}

type inboundSamlConfigIdpConfigPtrType InboundSamlConfigIdpConfigArgs

func InboundSamlConfigIdpConfigPtr(v *InboundSamlConfigIdpConfigArgs) InboundSamlConfigIdpConfigPtrInput {	return (*inboundSamlConfigIdpConfigPtrType)(v)
}

func (*inboundSamlConfigIdpConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundSamlConfigIdpConfig)(nil)).Elem()
}

func (i *inboundSamlConfigIdpConfigPtrType) ToInboundSamlConfigIdpConfigPtrOutput() InboundSamlConfigIdpConfigPtrOutput {
	return i.ToInboundSamlConfigIdpConfigPtrOutputWithContext(context.Background())
}

func (i *inboundSamlConfigIdpConfigPtrType) ToInboundSamlConfigIdpConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigIdpConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlConfigIdpConfigPtrOutput)
}

type InboundSamlConfigIdpConfigOutput struct { *pulumi.OutputState }

func (InboundSamlConfigIdpConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundSamlConfigIdpConfig)(nil)).Elem()
}

func (o InboundSamlConfigIdpConfigOutput) ToInboundSamlConfigIdpConfigOutput() InboundSamlConfigIdpConfigOutput {
	return o
}

func (o InboundSamlConfigIdpConfigOutput) ToInboundSamlConfigIdpConfigOutputWithContext(ctx context.Context) InboundSamlConfigIdpConfigOutput {
	return o
}

func (o InboundSamlConfigIdpConfigOutput) ToInboundSamlConfigIdpConfigPtrOutput() InboundSamlConfigIdpConfigPtrOutput {
	return o.ToInboundSamlConfigIdpConfigPtrOutputWithContext(context.Background())
}

func (o InboundSamlConfigIdpConfigOutput) ToInboundSamlConfigIdpConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigIdpConfigPtrOutput {
	return o.ApplyT(func(v InboundSamlConfigIdpConfig) *InboundSamlConfigIdpConfig {
		return &v
	}).(InboundSamlConfigIdpConfigPtrOutput)
}
func (o InboundSamlConfigIdpConfigOutput) IdpCertificates() identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificateArrayOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) []identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificate { return v.IdpCertificates }).(identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificateArrayOutput)
}

func (o InboundSamlConfigIdpConfigOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) string { return v.IdpEntityId }).(pulumi.StringOutput)
}

func (o InboundSamlConfigIdpConfigOutput) SignRequest() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) *bool { return v.SignRequest }).(pulumi.BoolPtrOutput)
}

func (o InboundSamlConfigIdpConfigOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) string { return v.SsoUrl }).(pulumi.StringOutput)
}

type InboundSamlConfigIdpConfigPtrOutput struct { *pulumi.OutputState}

func (InboundSamlConfigIdpConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundSamlConfigIdpConfig)(nil)).Elem()
}

func (o InboundSamlConfigIdpConfigPtrOutput) ToInboundSamlConfigIdpConfigPtrOutput() InboundSamlConfigIdpConfigPtrOutput {
	return o
}

func (o InboundSamlConfigIdpConfigPtrOutput) ToInboundSamlConfigIdpConfigPtrOutputWithContext(ctx context.Context) InboundSamlConfigIdpConfigPtrOutput {
	return o
}

func (o InboundSamlConfigIdpConfigPtrOutput) Elem() InboundSamlConfigIdpConfigOutput {
	return o.ApplyT(func (v *InboundSamlConfigIdpConfig) InboundSamlConfigIdpConfig { return *v }).(InboundSamlConfigIdpConfigOutput)
}

func (o InboundSamlConfigIdpConfigPtrOutput) IdpCertificates() identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificateArrayOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) []identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificate { return v.IdpCertificates }).(identityplatformInboundSamlConfigIdpConfigIdpCertificate.InboundSamlConfigIdpConfigIdpCertificateArrayOutput)
}

func (o InboundSamlConfigIdpConfigPtrOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) string { return v.IdpEntityId }).(pulumi.StringOutput)
}

func (o InboundSamlConfigIdpConfigPtrOutput) SignRequest() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) *bool { return v.SignRequest }).(pulumi.BoolPtrOutput)
}

func (o InboundSamlConfigIdpConfigPtrOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func (v InboundSamlConfigIdpConfig) string { return v.SsoUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InboundSamlConfigIdpConfigOutput{})
	pulumi.RegisterOutputType(InboundSamlConfigIdpConfigPtrOutput{})
}
