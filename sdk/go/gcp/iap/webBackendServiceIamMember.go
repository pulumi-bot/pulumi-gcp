// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:
//
// * `iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
// * `iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
// * `iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.
//
// > **Note:** `iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `iap.WebBackendServiceIamBinding` and `iap.WebBackendServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.
type WebBackendServiceIamMember struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringOutput `pulumi:"webBackendService"`
}

// NewWebBackendServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, args *WebBackendServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*WebBackendServiceIamMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.WebBackendService == nil {
		return nil, errors.New("missing required argument 'WebBackendService'")
	}
	if args == nil {
		args = &WebBackendServiceIamMemberArgs{}
	}
	var resource WebBackendServiceIamMember
	err := ctx.RegisterResource("gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebBackendServiceIamMember gets an existing WebBackendServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebBackendServiceIamMemberState, opts ...pulumi.ResourceOption) (*WebBackendServiceIamMember, error) {
	var resource WebBackendServiceIamMember
	err := ctx.ReadResource("gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebBackendServiceIamMember resources.
type webBackendServiceIamMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebBackendServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService *string `pulumi:"webBackendService"`
}

type WebBackendServiceIamMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringPtrInput
}

func (WebBackendServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamMemberState)(nil)).Elem()
}

type webBackendServiceIamMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebBackendServiceIamMemberCondition `pulumi:"condition"`
	Member    string                               `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService string `pulumi:"webBackendService"`
}

// The set of arguments for constructing a WebBackendServiceIamMember resource.
type WebBackendServiceIamMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringInput
}

func (WebBackendServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamMemberArgs)(nil)).Elem()
}

type WebBackendServiceIamMemberInput interface {
	pulumi.Input

	ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput
	ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput
}

func (WebBackendServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamMember)(nil)).Elem()
}

func (i WebBackendServiceIamMember) ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput {
	return i.ToWebBackendServiceIamMemberOutputWithContext(context.Background())
}

func (i WebBackendServiceIamMember) ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberOutput)
}

type WebBackendServiceIamMemberOutput struct {
	*pulumi.OutputState
}

func (WebBackendServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamMemberOutput)(nil)).Elem()
}

func (o WebBackendServiceIamMemberOutput) ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput {
	return o
}

func (o WebBackendServiceIamMemberOutput) ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebBackendServiceIamMemberOutput{})
}
