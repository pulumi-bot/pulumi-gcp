// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an Image resource.
//
// Google Compute Engine uses operating system images to create the root
// persistent disks for your instances. You specify an image when you create
// an instance. Images contain a boot loader, an operating system, and a
// root file system. Linux operating system images are also capable of
// running containers on Compute Engine.
//
// Images can be either public or custom.
//
// Public images are provided and maintained by Google, open-source
// communities, and third-party vendors. By default, all projects have
// access to these images and can use them to create instances.  Custom
// images are available only to your project. You can create a custom image
// from root persistent disks and other images. Then, use the custom image
// to create an instance.
//
// To get more information about Image, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/images)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/images)
//
// ## Example Usage
type Image struct {
	pulumi.CustomResourceState

	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes pulumi.IntOutput `pulumi:"archiveSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.IntOutput `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family pulumi.StringPtrOutput `pulumi:"family"`
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures ImageGuestOsFeatureArrayOutput `pulumi:"guestOsFeatures"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this Image.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Any applicable license URI.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk ImageRawDiskPtrOutput `pulumi:"rawDisk"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk pulumi.StringPtrOutput `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage pulumi.StringPtrOutput `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot pulumi.StringPtrOutput `pulumi:"sourceSnapshot"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		args = &ImageArgs{}
	}

	var resource Image
	err := ctx.RegisterResource("gcp:compute/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("gcp:compute/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes *int `pulumi:"archiveSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family *string `pulumi:"family"`
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures []ImageGuestOsFeature `pulumi:"guestOsFeatures"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this Image.
	Labels map[string]string `pulumi:"labels"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk *ImageRawDisk `pulumi:"rawDisk"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk *string `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
}

type ImageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.IntPtrInput
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family pulumi.StringPtrInput
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures ImageGuestOsFeatureArrayInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this Image.
	Labels pulumi.StringMapInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk ImageRawDiskPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk pulumi.StringPtrInput
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family *string `pulumi:"family"`
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures []ImageGuestOsFeature `pulumi:"guestOsFeatures"`
	// Labels to apply to this Image.
	Labels map[string]string `pulumi:"labels"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk *ImageRawDisk `pulumi:"rawDisk"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk *string `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.IntPtrInput
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	Family pulumi.StringPtrInput
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	GuestOsFeatures ImageGuestOsFeatureArrayInput
	// Labels to apply to this Image.
	Labels pulumi.StringMapInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The parameters of the raw disk image.
	// Structure is documented below.
	RawDisk ImageRawDiskPtrInput
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	SourceDisk pulumi.StringPtrInput
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (Image) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil)).Elem()
}

func (i Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct {
	*pulumi.OutputState
}

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOutput)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
}
