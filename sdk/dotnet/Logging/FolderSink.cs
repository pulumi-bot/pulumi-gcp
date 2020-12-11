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
    /// Manages a folder-level logging sink. For more information see
    /// [the official documentation](https://cloud.google.com/logging/docs/) and
    /// [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
    /// 
    /// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    /// granted to the credentials used with this provider.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var log_bucket = new Gcp.Storage.Bucket("log-bucket", new Gcp.Storage.BucketArgs
    ///         {
    ///         });
    ///         var my_folder = new Gcp.Organizations.Folder("my-folder", new Gcp.Organizations.FolderArgs
    ///         {
    ///             DisplayName = "My folder",
    ///             Parent = "organizations/123456",
    ///         });
    ///         var my_sink = new Gcp.Logging.FolderSink("my-sink", new Gcp.Logging.FolderSinkArgs
    ///         {
    ///             Description = "some explaination on what this is",
    ///             Folder = my_folder.Name,
    ///             Destination = log_bucket.Name.Apply(name =&gt; $"storage.googleapis.com/{name}"),
    ///             Filter = "resource.type = gce_instance AND severity &gt;= WARNING",
    ///         });
    ///         var log_writer = new Gcp.Projects.IAMBinding("log-writer", new Gcp.Projects.IAMBindingArgs
    ///         {
    ///             Role = "roles/storage.objectCreator",
    ///             Members = 
    ///             {
    ///                 my_sink.WriterIdentity,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Folder-level logging sinks can be imported using this format
    /// 
    /// ```sh
    ///  $ pulumi import gcp:logging/folderSink:FolderSink my_sink folders/{{folder_id}}/sinks/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:logging/folderSink:FolderSink")]
    public partial class FolderSink : Pulumi.CustomResource
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Output("bigqueryOptions")]
        public Output<Outputs.FolderSinkBigqueryOptions> BigqueryOptions { get; private set; } = null!;

        /// <summary>
        /// A description of this sink. The maximum length of the description is 8000 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// If set to True, then this sink is disabled and it does not export any log entries.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
        /// one of exclusion_filters it will not be exported.
        /// </summary>
        [Output("exclusions")]
        public Output<ImmutableArray<Outputs.FolderSinkExclusion>> Exclusions { get; private set; } = null!;

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Output("filter")]
        public Output<string?> Filter { get; private set; } = null!;

        /// <summary>
        /// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
        /// accepted.
        /// </summary>
        [Output("folder")]
        public Output<string> Folder { get; private set; } = null!;

        /// <summary>
        /// Whether or not to include children folders in the sink export. If true, logs
        /// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
        /// </summary>
        [Output("includeChildren")]
        public Output<bool?> IncludeChildren { get; private set; } = null!;

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Output("writerIdentity")]
        public Output<string> WriterIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a FolderSink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FolderSink(string name, FolderSinkArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/folderSink:FolderSink", name, args ?? new FolderSinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FolderSink(string name, Input<string> id, FolderSinkState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/folderSink:FolderSink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FolderSink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FolderSink Get(string name, Input<string> id, FolderSinkState? state = null, CustomResourceOptions? options = null)
        {
            return new FolderSink(name, id, state, options);
        }
    }

    public sealed class FolderSinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.FolderSinkBigqueryOptionsArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// A description of this sink. The maximum length of the description is 8000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// If set to True, then this sink is disabled and it does not export any log entries.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("exclusions")]
        private InputList<Inputs.FolderSinkExclusionArgs>? _exclusions;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
        /// one of exclusion_filters it will not be exported.
        /// </summary>
        public InputList<Inputs.FolderSinkExclusionArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.FolderSinkExclusionArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
        /// accepted.
        /// </summary>
        [Input("folder", required: true)]
        public Input<string> Folder { get; set; } = null!;

        /// <summary>
        /// Whether or not to include children folders in the sink export. If true, logs
        /// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
        /// </summary>
        [Input("includeChildren")]
        public Input<bool>? IncludeChildren { get; set; }

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FolderSinkArgs()
        {
        }
    }

    public sealed class FolderSinkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.FolderSinkBigqueryOptionsGetArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// A description of this sink. The maximum length of the description is 8000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// If set to True, then this sink is disabled and it does not export any log entries.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("exclusions")]
        private InputList<Inputs.FolderSinkExclusionGetArgs>? _exclusions;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
        /// one of exclusion_filters it will not be exported.
        /// </summary>
        public InputList<Inputs.FolderSinkExclusionGetArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.FolderSinkExclusionGetArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
        /// accepted.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// Whether or not to include children folders in the sink export. If true, logs
        /// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
        /// </summary>
        [Input("includeChildren")]
        public Input<bool>? IncludeChildren { get; set; }

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Input("writerIdentity")]
        public Input<string>? WriterIdentity { get; set; }

        public FolderSinkState()
        {
        }
    }
}
