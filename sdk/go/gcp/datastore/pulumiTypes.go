// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datastore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DataStoreIndexProperty struct {
	Direction string `pulumi:"direction"`
	Name      string `pulumi:"name"`
}

type DataStoreIndexPropertyInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput
	ToDataStoreIndexPropertyOutputWithContext(context.Context) DataStoreIndexPropertyOutput
}

type DataStoreIndexPropertyArgs struct {
	Direction pulumi.StringInput `pulumi:"direction"`
	Name      pulumi.StringInput `pulumi:"name"`
}

func (DataStoreIndexPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexProperty)(nil)).Elem()
}

func (i DataStoreIndexPropertyArgs) ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput {
	return i.ToDataStoreIndexPropertyOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArgs) ToDataStoreIndexPropertyOutputWithContext(ctx context.Context) DataStoreIndexPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyOutput)
}

type DataStoreIndexPropertyArrayInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput
	ToDataStoreIndexPropertyArrayOutputWithContext(context.Context) DataStoreIndexPropertyArrayOutput
}

type DataStoreIndexPropertyArray []DataStoreIndexPropertyInput

func (DataStoreIndexPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexProperty)(nil)).Elem()
}

func (i DataStoreIndexPropertyArray) ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput {
	return i.ToDataStoreIndexPropertyArrayOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArray) ToDataStoreIndexPropertyArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyArrayOutput)
}

type DataStoreIndexPropertyOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexProperty)(nil)).Elem()
}

func (o DataStoreIndexPropertyOutput) ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput {
	return o
}

func (o DataStoreIndexPropertyOutput) ToDataStoreIndexPropertyOutputWithContext(ctx context.Context) DataStoreIndexPropertyOutput {
	return o
}

func (o DataStoreIndexPropertyOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexProperty) string { return v.Direction }).(pulumi.StringOutput)
}

func (o DataStoreIndexPropertyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexProperty) string { return v.Name }).(pulumi.StringOutput)
}

type DataStoreIndexPropertyArrayOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexProperty)(nil)).Elem()
}

func (o DataStoreIndexPropertyArrayOutput) ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArrayOutput) ToDataStoreIndexPropertyArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArrayOutput) Index(i pulumi.IntInput) DataStoreIndexPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataStoreIndexProperty {
		return vs[0].([]DataStoreIndexProperty)[vs[1].(int)]
	}).(DataStoreIndexPropertyOutput)
}

type DataStoreIndexPropertyArgs struct {
	Direction string `pulumi:"direction"`
	Name      string `pulumi:"name"`
}

type DataStoreIndexPropertyArgsInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyArgsOutput() DataStoreIndexPropertyArgsOutput
	ToDataStoreIndexPropertyArgsOutputWithContext(context.Context) DataStoreIndexPropertyArgsOutput
}

type DataStoreIndexPropertyArgsArgs struct {
	Direction pulumi.StringInput `pulumi:"direction"`
	Name      pulumi.StringInput `pulumi:"name"`
}

func (DataStoreIndexPropertyArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexPropertyArgs)(nil)).Elem()
}

func (i DataStoreIndexPropertyArgsArgs) ToDataStoreIndexPropertyArgsOutput() DataStoreIndexPropertyArgsOutput {
	return i.ToDataStoreIndexPropertyArgsOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArgsArgs) ToDataStoreIndexPropertyArgsOutputWithContext(ctx context.Context) DataStoreIndexPropertyArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyArgsOutput)
}

type DataStoreIndexPropertyArgsArrayInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyArgsArrayOutput() DataStoreIndexPropertyArgsArrayOutput
	ToDataStoreIndexPropertyArgsArrayOutputWithContext(context.Context) DataStoreIndexPropertyArgsArrayOutput
}

type DataStoreIndexPropertyArgsArray []DataStoreIndexPropertyArgsInput

func (DataStoreIndexPropertyArgsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexPropertyArgs)(nil)).Elem()
}

func (i DataStoreIndexPropertyArgsArray) ToDataStoreIndexPropertyArgsArrayOutput() DataStoreIndexPropertyArgsArrayOutput {
	return i.ToDataStoreIndexPropertyArgsArrayOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArgsArray) ToDataStoreIndexPropertyArgsArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArgsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyArgsArrayOutput)
}

type DataStoreIndexPropertyArgsOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexPropertyArgs)(nil)).Elem()
}

func (o DataStoreIndexPropertyArgsOutput) ToDataStoreIndexPropertyArgsOutput() DataStoreIndexPropertyArgsOutput {
	return o
}

func (o DataStoreIndexPropertyArgsOutput) ToDataStoreIndexPropertyArgsOutputWithContext(ctx context.Context) DataStoreIndexPropertyArgsOutput {
	return o
}

func (o DataStoreIndexPropertyArgsOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexPropertyArgs) string { return v.Direction }).(pulumi.StringOutput)
}

func (o DataStoreIndexPropertyArgsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexPropertyArgs) string { return v.Name }).(pulumi.StringOutput)
}

type DataStoreIndexPropertyArgsArrayOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyArgsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexPropertyArgs)(nil)).Elem()
}

func (o DataStoreIndexPropertyArgsArrayOutput) ToDataStoreIndexPropertyArgsArrayOutput() DataStoreIndexPropertyArgsArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArgsArrayOutput) ToDataStoreIndexPropertyArgsArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArgsArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArgsArrayOutput) Index(i pulumi.IntInput) DataStoreIndexPropertyArgsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataStoreIndexPropertyArgs {
		return vs[0].([]DataStoreIndexPropertyArgs)[vs[1].(int)]
	}).(DataStoreIndexPropertyArgsOutput)
}

type DataStoreIndexPropertyState struct {
	Direction string `pulumi:"direction"`
	Name      string `pulumi:"name"`
}

type DataStoreIndexPropertyStateInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyStateOutput() DataStoreIndexPropertyStateOutput
	ToDataStoreIndexPropertyStateOutputWithContext(context.Context) DataStoreIndexPropertyStateOutput
}

type DataStoreIndexPropertyStateArgs struct {
	Direction pulumi.StringInput `pulumi:"direction"`
	Name      pulumi.StringInput `pulumi:"name"`
}

func (DataStoreIndexPropertyStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexPropertyState)(nil)).Elem()
}

func (i DataStoreIndexPropertyStateArgs) ToDataStoreIndexPropertyStateOutput() DataStoreIndexPropertyStateOutput {
	return i.ToDataStoreIndexPropertyStateOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyStateArgs) ToDataStoreIndexPropertyStateOutputWithContext(ctx context.Context) DataStoreIndexPropertyStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyStateOutput)
}

type DataStoreIndexPropertyStateArrayInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyStateArrayOutput() DataStoreIndexPropertyStateArrayOutput
	ToDataStoreIndexPropertyStateArrayOutputWithContext(context.Context) DataStoreIndexPropertyStateArrayOutput
}

type DataStoreIndexPropertyStateArray []DataStoreIndexPropertyStateInput

func (DataStoreIndexPropertyStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexPropertyState)(nil)).Elem()
}

func (i DataStoreIndexPropertyStateArray) ToDataStoreIndexPropertyStateArrayOutput() DataStoreIndexPropertyStateArrayOutput {
	return i.ToDataStoreIndexPropertyStateArrayOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyStateArray) ToDataStoreIndexPropertyStateArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyStateArrayOutput)
}

type DataStoreIndexPropertyStateOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexPropertyState)(nil)).Elem()
}

func (o DataStoreIndexPropertyStateOutput) ToDataStoreIndexPropertyStateOutput() DataStoreIndexPropertyStateOutput {
	return o
}

func (o DataStoreIndexPropertyStateOutput) ToDataStoreIndexPropertyStateOutputWithContext(ctx context.Context) DataStoreIndexPropertyStateOutput {
	return o
}

func (o DataStoreIndexPropertyStateOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexPropertyState) string { return v.Direction }).(pulumi.StringOutput)
}

func (o DataStoreIndexPropertyStateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexPropertyState) string { return v.Name }).(pulumi.StringOutput)
}

type DataStoreIndexPropertyStateArrayOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexPropertyState)(nil)).Elem()
}

func (o DataStoreIndexPropertyStateArrayOutput) ToDataStoreIndexPropertyStateArrayOutput() DataStoreIndexPropertyStateArrayOutput {
	return o
}

func (o DataStoreIndexPropertyStateArrayOutput) ToDataStoreIndexPropertyStateArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyStateArrayOutput {
	return o
}

func (o DataStoreIndexPropertyStateArrayOutput) Index(i pulumi.IntInput) DataStoreIndexPropertyStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataStoreIndexPropertyState {
		return vs[0].([]DataStoreIndexPropertyState)[vs[1].(int)]
	}).(DataStoreIndexPropertyStateOutput)
}

func init() {
	pulumi.RegisterOutputType(DataStoreIndexPropertyOutput{})
	pulumi.RegisterOutputType(DataStoreIndexPropertyArrayOutput{})
	pulumi.RegisterOutputType(DataStoreIndexPropertyArgsOutput{})
	pulumi.RegisterOutputType(DataStoreIndexPropertyArgsArrayOutput{})
	pulumi.RegisterOutputType(DataStoreIndexPropertyStateOutput{})
	pulumi.RegisterOutputType(DataStoreIndexPropertyStateArrayOutput{})
}
