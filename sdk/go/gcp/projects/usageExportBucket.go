// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows creation and management of a Google Cloud Platform project.
//
// Projects created with this resource must be associated with an Organization.
// See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
//
// The service account used to run this provider when creating a `organizations.Project`
// resource must have `roles/resourcemanager.projectCreator`. See the
// [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
// doc for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewProject(ctx, "myProject", &organizations.ProjectArgs{
// 			OrgId:     pulumi.String("1234567"),
// 			ProjectId: pulumi.String("your-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To create a project under a specific folder
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		department1, err := organizations.NewFolder(ctx, "department1", &organizations.FolderArgs{
// 			DisplayName: pulumi.String("Department 1"),
// 			Parent:      pulumi.String("organizations/1234567"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = organizations.NewProject(ctx, "myProject_in_a_folder", &organizations.ProjectArgs{
// 			ProjectId: pulumi.String("your-project-id"),
// 			FolderId:  department1.Name,
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
// Projects can be imported using the `project_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
// ```
type UsageExportBucket struct {
	pulumi.CustomResourceState

	// The bucket to store reports in.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// A prefix for the reports, for instance, the project name.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewUsageExportBucket registers a new resource with the given unique name, arguments, and options.
func NewUsageExportBucket(ctx *pulumi.Context,
	name string, args *UsageExportBucketArgs, opts ...pulumi.ResourceOption) (*UsageExportBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	var resource UsageExportBucket
	err := ctx.RegisterResource("gcp:projects/usageExportBucket:UsageExportBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUsageExportBucket gets an existing UsageExportBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUsageExportBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UsageExportBucketState, opts ...pulumi.ResourceOption) (*UsageExportBucket, error) {
	var resource UsageExportBucket
	err := ctx.ReadResource("gcp:projects/usageExportBucket:UsageExportBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UsageExportBucket resources.
type usageExportBucketState struct {
	// The bucket to store reports in.
	BucketName *string `pulumi:"bucketName"`
	// A prefix for the reports, for instance, the project name.
	Prefix *string `pulumi:"prefix"`
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type UsageExportBucketState struct {
	// The bucket to store reports in.
	BucketName pulumi.StringPtrInput
	// A prefix for the reports, for instance, the project name.
	Prefix pulumi.StringPtrInput
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (UsageExportBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*usageExportBucketState)(nil)).Elem()
}

type usageExportBucketArgs struct {
	// The bucket to store reports in.
	BucketName string `pulumi:"bucketName"`
	// A prefix for the reports, for instance, the project name.
	Prefix *string `pulumi:"prefix"`
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a UsageExportBucket resource.
type UsageExportBucketArgs struct {
	// The bucket to store reports in.
	BucketName pulumi.StringInput
	// A prefix for the reports, for instance, the project name.
	Prefix pulumi.StringPtrInput
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (UsageExportBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usageExportBucketArgs)(nil)).Elem()
}

type UsageExportBucketInput interface {
	pulumi.Input

	ToUsageExportBucketOutput() UsageExportBucketOutput
	ToUsageExportBucketOutputWithContext(ctx context.Context) UsageExportBucketOutput
}

func (*UsageExportBucket) ElementType() reflect.Type {
	return reflect.TypeOf((*UsageExportBucket)(nil))
}

func (i *UsageExportBucket) ToUsageExportBucketOutput() UsageExportBucketOutput {
	return i.ToUsageExportBucketOutputWithContext(context.Background())
}

func (i *UsageExportBucket) ToUsageExportBucketOutputWithContext(ctx context.Context) UsageExportBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsageExportBucketOutput)
}

type UsageExportBucketOutput struct {
	*pulumi.OutputState
}

func (UsageExportBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsageExportBucket)(nil))
}

func (o UsageExportBucketOutput) ToUsageExportBucketOutput() UsageExportBucketOutput {
	return o
}

func (o UsageExportBucketOutput) ToUsageExportBucketOutputWithContext(ctx context.Context) UsageExportBucketOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UsageExportBucketOutput{})
}
