// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_config_iam_member.html.markdown.
type ConfigIamMember struct {
	s *pulumi.ResourceState
}

// NewConfigIamMember registers a new resource with the given unique name, arguments, and options.
func NewConfigIamMember(ctx *pulumi.Context,
	name string, args *ConfigIamMemberArgs, opts ...pulumi.ResourceOpt) (*ConfigIamMember, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["config"] = nil
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["config"] = args.Config
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:runtimeconfig/configIamMember:ConfigIamMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigIamMember{s: s}, nil
}

// GetConfigIamMember gets an existing ConfigIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigIamMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigIamMemberState, opts ...pulumi.ResourceOpt) (*ConfigIamMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["config"] = state.Config
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:runtimeconfig/configIamMember:ConfigIamMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigIamMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ConfigIamMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ConfigIamMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *ConfigIamMember) Condition() *pulumi.Output {
	return r.s.State["condition"]
}

// Used to find the parent resource to bind the IAM policy to
func (r *ConfigIamMember) Config() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["config"])
}

// (Computed) The etag of the IAM policy.
func (r *ConfigIamMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *ConfigIamMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *ConfigIamMember) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *ConfigIamMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering ConfigIamMember resources.
type ConfigIamMemberState struct {
	Condition interface{}
	// Used to find the parent resource to bind the IAM policy to
	Config interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	Member interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a ConfigIamMember resource.
type ConfigIamMemberArgs struct {
	Condition interface{}
	// Used to find the parent resource to bind the IAM policy to
	Config interface{}
	Member interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
