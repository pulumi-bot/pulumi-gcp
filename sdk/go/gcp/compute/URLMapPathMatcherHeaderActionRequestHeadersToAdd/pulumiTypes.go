// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package URLMapPathMatcherHeaderActionRequestHeadersToAdd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type URLMapPathMatcherHeaderActionRequestHeadersToAdd struct {
	HeaderName string `pulumi:"headerName"`
	HeaderValue string `pulumi:"headerValue"`
	Replace bool `pulumi:"replace"`
}

type URLMapPathMatcherHeaderActionRequestHeadersToAddInput interface {
	pulumi.Input

	ToURLMapPathMatcherHeaderActionRequestHeadersToAddOutput() URLMapPathMatcherHeaderActionRequestHeadersToAddOutput
	ToURLMapPathMatcherHeaderActionRequestHeadersToAddOutputWithContext(context.Context) URLMapPathMatcherHeaderActionRequestHeadersToAddOutput
}

type URLMapPathMatcherHeaderActionRequestHeadersToAddArgs struct {
	HeaderName pulumi.StringInput `pulumi:"headerName"`
	HeaderValue pulumi.StringInput `pulumi:"headerValue"`
	Replace pulumi.BoolInput `pulumi:"replace"`
}

func (URLMapPathMatcherHeaderActionRequestHeadersToAddArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapPathMatcherHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (i URLMapPathMatcherHeaderActionRequestHeadersToAddArgs) ToURLMapPathMatcherHeaderActionRequestHeadersToAddOutput() URLMapPathMatcherHeaderActionRequestHeadersToAddOutput {
	return i.ToURLMapPathMatcherHeaderActionRequestHeadersToAddOutputWithContext(context.Background())
}

func (i URLMapPathMatcherHeaderActionRequestHeadersToAddArgs) ToURLMapPathMatcherHeaderActionRequestHeadersToAddOutputWithContext(ctx context.Context) URLMapPathMatcherHeaderActionRequestHeadersToAddOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapPathMatcherHeaderActionRequestHeadersToAddOutput)
}

type URLMapPathMatcherHeaderActionRequestHeadersToAddArrayInput interface {
	pulumi.Input

	ToURLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput() URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput
	ToURLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutputWithContext(context.Context) URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput
}

type URLMapPathMatcherHeaderActionRequestHeadersToAddArray []URLMapPathMatcherHeaderActionRequestHeadersToAddInput

func (URLMapPathMatcherHeaderActionRequestHeadersToAddArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapPathMatcherHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (i URLMapPathMatcherHeaderActionRequestHeadersToAddArray) ToURLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput() URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput {
	return i.ToURLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutputWithContext(context.Background())
}

func (i URLMapPathMatcherHeaderActionRequestHeadersToAddArray) ToURLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutputWithContext(ctx context.Context) URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput)
}

type URLMapPathMatcherHeaderActionRequestHeadersToAddOutput struct { *pulumi.OutputState }

func (URLMapPathMatcherHeaderActionRequestHeadersToAddOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*URLMapPathMatcherHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddOutput) ToURLMapPathMatcherHeaderActionRequestHeadersToAddOutput() URLMapPathMatcherHeaderActionRequestHeadersToAddOutput {
	return o
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddOutput) ToURLMapPathMatcherHeaderActionRequestHeadersToAddOutputWithContext(ctx context.Context) URLMapPathMatcherHeaderActionRequestHeadersToAddOutput {
	return o
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapPathMatcherHeaderActionRequestHeadersToAdd) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddOutput) HeaderValue() pulumi.StringOutput {
	return o.ApplyT(func (v URLMapPathMatcherHeaderActionRequestHeadersToAdd) string { return v.HeaderValue }).(pulumi.StringOutput)
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddOutput) Replace() pulumi.BoolOutput {
	return o.ApplyT(func (v URLMapPathMatcherHeaderActionRequestHeadersToAdd) bool { return v.Replace }).(pulumi.BoolOutput)
}

type URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput struct { *pulumi.OutputState}

func (URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]URLMapPathMatcherHeaderActionRequestHeadersToAdd)(nil)).Elem()
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput) ToURLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput() URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput {
	return o
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput) ToURLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutputWithContext(ctx context.Context) URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput {
	return o
}

func (o URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput) Index(i pulumi.IntInput) URLMapPathMatcherHeaderActionRequestHeadersToAddOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) URLMapPathMatcherHeaderActionRequestHeadersToAdd {
		return vs[0].([]URLMapPathMatcherHeaderActionRequestHeadersToAdd)[vs[1].(int)]
	}).(URLMapPathMatcherHeaderActionRequestHeadersToAddOutput)
}

func init() {
	pulumi.RegisterOutputType(URLMapPathMatcherHeaderActionRequestHeadersToAddOutput{})
	pulumi.RegisterOutputType(URLMapPathMatcherHeaderActionRequestHeadersToAddArrayOutput{})
}
