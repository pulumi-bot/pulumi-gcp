// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container
{
    /// <summary>
    /// Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
    /// the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
    /// and [the API reference](https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.nodePools).
    /// </summary>
    public partial class NodePool : Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration required by cluster autoscaler to adjust
        /// the size of the node pool to the current cluster usage. Structure is documented below.
        /// </summary>
        [Output("autoscaling")]
        public Output<Outputs.NodePoolAutoscaling?> Autoscaling { get; private set; } = null!;

        /// <summary>
        /// The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// The initial number of nodes for the pool. In
        /// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
        /// this will force recreation of the resource.
        /// </summary>
        [Output("initialNodeCount")]
        public Output<int> InitialNodeCount { get; private set; } = null!;

        /// <summary>
        /// The resource URLs of the managed instance groups associated with this node pool.
        /// </summary>
        [Output("instanceGroupUrls")]
        public Output<ImmutableArray<string>> InstanceGroupUrls { get; private set; } = null!;

        /// <summary>
        /// The location (region or zone) of the cluster.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Node management configuration, wherein auto-repair and
        /// auto-upgrade is configured. Structure is documented below.
        /// </summary>
        [Output("management")]
        public Output<Outputs.NodePoolManagement> Management { get; private set; } = null!;

        /// <summary>
        /// The maximum number of pods per node in this node pool.
        /// Note that this does not work on node pools which are "route-based" - that is, node
        /// pools belonging to clusters that do not have IP Aliasing enabled.
        /// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
        /// for more information.
        /// </summary>
        [Output("maxPodsPerNode")]
        public Output<int> MaxPodsPerNode { get; private set; } = null!;

        /// <summary>
        /// The name of the node pool. If left blank, the provider will
        /// auto-generate a unique name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name for the node pool beginning
        /// with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The node configuration of the pool. See
        /// gcp.container.Cluster for schema.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.NodePoolNodeConfig> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// The number of nodes per instance group. This field can be used to
        /// update the number of nodes per instance group but should not be used alongside `autoscaling`.
        /// </summary>
        [Output("nodeCount")]
        public Output<int> NodeCount { get; private set; } = null!;

        /// <summary>
        /// 
        /// The list of zones in which the node pool's nodes should be located. Nodes must
        /// be in the region of their regional cluster or in the same region as their
        /// cluster's zone for zonal clusters. If unspecified, the cluster-level
        /// `node_locations` will be used.
        /// </summary>
        [Output("nodeLocations")]
        public Output<ImmutableArray<string>> NodeLocations { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which to create the node pool. If blank,
        /// the provider-configured project will be used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("upgradeSettings")]
        public Output<Outputs.NodePoolUpgradeSettings> UpgradeSettings { get; private set; } = null!;

        /// <summary>
        /// The Kubernetes version for the nodes in this pool. Note that if this field
        /// and `auto_upgrade` are both specified, they will fight each other for what the node version should
        /// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
        /// recommended that you specify explicit versions as the provider will see spurious diffs
        /// when fuzzy versions are used. See the `gcp.container.getEngineVersions` data source's
        /// `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a NodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodePool(string name, NodePoolArgs args, CustomResourceOptions? options = null)
            : base("gcp:container/nodePool:NodePool", name, args ?? new NodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodePool(string name, Input<string> id, NodePoolState? state = null, CustomResourceOptions? options = null)
            : base("gcp:container/nodePool:NodePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodePool Get(string name, Input<string> id, NodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new NodePool(name, id, state, options);
        }
    }

    public sealed class NodePoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration required by cluster autoscaler to adjust
        /// the size of the node pool to the current cluster usage. Structure is documented below.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.NodePoolAutoscalingArgs>? Autoscaling { get; set; }

        /// <summary>
        /// The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        /// <summary>
        /// The initial number of nodes for the pool. In
        /// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
        /// this will force recreation of the resource.
        /// </summary>
        [Input("initialNodeCount")]
        public Input<int>? InitialNodeCount { get; set; }

        /// <summary>
        /// The location (region or zone) of the cluster.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Node management configuration, wherein auto-repair and
        /// auto-upgrade is configured. Structure is documented below.
        /// </summary>
        [Input("management")]
        public Input<Inputs.NodePoolManagementArgs>? Management { get; set; }

        /// <summary>
        /// The maximum number of pods per node in this node pool.
        /// Note that this does not work on node pools which are "route-based" - that is, node
        /// pools belonging to clusters that do not have IP Aliasing enabled.
        /// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
        /// for more information.
        /// </summary>
        [Input("maxPodsPerNode")]
        public Input<int>? MaxPodsPerNode { get; set; }

        /// <summary>
        /// The name of the node pool. If left blank, the provider will
        /// auto-generate a unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name for the node pool beginning
        /// with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The node configuration of the pool. See
        /// gcp.container.Cluster for schema.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.NodePoolNodeConfigArgs>? NodeConfig { get; set; }

        /// <summary>
        /// The number of nodes per instance group. This field can be used to
        /// update the number of nodes per instance group but should not be used alongside `autoscaling`.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("nodeLocations")]
        private InputList<string>? _nodeLocations;

        /// <summary>
        /// 
        /// The list of zones in which the node pool's nodes should be located. Nodes must
        /// be in the region of their regional cluster or in the same region as their
        /// cluster's zone for zonal clusters. If unspecified, the cluster-level
        /// `node_locations` will be used.
        /// </summary>
        public InputList<string> NodeLocations
        {
            get => _nodeLocations ?? (_nodeLocations = new InputList<string>());
            set => _nodeLocations = value;
        }

        /// <summary>
        /// The ID of the project in which to create the node pool. If blank,
        /// the provider-configured project will be used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("upgradeSettings")]
        public Input<Inputs.NodePoolUpgradeSettingsArgs>? UpgradeSettings { get; set; }

        /// <summary>
        /// The Kubernetes version for the nodes in this pool. Note that if this field
        /// and `auto_upgrade` are both specified, they will fight each other for what the node version should
        /// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
        /// recommended that you specify explicit versions as the provider will see spurious diffs
        /// when fuzzy versions are used. See the `gcp.container.getEngineVersions` data source's
        /// `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodePoolArgs()
        {
        }
    }

    public sealed class NodePoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration required by cluster autoscaler to adjust
        /// the size of the node pool to the current cluster usage. Structure is documented below.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.NodePoolAutoscalingGetArgs>? Autoscaling { get; set; }

        /// <summary>
        /// The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// The initial number of nodes for the pool. In
        /// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
        /// this will force recreation of the resource.
        /// </summary>
        [Input("initialNodeCount")]
        public Input<int>? InitialNodeCount { get; set; }

        [Input("instanceGroupUrls")]
        private InputList<string>? _instanceGroupUrls;

        /// <summary>
        /// The resource URLs of the managed instance groups associated with this node pool.
        /// </summary>
        public InputList<string> InstanceGroupUrls
        {
            get => _instanceGroupUrls ?? (_instanceGroupUrls = new InputList<string>());
            set => _instanceGroupUrls = value;
        }

        /// <summary>
        /// The location (region or zone) of the cluster.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Node management configuration, wherein auto-repair and
        /// auto-upgrade is configured. Structure is documented below.
        /// </summary>
        [Input("management")]
        public Input<Inputs.NodePoolManagementGetArgs>? Management { get; set; }

        /// <summary>
        /// The maximum number of pods per node in this node pool.
        /// Note that this does not work on node pools which are "route-based" - that is, node
        /// pools belonging to clusters that do not have IP Aliasing enabled.
        /// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
        /// for more information.
        /// </summary>
        [Input("maxPodsPerNode")]
        public Input<int>? MaxPodsPerNode { get; set; }

        /// <summary>
        /// The name of the node pool. If left blank, the provider will
        /// auto-generate a unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name for the node pool beginning
        /// with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The node configuration of the pool. See
        /// gcp.container.Cluster for schema.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.NodePoolNodeConfigGetArgs>? NodeConfig { get; set; }

        /// <summary>
        /// The number of nodes per instance group. This field can be used to
        /// update the number of nodes per instance group but should not be used alongside `autoscaling`.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("nodeLocations")]
        private InputList<string>? _nodeLocations;

        /// <summary>
        /// 
        /// The list of zones in which the node pool's nodes should be located. Nodes must
        /// be in the region of their regional cluster or in the same region as their
        /// cluster's zone for zonal clusters. If unspecified, the cluster-level
        /// `node_locations` will be used.
        /// </summary>
        public InputList<string> NodeLocations
        {
            get => _nodeLocations ?? (_nodeLocations = new InputList<string>());
            set => _nodeLocations = value;
        }

        /// <summary>
        /// The ID of the project in which to create the node pool. If blank,
        /// the provider-configured project will be used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("upgradeSettings")]
        public Input<Inputs.NodePoolUpgradeSettingsGetArgs>? UpgradeSettings { get; set; }

        /// <summary>
        /// The Kubernetes version for the nodes in this pool. Note that if this field
        /// and `auto_upgrade` are both specified, they will fight each other for what the node version should
        /// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
        /// recommended that you specify explicit versions as the provider will see spurious diffs
        /// when fuzzy versions are used. See the `gcp.container.getEngineVersions` data source's
        /// `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodePoolState()
        {
        }
    }
}
