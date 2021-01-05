// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new bucket in Google cloud storage service (GCS).
// Once a bucket has been created, its location can't be changed.
// [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied
// using the [`storage.BucketACL`](https://www.terraform.io/docs/providers/google/r/storage_bucket_acl.html) resource.
//
// For more information see
// [the official documentation](https://cloud.google.com/storage/docs/overview)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
//
// **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
// determined which will require enabling the compute api.
//
// ## Example Usage
// ### Creating A Private Bucket In Standard Storage, In The EU Region. Bucket Configured As Static Website And CORS Configurations
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
// 		_, err := storage.NewBucket(ctx, "static-site", &storage.BucketArgs{
// 			Cors: storage.BucketCorArray{
// 				&storage.BucketCorArgs{
// 					MaxAgeSeconds: pulumi.Int(3600),
// 					Methods: pulumi.StringArray{
// 						pulumi.String("GET"),
// 						pulumi.String("HEAD"),
// 						pulumi.String("PUT"),
// 						pulumi.String("POST"),
// 						pulumi.String("DELETE"),
// 					},
// 					Origins: pulumi.StringArray{
// 						pulumi.String("http://image-store.com"),
// 					},
// 					ResponseHeaders: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 				},
// 			},
// 			ForceDestroy:             pulumi.Bool(true),
// 			Location:                 pulumi.String("EU"),
// 			UniformBucketLevelAccess: pulumi.Bool(true),
// 			Website: &storage.BucketWebsiteArgs{
// 				MainPageSuffix: pulumi.String("index.html"),
// 				NotFoundPage:   pulumi.String("404.html"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Life Cycle Settings For Storage Bucket Objects
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
// 		_, err := storage.NewBucket(ctx, "auto-expire", &storage.BucketArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 			LifecycleRules: storage.BucketLifecycleRuleArray{
// 				&storage.BucketLifecycleRuleArgs{
// 					Action: &storage.BucketLifecycleRuleActionArgs{
// 						Type: pulumi.String("Delete"),
// 					},
// 					Condition: &storage.BucketLifecycleRuleConditionArgs{
// 						Age: pulumi.Int(3),
// 					},
// 				},
// 			},
// 			Location: pulumi.String("US"),
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
// Storage buckets can be imported using the `name` or
//
// `project/name`. If the project is not passed to the import command it will be inferred from the provider block or environment variables. If it cannot be inferred it will be queried from the Compute API (this will fail if the API is not enabled). e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucket:Bucket image-store image-store-bucket
// ```
//
// ```sh
//  $ pulumi import gcp:storage/bucket:Bucket image-store tf-test-project/image-store-bucket
// ```
//
//  `false` in state. If you've set it to `true` in config, run `terraform apply` to update the value set in state. If you delete this resource before updating the value, objects in the bucket will not be destroyed.
type Bucket struct {
	pulumi.CustomResourceState

	// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
	//
	// Deprecated: Please use the uniform_bucket_level_access as this field has been renamed by Google.
	BucketPolicyOnly pulumi.BoolOutput `pulumi:"bucketPolicyOnly"`
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  BucketCorArrayOutput `pulumi:"cors"`
	DefaultEventBasedHold pulumi.BoolPtrOutput `pulumi:"defaultEventBasedHold"`
	// The bucket's encryption configuration.
	Encryption BucketEncryptionPtrOutput `pulumi:"encryption"`
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// A set of key/value label pairs to assign to the bucket.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
	Logging BucketLoggingPtrOutput `pulumi:"logging"`
	// The name of the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays pulumi.BoolPtrOutput `pulumi:"requesterPays"`
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy BucketRetentionPolicyPtrOutput `pulumi:"retentionPolicy"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess pulumi.BoolOutput `pulumi:"uniformBucketLevelAccess"`
	// The base URL of the bucket, in the format `gs://<bucket-name>`.
	Url pulumi.StringOutput `pulumi:"url"`
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
	Versioning BucketVersioningPtrOutput `pulumi:"versioning"`
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website BucketWebsitePtrOutput `pulumi:"website"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

	var resource Bucket
	err := ctx.RegisterResource("gcp:storage/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("gcp:storage/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
	//
	// Deprecated: Please use the uniform_bucket_level_access as this field has been renamed by Google.
	BucketPolicyOnly *bool `pulumi:"bucketPolicyOnly"`
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  []BucketCor `pulumi:"cors"`
	DefaultEventBasedHold *bool       `pulumi:"defaultEventBasedHold"`
	// The bucket's encryption configuration.
	Encryption *BucketEncryption `pulumi:"encryption"`
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A set of key/value label pairs to assign to the bucket.
	Labels map[string]string `pulumi:"labels"`
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location *string `pulumi:"location"`
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
	Logging *BucketLogging `pulumi:"logging"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays *bool `pulumi:"requesterPays"`
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy *BucketRetentionPolicy `pulumi:"retentionPolicy"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass *string `pulumi:"storageClass"`
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess *bool `pulumi:"uniformBucketLevelAccess"`
	// The base URL of the bucket, in the format `gs://<bucket-name>`.
	Url *string `pulumi:"url"`
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
	Versioning *BucketVersioning `pulumi:"versioning"`
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website *BucketWebsite `pulumi:"website"`
}

type BucketState struct {
	// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
	//
	// Deprecated: Please use the uniform_bucket_level_access as this field has been renamed by Google.
	BucketPolicyOnly pulumi.BoolPtrInput
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  BucketCorArrayInput
	DefaultEventBasedHold pulumi.BoolPtrInput
	// The bucket's encryption configuration.
	Encryption BucketEncryptionPtrInput
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy pulumi.BoolPtrInput
	// A set of key/value label pairs to assign to the bucket.
	Labels pulumi.StringMapInput
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules BucketLifecycleRuleArrayInput
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location pulumi.StringPtrInput
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
	Logging BucketLoggingPtrInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays pulumi.BoolPtrInput
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy BucketRetentionPolicyPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass pulumi.StringPtrInput
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess pulumi.BoolPtrInput
	// The base URL of the bucket, in the format `gs://<bucket-name>`.
	Url pulumi.StringPtrInput
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
	Versioning BucketVersioningPtrInput
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website BucketWebsitePtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
	//
	// Deprecated: Please use the uniform_bucket_level_access as this field has been renamed by Google.
	BucketPolicyOnly *bool `pulumi:"bucketPolicyOnly"`
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  []BucketCor `pulumi:"cors"`
	DefaultEventBasedHold *bool       `pulumi:"defaultEventBasedHold"`
	// The bucket's encryption configuration.
	Encryption *BucketEncryption `pulumi:"encryption"`
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A set of key/value label pairs to assign to the bucket.
	Labels map[string]string `pulumi:"labels"`
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location *string `pulumi:"location"`
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
	Logging *BucketLogging `pulumi:"logging"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays *bool `pulumi:"requesterPays"`
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy *BucketRetentionPolicy `pulumi:"retentionPolicy"`
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass *string `pulumi:"storageClass"`
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess *bool `pulumi:"uniformBucketLevelAccess"`
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
	Versioning *BucketVersioning `pulumi:"versioning"`
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website *BucketWebsite `pulumi:"website"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
	//
	// Deprecated: Please use the uniform_bucket_level_access as this field has been renamed by Google.
	BucketPolicyOnly pulumi.BoolPtrInput
	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors                  BucketCorArrayInput
	DefaultEventBasedHold pulumi.BoolPtrInput
	// The bucket's encryption configuration.
	Encryption BucketEncryptionPtrInput
	// When deleting a bucket, this
	// boolean option will delete all contained objects. If you try to delete a
	// bucket that contains objects, the provider will fail that run.
	ForceDestroy pulumi.BoolPtrInput
	// A set of key/value label pairs to assign to the bucket.
	Labels pulumi.StringMapInput
	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules BucketLifecycleRuleArrayInput
	// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
	Location pulumi.StringPtrInput
	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
	Logging BucketLoggingPtrInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays pulumi.BoolPtrInput
	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy BucketRetentionPolicyPtrInput
	// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass pulumi.StringPtrInput
	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess pulumi.BoolPtrInput
	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
	Versioning BucketVersioningPtrInput
	// Configuration if the bucket acts as a website. Structure is documented below.
	Website BucketWebsitePtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

type BucketOutput struct {
	*pulumi.OutputState
}

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketOutput{})
}
