// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package FirewallAllow

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type FirewallAllow struct {
	Ports []string `pulumi:"ports"`
	Protocol string `pulumi:"protocol"`
}

type FirewallAllowInput interface {
	pulumi.Input

	ToFirewallAllowOutput() FirewallAllowOutput
	ToFirewallAllowOutputWithContext(context.Context) FirewallAllowOutput
}

type FirewallAllowArgs struct {
	Ports pulumi.StringArrayInput `pulumi:"ports"`
	Protocol pulumi.StringInput `pulumi:"protocol"`
}

func (FirewallAllowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallAllow)(nil)).Elem()
}

func (i FirewallAllowArgs) ToFirewallAllowOutput() FirewallAllowOutput {
	return i.ToFirewallAllowOutputWithContext(context.Background())
}

func (i FirewallAllowArgs) ToFirewallAllowOutputWithContext(ctx context.Context) FirewallAllowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAllowOutput)
}

type FirewallAllowArrayInput interface {
	pulumi.Input

	ToFirewallAllowArrayOutput() FirewallAllowArrayOutput
	ToFirewallAllowArrayOutputWithContext(context.Context) FirewallAllowArrayOutput
}

type FirewallAllowArray []FirewallAllowInput

func (FirewallAllowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallAllow)(nil)).Elem()
}

func (i FirewallAllowArray) ToFirewallAllowArrayOutput() FirewallAllowArrayOutput {
	return i.ToFirewallAllowArrayOutputWithContext(context.Background())
}

func (i FirewallAllowArray) ToFirewallAllowArrayOutputWithContext(ctx context.Context) FirewallAllowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAllowArrayOutput)
}

type FirewallAllowOutput struct { *pulumi.OutputState }

func (FirewallAllowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallAllow)(nil)).Elem()
}

func (o FirewallAllowOutput) ToFirewallAllowOutput() FirewallAllowOutput {
	return o
}

func (o FirewallAllowOutput) ToFirewallAllowOutputWithContext(ctx context.Context) FirewallAllowOutput {
	return o
}

func (o FirewallAllowOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func (v FirewallAllow) []string { return v.Ports }).(pulumi.StringArrayOutput)
}

func (o FirewallAllowOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func (v FirewallAllow) string { return v.Protocol }).(pulumi.StringOutput)
}

type FirewallAllowArrayOutput struct { *pulumi.OutputState}

func (FirewallAllowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallAllow)(nil)).Elem()
}

func (o FirewallAllowArrayOutput) ToFirewallAllowArrayOutput() FirewallAllowArrayOutput {
	return o
}

func (o FirewallAllowArrayOutput) ToFirewallAllowArrayOutputWithContext(ctx context.Context) FirewallAllowArrayOutput {
	return o
}

func (o FirewallAllowArrayOutput) Index(i pulumi.IntInput) FirewallAllowOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) FirewallAllow {
		return vs[0].([]FirewallAllow)[vs[1].(int)]
	}).(FirewallAllowOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallAllowOutput{})
	pulumi.RegisterOutputType(FirewallAllowArrayOutput{})
}
