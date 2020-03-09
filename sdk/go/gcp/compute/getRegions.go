// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides access to available Google Compute regions for a given project.
// See more about [regions and regions](https://cloud.google.com/compute/docs/regions-zones/) in the upstream docs.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_compute_regions.html.markdown.
func GetRegions(ctx *pulumi.Context, args *GetRegionsArgs, opts ...pulumi.InvokeOption) (*GetRegionsResult, error) {
	var rv GetRegionsResult
	err := ctx.Invoke("gcp:compute/getRegions:getRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegions.
type GetRegionsArgs struct {
	// Project from which to list available regions. Defaults to project declared in the provider.
	Project *string `pulumi:"project"`
	// Allows to filter list of regions based on their current status. Status can be either `UP` or `DOWN`.
	// Defaults to no filtering (all available regions - both `UP` and `DOWN`).
	Status *string `pulumi:"status"`
}


// A collection of values returned by getRegions.
type GetRegionsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of regions available in the given project
	Names []string `pulumi:"names"`
	Project string `pulumi:"project"`
	Status *string `pulumi:"status"`
}

