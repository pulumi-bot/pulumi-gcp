// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
// 
// * `google_pubsub_subscription_iam_policy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
// * `google_pubsub_subscription_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
// * `google_pubsub_subscription_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
// 
// > **Note:** `google_pubsub_subscription_iam_policy` **cannot** be used in conjunction with `google_pubsub_subscription_iam_binding` and `google_pubsub_subscription_iam_member` or they will fight over what your policy should be.
// 
// > **Note:** `google_pubsub_subscription_iam_binding` resources **can be** used in conjunction with `google_pubsub_subscription_iam_member` resources **only if** they do not grant privilege to the same role.
type SubscriptionIAMMember struct {
	s *pulumi.ResourceState
}

// NewSubscriptionIAMMember registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionIAMMember(ctx *pulumi.Context,
	name string, args *SubscriptionIAMMemberArgs, opts ...pulumi.ResourceOpt) (*SubscriptionIAMMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Subscription == nil {
		return nil, errors.New("missing required argument 'Subscription'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
		inputs["subscription"] = nil
	} else {
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["role"] = args.Role
		inputs["subscription"] = args.Subscription
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubscriptionIAMMember{s: s}, nil
}

// GetSubscriptionIAMMember gets an existing SubscriptionIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubscriptionIAMMemberState, opts ...pulumi.ResourceOpt) (*SubscriptionIAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["role"] = state.Role
		inputs["subscription"] = state.Subscription
	}
	s, err := ctx.ReadResource("gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubscriptionIAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SubscriptionIAMMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SubscriptionIAMMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the subscription's IAM policy.
func (r *SubscriptionIAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *SubscriptionIAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *SubscriptionIAMMember) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `google_pubsub_subscription_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *SubscriptionIAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// The subscription name or id to bind to attach IAM policy to.
func (r *SubscriptionIAMMember) Subscription() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subscription"])
}

// Input properties used for looking up and filtering SubscriptionIAMMember resources.
type SubscriptionIAMMemberState struct {
	// (Computed) The etag of the subscription's IAM policy.
	Etag interface{}
	Member interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_pubsub_subscription_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The subscription name or id to bind to attach IAM policy to.
	Subscription interface{}
}

// The set of arguments for constructing a SubscriptionIAMMember resource.
type SubscriptionIAMMemberArgs struct {
	Member interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_pubsub_subscription_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The subscription name or id to bind to attach IAM policy to.
	Subscription interface{}
}
