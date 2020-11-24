// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for BigQuery dataset. Each of these resources serves a different use case:
//
// * `bigquery.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
// * `bigquery.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
// * `bigquery.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
//
// These resources are intended to convert the permissions system for BigQuery datasets to the standard IAM interface. For advanced usages, including [creating authorized views](https://cloud.google.com/bigquery/docs/share-access-views), please use either `bigquery.DatasetAccess` or the `access` field on `bigquery.Dataset`.
//
// > **Note:** These resources **cannot** be used with `bigquery.DatasetAccess` resources or the `access` field on `bigquery.Dataset` or they will fight over what the policy should be.
//
// > **Note:** Using any of these resources will remove any authorized view permissions from the dataset. To assign and preserve authorized view permissions use the `bigquery.DatasetAccess` instead.
//
// > **Note:** Legacy BigQuery roles `OWNER` `WRITER` and `READER` **cannot** be used with any of these IAM resources. Instead use the full role form of: `roles/bigquery.dataOwner` `roles/bigquery.dataEditor` and `roles/bigquery.dataViewer`.
//
// > **Note:** `bigquery.DatasetIamPolicy` **cannot** be used in conjunction with `bigquery.DatasetIamBinding` and `bigquery.DatasetIamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigquery.DatasetIamBinding` resources **can be** used in conjunction with `bigquery.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigquery\_dataset\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		owner, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/dataOwner",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewDatasetIamPolicy(ctx, "dataset", &bigquery.DatasetIamPolicyArgs{
// 			DatasetId:  pulumi.String("your-dataset-id"),
// 			PolicyData: pulumi.String(owner.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_bigquery\_dataset\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewDatasetIamBinding(ctx, "reader", &bigquery.DatasetIamBindingArgs{
// 			DatasetId: pulumi.String("your-dataset-id"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/bigquery.dataViewer"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_bigquery\_dataset\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewDatasetIamMember(ctx, "editor", &bigquery.DatasetIamMemberArgs{
// 			DatasetId: pulumi.String("your-dataset-id"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Role:      pulumi.String("roles/bigquery.dataEditor"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `dataset_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/datasetIamMember:DatasetIamMember dataset_iam "projects/your-project-id/datasets/dataset-id roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `dataset_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/datasetIamMember:DatasetIamMember dataset_iam "projects/your-project-id/datasets/dataset-id roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `dataset_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/datasetIamMember:DatasetIamMember dataset_iam projects/your-project-id/datasets/dataset-id
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatasetIamMember struct {
	pulumi.CustomResourceState

	Condition DatasetIamMemberConditionPtrOutput `pulumi:"condition"`
	// The dataset ID.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringOutput `pulumi:"etag"`
	Member  pulumi.StringOutput `pulumi:"member"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatasetIamMember registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamMember(ctx *pulumi.Context,
	name string, args *DatasetIamMemberArgs, opts ...pulumi.ResourceOption) (*DatasetIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DatasetIamMember
	err := ctx.RegisterResource("gcp:bigquery/datasetIamMember:DatasetIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamMember gets an existing DatasetIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamMemberState, opts ...pulumi.ResourceOption) (*DatasetIamMember, error) {
	var resource DatasetIamMember
	err := ctx.ReadResource("gcp:bigquery/datasetIamMember:DatasetIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamMember resources.
type datasetIamMemberState struct {
	Condition *DatasetIamMemberCondition `pulumi:"condition"`
	// The dataset ID.
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    *string `pulumi:"etag"`
	Member  *string `pulumi:"member"`
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatasetIamMemberState struct {
	Condition DatasetIamMemberConditionPtrInput
	// The dataset ID.
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringPtrInput
	Member  pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatasetIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamMemberState)(nil)).Elem()
}

type datasetIamMemberArgs struct {
	Condition *DatasetIamMemberCondition `pulumi:"condition"`
	// The dataset ID.
	DatasetId string  `pulumi:"datasetId"`
	Member    string  `pulumi:"member"`
	Project   *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatasetIamMember resource.
type DatasetIamMemberArgs struct {
	Condition DatasetIamMemberConditionPtrInput
	// The dataset ID.
	DatasetId pulumi.StringInput
	Member    pulumi.StringInput
	Project   pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatasetIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamMemberArgs)(nil)).Elem()
}

type DatasetIamMemberInput interface {
	pulumi.Input

	ToDatasetIamMemberOutput() DatasetIamMemberOutput
	ToDatasetIamMemberOutputWithContext(ctx context.Context) DatasetIamMemberOutput
}

func (DatasetIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamMember)(nil)).Elem()
}

func (i DatasetIamMember) ToDatasetIamMemberOutput() DatasetIamMemberOutput {
	return i.ToDatasetIamMemberOutputWithContext(context.Background())
}

func (i DatasetIamMember) ToDatasetIamMemberOutputWithContext(ctx context.Context) DatasetIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamMemberOutput)
}

type DatasetIamMemberOutput struct {
	*pulumi.OutputState
}

func (DatasetIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamMemberOutput)(nil)).Elem()
}

func (o DatasetIamMemberOutput) ToDatasetIamMemberOutput() DatasetIamMemberOutput {
	return o
}

func (o DatasetIamMemberOutput) ToDatasetIamMemberOutputWithContext(ctx context.Context) DatasetIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetIamMemberOutput{})
}
