// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get a Compute Region Instance Group within GCE.
// For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroups).
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
// 		opt0 := "instance-group-name"
// 		_, err := compute.GetRegionInstanceGroup(ctx, &compute.GetRegionInstanceGroupArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// The most common use of this datasource will be to fetch information about the instances inside regional managed instance groups, for instance:
func GetRegionInstanceGroup(ctx *pulumi.Context, args *GetRegionInstanceGroupArgs, opts ...pulumi.InvokeOption) (*GetRegionInstanceGroupResult, error) {
	var rv GetRegionInstanceGroupResult
	err := ctx.Invoke("gcp:compute/getRegionInstanceGroup:getRegionInstanceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegionInstanceGroup.
type GetRegionInstanceGroupArgs struct {
	// The name of the instance group.  One of `name` or `selfLink` must be provided.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If `selfLink` is provided, this value is ignored.  If neither `selfLink`
	// nor `project` are provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the resource belongs.  If `selfLink`
	// is provided, this value is ignored.  If neither `selfLink` nor `region` are
	// provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The link to the instance group.  One of `name` or `selfLink` must be provided.
	SelfLink *string `pulumi:"selfLink"`
}

// A collection of values returned by getRegionInstanceGroup.
type GetRegionInstanceGroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of instances in the group, as a list of resources, each containing:
	Instances []GetRegionInstanceGroupInstance `pulumi:"instances"`
	// String port name
	Name     string `pulumi:"name"`
	Project  string `pulumi:"project"`
	Region   string `pulumi:"region"`
	SelfLink string `pulumi:"selfLink"`
	// The number of instances in the group.
	Size int `pulumi:"size"`
}
