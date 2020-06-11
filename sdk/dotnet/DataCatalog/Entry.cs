// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataCatalog
{
    /// <summary>
    /// Entry Metadata. A Data Catalog Entry resource represents another resource in Google Cloud Platform
    /// (such as a BigQuery dataset or a Pub/Sub topic) or outside of Google Cloud Platform. Clients can use
    /// the linkedResource field in the Entry resource to refer to the original resource ID of the source system.
    /// 
    /// An Entry resource contains resource details, such as its schema. An Entry can also be used to attach
    /// flexible metadata, such as a Tag.
    /// 
    /// 
    /// To get more information about Entry, see:
    /// 
    /// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
    /// 
    /// ## Example Usage
    /// 
    /// ### Data Catalog Entry Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var entryGroup = new Gcp.DataCatalog.EntryGroup("entryGroup", new Gcp.DataCatalog.EntryGroupArgs
    ///         {
    ///             EntryGroupId = "my_group",
    ///         });
    ///         var basicEntry = new Gcp.DataCatalog.Entry("basicEntry", new Gcp.DataCatalog.EntryArgs
    ///         {
    ///             EntryGroup = entryGroup.Id,
    ///             EntryId = "my_entry",
    ///             UserSpecifiedType = "my_custom_type",
    ///             UserSpecifiedSystem = "SomethingExternal",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Data Catalog Entry Full
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var entryGroup = new Gcp.DataCatalog.EntryGroup("entryGroup", new Gcp.DataCatalog.EntryGroupArgs
    ///         {
    ///             EntryGroupId = "my_group",
    ///         });
    ///         var basicEntry = new Gcp.DataCatalog.Entry("basicEntry", new Gcp.DataCatalog.EntryArgs
    ///         {
    ///             EntryGroup = entryGroup.Id,
    ///             EntryId = "my_entry",
    ///             UserSpecifiedType = "my_user_specified_type",
    ///             UserSpecifiedSystem = "Something_custom",
    ///             LinkedResource = "my/linked/resource",
    ///             DisplayName = "my custom type entry",
    ///             Description = "a custom type entry for a user specified system",
    ///             Schema = @"{
    ///   ""columns"": [
    ///     {
    ///       ""column"": ""first_name"",
    ///       ""description"": ""First name"",
    ///       ""mode"": ""REQUIRED"",
    ///       ""type"": ""STRING""
    ///     },
    ///     {
    ///       ""column"": ""last_name"",
    ///       ""description"": ""Last name"",
    ///       ""mode"": ""REQUIRED"",
    ///       ""type"": ""STRING""
    ///     },
    ///     {
    ///       ""column"": ""address"",
    ///       ""description"": ""Address"",
    ///       ""mode"": ""REPEATED"",
    ///       ""subcolumns"": [
    ///         {
    ///           ""column"": ""city"",
    ///           ""description"": ""City"",
    ///           ""mode"": ""NULLABLE"",
    ///           ""type"": ""STRING""
    ///         },
    ///         {
    ///           ""column"": ""state"",
    ///           ""description"": ""State"",
    ///           ""mode"": ""NULLABLE"",
    ///           ""type"": ""STRING""
    ///         }
    ///       ],
    ///       ""type"": ""RECORD""
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Entry : Pulumi.CustomResource
    {
        /// <summary>
        /// Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
        /// https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        /// </summary>
        [Output("bigqueryDateShardedSpec")]
        public Output<Outputs.EntryBigqueryDateShardedSpec> BigqueryDateShardedSpec { get; private set; } = null!;

        /// <summary>
        /// Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
        /// </summary>
        [Output("bigqueryTableSpec")]
        public Output<Outputs.EntryBigqueryTableSpec> BigqueryTableSpec { get; private set; } = null!;

        /// <summary>
        /// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Display information such as title and description. A short name to identify the entry,
        /// for example, "Analytics Data - Jan 2011".
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the entry group this entry is in.
        /// </summary>
        [Output("entryGroup")]
        public Output<string> EntryGroup { get; private set; } = null!;

        /// <summary>
        /// The id of the entry to create.
        /// </summary>
        [Output("entryId")]
        public Output<string> EntryId { get; private set; } = null!;

        /// <summary>
        /// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
        /// </summary>
        [Output("gcsFilesetSpec")]
        public Output<Outputs.EntryGcsFilesetSpec?> GcsFilesetSpec { get; private set; } = null!;

        /// <summary>
        /// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
        /// </summary>
        [Output("integratedSystem")]
        public Output<string> IntegratedSystem { get; private set; } = null!;

        /// <summary>
        /// The resource this metadata entry refers to.
        /// For Google Cloud Platform resources, linkedResource is the full name of the resource.
        /// For example, the linkedResource for a table resource from BigQuery is:
        /// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
        /// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
        /// this field is optional and defaults to an empty string.
        /// </summary>
        [Output("linkedResource")]
        public Output<string> LinkedResource { get; private set; } = null!;

        /// <summary>
        /// The Data Catalog resource name of the entry in URL format. Example:
        /// projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
        /// child resources may not actually be stored in the location in this name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
        /// attached to it. See
        /// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
        /// for what fields this schema can contain.
        /// </summary>
        [Output("schema")]
        public Output<string?> Schema { get; private set; } = null!;

        /// <summary>
        /// The type of the entry. Only used for Entries with types in the EntryType enum.
        /// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// This field indicates the entry's source system that Data Catalog does not integrate with.
        /// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
        /// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        /// </summary>
        [Output("userSpecifiedSystem")]
        public Output<string?> UserSpecifiedSystem { get; private set; } = null!;

        /// <summary>
        /// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
        /// When creating an entry, users should check the enum values first, if nothing matches the entry
        /// to be created, then provide a custom value, for example "my_special_type".
        /// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
        /// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        /// </summary>
        [Output("userSpecifiedType")]
        public Output<string?> UserSpecifiedType { get; private set; } = null!;


        /// <summary>
        /// Create a Entry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Entry(string name, EntryArgs args, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/entry:Entry", name, args ?? new EntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Entry(string name, Input<string> id, EntryState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/entry:Entry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Entry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Entry Get(string name, Input<string> id, EntryState? state = null, CustomResourceOptions? options = null)
        {
            return new Entry(name, id, state, options);
        }
    }

    public sealed class EntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display information such as title and description. A short name to identify the entry,
        /// for example, "Analytics Data - Jan 2011".
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the entry group this entry is in.
        /// </summary>
        [Input("entryGroup", required: true)]
        public Input<string> EntryGroup { get; set; } = null!;

        /// <summary>
        /// The id of the entry to create.
        /// </summary>
        [Input("entryId", required: true)]
        public Input<string> EntryId { get; set; } = null!;

        /// <summary>
        /// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
        /// </summary>
        [Input("gcsFilesetSpec")]
        public Input<Inputs.EntryGcsFilesetSpecArgs>? GcsFilesetSpec { get; set; }

        /// <summary>
        /// The resource this metadata entry refers to.
        /// For Google Cloud Platform resources, linkedResource is the full name of the resource.
        /// For example, the linkedResource for a table resource from BigQuery is:
        /// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
        /// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
        /// this field is optional and defaults to an empty string.
        /// </summary>
        [Input("linkedResource")]
        public Input<string>? LinkedResource { get; set; }

        /// <summary>
        /// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
        /// attached to it. See
        /// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
        /// for what fields this schema can contain.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// The type of the entry. Only used for Entries with types in the EntryType enum.
        /// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// This field indicates the entry's source system that Data Catalog does not integrate with.
        /// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
        /// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        /// </summary>
        [Input("userSpecifiedSystem")]
        public Input<string>? UserSpecifiedSystem { get; set; }

        /// <summary>
        /// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
        /// When creating an entry, users should check the enum values first, if nothing matches the entry
        /// to be created, then provide a custom value, for example "my_special_type".
        /// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
        /// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        /// </summary>
        [Input("userSpecifiedType")]
        public Input<string>? UserSpecifiedType { get; set; }

        public EntryArgs()
        {
        }
    }

    public sealed class EntryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
        /// https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        /// </summary>
        [Input("bigqueryDateShardedSpec")]
        public Input<Inputs.EntryBigqueryDateShardedSpecGetArgs>? BigqueryDateShardedSpec { get; set; }

        /// <summary>
        /// Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
        /// </summary>
        [Input("bigqueryTableSpec")]
        public Input<Inputs.EntryBigqueryTableSpecGetArgs>? BigqueryTableSpec { get; set; }

        /// <summary>
        /// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display information such as title and description. A short name to identify the entry,
        /// for example, "Analytics Data - Jan 2011".
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the entry group this entry is in.
        /// </summary>
        [Input("entryGroup")]
        public Input<string>? EntryGroup { get; set; }

        /// <summary>
        /// The id of the entry to create.
        /// </summary>
        [Input("entryId")]
        public Input<string>? EntryId { get; set; }

        /// <summary>
        /// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
        /// </summary>
        [Input("gcsFilesetSpec")]
        public Input<Inputs.EntryGcsFilesetSpecGetArgs>? GcsFilesetSpec { get; set; }

        /// <summary>
        /// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
        /// </summary>
        [Input("integratedSystem")]
        public Input<string>? IntegratedSystem { get; set; }

        /// <summary>
        /// The resource this metadata entry refers to.
        /// For Google Cloud Platform resources, linkedResource is the full name of the resource.
        /// For example, the linkedResource for a table resource from BigQuery is:
        /// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
        /// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
        /// this field is optional and defaults to an empty string.
        /// </summary>
        [Input("linkedResource")]
        public Input<string>? LinkedResource { get; set; }

        /// <summary>
        /// The Data Catalog resource name of the entry in URL format. Example:
        /// projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
        /// child resources may not actually be stored in the location in this name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
        /// attached to it. See
        /// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
        /// for what fields this schema can contain.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// The type of the entry. Only used for Entries with types in the EntryType enum.
        /// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// This field indicates the entry's source system that Data Catalog does not integrate with.
        /// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
        /// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        /// </summary>
        [Input("userSpecifiedSystem")]
        public Input<string>? UserSpecifiedSystem { get; set; }

        /// <summary>
        /// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
        /// When creating an entry, users should check the enum values first, if nothing matches the entry
        /// to be created, then provide a custom value, for example "my_special_type".
        /// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
        /// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        /// </summary>
        [Input("userSpecifiedType")]
        public Input<string>? UserSpecifiedType { get; set; }

        public EntryState()
        {
        }
    }
}
