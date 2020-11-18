// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:
//
// * `iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
// * `iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
// * `iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.
//
// > **Note:** `iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `iap.WebTypeAppEngingIamBinding` and `iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.
type WebTypeAppEngingIamBinding struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeAppEngingIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewWebTypeAppEngingIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWebTypeAppEngingIamBinding(ctx *pulumi.Context,
	name string, args *WebTypeAppEngingIamBindingArgs, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource WebTypeAppEngingIamBinding
	err := ctx.RegisterResource("gcp:iap/webTypeAppEngingIamBinding:WebTypeAppEngingIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeAppEngingIamBinding gets an existing WebTypeAppEngingIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeAppEngingIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeAppEngingIamBindingState, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamBinding, error) {
	var resource WebTypeAppEngingIamBinding
	err := ctx.ReadResource("gcp:iap/webTypeAppEngingIamBinding:WebTypeAppEngingIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeAppEngingIamBinding resources.
type webTypeAppEngingIamBindingState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeAppEngingIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type WebTypeAppEngingIamBindingState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeAppEngingIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (WebTypeAppEngingIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamBindingState)(nil)).Elem()
}

type webTypeAppEngingIamBindingArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeAppEngingIamBindingCondition `pulumi:"condition"`
	Members   []string                             `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a WebTypeAppEngingIamBinding resource.
type WebTypeAppEngingIamBindingArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeAppEngingIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (WebTypeAppEngingIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamBindingArgs)(nil)).Elem()
}

type WebTypeAppEngingIamBindingInput interface {
	pulumi.Input

	ToWebTypeAppEngingIamBindingOutput() WebTypeAppEngingIamBindingOutput
	ToWebTypeAppEngingIamBindingOutputWithContext(ctx context.Context) WebTypeAppEngingIamBindingOutput
}

func (WebTypeAppEngingIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeAppEngingIamBinding)(nil)).Elem()
}

func (i WebTypeAppEngingIamBinding) ToWebTypeAppEngingIamBindingOutput() WebTypeAppEngingIamBindingOutput {
	return i.ToWebTypeAppEngingIamBindingOutputWithContext(context.Background())
}

func (i WebTypeAppEngingIamBinding) ToWebTypeAppEngingIamBindingOutputWithContext(ctx context.Context) WebTypeAppEngingIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeAppEngingIamBindingOutput)
}

type WebTypeAppEngingIamBindingOutput struct {
	*pulumi.OutputState
}

func (WebTypeAppEngingIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeAppEngingIamBindingOutput)(nil)).Elem()
}

func (o WebTypeAppEngingIamBindingOutput) ToWebTypeAppEngingIamBindingOutput() WebTypeAppEngingIamBindingOutput {
	return o
}

func (o WebTypeAppEngingIamBindingOutput) ToWebTypeAppEngingIamBindingOutputWithContext(ctx context.Context) WebTypeAppEngingIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebTypeAppEngingIamBindingOutput{})
}
