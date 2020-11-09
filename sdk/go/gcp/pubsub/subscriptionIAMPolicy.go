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
// ## Import
//
// Pubsub subscription IAM resources can be imported using the project, subscription name, role and member.
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy editor projects/{your-project-id}/subscriptions/{your-subscription-name}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type SubscriptionIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the subscription's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringOutput `pulumi:"subscription"`
}

// NewSubscriptionIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionIAMPolicy(ctx *pulumi.Context,
	name string, args *SubscriptionIAMPolicyArgs, opts ...pulumi.ResourceOption) (*SubscriptionIAMPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.Subscription == nil {
		return nil, errors.New("missing required argument 'Subscription'")
	}
	if args == nil {
		args = &SubscriptionIAMPolicyArgs{}
	}
	var resource SubscriptionIAMPolicy
	err := ctx.RegisterResource("gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionIAMPolicy gets an existing SubscriptionIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionIAMPolicyState, opts ...pulumi.ResourceOption) (*SubscriptionIAMPolicy, error) {
	var resource SubscriptionIAMPolicy
	err := ctx.ReadResource("gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionIAMPolicy resources.
type subscriptionIAMPolicyState struct {
	// (Computed) The etag of the subscription's IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription *string `pulumi:"subscription"`
}

type SubscriptionIAMPolicyState struct {
	// (Computed) The etag of the subscription's IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringPtrInput
}

func (SubscriptionIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMPolicyState)(nil)).Elem()
}

type subscriptionIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription string `pulumi:"subscription"`
}

// The set of arguments for constructing a SubscriptionIAMPolicy resource.
type SubscriptionIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringInput
}

func (SubscriptionIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMPolicyArgs)(nil)).Elem()
}
