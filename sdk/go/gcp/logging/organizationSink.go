// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a organization-level logging sink. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/) and
// [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
//
// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
// granted to the credentials used with this provider.
type OrganizationSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions OrganizationSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions OrganizationSinkExclusionArrayOutput `pulumi:"exclusions"`
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
	Destination *string `pulumi:"destination"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions []OrganizationSinkExclusion `pulumi:"exclusions"`
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
	Destination pulumi.StringPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions OrganizationSinkExclusionArrayInput
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
	Destination string `pulumi:"destination"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions []OrganizationSinkExclusion `pulumi:"exclusions"`
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
	Destination pulumi.StringInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
	// one of exclusion_filters it will not be exported.
	Exclusions OrganizationSinkExclusionArrayInput
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

type OrganizationSinkInput interface {
	pulumi.Input

	ToOrganizationSinkOutput() OrganizationSinkOutput
	ToOrganizationSinkOutputWithContext(ctx context.Context) OrganizationSinkOutput
}

func (OrganizationSink) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSink)(nil)).Elem()
}

func (i OrganizationSink) ToOrganizationSinkOutput() OrganizationSinkOutput {
	return i.ToOrganizationSinkOutputWithContext(context.Background())
}

func (i OrganizationSink) ToOrganizationSinkOutputWithContext(ctx context.Context) OrganizationSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkOutput)
}

type OrganizationSinkOutput struct {
	*pulumi.OutputState
}

func (OrganizationSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSinkOutput)(nil)).Elem()
}

func (o OrganizationSinkOutput) ToOrganizationSinkOutput() OrganizationSinkOutput {
	return o
}

func (o OrganizationSinkOutput) ToOrganizationSinkOutputWithContext(ctx context.Context) OrganizationSinkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationSinkOutput{})
}
