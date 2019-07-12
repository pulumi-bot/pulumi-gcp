// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_tunnel_instance_iam_policy.html.markdown.
type TunnelInstanceIAMPolicy struct {
	s *pulumi.ResourceState
}

// NewTunnelInstanceIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewTunnelInstanceIAMPolicy(ctx *pulumi.Context,
	name string, args *TunnelInstanceIAMPolicyArgs, opts ...pulumi.ResourceOpt) (*TunnelInstanceIAMPolicy, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["instance"] = nil
		inputs["policyData"] = nil
		inputs["project"] = nil
		inputs["zone"] = nil
	} else {
		inputs["instance"] = args.Instance
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
		inputs["zone"] = args.Zone
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:iap/tunnelInstanceIAMPolicy:TunnelInstanceIAMPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TunnelInstanceIAMPolicy{s: s}, nil
}

// GetTunnelInstanceIAMPolicy gets an existing TunnelInstanceIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTunnelInstanceIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TunnelInstanceIAMPolicyState, opts ...pulumi.ResourceOpt) (*TunnelInstanceIAMPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["instance"] = state.Instance
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:iap/tunnelInstanceIAMPolicy:TunnelInstanceIAMPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TunnelInstanceIAMPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TunnelInstanceIAMPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TunnelInstanceIAMPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the instance's IAM policy.
func (r *TunnelInstanceIAMPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The name of the instance.
func (r *TunnelInstanceIAMPolicy) Instance() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instance"])
}

// The policy data generated by
// a `google_iam_policy` data source.
func (r *TunnelInstanceIAMPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *TunnelInstanceIAMPolicy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The zone of the instance. If
// unspecified, this defaults to the zone configured in the provider.
func (r *TunnelInstanceIAMPolicy) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering TunnelInstanceIAMPolicy resources.
type TunnelInstanceIAMPolicyState struct {
	// (Computed) The etag of the instance's IAM policy.
	Etag interface{}
	// The name of the instance.
	Instance interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The zone of the instance. If
	// unspecified, this defaults to the zone configured in the provider.
	Zone interface{}
}

// The set of arguments for constructing a TunnelInstanceIAMPolicy resource.
type TunnelInstanceIAMPolicyArgs struct {
	// The name of the instance.
	Instance interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The zone of the instance. If
	// unspecified, this defaults to the zone configured in the provider.
	Zone interface{}
}
