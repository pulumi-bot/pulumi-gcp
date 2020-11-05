// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Run Service. Each of these resources serves a different use case:
//
// * `cloudrun.IamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `cloudrun.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `cloudrun.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// > **Note:** `cloudrun.IamPolicy` **cannot** be used in conjunction with `cloudrun.IamBinding` and `cloudrun.IamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudrun.IamBinding` resources **can be** used in conjunction with `cloudrun.IamMember` resources **only if** they do not grant privilege to the same role.
type IamMember struct {
	pulumi.CustomResourceState

	Condition IamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIamMember registers a new resource with the given unique name, arguments, and options.
func NewIamMember(ctx *pulumi.Context,
	name string, args *IamMemberArgs, opts ...pulumi.ResourceOption) (*IamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource IamMember
	err := ctx.RegisterResource("gcp:cloudrun/iamMember:IamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamMember gets an existing IamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamMemberState, opts ...pulumi.ResourceOption) (*IamMember, error) {
	var resource IamMember
	err := ctx.ReadResource("gcp:cloudrun/iamMember:IamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamMember resources.
type iamMemberState struct {
	Condition *IamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
}

type IamMemberState struct {
	Condition IamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
}

func (IamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamMemberState)(nil)).Elem()
}

type iamMemberArgs struct {
	Condition *IamMemberCondition `pulumi:"condition"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   string  `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamMember resource.
type IamMemberArgs struct {
	Condition IamMemberConditionPtrInput
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
}

func (IamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamMemberArgs)(nil)).Elem()
}
