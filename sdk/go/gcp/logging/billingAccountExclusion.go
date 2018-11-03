// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a billing account logging exclusion. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/) and
// [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
// 
// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
// granted to the credentials used with Terraform.
type BillingAccountExclusion struct {
	s *pulumi.ResourceState
}

// NewBillingAccountExclusion registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountExclusion(ctx *pulumi.Context,
	name string, args *BillingAccountExclusionArgs, opts ...pulumi.ResourceOpt) (*BillingAccountExclusion, error) {
	if args == nil || args.BillingAccount == nil {
		return nil, errors.New("missing required argument 'BillingAccount'")
	}
	if args == nil || args.Filter == nil {
		return nil, errors.New("missing required argument 'Filter'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["billingAccount"] = nil
		inputs["description"] = nil
		inputs["disabled"] = nil
		inputs["filter"] = nil
		inputs["name"] = nil
	} else {
		inputs["billingAccount"] = args.BillingAccount
		inputs["description"] = args.Description
		inputs["disabled"] = args.Disabled
		inputs["filter"] = args.Filter
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BillingAccountExclusion{s: s}, nil
}

// GetBillingAccountExclusion gets an existing BillingAccountExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountExclusion(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BillingAccountExclusionState, opts ...pulumi.ResourceOpt) (*BillingAccountExclusion, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["billingAccount"] = state.BillingAccount
		inputs["description"] = state.Description
		inputs["disabled"] = state.Disabled
		inputs["filter"] = state.Filter
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BillingAccountExclusion{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BillingAccountExclusion) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BillingAccountExclusion) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The billing account to create the exclusion for.
func (r *BillingAccountExclusion) BillingAccount() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["billingAccount"])
}

// A human-readable description.
func (r *BillingAccountExclusion) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Whether this exclusion rule should be disabled or not. This defaults to
// false.
func (r *BillingAccountExclusion) Disabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disabled"])
}

// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
// write a filter.
func (r *BillingAccountExclusion) Filter() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["filter"])
}

// The name of the logging exclusion.
func (r *BillingAccountExclusion) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering BillingAccountExclusion resources.
type BillingAccountExclusionState struct {
	// The billing account to create the exclusion for.
	BillingAccount interface{}
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
}

// The set of arguments for constructing a BillingAccountExclusion resource.
type BillingAccountExclusionArgs struct {
	// The billing account to create the exclusion for.
	BillingAccount interface{}
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
}
