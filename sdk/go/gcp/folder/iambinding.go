// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows creation and management of a single binding within IAM policy for
// an existing Google Cloud Platform folder.
//
// > **Note:** This resource _must not_ be used in conjunction with
//    `folder.IAMPolicy` or they will fight over what your policy
//    should be.
//
// > **Note:** On create, this resource will overwrite members of any existing roles.
//     Use `pulumi import` and inspect the output to ensure
//     your existing members are preserved.
type IAMBinding struct {
	pulumi.CustomResourceState

	Condition IAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the folder's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource IAMBinding
	err := ctx.RegisterResource("gcp:folder/iAMBinding:IAMBinding", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:folder/iAMBinding:IAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMBinding resources.
type iambindingState struct {
	Condition *IAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the folder's IAM policy.
	Etag *string `pulumi:"etag"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder *string `pulumi:"folder"`
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type IAMBindingState struct {
	Condition IAMBindingConditionPtrInput
	// (Computed) The etag of the folder's IAM policy.
	Etag pulumi.StringPtrInput
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringPtrInput
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (IAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingState)(nil)).Elem()
}

type iambindingArgs struct {
	Condition *IAMBindingCondition `pulumi:"condition"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder string `pulumi:"folder"`
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	Condition IAMBindingConditionPtrInput
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringInput
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that is associated with a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (IAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingArgs)(nil)).Elem()
}

type IAMBindingInput interface {
	pulumi.Input

	ToIAMBindingOutput() IAMBindingOutput
	ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput
}

func (IAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMBinding)(nil)).Elem()
}

func (i IAMBinding) ToIAMBindingOutput() IAMBindingOutput {
	return i.ToIAMBindingOutputWithContext(context.Background())
}

func (i IAMBinding) ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMBindingOutput)
}

type IAMBindingOutput struct {
	*pulumi.OutputState
}

func (IAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMBindingOutput)(nil)).Elem()
}

func (o IAMBindingOutput) ToIAMBindingOutput() IAMBindingOutput {
	return o
}

func (o IAMBindingOutput) ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IAMBindingOutput{})
}
