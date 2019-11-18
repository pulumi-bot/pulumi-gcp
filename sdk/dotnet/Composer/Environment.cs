// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer
{
    /// <summary>
    /// An environment for running orchestration tasks.
    /// 
    /// Environments run Apache Airflow software on Google infrastructure.
    /// 
    /// To get more information about Environments, see:
    /// 
    /// * [API documentation](https://cloud.google.com/composer/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/composer/docs)
    ///     * [Configuring Shared VPC for Composer Environments](https://cloud.google.com/composer/docs/how-to/managing/configuring-shared-vpc)
    /// * [Apache Airflow Documentation](http://airflow.apache.org/)
    /// 
    /// &gt; **Warning:** We **STRONGLY** recommend  you read the [GCP guides](https://cloud.google.com/composer/docs/how-to)
    ///   as the Environment resource requires a long deployment process and involves several layers of GCP infrastructure, 
    ///   including a Kubernetes Engine cluster, Cloud Storage, and Compute networking resources. Due to limitations of the API,
    ///   this provider will not be able to automatically find or manage many of these underlying resources. In particular:
    ///   * It can take up to one hour to create or update an environment resource. In addition, GCP may only detect some 
    ///     errors in configuration when they are used (e.g. ~40-50 minutes into the creation process), and is prone to limited
    ///     error reporting. If you encounter confusing or uninformative errors, please verify your configuration is valid 
    ///     against GCP Cloud Composer before filing bugs against this provider. 
    ///   * **Environments create Google Cloud Storage buckets that do not get cleaned up automatically** on environment 
    ///     deletion. [More about Composer's use of Cloud Storage](https://cloud.google.com/composer/docs/concepts/cloud-storage).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/composer_environment.html.markdown.
    /// </summary>
    public partial class Environment : Pulumi.CustomResource
    {
        [Output("config")]
        public Output<Outputs.EnvironmentConfig> Config { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:composer/environment:Environment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:composer/environment:Environment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, state, options);
        }
    }

    public sealed class EnvironmentArgs : Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.EnvironmentConfigArgs>? Config { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public EnvironmentArgs()
        {
        }
    }

    public sealed class EnvironmentState : Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.EnvironmentConfigGetArgs>? Config { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public EnvironmentState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EnvironmentConfigArgs : Pulumi.ResourceArgs
    {
        [Input("airflowUri")]
        public Input<string>? AirflowUri { get; set; }

        [Input("dagGcsPrefix")]
        public Input<string>? DagGcsPrefix { get; set; }

        [Input("gkeCluster")]
        public Input<string>? GkeCluster { get; set; }

        [Input("nodeConfig")]
        public Input<EnvironmentConfigNodeConfigArgs>? NodeConfig { get; set; }

        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("privateEnvironmentConfig")]
        public Input<EnvironmentConfigPrivateEnvironmentConfigArgs>? PrivateEnvironmentConfig { get; set; }

        [Input("softwareConfig")]
        public Input<EnvironmentConfigSoftwareConfigArgs>? SoftwareConfig { get; set; }

        public EnvironmentConfigArgs()
        {
        }
    }

    public sealed class EnvironmentConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("airflowUri")]
        public Input<string>? AirflowUri { get; set; }

        [Input("dagGcsPrefix")]
        public Input<string>? DagGcsPrefix { get; set; }

        [Input("gkeCluster")]
        public Input<string>? GkeCluster { get; set; }

        [Input("nodeConfig")]
        public Input<EnvironmentConfigNodeConfigGetArgs>? NodeConfig { get; set; }

        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("privateEnvironmentConfig")]
        public Input<EnvironmentConfigPrivateEnvironmentConfigGetArgs>? PrivateEnvironmentConfig { get; set; }

        [Input("softwareConfig")]
        public Input<EnvironmentConfigSoftwareConfigGetArgs>? SoftwareConfig { get; set; }

        public EnvironmentConfigGetArgs()
        {
        }
    }

    public sealed class EnvironmentConfigNodeConfigArgs : Pulumi.ResourceArgs
    {
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        [Input("ipAllocationPolicy")]
        public Input<EnvironmentConfigNodeConfigIpAllocationPolicyArgs>? IpAllocationPolicy { get; set; }

        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("oauthScopes")]
        private InputList<string>? _oauthScopes;
        public InputList<string> OauthScopes
        {
            get => _oauthScopes ?? (_oauthScopes = new InputList<string>());
            set => _oauthScopes = value;
        }

        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public EnvironmentConfigNodeConfigArgs()
        {
        }
    }

    public sealed class EnvironmentConfigNodeConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        [Input("ipAllocationPolicy")]
        public Input<EnvironmentConfigNodeConfigIpAllocationPolicyGetArgs>? IpAllocationPolicy { get; set; }

        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("oauthScopes")]
        private InputList<string>? _oauthScopes;
        public InputList<string> OauthScopes
        {
            get => _oauthScopes ?? (_oauthScopes = new InputList<string>());
            set => _oauthScopes = value;
        }

        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public EnvironmentConfigNodeConfigGetArgs()
        {
        }
    }

    public sealed class EnvironmentConfigNodeConfigIpAllocationPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("clusterIpv4CidrBlock")]
        public Input<string>? ClusterIpv4CidrBlock { get; set; }

        [Input("clusterSecondaryRangeName")]
        public Input<string>? ClusterSecondaryRangeName { get; set; }

        [Input("servicesIpv4CidrBlock")]
        public Input<string>? ServicesIpv4CidrBlock { get; set; }

        [Input("servicesSecondaryRangeName")]
        public Input<string>? ServicesSecondaryRangeName { get; set; }

        [Input("useIpAliases")]
        public Input<bool>? UseIpAliases { get; set; }

        public EnvironmentConfigNodeConfigIpAllocationPolicyArgs()
        {
        }
    }

    public sealed class EnvironmentConfigNodeConfigIpAllocationPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("clusterIpv4CidrBlock")]
        public Input<string>? ClusterIpv4CidrBlock { get; set; }

        [Input("clusterSecondaryRangeName")]
        public Input<string>? ClusterSecondaryRangeName { get; set; }

        [Input("servicesIpv4CidrBlock")]
        public Input<string>? ServicesIpv4CidrBlock { get; set; }

        [Input("servicesSecondaryRangeName")]
        public Input<string>? ServicesSecondaryRangeName { get; set; }

        [Input("useIpAliases")]
        public Input<bool>? UseIpAliases { get; set; }

        public EnvironmentConfigNodeConfigIpAllocationPolicyGetArgs()
        {
        }
    }

    public sealed class EnvironmentConfigPrivateEnvironmentConfigArgs : Pulumi.ResourceArgs
    {
        [Input("enablePrivateEndpoint")]
        public Input<bool>? EnablePrivateEndpoint { get; set; }

        [Input("masterIpv4CidrBlock")]
        public Input<string>? MasterIpv4CidrBlock { get; set; }

        public EnvironmentConfigPrivateEnvironmentConfigArgs()
        {
        }
    }

    public sealed class EnvironmentConfigPrivateEnvironmentConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("enablePrivateEndpoint")]
        public Input<bool>? EnablePrivateEndpoint { get; set; }

        [Input("masterIpv4CidrBlock")]
        public Input<string>? MasterIpv4CidrBlock { get; set; }

        public EnvironmentConfigPrivateEnvironmentConfigGetArgs()
        {
        }
    }

    public sealed class EnvironmentConfigSoftwareConfigArgs : Pulumi.ResourceArgs
    {
        [Input("airflowConfigOverrides")]
        private InputMap<string>? _airflowConfigOverrides;
        public InputMap<string> AirflowConfigOverrides
        {
            get => _airflowConfigOverrides ?? (_airflowConfigOverrides = new InputMap<string>());
            set => _airflowConfigOverrides = value;
        }

        [Input("envVariables")]
        private InputMap<string>? _envVariables;
        public InputMap<string> EnvVariables
        {
            get => _envVariables ?? (_envVariables = new InputMap<string>());
            set => _envVariables = value;
        }

        [Input("imageVersion")]
        public Input<string>? ImageVersion { get; set; }

        [Input("pypiPackages")]
        private InputMap<string>? _pypiPackages;
        public InputMap<string> PypiPackages
        {
            get => _pypiPackages ?? (_pypiPackages = new InputMap<string>());
            set => _pypiPackages = value;
        }

        [Input("pythonVersion")]
        public Input<string>? PythonVersion { get; set; }

        public EnvironmentConfigSoftwareConfigArgs()
        {
        }
    }

    public sealed class EnvironmentConfigSoftwareConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("airflowConfigOverrides")]
        private InputMap<string>? _airflowConfigOverrides;
        public InputMap<string> AirflowConfigOverrides
        {
            get => _airflowConfigOverrides ?? (_airflowConfigOverrides = new InputMap<string>());
            set => _airflowConfigOverrides = value;
        }

        [Input("envVariables")]
        private InputMap<string>? _envVariables;
        public InputMap<string> EnvVariables
        {
            get => _envVariables ?? (_envVariables = new InputMap<string>());
            set => _envVariables = value;
        }

        [Input("imageVersion")]
        public Input<string>? ImageVersion { get; set; }

        [Input("pypiPackages")]
        private InputMap<string>? _pypiPackages;
        public InputMap<string> PypiPackages
        {
            get => _pypiPackages ?? (_pypiPackages = new InputMap<string>());
            set => _pypiPackages = value;
        }

        [Input("pythonVersion")]
        public Input<string>? PythonVersion { get; set; }

        public EnvironmentConfigSoftwareConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EnvironmentConfig
    {
        public readonly string AirflowUri;
        public readonly string DagGcsPrefix;
        public readonly string GkeCluster;
        public readonly EnvironmentConfigNodeConfig NodeConfig;
        public readonly int NodeCount;
        public readonly EnvironmentConfigPrivateEnvironmentConfig PrivateEnvironmentConfig;
        public readonly EnvironmentConfigSoftwareConfig SoftwareConfig;

        [OutputConstructor]
        private EnvironmentConfig(
            string airflowUri,
            string dagGcsPrefix,
            string gkeCluster,
            EnvironmentConfigNodeConfig nodeConfig,
            int nodeCount,
            EnvironmentConfigPrivateEnvironmentConfig privateEnvironmentConfig,
            EnvironmentConfigSoftwareConfig softwareConfig)
        {
            AirflowUri = airflowUri;
            DagGcsPrefix = dagGcsPrefix;
            GkeCluster = gkeCluster;
            NodeConfig = nodeConfig;
            NodeCount = nodeCount;
            PrivateEnvironmentConfig = privateEnvironmentConfig;
            SoftwareConfig = softwareConfig;
        }
    }

    [OutputType]
    public sealed class EnvironmentConfigNodeConfig
    {
        public readonly int DiskSizeGb;
        public readonly EnvironmentConfigNodeConfigIpAllocationPolicy IpAllocationPolicy;
        public readonly string MachineType;
        public readonly string Network;
        public readonly ImmutableArray<string> OauthScopes;
        public readonly string ServiceAccount;
        public readonly string? Subnetwork;
        public readonly ImmutableArray<string> Tags;
        public readonly string Zone;

        [OutputConstructor]
        private EnvironmentConfigNodeConfig(
            int diskSizeGb,
            EnvironmentConfigNodeConfigIpAllocationPolicy ipAllocationPolicy,
            string machineType,
            string network,
            ImmutableArray<string> oauthScopes,
            string serviceAccount,
            string? subnetwork,
            ImmutableArray<string> tags,
            string zone)
        {
            DiskSizeGb = diskSizeGb;
            IpAllocationPolicy = ipAllocationPolicy;
            MachineType = machineType;
            Network = network;
            OauthScopes = oauthScopes;
            ServiceAccount = serviceAccount;
            Subnetwork = subnetwork;
            Tags = tags;
            Zone = zone;
        }
    }

    [OutputType]
    public sealed class EnvironmentConfigNodeConfigIpAllocationPolicy
    {
        public readonly string? ClusterIpv4CidrBlock;
        public readonly string? ClusterSecondaryRangeName;
        public readonly string? ServicesIpv4CidrBlock;
        public readonly string? ServicesSecondaryRangeName;
        public readonly bool? UseIpAliases;

        [OutputConstructor]
        private EnvironmentConfigNodeConfigIpAllocationPolicy(
            string? clusterIpv4CidrBlock,
            string? clusterSecondaryRangeName,
            string? servicesIpv4CidrBlock,
            string? servicesSecondaryRangeName,
            bool? useIpAliases)
        {
            ClusterIpv4CidrBlock = clusterIpv4CidrBlock;
            ClusterSecondaryRangeName = clusterSecondaryRangeName;
            ServicesIpv4CidrBlock = servicesIpv4CidrBlock;
            ServicesSecondaryRangeName = servicesSecondaryRangeName;
            UseIpAliases = useIpAliases;
        }
    }

    [OutputType]
    public sealed class EnvironmentConfigPrivateEnvironmentConfig
    {
        public readonly bool? EnablePrivateEndpoint;
        public readonly string? MasterIpv4CidrBlock;

        [OutputConstructor]
        private EnvironmentConfigPrivateEnvironmentConfig(
            bool? enablePrivateEndpoint,
            string? masterIpv4CidrBlock)
        {
            EnablePrivateEndpoint = enablePrivateEndpoint;
            MasterIpv4CidrBlock = masterIpv4CidrBlock;
        }
    }

    [OutputType]
    public sealed class EnvironmentConfigSoftwareConfig
    {
        public readonly ImmutableDictionary<string, string>? AirflowConfigOverrides;
        public readonly ImmutableDictionary<string, string>? EnvVariables;
        public readonly string ImageVersion;
        public readonly ImmutableDictionary<string, string>? PypiPackages;
        public readonly string PythonVersion;

        [OutputConstructor]
        private EnvironmentConfigSoftwareConfig(
            ImmutableDictionary<string, string>? airflowConfigOverrides,
            ImmutableDictionary<string, string>? envVariables,
            string imageVersion,
            ImmutableDictionary<string, string>? pypiPackages,
            string pythonVersion)
        {
            AirflowConfigOverrides = airflowConfigOverrides;
            EnvVariables = envVariables;
            ImageVersion = imageVersion;
            PypiPackages = pypiPackages;
            PythonVersion = pythonVersion;
        }
    }
    }
}
