// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a billing account level logging bucket config. For more information see
// [the official logging documentation](https://cloud.google.com/logging/docs/) and
// [Storing Logs](https://cloud.google.com/logging/docs/storage).
//
// > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/logging"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "00AA00-000AAA-00AA0A"
// 		_default, err := organizations.GetBillingAccount(ctx, &organizations.GetBillingAccountArgs{
// 			BillingAccount: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewBillingAccountBucketConfig(ctx, "basic", &logging.BillingAccountBucketConfigArgs{
// 			BillingAccount: pulumi.String(_default.BillingAccount),
// 			Location:       pulumi.String("global"),
// 			RetentionDays:  pulumi.Int(30),
// 			BucketId:       pulumi.String("_Default"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type BillingAccountBucketConfig struct {
	pulumi.CustomResourceState

	// The parent resource that contains the logging bucket.
	BillingAccount pulumi.StringOutput `pulumi:"billingAccount"`
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringOutput `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewBillingAccountBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountBucketConfig(ctx *pulumi.Context,
	name string, args *BillingAccountBucketConfigArgs, opts ...pulumi.ResourceOption) (*BillingAccountBucketConfig, error) {
	if args == nil || args.BillingAccount == nil {
		return nil, errors.New("missing required argument 'BillingAccount'")
	}
	if args == nil || args.BucketId == nil {
		return nil, errors.New("missing required argument 'BucketId'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &BillingAccountBucketConfigArgs{}
	}
	var resource BillingAccountBucketConfig
	err := ctx.RegisterResource("gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountBucketConfig gets an existing BillingAccountBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountBucketConfigState, opts ...pulumi.ResourceOption) (*BillingAccountBucketConfig, error) {
	var resource BillingAccountBucketConfig
	err := ctx.ReadResource("gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountBucketConfig resources.
type billingAccountBucketConfigState struct {
	// The parent resource that contains the logging bucket.
	BillingAccount *string `pulumi:"billingAccount"`
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location *string `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name *string `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

type BillingAccountBucketConfigState struct {
	// The parent resource that contains the logging bucket.
	BillingAccount pulumi.StringPtrInput
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringPtrInput
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (BillingAccountBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketConfigState)(nil)).Elem()
}

type billingAccountBucketConfigArgs struct {
	// The parent resource that contains the logging bucket.
	BillingAccount string `pulumi:"billingAccount"`
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location string `pulumi:"location"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a BillingAccountBucketConfig resource.
type BillingAccountBucketConfigArgs struct {
	// The parent resource that contains the logging bucket.
	BillingAccount pulumi.StringInput
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (BillingAccountBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketConfigArgs)(nil)).Elem()
}
