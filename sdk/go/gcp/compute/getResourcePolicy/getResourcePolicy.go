// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getResourcePolicy

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_resource_policy.html.markdown.
func GetResourcePolicy(ctx *pulumi.Context, args *GetResourcePolicyArgs, opts ...pulumi.InvokeOption) (*GetResourcePolicyResult, error) {
	var rv GetResourcePolicyResult
	err := ctx.Invoke("gcp:compute/getResourcePolicy:getResourcePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourcePolicy.
type GetResourcePolicyArgs struct {
	// The name of the Resource Policy.
	Name string `pulumi:"name"`
	// Project from which to list the Resource Policy. Defaults to project declared in the provider.
	Project *string `pulumi:"project"`
	// Region where the Resource Policy resides.
	Region string `pulumi:"region"`
}


// A collection of values returned by getResourcePolicy.
type GetResourcePolicyResult struct {
	// Description of this Resource Policy.
	Description string `pulumi:"description"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region string `pulumi:"region"`
	// The URI of the resource.
	SelfLink string `pulumi:"selfLink"`
}

