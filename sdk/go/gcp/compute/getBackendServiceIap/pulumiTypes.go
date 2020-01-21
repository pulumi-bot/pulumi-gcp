// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getBackendServiceIap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetBackendServiceIap struct {
	Oauth2ClientId string `pulumi:"oauth2ClientId"`
	Oauth2ClientSecret string `pulumi:"oauth2ClientSecret"`
	Oauth2ClientSecretSha256 string `pulumi:"oauth2ClientSecretSha256"`
}

type GetBackendServiceIapInput interface {
	pulumi.Input

	ToGetBackendServiceIapOutput() GetBackendServiceIapOutput
	ToGetBackendServiceIapOutputWithContext(context.Context) GetBackendServiceIapOutput
}

type GetBackendServiceIapArgs struct {
	Oauth2ClientId pulumi.StringInput `pulumi:"oauth2ClientId"`
	Oauth2ClientSecret pulumi.StringInput `pulumi:"oauth2ClientSecret"`
	Oauth2ClientSecretSha256 pulumi.StringInput `pulumi:"oauth2ClientSecretSha256"`
}

func (GetBackendServiceIapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendServiceIap)(nil)).Elem()
}

func (i GetBackendServiceIapArgs) ToGetBackendServiceIapOutput() GetBackendServiceIapOutput {
	return i.ToGetBackendServiceIapOutputWithContext(context.Background())
}

func (i GetBackendServiceIapArgs) ToGetBackendServiceIapOutputWithContext(ctx context.Context) GetBackendServiceIapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackendServiceIapOutput)
}

type GetBackendServiceIapArrayInput interface {
	pulumi.Input

	ToGetBackendServiceIapArrayOutput() GetBackendServiceIapArrayOutput
	ToGetBackendServiceIapArrayOutputWithContext(context.Context) GetBackendServiceIapArrayOutput
}

type GetBackendServiceIapArray []GetBackendServiceIapInput

func (GetBackendServiceIapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackendServiceIap)(nil)).Elem()
}

func (i GetBackendServiceIapArray) ToGetBackendServiceIapArrayOutput() GetBackendServiceIapArrayOutput {
	return i.ToGetBackendServiceIapArrayOutputWithContext(context.Background())
}

func (i GetBackendServiceIapArray) ToGetBackendServiceIapArrayOutputWithContext(ctx context.Context) GetBackendServiceIapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackendServiceIapArrayOutput)
}

type GetBackendServiceIapOutput struct { *pulumi.OutputState }

func (GetBackendServiceIapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendServiceIap)(nil)).Elem()
}

func (o GetBackendServiceIapOutput) ToGetBackendServiceIapOutput() GetBackendServiceIapOutput {
	return o
}

func (o GetBackendServiceIapOutput) ToGetBackendServiceIapOutputWithContext(ctx context.Context) GetBackendServiceIapOutput {
	return o
}

func (o GetBackendServiceIapOutput) Oauth2ClientId() pulumi.StringOutput {
	return o.ApplyT(func (v GetBackendServiceIap) string { return v.Oauth2ClientId }).(pulumi.StringOutput)
}

func (o GetBackendServiceIapOutput) Oauth2ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func (v GetBackendServiceIap) string { return v.Oauth2ClientSecret }).(pulumi.StringOutput)
}

func (o GetBackendServiceIapOutput) Oauth2ClientSecretSha256() pulumi.StringOutput {
	return o.ApplyT(func (v GetBackendServiceIap) string { return v.Oauth2ClientSecretSha256 }).(pulumi.StringOutput)
}

type GetBackendServiceIapArrayOutput struct { *pulumi.OutputState}

func (GetBackendServiceIapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackendServiceIap)(nil)).Elem()
}

func (o GetBackendServiceIapArrayOutput) ToGetBackendServiceIapArrayOutput() GetBackendServiceIapArrayOutput {
	return o
}

func (o GetBackendServiceIapArrayOutput) ToGetBackendServiceIapArrayOutputWithContext(ctx context.Context) GetBackendServiceIapArrayOutput {
	return o
}

func (o GetBackendServiceIapArrayOutput) Index(i pulumi.IntInput) GetBackendServiceIapOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetBackendServiceIap {
		return vs[0].([]GetBackendServiceIap)[vs[1].(int)]
	}).(GetBackendServiceIapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackendServiceIapOutput{})
	pulumi.RegisterOutputType(GetBackendServiceIapArrayOutput{})
}
