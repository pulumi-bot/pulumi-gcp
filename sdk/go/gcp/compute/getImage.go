// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information about a Google Compute Image. Check that your service account has the `compute.imageUser` role if you want to share [custom images](https://cloud.google.com/compute/docs/images/sharing-images-across-projects) from another project. If you want to use [public images][pubimg], do not forget to specify the dedicated project. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/images) and its [API](https://cloud.google.com/compute/docs/reference/latest/images).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "debian-9"
// 		opt1 := "debian-cloud"
// 		myImage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Family:  &opt0,
// 			Project: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewInstance(ctx, "_default", &compute.InstanceArgs{
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String(myImage.SelfLink),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	var rv LookupImageResult
	err := ctx.Invoke("gcp:compute/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type LookupImageArgs struct {
	// The family name of the image.
	Family *string `pulumi:"family"`
	// or `family` - (Required) The name of a specific image or a family.
	// Exactly one of `name` of `family` must be specified. If `name` is specified, it will fetch
	// the corresponding image. If `family` is specified, it will returns the latest image
	// that is part of an image family and is not deprecated.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used. If you are using a
	// [public base image][pubimg], be sure to specify the correct Image Project.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getImage.
type LookupImageResult struct {
	// The size of the image tar.gz archive stored in Google Cloud Storage in bytes.
	ArchiveSizeBytes int `pulumi:"archiveSizeBytes"`
	// The creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this image.
	Description string `pulumi:"description"`
	// The size of the image when restored onto a persistent disk in gigabytes.
	DiskSizeGb int `pulumi:"diskSizeGb"`
	// The family name of the image.
	Family string `pulumi:"family"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
	// encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
	// that protects this image.
	ImageEncryptionKeySha256 string `pulumi:"imageEncryptionKeySha256"`
	// The unique identifier for the image.
	ImageId string `pulumi:"imageId"`
	// A fingerprint for the labels being applied to this image.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// A map of labels applied to this image.
	Labels map[string]string `pulumi:"labels"`
	// A list of applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// The name of the image.
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	// The URI of the image.
	SelfLink string `pulumi:"selfLink"`
	// The URL of the source disk used to create this image.
	SourceDisk string `pulumi:"sourceDisk"`
	// The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
	// encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
	// that protects this image.
	SourceDiskEncryptionKeySha256 string `pulumi:"sourceDiskEncryptionKeySha256"`
	// The ID value of the disk used to create this image.
	SourceDiskId string `pulumi:"sourceDiskId"`
	// The ID value of the image used to create this image.
	SourceImageId string `pulumi:"sourceImageId"`
	// The status of the image. Possible values are **FAILED**, **PENDING**, or **READY**.
	Status string `pulumi:"status"`
}
