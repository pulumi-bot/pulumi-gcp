// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getInstanceSchedulingNodeAffinity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetInstanceSchedulingNodeAffinity struct {
	Key string `pulumi:"key"`
	Operator string `pulumi:"operator"`
	Values []string `pulumi:"values"`
}

type GetInstanceSchedulingNodeAffinityInput interface {
	pulumi.Input

	ToGetInstanceSchedulingNodeAffinityOutput() GetInstanceSchedulingNodeAffinityOutput
	ToGetInstanceSchedulingNodeAffinityOutputWithContext(context.Context) GetInstanceSchedulingNodeAffinityOutput
}

type GetInstanceSchedulingNodeAffinityArgs struct {
	Key pulumi.StringInput `pulumi:"key"`
	Operator pulumi.StringInput `pulumi:"operator"`
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GetInstanceSchedulingNodeAffinityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceSchedulingNodeAffinity)(nil)).Elem()
}

func (i GetInstanceSchedulingNodeAffinityArgs) ToGetInstanceSchedulingNodeAffinityOutput() GetInstanceSchedulingNodeAffinityOutput {
	return i.ToGetInstanceSchedulingNodeAffinityOutputWithContext(context.Background())
}

func (i GetInstanceSchedulingNodeAffinityArgs) ToGetInstanceSchedulingNodeAffinityOutputWithContext(ctx context.Context) GetInstanceSchedulingNodeAffinityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceSchedulingNodeAffinityOutput)
}

type GetInstanceSchedulingNodeAffinityArrayInput interface {
	pulumi.Input

	ToGetInstanceSchedulingNodeAffinityArrayOutput() GetInstanceSchedulingNodeAffinityArrayOutput
	ToGetInstanceSchedulingNodeAffinityArrayOutputWithContext(context.Context) GetInstanceSchedulingNodeAffinityArrayOutput
}

type GetInstanceSchedulingNodeAffinityArray []GetInstanceSchedulingNodeAffinityInput

func (GetInstanceSchedulingNodeAffinityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceSchedulingNodeAffinity)(nil)).Elem()
}

func (i GetInstanceSchedulingNodeAffinityArray) ToGetInstanceSchedulingNodeAffinityArrayOutput() GetInstanceSchedulingNodeAffinityArrayOutput {
	return i.ToGetInstanceSchedulingNodeAffinityArrayOutputWithContext(context.Background())
}

func (i GetInstanceSchedulingNodeAffinityArray) ToGetInstanceSchedulingNodeAffinityArrayOutputWithContext(ctx context.Context) GetInstanceSchedulingNodeAffinityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceSchedulingNodeAffinityArrayOutput)
}

type GetInstanceSchedulingNodeAffinityOutput struct { *pulumi.OutputState }

func (GetInstanceSchedulingNodeAffinityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceSchedulingNodeAffinity)(nil)).Elem()
}

func (o GetInstanceSchedulingNodeAffinityOutput) ToGetInstanceSchedulingNodeAffinityOutput() GetInstanceSchedulingNodeAffinityOutput {
	return o
}

func (o GetInstanceSchedulingNodeAffinityOutput) ToGetInstanceSchedulingNodeAffinityOutputWithContext(ctx context.Context) GetInstanceSchedulingNodeAffinityOutput {
	return o
}

func (o GetInstanceSchedulingNodeAffinityOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceSchedulingNodeAffinity) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetInstanceSchedulingNodeAffinityOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func (v GetInstanceSchedulingNodeAffinity) string { return v.Operator }).(pulumi.StringOutput)
}

func (o GetInstanceSchedulingNodeAffinityOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetInstanceSchedulingNodeAffinity) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GetInstanceSchedulingNodeAffinityArrayOutput struct { *pulumi.OutputState}

func (GetInstanceSchedulingNodeAffinityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceSchedulingNodeAffinity)(nil)).Elem()
}

func (o GetInstanceSchedulingNodeAffinityArrayOutput) ToGetInstanceSchedulingNodeAffinityArrayOutput() GetInstanceSchedulingNodeAffinityArrayOutput {
	return o
}

func (o GetInstanceSchedulingNodeAffinityArrayOutput) ToGetInstanceSchedulingNodeAffinityArrayOutputWithContext(ctx context.Context) GetInstanceSchedulingNodeAffinityArrayOutput {
	return o
}

func (o GetInstanceSchedulingNodeAffinityArrayOutput) Index(i pulumi.IntInput) GetInstanceSchedulingNodeAffinityOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetInstanceSchedulingNodeAffinity {
		return vs[0].([]GetInstanceSchedulingNodeAffinity)[vs[1].(int)]
	}).(GetInstanceSchedulingNodeAffinityOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceSchedulingNodeAffinityOutput{})
	pulumi.RegisterOutputType(GetInstanceSchedulingNodeAffinityArrayOutput{})
}
