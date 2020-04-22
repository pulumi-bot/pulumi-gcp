// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Persistent disks are durable storage devices that function similarly to
// the physical disks in a desktop or a server. Compute Engine manages the
// hardware behind these devices to ensure data redundancy and optimize
// performance for you. Persistent disks are available as either standard
// hard disk drives (HDD) or solid-state drives (SSD).
//
// Persistent disks are located independently from your virtual machine
// instances, so you can detach or move persistent disks to keep your data
// even after you delete your instances. Persistent disk performance scales
// automatically with size, so you can resize your existing persistent disks
// or add more persistent disks to an instance to meet your performance and
// storage space requirements.
//
// Add a persistent disk to your instance when you need reliable and
// affordable storage with consistent performance characteristics.
//
//
// To get more information about RegionDisk, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/regionDisks)
// * How-to Guides
//     * [Adding or Resizing Regional Persistent Disks](https://cloud.google.com/compute/docs/disks/regional-persistent-disk)
//
// > **Warning:** All arguments including the disk encryption key will be stored in the raw
// state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
type RegionDisk struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// -
	// (Optional)
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.  Structure is documented below.
	DiskEncryptionKey RegionDiskDiskEncryptionKeyPtrOutput `pulumi:"diskEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// -
	// (Optional)
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp pulumi.StringOutput `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp pulumi.StringOutput `pulumi:"lastDetachTimestamp"`
	// -
	// (Required)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// -
	// (Optional)
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntOutput `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -
	// (Optional)
	// A reference to the region where the disk resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// -
	// (Required)
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones pulumi.StringArrayOutput `pulumi:"replicaZones"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// -
	// (Optional)
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size pulumi.IntOutput `pulumi:"size"`
	// -
	// (Optional)
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot pulumi.StringPtrOutput `pulumi:"snapshot"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceSnapshotEncryptionKey RegionDiskSourceSnapshotEncryptionKeyPtrOutput `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId pulumi.StringOutput `pulumi:"sourceSnapshotId"`
	// -
	// (Optional)
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewRegionDisk registers a new resource with the given unique name, arguments, and options.
func NewRegionDisk(ctx *pulumi.Context,
	name string, args *RegionDiskArgs, opts ...pulumi.ResourceOption) (*RegionDisk, error) {
	if args == nil || args.ReplicaZones == nil {
		return nil, errors.New("missing required argument 'ReplicaZones'")
	}
	if args == nil {
		args = &RegionDiskArgs{}
	}
	var resource RegionDisk
	err := ctx.RegisterResource("gcp:compute/regionDisk:RegionDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionDisk gets an existing RegionDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionDiskState, opts ...pulumi.ResourceOption) (*RegionDisk, error) {
	var resource RegionDisk
	err := ctx.ReadResource("gcp:compute/regionDisk:RegionDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionDisk resources.
type regionDiskState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.  Structure is documented below.
	DiskEncryptionKey *RegionDiskDiskEncryptionKey `pulumi:"diskEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// -
	// (Optional)
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp *string `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp *string `pulumi:"lastDetachTimestamp"`
	// -
	// (Required)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// -
	// (Optional)
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *int `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
	// -
	// (Required)
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones []string `pulumi:"replicaZones"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// -
	// (Optional)
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size *int `pulumi:"size"`
	// -
	// (Optional)
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot *string `pulumi:"snapshot"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceSnapshotEncryptionKey *RegionDiskSourceSnapshotEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId *string `pulumi:"sourceSnapshotId"`
	// -
	// (Optional)
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `pulumi:"type"`
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users []string `pulumi:"users"`
}

type RegionDiskState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.  Structure is documented below.
	DiskEncryptionKey RegionDiskDiskEncryptionKeyPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// -
	// (Optional)
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp pulumi.StringPtrInput
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp pulumi.StringPtrInput
	// -
	// (Required)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// -
	// (Optional)
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
	// -
	// (Required)
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones pulumi.StringArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// -
	// (Optional)
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size pulumi.IntPtrInput
	// -
	// (Optional)
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot pulumi.StringPtrInput
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceSnapshotEncryptionKey RegionDiskSourceSnapshotEncryptionKeyPtrInput
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId pulumi.StringPtrInput
	// -
	// (Optional)
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrInput
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users pulumi.StringArrayInput
}

func (RegionDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskState)(nil)).Elem()
}

type regionDiskArgs struct {
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.  Structure is documented below.
	DiskEncryptionKey *RegionDiskDiskEncryptionKey `pulumi:"diskEncryptionKey"`
	// -
	// (Optional)
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// -
	// (Required)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// -
	// (Optional)
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *int `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
	// -
	// (Required)
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones []string `pulumi:"replicaZones"`
	// -
	// (Optional)
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size *int `pulumi:"size"`
	// -
	// (Optional)
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot *string `pulumi:"snapshot"`
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceSnapshotEncryptionKey *RegionDiskSourceSnapshotEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// -
	// (Optional)
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a RegionDisk resource.
type RegionDiskArgs struct {
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.  Structure is documented below.
	DiskEncryptionKey RegionDiskDiskEncryptionKeyPtrInput
	// -
	// (Optional)
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// -
	// (Required)
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// -
	// (Optional)
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
	// -
	// (Required)
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones pulumi.StringArrayInput
	// -
	// (Optional)
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size pulumi.IntPtrInput
	// -
	// (Optional)
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot pulumi.StringPtrInput
	// -
	// (Optional)
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.  Structure is documented below.
	SourceSnapshotEncryptionKey RegionDiskSourceSnapshotEncryptionKeyPtrInput
	// -
	// (Optional)
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrInput
}

func (RegionDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskArgs)(nil)).Elem()
}
