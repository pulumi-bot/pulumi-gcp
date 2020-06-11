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
    /// Represents a data transfer configuration. A transfer configuration
    /// contains all metadata needed to perform a data transfer.
    /// 
    /// 
    /// To get more information about Config, see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigquery/docs/reference/datatransfer/rest/v1/projects.locations.transferConfigs/create)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/bigquery/docs/reference/datatransfer/rest/)
    /// 
    /// ## Example Usage
    /// 
    /// ### Scheduled Query
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project = Output.Create(Gcp.Organizations.GetProject.InvokeAsync());
    ///         var permissions = new Gcp.Projects.IAMMember("permissions", new Gcp.Projects.IAMMemberArgs
    ///         {
    ///             Role = "roles/iam.serviceAccountShortTermTokenMinter",
    ///             Member = project.Apply(project =&gt; $"serviceAccount:service-{project.Number}@gcp-sa-bigquerydatatransfer.iam.gserviceaccount.com"),
    ///         });
    ///         var myDataset = new Gcp.BigQuery.Dataset("myDataset", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "my_dataset",
    ///             FriendlyName = "foo",
    ///             Description = "bar",
    ///             Location = "asia-northeast1",
    ///         });
    ///         var queryConfig = new Gcp.BigQuery.DataTransferConfig("queryConfig", new Gcp.BigQuery.DataTransferConfigArgs
    ///         {
    ///             DisplayName = "my-query",
    ///             Location = "asia-northeast1",
    ///             DataSourceId = "scheduled_query",
    ///             Schedule = "first sunday of quarter 00:00",
    ///             DestinationDatasetId = myDataset.DatasetId,
    ///             Params = 
    ///             {
    ///                 { "destination_table_name_template", "my-table" },
    ///                 { "write_disposition", "WRITE_APPEND" },
    ///                 { "query", "SELECT name FROM tabl WHERE x = 'y'" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DataTransferConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The number of days to look back to automatically refresh the data.
        /// For example, if dataRefreshWindowDays = 10, then every day BigQuery
        /// reingests data for [today-10, today-1], rather than ingesting data for
        /// just [today-1]. Only valid if the data source supports the feature.
        /// Set the value to 0 to use the default value.
        /// </summary>
        [Output("dataRefreshWindowDays")]
        public Output<int?> DataRefreshWindowDays { get; private set; } = null!;

        /// <summary>
        /// The data source id. Cannot be changed once the transfer config is created.
        /// </summary>
        [Output("dataSourceId")]
        public Output<string> DataSourceId { get; private set; } = null!;

        /// <summary>
        /// The BigQuery target dataset id.
        /// </summary>
        [Output("destinationDatasetId")]
        public Output<string> DestinationDatasetId { get; private set; } = null!;

        /// <summary>
        /// When set to true, no runs are scheduled for a given transfer.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// The user specified display name for the transfer config.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The geographic location where the transfer config should reside.
        /// Examples: US, EU, asia-northeast1. The default value is US.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the transfer config. Transfer config names have the form
        /// projects/{projectId}/locations/{location}/transferConfigs/{configId}. Where configId is usually a uuid, but this is not
        /// required. The name is ignored when creating a transfer config.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// These parameters are specific to each data source.
        /// </summary>
        [Output("params")]
        public Output<ImmutableDictionary<string, string>> Params { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Data transfer schedule. If the data source does not support a custom
        /// schedule, this should be empty. If it is empty, the default value for
        /// the data source will be used. The specified times are in UTC. Examples
        /// of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
        /// jun 13:15, and first sunday of quarter 00:00. See more explanation
        /// about the format here:
        /// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
        /// NOTE: the granularity should be at least 8 hours, or less frequent.
        /// </summary>
        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Optional service account name. If this field is set, transfer config will
        /// be created with this service account credentials. It requires that
        /// requesting user calling this API has permissions to act as this service account.
        /// </summary>
        [Output("serviceAccountName")]
        public Output<string?> ServiceAccountName { get; private set; } = null!;


        /// <summary>
        /// Create a DataTransferConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataTransferConfig(string name, DataTransferConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/dataTransferConfig:DataTransferConfig", name, args ?? new DataTransferConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataTransferConfig(string name, Input<string> id, DataTransferConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/dataTransferConfig:DataTransferConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataTransferConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataTransferConfig Get(string name, Input<string> id, DataTransferConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DataTransferConfig(name, id, state, options);
        }
    }

    public sealed class DataTransferConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days to look back to automatically refresh the data.
        /// For example, if dataRefreshWindowDays = 10, then every day BigQuery
        /// reingests data for [today-10, today-1], rather than ingesting data for
        /// just [today-1]. Only valid if the data source supports the feature.
        /// Set the value to 0 to use the default value.
        /// </summary>
        [Input("dataRefreshWindowDays")]
        public Input<int>? DataRefreshWindowDays { get; set; }

        /// <summary>
        /// The data source id. Cannot be changed once the transfer config is created.
        /// </summary>
        [Input("dataSourceId", required: true)]
        public Input<string> DataSourceId { get; set; } = null!;

        /// <summary>
        /// The BigQuery target dataset id.
        /// </summary>
        [Input("destinationDatasetId", required: true)]
        public Input<string> DestinationDatasetId { get; set; } = null!;

        /// <summary>
        /// When set to true, no runs are scheduled for a given transfer.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The user specified display name for the transfer config.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The geographic location where the transfer config should reside.
        /// Examples: US, EU, asia-northeast1. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("params", required: true)]
        private InputMap<string>? _params;

        /// <summary>
        /// These parameters are specific to each data source.
        /// </summary>
        public InputMap<string> Params
        {
            get => _params ?? (_params = new InputMap<string>());
            set => _params = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Data transfer schedule. If the data source does not support a custom
        /// schedule, this should be empty. If it is empty, the default value for
        /// the data source will be used. The specified times are in UTC. Examples
        /// of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
        /// jun 13:15, and first sunday of quarter 00:00. See more explanation
        /// about the format here:
        /// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
        /// NOTE: the granularity should be at least 8 hours, or less frequent.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Optional service account name. If this field is set, transfer config will
        /// be created with this service account credentials. It requires that
        /// requesting user calling this API has permissions to act as this service account.
        /// </summary>
        [Input("serviceAccountName")]
        public Input<string>? ServiceAccountName { get; set; }

        public DataTransferConfigArgs()
        {
        }
    }

    public sealed class DataTransferConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days to look back to automatically refresh the data.
        /// For example, if dataRefreshWindowDays = 10, then every day BigQuery
        /// reingests data for [today-10, today-1], rather than ingesting data for
        /// just [today-1]. Only valid if the data source supports the feature.
        /// Set the value to 0 to use the default value.
        /// </summary>
        [Input("dataRefreshWindowDays")]
        public Input<int>? DataRefreshWindowDays { get; set; }

        /// <summary>
        /// The data source id. Cannot be changed once the transfer config is created.
        /// </summary>
        [Input("dataSourceId")]
        public Input<string>? DataSourceId { get; set; }

        /// <summary>
        /// The BigQuery target dataset id.
        /// </summary>
        [Input("destinationDatasetId")]
        public Input<string>? DestinationDatasetId { get; set; }

        /// <summary>
        /// When set to true, no runs are scheduled for a given transfer.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The user specified display name for the transfer config.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The geographic location where the transfer config should reside.
        /// Examples: US, EU, asia-northeast1. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the transfer config. Transfer config names have the form
        /// projects/{projectId}/locations/{location}/transferConfigs/{configId}. Where configId is usually a uuid, but this is not
        /// required. The name is ignored when creating a transfer config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("params")]
        private InputMap<string>? _params;

        /// <summary>
        /// These parameters are specific to each data source.
        /// </summary>
        public InputMap<string> Params
        {
            get => _params ?? (_params = new InputMap<string>());
            set => _params = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Data transfer schedule. If the data source does not support a custom
        /// schedule, this should be empty. If it is empty, the default value for
        /// the data source will be used. The specified times are in UTC. Examples
        /// of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
        /// jun 13:15, and first sunday of quarter 00:00. See more explanation
        /// about the format here:
        /// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
        /// NOTE: the granularity should be at least 8 hours, or less frequent.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Optional service account name. If this field is set, transfer config will
        /// be created with this service account credentials. It requires that
        /// requesting user calling this API has permissions to act as this service account.
        /// </summary>
        [Input("serviceAccountName")]
        public Input<string>? ServiceAccountName { get; set; }

        public DataTransferConfigState()
        {
        }
    }
}
