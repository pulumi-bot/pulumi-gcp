// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
	Filter *string `pulumi:"filter"`
	// The name of the image.
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
	Family string  `pulumi:"family"`
	Filter *string `pulumi:"filter"`
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

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			return *r, err
		}).(LookupImageResultOutput)
}

// A collection of arguments for invoking getImage.
type LookupImageOutputArgs struct {
	// The family name of the image.
	Family pulumi.StringPtrInput `pulumi:"family"`
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The name of the image.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used. If you are using a
	// [public base image][pubimg], be sure to specify the correct Image Project.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

// A collection of values returned by getImage.
type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

// The size of the image tar.gz archive stored in Google Cloud Storage in bytes.
func (o LookupImageResultOutput) ArchiveSizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.ArchiveSizeBytes }).(pulumi.IntOutput)
}

// The creation timestamp in RFC3339 text format.
func (o LookupImageResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this image.
func (o LookupImageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Description }).(pulumi.StringOutput)
}

// The size of the image when restored onto a persistent disk in gigabytes.
func (o LookupImageResultOutput) DiskSizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.DiskSizeGb }).(pulumi.IntOutput)
}

// The family name of the image.
func (o LookupImageResultOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Family }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
// encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
// that protects this image.
func (o LookupImageResultOutput) ImageEncryptionKeySha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ImageEncryptionKeySha256 }).(pulumi.StringOutput)
}

// The unique identifier for the image.
func (o LookupImageResultOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ImageId }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this image.
func (o LookupImageResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// A map of labels applied to this image.
func (o LookupImageResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// A list of applicable license URI.
func (o LookupImageResultOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.Licenses }).(pulumi.StringArrayOutput)
}

// The name of the image.
func (o LookupImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Project }).(pulumi.StringOutput)
}

// The URI of the image.
func (o LookupImageResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The URL of the source disk used to create this image.
func (o LookupImageResultOutput) SourceDisk() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceDisk }).(pulumi.StringOutput)
}

// The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
// encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
// that protects this image.
func (o LookupImageResultOutput) SourceDiskEncryptionKeySha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceDiskEncryptionKeySha256 }).(pulumi.StringOutput)
}

// The ID value of the disk used to create this image.
func (o LookupImageResultOutput) SourceDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceDiskId }).(pulumi.StringOutput)
}

// The ID value of the image used to create this image.
func (o LookupImageResultOutput) SourceImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceImageId }).(pulumi.StringOutput)
}

// The status of the image. Possible values are **FAILED**, **PENDING**, or **READY**.
func (o LookupImageResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
