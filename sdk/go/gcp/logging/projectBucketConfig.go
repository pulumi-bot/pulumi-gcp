// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a project-level logging bucket config. For more information see
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/logging"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewProject(ctx, "_default", &organizations.ProjectArgs{
// 			ProjectId: pulumi.String("your-project-id"),
// 			OrgId:     pulumi.String("123456789"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewProjectBucketConfig(ctx, "basic", &logging.ProjectBucketConfigArgs{
// 			Project:       _default.Name,
// 			Location:      pulumi.String("global"),
// 			RetentionDays: pulumi.Int(30),
// 			BucketId:      pulumi.String("_Default"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Create logging bucket with customId
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/logging"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logging.NewProjectBucketConfig(ctx, "basic", &logging.ProjectBucketConfigArgs{
// 			BucketId:      pulumi.String("custom-bucket"),
// 			Location:      pulumi.String("global"),
// 			Project:       pulumi.String("project_id"),
// 			RetentionDays: pulumi.Int(30),
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
// This resource can be imported using the following format
//
// ```sh
//  $ pulumi import gcp:logging/projectBucketConfig:ProjectBucketConfig default projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
// ```
type ProjectBucketConfig struct {
	pulumi.CustomResourceState

	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Project pulumi.StringOutput `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewProjectBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewProjectBucketConfig(ctx *pulumi.Context,
	name string, args *ProjectBucketConfigArgs, opts ...pulumi.ResourceOption) (*ProjectBucketConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource ProjectBucketConfig
	err := ctx.RegisterResource("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectBucketConfig gets an existing ProjectBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectBucketConfigState, opts ...pulumi.ResourceOption) (*ProjectBucketConfig, error) {
	var resource ProjectBucketConfig
	err := ctx.ReadResource("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectBucketConfig resources.
type projectBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location *string `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name *string `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Project *string `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

type ProjectBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket.
	Location pulumi.StringPtrInput
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Project pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (ProjectBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBucketConfigState)(nil)).Elem()
}

type projectBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The location of the bucket.
	Location string `pulumi:"location"`
	// The parent resource that contains the logging bucket.
	Project string `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a ProjectBucketConfig resource.
type ProjectBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The location of the bucket.
	Location pulumi.StringInput
	// The parent resource that contains the logging bucket.
	Project pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (ProjectBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBucketConfigArgs)(nil)).Elem()
}

type ProjectBucketConfigInput interface {
	pulumi.Input

	ToProjectBucketConfigOutput() ProjectBucketConfigOutput
	ToProjectBucketConfigOutputWithContext(ctx context.Context) ProjectBucketConfigOutput
}

func (*ProjectBucketConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectBucketConfig)(nil))
}

func (i *ProjectBucketConfig) ToProjectBucketConfigOutput() ProjectBucketConfigOutput {
	return i.ToProjectBucketConfigOutputWithContext(context.Background())
}

func (i *ProjectBucketConfig) ToProjectBucketConfigOutputWithContext(ctx context.Context) ProjectBucketConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectBucketConfigOutput)
}

type ProjectBucketConfigOutput struct {
	*pulumi.OutputState
}

func (ProjectBucketConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectBucketConfig)(nil))
}

func (o ProjectBucketConfigOutput) ToProjectBucketConfigOutput() ProjectBucketConfigOutput {
	return o
}

func (o ProjectBucketConfigOutput) ToProjectBucketConfigOutputWithContext(ctx context.Context) ProjectBucketConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectBucketConfigOutput{})
}
