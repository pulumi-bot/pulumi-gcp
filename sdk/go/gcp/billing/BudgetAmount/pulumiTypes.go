// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BudgetAmount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/billing/BudgetAmountSpecifiedAmount"
)

type BudgetAmount struct {
	SpecifiedAmount billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmount `pulumi:"specifiedAmount"`
}

type BudgetAmountInput interface {
	pulumi.Input

	ToBudgetAmountOutput() BudgetAmountOutput
	ToBudgetAmountOutputWithContext(context.Context) BudgetAmountOutput
}

type BudgetAmountArgs struct {
	SpecifiedAmount billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmountInput `pulumi:"specifiedAmount"`
}

func (BudgetAmountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetAmount)(nil)).Elem()
}

func (i BudgetAmountArgs) ToBudgetAmountOutput() BudgetAmountOutput {
	return i.ToBudgetAmountOutputWithContext(context.Background())
}

func (i BudgetAmountArgs) ToBudgetAmountOutputWithContext(ctx context.Context) BudgetAmountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetAmountOutput)
}

func (i BudgetAmountArgs) ToBudgetAmountPtrOutput() BudgetAmountPtrOutput {
	return i.ToBudgetAmountPtrOutputWithContext(context.Background())
}

func (i BudgetAmountArgs) ToBudgetAmountPtrOutputWithContext(ctx context.Context) BudgetAmountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetAmountOutput).ToBudgetAmountPtrOutputWithContext(ctx)
}

type BudgetAmountPtrInput interface {
	pulumi.Input

	ToBudgetAmountPtrOutput() BudgetAmountPtrOutput
	ToBudgetAmountPtrOutputWithContext(context.Context) BudgetAmountPtrOutput
}

type budgetAmountPtrType BudgetAmountArgs

func BudgetAmountPtr(v *BudgetAmountArgs) BudgetAmountPtrInput {	return (*budgetAmountPtrType)(v)
}

func (*budgetAmountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetAmount)(nil)).Elem()
}

func (i *budgetAmountPtrType) ToBudgetAmountPtrOutput() BudgetAmountPtrOutput {
	return i.ToBudgetAmountPtrOutputWithContext(context.Background())
}

func (i *budgetAmountPtrType) ToBudgetAmountPtrOutputWithContext(ctx context.Context) BudgetAmountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetAmountPtrOutput)
}

type BudgetAmountOutput struct { *pulumi.OutputState }

func (BudgetAmountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetAmount)(nil)).Elem()
}

func (o BudgetAmountOutput) ToBudgetAmountOutput() BudgetAmountOutput {
	return o
}

func (o BudgetAmountOutput) ToBudgetAmountOutputWithContext(ctx context.Context) BudgetAmountOutput {
	return o
}

func (o BudgetAmountOutput) ToBudgetAmountPtrOutput() BudgetAmountPtrOutput {
	return o.ToBudgetAmountPtrOutputWithContext(context.Background())
}

func (o BudgetAmountOutput) ToBudgetAmountPtrOutputWithContext(ctx context.Context) BudgetAmountPtrOutput {
	return o.ApplyT(func(v BudgetAmount) *BudgetAmount {
		return &v
	}).(BudgetAmountPtrOutput)
}
func (o BudgetAmountOutput) SpecifiedAmount() billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmountOutput {
	return o.ApplyT(func (v BudgetAmount) billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmount { return v.SpecifiedAmount }).(billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmountOutput)
}

type BudgetAmountPtrOutput struct { *pulumi.OutputState}

func (BudgetAmountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetAmount)(nil)).Elem()
}

func (o BudgetAmountPtrOutput) ToBudgetAmountPtrOutput() BudgetAmountPtrOutput {
	return o
}

func (o BudgetAmountPtrOutput) ToBudgetAmountPtrOutputWithContext(ctx context.Context) BudgetAmountPtrOutput {
	return o
}

func (o BudgetAmountPtrOutput) Elem() BudgetAmountOutput {
	return o.ApplyT(func (v *BudgetAmount) BudgetAmount { return *v }).(BudgetAmountOutput)
}

func (o BudgetAmountPtrOutput) SpecifiedAmount() billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmountOutput {
	return o.ApplyT(func (v BudgetAmount) billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmount { return v.SpecifiedAmount }).(billingBudgetAmountSpecifiedAmount.BudgetAmountSpecifiedAmountOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetAmountOutput{})
	pulumi.RegisterOutputType(BudgetAmountPtrOutput{})
}
