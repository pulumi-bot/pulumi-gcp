// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
//
// ## Example Usage
//
// Example picture stored within a folder.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.LookupBucketObject(ctx, &storage.LookupBucketObjectArgs{
// 			Bucket: "image-store",
// 			Name:   "folder/butterfly01.jpg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupBucketObject(ctx *pulumi.Context, args *LookupBucketObjectArgs, opts ...pulumi.InvokeOption) (*LookupBucketObjectResult, error) {
	var rv LookupBucketObjectResult
	err := ctx.Invoke("gcp:storage/getBucketObject:getBucketObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketObject.
type LookupBucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket *string `pulumi:"bucket"`
	// The name of the object.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getBucketObject.
type LookupBucketObjectResult struct {
	Bucket *string `pulumi:"bucket"`
	// (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl string `pulumi:"cacheControl"`
	Content      string `pulumi:"content"`
	// (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition string `pulumi:"contentDisposition"`
	// (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding string `pulumi:"contentEncoding"`
	// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage string `pulumi:"contentLanguage"`
	// (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType string `pulumi:"contentType"`
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c        string `pulumi:"crc32c"`
	DetectMd5hash string `pulumi:"detectMd5hash"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash    string            `pulumi:"md5hash"`
	Metadata   map[string]string `pulumi:"metadata"`
	Name       *string           `pulumi:"name"`
	OutputName string            `pulumi:"outputName"`
	// (Computed) A url reference to this object.
	SelfLink string `pulumi:"selfLink"`
	Source   string `pulumi:"source"`
	// (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass string `pulumi:"storageClass"`
}
