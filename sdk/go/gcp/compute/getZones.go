// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
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
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zones available in the given region
	Names   []string `pulumi:"names"`
	Project string   `pulumi:"project"`
	Region  *string  `pulumi:"region"`
	Status  *string  `pulumi:"status"`
}
