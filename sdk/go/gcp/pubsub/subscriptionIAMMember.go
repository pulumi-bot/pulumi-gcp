// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/editor",
// 					"members": []string{
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = pubsub.NewSubscriptionIAMBinding(ctx, "editor", &pubsub.SubscriptionIAMBindingArgs{
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = pubsub.NewSubscriptionIAMMember(ctx, "editor", &pubsub.SubscriptionIAMMemberArgs{
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
type SubscriptionIAMMember struct {
	pulumi.CustomResourceState

	Condition SubscriptionIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the subscription's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
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

// NewSubscriptionIAMMember registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionIAMMember(ctx *pulumi.Context,
	name string, args *SubscriptionIAMMemberArgs, opts ...pulumi.ResourceOption) (*SubscriptionIAMMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Subscription == nil {
		return nil, errors.New("missing required argument 'Subscription'")
	}
	if args == nil {
		args = &SubscriptionIAMMemberArgs{}
	}
	var resource SubscriptionIAMMember
	err := ctx.RegisterResource("gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionIAMMember gets an existing SubscriptionIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionIAMMemberState, opts ...pulumi.ResourceOption) (*SubscriptionIAMMember, error) {
	var resource SubscriptionIAMMember
	err := ctx.ReadResource("gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionIAMMember resources.
type subscriptionIAMMemberState struct {
	Condition *SubscriptionIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the subscription's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
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

type SubscriptionIAMMemberState struct {
	Condition SubscriptionIAMMemberConditionPtrInput
	// (Computed) The etag of the subscription's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
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

func (SubscriptionIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMMemberState)(nil)).Elem()
}

type subscriptionIAMMemberArgs struct {
	Condition *SubscriptionIAMMemberCondition `pulumi:"condition"`
	Member    string                          `pulumi:"member"`
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

// The set of arguments for constructing a SubscriptionIAMMember resource.
type SubscriptionIAMMemberArgs struct {
	Condition SubscriptionIAMMemberConditionPtrInput
	Member    pulumi.StringInput
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

func (SubscriptionIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMMemberArgs)(nil)).Elem()
}
