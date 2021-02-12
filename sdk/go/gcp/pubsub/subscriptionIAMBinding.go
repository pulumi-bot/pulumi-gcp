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
//
// ## google\_pubsub\_subscription\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewSubscriptionIAMPolicy(ctx, "editor", &pubsub.SubscriptionIAMPolicyArgs{
// 			Subscription: pulumi.String("your-subscription-name"),
// 			PolicyData:   pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_subscription\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewSubscriptionIAMBinding(ctx, "editor", &pubsub.SubscriptionIAMBindingArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role:         pulumi.String("roles/editor"),
// 			Subscription: pulumi.String("your-subscription-name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_subscription\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewSubscriptionIAMMember(ctx, "editor", &pubsub.SubscriptionIAMMemberArgs{
// 			Member:       pulumi.String("user:jane@example.com"),
// 			Role:         pulumi.String("roles/editor"),
// 			Subscription: pulumi.String("your-subscription-name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Pubsub subscription IAM resources can be imported using the project, subscription name, role and member.
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding editor projects/{your-project-id}/subscriptions/{your-subscription-name}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
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

func (*SubscriptionIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionIAMBinding)(nil))
}

func (i *SubscriptionIAMBinding) ToSubscriptionIAMBindingOutput() SubscriptionIAMBindingOutput {
	return i.ToSubscriptionIAMBindingOutputWithContext(context.Background())
}

func (i *SubscriptionIAMBinding) ToSubscriptionIAMBindingOutputWithContext(ctx context.Context) SubscriptionIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMBindingOutput)
}

func (i *SubscriptionIAMBinding) ToSubscriptionIAMBindingPtrOutput() SubscriptionIAMBindingPtrOutput {
	return i.ToSubscriptionIAMBindingPtrOutputWithContext(context.Background())
}

func (i *SubscriptionIAMBinding) ToSubscriptionIAMBindingPtrOutputWithContext(ctx context.Context) SubscriptionIAMBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMBindingPtrOutput)
}

type SubscriptionIAMBindingPtrInput interface {
	pulumi.Input

	ToSubscriptionIAMBindingPtrOutput() SubscriptionIAMBindingPtrOutput
	ToSubscriptionIAMBindingPtrOutputWithContext(ctx context.Context) SubscriptionIAMBindingPtrOutput
}

type subscriptionIAMBindingPtrType SubscriptionIAMBindingArgs

func (*subscriptionIAMBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionIAMBinding)(nil))
}

func (i *subscriptionIAMBindingPtrType) ToSubscriptionIAMBindingPtrOutput() SubscriptionIAMBindingPtrOutput {
	return i.ToSubscriptionIAMBindingPtrOutputWithContext(context.Background())
}

func (i *subscriptionIAMBindingPtrType) ToSubscriptionIAMBindingPtrOutputWithContext(ctx context.Context) SubscriptionIAMBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMBindingPtrOutput)
}

// SubscriptionIAMBindingArrayInput is an input type that accepts SubscriptionIAMBindingArray and SubscriptionIAMBindingArrayOutput values.
// You can construct a concrete instance of `SubscriptionIAMBindingArrayInput` via:
//
//          SubscriptionIAMBindingArray{ SubscriptionIAMBindingArgs{...} }
type SubscriptionIAMBindingArrayInput interface {
	pulumi.Input

	ToSubscriptionIAMBindingArrayOutput() SubscriptionIAMBindingArrayOutput
	ToSubscriptionIAMBindingArrayOutputWithContext(context.Context) SubscriptionIAMBindingArrayOutput
}

type SubscriptionIAMBindingArray []SubscriptionIAMBindingInput

func (SubscriptionIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SubscriptionIAMBinding)(nil))
}

func (i SubscriptionIAMBindingArray) ToSubscriptionIAMBindingArrayOutput() SubscriptionIAMBindingArrayOutput {
	return i.ToSubscriptionIAMBindingArrayOutputWithContext(context.Background())
}

func (i SubscriptionIAMBindingArray) ToSubscriptionIAMBindingArrayOutputWithContext(ctx context.Context) SubscriptionIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMBindingArrayOutput)
}

// SubscriptionIAMBindingMapInput is an input type that accepts SubscriptionIAMBindingMap and SubscriptionIAMBindingMapOutput values.
// You can construct a concrete instance of `SubscriptionIAMBindingMapInput` via:
//
//          SubscriptionIAMBindingMap{ "key": SubscriptionIAMBindingArgs{...} }
type SubscriptionIAMBindingMapInput interface {
	pulumi.Input

	ToSubscriptionIAMBindingMapOutput() SubscriptionIAMBindingMapOutput
	ToSubscriptionIAMBindingMapOutputWithContext(context.Context) SubscriptionIAMBindingMapOutput
}

type SubscriptionIAMBindingMap map[string]SubscriptionIAMBindingInput

func (SubscriptionIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SubscriptionIAMBinding)(nil))
}

func (i SubscriptionIAMBindingMap) ToSubscriptionIAMBindingMapOutput() SubscriptionIAMBindingMapOutput {
	return i.ToSubscriptionIAMBindingMapOutputWithContext(context.Background())
}

func (i SubscriptionIAMBindingMap) ToSubscriptionIAMBindingMapOutputWithContext(ctx context.Context) SubscriptionIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMBindingMapOutput)
}

type SubscriptionIAMBindingOutput struct {
	*pulumi.OutputState
}

func (SubscriptionIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionIAMBinding)(nil))
}

func (o SubscriptionIAMBindingOutput) ToSubscriptionIAMBindingOutput() SubscriptionIAMBindingOutput {
	return o
}

func (o SubscriptionIAMBindingOutput) ToSubscriptionIAMBindingOutputWithContext(ctx context.Context) SubscriptionIAMBindingOutput {
	return o
}

func (o SubscriptionIAMBindingOutput) ToSubscriptionIAMBindingPtrOutput() SubscriptionIAMBindingPtrOutput {
	return o.ToSubscriptionIAMBindingPtrOutputWithContext(context.Background())
}

func (o SubscriptionIAMBindingOutput) ToSubscriptionIAMBindingPtrOutputWithContext(ctx context.Context) SubscriptionIAMBindingPtrOutput {
	return o.ApplyT(func(v SubscriptionIAMBinding) *SubscriptionIAMBinding {
		return &v
	}).(SubscriptionIAMBindingPtrOutput)
}

type SubscriptionIAMBindingPtrOutput struct {
	*pulumi.OutputState
}

func (SubscriptionIAMBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionIAMBinding)(nil))
}

func (o SubscriptionIAMBindingPtrOutput) ToSubscriptionIAMBindingPtrOutput() SubscriptionIAMBindingPtrOutput {
	return o
}

func (o SubscriptionIAMBindingPtrOutput) ToSubscriptionIAMBindingPtrOutputWithContext(ctx context.Context) SubscriptionIAMBindingPtrOutput {
	return o
}

type SubscriptionIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionIAMBinding)(nil))
}

func (o SubscriptionIAMBindingArrayOutput) ToSubscriptionIAMBindingArrayOutput() SubscriptionIAMBindingArrayOutput {
	return o
}

func (o SubscriptionIAMBindingArrayOutput) ToSubscriptionIAMBindingArrayOutputWithContext(ctx context.Context) SubscriptionIAMBindingArrayOutput {
	return o
}

func (o SubscriptionIAMBindingArrayOutput) Index(i pulumi.IntInput) SubscriptionIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionIAMBinding {
		return vs[0].([]SubscriptionIAMBinding)[vs[1].(int)]
	}).(SubscriptionIAMBindingOutput)
}

type SubscriptionIAMBindingMapOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SubscriptionIAMBinding)(nil))
}

func (o SubscriptionIAMBindingMapOutput) ToSubscriptionIAMBindingMapOutput() SubscriptionIAMBindingMapOutput {
	return o
}

func (o SubscriptionIAMBindingMapOutput) ToSubscriptionIAMBindingMapOutputWithContext(ctx context.Context) SubscriptionIAMBindingMapOutput {
	return o
}

func (o SubscriptionIAMBindingMapOutput) MapIndex(k pulumi.StringInput) SubscriptionIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SubscriptionIAMBinding {
		return vs[0].(map[string]SubscriptionIAMBinding)[vs[1].(string)]
	}).(SubscriptionIAMBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionIAMBindingOutput{})
	pulumi.RegisterOutputType(SubscriptionIAMBindingPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionIAMBindingMapOutput{})
}
