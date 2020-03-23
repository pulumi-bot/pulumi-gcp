// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures the Google Compute Engine
// [Default Network Tier](https://cloud.google.com/network-tiers/docs/using-network-service-tiers#setting_the_tier_for_all_resources_in_a_project)
// for a project.
//
// For more information, see,
// [the Project API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/projects/setDefaultNetworkTier).
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_project_default_network_tier.html.markdown.
type ProjectDefaultNetworkTier struct {
	pulumi.CustomResourceState

	// The default network tier to be configured for the project.
	// This field can take the following values: `PREMIUM` or `STANDARD`.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectDefaultNetworkTier registers a new resource with the given unique name, arguments, and options.
func NewProjectDefaultNetworkTier(ctx *pulumi.Context,
	name string, args *ProjectDefaultNetworkTierArgs, opts ...pulumi.ResourceOption) (*ProjectDefaultNetworkTier, error) {
	if args == nil || args.NetworkTier == nil {
		return nil, errors.New("missing required argument 'NetworkTier'")
	}
	if args == nil {
		args = &ProjectDefaultNetworkTierArgs{}
	}
	var resource ProjectDefaultNetworkTier
	err := ctx.RegisterResource("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectDefaultNetworkTier gets an existing ProjectDefaultNetworkTier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectDefaultNetworkTier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectDefaultNetworkTierState, opts ...pulumi.ResourceOption) (*ProjectDefaultNetworkTier, error) {
	var resource ProjectDefaultNetworkTier
	err := ctx.ReadResource("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectDefaultNetworkTier resources.
type projectDefaultNetworkTierState struct {
	// The default network tier to be configured for the project.
	// This field can take the following values: `PREMIUM` or `STANDARD`.
	NetworkTier *string `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ProjectDefaultNetworkTierState struct {
	// The default network tier to be configured for the project.
	// This field can take the following values: `PREMIUM` or `STANDARD`.
	NetworkTier pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectDefaultNetworkTierState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectDefaultNetworkTierState)(nil)).Elem()
}

type projectDefaultNetworkTierArgs struct {
	// The default network tier to be configured for the project.
	// This field can take the following values: `PREMIUM` or `STANDARD`.
	NetworkTier string `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectDefaultNetworkTier resource.
type ProjectDefaultNetworkTierArgs struct {
	// The default network tier to be configured for the project.
	// This field can take the following values: `PREMIUM` or `STANDARD`.
	NetworkTier pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectDefaultNetworkTierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectDefaultNetworkTierArgs)(nil)).Elem()
}

