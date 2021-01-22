// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigTable
{
    /// <summary>
    /// Creates a Google Cloud Bigtable table inside an instance. For more information see
    /// [the official documentation](https://cloud.google.com/bigtable/) and
    /// [API](https://cloud.google.com/bigtable/docs/go/reference).
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
    ///         var instance = new Gcp.BigTable.Instance("instance", new Gcp.BigTable.InstanceArgs
    ///         {
    ///             Clusters = 
    ///             {
    ///                 new Gcp.BigTable.Inputs.InstanceClusterArgs
    ///                 {
    ///                     ClusterId = "tf-instance-cluster",
    ///                     Zone = "us-central1-b",
    ///                     NumNodes = 3,
    ///                     StorageType = "HDD",
    ///                 },
    ///             },
    ///         });
    ///         var table = new Gcp.BigTable.Table("table", new Gcp.BigTable.TableArgs
    ///         {
    ///             InstanceName = instance.Name,
    ///             SplitKeys = 
    ///             {
    ///                 "a",
    ///                 "b",
    ///                 "c",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Bigtable Tables can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/table:Table default projects/{{project}}/instances/{{instance_name}}/tables/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/table:Table default {{project}}/{{instance_name}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/table:Table default {{instance_name}}/{{name}}
    /// ```
    /// 
    ///  The following fields can't be read and will show diffs if set in config when imported- `split_keys`
    /// </summary>
    [GcpResourceType("gcp:bigtable/table:Table")]
    public partial class Table : Pulumi.CustomResource
    {
        /// <summary>
        /// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
        /// </summary>
        [Output("columnFamilies")]
        public Output<ImmutableArray<Outputs.TableColumnFamily>> ColumnFamilies { get; private set; } = null!;

        /// <summary>
        /// The name of the Bigtable instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The name of the table.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A list of predefined keys to split the table on.
        /// !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
        /// to delete/recreate the entire `gcp.bigtable.Table` resource.
        /// </summary>
        [Output("splitKeys")]
        public Output<ImmutableArray<string>> SplitKeys { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigtable/table:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, TableState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigtable/table:Table", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, TableState? state = null, CustomResourceOptions? options = null)
        {
            return new Table(name, id, state, options);
        }
    }

    public sealed class TableArgs : Pulumi.ResourceArgs
    {
        [Input("columnFamilies")]
        private InputList<Inputs.TableColumnFamilyArgs>? _columnFamilies;

        /// <summary>
        /// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
        /// </summary>
        public InputList<Inputs.TableColumnFamilyArgs> ColumnFamilies
        {
            get => _columnFamilies ?? (_columnFamilies = new InputList<Inputs.TableColumnFamilyArgs>());
            set => _columnFamilies = value;
        }

        /// <summary>
        /// The name of the Bigtable instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the table.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("splitKeys")]
        private InputList<string>? _splitKeys;

        /// <summary>
        /// A list of predefined keys to split the table on.
        /// !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
        /// to delete/recreate the entire `gcp.bigtable.Table` resource.
        /// </summary>
        public InputList<string> SplitKeys
        {
            get => _splitKeys ?? (_splitKeys = new InputList<string>());
            set => _splitKeys = value;
        }

        public TableArgs()
        {
        }
    }

    public sealed class TableState : Pulumi.ResourceArgs
    {
        [Input("columnFamilies")]
        private InputList<Inputs.TableColumnFamilyGetArgs>? _columnFamilies;

        /// <summary>
        /// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
        /// </summary>
        public InputList<Inputs.TableColumnFamilyGetArgs> ColumnFamilies
        {
            get => _columnFamilies ?? (_columnFamilies = new InputList<Inputs.TableColumnFamilyGetArgs>());
            set => _columnFamilies = value;
        }

        /// <summary>
        /// The name of the Bigtable instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The name of the table.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("splitKeys")]
        private InputList<string>? _splitKeys;

        /// <summary>
        /// A list of predefined keys to split the table on.
        /// !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
        /// to delete/recreate the entire `gcp.bigtable.Table` resource.
        /// </summary>
        public InputList<string> SplitKeys
        {
            get => _splitKeys ?? (_splitKeys = new InputList<string>());
            set => _splitKeys = value;
        }

        public TableState()
        {
        }
    }
}
