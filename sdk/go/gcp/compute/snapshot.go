// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a Persistent Disk Snapshot resource.
//
// Use snapshots to back up data from your persistent disks. Snapshots are
// different from public images and custom images, which are used primarily
// to create instances or configure instance templates. Snapshots are useful
// for periodic backup of the data on your persistent disks. You can create
// snapshots from persistent disks even while they are attached to running
// instances.
//
// Snapshots are incremental, so you can create regular snapshots on a
// persistent disk faster and at a much lower cost than if you regularly
// created a full image of the disk.
//
//
// To get more information about Snapshot, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
type Snapshot struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Size of the snapshot, specified in GB.
	DiskSizeGb pulumi.IntOutput `pulumi:"diskSizeGb"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// -
	// (Optional)
	// Labels to apply to this Snapshot.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
	// attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
	// encryption key.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// -
	// (Required)
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the snapshot. Required if the
	// source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
	SnapshotEncryptionKey SnapshotSnapshotEncryptionKeyPtrOutput `pulumi:"snapshotEncryptionKey"`
	// The unique identifier for the resource.
	SnapshotId pulumi.IntOutput `pulumi:"snapshotId"`
	// -
	// (Required)
	// A reference to the disk used to create this snapshot.
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceDiskEncryptionKey SnapshotSourceDiskEncryptionKeyPtrOutput `pulumi:"sourceDiskEncryptionKey"`
	SourceDiskLink          pulumi.StringOutput                      `pulumi:"sourceDiskLink"`
	// A size of the the storage used by the snapshot. As snapshots share storage, this number is expected to change with
	// snapshot creation/deletion.
	StorageBytes pulumi.IntOutput `pulumi:"storageBytes"`
	// -
	// (Optional)
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil || args.SourceDisk == nil {
		return nil, errors.New("missing required argument 'SourceDisk'")
	}
	if args == nil {
		args = &SnapshotArgs{}
	}
	var resource Snapshot
	err := ctx.RegisterResource("gcp:compute/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("gcp:compute/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Size of the snapshot, specified in GB.
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// -
	// (Optional)
	// Labels to apply to this Snapshot.
	Labels map[string]string `pulumi:"labels"`
	// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
	// attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
	// encryption key.
	Licenses []string `pulumi:"licenses"`
	// -
	// (Required)
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the snapshot. Required if the
	// source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
	SnapshotEncryptionKey *SnapshotSnapshotEncryptionKey `pulumi:"snapshotEncryptionKey"`
	// The unique identifier for the resource.
	SnapshotId *int `pulumi:"snapshotId"`
	// -
	// (Required)
	// A reference to the disk used to create this snapshot.
	SourceDisk *string `pulumi:"sourceDisk"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceDiskEncryptionKey *SnapshotSourceDiskEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	SourceDiskLink          *string                          `pulumi:"sourceDiskLink"`
	// A size of the the storage used by the snapshot. As snapshots share storage, this number is expected to change with
	// snapshot creation/deletion.
	StorageBytes *int `pulumi:"storageBytes"`
	// -
	// (Optional)
	// A reference to the zone where the disk is hosted.
	Zone *string `pulumi:"zone"`
}

type SnapshotState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Size of the snapshot, specified in GB.
	DiskSizeGb pulumi.IntPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// -
	// (Optional)
	// Labels to apply to this Snapshot.
	Labels pulumi.StringMapInput
	// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
	// attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
	// encryption key.
	Licenses pulumi.StringArrayInput
	// -
	// (Required)
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// -
	// (Optional)
	// The customer-supplied encryption key of the snapshot. Required if the
	// source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
	SnapshotEncryptionKey SnapshotSnapshotEncryptionKeyPtrInput
	// The unique identifier for the resource.
	SnapshotId pulumi.IntPtrInput
	// -
	// (Required)
	// A reference to the disk used to create this snapshot.
	SourceDisk pulumi.StringPtrInput
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceDiskEncryptionKey SnapshotSourceDiskEncryptionKeyPtrInput
	SourceDiskLink          pulumi.StringPtrInput
	// A size of the the storage used by the snapshot. As snapshots share storage, this number is expected to change with
	// snapshot creation/deletion.
	StorageBytes pulumi.IntPtrInput
	// -
	// (Optional)
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// Labels to apply to this Snapshot.
	Labels map[string]string `pulumi:"labels"`
	// -
	// (Required)
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the snapshot. Required if the
	// source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
	SnapshotEncryptionKey *SnapshotSnapshotEncryptionKey `pulumi:"snapshotEncryptionKey"`
	// -
	// (Required)
	// A reference to the disk used to create this snapshot.
	SourceDisk string `pulumi:"sourceDisk"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceDiskEncryptionKey *SnapshotSourceDiskEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	// -
	// (Optional)
	// A reference to the zone where the disk is hosted.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// Labels to apply to this Snapshot.
	Labels pulumi.StringMapInput
	// -
	// (Required)
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// The customer-supplied encryption key of the snapshot. Required if the
	// source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
	SnapshotEncryptionKey SnapshotSnapshotEncryptionKeyPtrInput
	// -
	// (Required)
	// A reference to the disk used to create this snapshot.
	SourceDisk pulumi.StringInput
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceDiskEncryptionKey SnapshotSourceDiskEncryptionKeyPtrInput
	// -
	// (Optional)
	// A reference to the zone where the disk is hosted.
	Zone pulumi.StringPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}
