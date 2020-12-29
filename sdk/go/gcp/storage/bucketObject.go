// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new object inside an existing bucket in Google cloud storage service (GCS).
// [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied using the `storage.ObjectACL` resource.
//  For more information see
// [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
//
// ## Example Usage
//
// Example creating a public object in an existing `image-store` bucket.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.NewBucketObject(ctx, "picture", &storage.BucketObjectArgs{
// 			Bucket: pulumi.String("image-store"),
// 			Source: pulumi.NewFileAsset("/images/nature/garden-tiger-moth.jpg"),
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
// This resource does not support import.
type BucketObject struct {
	pulumi.CustomResourceState

	// The name of the containing bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrOutput `pulumi:"cacheControl"`
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition pulumi.StringPtrOutput `pulumi:"contentDisposition"`
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding pulumi.StringPtrOutput `pulumi:"contentEncoding"`
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage pulumi.StringPtrOutput `pulumi:"contentLanguage"`
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c        pulumi.StringOutput    `pulumi:"crc32c"`
	DetectMd5hash pulumi.StringPtrOutput `pulumi:"detectMd5hash"`
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash pulumi.StringOutput `pulumi:"md5hash"`
	// (Computed) A url reference to download this object.
	MediaLink pulumi.StringOutput `pulumi:"mediaLink"`
	// User-provided metadata, in key/value pairs.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Computed) The name of the object. Use this field in interpolations with `storage.ObjectACL` to recreate
	// `storage.ObjectACL` resources when your `storage.BucketObject` is recreated.
	OutputName pulumi.StringOutput `pulumi:"outputName"`
	// (Computed) A url reference to this object.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	Source pulumi.AssetOrArchiveOutput `pulumi:"source"`
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
}

// NewBucketObject registers a new resource with the given unique name, arguments, and options.
func NewBucketObject(ctx *pulumi.Context,
	name string, args *BucketObjectArgs, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource BucketObject
	err := ctx.RegisterResource("gcp:storage/bucketObject:BucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObject gets an existing BucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectState, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	var resource BucketObject
	err := ctx.ReadResource("gcp:storage/bucketObject:BucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObject resources.
type bucketObjectState struct {
	// The name of the containing bucket.
	Bucket *string `pulumi:"bucket"`
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl *string `pulumi:"cacheControl"`
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content *string `pulumi:"content"`
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType *string `pulumi:"contentType"`
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c        *string `pulumi:"crc32c"`
	DetectMd5hash *string `pulumi:"detectMd5hash"`
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash *string `pulumi:"md5hash"`
	// (Computed) A url reference to download this object.
	MediaLink *string `pulumi:"mediaLink"`
	// User-provided metadata, in key/value pairs.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name *string `pulumi:"name"`
	// (Computed) The name of the object. Use this field in interpolations with `storage.ObjectACL` to recreate
	// `storage.ObjectACL` resources when your `storage.BucketObject` is recreated.
	OutputName *string `pulumi:"outputName"`
	// (Computed) A url reference to this object.
	SelfLink *string `pulumi:"selfLink"`
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass *string `pulumi:"storageClass"`
}

type BucketObjectState struct {
	// The name of the containing bucket.
	Bucket pulumi.StringPtrInput
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrInput
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content pulumi.StringPtrInput
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition pulumi.StringPtrInput
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding pulumi.StringPtrInput
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage pulumi.StringPtrInput
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringPtrInput
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c        pulumi.StringPtrInput
	DetectMd5hash pulumi.StringPtrInput
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash pulumi.StringPtrInput
	// (Computed) A url reference to download this object.
	MediaLink pulumi.StringPtrInput
	// User-provided metadata, in key/value pairs.
	Metadata pulumi.StringMapInput
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name pulumi.StringPtrInput
	// (Computed) The name of the object. Use this field in interpolations with `storage.ObjectACL` to recreate
	// `storage.ObjectACL` resources when your `storage.BucketObject` is recreated.
	OutputName pulumi.StringPtrInput
	// (Computed) A url reference to this object.
	SelfLink pulumi.StringPtrInput
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	Source pulumi.AssetOrArchiveInput
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass pulumi.StringPtrInput
}

func (BucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectState)(nil)).Elem()
}

type bucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket string `pulumi:"bucket"`
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl *string `pulumi:"cacheControl"`
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content *string `pulumi:"content"`
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType   *string `pulumi:"contentType"`
	DetectMd5hash *string `pulumi:"detectMd5hash"`
	// User-provided metadata, in key/value pairs.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name *string `pulumi:"name"`
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass *string `pulumi:"storageClass"`
}

// The set of arguments for constructing a BucketObject resource.
type BucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket pulumi.StringInput
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrInput
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content pulumi.StringPtrInput
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition pulumi.StringPtrInput
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding pulumi.StringPtrInput
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage pulumi.StringPtrInput
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType   pulumi.StringPtrInput
	DetectMd5hash pulumi.StringPtrInput
	// User-provided metadata, in key/value pairs.
	Metadata pulumi.StringMapInput
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name pulumi.StringPtrInput
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	Source pulumi.AssetOrArchiveInput
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass pulumi.StringPtrInput
}

func (BucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectArgs)(nil)).Elem()
}

type BucketObjectInput interface {
	pulumi.Input

	ToBucketObjectOutput() BucketObjectOutput
	ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput
}

func (*BucketObject) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketObject)(nil))
}

func (i *BucketObject) ToBucketObjectOutput() BucketObjectOutput {
	return i.ToBucketObjectOutputWithContext(context.Background())
}

func (i *BucketObject) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectOutput)
}

type BucketObjectOutput struct {
	*pulumi.OutputState
}

func (BucketObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketObject)(nil))
}

func (o BucketObjectOutput) ToBucketObjectOutput() BucketObjectOutput {
	return o
}

func (o BucketObjectOutput) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketObjectOutput{})
}
