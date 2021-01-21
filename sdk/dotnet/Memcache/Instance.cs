// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Memcache
{
    /// <summary>
    /// A Google Cloud Memcache instance.
    /// 
    /// To get more information about Instance, see:
    /// 
    /// * [API documentation](https://cloud.google.com/memorystore/docs/memcached/reference/rest)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/memcache/docs/creating-instances)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Instance can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default {{name}}
    /// ```
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// The full name of the GCE network to connect the instance to.  If not provided,
        /// 'default' will be used.
        /// </summary>
        [Output("authorizedNetwork")]
        public Output<string?> AuthorizedNetwork { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Endpoint for Discovery API
        /// </summary>
        [Output("discoveryEndpoint")]
        public Output<string> DiscoveryEndpoint { get; private set; } = null!;

        /// <summary>
        /// A user-visible name for the instance.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The full version of memcached server running on this instance.
        /// </summary>
        [Output("memcacheFullVersion")]
        public Output<string> MemcacheFullVersion { get; private set; } = null!;

        /// <summary>
        /// Additional information about the instance state, if available.
        /// </summary>
        [Output("memcacheNodes")]
        public Output<ImmutableArray<Outputs.InstanceMemcacheNode>> MemcacheNodes { get; private set; } = null!;

        /// <summary>
        /// User-specified parameters for this memcache instance.
        /// Structure is documented below.
        /// </summary>
        [Output("memcacheParameters")]
        public Output<Outputs.InstanceMemcacheParameters?> MemcacheParameters { get; private set; } = null!;

        /// <summary>
        /// The major version of Memcached software. If not provided, latest supported version will be used.
        /// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
        /// determined by our system based on the latest supported minor version.
        /// Default value is `MEMCACHE_1_5`.
        /// Possible values are `MEMCACHE_1_5`.
        /// </summary>
        [Output("memcacheVersion")]
        public Output<string?> MemcacheVersion { get; private set; } = null!;

        /// <summary>
        /// The resource name of the instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration for memcache nodes.
        /// Structure is documented below.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.InstanceNodeConfig> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// Number of nodes in the memcache instance.
        /// </summary>
        [Output("nodeCount")]
        public Output<int> NodeCount { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region of the Memcache instance. If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Zones where memcache nodes should be provisioned.  If not
        /// provided, all zones will be used.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:memcache/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:memcache/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full name of the GCE network to connect the instance to.  If not provided,
        /// 'default' will be used.
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// A user-visible name for the instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// User-specified parameters for this memcache instance.
        /// Structure is documented below.
        /// </summary>
        [Input("memcacheParameters")]
        public Input<Inputs.InstanceMemcacheParametersArgs>? MemcacheParameters { get; set; }

        /// <summary>
        /// The major version of Memcached software. If not provided, latest supported version will be used.
        /// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
        /// determined by our system based on the latest supported minor version.
        /// Default value is `MEMCACHE_1_5`.
        /// Possible values are `MEMCACHE_1_5`.
        /// </summary>
        [Input("memcacheVersion")]
        public Input<string>? MemcacheVersion { get; set; }

        /// <summary>
        /// The resource name of the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration for memcache nodes.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeConfig", required: true)]
        public Input<Inputs.InstanceNodeConfigArgs> NodeConfig { get; set; } = null!;

        /// <summary>
        /// Number of nodes in the memcache instance.
        /// </summary>
        [Input("nodeCount", required: true)]
        public Input<int> NodeCount { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Memcache instance. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Zones where memcache nodes should be provisioned.  If not
        /// provided, all zones will be used.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full name of the GCE network to connect the instance to.  If not provided,
        /// 'default' will be used.
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Endpoint for Discovery API
        /// </summary>
        [Input("discoveryEndpoint")]
        public Input<string>? DiscoveryEndpoint { get; set; }

        /// <summary>
        /// A user-visible name for the instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The full version of memcached server running on this instance.
        /// </summary>
        [Input("memcacheFullVersion")]
        public Input<string>? MemcacheFullVersion { get; set; }

        [Input("memcacheNodes")]
        private InputList<Inputs.InstanceMemcacheNodeGetArgs>? _memcacheNodes;

        /// <summary>
        /// Additional information about the instance state, if available.
        /// </summary>
        public InputList<Inputs.InstanceMemcacheNodeGetArgs> MemcacheNodes
        {
            get => _memcacheNodes ?? (_memcacheNodes = new InputList<Inputs.InstanceMemcacheNodeGetArgs>());
            set => _memcacheNodes = value;
        }

        /// <summary>
        /// User-specified parameters for this memcache instance.
        /// Structure is documented below.
        /// </summary>
        [Input("memcacheParameters")]
        public Input<Inputs.InstanceMemcacheParametersGetArgs>? MemcacheParameters { get; set; }

        /// <summary>
        /// The major version of Memcached software. If not provided, latest supported version will be used.
        /// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
        /// determined by our system based on the latest supported minor version.
        /// Default value is `MEMCACHE_1_5`.
        /// Possible values are `MEMCACHE_1_5`.
        /// </summary>
        [Input("memcacheVersion")]
        public Input<string>? MemcacheVersion { get; set; }

        /// <summary>
        /// The resource name of the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration for memcache nodes.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.InstanceNodeConfigGetArgs>? NodeConfig { get; set; }

        /// <summary>
        /// Number of nodes in the memcache instance.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Memcache instance. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Zones where memcache nodes should be provisioned.  If not
        /// provided, all zones will be used.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public InstanceState()
        {
        }
    }
}
