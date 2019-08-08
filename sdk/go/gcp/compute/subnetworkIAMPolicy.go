// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_subnetwork_iam_policy.html.markdown.
type SubnetworkIAMPolicy struct {
	s *pulumi.ResourceState
}

// NewSubnetworkIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewSubnetworkIAMPolicy(ctx *pulumi.Context,
	name string, args *SubnetworkIAMPolicyArgs, opts ...pulumi.ResourceOpt) (*SubnetworkIAMPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.Subnetwork == nil {
		return nil, errors.New("missing required argument 'Subnetwork'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["policyData"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["subnetwork"] = nil
	} else {
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["subnetwork"] = args.Subnetwork
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetworkIAMPolicy{s: s}, nil
}

// GetSubnetworkIAMPolicy gets an existing SubnetworkIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetworkIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetworkIAMPolicyState, opts ...pulumi.ResourceOpt) (*SubnetworkIAMPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["subnetwork"] = state.Subnetwork
	}
	s, err := ctx.ReadResource("gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetworkIAMPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SubnetworkIAMPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SubnetworkIAMPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the subnetwork's IAM policy.
func (r *SubnetworkIAMPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (r *SubnetworkIAMPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *SubnetworkIAMPolicy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The region of the subnetwork. If
// unspecified, this defaults to the region configured in the provider.
func (r *SubnetworkIAMPolicy) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The name of the subnetwork.
func (r *SubnetworkIAMPolicy) Subnetwork() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetwork"])
}

// Input properties used for looking up and filtering SubnetworkIAMPolicy resources.
type SubnetworkIAMPolicyState struct {
	// (Computed) The etag of the subnetwork's IAM policy.
	Etag interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region of the subnetwork. If
	// unspecified, this defaults to the region configured in the provider.
	Region interface{}
	// The name of the subnetwork.
	Subnetwork interface{}
}

// The set of arguments for constructing a SubnetworkIAMPolicy resource.
type SubnetworkIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region of the subnetwork. If
	// unspecified, this defaults to the region configured in the provider.
	Region interface{}
	// The name of the subnetwork.
	Subnetwork interface{}
}
