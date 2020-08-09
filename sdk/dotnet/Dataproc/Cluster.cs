// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc
{
    /// <summary>
    /// Manages a Cloud Dataproc cluster resource within GCP. For more information see
    /// [the official dataproc documentation](https://cloud.google.com/dataproc/).
    /// 
    /// !&gt; **Warning:** Due to limitations of the API, all arguments except
    /// `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updatable. Changing others will cause recreation of the
    /// whole cluster!
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var simplecluster = new Gcp.Dataproc.Cluster("simplecluster", new Gcp.Dataproc.ClusterArgs
    ///         {
    ///             Region = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Advanced
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mycluster = new Gcp.Dataproc.Cluster("mycluster", new Gcp.Dataproc.ClusterArgs
    ///         {
    ///             ClusterConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigArgs
    ///             {
    ///                 GceClusterConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigGceClusterConfigArgs
    ///                 {
    ///                     ServiceAccountScopes = 
    ///                     {
    ///                         "https://www.googleapis.com/auth/monitoring",
    ///                         "useraccounts-ro",
    ///                         "storage-rw",
    ///                         "logging-write",
    ///                     },
    ///                     Tags = 
    ///                     {
    ///                         "foo",
    ///                         "bar",
    ///                     },
    ///                 },
    ///                 InitializationActions = 
    ///                 {
    ///                     new Gcp.Dataproc.Inputs.ClusterClusterConfigInitializationActionArgs
    ///                     {
    ///                         Script = "gs://dataproc-initialization-actions/stackdriver/stackdriver.sh",
    ///                         TimeoutSec = 500,
    ///                     },
    ///                 },
    ///                 MasterConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigMasterConfigArgs
    ///                 {
    ///                     DiskConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigMasterConfigDiskConfigArgs
    ///                     {
    ///                         BootDiskSizeGb = 15,
    ///                         BootDiskType = "pd-ssd",
    ///                     },
    ///                     MachineType = "n1-standard-1",
    ///                     NumInstances = 1,
    ///                 },
    ///                 PreemptibleWorkerConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigPreemptibleWorkerConfigArgs
    ///                 {
    ///                     NumInstances = 0,
    ///                 },
    ///                 SoftwareConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigSoftwareConfigArgs
    ///                 {
    ///                     ImageVersion = "1.3.7-deb9",
    ///                     OverrideProperties = 
    ///                     {
    ///                         { "dataproc:dataproc.allow.zero.workers", "true" },
    ///                     },
    ///                 },
    ///                 StagingBucket = "dataproc-staging-bucket",
    ///                 WorkerConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigWorkerConfigArgs
    ///                 {
    ///                     DiskConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigWorkerConfigDiskConfigArgs
    ///                     {
    ///                         BootDiskSizeGb = 15,
    ///                         NumLocalSsds = 1,
    ///                     },
    ///                     MachineType = "n1-standard-1",
    ///                     MinCpuPlatform = "Intel Skylake",
    ///                     NumInstances = 2,
    ///                 },
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///             Region = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Using A GPU Accelerator
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var acceleratedCluster = new Gcp.Dataproc.Cluster("acceleratedCluster", new Gcp.Dataproc.ClusterArgs
    ///         {
    ///             ClusterConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigArgs
    ///             {
    ///                 GceClusterConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigGceClusterConfigArgs
    ///                 {
    ///                     Zone = "us-central1-a",
    ///                 },
    ///                 MasterConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigMasterConfigArgs
    ///                 {
    ///                     Accelerators = 
    ///                     {
    ///                         new Gcp.Dataproc.Inputs.ClusterClusterConfigMasterConfigAcceleratorArgs
    ///                         {
    ///                             AcceleratorCount = 1,
    ///                             AcceleratorType = "nvidia-tesla-k80",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Region = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Allows you to configure various aspects of the cluster.
        /// Structure defined below.
        /// </summary>
        [Output("clusterConfig")]
        public Output<Outputs.ClusterClusterConfig> ClusterConfig { get; private set; } = null!;

        /// <summary>
        /// The list of labels (key/value pairs) to be applied to
        /// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
        /// which is the name of the cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the cluster, unique within the project and
        /// zone.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the `cluster` will exist. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region in which the cluster and associated nodes will be created in.
        /// Defaults to `global`.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows you to configure various aspects of the cluster.
        /// Structure defined below.
        /// </summary>
        [Input("clusterConfig")]
        public Input<Inputs.ClusterClusterConfigArgs>? ClusterConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The list of labels (key/value pairs) to be applied to
        /// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
        /// which is the name of the cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the cluster, unique within the project and
        /// zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the `cluster` will exist. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the cluster and associated nodes will be created in.
        /// Defaults to `global`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows you to configure various aspects of the cluster.
        /// Structure defined below.
        /// </summary>
        [Input("clusterConfig")]
        public Input<Inputs.ClusterClusterConfigGetArgs>? ClusterConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The list of labels (key/value pairs) to be applied to
        /// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
        /// which is the name of the cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the cluster, unique within the project and
        /// zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the `cluster` will exist. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the cluster and associated nodes will be created in.
        /// Defaults to `global`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ClusterState()
        {
        }
    }
}
