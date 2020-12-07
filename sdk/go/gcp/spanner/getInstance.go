// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get a spanner instance from Google Cloud by its name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.LookupInstance(ctx, &spanner.LookupInstanceArgs{
// 			Name: "bar",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("gcp:spanner/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	Config      *string           `pulumi:"config"`
	DisplayName *string           `pulumi:"displayName"`
	Labels      map[string]string `pulumi:"labels"`
	// The name of the spanner instance.
	Name     string `pulumi:"name"`
	NumNodes *int   `pulumi:"numNodes"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	Config      *string `pulumi:"config"`
	DisplayName *string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id       string            `pulumi:"id"`
	Labels   map[string]string `pulumi:"labels"`
	Name     string            `pulumi:"name"`
	NumNodes *int              `pulumi:"numNodes"`
	Project  *string           `pulumi:"project"`
	State    string            `pulumi:"state"`
}
