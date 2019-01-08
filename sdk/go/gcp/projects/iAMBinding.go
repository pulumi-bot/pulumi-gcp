// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
// 
// * `google_project_iam_policy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
// * `google_project_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
// * `google_project_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
// 
// > **Note:** `google_project_iam_policy` **cannot** be used in conjunction with `google_project_iam_binding` and `google_project_iam_member` or they will fight over what your policy should be.
// 
// > **Note:** `google_project_iam_binding` resources **can be** used in conjunction with `google_project_iam_member` resources **only if** they do not grant privilege to the same role.
type IAMBinding struct {
	s *pulumi.ResourceState
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOpt) (*IAMBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["members"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
	} else {
		inputs["members"] = args.Members
		inputs["project"] = args.Project
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:projects/iAMBinding:IAMBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMBinding{s: s}, nil
}

// GetIAMBinding gets an existing IAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMBindingState, opts ...pulumi.ResourceOpt) (*IAMBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["project"] = state.Project
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:projects/iAMBinding:IAMBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the project's IAM policy.
func (r *IAMBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *IAMBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The project ID. If not specified, uses the
// ID of the project configured with the provider.
func (r *IAMBinding) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IAMBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering IAMBinding resources.
type IAMBindingState struct {
	// (Computed) The etag of the project's IAM policy.
	Etag interface{}
	Members interface{}
	// The project ID. If not specified, uses the
	// ID of the project configured with the provider.
	Project interface{}
	// The role that should be applied. Only one
	// `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	Members interface{}
	// The project ID. If not specified, uses the
	// ID of the project configured with the provider.
	Project interface{}
	// The role that should be applied. Only one
	// `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
