// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
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
