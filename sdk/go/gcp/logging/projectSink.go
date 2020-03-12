// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package logging

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ProjectSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
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
	// The name of the logging sink.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrOutput `pulumi:"uniqueWriterIdentity"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewProjectSink registers a new resource with the given unique name, arguments, and options.
func NewProjectSink(ctx *pulumi.Context,
	name string, args *ProjectSinkArgs, opts ...pulumi.ResourceOption) (*ProjectSink, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	if args == nil {
		args = &ProjectSinkArgs{}
	}
	var resource ProjectSink
	err := ctx.RegisterResource("gcp:logging/projectSink:ProjectSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectSink gets an existing ProjectSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectSinkState, opts ...pulumi.ResourceOption) (*ProjectSink, error) {
	var resource ProjectSink
	err := ctx.ReadResource("gcp:logging/projectSink:ProjectSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectSink resources.
type projectSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *ProjectSinkBigqueryOptions `pulumi:"bigqueryOptions"`
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
	// The name of the logging sink.
	Name *string `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity *bool `pulumi:"uniqueWriterIdentity"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type ProjectSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsPtrInput
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
	// The name of the logging sink.
	Name pulumi.StringPtrInput
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (ProjectSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSinkState)(nil)).Elem()
}

type projectSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *ProjectSinkBigqueryOptions `pulumi:"bigqueryOptions"`
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
	// The name of the logging sink.
	Name *string `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity *bool `pulumi:"uniqueWriterIdentity"`
}

// The set of arguments for constructing a ProjectSink resource.
type ProjectSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsPtrInput
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
	// The name of the logging sink.
	Name pulumi.StringPtrInput
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrInput
}

func (ProjectSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSinkArgs)(nil)).Elem()
}

