// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
//
// * `pubsub.SubscriptionIAMPolicy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
// * `pubsub.SubscriptionIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
// * `pubsub.SubscriptionIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
//
// > **Note:** `pubsub.SubscriptionIAMPolicy` **cannot** be used in conjunction with `pubsub.SubscriptionIAMBinding` and `pubsub.SubscriptionIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `pubsub.SubscriptionIAMBinding` resources **can be** used in conjunction with `pubsub.SubscriptionIAMMember` resources **only if** they do not grant privilege to the same role.
type SubscriptionIAMBinding struct {
	pulumi.CustomResourceState

	Condition SubscriptionIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the subscription's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringOutput `pulumi:"subscription"`
}

// NewSubscriptionIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionIAMBinding(ctx *pulumi.Context,
	name string, args *SubscriptionIAMBindingArgs, opts ...pulumi.ResourceOption) (*SubscriptionIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Subscription == nil {
		return nil, errors.New("invalid value for required argument 'Subscription'")
	}
	var resource SubscriptionIAMBinding
	err := ctx.RegisterResource("gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionIAMBinding gets an existing SubscriptionIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionIAMBindingState, opts ...pulumi.ResourceOption) (*SubscriptionIAMBinding, error) {
	var resource SubscriptionIAMBinding
	err := ctx.ReadResource("gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionIAMBinding resources.
type subscriptionIAMBindingState struct {
	Condition *SubscriptionIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the subscription's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription *string `pulumi:"subscription"`
}

type SubscriptionIAMBindingState struct {
	Condition SubscriptionIAMBindingConditionPtrInput
	// (Computed) The etag of the subscription's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringPtrInput
}

func (SubscriptionIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMBindingState)(nil)).Elem()
}

type subscriptionIAMBindingArgs struct {
	Condition *SubscriptionIAMBindingCondition `pulumi:"condition"`
	Members   []string                         `pulumi:"members"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription string `pulumi:"subscription"`
}

// The set of arguments for constructing a SubscriptionIAMBinding resource.
type SubscriptionIAMBindingArgs struct {
	Condition SubscriptionIAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringInput
}

func (SubscriptionIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMBindingArgs)(nil)).Elem()
}

type SubscriptionIAMBindingInput interface {
	pulumi.Input

	ToSubscriptionIAMBindingOutput() SubscriptionIAMBindingOutput
	ToSubscriptionIAMBindingOutputWithContext(ctx context.Context) SubscriptionIAMBindingOutput
}

func (SubscriptionIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionIAMBinding)(nil)).Elem()
}

func (i SubscriptionIAMBinding) ToSubscriptionIAMBindingOutput() SubscriptionIAMBindingOutput {
	return i.ToSubscriptionIAMBindingOutputWithContext(context.Background())
}

func (i SubscriptionIAMBinding) ToSubscriptionIAMBindingOutputWithContext(ctx context.Context) SubscriptionIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMBindingOutput)
}

type SubscriptionIAMBindingOutput struct {
	*pulumi.OutputState
}

func (SubscriptionIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionIAMBindingOutput)(nil)).Elem()
}

func (o SubscriptionIAMBindingOutput) ToSubscriptionIAMBindingOutput() SubscriptionIAMBindingOutput {
	return o
}

func (o SubscriptionIAMBindingOutput) ToSubscriptionIAMBindingOutputWithContext(ctx context.Context) SubscriptionIAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionIAMBindingOutput{})
}
