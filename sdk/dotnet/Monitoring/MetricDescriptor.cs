// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// Defines a metric type and its schema. Once a metric descriptor is created, deleting or altering it stops data collection and makes the metric type's existing data unusable.
    /// 
    /// To get more information about MetricDescriptor, see:
    /// 
    /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/monitoring/custom-metrics/)
    /// 
    /// ## Example Usage
    /// ### Monitoring Metric Descriptor Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var basic = new Gcp.Monitoring.MetricDescriptor("basic", new Gcp.Monitoring.MetricDescriptorArgs
    ///         {
    ///             Description = "Daily sales records from all branch stores.",
    ///             DisplayName = "metric-descriptor",
    ///             Labels = 
    ///             {
    ///                 new Gcp.Monitoring.Inputs.MetricDescriptorLabelArgs
    ///                 {
    ///                     Description = "The ID of the store.",
    ///                     Key = "store_id",
    ///                     ValueType = "STRING",
    ///                 },
    ///             },
    ///             LaunchStage = "BETA",
    ///             Metadata = new Gcp.Monitoring.Inputs.MetricDescriptorMetadataArgs
    ///             {
    ///                 IngestDelay = "30s",
    ///                 SamplePeriod = "60s",
    ///             },
    ///             MetricKind = "GAUGE",
    ///             Type = "custom.googleapis.com/stores/daily_sales",
    ///             Unit = "{USD}",
    ///             ValueType = "DOUBLE",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Monitoring Metric Descriptor Alert
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var withAlert = new Gcp.Monitoring.MetricDescriptor("withAlert", new Gcp.Monitoring.MetricDescriptorArgs
    ///         {
    ///             Description = "Daily sales records from all branch stores.",
    ///             DisplayName = "metric-descriptor",
    ///             MetricKind = "GAUGE",
    ///             Type = "custom.googleapis.com/stores/daily_sales",
    ///             Unit = "{USD}",
    ///             ValueType = "DOUBLE",
    ///         });
    ///         var alertPolicy = new Gcp.Monitoring.AlertPolicy("alertPolicy", new Gcp.Monitoring.AlertPolicyArgs
    ///         {
    ///             Combiner = "OR",
    ///             Conditions = 
    ///             {
    ///                 new Gcp.Monitoring.Inputs.AlertPolicyConditionArgs
    ///                 {
    ///                     ConditionThreshold = new Gcp.Monitoring.Inputs.AlertPolicyConditionConditionThresholdArgs
    ///                     {
    ///                         Comparison = "COMPARISON_GT",
    ///                         Duration = "60s",
    ///                         Filter = withAlert.Type.Apply(type =&gt; $"metric.type=\"{type}\" AND resource.type=\"gce_instance\""),
    ///                     },
    ///                     DisplayName = "test condition",
    ///                 },
    ///             },
    ///             DisplayName = "metric-descriptor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// MetricDescriptor can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:monitoring/metricDescriptor:MetricDescriptor default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:monitoring/metricDescriptor:MetricDescriptor")]
    public partial class MetricDescriptor : Pulumi.CustomResource
    {
        /// <summary>
        /// A human-readable description for the label.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
        /// Structure is documented below.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.MetricDescriptorLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// The launch stage of the metric definition.
        /// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
        /// </summary>
        [Output("launchStage")]
        public Output<string?> LaunchStage { get; private set; } = null!;

        /// <summary>
        /// Metadata which can be used to guide usage of the metric.
        /// Structure is documented below.
        /// </summary>
        [Output("metadata")]
        public Output<Outputs.MetricDescriptorMetadata?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
        /// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
        /// </summary>
        [Output("metricKind")]
        public Output<string> MetricKind { get; private set; } = null!;

        /// <summary>
        /// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
        /// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
        /// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
        /// this list.
        /// </summary>
        [Output("monitoredResourceTypes")]
        public Output<ImmutableArray<string>> MonitoredResourceTypes { get; private set; } = null!;

        /// <summary>
        /// The resource name of the metric descriptor.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The units in which the metric value is reported. It is only applicable if the
        /// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
        /// the stored metric values.
        /// Different systems may scale the values to be more easily displayed (so a value of
        /// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
        /// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
        /// thousands of bytes, no matter how it may be displayed.
        /// If you want a custom metric to record the exact number of CPU-seconds used by a job,
        /// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
        /// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
        /// 12005.
        /// Alternatively, if you want a custom metric to record data in a more granular way, you
        /// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
        /// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
        /// The supported units are a subset of The Unified Code for Units of Measure standard.
        /// More info can be found in the API documentation
        /// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
        /// </summary>
        [Output("unit")]
        public Output<string?> Unit { get; private set; } = null!;

        /// <summary>
        /// The type of data that can be assigned to the label.
        /// Default value is `STRING`.
        /// Possible values are `STRING`, `BOOL`, and `INT64`.
        /// </summary>
        [Output("valueType")]
        public Output<string> ValueType { get; private set; } = null!;


        /// <summary>
        /// Create a MetricDescriptor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetricDescriptor(string name, MetricDescriptorArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/metricDescriptor:MetricDescriptor", name, args ?? new MetricDescriptorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetricDescriptor(string name, Input<string> id, MetricDescriptorState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/metricDescriptor:MetricDescriptor", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MetricDescriptor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetricDescriptor Get(string name, Input<string> id, MetricDescriptorState? state = null, CustomResourceOptions? options = null)
        {
            return new MetricDescriptor(name, id, state, options);
        }
    }

    public sealed class MetricDescriptorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description for the label.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("labels")]
        private InputList<Inputs.MetricDescriptorLabelArgs>? _labels;

        /// <summary>
        /// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.MetricDescriptorLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.MetricDescriptorLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The launch stage of the metric definition.
        /// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
        /// </summary>
        [Input("launchStage")]
        public Input<string>? LaunchStage { get; set; }

        /// <summary>
        /// Metadata which can be used to guide usage of the metric.
        /// Structure is documented below.
        /// </summary>
        [Input("metadata")]
        public Input<Inputs.MetricDescriptorMetadataArgs>? Metadata { get; set; }

        /// <summary>
        /// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
        /// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
        /// </summary>
        [Input("metricKind", required: true)]
        public Input<string> MetricKind { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The units in which the metric value is reported. It is only applicable if the
        /// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
        /// the stored metric values.
        /// Different systems may scale the values to be more easily displayed (so a value of
        /// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
        /// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
        /// thousands of bytes, no matter how it may be displayed.
        /// If you want a custom metric to record the exact number of CPU-seconds used by a job,
        /// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
        /// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
        /// 12005.
        /// Alternatively, if you want a custom metric to record data in a more granular way, you
        /// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
        /// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
        /// The supported units are a subset of The Unified Code for Units of Measure standard.
        /// More info can be found in the API documentation
        /// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// The type of data that can be assigned to the label.
        /// Default value is `STRING`.
        /// Possible values are `STRING`, `BOOL`, and `INT64`.
        /// </summary>
        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        public MetricDescriptorArgs()
        {
        }
    }

    public sealed class MetricDescriptorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description for the label.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputList<Inputs.MetricDescriptorLabelGetArgs>? _labels;

        /// <summary>
        /// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.MetricDescriptorLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.MetricDescriptorLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The launch stage of the metric definition.
        /// Possible values are `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
        /// </summary>
        [Input("launchStage")]
        public Input<string>? LaunchStage { get; set; }

        /// <summary>
        /// Metadata which can be used to guide usage of the metric.
        /// Structure is documented below.
        /// </summary>
        [Input("metadata")]
        public Input<Inputs.MetricDescriptorMetadataGetArgs>? Metadata { get; set; }

        /// <summary>
        /// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
        /// Possible values are `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, and `CUMULATIVE`.
        /// </summary>
        [Input("metricKind")]
        public Input<string>? MetricKind { get; set; }

        [Input("monitoredResourceTypes")]
        private InputList<string>? _monitoredResourceTypes;

        /// <summary>
        /// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
        /// is associated with this metric type can only be associated with one of the monitored resource types listed here. This
        /// field allows time series to be associated with the intersection of this metric type and the monitored resource types in
        /// this list.
        /// </summary>
        public InputList<string> MonitoredResourceTypes
        {
            get => _monitoredResourceTypes ?? (_monitoredResourceTypes = new InputList<string>());
            set => _monitoredResourceTypes = value;
        }

        /// <summary>
        /// The resource name of the metric descriptor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The units in which the metric value is reported. It is only applicable if the
        /// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
        /// the stored metric values.
        /// Different systems may scale the values to be more easily displayed (so a value of
        /// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
        /// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
        /// thousands of bytes, no matter how it may be displayed.
        /// If you want a custom metric to record the exact number of CPU-seconds used by a job,
        /// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
        /// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
        /// 12005.
        /// Alternatively, if you want a custom metric to record data in a more granular way, you
        /// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
        /// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
        /// The supported units are a subset of The Unified Code for Units of Measure standard.
        /// More info can be found in the API documentation
        /// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// The type of data that can be assigned to the label.
        /// Default value is `STRING`.
        /// Possible values are `STRING`, `BOOL`, and `INT64`.
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        public MetricDescriptorState()
        {
        }
    }
}
