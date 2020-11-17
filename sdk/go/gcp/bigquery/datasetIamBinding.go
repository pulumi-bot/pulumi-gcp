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
type DatasetIamBinding struct {
	pulumi.CustomResourceState

	Condition DatasetIamBindingConditionPtrOutput `pulumi:"condition"`
	// The dataset ID.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	Project pulumi.StringOutput      `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatasetIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamBinding(ctx *pulumi.Context,
	name string, args *DatasetIamBindingArgs, opts ...pulumi.ResourceOption) (*DatasetIamBinding, error) {
	if args == nil || args.DatasetId == nil {
		return nil, errors.New("missing required argument 'DatasetId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &DatasetIamBindingArgs{}
	}
	var resource DatasetIamBinding
	err := ctx.RegisterResource("gcp:bigquery/datasetIamBinding:DatasetIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamBinding gets an existing DatasetIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamBindingState, opts ...pulumi.ResourceOption) (*DatasetIamBinding, error) {
	var resource DatasetIamBinding
	err := ctx.ReadResource("gcp:bigquery/datasetIamBinding:DatasetIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamBinding resources.
type datasetIamBindingState struct {
	Condition *DatasetIamBindingCondition `pulumi:"condition"`
	// The dataset ID.
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	Project *string  `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatasetIamBindingState struct {
	Condition DatasetIamBindingConditionPtrInput
	// The dataset ID.
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatasetIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamBindingState)(nil)).Elem()
}

type datasetIamBindingArgs struct {
	Condition *DatasetIamBindingCondition `pulumi:"condition"`
	// The dataset ID.
	DatasetId string   `pulumi:"datasetId"`
	Members   []string `pulumi:"members"`
	Project   *string  `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatasetIamBinding resource.
type DatasetIamBindingArgs struct {
	Condition DatasetIamBindingConditionPtrInput
	// The dataset ID.
	DatasetId pulumi.StringInput
	Members   pulumi.StringArrayInput
	Project   pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatasetIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamBindingArgs)(nil)).Elem()
}

type DatasetIamBindingInput interface {
	pulumi.Input

	ToDatasetIamBindingOutput() DatasetIamBindingOutput
	ToDatasetIamBindingOutputWithContext(ctx context.Context) DatasetIamBindingOutput
}

func (DatasetIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamBinding)(nil)).Elem()
}

func (i DatasetIamBinding) ToDatasetIamBindingOutput() DatasetIamBindingOutput {
	return i.ToDatasetIamBindingOutputWithContext(context.Background())
}

func (i DatasetIamBinding) ToDatasetIamBindingOutputWithContext(ctx context.Context) DatasetIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamBindingOutput)
}

type DatasetIamBindingOutput struct {
	*pulumi.OutputState
}

func (DatasetIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamBindingOutput)(nil)).Elem()
}

func (o DatasetIamBindingOutput) ToDatasetIamBindingOutput() DatasetIamBindingOutput {
	return o
}

func (o DatasetIamBindingOutput) ToDatasetIamBindingOutputWithContext(ctx context.Context) DatasetIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetIamBindingOutput{})
}
