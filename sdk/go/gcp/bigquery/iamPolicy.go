// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:
//
// * `bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
// * `bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
// * `bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
//
// > **Note:** `bigquery.IamPolicy` **cannot** be used in conjunction with `bigquery.IamBinding` and `bigquery.IamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigquery.IamBinding` resources **can be** used in conjunction with `bigquery.IamMember` resources **only if** they do not grant privilege to the same role.
type IamPolicy struct {
	pulumi.CustomResourceState

	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	TableId pulumi.StringOutput `pulumi:"tableId"`
}

// NewIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewIamPolicy(ctx *pulumi.Context,
	name string, args *IamPolicyArgs, opts ...pulumi.ResourceOption) (*IamPolicy, error) {
	if args == nil || args.DatasetId == nil {
		return nil, errors.New("missing required argument 'DatasetId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.TableId == nil {
		return nil, errors.New("missing required argument 'TableId'")
	}
	if args == nil {
		args = &IamPolicyArgs{}
	}
	var resource IamPolicy
	err := ctx.RegisterResource("gcp:bigquery/iamPolicy:IamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamPolicy gets an existing IamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamPolicyState, opts ...pulumi.ResourceOption) (*IamPolicy, error) {
	var resource IamPolicy
	err := ctx.ReadResource("gcp:bigquery/iamPolicy:IamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamPolicy resources.
type iamPolicyState struct {
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	TableId *string `pulumi:"tableId"`
}

type IamPolicyState struct {
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	TableId pulumi.StringPtrInput
}

func (IamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyState)(nil)).Elem()
}

type iamPolicyArgs struct {
	DatasetId string `pulumi:"datasetId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	TableId string  `pulumi:"tableId"`
}

// The set of arguments for constructing a IamPolicy resource.
type IamPolicyArgs struct {
	DatasetId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	TableId pulumi.StringInput
}

func (IamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyArgs)(nil)).Elem()
}

type IamPolicyInput interface {
	pulumi.Input

	ToIamPolicyOutput() IamPolicyOutput
	ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput
}

func (IamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*IamPolicy)(nil)).Elem()
}

func (i IamPolicy) ToIamPolicyOutput() IamPolicyOutput {
	return i.ToIamPolicyOutputWithContext(context.Background())
}

func (i IamPolicy) ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyOutput)
}

type IamPolicyOutput struct {
	*pulumi.OutputState
}

func (IamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamPolicyOutput)(nil)).Elem()
}

func (o IamPolicyOutput) ToIamPolicyOutput() IamPolicyOutput {
	return o
}

func (o IamPolicyOutput) ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IamPolicyOutput{})
}
