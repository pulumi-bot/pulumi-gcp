// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get the IP address from a static address. For more information see
// the official [API](https://cloud.google.com/compute/docs/reference/latest/addresses/get) documentation.
//
// {{% examples %}}
// ## Example Usage
// {{% example %}}
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/dns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myAddress, err := compute.LookupAddress(ctx, &compute.LookupAddressArgs{
// 			Name: "foobar",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		prod, err := dns.NewManagedZone(ctx, "prod", &dns.ManagedZoneArgs{
// 			DnsName: pulumi.String("prod.mydomain.com."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewRecordSet(ctx, "frontend", &dns.RecordSetArgs{
// 			Type:        pulumi.String("A"),
// 			Ttl:         pulumi.Int(300),
// 			ManagedZone: prod.Name,
// 			Rrdatas: pulumi.StringArray{
// 				pulumi.String(myAddress.Address),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// {{% /example %}}
// {{% /examples %}}
func LookupAddress(ctx *pulumi.Context, args *LookupAddressArgs, opts ...pulumi.InvokeOption) (*LookupAddressResult, error) {
	var rv LookupAddressResult
	err := ctx.Invoke("gcp:compute/getAddress:getAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddress.
type LookupAddressArgs struct {
	// A unique name for the resource, required by GCE.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created address reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getAddress.
type LookupAddressResult struct {
	// The IP of the created resource.
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	Region  string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink string `pulumi:"selfLink"`
	// Indicates if the address is used. Possible values are: RESERVED or IN_USE.
	Status string `pulumi:"status"`
}
