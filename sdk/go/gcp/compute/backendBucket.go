// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S)
// load balancing.
//
// An HTTP(S) load balancer can direct traffic to specified URLs to a
// backend bucket rather than a backend service. It can send requests for
// static content to a Cloud Storage bucket and requests for dynamic content
// to a virtual machine instance.
//
//
// To get more information about BackendBucket, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendBuckets)
// * How-to Guides
//     * [Using a Cloud Storage bucket as a load balancer backend](https://cloud.google.com/compute/docs/load-balancing/http/backend-bucket)
type BackendBucket struct {
	pulumi.CustomResourceState

	// Cloud Storage bucket name.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// Cloud CDN configuration for this Backend Bucket.
	CdnPolicy BackendBucketCdnPolicyOutput `pulumi:"cdnPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn pulumi.BoolPtrOutput `pulumi:"enableCdn"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewBackendBucket registers a new resource with the given unique name, arguments, and options.
func NewBackendBucket(ctx *pulumi.Context,
	name string, args *BackendBucketArgs, opts ...pulumi.ResourceOption) (*BackendBucket, error) {
	if args == nil || args.BucketName == nil {
		return nil, errors.New("missing required argument 'BucketName'")
	}
	if args == nil {
		args = &BackendBucketArgs{}
	}
	var resource BackendBucket
	err := ctx.RegisterResource("gcp:compute/backendBucket:BackendBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendBucket gets an existing BackendBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendBucketState, opts ...pulumi.ResourceOption) (*BackendBucket, error) {
	var resource BackendBucket
	err := ctx.ReadResource("gcp:compute/backendBucket:BackendBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendBucket resources.
type backendBucketState struct {
	// Cloud Storage bucket name.
	BucketName *string `pulumi:"bucketName"`
	// Cloud CDN configuration for this Backend Bucket.
	CdnPolicy *BackendBucketCdnPolicy `pulumi:"cdnPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn *bool `pulumi:"enableCdn"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type BackendBucketState struct {
	// Cloud Storage bucket name.
	BucketName pulumi.StringPtrInput
	// Cloud CDN configuration for this Backend Bucket.
	CdnPolicy BackendBucketCdnPolicyPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (BackendBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketState)(nil)).Elem()
}

type backendBucketArgs struct {
	// Cloud Storage bucket name.
	BucketName string `pulumi:"bucketName"`
	// Cloud CDN configuration for this Backend Bucket.
	CdnPolicy *BackendBucketCdnPolicyArgs `pulumi:"cdnPolicy"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn *bool `pulumi:"enableCdn"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackendBucket resource.
type BackendBucketArgs struct {
	// Cloud Storage bucket name.
	BucketName pulumi.StringInput
	// Cloud CDN configuration for this Backend Bucket.
	CdnPolicy BackendBucketCdnPolicyArgsPtrInput
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketArgs)(nil)).Elem()
}
