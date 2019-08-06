// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sourcerepo

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for Pubsub Topic. Each of these resources serves a different use case:
// 
// * `pubsub.TopicIAMPolicy`: Authoritative. Sets the IAM policy for the topic and replaces any existing policy already attached.
// * `pubsub.TopicIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the topic are preserved.
// * `pubsub.TopicIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the topic are preserved.
// 
// > **Note:** `pubsub.TopicIAMPolicy` **cannot** be used in conjunction with `pubsub.TopicIAMBinding` and `pubsub.TopicIAMMember` or they will fight over what your policy should be.
// 
// > **Note:** `pubsub.TopicIAMBinding` resources **can be** used in conjunction with `pubsub.TopicIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/sourcerepo_repository_iam_member.html.markdown.
type RepositoryIamMember struct {
	s *pulumi.ResourceState
}

// NewRepositoryIamMember registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamMember(ctx *pulumi.Context,
	name string, args *RepositoryIamMemberArgs, opts ...pulumi.ResourceOpt) (*RepositoryIamMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["repository"] = nil
		inputs["role"] = nil
	} else {
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["repository"] = args.Repository
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:sourcerepo/repositoryIamMember:RepositoryIamMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RepositoryIamMember{s: s}, nil
}

// GetRepositoryIamMember gets an existing RepositoryIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RepositoryIamMemberState, opts ...pulumi.ResourceOpt) (*RepositoryIamMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["repository"] = state.Repository
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:sourcerepo/repositoryIamMember:RepositoryIamMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RepositoryIamMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RepositoryIamMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RepositoryIamMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the topic's IAM policy.
func (r *RepositoryIamMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *RepositoryIamMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *RepositoryIamMember) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *RepositoryIamMember) Repository() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["repository"])
}

// The role that should be applied. Only one
// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *RepositoryIamMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering RepositoryIamMember resources.
type RepositoryIamMemberState struct {
	// (Computed) The etag of the topic's IAM policy.
	Etag interface{}
	Member interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	Repository interface{}
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a RepositoryIamMember resource.
type RepositoryIamMemberArgs struct {
	Member interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	Repository interface{}
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
