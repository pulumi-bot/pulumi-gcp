// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Spanner
{
    /// <summary>
    /// A Cloud Spanner Database which is hosted on a Spanner instance.
    /// 
    /// 
    /// To get more information about Database, see:
    /// 
    /// * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/spanner/)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database.html.markdown.
    /// </summary>
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional list of DDL statements to run inside the newly created
        /// database. Statements can create tables, indexes, etc. These statements
        /// execute atomically with the creation of the database: if there is an
        /// error in any statement, the database is not created.
        /// </summary>
        [Output("ddls")]
        public Output<ImmutableArray<string>> Ddls { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The instance to create the database on.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// A unique identifier for the database, which cannot be changed after
        /// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
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
        /// An explanation of the status of the database.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("gcp:spanner/database:Database", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("gcp:spanner/database:Database", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        [Input("ddls")]
        private InputList<string>? _ddls;

        /// <summary>
        /// -
        /// (Optional)
        /// An optional list of DDL statements to run inside the newly created
        /// database. Statements can create tables, indexes, etc. These statements
        /// execute atomically with the creation of the database: if there is an
        /// error in any statement, the database is not created.
        /// </summary>
        public InputList<string> Ddls
        {
            get => _ddls ?? (_ddls = new InputList<string>());
            set => _ddls = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// The instance to create the database on.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// A unique identifier for the database, which cannot be changed after
        /// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public DatabaseArgs()
        {
        }
    }

    public sealed class DatabaseState : Pulumi.ResourceArgs
    {
        [Input("ddls")]
        private InputList<string>? _ddls;

        /// <summary>
        /// -
        /// (Optional)
        /// An optional list of DDL statements to run inside the newly created
        /// database. Statements can create tables, indexes, etc. These statements
        /// execute atomically with the creation of the database: if there is an
        /// error in any statement, the database is not created.
        /// </summary>
        public InputList<string> Ddls
        {
            get => _ddls ?? (_ddls = new InputList<string>());
            set => _ddls = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// The instance to create the database on.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// A unique identifier for the database, which cannot be changed after
        /// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
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
        /// An explanation of the status of the database.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public DatabaseState()
        {
        }
    }
}
