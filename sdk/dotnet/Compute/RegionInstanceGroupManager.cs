// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// The Google Compute Engine Regional Instance Group Manager API creates and manages pools
    /// of homogeneous Compute Engine virtual machine instances from a common instance
    /// template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
    /// and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
    /// 
    /// &gt; **Note:** Use [gcp.compute.InstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a single-zone instance group manager.
    /// </summary>
    public partial class RegionInstanceGroupManager : Pulumi.CustomResource
    {
        /// <summary>
        /// The autohealing policies for this managed instance
        /// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        /// </summary>
        [Output("autoHealingPolicies")]
        public Output<Outputs.RegionInstanceGroupManagerAutoHealingPolicies?> AutoHealingPolicies { get; private set; } = null!;

        /// <summary>
        /// The base instance name to use for
        /// instances in this group. The value must be a valid
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
        /// are lowercase letters, numbers, and hyphens (-). Instances are named by
        /// appending a hyphen and a random four-character string to the base instance
        /// name.
        /// </summary>
        [Output("baseInstanceName")]
        public Output<string> BaseInstanceName { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the instance
        /// group manager.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The distribution policy for this managed instance
        /// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
        /// - - -
        /// </summary>
        [Output("distributionPolicyZones")]
        public Output<ImmutableArray<string>> DistributionPolicyZones { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the instance group manager.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The full URL of the instance group created by the manager.
        /// </summary>
        [Output("instanceGroup")]
        public Output<string> InstanceGroup { get; private set; } = null!;

        /// <summary>
        /// The name of the instance group manager. Must be 1-63
        /// characters long and comply with
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
        /// include lowercase letters, numbers, and hyphens.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The named port configuration. See the section below
        /// for details on configuration.
        /// </summary>
        [Output("namedPorts")]
        public Output<ImmutableArray<Outputs.RegionInstanceGroupManagerNamedPort>> NamedPorts { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region where the managed instance group resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URL of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The full URL of all target pools to which new
        /// instances in the group are added. Updating the target pools attribute does
        /// not affect existing instances.
        /// </summary>
        [Output("targetPools")]
        public Output<ImmutableArray<string>> TargetPools { get; private set; } = null!;

        /// <summary>
        /// The target number of running instances for this managed
        /// instance group. This value should always be explicitly set unless this resource is attached to
        /// an autoscaler, in which case it should never be set. Defaults to `0`.
        /// </summary>
        [Output("targetSize")]
        public Output<int> TargetSize { get; private set; } = null!;

        /// <summary>
        /// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
        /// </summary>
        [Output("updatePolicy")]
        public Output<Outputs.RegionInstanceGroupManagerUpdatePolicy> UpdatePolicy { get; private set; } = null!;

        /// <summary>
        /// Application versions managed by this instance group. Each
        /// version deals with a specific instance template, allowing canary release scenarios.
        /// Structure is documented below.
        /// </summary>
        [Output("versions")]
        public Output<ImmutableArray<Outputs.RegionInstanceGroupManagerVersion>> Versions { get; private set; } = null!;

        /// <summary>
        /// Whether to wait for all instances to be created/updated before
        /// returning. Note that if this is set to true and the operation does not succeed, the provider will
        /// continue trying until it times out.
        /// </summary>
        [Output("waitForInstances")]
        public Output<bool?> WaitForInstances { get; private set; } = null!;


        /// <summary>
        /// Create a RegionInstanceGroupManager resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionInstanceGroupManager(string name, RegionInstanceGroupManagerArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager", name, args ?? new RegionInstanceGroupManagerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionInstanceGroupManager(string name, Input<string> id, RegionInstanceGroupManagerState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegionInstanceGroupManager resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionInstanceGroupManager Get(string name, Input<string> id, RegionInstanceGroupManagerState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionInstanceGroupManager(name, id, state, options);
        }
    }

    public sealed class RegionInstanceGroupManagerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The autohealing policies for this managed instance
        /// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        /// </summary>
        [Input("autoHealingPolicies")]
        public Input<Inputs.RegionInstanceGroupManagerAutoHealingPoliciesArgs>? AutoHealingPolicies { get; set; }

        /// <summary>
        /// The base instance name to use for
        /// instances in this group. The value must be a valid
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
        /// are lowercase letters, numbers, and hyphens (-). Instances are named by
        /// appending a hyphen and a random four-character string to the base instance
        /// name.
        /// </summary>
        [Input("baseInstanceName", required: true)]
        public Input<string> BaseInstanceName { get; set; } = null!;

        /// <summary>
        /// An optional textual description of the instance
        /// group manager.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("distributionPolicyZones")]
        private InputList<string>? _distributionPolicyZones;

        /// <summary>
        /// The distribution policy for this managed instance
        /// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
        /// - - -
        /// </summary>
        public InputList<string> DistributionPolicyZones
        {
            get => _distributionPolicyZones ?? (_distributionPolicyZones = new InputList<string>());
            set => _distributionPolicyZones = value;
        }

        /// <summary>
        /// The name of the instance group manager. Must be 1-63
        /// characters long and comply with
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
        /// include lowercase letters, numbers, and hyphens.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namedPorts")]
        private InputList<Inputs.RegionInstanceGroupManagerNamedPortArgs>? _namedPorts;

        /// <summary>
        /// The named port configuration. See the section below
        /// for details on configuration.
        /// </summary>
        public InputList<Inputs.RegionInstanceGroupManagerNamedPortArgs> NamedPorts
        {
            get => _namedPorts ?? (_namedPorts = new InputList<Inputs.RegionInstanceGroupManagerNamedPortArgs>());
            set => _namedPorts = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region where the managed instance group resides.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("targetPools")]
        private InputList<string>? _targetPools;

        /// <summary>
        /// The full URL of all target pools to which new
        /// instances in the group are added. Updating the target pools attribute does
        /// not affect existing instances.
        /// </summary>
        public InputList<string> TargetPools
        {
            get => _targetPools ?? (_targetPools = new InputList<string>());
            set => _targetPools = value;
        }

        /// <summary>
        /// The target number of running instances for this managed
        /// instance group. This value should always be explicitly set unless this resource is attached to
        /// an autoscaler, in which case it should never be set. Defaults to `0`.
        /// </summary>
        [Input("targetSize")]
        public Input<int>? TargetSize { get; set; }

        /// <summary>
        /// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
        /// </summary>
        [Input("updatePolicy")]
        public Input<Inputs.RegionInstanceGroupManagerUpdatePolicyArgs>? UpdatePolicy { get; set; }

        [Input("versions", required: true)]
        private InputList<Inputs.RegionInstanceGroupManagerVersionArgs>? _versions;

        /// <summary>
        /// Application versions managed by this instance group. Each
        /// version deals with a specific instance template, allowing canary release scenarios.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RegionInstanceGroupManagerVersionArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.RegionInstanceGroupManagerVersionArgs>());
            set => _versions = value;
        }

        /// <summary>
        /// Whether to wait for all instances to be created/updated before
        /// returning. Note that if this is set to true and the operation does not succeed, the provider will
        /// continue trying until it times out.
        /// </summary>
        [Input("waitForInstances")]
        public Input<bool>? WaitForInstances { get; set; }

        public RegionInstanceGroupManagerArgs()
        {
        }
    }

    public sealed class RegionInstanceGroupManagerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The autohealing policies for this managed instance
        /// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        /// </summary>
        [Input("autoHealingPolicies")]
        public Input<Inputs.RegionInstanceGroupManagerAutoHealingPoliciesGetArgs>? AutoHealingPolicies { get; set; }

        /// <summary>
        /// The base instance name to use for
        /// instances in this group. The value must be a valid
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
        /// are lowercase letters, numbers, and hyphens (-). Instances are named by
        /// appending a hyphen and a random four-character string to the base instance
        /// name.
        /// </summary>
        [Input("baseInstanceName")]
        public Input<string>? BaseInstanceName { get; set; }

        /// <summary>
        /// An optional textual description of the instance
        /// group manager.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("distributionPolicyZones")]
        private InputList<string>? _distributionPolicyZones;

        /// <summary>
        /// The distribution policy for this managed instance
        /// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
        /// - - -
        /// </summary>
        public InputList<string> DistributionPolicyZones
        {
            get => _distributionPolicyZones ?? (_distributionPolicyZones = new InputList<string>());
            set => _distributionPolicyZones = value;
        }

        /// <summary>
        /// The fingerprint of the instance group manager.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The full URL of the instance group created by the manager.
        /// </summary>
        [Input("instanceGroup")]
        public Input<string>? InstanceGroup { get; set; }

        /// <summary>
        /// The name of the instance group manager. Must be 1-63
        /// characters long and comply with
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
        /// include lowercase letters, numbers, and hyphens.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namedPorts")]
        private InputList<Inputs.RegionInstanceGroupManagerNamedPortGetArgs>? _namedPorts;

        /// <summary>
        /// The named port configuration. See the section below
        /// for details on configuration.
        /// </summary>
        public InputList<Inputs.RegionInstanceGroupManagerNamedPortGetArgs> NamedPorts
        {
            get => _namedPorts ?? (_namedPorts = new InputList<Inputs.RegionInstanceGroupManagerNamedPortGetArgs>());
            set => _namedPorts = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region where the managed instance group resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URL of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("targetPools")]
        private InputList<string>? _targetPools;

        /// <summary>
        /// The full URL of all target pools to which new
        /// instances in the group are added. Updating the target pools attribute does
        /// not affect existing instances.
        /// </summary>
        public InputList<string> TargetPools
        {
            get => _targetPools ?? (_targetPools = new InputList<string>());
            set => _targetPools = value;
        }

        /// <summary>
        /// The target number of running instances for this managed
        /// instance group. This value should always be explicitly set unless this resource is attached to
        /// an autoscaler, in which case it should never be set. Defaults to `0`.
        /// </summary>
        [Input("targetSize")]
        public Input<int>? TargetSize { get; set; }

        /// <summary>
        /// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
        /// </summary>
        [Input("updatePolicy")]
        public Input<Inputs.RegionInstanceGroupManagerUpdatePolicyGetArgs>? UpdatePolicy { get; set; }

        [Input("versions")]
        private InputList<Inputs.RegionInstanceGroupManagerVersionGetArgs>? _versions;

        /// <summary>
        /// Application versions managed by this instance group. Each
        /// version deals with a specific instance template, allowing canary release scenarios.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RegionInstanceGroupManagerVersionGetArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.RegionInstanceGroupManagerVersionGetArgs>());
            set => _versions = value;
        }

        /// <summary>
        /// Whether to wait for all instances to be created/updated before
        /// returning. Note that if this is set to true and the operation does not succeed, the provider will
        /// continue trying until it times out.
        /// </summary>
        [Input("waitForInstances")]
        public Input<bool>? WaitForInstances { get; set; }

        public RegionInstanceGroupManagerState()
        {
        }
    }
}
