// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Binary Authorization Attestor. Each of these resources serves a different use case:
//
// * `binaryauthorization.AttestorIamPolicy`: Authoritative. Sets the IAM policy for the attestor and replaces any existing policy already attached.
// * `binaryauthorization.AttestorIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the attestor are preserved.
// * `binaryauthorization.AttestorIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the attestor are preserved.
//
// > **Note:** `binaryauthorization.AttestorIamPolicy` **cannot** be used in conjunction with `binaryauthorization.AttestorIamBinding` and `binaryauthorization.AttestorIamMember` or they will fight over what your policy should be.
//
// > **Note:** `binaryauthorization.AttestorIamBinding` resources **can be** used in conjunction with `binaryauthorization.AttestorIamMember` resources **only if** they do not grant privilege to the same role.
type AttestorIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Attestor pulumi.StringOutput `pulumi:"attestor"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewAttestorIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAttestorIamPolicy(ctx *pulumi.Context,
	name string, args *AttestorIamPolicyArgs, opts ...pulumi.ResourceOption) (*AttestorIamPolicy, error) {
	if args == nil || args.Attestor == nil {
		return nil, errors.New("missing required argument 'Attestor'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &AttestorIamPolicyArgs{}
	}
	var resource AttestorIamPolicy
	err := ctx.RegisterResource("gcp:binaryauthorization/attestorIamPolicy:AttestorIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestorIamPolicy gets an existing AttestorIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestorIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorIamPolicyState, opts ...pulumi.ResourceOption) (*AttestorIamPolicy, error) {
	var resource AttestorIamPolicy
	err := ctx.ReadResource("gcp:binaryauthorization/attestorIamPolicy:AttestorIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestorIamPolicy resources.
type attestorIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor *string `pulumi:"attestor"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type AttestorIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (AttestorIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamPolicyState)(nil)).Elem()
}

type attestorIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor string `pulumi:"attestor"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a AttestorIamPolicy resource.
type AttestorIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (AttestorIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamPolicyArgs)(nil)).Elem()
}

type AttestorIamPolicyInput interface {
	pulumi.Input

	ToAttestorIamPolicyOutput() AttestorIamPolicyOutput
	ToAttestorIamPolicyOutputWithContext(ctx context.Context) AttestorIamPolicyOutput
}

func (AttestorIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamPolicy)(nil)).Elem()
}

func (i AttestorIamPolicy) ToAttestorIamPolicyOutput() AttestorIamPolicyOutput {
	return i.ToAttestorIamPolicyOutputWithContext(context.Background())
}

func (i AttestorIamPolicy) ToAttestorIamPolicyOutputWithContext(ctx context.Context) AttestorIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamPolicyOutput)
}

type AttestorIamPolicyOutput struct {
	*pulumi.OutputState
}

func (AttestorIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamPolicyOutput)(nil)).Elem()
}

func (o AttestorIamPolicyOutput) ToAttestorIamPolicyOutput() AttestorIamPolicyOutput {
	return o
}

func (o AttestorIamPolicyOutput) ToAttestorIamPolicyOutputWithContext(ctx context.Context) AttestorIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AttestorIamPolicyOutput{})
}
