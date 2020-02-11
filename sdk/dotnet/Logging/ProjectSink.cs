// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_project_sink.html.markdown.
    /// </summary>
    public partial class ProjectSink : Pulumi.CustomResource
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Output("bigqueryOptions")]
        public Output<Outputs.ProjectSinkBigqueryOptions> BigqueryOptions { get; private set; } = null!;

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Output("filter")]
        public Output<string?> Filter { get; private set; } = null!;

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to create the sink in. If omitted, the project associated with the provider is
        /// used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Whether or not to create a unique identity associated with this sink. If `false`
        /// (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
        /// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
        /// must set `unique_writer_identity` to true.
        /// </summary>
        [Output("uniqueWriterIdentity")]
        public Output<bool?> UniqueWriterIdentity { get; private set; } = null!;

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Output("writerIdentity")]
        public Output<string> WriterIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectSink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectSink(string name, ProjectSinkArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/projectSink:ProjectSink", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ProjectSink(string name, Input<string> id, ProjectSinkState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/projectSink:ProjectSink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectSink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectSink Get(string name, Input<string> id, ProjectSinkState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectSink(name, id, state, options);
        }
    }

    public sealed class ProjectSinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.ProjectSinkBigqueryOptionsArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project to create the sink in. If omitted, the project associated with the provider is
        /// used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Whether or not to create a unique identity associated with this sink. If `false`
        /// (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
        /// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
        /// must set `unique_writer_identity` to true.
        /// </summary>
        [Input("uniqueWriterIdentity")]
        public Input<bool>? UniqueWriterIdentity { get; set; }

        public ProjectSinkArgs()
        {
        }
    }

    public sealed class ProjectSinkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.ProjectSinkBigqueryOptionsGetArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project to create the sink in. If omitted, the project associated with the provider is
        /// used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Whether or not to create a unique identity associated with this sink. If `false`
        /// (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
        /// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
        /// must set `unique_writer_identity` to true.
        /// </summary>
        [Input("uniqueWriterIdentity")]
        public Input<bool>? UniqueWriterIdentity { get; set; }

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Input("writerIdentity")]
        public Input<string>? WriterIdentity { get; set; }

        public ProjectSinkState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ProjectSinkBigqueryOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("usePartitionedTables", required: true)]
        public Input<bool> UsePartitionedTables { get; set; } = null!;

        public ProjectSinkBigqueryOptionsArgs()
        {
        }
    }

    public sealed class ProjectSinkBigqueryOptionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("usePartitionedTables", required: true)]
        public Input<bool> UsePartitionedTables { get; set; } = null!;

        public ProjectSinkBigqueryOptionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ProjectSinkBigqueryOptions
    {
        public readonly bool UsePartitionedTables;

        [OutputConstructor]
        private ProjectSinkBigqueryOptions(bool usePartitionedTables)
        {
            UsePartitionedTables = usePartitionedTables;
        }
    }
    }
}
