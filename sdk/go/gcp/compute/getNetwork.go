// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a network within GCE from its name.
//
// ## Example Usage
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
// 		_, err := compute.LookupNetwork(ctx, &compute.LookupNetworkArgs{
// 			Name: "default-us-east1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	var rv LookupNetworkResult
	err := ctx.Invoke("gcp:compute/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// The name of the network.
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// Description of this network.
	Description string `pulumi:"description"`
	// The IP address of the gateway.
	GatewayIpv4 string `pulumi:"gatewayIpv4"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The URI of the resource.
	SelfLink string `pulumi:"selfLink"`
	// the list of subnetworks which belong to the network
	SubnetworksSelfLinks []string `pulumi:"subnetworksSelfLinks"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResult, error) {
			args := v.(LookupNetworkArgs)
			r, err := LookupNetwork(ctx, &args, opts...)
			return *r, err
		}).(LookupNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkOutputArgs struct {
	// The name of the network.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

// Description of this network.
func (o LookupNetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

// The IP address of the gateway.
func (o LookupNetworkResultOutput) GatewayIpv4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.GatewayIpv4 }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// The URI of the resource.
func (o LookupNetworkResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// the list of subnetworks which belong to the network
func (o LookupNetworkResultOutput) SubnetworksSelfLinks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.SubnetworksSelfLinks }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}
