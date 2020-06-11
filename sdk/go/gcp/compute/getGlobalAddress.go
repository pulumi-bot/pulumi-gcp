// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get the IP address from a static address reserved for a Global Forwarding Rule which are only used for HTTP load balancing. For more information see
// the official [API](https://cloud.google.com/compute/docs/reference/latest/globalAddresses) documentation.
//
// ## Example Usage
//
//
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
// 		myAddress, err := compute.LookupGlobalAddress(ctx, &compute.LookupGlobalAddressArgs{
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
// 		frontend, err := dns.NewRecordSet(ctx, "frontend", &dns.RecordSetArgs{
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
func LookupGlobalAddress(ctx *pulumi.Context, args *LookupGlobalAddressArgs, opts ...pulumi.InvokeOption) (*LookupGlobalAddressResult, error) {
	var rv LookupGlobalAddressResult
	err := ctx.Invoke("gcp:compute/getGlobalAddress:getGlobalAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGlobalAddress.
type LookupGlobalAddressArgs struct {
	// A unique name for the resource, required by GCE.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getGlobalAddress.
type LookupGlobalAddressResult struct {
	// The IP of the created resource.
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink string `pulumi:"selfLink"`
	// Indicates if the address is used. Possible values are: RESERVED or IN_USE.
	Status string `pulumi:"status"`
}
