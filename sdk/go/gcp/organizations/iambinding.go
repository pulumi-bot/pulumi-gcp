// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of a single binding within IAM policy for
// an existing Google Cloud Platform Organization.
//
// > **Note:** This resource __must not__ be used in conjunction with
//    `organizations.IAMMember` for the __same role__ or they will fight over
//    what your policy should be.
//
// > **Note:** On create, this resource will overwrite members of any existing roles.
//     Use `pulumi import` and inspect the `output to ensure
//     your existing members are preserved.
type IAMBinding struct {
	pulumi.CustomResourceState

	Condition IAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &IAMBindingArgs{}
	}
	var resource IAMBinding
	err := ctx.RegisterResource("gcp:organizations/iAMBinding:IAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMBinding gets an existing IAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMBindingState, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	var resource IAMBinding
	err := ctx.ReadResource("gcp:organizations/iAMBinding:IAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMBinding resources.
type iambindingState struct {
	Condition *IAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag *string `pulumi:"etag"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId *string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type IAMBindingState struct {
	Condition IAMBindingConditionPtrInput
	// (Computed) The etag of the organization's IAM policy.
	Etag pulumi.StringPtrInput
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (IAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingState)(nil)).Elem()
}

type iambindingArgs struct {
	Condition *IAMBindingCondition `pulumi:"condition"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	Condition IAMBindingConditionPtrInput
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (IAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingArgs)(nil)).Elem()
}
