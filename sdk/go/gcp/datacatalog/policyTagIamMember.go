// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* {{policy_tag}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor {{policy_tag}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type PolicyTagIamMember struct {
	pulumi.CustomResourceState

	Condition PolicyTagIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringOutput `pulumi:"policyTag"`
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewPolicyTagIamMember registers a new resource with the given unique name, arguments, and options.
func NewPolicyTagIamMember(ctx *pulumi.Context,
	name string, args *PolicyTagIamMemberArgs, opts ...pulumi.ResourceOption) (*PolicyTagIamMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.PolicyTag == nil {
		return nil, errors.New("missing required argument 'PolicyTag'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &PolicyTagIamMemberArgs{}
	}
	var resource PolicyTagIamMember
	err := ctx.RegisterResource("gcp:datacatalog/policyTagIamMember:PolicyTagIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTagIamMember gets an existing PolicyTagIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTagIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTagIamMemberState, opts ...pulumi.ResourceOption) (*PolicyTagIamMember, error) {
	var resource PolicyTagIamMember
	err := ctx.ReadResource("gcp:datacatalog/policyTagIamMember:PolicyTagIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTagIamMember resources.
type policyTagIamMemberState struct {
	Condition *PolicyTagIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag *string `pulumi:"policyTag"`
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type PolicyTagIamMemberState struct {
	Condition PolicyTagIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (PolicyTagIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamMemberState)(nil)).Elem()
}

type policyTagIamMemberArgs struct {
	Condition *PolicyTagIamMemberCondition `pulumi:"condition"`
	Member    string                       `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag string `pulumi:"policyTag"`
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a PolicyTagIamMember resource.
type PolicyTagIamMemberArgs struct {
	Condition PolicyTagIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringInput
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (PolicyTagIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamMemberArgs)(nil)).Elem()
}

type PolicyTagIamMemberInput interface {
	pulumi.Input

	ToPolicyTagIamMemberOutput() PolicyTagIamMemberOutput
	ToPolicyTagIamMemberOutputWithContext(ctx context.Context) PolicyTagIamMemberOutput
}

func (PolicyTagIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTagIamMember)(nil)).Elem()
}

func (i PolicyTagIamMember) ToPolicyTagIamMemberOutput() PolicyTagIamMemberOutput {
	return i.ToPolicyTagIamMemberOutputWithContext(context.Background())
}

func (i PolicyTagIamMember) ToPolicyTagIamMemberOutputWithContext(ctx context.Context) PolicyTagIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamMemberOutput)
}

type PolicyTagIamMemberOutput struct {
	*pulumi.OutputState
}

func (PolicyTagIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTagIamMemberOutput)(nil)).Elem()
}

func (o PolicyTagIamMemberOutput) ToPolicyTagIamMemberOutput() PolicyTagIamMemberOutput {
	return o
}

func (o PolicyTagIamMemberOutput) ToPolicyTagIamMemberOutputWithContext(ctx context.Context) PolicyTagIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyTagIamMemberOutput{})
}
