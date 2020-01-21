// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package URLMapHeaderActionRequestHeadersToAdd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type URLMapHeaderActionRequestHeadersToAdd struct {
	HeaderName string `pulumi:"headerName"`
	HeaderValue string `pulumi:"headerValue"`
	Replace bool `pulumi:"replace"`
}

type URLMapHeaderActionRequestHeadersToAddInput interface {
	pulumi.Input

	ToURLMapHeaderActionRequestHeadersToAddOutput() URLMapHeaderActionRequestHeadersToAddOutput
	ToURLMapHeaderActionRequestHeadersToAddOutputWithContext(context.Context) URLMapHeaderActionRequestHeadersToAddOutput
}

type URLMapHeaderActionRequestHeadersToAddArgs struct {
	HeaderName pulumi.StringInput `pulumi:"headerName"`
	HeaderValue pulumi.StringInput `pulumi:"headerValue"`
	Replace pulumi.BoolInput `pulumi:"replace"`
}

func (URLMapHeaderActionRequestHeadersToAddArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (i URLMapHeaderActionRequestHeadersToAddArgs) ToURLMapHeaderActionRequestHeadersToAddOutput() URLMapHeaderActionRequestHeadersToAddOutput {
	return i.ToURLMapHeaderActionRequestHeadersToAddOutputWithContext(context.Background())
}

func (i URLMapHeaderActionRequestHeadersToAddArgs) ToURLMapHeaderActionRequestHeadersToAddOutputWithContext(ctx context.Context) URLMapHeaderActionRequestHeadersToAddOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapHeaderActionRequestHeadersToAddOutput)
}

type URLMapHeaderActionRequestHeadersToAddArrayInput interface {
	pulumi.Input

	ToURLMapHeaderActionRequestHeadersToAddArrayOutput() URLMapHeaderActionRequestHeadersToAddArrayOutput
	ToURLMapHeaderActionRequestHeadersToAddArrayOutputWithContext(context.Context) URLMapHeaderActionRequestHeadersToAddArrayOutput
}

type URLMapHeaderActionRequestHeadersToAddArray []URLMapHeaderActionRequestHeadersToAddInput

func (URLMapHeaderActionRequestHeadersToAddArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (i URLMapHeaderActionRequestHeadersToAddArray) ToURLMapHeaderActionRequestHeadersToAddArrayOutput() URLMapHeaderActionRequestHeadersToAddArrayOutput {
	return i.ToURLMapHeaderActionRequestHeadersToAddArrayOutputWithContext(context.Background())
}

func (i URLMapHeaderActionRequestHeadersToAddArray) ToURLMapHeaderActionRequestHeadersToAddArrayOutputWithContext(ctx context.Context) URLMapHeaderActionRequestHeadersToAddArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapHeaderActionRequestHeadersToAddArrayOutput)
}

type URLMapHeaderActionRequestHeadersToAddOutput struct { *pulumi.OutputState }

func (URLMapHeaderActionRequestHeadersToAddOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (o URLMapHeaderActionRequestHeadersToAddOutput) ToURLMapHeaderActionRequestHeadersToAddOutput() URLMapHeaderActionRequestHeadersToAddOutput {
	return o
}

func (o URLMapHeaderActionRequestHeadersToAddOutput) ToURLMapHeaderActionRequestHeadersToAddOutputWithContext(ctx context.Context) URLMapHeaderActionRequestHeadersToAddOutput {
	return o
}

func (o URLMapHeaderActionRequestHeadersToAddOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapHeaderActionRequestHeadersToAdd) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o URLMapHeaderActionRequestHeadersToAddOutput) HeaderValue() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapHeaderActionRequestHeadersToAdd) string { return v.HeaderValue }).(pulumi.StringOutput)
}

func (o URLMapHeaderActionRequestHeadersToAddOutput) Replace() pulumi.BoolOutput {
	return o.ApplyT(func (v URLMapHeaderActionRequestHeadersToAdd) bool { return v.Replace }).(pulumi.BoolOutput)
}

type URLMapHeaderActionRequestHeadersToAddArrayOutput struct { *pulumi.OutputState}

func (URLMapHeaderActionRequestHeadersToAddArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (o URLMapHeaderActionRequestHeadersToAddArrayOutput) ToURLMapHeaderActionRequestHeadersToAddArrayOutput() URLMapHeaderActionRequestHeadersToAddArrayOutput {
	return o
}

func (o URLMapHeaderActionRequestHeadersToAddArrayOutput) ToURLMapHeaderActionRequestHeadersToAddArrayOutputWithContext(ctx context.Context) URLMapHeaderActionRequestHeadersToAddArrayOutput {
	return o
}

func (o URLMapHeaderActionRequestHeadersToAddArrayOutput) Index(i pulumi.IntInput) URLMapHeaderActionRequestHeadersToAddOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) URLMapHeaderActionRequestHeadersToAdd {
		return vs[0].([]URLMapHeaderActionRequestHeadersToAdd)[vs[1].(int)]
	}).(URLMapHeaderActionRequestHeadersToAddOutput)
}

func init() {
	pulumi.RegisterOutputType(URLMapHeaderActionRequestHeadersToAddOutput{})
	pulumi.RegisterOutputType(URLMapHeaderActionRequestHeadersToAddArrayOutput{})
}
