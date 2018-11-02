// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource to configure permissions for who can edit the service account. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the google_project_iam set of resources.
// 
// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
// 
// * `google_service_account_iam_policy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
// * `google_service_account_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
// * `google_service_account_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
// 
// ~> **Note:** `google_service_account_iam_policy` **cannot** be used in conjunction with `google_service_account_iam_binding` and `google_service_account_iam_member` or they will fight over what your policy should be.
// 
// ~> **Note:** `google_service_account_iam_binding` resources **can be** used in conjunction with `google_service_account_iam_member` resources **only if** they do not grant privilege to the same role.
type IAMMember struct {
	s *pulumi.ResourceState
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOpt) (*IAMMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServiceAccountId == nil {
		return nil, errors.New("missing required argument 'ServiceAccountId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["member"] = nil
		inputs["role"] = nil
		inputs["serviceAccountId"] = nil
	} else {
		inputs["member"] = args.Member
		inputs["role"] = args.Role
		inputs["serviceAccountId"] = args.ServiceAccountId
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:serviceAccount/iAMMember:IAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMMember{s: s}, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMMemberState, opts ...pulumi.ResourceOpt) (*IAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["role"] = state.Role
		inputs["serviceAccountId"] = state.ServiceAccountId
	}
	s, err := ctx.ReadResource("gcp:serviceAccount/iAMMember:IAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the service account IAM policy.
func (r *IAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *IAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The role that should be applied. Only one
// `google_service_account_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// The service account id to apply policy to.
func (r *IAMMember) ServiceAccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceAccountId"])
}

// Input properties used for looking up and filtering IAMMember resources.
type IAMMemberState struct {
	// (Computed) The etag of the service account IAM policy.
	Etag interface{}
	Member interface{}
	// The role that should be applied. Only one
	// `google_service_account_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The service account id to apply policy to.
	ServiceAccountId interface{}
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	Member interface{}
	// The role that should be applied. Only one
	// `google_service_account_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The service account id to apply policy to.
	ServiceAccountId interface{}
}
