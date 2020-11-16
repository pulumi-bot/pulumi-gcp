// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Functions CloudFunction. Each of these resources serves a different use case:
//
// * `cloudfunctions.FunctionIamPolicy`: Authoritative. Sets the IAM policy for the cloudfunction and replaces any existing policy already attached.
// * `cloudfunctions.FunctionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cloudfunction are preserved.
// * `cloudfunctions.FunctionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cloudfunction are preserved.
//
// > **Note:** `cloudfunctions.FunctionIamPolicy` **cannot** be used in conjunction with `cloudfunctions.FunctionIamBinding` and `cloudfunctions.FunctionIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudfunctions.FunctionIamBinding` resources **can be** used in conjunction with `cloudfunctions.FunctionIamMember` resources **only if** they do not grant privilege to the same role.
type FunctionIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringOutput `pulumi:"cloudFunction"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewFunctionIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamPolicy(ctx *pulumi.Context,
	name string, args *FunctionIamPolicyArgs, opts ...pulumi.ResourceOption) (*FunctionIamPolicy, error) {
	if args == nil || args.CloudFunction == nil {
		return nil, errors.New("missing required argument 'CloudFunction'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &FunctionIamPolicyArgs{}
	}
	var resource FunctionIamPolicy
	err := ctx.RegisterResource("gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamPolicy gets an existing FunctionIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamPolicyState, opts ...pulumi.ResourceOption) (*FunctionIamPolicy, error) {
	var resource FunctionIamPolicy
	err := ctx.ReadResource("gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamPolicy resources.
type functionIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction *string `pulumi:"cloudFunction"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
}

type FunctionIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
}

func (FunctionIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamPolicyState)(nil)).Elem()
}

type functionIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction string `pulumi:"cloudFunction"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a FunctionIamPolicy resource.
type FunctionIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
}

func (FunctionIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamPolicyArgs)(nil)).Elem()
}

type FunctionIamPolicyInput interface {
	pulumi.Input

	ToFunctionIamPolicyOutput() FunctionIamPolicyOutput
	ToFunctionIamPolicyOutputWithContext(ctx context.Context) FunctionIamPolicyOutput
}

func (FunctionIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionIamPolicy)(nil)).Elem()
}

func (i FunctionIamPolicy) ToFunctionIamPolicyOutput() FunctionIamPolicyOutput {
	return i.ToFunctionIamPolicyOutputWithContext(context.Background())
}

func (i FunctionIamPolicy) ToFunctionIamPolicyOutputWithContext(ctx context.Context) FunctionIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamPolicyOutput)
}

type FunctionIamPolicyOutput struct {
	*pulumi.OutputState
}

func (FunctionIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionIamPolicyOutput)(nil)).Elem()
}

func (o FunctionIamPolicyOutput) ToFunctionIamPolicyOutput() FunctionIamPolicyOutput {
	return o
}

func (o FunctionIamPolicyOutput) ToFunctionIamPolicyOutputWithContext(ctx context.Context) FunctionIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FunctionIamPolicyOutput{})
}
