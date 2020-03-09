// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DiskResourcePolicyAttachment struct {
	pulumi.CustomResourceState

	// The name of the disk in which the resource policies are attached to.
	Disk pulumi.StringOutput `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the zone where the disk resides.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDiskResourcePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, args *DiskResourcePolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*DiskResourcePolicyAttachment, error) {
	if args == nil || args.Disk == nil {
		return nil, errors.New("missing required argument 'Disk'")
	}
	if args == nil {
		args = &DiskResourcePolicyAttachmentArgs{}
	}
	var resource DiskResourcePolicyAttachment
	err := ctx.RegisterResource("gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskResourcePolicyAttachment gets an existing DiskResourcePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskResourcePolicyAttachmentState, opts ...pulumi.ResourceOption) (*DiskResourcePolicyAttachment, error) {
	var resource DiskResourcePolicyAttachment
	err := ctx.ReadResource("gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskResourcePolicyAttachment resources.
type diskResourcePolicyAttachmentState struct {
	// The name of the disk in which the resource policies are attached to.
	Disk *string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the zone where the disk resides.
	Zone *string `pulumi:"zone"`
}

type DiskResourcePolicyAttachmentState struct {
	// The name of the disk in which the resource policies are attached to.
	Disk pulumi.StringPtrInput
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the zone where the disk resides.
	Zone pulumi.StringPtrInput
}

func (DiskResourcePolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskResourcePolicyAttachmentState)(nil)).Elem()
}

type diskResourcePolicyAttachmentArgs struct {
	// The name of the disk in which the resource policies are attached to.
	Disk string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the zone where the disk resides.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a DiskResourcePolicyAttachment resource.
type DiskResourcePolicyAttachmentArgs struct {
	// The name of the disk in which the resource policies are attached to.
	Disk pulumi.StringInput
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the zone where the disk resides.
	Zone pulumi.StringPtrInput
}

func (DiskResourcePolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskResourcePolicyAttachmentArgs)(nil)).Elem()
}

