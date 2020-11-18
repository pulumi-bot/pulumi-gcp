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
type AttestorIamMember struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringOutput                 `pulumi:"attestor"`
	Condition AttestorIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAttestorIamMember registers a new resource with the given unique name, arguments, and options.
func NewAttestorIamMember(ctx *pulumi.Context,
	name string, args *AttestorIamMemberArgs, opts ...pulumi.ResourceOption) (*AttestorIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Attestor == nil {
		return nil, errors.New("invalid value for required argument 'Attestor'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource AttestorIamMember
	err := ctx.RegisterResource("gcp:binaryauthorization/attestorIamMember:AttestorIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestorIamMember gets an existing AttestorIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestorIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorIamMemberState, opts ...pulumi.ResourceOption) (*AttestorIamMember, error) {
	var resource AttestorIamMember
	err := ctx.ReadResource("gcp:binaryauthorization/attestorIamMember:AttestorIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestorIamMember resources.
type attestorIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  *string                     `pulumi:"attestor"`
	Condition *AttestorIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type AttestorIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringPtrInput
	Condition AttestorIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (AttestorIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamMemberState)(nil)).Elem()
}

type attestorIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  string                      `pulumi:"attestor"`
	Condition *AttestorIamMemberCondition `pulumi:"condition"`
	Member    string                      `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AttestorIamMember resource.
type AttestorIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringInput
	Condition AttestorIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (AttestorIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamMemberArgs)(nil)).Elem()
}

type AttestorIamMemberInput interface {
	pulumi.Input

	ToAttestorIamMemberOutput() AttestorIamMemberOutput
	ToAttestorIamMemberOutputWithContext(ctx context.Context) AttestorIamMemberOutput
}

func (AttestorIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamMember)(nil)).Elem()
}

func (i AttestorIamMember) ToAttestorIamMemberOutput() AttestorIamMemberOutput {
	return i.ToAttestorIamMemberOutputWithContext(context.Background())
}

func (i AttestorIamMember) ToAttestorIamMemberOutputWithContext(ctx context.Context) AttestorIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamMemberOutput)
}

type AttestorIamMemberOutput struct {
	*pulumi.OutputState
}

func (AttestorIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorIamMemberOutput)(nil)).Elem()
}

func (o AttestorIamMemberOutput) ToAttestorIamMemberOutput() AttestorIamMemberOutput {
	return o
}

func (o AttestorIamMemberOutput) ToAttestorIamMemberOutputWithContext(ctx context.Context) AttestorIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AttestorIamMemberOutput{})
}
