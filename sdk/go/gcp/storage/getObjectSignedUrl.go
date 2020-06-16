// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Google Cloud storage signed URL data source generates a signed URL for a given storage object. Signed URLs provide a way to give time-limited read or write access to anyone in possession of the URL, regardless of whether they have a Google account.
//
// For more info about signed URL's is available [here](https://cloud.google.com/storage/docs/access-control/signed-urls).
//
// ## Example Usage
//
//
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
// 		_, err := storage.LookupObjectSignedUrl(ctx, &storage.LookupObjectSignedUrlArgs{
// 			Bucket: "install_binaries",
// 			Path:   "path/to/install_file.bin",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewInstance(ctx, "vm", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetObjectSignedUrl(ctx *pulumi.Context, args *GetObjectSignedUrlArgs, opts ...pulumi.InvokeOption) (*GetObjectSignedUrlResult, error) {
	var rv GetObjectSignedUrlResult
	err := ctx.Invoke("gcp:storage/getObjectSignedUrl:getObjectSignedUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjectSignedUrl.
type GetObjectSignedUrlArgs struct {
	// The name of the bucket to read the object from
	Bucket string `pulumi:"bucket"`
	// The [MD5 digest](https://cloud.google.com/storage/docs/hashes-etags#_MD5) value in Base64.
	// Typically retrieved from `google_storage_bucket_object.object.md5hash` attribute.
	// If you provide this in the datasource, the client (e.g. browser, curl) must provide the `Content-MD5` HTTP header with this same value in its request.
	ContentMd5 *string `pulumi:"contentMd5"`
	// If you specify this in the datasource, the client must provide the `Content-Type` HTTP header with the same value in its request.
	ContentType *string `pulumi:"contentType"`
	// What Google service account credentials json should be used to sign the URL.
	// This data source checks the following locations for credentials, in order of preference: data source `credentials` attribute, provider `credentials` attribute and finally the GOOGLE_APPLICATION_CREDENTIALS environment variable.
	Credentials *string `pulumi:"credentials"`
	// For how long shall the signed URL be valid (defaults to 1 hour - i.e. `1h`).
	// See [here](https://golang.org/pkg/time/#ParseDuration) for info on valid duration formats.
	Duration *string `pulumi:"duration"`
	// As needed. The server checks to make sure that the client provides matching values in requests using the signed URL.
	// Any header starting with `x-goog-` is accepted but see the [Google Docs](https://cloud.google.com/storage/docs/xml-api/reference-headers) for list of headers that are supported by Google.
	ExtensionHeaders map[string]string `pulumi:"extensionHeaders"`
	// What HTTP Method will the signed URL allow (defaults to `GET`)
	HttpMethod *string `pulumi:"httpMethod"`
	// The full path to the object inside the bucket
	Path string `pulumi:"path"`
}

// A collection of values returned by getObjectSignedUrl.
type GetObjectSignedUrlResult struct {
	Bucket           string            `pulumi:"bucket"`
	ContentMd5       *string           `pulumi:"contentMd5"`
	ContentType      *string           `pulumi:"contentType"`
	Credentials      *string           `pulumi:"credentials"`
	Duration         *string           `pulumi:"duration"`
	ExtensionHeaders map[string]string `pulumi:"extensionHeaders"`
	HttpMethod       *string           `pulumi:"httpMethod"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Path string `pulumi:"path"`
	// The signed URL that can be used to access the storage object without authentication.
	SignedUrl string `pulumi:"signedUrl"`
}
