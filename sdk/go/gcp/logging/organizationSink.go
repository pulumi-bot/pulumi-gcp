// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package logging

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type OrganizationSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions OrganizationSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	// {{% examples %}}
	// {{% /examples %}}
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren pulumi.BoolPtrOutput `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name pulumi.StringOutput `pulumi:"name"`
	// The numeric ID of the organization to be exported to the sink.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewOrganizationSink registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSink(ctx *pulumi.Context,
	name string, args *OrganizationSinkArgs, opts ...pulumi.ResourceOption) (*OrganizationSink, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil {
		args = &OrganizationSinkArgs{}
	}
	var resource OrganizationSink
	err := ctx.RegisterResource("gcp:logging/organizationSink:OrganizationSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSink gets an existing OrganizationSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSinkState, opts ...pulumi.ResourceOption) (*OrganizationSink, error) {
	var resource OrganizationSink
	err := ctx.ReadResource("gcp:logging/organizationSink:OrganizationSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSink resources.
type organizationSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *OrganizationSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	// {{% examples %}}
	// {{% /examples %}}
	// The writer associated with the sink must have access to write to the above resource.
	Destination *string `pulumi:"destination"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
	// The numeric ID of the organization to be exported to the sink.
	OrgId *string `pulumi:"orgId"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type OrganizationSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions OrganizationSinkBigqueryOptionsPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	// {{% examples %}}
	// {{% /examples %}}
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringPtrInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren pulumi.BoolPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
	// The numeric ID of the organization to be exported to the sink.
	OrgId pulumi.StringPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (OrganizationSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSinkState)(nil)).Elem()
}

type organizationSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *OrganizationSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	// {{% examples %}}
	// {{% /examples %}}
	// The writer associated with the sink must have access to write to the above resource.
	Destination string `pulumi:"destination"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
	// The numeric ID of the organization to be exported to the sink.
	OrgId string `pulumi:"orgId"`
}

// The set of arguments for constructing a OrganizationSink resource.
type OrganizationSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions OrganizationSinkBigqueryOptionsPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	// {{% examples %}}
	// {{% /examples %}}
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren pulumi.BoolPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
	// The numeric ID of the organization to be exported to the sink.
	OrgId pulumi.StringInput
}

func (OrganizationSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSinkArgs)(nil)).Elem()
}

