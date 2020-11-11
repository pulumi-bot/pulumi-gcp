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
type FunctionIamBinding struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringOutput                  `pulumi:"cloudFunction"`
	Condition     FunctionIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFunctionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamBinding(ctx *pulumi.Context,
	name string, args *FunctionIamBindingArgs, opts ...pulumi.ResourceOption) (*FunctionIamBinding, error) {
	if args == nil || args.CloudFunction == nil {
		return nil, errors.New("missing required argument 'CloudFunction'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &FunctionIamBindingArgs{}
	}
	var resource FunctionIamBinding
	err := ctx.RegisterResource("gcp:cloudfunctions/functionIamBinding:FunctionIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamBinding gets an existing FunctionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamBindingState, opts ...pulumi.ResourceOption) (*FunctionIamBinding, error) {
	var resource FunctionIamBinding
	err := ctx.ReadResource("gcp:cloudfunctions/functionIamBinding:FunctionIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamBinding resources.
type functionIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction *string                      `pulumi:"cloudFunction"`
	Condition     *FunctionIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FunctionIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringPtrInput
	Condition     FunctionIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FunctionIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamBindingState)(nil)).Elem()
}

type functionIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction string                       `pulumi:"cloudFunction"`
	Condition     *FunctionIamBindingCondition `pulumi:"condition"`
	Members       []string                     `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FunctionIamBinding resource.
type FunctionIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringInput
	Condition     FunctionIamBindingConditionPtrInput
	Members       pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FunctionIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamBindingArgs)(nil)).Elem()
}

type FunctionIamBindingInput interface {
	pulumi.Input

	ToFunctionIamBindingOutput() FunctionIamBindingOutput
	ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput
}

func (FunctionIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionIamBinding)(nil)).Elem()
}

func (i FunctionIamBinding) ToFunctionIamBindingOutput() FunctionIamBindingOutput {
	return i.ToFunctionIamBindingOutputWithContext(context.Background())
}

func (i FunctionIamBinding) ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamBindingOutput)
}

type FunctionIamBindingOutput struct {
	*pulumi.OutputState
}

func (FunctionIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionIamBindingOutput)(nil)).Elem()
}

func (o FunctionIamBindingOutput) ToFunctionIamBindingOutput() FunctionIamBindingOutput {
	return o
}

func (o FunctionIamBindingOutput) ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FunctionIamBindingOutput{})
}
