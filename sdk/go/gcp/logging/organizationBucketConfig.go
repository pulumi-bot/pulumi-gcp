// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type OrganizationBucketConfig struct {
	pulumi.CustomResourceState

	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewOrganizationBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewOrganizationBucketConfig(ctx *pulumi.Context,
	name string, args *OrganizationBucketConfigArgs, opts ...pulumi.ResourceOption) (*OrganizationBucketConfig, error) {
	if args == nil || args.BucketId == nil {
		return nil, errors.New("missing required argument 'BucketId'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Organization == nil {
		return nil, errors.New("missing required argument 'Organization'")
	}
	if args == nil {
		args = &OrganizationBucketConfigArgs{}
	}
	var resource OrganizationBucketConfig
	err := ctx.RegisterResource("gcp:logging/organizationBucketConfig:OrganizationBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationBucketConfig gets an existing OrganizationBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationBucketConfigState, opts ...pulumi.ResourceOption) (*OrganizationBucketConfig, error) {
	var resource OrganizationBucketConfig
	err := ctx.ReadResource("gcp:logging/organizationBucketConfig:OrganizationBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationBucketConfig resources.
type organizationBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location *string `pulumi:"location"`
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name *string `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Organization *string `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

type OrganizationBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringPtrInput
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketConfigState)(nil)).Elem()
}

type organizationBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location string `pulumi:"location"`
	// The parent resource that contains the logging bucket.
	Organization string `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a OrganizationBucketConfig resource.
type OrganizationBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringInput
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketConfigArgs)(nil)).Elem()
}
