// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getInstanceServiceAccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetInstanceServiceAccount struct {
	// The service account e-mail address.
	Email string `pulumi:"email"`
	// A list of service scopes.
	Scopes []string `pulumi:"scopes"`
}

type GetInstanceServiceAccountInput interface {
	pulumi.Input

	ToGetInstanceServiceAccountOutput() GetInstanceServiceAccountOutput
	ToGetInstanceServiceAccountOutputWithContext(context.Context) GetInstanceServiceAccountOutput
}

type GetInstanceServiceAccountArgs struct {
	// The service account e-mail address.
	Email pulumi.StringInput `pulumi:"email"`
	// A list of service scopes.
	Scopes pulumi.StringArrayInput `pulumi:"scopes"`
}

func (GetInstanceServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceServiceAccount)(nil)).Elem()
}

func (i GetInstanceServiceAccountArgs) ToGetInstanceServiceAccountOutput() GetInstanceServiceAccountOutput {
	return i.ToGetInstanceServiceAccountOutputWithContext(context.Background())
}

func (i GetInstanceServiceAccountArgs) ToGetInstanceServiceAccountOutputWithContext(ctx context.Context) GetInstanceServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceServiceAccountOutput)
}

type GetInstanceServiceAccountArrayInput interface {
	pulumi.Input

	ToGetInstanceServiceAccountArrayOutput() GetInstanceServiceAccountArrayOutput
	ToGetInstanceServiceAccountArrayOutputWithContext(context.Context) GetInstanceServiceAccountArrayOutput
}

type GetInstanceServiceAccountArray []GetInstanceServiceAccountInput

func (GetInstanceServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceServiceAccount)(nil)).Elem()
}

func (i GetInstanceServiceAccountArray) ToGetInstanceServiceAccountArrayOutput() GetInstanceServiceAccountArrayOutput {
	return i.ToGetInstanceServiceAccountArrayOutputWithContext(context.Background())
}

func (i GetInstanceServiceAccountArray) ToGetInstanceServiceAccountArrayOutputWithContext(ctx context.Context) GetInstanceServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceServiceAccountArrayOutput)
}

type GetInstanceServiceAccountOutput struct { *pulumi.OutputState }

func (GetInstanceServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceServiceAccount)(nil)).Elem()
}

func (o GetInstanceServiceAccountOutput) ToGetInstanceServiceAccountOutput() GetInstanceServiceAccountOutput {
	return o
}

func (o GetInstanceServiceAccountOutput) ToGetInstanceServiceAccountOutputWithContext(ctx context.Context) GetInstanceServiceAccountOutput {
	return o
}

// The service account e-mail address.
func (o GetInstanceServiceAccountOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceServiceAccount) string { return v.Email }).(pulumi.StringOutput)
}

// A list of service scopes.
func (o GetInstanceServiceAccountOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetInstanceServiceAccount) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type GetInstanceServiceAccountArrayOutput struct { *pulumi.OutputState}

func (GetInstanceServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceServiceAccount)(nil)).Elem()
}

func (o GetInstanceServiceAccountArrayOutput) ToGetInstanceServiceAccountArrayOutput() GetInstanceServiceAccountArrayOutput {
	return o
}

func (o GetInstanceServiceAccountArrayOutput) ToGetInstanceServiceAccountArrayOutputWithContext(ctx context.Context) GetInstanceServiceAccountArrayOutput {
	return o
}

func (o GetInstanceServiceAccountArrayOutput) Index(i pulumi.IntInput) GetInstanceServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetInstanceServiceAccount {
		return vs[0].([]GetInstanceServiceAccount)[vs[1].(int)]
	}).(GetInstanceServiceAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceServiceAccountOutput{})
	pulumi.RegisterOutputType(GetInstanceServiceAccountArrayOutput{})
}
