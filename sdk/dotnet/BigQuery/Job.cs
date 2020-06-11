// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// Jobs are actions that BigQuery runs on your behalf to load data, export data, query data, or copy data.
    /// Once a BigQuery job is created, it cannot be changed or deleted.
    /// 
    /// 
    /// 
    /// ## Example Usage
    /// 
    /// ### Bigquery Job Query
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bar = new Gcp.BigQuery.Dataset("bar", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "job_query_dataset",
    ///             FriendlyName = "test",
    ///             Description = "This is a test description",
    ///             Location = "US",
    ///         });
    ///         var foo = new Gcp.BigQuery.Table("foo", new Gcp.BigQuery.TableArgs
    ///         {
    ///             DatasetId = bar.DatasetId,
    ///             TableId = "job_query_table",
    ///         });
    ///         var job = new Gcp.BigQuery.Job("job", new Gcp.BigQuery.JobArgs
    ///         {
    ///             JobId = "job_query",
    ///             Labels = 
    ///             {
    ///                 { "example-label", "example-value" },
    ///             },
    ///             Query = new Gcp.BigQuery.Inputs.JobQueryArgs
    ///             {
    ///                 Query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
    ///                 Destination_table = 
    ///                 {
    ///                     { "projectId", foo.Project },
    ///                     { "datasetId", foo.DatasetId },
    ///                     { "tableId", foo.TableId },
    ///                 },
    ///                 AllowLargeResults = true,
    ///                 FlattenResults = true,
    ///                 Script_options = 
    ///                 {
    ///                     { "keyResultStatement", "LAST" },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Bigquery Job Query Table Reference
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bar = new Gcp.BigQuery.Dataset("bar", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "job_query_dataset",
    ///             FriendlyName = "test",
    ///             Description = "This is a test description",
    ///             Location = "US",
    ///         });
    ///         var foo = new Gcp.BigQuery.Table("foo", new Gcp.BigQuery.TableArgs
    ///         {
    ///             DatasetId = bar.DatasetId,
    ///             TableId = "job_query_table",
    ///         });
    ///         var job = new Gcp.BigQuery.Job("job", new Gcp.BigQuery.JobArgs
    ///         {
    ///             JobId = "job_query",
    ///             Labels = 
    ///             {
    ///                 { "example-label", "example-value" },
    ///             },
    ///             Query = new Gcp.BigQuery.Inputs.JobQueryArgs
    ///             {
    ///                 Query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
    ///                 Destination_table = 
    ///                 {
    ///                     { "tableId", foo.Id },
    ///                 },
    ///                 Default_dataset = 
    ///                 {
    ///                     { "datasetId", bar.Id },
    ///                 },
    ///                 AllowLargeResults = true,
    ///                 FlattenResults = true,
    ///                 Script_options = 
    ///                 {
    ///                     { "keyResultStatement", "LAST" },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Bigquery Job Load
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bar = new Gcp.BigQuery.Dataset("bar", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "job_load_dataset",
    ///             FriendlyName = "test",
    ///             Description = "This is a test description",
    ///             Location = "US",
    ///         });
    ///         var foo = new Gcp.BigQuery.Table("foo", new Gcp.BigQuery.TableArgs
    ///         {
    ///             DatasetId = bar.DatasetId,
    ///             TableId = "job_load_table",
    ///         });
    ///         var job = new Gcp.BigQuery.Job("job", new Gcp.BigQuery.JobArgs
    ///         {
    ///             JobId = "job_load",
    ///             Labels = 
    ///             {
    ///                 { "my_job", "load" },
    ///             },
    ///             Load = new Gcp.BigQuery.Inputs.JobLoadArgs
    ///             {
    ///                 SourceUris = 
    ///                 {
    ///                     "gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv",
    ///                 },
    ///                 Destination_table = 
    ///                 {
    ///                     { "projectId", foo.Project },
    ///                     { "datasetId", foo.DatasetId },
    ///                     { "tableId", foo.TableId },
    ///                 },
    ///                 SkipLeadingRows = 1,
    ///                 SchemaUpdateOptions = 
    ///                 {
    ///                     "ALLOW_FIELD_RELAXATION",
    ///                     "ALLOW_FIELD_ADDITION",
    ///                 },
    ///                 WriteDisposition = "WRITE_APPEND",
    ///                 Autodetect = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Bigquery Job Extract
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var source_oneDataset = new Gcp.BigQuery.Dataset("source-oneDataset", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "job_extract_dataset",
    ///             FriendlyName = "test",
    ///             Description = "This is a test description",
    ///             Location = "US",
    ///         });
    ///         var source_oneTable = new Gcp.BigQuery.Table("source-oneTable", new Gcp.BigQuery.TableArgs
    ///         {
    ///             DatasetId = source_oneDataset.DatasetId,
    ///             TableId = "job_extract_table",
    ///             Schema = @"[
    ///   {
    ///     ""name"": ""name"",
    ///     ""type"": ""STRING"",
    ///     ""mode"": ""NULLABLE""
    ///   },
    ///   {
    ///     ""name"": ""post_abbr"",
    ///     ""type"": ""STRING"",
    ///     ""mode"": ""NULLABLE""
    ///   },
    ///   {
    ///     ""name"": ""date"",
    ///     ""type"": ""DATE"",
    ///     ""mode"": ""NULLABLE""
    ///   }
    /// ]
    /// ",
    ///         });
    ///         var dest = new Gcp.Storage.Bucket("dest", new Gcp.Storage.BucketArgs
    ///         {
    ///             ForceDestroy = true,
    ///         });
    ///         var job = new Gcp.BigQuery.Job("job", new Gcp.BigQuery.JobArgs
    ///         {
    ///             JobId = "job_extract",
    ///             Extract = new Gcp.BigQuery.Inputs.JobExtractArgs
    ///             {
    ///                 DestinationUris = 
    ///                 {
    ///                     dest.Url.Apply(url =&gt; $"{url}/extract"),
    ///                 },
    ///                 Source_table = 
    ///                 {
    ///                     { "projectId", source_oneTable.Project },
    ///                     { "datasetId", source_oneTable.DatasetId },
    ///                     { "tableId", source_oneTable.TableId },
    ///                 },
    ///                 DestinationFormat = "NEWLINE_DELIMITED_JSON",
    ///                 Compression = "GZIP",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Job : Pulumi.CustomResource
    {
        /// <summary>
        /// Copies a table.  Structure is documented below.
        /// </summary>
        [Output("copy")]
        public Output<Outputs.JobCopy?> Copy { get; private set; } = null!;

        /// <summary>
        /// Configures an extract job.  Structure is documented below.
        /// </summary>
        [Output("extract")]
        public Output<Outputs.JobExtract?> Extract { get; private set; } = null!;

        /// <summary>
        /// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;

        /// <summary>
        /// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
        /// </summary>
        [Output("jobTimeoutMs")]
        public Output<string?> JobTimeoutMs { get; private set; } = null!;

        /// <summary>
        /// The type of the job.
        /// </summary>
        [Output("jobType")]
        public Output<string> JobType { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this job. You can use these to organize and group your jobs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Configures a load job.  Structure is documented below.
        /// </summary>
        [Output("load")]
        public Output<Outputs.JobLoad?> Load { get; private set; } = null!;

        /// <summary>
        /// The geographic location of the job. The default value is US.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Configures a query job.  Structure is documented below.
        /// </summary>
        [Output("query")]
        public Output<Outputs.JobQuery?> Query { get; private set; } = null!;

        /// <summary>
        /// Email address of the user who ran the job.
        /// </summary>
        [Output("userEmail")]
        public Output<string> UserEmail { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/job:Job", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
        {
            return new Job(name, id, state, options);
        }
    }

    public sealed class JobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Copies a table.  Structure is documented below.
        /// </summary>
        [Input("copy")]
        public Input<Inputs.JobCopyArgs>? Copy { get; set; }

        /// <summary>
        /// Configures an extract job.  Structure is documented below.
        /// </summary>
        [Input("extract")]
        public Input<Inputs.JobExtractArgs>? Extract { get; set; }

        /// <summary>
        /// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        /// <summary>
        /// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
        /// </summary>
        [Input("jobTimeoutMs")]
        public Input<string>? JobTimeoutMs { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this job. You can use these to organize and group your jobs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configures a load job.  Structure is documented below.
        /// </summary>
        [Input("load")]
        public Input<Inputs.JobLoadArgs>? Load { get; set; }

        /// <summary>
        /// The geographic location of the job. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Configures a query job.  Structure is documented below.
        /// </summary>
        [Input("query")]
        public Input<Inputs.JobQueryArgs>? Query { get; set; }

        public JobArgs()
        {
        }
    }

    public sealed class JobState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Copies a table.  Structure is documented below.
        /// </summary>
        [Input("copy")]
        public Input<Inputs.JobCopyGetArgs>? Copy { get; set; }

        /// <summary>
        /// Configures an extract job.  Structure is documented below.
        /// </summary>
        [Input("extract")]
        public Input<Inputs.JobExtractGetArgs>? Extract { get; set; }

        /// <summary>
        /// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        /// <summary>
        /// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
        /// </summary>
        [Input("jobTimeoutMs")]
        public Input<string>? JobTimeoutMs { get; set; }

        /// <summary>
        /// The type of the job.
        /// </summary>
        [Input("jobType")]
        public Input<string>? JobType { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this job. You can use these to organize and group your jobs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configures a load job.  Structure is documented below.
        /// </summary>
        [Input("load")]
        public Input<Inputs.JobLoadGetArgs>? Load { get; set; }

        /// <summary>
        /// The geographic location of the job. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Configures a query job.  Structure is documented below.
        /// </summary>
        [Input("query")]
        public Input<Inputs.JobQueryGetArgs>? Query { get; set; }

        /// <summary>
        /// Email address of the user who ran the job.
        /// </summary>
        [Input("userEmail")]
        public Input<string>? UserEmail { get; set; }

        public JobState()
        {
        }
    }
}
