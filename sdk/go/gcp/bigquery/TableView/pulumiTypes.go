// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TableView

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TableView struct {
	Query string `pulumi:"query"`
	UseLegacySql *bool `pulumi:"useLegacySql"`
}

type TableViewInput interface {
	pulumi.Input

	ToTableViewOutput() TableViewOutput
	ToTableViewOutputWithContext(context.Context) TableViewOutput
}

type TableViewArgs struct {
	Query pulumi.StringInput `pulumi:"query"`
	UseLegacySql pulumi.BoolPtrInput `pulumi:"useLegacySql"`
}

func (TableViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableView)(nil)).Elem()
}

func (i TableViewArgs) ToTableViewOutput() TableViewOutput {
	return i.ToTableViewOutputWithContext(context.Background())
}

func (i TableViewArgs) ToTableViewOutputWithContext(ctx context.Context) TableViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableViewOutput)
}

func (i TableViewArgs) ToTableViewPtrOutput() TableViewPtrOutput {
	return i.ToTableViewPtrOutputWithContext(context.Background())
}

func (i TableViewArgs) ToTableViewPtrOutputWithContext(ctx context.Context) TableViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableViewOutput).ToTableViewPtrOutputWithContext(ctx)
}

type TableViewPtrInput interface {
	pulumi.Input

	ToTableViewPtrOutput() TableViewPtrOutput
	ToTableViewPtrOutputWithContext(context.Context) TableViewPtrOutput
}

type tableViewPtrType TableViewArgs

func TableViewPtr(v *TableViewArgs) TableViewPtrInput {	return (*tableViewPtrType)(v)
}

func (*tableViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableView)(nil)).Elem()
}

func (i *tableViewPtrType) ToTableViewPtrOutput() TableViewPtrOutput {
	return i.ToTableViewPtrOutputWithContext(context.Background())
}

func (i *tableViewPtrType) ToTableViewPtrOutputWithContext(ctx context.Context) TableViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableViewPtrOutput)
}

type TableViewOutput struct { *pulumi.OutputState }

func (TableViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableView)(nil)).Elem()
}

func (o TableViewOutput) ToTableViewOutput() TableViewOutput {
	return o
}

func (o TableViewOutput) ToTableViewOutputWithContext(ctx context.Context) TableViewOutput {
	return o
}

func (o TableViewOutput) ToTableViewPtrOutput() TableViewPtrOutput {
	return o.ToTableViewPtrOutputWithContext(context.Background())
}

func (o TableViewOutput) ToTableViewPtrOutputWithContext(ctx context.Context) TableViewPtrOutput {
	return o.ApplyT(func(v TableView) *TableView {
		return &v
	}).(TableViewPtrOutput)
}
func (o TableViewOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func (v TableView) string { return v.Query }).(pulumi.StringOutput)
}

func (o TableViewOutput) UseLegacySql() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v TableView) *bool { return v.UseLegacySql }).(pulumi.BoolPtrOutput)
}

type TableViewPtrOutput struct { *pulumi.OutputState}

func (TableViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableView)(nil)).Elem()
}

func (o TableViewPtrOutput) ToTableViewPtrOutput() TableViewPtrOutput {
	return o
}

func (o TableViewPtrOutput) ToTableViewPtrOutputWithContext(ctx context.Context) TableViewPtrOutput {
	return o
}

func (o TableViewPtrOutput) Elem() TableViewOutput {
	return o.ApplyT(func (v *TableView) TableView { return *v }).(TableViewOutput)
}

func (o TableViewPtrOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func (v TableView) string { return v.Query }).(pulumi.StringOutput)
}

func (o TableViewPtrOutput) UseLegacySql() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v TableView) *bool { return v.UseLegacySql }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TableViewOutput{})
	pulumi.RegisterOutputType(TableViewPtrOutput{})
}
