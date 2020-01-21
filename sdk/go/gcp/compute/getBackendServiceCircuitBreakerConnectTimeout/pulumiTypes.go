// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getBackendServiceCircuitBreakerConnectTimeout

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetBackendServiceCircuitBreakerConnectTimeout struct {
	Nanos int `pulumi:"nanos"`
	Seconds int `pulumi:"seconds"`
}

type GetBackendServiceCircuitBreakerConnectTimeoutInput interface {
	pulumi.Input

	ToGetBackendServiceCircuitBreakerConnectTimeoutOutput() GetBackendServiceCircuitBreakerConnectTimeoutOutput
	ToGetBackendServiceCircuitBreakerConnectTimeoutOutputWithContext(context.Context) GetBackendServiceCircuitBreakerConnectTimeoutOutput
}

type GetBackendServiceCircuitBreakerConnectTimeoutArgs struct {
	Nanos pulumi.IntInput `pulumi:"nanos"`
	Seconds pulumi.IntInput `pulumi:"seconds"`
}

func (GetBackendServiceCircuitBreakerConnectTimeoutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendServiceCircuitBreakerConnectTimeout)(nil)).Elem()
}

func (i GetBackendServiceCircuitBreakerConnectTimeoutArgs) ToGetBackendServiceCircuitBreakerConnectTimeoutOutput() GetBackendServiceCircuitBreakerConnectTimeoutOutput {
	return i.ToGetBackendServiceCircuitBreakerConnectTimeoutOutputWithContext(context.Background())
}

func (i GetBackendServiceCircuitBreakerConnectTimeoutArgs) ToGetBackendServiceCircuitBreakerConnectTimeoutOutputWithContext(ctx context.Context) GetBackendServiceCircuitBreakerConnectTimeoutOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackendServiceCircuitBreakerConnectTimeoutOutput)
}

type GetBackendServiceCircuitBreakerConnectTimeoutArrayInput interface {
	pulumi.Input

	ToGetBackendServiceCircuitBreakerConnectTimeoutArrayOutput() GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput
	ToGetBackendServiceCircuitBreakerConnectTimeoutArrayOutputWithContext(context.Context) GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput
}

type GetBackendServiceCircuitBreakerConnectTimeoutArray []GetBackendServiceCircuitBreakerConnectTimeoutInput

func (GetBackendServiceCircuitBreakerConnectTimeoutArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackendServiceCircuitBreakerConnectTimeout)(nil)).Elem()
}

func (i GetBackendServiceCircuitBreakerConnectTimeoutArray) ToGetBackendServiceCircuitBreakerConnectTimeoutArrayOutput() GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput {
	return i.ToGetBackendServiceCircuitBreakerConnectTimeoutArrayOutputWithContext(context.Background())
}

func (i GetBackendServiceCircuitBreakerConnectTimeoutArray) ToGetBackendServiceCircuitBreakerConnectTimeoutArrayOutputWithContext(ctx context.Context) GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput)
}

type GetBackendServiceCircuitBreakerConnectTimeoutOutput struct { *pulumi.OutputState }

func (GetBackendServiceCircuitBreakerConnectTimeoutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendServiceCircuitBreakerConnectTimeout)(nil)).Elem()
}

func (o GetBackendServiceCircuitBreakerConnectTimeoutOutput) ToGetBackendServiceCircuitBreakerConnectTimeoutOutput() GetBackendServiceCircuitBreakerConnectTimeoutOutput {
	return o
}

func (o GetBackendServiceCircuitBreakerConnectTimeoutOutput) ToGetBackendServiceCircuitBreakerConnectTimeoutOutputWithContext(ctx context.Context) GetBackendServiceCircuitBreakerConnectTimeoutOutput {
	return o
}

func (o GetBackendServiceCircuitBreakerConnectTimeoutOutput) Nanos() pulumi.IntOutput {
	return o.ApplyT(func (v GetBackendServiceCircuitBreakerConnectTimeout) int { return v.Nanos }).(pulumi.IntOutput)
}

func (o GetBackendServiceCircuitBreakerConnectTimeoutOutput) Seconds() pulumi.IntOutput {
	return o.ApplyT(func (v GetBackendServiceCircuitBreakerConnectTimeout) int { return v.Seconds }).(pulumi.IntOutput)
}

type GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput struct { *pulumi.OutputState}

func (GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackendServiceCircuitBreakerConnectTimeout)(nil)).Elem()
}

func (o GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput) ToGetBackendServiceCircuitBreakerConnectTimeoutArrayOutput() GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput {
	return o
}

func (o GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput) ToGetBackendServiceCircuitBreakerConnectTimeoutArrayOutputWithContext(ctx context.Context) GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput {
	return o
}

func (o GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput) Index(i pulumi.IntInput) GetBackendServiceCircuitBreakerConnectTimeoutOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetBackendServiceCircuitBreakerConnectTimeout {
		return vs[0].([]GetBackendServiceCircuitBreakerConnectTimeout)[vs[1].(int)]
	}).(GetBackendServiceCircuitBreakerConnectTimeoutOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackendServiceCircuitBreakerConnectTimeoutOutput{})
	pulumi.RegisterOutputType(GetBackendServiceCircuitBreakerConnectTimeoutArrayOutput{})
}
