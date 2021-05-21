// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides access to available Google Compute zones in a region for a given project.
// See more about [regions and zones](https://cloud.google.com/compute/docs/regions-zones/regions-zones) in the upstream docs.
func GetZones(ctx *pulumi.Context, args *GetZonesArgs, opts ...pulumi.InvokeOption) (*GetZonesResult, error) {
	var rv GetZonesResult
	err := ctx.Invoke("gcp:compute/getZones:getZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// Project from which to list available zones. Defaults to project declared in the provider.
	Project *string `pulumi:"project"`
	// Region from which to list available zones. Defaults to region declared in the provider.
	Region *string `pulumi:"region"`
	// Allows to filter list of zones based on their current status. Status can be either `UP` or `DOWN`.
	// Defaults to no filtering (all available zones - both `UP` and `DOWN`).
	Status *string `pulumi:"status"`
}

// A collection of values returned by getZones.
type GetZonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zones available in the given region
	Names   []string `pulumi:"names"`
	Project string   `pulumi:"project"`
	Region  *string  `pulumi:"region"`
	Status  *string  `pulumi:"status"`
}

func GetZonesApply(ctx *pulumi.Context, args GetZonesApplyInput, opts ...pulumi.InvokeOption) GetZonesResultOutput {
	return args.ToGetZonesApplyOutput().ApplyT(func(v GetZonesArgs) (GetZonesResult, error) {
		r, err := GetZones(ctx, &v, opts...)
		return *r, err

	}).(GetZonesResultOutput)
}

// GetZonesApplyInput is an input type that accepts GetZonesApplyArgs and GetZonesApplyOutput values.
// You can construct a concrete instance of `GetZonesApplyInput` via:
//
//          GetZonesApplyArgs{...}
type GetZonesApplyInput interface {
	pulumi.Input

	ToGetZonesApplyOutput() GetZonesApplyOutput
	ToGetZonesApplyOutputWithContext(context.Context) GetZonesApplyOutput
}

// A collection of arguments for invoking getZones.
type GetZonesApplyArgs struct {
	// Project from which to list available zones. Defaults to project declared in the provider.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// Region from which to list available zones. Defaults to region declared in the provider.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Allows to filter list of zones based on their current status. Status can be either `UP` or `DOWN`.
	// Defaults to no filtering (all available zones - both `UP` and `DOWN`).
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetZonesApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesArgs)(nil)).Elem()
}

func (i GetZonesApplyArgs) ToGetZonesApplyOutput() GetZonesApplyOutput {
	return i.ToGetZonesApplyOutputWithContext(context.Background())
}

func (i GetZonesApplyArgs) ToGetZonesApplyOutputWithContext(ctx context.Context) GetZonesApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZonesApplyOutput)
}

// A collection of arguments for invoking getZones.
type GetZonesApplyOutput struct{ *pulumi.OutputState }

func (GetZonesApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesArgs)(nil)).Elem()
}

func (o GetZonesApplyOutput) ToGetZonesApplyOutput() GetZonesApplyOutput {
	return o
}

func (o GetZonesApplyOutput) ToGetZonesApplyOutputWithContext(ctx context.Context) GetZonesApplyOutput {
	return o
}

// Project from which to list available zones. Defaults to project declared in the provider.
func (o GetZonesApplyOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesArgs) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// Region from which to list available zones. Defaults to region declared in the provider.
func (o GetZonesApplyOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesArgs) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// Allows to filter list of zones based on their current status. Status can be either `UP` or `DOWN`.
// Defaults to no filtering (all available zones - both `UP` and `DOWN`).
func (o GetZonesApplyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesArgs) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getZones.
type GetZonesResultOutput struct{ *pulumi.OutputState }

func (GetZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesResult)(nil)).Elem()
}

func (o GetZonesResultOutput) ToGetZonesResultOutput() GetZonesResultOutput {
	return o
}

func (o GetZonesResultOutput) ToGetZonesResultOutputWithContext(ctx context.Context) GetZonesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of zones available in the given region
func (o GetZonesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetZonesResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetZonesResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZonesApplyOutput{})
	pulumi.RegisterOutputType(GetZonesResultOutput{})
}
