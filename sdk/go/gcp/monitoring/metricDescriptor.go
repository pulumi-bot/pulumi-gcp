// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Defines a metric type and its schema. Once a metric descriptor is created, deleting or altering it stops data collection and makes the metric type's existing data unusable.
//
// To get more information about MetricDescriptor, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/monitoring/custom-metrics/)
//
// ## Example Usage
type MetricDescriptor struct {
	pulumi.CustomResourceState

	// A human-readable description for the label.
	Description pulumi.StringOutput `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	Labels MetricDescriptorLabelArrayOutput `pulumi:"labels"`
	// The launch stage of the metric definition.
	// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage pulumi.StringPtrOutput `pulumi:"launchStage"`
	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	Metadata MetricDescriptorMetadataPtrOutput `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
	MetricKind pulumi.StringOutput `pulumi:"metricKind"`
	// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
	// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
	// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
	// this list.
	MonitoredResourceTypes pulumi.StringArrayOutput `pulumi:"monitoredResourceTypes"`
	// The resource name of the metric descriptor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relativeMetricName is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type pulumi.StringOutput `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit pulumi.StringPtrOutput `pulumi:"unit"`
	// The type of data that can be assigned to the label.
	// Default value is `STRING`.
	// Possible values are `STRING`, `BOOL`, and `INT64`.
	ValueType pulumi.StringOutput `pulumi:"valueType"`
}

// NewMetricDescriptor registers a new resource with the given unique name, arguments, and options.
func NewMetricDescriptor(ctx *pulumi.Context,
	name string, args *MetricDescriptorArgs, opts ...pulumi.ResourceOption) (*MetricDescriptor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MetricKind == nil {
		return nil, errors.New("invalid value for required argument 'MetricKind'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.ValueType == nil {
		return nil, errors.New("invalid value for required argument 'ValueType'")
	}
	var resource MetricDescriptor
	err := ctx.RegisterResource("gcp:monitoring/metricDescriptor:MetricDescriptor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricDescriptor gets an existing MetricDescriptor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricDescriptor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricDescriptorState, opts ...pulumi.ResourceOption) (*MetricDescriptor, error) {
	var resource MetricDescriptor
	err := ctx.ReadResource("gcp:monitoring/metricDescriptor:MetricDescriptor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricDescriptor resources.
type metricDescriptorState struct {
	// A human-readable description for the label.
	Description *string `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	DisplayName *string `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	Labels []MetricDescriptorLabel `pulumi:"labels"`
	// The launch stage of the metric definition.
	// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage *string `pulumi:"launchStage"`
	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	Metadata *MetricDescriptorMetadata `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
	MetricKind *string `pulumi:"metricKind"`
	// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
	// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
	// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
	// this list.
	MonitoredResourceTypes []string `pulumi:"monitoredResourceTypes"`
	// The resource name of the metric descriptor.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relativeMetricName is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type *string `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit *string `pulumi:"unit"`
	// The type of data that can be assigned to the label.
	// Default value is `STRING`.
	// Possible values are `STRING`, `BOOL`, and `INT64`.
	ValueType *string `pulumi:"valueType"`
}

type MetricDescriptorState struct {
	// A human-readable description for the label.
	Description pulumi.StringPtrInput
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	DisplayName pulumi.StringPtrInput
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	Labels MetricDescriptorLabelArrayInput
	// The launch stage of the metric definition.
	// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage pulumi.StringPtrInput
	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	Metadata MetricDescriptorMetadataPtrInput
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
	MetricKind pulumi.StringPtrInput
	// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
	// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
	// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
	// this list.
	MonitoredResourceTypes pulumi.StringArrayInput
	// The resource name of the metric descriptor.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relativeMetricName is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type pulumi.StringPtrInput
	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit pulumi.StringPtrInput
	// The type of data that can be assigned to the label.
	// Default value is `STRING`.
	// Possible values are `STRING`, `BOOL`, and `INT64`.
	ValueType pulumi.StringPtrInput
}

func (MetricDescriptorState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricDescriptorState)(nil)).Elem()
}

type metricDescriptorArgs struct {
	// A human-readable description for the label.
	Description string `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	DisplayName string `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	Labels []MetricDescriptorLabel `pulumi:"labels"`
	// The launch stage of the metric definition.
	// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage *string `pulumi:"launchStage"`
	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	Metadata *MetricDescriptorMetadata `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
	MetricKind string `pulumi:"metricKind"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relativeMetricName is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type string `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit *string `pulumi:"unit"`
	// The type of data that can be assigned to the label.
	// Default value is `STRING`.
	// Possible values are `STRING`, `BOOL`, and `INT64`.
	ValueType string `pulumi:"valueType"`
}

// The set of arguments for constructing a MetricDescriptor resource.
type MetricDescriptorArgs struct {
	// A human-readable description for the label.
	Description pulumi.StringInput
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	DisplayName pulumi.StringInput
	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	Labels MetricDescriptorLabelArrayInput
	// The launch stage of the metric definition.
	// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage pulumi.StringPtrInput
	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	Metadata MetricDescriptorMetadataPtrInput
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
	MetricKind pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relativeMetricName is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type pulumi.StringInput
	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit pulumi.StringPtrInput
	// The type of data that can be assigned to the label.
	// Default value is `STRING`.
	// Possible values are `STRING`, `BOOL`, and `INT64`.
	ValueType pulumi.StringInput
}

func (MetricDescriptorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricDescriptorArgs)(nil)).Elem()
}
