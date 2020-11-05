// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Adds existing resource policies to a disk. You can only add one policy
// which will be applied to this disk for scheduling snapshot creation.
//
// > **Note:** This resource does not support zonal disks (`compute.Disk`). For zonal disks, please refer to the `compute.DiskResourcePolicyAttachment` resource.
//
// ## Example Usage
type RegionDiskResourcePolicyAttachment struct {
	pulumi.CustomResourceState

	// The name of the regional disk in which the resource policies are attached to.
	Disk pulumi.StringOutput `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewRegionDiskResourcePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewRegionDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, args *RegionDiskResourcePolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*RegionDiskResourcePolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Disk == nil {
		return nil, errors.New("invalid value for required argument 'Disk'")
	}
	var resource RegionDiskResourcePolicyAttachment
	err := ctx.RegisterResource("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionDiskResourcePolicyAttachment gets an existing RegionDiskResourcePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionDiskResourcePolicyAttachmentState, opts ...pulumi.ResourceOption) (*RegionDiskResourcePolicyAttachment, error) {
	var resource RegionDiskResourcePolicyAttachment
	err := ctx.ReadResource("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionDiskResourcePolicyAttachment resources.
type regionDiskResourcePolicyAttachmentState struct {
	// The name of the regional disk in which the resource policies are attached to.
	Disk *string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
}

type RegionDiskResourcePolicyAttachmentState struct {
	// The name of the regional disk in which the resource policies are attached to.
	Disk pulumi.StringPtrInput
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
}

func (RegionDiskResourcePolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskResourcePolicyAttachmentState)(nil)).Elem()
}

type regionDiskResourcePolicyAttachmentArgs struct {
	// The name of the regional disk in which the resource policies are attached to.
	Disk string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RegionDiskResourcePolicyAttachment resource.
type RegionDiskResourcePolicyAttachmentArgs struct {
	// The name of the regional disk in which the resource policies are attached to.
	Disk pulumi.StringInput
	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
}

func (RegionDiskResourcePolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskResourcePolicyAttachmentArgs)(nil)).Elem()
}
