// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
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
// To get more information about RegionDisk, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionDisks)
// * How-to Guides
//     * [Adding or Resizing Regional Persistent Disks](https://cloud.google.com/compute/docs/disks/regional-persistent-disk)
//
// > **Warning:** All arguments including `disk_encryption_key.raw_key` will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
//
// ## Example Usage
// ### Region Disk Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		disk, err := compute.NewDisk(ctx, "disk", &compute.DiskArgs{
// 			Image: pulumi.String("debian-cloud/debian-9"),
// 			Size:  pulumi.Int(50),
// 			Type:  pulumi.String("pd-ssd"),
// 			Zone:  pulumi.String("us-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		snapdisk, err := compute.NewSnapshot(ctx, "snapdisk", &compute.SnapshotArgs{
// 			SourceDisk: disk.Name,
// 			Zone:       pulumi.String("us-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionDisk(ctx, "regiondisk", &compute.RegionDiskArgs{
// 			Snapshot:               snapdisk.ID(),
// 			Type:                   pulumi.String("pd-ssd"),
// 			Region:                 pulumi.String("us-central1"),
// 			PhysicalBlockSizeBytes: pulumi.Int(4096),
// 			ReplicaZones: pulumi.StringArray{
// 				pulumi.String("us-central1-a"),
// 				pulumi.String("us-central1-f"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// RegionDisk can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/regionDisk:RegionDisk default projects/{{project}}/regions/{{region}}/disks/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionDisk:RegionDisk default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionDisk:RegionDisk default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionDisk:RegionDisk default {{name}}
// ```
type RegionDisk struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey RegionDiskDiskEncryptionKeyPtrOutput `pulumi:"diskEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp pulumi.StringOutput `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp pulumi.StringOutput `pulumi:"lastDetachTimestamp"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntOutput `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones pulumi.StringArrayOutput `pulumi:"replicaZones"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size pulumi.IntOutput `pulumi:"size"`
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot pulumi.StringPtrOutput `pulumi:"snapshot"`
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey RegionDiskSourceSnapshotEncryptionKeyPtrOutput `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId pulumi.StringOutput `pulumi:"sourceSnapshotId"`
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewRegionDisk registers a new resource with the given unique name, arguments, and options.
func NewRegionDisk(ctx *pulumi.Context,
	name string, args *RegionDiskArgs, opts ...pulumi.ResourceOption) (*RegionDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReplicaZones == nil {
		return nil, errors.New("invalid value for required argument 'ReplicaZones'")
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
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey *RegionDiskDiskEncryptionKey `pulumi:"diskEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp *string `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp *string `pulumi:"lastDetachTimestamp"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *int `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones []string `pulumi:"replicaZones"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size *int `pulumi:"size"`
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot *string `pulumi:"snapshot"`
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey *RegionDiskSourceSnapshotEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId *string `pulumi:"sourceSnapshotId"`
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `pulumi:"type"`
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users []string `pulumi:"users"`
}

type RegionDiskState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey RegionDiskDiskEncryptionKeyPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp pulumi.StringPtrInput
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones pulumi.StringArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size pulumi.IntPtrInput
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey RegionDiskSourceSnapshotEncryptionKeyPtrInput
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId pulumi.StringPtrInput
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
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey *RegionDiskDiskEncryptionKey `pulumi:"diskEncryptionKey"`
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *int `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the disk resides.
	Region *string `pulumi:"region"`
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones []string `pulumi:"replicaZones"`
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size *int `pulumi:"size"`
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot *string `pulumi:"snapshot"`
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey *RegionDiskSourceSnapshotEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a RegionDisk resource.
type RegionDiskArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey RegionDiskDiskEncryptionKeyPtrInput
	// Labels to apply to this disk.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the disk resides.
	Region pulumi.StringPtrInput
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones pulumi.StringArrayInput
	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the sourceImage or
	// sourceSnapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with sourceImage or sourceSnapshot,
	// the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size pulumi.IntPtrInput
	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. For example, the following are
	// valid values:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
	// * `projects/project/global/snapshots/snapshot`
	// * `global/snapshots/snapshot`
	// * `snapshot`
	Snapshot pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey RegionDiskSourceSnapshotEncryptionKeyPtrInput
	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrInput
}

func (RegionDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskArgs)(nil)).Elem()
}

type RegionDiskInput interface {
	pulumi.Input

	ToRegionDiskOutput() RegionDiskOutput
	ToRegionDiskOutputWithContext(ctx context.Context) RegionDiskOutput
}

func (*RegionDisk) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDisk)(nil))
}

func (i *RegionDisk) ToRegionDiskOutput() RegionDiskOutput {
	return i.ToRegionDiskOutputWithContext(context.Background())
}

func (i *RegionDisk) ToRegionDiskOutputWithContext(ctx context.Context) RegionDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskOutput)
}

func (i *RegionDisk) ToRegionDiskPtrOutput() RegionDiskPtrOutput {
	return i.ToRegionDiskPtrOutputWithContext(context.Background())
}

func (i *RegionDisk) ToRegionDiskPtrOutputWithContext(ctx context.Context) RegionDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskPtrOutput)
}

type RegionDiskPtrInput interface {
	pulumi.Input

	ToRegionDiskPtrOutput() RegionDiskPtrOutput
	ToRegionDiskPtrOutputWithContext(ctx context.Context) RegionDiskPtrOutput
}

type regionDiskPtrType RegionDiskArgs

func (*regionDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDisk)(nil))
}

func (i *regionDiskPtrType) ToRegionDiskPtrOutput() RegionDiskPtrOutput {
	return i.ToRegionDiskPtrOutputWithContext(context.Background())
}

func (i *regionDiskPtrType) ToRegionDiskPtrOutputWithContext(ctx context.Context) RegionDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskOutput).ToRegionDiskPtrOutput()
}

// RegionDiskArrayInput is an input type that accepts RegionDiskArray and RegionDiskArrayOutput values.
// You can construct a concrete instance of `RegionDiskArrayInput` via:
//
//          RegionDiskArray{ RegionDiskArgs{...} }
type RegionDiskArrayInput interface {
	pulumi.Input

	ToRegionDiskArrayOutput() RegionDiskArrayOutput
	ToRegionDiskArrayOutputWithContext(context.Context) RegionDiskArrayOutput
}

type RegionDiskArray []RegionDiskInput

func (RegionDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionDisk)(nil))
}

func (i RegionDiskArray) ToRegionDiskArrayOutput() RegionDiskArrayOutput {
	return i.ToRegionDiskArrayOutputWithContext(context.Background())
}

func (i RegionDiskArray) ToRegionDiskArrayOutputWithContext(ctx context.Context) RegionDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskArrayOutput)
}

// RegionDiskMapInput is an input type that accepts RegionDiskMap and RegionDiskMapOutput values.
// You can construct a concrete instance of `RegionDiskMapInput` via:
//
//          RegionDiskMap{ "key": RegionDiskArgs{...} }
type RegionDiskMapInput interface {
	pulumi.Input

	ToRegionDiskMapOutput() RegionDiskMapOutput
	ToRegionDiskMapOutputWithContext(context.Context) RegionDiskMapOutput
}

type RegionDiskMap map[string]RegionDiskInput

func (RegionDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegionDisk)(nil))
}

func (i RegionDiskMap) ToRegionDiskMapOutput() RegionDiskMapOutput {
	return i.ToRegionDiskMapOutputWithContext(context.Background())
}

func (i RegionDiskMap) ToRegionDiskMapOutputWithContext(ctx context.Context) RegionDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskMapOutput)
}

type RegionDiskOutput struct {
	*pulumi.OutputState
}

func (RegionDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDisk)(nil))
}

func (o RegionDiskOutput) ToRegionDiskOutput() RegionDiskOutput {
	return o
}

func (o RegionDiskOutput) ToRegionDiskOutputWithContext(ctx context.Context) RegionDiskOutput {
	return o
}

func (o RegionDiskOutput) ToRegionDiskPtrOutput() RegionDiskPtrOutput {
	return o.ToRegionDiskPtrOutputWithContext(context.Background())
}

func (o RegionDiskOutput) ToRegionDiskPtrOutputWithContext(ctx context.Context) RegionDiskPtrOutput {
	return o.ApplyT(func(v RegionDisk) *RegionDisk {
		return &v
	}).(RegionDiskPtrOutput)
}

type RegionDiskPtrOutput struct {
	*pulumi.OutputState
}

func (RegionDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDisk)(nil))
}

func (o RegionDiskPtrOutput) ToRegionDiskPtrOutput() RegionDiskPtrOutput {
	return o
}

func (o RegionDiskPtrOutput) ToRegionDiskPtrOutputWithContext(ctx context.Context) RegionDiskPtrOutput {
	return o
}

type RegionDiskArrayOutput struct{ *pulumi.OutputState }

func (RegionDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionDisk)(nil))
}

func (o RegionDiskArrayOutput) ToRegionDiskArrayOutput() RegionDiskArrayOutput {
	return o
}

func (o RegionDiskArrayOutput) ToRegionDiskArrayOutputWithContext(ctx context.Context) RegionDiskArrayOutput {
	return o
}

func (o RegionDiskArrayOutput) Index(i pulumi.IntInput) RegionDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegionDisk {
		return vs[0].([]RegionDisk)[vs[1].(int)]
	}).(RegionDiskOutput)
}

type RegionDiskMapOutput struct{ *pulumi.OutputState }

func (RegionDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegionDisk)(nil))
}

func (o RegionDiskMapOutput) ToRegionDiskMapOutput() RegionDiskMapOutput {
	return o
}

func (o RegionDiskMapOutput) ToRegionDiskMapOutputWithContext(ctx context.Context) RegionDiskMapOutput {
	return o
}

func (o RegionDiskMapOutput) MapIndex(k pulumi.StringInput) RegionDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RegionDisk {
		return vs[0].(map[string]RegionDisk)[vs[1].(string)]
	}).(RegionDiskOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionDiskOutput{})
	pulumi.RegisterOutputType(RegionDiskPtrOutput{})
	pulumi.RegisterOutputType(RegionDiskArrayOutput{})
	pulumi.RegisterOutputType(RegionDiskMapOutput{})
}
