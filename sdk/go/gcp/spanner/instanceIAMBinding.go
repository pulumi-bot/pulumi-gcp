// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:
//
// * `spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
//
// > **Warning:** It's entirely possibly to lock yourself out of your instance using `spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
//
// * `spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `spanner.InstanceIAMBinding` and `spanner.InstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `spanner.InstanceIAMBinding` resources **can be** used in conjunction with `spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
type InstanceIAMBinding struct {
	pulumi.CustomResourceState

	Condition InstanceIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the instance.
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMBinding(ctx *pulumi.Context,
	name string, args *InstanceIAMBindingArgs, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &InstanceIAMBindingArgs{}
	}
	var resource InstanceIAMBinding
	err := ctx.RegisterResource("gcp:spanner/instanceIAMBinding:InstanceIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIAMBinding gets an existing InstanceIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIAMBindingState, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
	var resource InstanceIAMBinding
	err := ctx.ReadResource("gcp:spanner/instanceIAMBinding:InstanceIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIAMBinding resources.
type instanceIAMBindingState struct {
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the instance's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the instance.
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type InstanceIAMBindingState struct {
	Condition InstanceIAMBindingConditionPtrInput
	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the instance.
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (InstanceIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingState)(nil)).Elem()
}

type instanceIAMBindingArgs struct {
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// The name of the instance.
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIAMBinding resource.
type InstanceIAMBindingArgs struct {
	Condition InstanceIAMBindingConditionPtrInput
	// The name of the instance.
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (InstanceIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingArgs)(nil)).Elem()
}

type InstanceIAMBindingInput interface {
	pulumi.Input

	ToInstanceIAMBindingOutput() InstanceIAMBindingOutput
	ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput
}

func (InstanceIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMBinding)(nil)).Elem()
}

func (i InstanceIAMBinding) ToInstanceIAMBindingOutput() InstanceIAMBindingOutput {
	return i.ToInstanceIAMBindingOutputWithContext(context.Background())
}

func (i InstanceIAMBinding) ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingOutput)
}

type InstanceIAMBindingOutput struct {
	*pulumi.OutputState
}

func (InstanceIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMBindingOutput)(nil)).Elem()
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingOutput() InstanceIAMBindingOutput {
	return o
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceIAMBindingOutput{})
}
