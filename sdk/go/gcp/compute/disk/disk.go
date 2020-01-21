// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package disk

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/DiskDiskEncryptionKey"
	"https:/github.com/pulumi/pulumi-gcp/compute/DiskSourceImageEncryptionKey"
	"https:/github.com/pulumi/pulumi-gcp/compute/DiskSourceSnapshotEncryptionKey"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_disk.html.markdown.
type Disk struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey computeDiskDiskEncryptionKey.DiskDiskEncryptionKeyPtrOutput `pulumi:"diskEncryptionKey"`
	// The image from which to initialize this disk. This can be one of: the image's 'self_link',
	// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}', 'global/images/{image}',
	// 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}', '{family}', or '{image}'.
	// If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data
	// source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image 'centos-6-v20180104' includes
	// its family name 'centos-6'. These images can be referred by family name here.
	Image pulumi.StringPtrOutput `pulumi:"image"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this disk. A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp pulumi.StringOutput `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp pulumi.StringOutput `pulumi:"lastDetachTimestamp"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntOutput `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies pulumi.StringArrayOutput `pulumi:"resourcePolicies"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field along
	// with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
	Size pulumi.IntOutput `pulumi:"size"`
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot pulumi.StringPtrOutput `pulumi:"snapshot"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a
	// customer-supplied encryption key.
	SourceImageEncryptionKey computeDiskSourceImageEncryptionKey.DiskSourceImageEncryptionKeyPtrOutput `pulumi:"sourceImageEncryptionKey"`
	// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
	// persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
	// under the same name, the source image ID would identify the exact version of the image that was used.
	SourceImageId pulumi.StringOutput `pulumi:"sourceImageId"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey computeDiskSourceSnapshotEncryptionKey.DiskSourceSnapshotEncryptionKeyPtrOutput `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId pulumi.StringOutput `pulumi:"sourceSnapshotId"`
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users pulumi.StringArrayOutput `pulumi:"users"`
	// A reference to the zone where the disk resides.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil {
		args = &DiskArgs{}
	}
	var resource Disk
	err := ctx.RegisterResource("gcp:compute/disk:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisk gets an existing Disk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("gcp:compute/disk:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Disk resources.
type diskState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey *computeDiskDiskEncryptionKey.DiskDiskEncryptionKey `pulumi:"diskEncryptionKey"`
	// The image from which to initialize this disk. This can be one of: the image's 'self_link',
	// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}', 'global/images/{image}',
	// 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}', '{family}', or '{image}'.
	// If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data
	// source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image 'centos-6-v20180104' includes
	// its family name 'centos-6'. These images can be referred by family name here.
	Image *string `pulumi:"image"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this disk. A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp *string `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp *string `pulumi:"lastDetachTimestamp"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes *int `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies []string `pulumi:"resourcePolicies"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field along
	// with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
	Size *int `pulumi:"size"`
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot *string `pulumi:"snapshot"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a
	// customer-supplied encryption key.
	SourceImageEncryptionKey *computeDiskSourceImageEncryptionKey.DiskSourceImageEncryptionKey `pulumi:"sourceImageEncryptionKey"`
	// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
	// persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
	// under the same name, the source image ID would identify the exact version of the image that was used.
	SourceImageId *string `pulumi:"sourceImageId"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey *computeDiskSourceSnapshotEncryptionKey.DiskSourceSnapshotEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId *string `pulumi:"sourceSnapshotId"`
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.
	Type *string `pulumi:"type"`
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users []string `pulumi:"users"`
	// A reference to the zone where the disk resides.
	Zone *string `pulumi:"zone"`
}

type DiskState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey computeDiskDiskEncryptionKey.DiskDiskEncryptionKeyPtrInput
	// The image from which to initialize this disk. This can be one of: the image's 'self_link',
	// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}', 'global/images/{image}',
	// 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}', '{family}', or '{image}'.
	// If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data
	// source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image 'centos-6-v20180104' includes
	// its family name 'centos-6'. These images can be referred by family name here.
	Image pulumi.StringPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this disk. A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp pulumi.StringPtrInput
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies pulumi.StringArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field along
	// with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
	Size pulumi.IntPtrInput
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot pulumi.StringPtrInput
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a
	// customer-supplied encryption key.
	SourceImageEncryptionKey computeDiskSourceImageEncryptionKey.DiskSourceImageEncryptionKeyPtrInput
	// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
	// persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
	// under the same name, the source image ID would identify the exact version of the image that was used.
	SourceImageId pulumi.StringPtrInput
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey computeDiskSourceSnapshotEncryptionKey.DiskSourceSnapshotEncryptionKeyPtrInput
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
	// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId pulumi.StringPtrInput
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrInput
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users pulumi.StringArrayInput
	// A reference to the zone where the disk resides.
	Zone pulumi.StringPtrInput
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey *computeDiskDiskEncryptionKey.DiskDiskEncryptionKey `pulumi:"diskEncryptionKey"`
	// The image from which to initialize this disk. This can be one of: the image's 'self_link',
	// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}', 'global/images/{image}',
	// 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}', '{family}', or '{image}'.
	// If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data
	// source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image 'centos-6-v20180104' includes
	// its family name 'centos-6'. These images can be referred by family name here.
	Image *string `pulumi:"image"`
	// Labels to apply to this disk. A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes *int `pulumi:"physicalBlockSizeBytes"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies []string `pulumi:"resourcePolicies"`
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field along
	// with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
	Size *int `pulumi:"size"`
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot *string `pulumi:"snapshot"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a
	// customer-supplied encryption key.
	SourceImageEncryptionKey *computeDiskSourceImageEncryptionKey.DiskSourceImageEncryptionKey `pulumi:"sourceImageEncryptionKey"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey *computeDiskSourceSnapshotEncryptionKey.DiskSourceSnapshotEncryptionKey `pulumi:"sourceSnapshotEncryptionKey"`
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.
	Type *string `pulumi:"type"`
	// A reference to the zone where the disk resides.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey computeDiskDiskEncryptionKey.DiskDiskEncryptionKeyPtrInput
	// The image from which to initialize this disk. This can be one of: the image's 'self_link',
	// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}', 'global/images/{image}',
	// 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}', '{family}', or '{image}'.
	// If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data
	// source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image 'centos-6-v20180104' includes
	// its family name 'centos-6'. These images can be referred by family name here.
	Image pulumi.StringPtrInput
	// Labels to apply to this disk. A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies pulumi.StringArrayInput
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field along
	// with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
	Size pulumi.IntPtrInput
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot pulumi.StringPtrInput
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a
	// customer-supplied encryption key.
	SourceImageEncryptionKey computeDiskSourceImageEncryptionKey.DiskSourceImageEncryptionKeyPtrInput
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey computeDiskSourceSnapshotEncryptionKey.DiskSourceSnapshotEncryptionKeyPtrInput
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.
	Type pulumi.StringPtrInput
	// A reference to the zone where the disk resides.
	Zone pulumi.StringPtrInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}

