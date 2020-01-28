// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_backend_bucket_signed_url_key.html.markdown.
type BackendBucketSignedUrlKey struct {
	pulumi.CustomResourceState

	// The backend bucket this signed URL key belongs.
	BackendBucket pulumi.StringOutput `pulumi:"backendBucket"`
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue pulumi.StringOutput `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewBackendBucketSignedUrlKey registers a new resource with the given unique name, arguments, and options.
func NewBackendBucketSignedUrlKey(ctx *pulumi.Context,
	name string, args *BackendBucketSignedUrlKeyArgs, opts ...pulumi.ResourceOption) (*BackendBucketSignedUrlKey, error) {
	if args == nil || args.BackendBucket == nil {
		return nil, errors.New("missing required argument 'BackendBucket'")
	}
	if args == nil || args.KeyValue == nil {
		return nil, errors.New("missing required argument 'KeyValue'")
	}
	if args == nil {
		args = &BackendBucketSignedUrlKeyArgs{}
	}
	var resource BackendBucketSignedUrlKey
	err := ctx.RegisterResource("gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendBucketSignedUrlKey gets an existing BackendBucketSignedUrlKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendBucketSignedUrlKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendBucketSignedUrlKeyState, opts ...pulumi.ResourceOption) (*BackendBucketSignedUrlKey, error) {
	var resource BackendBucketSignedUrlKey
	err := ctx.ReadResource("gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendBucketSignedUrlKey resources.
type backendBucketSignedUrlKeyState struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket *string `pulumi:"backendBucket"`
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue *string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type BackendBucketSignedUrlKeyState struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket pulumi.StringPtrInput
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue pulumi.StringPtrInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketSignedUrlKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketSignedUrlKeyState)(nil)).Elem()
}

type backendBucketSignedUrlKeyArgs struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket string `pulumi:"backendBucket"`
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackendBucketSignedUrlKey resource.
type BackendBucketSignedUrlKeyArgs struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket pulumi.StringInput
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue pulumi.StringInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketSignedUrlKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketSignedUrlKeyArgs)(nil)).Elem()
}

