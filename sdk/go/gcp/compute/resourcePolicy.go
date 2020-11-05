// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A policy that can be attached to a resource to specify or schedule actions on that resource.
//
// ## Example Usage
type ResourcePolicy struct {
	pulumi.CustomResourceState

	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyPtrOutput `pulumi:"groupPlacementPolicy"`
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where resource policy resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyPtrOutput `pulumi:"snapshotSchedulePolicy"`
}

// NewResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicy(ctx *pulumi.Context,
	name string, args *ResourcePolicyArgs, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	if args == nil {
		args = &ResourcePolicyArgs{}
	}
	var resource ResourcePolicy
	err := ctx.RegisterResource("gcp:compute/resourcePolicy:ResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicy gets an existing ResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyState, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	var resource ResourcePolicy
	err := ctx.ReadResource("gcp:compute/resourcePolicy:ResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicy resources.
type resourcePolicyState struct {
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicy `pulumi:"groupPlacementPolicy"`
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where resource policy resides.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicy `pulumi:"snapshotSchedulePolicy"`
}

type ResourcePolicyState struct {
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyPtrInput
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where resource policy resides.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyPtrInput
}

func (ResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyState)(nil)).Elem()
}

type resourcePolicyArgs struct {
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicy `pulumi:"groupPlacementPolicy"`
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where resource policy resides.
	Region *string `pulumi:"region"`
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicy `pulumi:"snapshotSchedulePolicy"`
}

// The set of arguments for constructing a ResourcePolicy resource.
type ResourcePolicyArgs struct {
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyPtrInput
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where resource policy resides.
	Region pulumi.StringPtrInput
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyPtrInput
}

func (ResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyArgs)(nil)).Elem()
}
