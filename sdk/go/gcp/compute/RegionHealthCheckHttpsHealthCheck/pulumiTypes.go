// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegionHealthCheckHttpsHealthCheck

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegionHealthCheckHttpsHealthCheck struct {
	Host *string `pulumi:"host"`
	Port *int `pulumi:"port"`
	PortName *string `pulumi:"portName"`
	PortSpecification *string `pulumi:"portSpecification"`
	ProxyHeader *string `pulumi:"proxyHeader"`
	RequestPath *string `pulumi:"requestPath"`
	Response *string `pulumi:"response"`
}

type RegionHealthCheckHttpsHealthCheckInput interface {
	pulumi.Input

	ToRegionHealthCheckHttpsHealthCheckOutput() RegionHealthCheckHttpsHealthCheckOutput
	ToRegionHealthCheckHttpsHealthCheckOutputWithContext(context.Context) RegionHealthCheckHttpsHealthCheckOutput
}

type RegionHealthCheckHttpsHealthCheckArgs struct {
	Host pulumi.StringPtrInput `pulumi:"host"`
	Port pulumi.IntPtrInput `pulumi:"port"`
	PortName pulumi.StringPtrInput `pulumi:"portName"`
	PortSpecification pulumi.StringPtrInput `pulumi:"portSpecification"`
	ProxyHeader pulumi.StringPtrInput `pulumi:"proxyHeader"`
	RequestPath pulumi.StringPtrInput `pulumi:"requestPath"`
	Response pulumi.StringPtrInput `pulumi:"response"`
}

func (RegionHealthCheckHttpsHealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionHealthCheckHttpsHealthCheck)(nil)).Elem()
}

func (i RegionHealthCheckHttpsHealthCheckArgs) ToRegionHealthCheckHttpsHealthCheckOutput() RegionHealthCheckHttpsHealthCheckOutput {
	return i.ToRegionHealthCheckHttpsHealthCheckOutputWithContext(context.Background())
}

func (i RegionHealthCheckHttpsHealthCheckArgs) ToRegionHealthCheckHttpsHealthCheckOutputWithContext(ctx context.Context) RegionHealthCheckHttpsHealthCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionHealthCheckHttpsHealthCheckOutput)
}

func (i RegionHealthCheckHttpsHealthCheckArgs) ToRegionHealthCheckHttpsHealthCheckPtrOutput() RegionHealthCheckHttpsHealthCheckPtrOutput {
	return i.ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(context.Background())
}

func (i RegionHealthCheckHttpsHealthCheckArgs) ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(ctx context.Context) RegionHealthCheckHttpsHealthCheckPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionHealthCheckHttpsHealthCheckOutput).ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(ctx)
}

type RegionHealthCheckHttpsHealthCheckPtrInput interface {
	pulumi.Input

	ToRegionHealthCheckHttpsHealthCheckPtrOutput() RegionHealthCheckHttpsHealthCheckPtrOutput
	ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(context.Context) RegionHealthCheckHttpsHealthCheckPtrOutput
}

type regionHealthCheckHttpsHealthCheckPtrType RegionHealthCheckHttpsHealthCheckArgs

func RegionHealthCheckHttpsHealthCheckPtr(v *RegionHealthCheckHttpsHealthCheckArgs) RegionHealthCheckHttpsHealthCheckPtrInput {	return (*regionHealthCheckHttpsHealthCheckPtrType)(v)
}

func (*regionHealthCheckHttpsHealthCheckPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionHealthCheckHttpsHealthCheck)(nil)).Elem()
}

func (i *regionHealthCheckHttpsHealthCheckPtrType) ToRegionHealthCheckHttpsHealthCheckPtrOutput() RegionHealthCheckHttpsHealthCheckPtrOutput {
	return i.ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(context.Background())
}

func (i *regionHealthCheckHttpsHealthCheckPtrType) ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(ctx context.Context) RegionHealthCheckHttpsHealthCheckPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionHealthCheckHttpsHealthCheckPtrOutput)
}

type RegionHealthCheckHttpsHealthCheckOutput struct { *pulumi.OutputState }

func (RegionHealthCheckHttpsHealthCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionHealthCheckHttpsHealthCheck)(nil)).Elem()
}

func (o RegionHealthCheckHttpsHealthCheckOutput) ToRegionHealthCheckHttpsHealthCheckOutput() RegionHealthCheckHttpsHealthCheckOutput {
	return o
}

func (o RegionHealthCheckHttpsHealthCheckOutput) ToRegionHealthCheckHttpsHealthCheckOutputWithContext(ctx context.Context) RegionHealthCheckHttpsHealthCheckOutput {
	return o
}

func (o RegionHealthCheckHttpsHealthCheckOutput) ToRegionHealthCheckHttpsHealthCheckPtrOutput() RegionHealthCheckHttpsHealthCheckPtrOutput {
	return o.ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(context.Background())
}

func (o RegionHealthCheckHttpsHealthCheckOutput) ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(ctx context.Context) RegionHealthCheckHttpsHealthCheckPtrOutput {
	return o.ApplyT(func(v RegionHealthCheckHttpsHealthCheck) *RegionHealthCheckHttpsHealthCheck {
		return &v
	}).(RegionHealthCheckHttpsHealthCheckPtrOutput)
}
func (o RegionHealthCheckHttpsHealthCheckOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckOutput) PortName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.PortName }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckOutput) PortSpecification() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.PortSpecification }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckOutput) ProxyHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.ProxyHeader }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckOutput) RequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.RequestPath }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckOutput) Response() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.Response }).(pulumi.StringPtrOutput)
}

type RegionHealthCheckHttpsHealthCheckPtrOutput struct { *pulumi.OutputState}

func (RegionHealthCheckHttpsHealthCheckPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionHealthCheckHttpsHealthCheck)(nil)).Elem()
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) ToRegionHealthCheckHttpsHealthCheckPtrOutput() RegionHealthCheckHttpsHealthCheckPtrOutput {
	return o
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) ToRegionHealthCheckHttpsHealthCheckPtrOutputWithContext(ctx context.Context) RegionHealthCheckHttpsHealthCheckPtrOutput {
	return o
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) Elem() RegionHealthCheckHttpsHealthCheckOutput {
	return o.ApplyT(func (v *RegionHealthCheckHttpsHealthCheck) RegionHealthCheckHttpsHealthCheck { return *v }).(RegionHealthCheckHttpsHealthCheckOutput)
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) PortName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.PortName }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) PortSpecification() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.PortSpecification }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) ProxyHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.ProxyHeader }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) RequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.RequestPath }).(pulumi.StringPtrOutput)
}

func (o RegionHealthCheckHttpsHealthCheckPtrOutput) Response() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegionHealthCheckHttpsHealthCheck) *string { return v.Response }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionHealthCheckHttpsHealthCheckOutput{})
	pulumi.RegisterOutputType(RegionHealthCheckHttpsHealthCheckPtrOutput{})
}
