// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql
{
    /// <summary>
    /// A source representation instance is a Cloud SQL instance that represents
    /// the source database server to the Cloud SQL replica. It is visible in the
    /// Cloud Console and appears the same as a regular Cloud SQL instance, but it
    /// contains no data, requires no configuration or maintenance, and does not
    /// affect billing. You cannot update the source representation instance.
    /// 
    /// ## Example Usage
    /// ### Sql Source Representation Instance Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var instance = new Gcp.Sql.SourceRepresentationInstance("instance", new Gcp.Sql.SourceRepresentationInstanceArgs
    ///         {
    ///             DatabaseVersion = "MYSQL_5_7",
    ///             Host = "10.20.30.40",
    ///             Port = 3306,
    ///             Region = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SourceRepresentationInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// The MySQL version running on your source database server.
        /// </summary>
        [Output("databaseVersion")]
        public Output<string> DatabaseVersion { get; private set; } = null!;

        /// <summary>
        /// The externally accessible IPv4 address for the source database server.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// The name of the source representation instance. Use any valid Cloud SQL instance name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The externally accessible port for the source database server.
        /// Defaults to 3306.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The Region in which the created instance should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a SourceRepresentationInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceRepresentationInstance(string name, SourceRepresentationInstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, args ?? new SourceRepresentationInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceRepresentationInstance(string name, Input<string> id, SourceRepresentationInstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SourceRepresentationInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceRepresentationInstance Get(string name, Input<string> id, SourceRepresentationInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new SourceRepresentationInstance(name, id, state, options);
        }
    }

    public sealed class SourceRepresentationInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The MySQL version running on your source database server.
        /// </summary>
        [Input("databaseVersion", required: true)]
        public Input<string> DatabaseVersion { get; set; } = null!;

        /// <summary>
        /// The externally accessible IPv4 address for the source database server.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// The name of the source representation instance. Use any valid Cloud SQL instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The externally accessible port for the source database server.
        /// Defaults to 3306.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created instance should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SourceRepresentationInstanceArgs()
        {
        }
    }

    public sealed class SourceRepresentationInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The MySQL version running on your source database server.
        /// </summary>
        [Input("databaseVersion")]
        public Input<string>? DatabaseVersion { get; set; }

        /// <summary>
        /// The externally accessible IPv4 address for the source database server.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The name of the source representation instance. Use any valid Cloud SQL instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The externally accessible port for the source database server.
        /// Defaults to 3306.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created instance should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SourceRepresentationInstanceState()
        {
        }
    }
}
