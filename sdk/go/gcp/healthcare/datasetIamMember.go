// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare dataset. Each of these resources serves a different use case:
//
// * `healthcare.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
// * `healthcare.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
// * `healthcare.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
//
// > **Note:** `healthcare.DatasetIamPolicy` **cannot** be used in conjunction with `healthcare.DatasetIamBinding` and `healthcare.DatasetIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.DatasetIamBinding` resources **can be** used in conjunction with `healthcare.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
type DatasetIamMember struct {
	pulumi.CustomResourceState

	Condition DatasetIamMemberConditionPtrOutput `pulumi:"condition"`
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatasetIamMember registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamMember(ctx *pulumi.Context,
	name string, args *DatasetIamMemberArgs, opts ...pulumi.ResourceOption) (*DatasetIamMember, error) {
	if args == nil || args.DatasetId == nil {
		return nil, errors.New("missing required argument 'DatasetId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &DatasetIamMemberArgs{}
	}
	var resource DatasetIamMember
	err := ctx.RegisterResource("gcp:healthcare/datasetIamMember:DatasetIamMember", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:healthcare/datasetIamMember:DatasetIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamMember resources.
type datasetIamMemberState struct {
	Condition *DatasetIamMemberCondition `pulumi:"condition"`
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatasetIamMemberState struct {
	Condition DatasetIamMemberConditionPtrInput
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the dataset's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatasetIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamMemberState)(nil)).Elem()
}

type datasetIamMemberArgs struct {
	Condition *DatasetIamMemberCondition `pulumi:"condition"`
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId string `pulumi:"datasetId"`
	Member    string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatasetIamMember resource.
type DatasetIamMemberArgs struct {
	Condition DatasetIamMemberConditionPtrInput
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringInput
	Member    pulumi.StringInput
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
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
