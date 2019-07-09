// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type OrganizationExclusion struct {
	s *pulumi.ResourceState
}

// NewOrganizationExclusion registers a new resource with the given unique name, arguments, and options.
func NewOrganizationExclusion(ctx *pulumi.Context,
	name string, args *OrganizationExclusionArgs, opts ...pulumi.ResourceOpt) (*OrganizationExclusion, error) {
	if args == nil || args.Filter == nil {
		return nil, errors.New("missing required argument 'Filter'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["disabled"] = nil
		inputs["filter"] = nil
		inputs["name"] = nil
		inputs["orgId"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["disabled"] = args.Disabled
		inputs["filter"] = args.Filter
		inputs["name"] = args.Name
		inputs["orgId"] = args.OrgId
	}
	s, err := ctx.RegisterResource("gcp:logging/organizationExclusion:OrganizationExclusion", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationExclusion{s: s}, nil
}

// GetOrganizationExclusion gets an existing OrganizationExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationExclusion(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OrganizationExclusionState, opts ...pulumi.ResourceOpt) (*OrganizationExclusion, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["disabled"] = state.Disabled
		inputs["filter"] = state.Filter
		inputs["name"] = state.Name
		inputs["orgId"] = state.OrgId
	}
	s, err := ctx.ReadResource("gcp:logging/organizationExclusion:OrganizationExclusion", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationExclusion{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OrganizationExclusion) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OrganizationExclusion) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A human-readable description.
func (r *OrganizationExclusion) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Whether this exclusion rule should be disabled or not. This defaults to
// false.
func (r *OrganizationExclusion) Disabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disabled"])
}

// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
// write a filter.
func (r *OrganizationExclusion) Filter() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["filter"])
}

// The name of the logging exclusion.
func (r *OrganizationExclusion) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The organization to create the exclusion in.
func (r *OrganizationExclusion) OrgId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["orgId"])
}

// Input properties used for looking up and filtering OrganizationExclusion resources.
type OrganizationExclusionState struct {
	// A human-readable description.
	Description interface{}
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled interface{}
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter interface{}
	// The name of the logging exclusion.
	Name interface{}
	// The organization to create the exclusion in.
	OrgId interface{}
}

// The set of arguments for constructing a OrganizationExclusion resource.
type OrganizationExclusionArgs struct {
	// A human-readable description.
	Description interface{}
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled interface{}
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter interface{}
	// The name of the logging exclusion.
	Name interface{}
	// The organization to create the exclusion in.
	OrgId interface{}
}
