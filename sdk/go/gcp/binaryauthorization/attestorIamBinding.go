// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"fmt"
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
type AttestorIamBinding struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringOutput                  `pulumi:"attestor"`
	Condition AttestorIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAttestorIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAttestorIamBinding(ctx *pulumi.Context,
	name string, args *AttestorIamBindingArgs, opts ...pulumi.ResourceOption) (*AttestorIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Attestor == nil {
		return nil, errors.New("invalid value for required argument 'Attestor'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource AttestorIamBinding
	err := ctx.RegisterResource("gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestorIamBinding gets an existing AttestorIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestorIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorIamBindingState, opts ...pulumi.ResourceOption) (*AttestorIamBinding, error) {
	var resource AttestorIamBinding
	err := ctx.ReadResource("gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestorIamBinding resources.
type attestorIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  *string                      `pulumi:"attestor"`
	Condition *AttestorIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type AttestorIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringPtrInput
	Condition AttestorIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (AttestorIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamBindingState)(nil)).Elem()
}

type attestorIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  string                       `pulumi:"attestor"`
	Condition *AttestorIamBindingCondition `pulumi:"condition"`
	Members   []string                     `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AttestorIamBinding resource.
type AttestorIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringInput
	Condition AttestorIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (AttestorIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamBindingArgs)(nil)).Elem()
}
