// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the IP addresses from different special IP ranges on Google Cloud Platform.
//
// ## Example Usage
// ### Cloud Ranges
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		netblock, err := compute.GetNetblockIPRanges(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cidrBlocks", netblock.CidrBlocks)
// 		ctx.Export("cidrBlocksIpv4", netblock.CidrBlocksIpv4s)
// 		ctx.Export("cidrBlocksIpv6", netblock.CidrBlocksIpv6s)
// 		return nil
// 	})
// }
// ```
// ### Allow Health Checks
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "legacy-health-checkers"
// 		legacy_hcs, err := compute.GetNetblockIPRanges(ctx, &compute.GetNetblockIPRangesArgs{
// 			RangeType: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetwork(ctx, "_default", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewFirewall(ctx, "allow_hcs", &compute.FirewallArgs{
// 			Network: _default.Name,
// 			Allows: compute.FirewallAllowArray{
// 				&compute.FirewallAllowArgs{
// 					Protocol: pulumi.String("tcp"),
// 					Ports: pulumi.StringArray{
// 						pulumi.String("80"),
// 					},
// 				},
// 			},
// 			SourceRanges: toPulumiStringArray(legacy_hcs.CidrBlocksIpv4s),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// func toPulumiStringArray(arr []string) pulumi.StringArray {
// 	var pulumiArr pulumi.StringArray
// 	for _, v := range arr {
// 		pulumiArr = append(pulumiArr, pulumi.String(v))
// 	}
// 	return pulumiArr
// }
// ```
func GetNetblockIPRanges(ctx *pulumi.Context, args *GetNetblockIPRangesArgs, opts ...pulumi.InvokeOption) (*GetNetblockIPRangesResult, error) {
	var rv GetNetblockIPRangesResult
	err := ctx.Invoke("gcp:compute/getNetblockIPRanges:getNetblockIPRanges", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetblockIPRanges.
type GetNetblockIPRangesArgs struct {
	// The type of range for which to provide results.
	RangeType *string `pulumi:"rangeType"`
}

// A collection of values returned by getNetblockIPRanges.
type GetNetblockIPRangesResult struct {
	// Retrieve list of all CIDR blocks.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// Retrieve list of the IPv4 CIDR blocks
	CidrBlocksIpv4s []string `pulumi:"cidrBlocksIpv4s"`
	// Retrieve list of the IPv6 CIDR blocks, if available.
	CidrBlocksIpv6s []string `pulumi:"cidrBlocksIpv6s"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	RangeType *string `pulumi:"rangeType"`
}

func GetNetblockIPRangesApply(ctx *pulumi.Context, args GetNetblockIPRangesApplyInput, opts ...pulumi.InvokeOption) GetNetblockIPRangesResultOutput {
	return args.ToGetNetblockIPRangesApplyOutput().ApplyT(func(v GetNetblockIPRangesArgs) (GetNetblockIPRangesResult, error) {
		r, err := GetNetblockIPRanges(ctx, &v, opts...)
		return *r, err

	}).(GetNetblockIPRangesResultOutput)
}

// GetNetblockIPRangesApplyInput is an input type that accepts GetNetblockIPRangesApplyArgs and GetNetblockIPRangesApplyOutput values.
// You can construct a concrete instance of `GetNetblockIPRangesApplyInput` via:
//
//          GetNetblockIPRangesApplyArgs{...}
type GetNetblockIPRangesApplyInput interface {
	pulumi.Input

	ToGetNetblockIPRangesApplyOutput() GetNetblockIPRangesApplyOutput
	ToGetNetblockIPRangesApplyOutputWithContext(context.Context) GetNetblockIPRangesApplyOutput
}

// A collection of arguments for invoking getNetblockIPRanges.
type GetNetblockIPRangesApplyArgs struct {
	// The type of range for which to provide results.
	RangeType pulumi.StringPtrInput `pulumi:"rangeType"`
}

func (GetNetblockIPRangesApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetblockIPRangesArgs)(nil)).Elem()
}

func (i GetNetblockIPRangesApplyArgs) ToGetNetblockIPRangesApplyOutput() GetNetblockIPRangesApplyOutput {
	return i.ToGetNetblockIPRangesApplyOutputWithContext(context.Background())
}

func (i GetNetblockIPRangesApplyArgs) ToGetNetblockIPRangesApplyOutputWithContext(ctx context.Context) GetNetblockIPRangesApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNetblockIPRangesApplyOutput)
}

// A collection of arguments for invoking getNetblockIPRanges.
type GetNetblockIPRangesApplyOutput struct{ *pulumi.OutputState }

func (GetNetblockIPRangesApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetblockIPRangesArgs)(nil)).Elem()
}

func (o GetNetblockIPRangesApplyOutput) ToGetNetblockIPRangesApplyOutput() GetNetblockIPRangesApplyOutput {
	return o
}

func (o GetNetblockIPRangesApplyOutput) ToGetNetblockIPRangesApplyOutputWithContext(ctx context.Context) GetNetblockIPRangesApplyOutput {
	return o
}

// The type of range for which to provide results.
func (o GetNetblockIPRangesApplyOutput) RangeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetblockIPRangesArgs) *string { return v.RangeType }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getNetblockIPRanges.
type GetNetblockIPRangesResultOutput struct{ *pulumi.OutputState }

func (GetNetblockIPRangesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetblockIPRangesResult)(nil)).Elem()
}

func (o GetNetblockIPRangesResultOutput) ToGetNetblockIPRangesResultOutput() GetNetblockIPRangesResultOutput {
	return o
}

func (o GetNetblockIPRangesResultOutput) ToGetNetblockIPRangesResultOutputWithContext(ctx context.Context) GetNetblockIPRangesResultOutput {
	return o
}

// Retrieve list of all CIDR blocks.
func (o GetNetblockIPRangesResultOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetblockIPRangesResult) []string { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

// Retrieve list of the IPv4 CIDR blocks
func (o GetNetblockIPRangesResultOutput) CidrBlocksIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetblockIPRangesResult) []string { return v.CidrBlocksIpv4s }).(pulumi.StringArrayOutput)
}

// Retrieve list of the IPv6 CIDR blocks, if available.
func (o GetNetblockIPRangesResultOutput) CidrBlocksIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetblockIPRangesResult) []string { return v.CidrBlocksIpv6s }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNetblockIPRangesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetblockIPRangesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNetblockIPRangesResultOutput) RangeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetblockIPRangesResult) *string { return v.RangeType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetblockIPRangesApplyOutput{})
	pulumi.RegisterOutputType(GetNetblockIPRangesResultOutput{})
}
