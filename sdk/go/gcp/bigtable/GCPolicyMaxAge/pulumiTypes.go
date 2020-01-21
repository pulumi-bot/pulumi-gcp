// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package GCPolicyMaxAge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GCPolicyMaxAge struct {
	// Number of days before applying GC policy.
	Days int `pulumi:"days"`
}

type GCPolicyMaxAgeInput interface {
	pulumi.Input

	ToGCPolicyMaxAgeOutput() GCPolicyMaxAgeOutput
	ToGCPolicyMaxAgeOutputWithContext(context.Context) GCPolicyMaxAgeOutput
}

type GCPolicyMaxAgeArgs struct {
	// Number of days before applying GC policy.
	Days pulumi.IntInput `pulumi:"days"`
}

func (GCPolicyMaxAgeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPolicyMaxAge)(nil)).Elem()
}

func (i GCPolicyMaxAgeArgs) ToGCPolicyMaxAgeOutput() GCPolicyMaxAgeOutput {
	return i.ToGCPolicyMaxAgeOutputWithContext(context.Background())
}

func (i GCPolicyMaxAgeArgs) ToGCPolicyMaxAgeOutputWithContext(ctx context.Context) GCPolicyMaxAgeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyMaxAgeOutput)
}

type GCPolicyMaxAgeArrayInput interface {
	pulumi.Input

	ToGCPolicyMaxAgeArrayOutput() GCPolicyMaxAgeArrayOutput
	ToGCPolicyMaxAgeArrayOutputWithContext(context.Context) GCPolicyMaxAgeArrayOutput
}

type GCPolicyMaxAgeArray []GCPolicyMaxAgeInput

func (GCPolicyMaxAgeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GCPolicyMaxAge)(nil)).Elem()
}

func (i GCPolicyMaxAgeArray) ToGCPolicyMaxAgeArrayOutput() GCPolicyMaxAgeArrayOutput {
	return i.ToGCPolicyMaxAgeArrayOutputWithContext(context.Background())
}

func (i GCPolicyMaxAgeArray) ToGCPolicyMaxAgeArrayOutputWithContext(ctx context.Context) GCPolicyMaxAgeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyMaxAgeArrayOutput)
}

type GCPolicyMaxAgeOutput struct { *pulumi.OutputState }

func (GCPolicyMaxAgeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPolicyMaxAge)(nil)).Elem()
}

func (o GCPolicyMaxAgeOutput) ToGCPolicyMaxAgeOutput() GCPolicyMaxAgeOutput {
	return o
}

func (o GCPolicyMaxAgeOutput) ToGCPolicyMaxAgeOutputWithContext(ctx context.Context) GCPolicyMaxAgeOutput {
	return o
}

// Number of days before applying GC policy.
func (o GCPolicyMaxAgeOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func (v GCPolicyMaxAge) int { return v.Days }).(pulumi.IntOutput)
}

type GCPolicyMaxAgeArrayOutput struct { *pulumi.OutputState}

func (GCPolicyMaxAgeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GCPolicyMaxAge)(nil)).Elem()
}

func (o GCPolicyMaxAgeArrayOutput) ToGCPolicyMaxAgeArrayOutput() GCPolicyMaxAgeArrayOutput {
	return o
}

func (o GCPolicyMaxAgeArrayOutput) ToGCPolicyMaxAgeArrayOutputWithContext(ctx context.Context) GCPolicyMaxAgeArrayOutput {
	return o
}

func (o GCPolicyMaxAgeArrayOutput) Index(i pulumi.IntInput) GCPolicyMaxAgeOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GCPolicyMaxAge {
		return vs[0].([]GCPolicyMaxAge)[vs[1].(int)]
	}).(GCPolicyMaxAgeOutput)
}

func init() {
	pulumi.RegisterOutputType(GCPolicyMaxAgeOutput{})
	pulumi.RegisterOutputType(GCPolicyMaxAgeArrayOutput{})
}
