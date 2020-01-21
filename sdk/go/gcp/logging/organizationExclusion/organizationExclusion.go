// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package organizationExclusion

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an organization-level logging exclusion. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/) and
// [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
// 
// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
// granted to the credentials used with this provider.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_organization_exclusion.html.markdown.
type OrganizationExclusion struct {
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
	// The organization to create the exclusion in.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
}

// NewOrganizationExclusion registers a new resource with the given unique name, arguments, and options.
func NewOrganizationExclusion(ctx *pulumi.Context,
	name string, args *OrganizationExclusionArgs, opts ...pulumi.ResourceOption) (*OrganizationExclusion, error) {
	if args == nil || args.Filter == nil {
		return nil, errors.New("missing required argument 'Filter'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil {
		args = &OrganizationExclusionArgs{}
	}
	var resource OrganizationExclusion
	err := ctx.RegisterResource("gcp:logging/organizationExclusion:OrganizationExclusion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationExclusion gets an existing OrganizationExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationExclusionState, opts ...pulumi.ResourceOption) (*OrganizationExclusion, error) {
	var resource OrganizationExclusion
	err := ctx.ReadResource("gcp:logging/organizationExclusion:OrganizationExclusion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationExclusion resources.
type organizationExclusionState struct {
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
	// The organization to create the exclusion in.
	OrgId *string `pulumi:"orgId"`
}

type OrganizationExclusionState struct {
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
	// The organization to create the exclusion in.
	OrgId pulumi.StringPtrInput
}

func (OrganizationExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationExclusionState)(nil)).Elem()
}

type organizationExclusionArgs struct {
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
	// The organization to create the exclusion in.
	OrgId string `pulumi:"orgId"`
}

// The set of arguments for constructing a OrganizationExclusion resource.
type OrganizationExclusionArgs struct {
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
	// The organization to create the exclusion in.
	OrgId pulumi.StringInput
}

func (OrganizationExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationExclusionArgs)(nil)).Elem()
}

