// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a project-level logging exclusion. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/) and
// [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
//
// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
// granted to the credentials used with this provider.
type ProjectExclusion struct {
	pulumi.CustomResourceState

	// A human-readable description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// The name of the logging exclusion.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectExclusion registers a new resource with the given unique name, arguments, and options.
func NewProjectExclusion(ctx *pulumi.Context,
	name string, args *ProjectExclusionArgs, opts ...pulumi.ResourceOption) (*ProjectExclusion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	var resource ProjectExclusion
	err := ctx.RegisterResource("gcp:logging/projectExclusion:ProjectExclusion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectExclusion gets an existing ProjectExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectExclusionState, opts ...pulumi.ResourceOption) (*ProjectExclusion, error) {
	var resource ProjectExclusion
	err := ctx.ReadResource("gcp:logging/projectExclusion:ProjectExclusion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectExclusion resources.
type projectExclusionState struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
}

type ProjectExclusionState struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
}

func (ProjectExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectExclusionState)(nil)).Elem()
}

type projectExclusionArgs struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter string `pulumi:"filter"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectExclusion resource.
type ProjectExclusionArgs struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
}

func (ProjectExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectExclusionArgs)(nil)).Elem()
}
