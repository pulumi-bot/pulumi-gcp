// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ServiceApi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/endpoints/ServiceApiMethod"
)

type ServiceApi struct {
	Methods []endpointsServiceApiMethod.ServiceApiMethod `pulumi:"methods"`
	Name *string `pulumi:"name"`
	Syntax *string `pulumi:"syntax"`
	Version *string `pulumi:"version"`
}

type ServiceApiInput interface {
	pulumi.Input

	ToServiceApiOutput() ServiceApiOutput
	ToServiceApiOutputWithContext(context.Context) ServiceApiOutput
}

type ServiceApiArgs struct {
	Methods endpointsServiceApiMethod.ServiceApiMethodArrayInput `pulumi:"methods"`
	Name pulumi.StringPtrInput `pulumi:"name"`
	Syntax pulumi.StringPtrInput `pulumi:"syntax"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ServiceApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceApi)(nil)).Elem()
}

func (i ServiceApiArgs) ToServiceApiOutput() ServiceApiOutput {
	return i.ToServiceApiOutputWithContext(context.Background())
}

func (i ServiceApiArgs) ToServiceApiOutputWithContext(ctx context.Context) ServiceApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceApiOutput)
}

type ServiceApiArrayInput interface {
	pulumi.Input

	ToServiceApiArrayOutput() ServiceApiArrayOutput
	ToServiceApiArrayOutputWithContext(context.Context) ServiceApiArrayOutput
}

type ServiceApiArray []ServiceApiInput

func (ServiceApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceApi)(nil)).Elem()
}

func (i ServiceApiArray) ToServiceApiArrayOutput() ServiceApiArrayOutput {
	return i.ToServiceApiArrayOutputWithContext(context.Background())
}

func (i ServiceApiArray) ToServiceApiArrayOutputWithContext(ctx context.Context) ServiceApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceApiArrayOutput)
}

type ServiceApiOutput struct { *pulumi.OutputState }

func (ServiceApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceApi)(nil)).Elem()
}

func (o ServiceApiOutput) ToServiceApiOutput() ServiceApiOutput {
	return o
}

func (o ServiceApiOutput) ToServiceApiOutputWithContext(ctx context.Context) ServiceApiOutput {
	return o
}

func (o ServiceApiOutput) Methods() endpointsServiceApiMethod.ServiceApiMethodArrayOutput {
	return o.ApplyT(func (v ServiceApi) []endpointsServiceApiMethod.ServiceApiMethod { return v.Methods }).(endpointsServiceApiMethod.ServiceApiMethodArrayOutput)
}

func (o ServiceApiOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApi) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceApiOutput) Syntax() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApi) *string { return v.Syntax }).(pulumi.StringPtrOutput)
}

func (o ServiceApiOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceApi) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ServiceApiArrayOutput struct { *pulumi.OutputState}

func (ServiceApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceApi)(nil)).Elem()
}

func (o ServiceApiArrayOutput) ToServiceApiArrayOutput() ServiceApiArrayOutput {
	return o
}

func (o ServiceApiArrayOutput) ToServiceApiArrayOutputWithContext(ctx context.Context) ServiceApiArrayOutput {
	return o
}

func (o ServiceApiArrayOutput) Index(i pulumi.IntInput) ServiceApiOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ServiceApi {
		return vs[0].([]ServiceApi)[vs[1].(int)]
	}).(ServiceApiOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceApiOutput{})
	pulumi.RegisterOutputType(ServiceApiArrayOutput{})
}
