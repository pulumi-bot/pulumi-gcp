// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public partial class Autoscalar : Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for
        /// an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are
        /// specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        /// </summary>
        [Output("autoscalingPolicy")]
        public Output<Outputs.AutoscalarAutoscalingPolicy> AutoscalingPolicy { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        /// <summary>
        /// URL of the zone where the instance group resides.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Autoscalar resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Autoscalar(string name, AutoscalarArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/autoscalar:Autoscalar", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Autoscalar(string name, Input<string> id, AutoscalarState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/autoscalar:Autoscalar", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Autoscalar resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Autoscalar Get(string name, Input<string> id, AutoscalarState? state = null, CustomResourceOptions? options = null)
        {
            return new Autoscalar(name, id, state, options);
        }
    }

    public sealed class AutoscalarArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for
        /// an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are
        /// specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        /// </summary>
        [Input("autoscalingPolicy", required: true)]
        public Input<Inputs.AutoscalarAutoscalingPolicyArgs> AutoscalingPolicy { get; set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// URL of the zone where the instance group resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AutoscalarArgs()
        {
        }
    }

    public sealed class AutoscalarState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for
        /// an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are
        /// specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<Inputs.AutoscalarAutoscalingPolicyGetArgs>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// URL of the zone where the instance group resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AutoscalarState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class AutoscalarAutoscalingPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("cooldownPeriod")]
        public Input<int>? CooldownPeriod { get; set; }

        [Input("cpuUtilization")]
        public Input<AutoscalarAutoscalingPolicyCpuUtilizationArgs>? CpuUtilization { get; set; }

        [Input("loadBalancingUtilization")]
        public Input<AutoscalarAutoscalingPolicyLoadBalancingUtilizationArgs>? LoadBalancingUtilization { get; set; }

        [Input("maxReplicas", required: true)]
        public Input<int> MaxReplicas { get; set; } = null!;

        [Input("metrics")]
        private InputList<AutoscalarAutoscalingPolicyMetricsArgs>? _metrics;
        public InputList<AutoscalarAutoscalingPolicyMetricsArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<AutoscalarAutoscalingPolicyMetricsArgs>());
            set => _metrics = value;
        }

        [Input("minReplicas", required: true)]
        public Input<int> MinReplicas { get; set; } = null!;

        public AutoscalarAutoscalingPolicyArgs()
        {
        }
    }

    public sealed class AutoscalarAutoscalingPolicyCpuUtilizationArgs : Pulumi.ResourceArgs
    {
        [Input("target", required: true)]
        public Input<double> Target { get; set; } = null!;

        public AutoscalarAutoscalingPolicyCpuUtilizationArgs()
        {
        }
    }

    public sealed class AutoscalarAutoscalingPolicyCpuUtilizationGetArgs : Pulumi.ResourceArgs
    {
        [Input("target", required: true)]
        public Input<double> Target { get; set; } = null!;

        public AutoscalarAutoscalingPolicyCpuUtilizationGetArgs()
        {
        }
    }

    public sealed class AutoscalarAutoscalingPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("cooldownPeriod")]
        public Input<int>? CooldownPeriod { get; set; }

        [Input("cpuUtilization")]
        public Input<AutoscalarAutoscalingPolicyCpuUtilizationGetArgs>? CpuUtilization { get; set; }

        [Input("loadBalancingUtilization")]
        public Input<AutoscalarAutoscalingPolicyLoadBalancingUtilizationGetArgs>? LoadBalancingUtilization { get; set; }

        [Input("maxReplicas", required: true)]
        public Input<int> MaxReplicas { get; set; } = null!;

        [Input("metrics")]
        private InputList<AutoscalarAutoscalingPolicyMetricsGetArgs>? _metrics;
        public InputList<AutoscalarAutoscalingPolicyMetricsGetArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<AutoscalarAutoscalingPolicyMetricsGetArgs>());
            set => _metrics = value;
        }

        [Input("minReplicas", required: true)]
        public Input<int> MinReplicas { get; set; } = null!;

        public AutoscalarAutoscalingPolicyGetArgs()
        {
        }
    }

    public sealed class AutoscalarAutoscalingPolicyLoadBalancingUtilizationArgs : Pulumi.ResourceArgs
    {
        [Input("target", required: true)]
        public Input<double> Target { get; set; } = null!;

        public AutoscalarAutoscalingPolicyLoadBalancingUtilizationArgs()
        {
        }
    }

    public sealed class AutoscalarAutoscalingPolicyLoadBalancingUtilizationGetArgs : Pulumi.ResourceArgs
    {
        [Input("target", required: true)]
        public Input<double> Target { get; set; } = null!;

        public AutoscalarAutoscalingPolicyLoadBalancingUtilizationGetArgs()
        {
        }
    }

    public sealed class AutoscalarAutoscalingPolicyMetricsArgs : Pulumi.ResourceArgs
    {
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("singleInstanceAssignment")]
        public Input<double>? SingleInstanceAssignment { get; set; }

        [Input("target")]
        public Input<double>? Target { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public AutoscalarAutoscalingPolicyMetricsArgs()
        {
        }
    }

    public sealed class AutoscalarAutoscalingPolicyMetricsGetArgs : Pulumi.ResourceArgs
    {
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("singleInstanceAssignment")]
        public Input<double>? SingleInstanceAssignment { get; set; }

        [Input("target")]
        public Input<double>? Target { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public AutoscalarAutoscalingPolicyMetricsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AutoscalarAutoscalingPolicy
    {
        public readonly int? CooldownPeriod;
        public readonly AutoscalarAutoscalingPolicyCpuUtilization CpuUtilization;
        public readonly AutoscalarAutoscalingPolicyLoadBalancingUtilization? LoadBalancingUtilization;
        public readonly int MaxReplicas;
        public readonly ImmutableArray<AutoscalarAutoscalingPolicyMetrics> Metrics;
        public readonly int MinReplicas;

        [OutputConstructor]
        private AutoscalarAutoscalingPolicy(
            int? cooldownPeriod,
            AutoscalarAutoscalingPolicyCpuUtilization cpuUtilization,
            AutoscalarAutoscalingPolicyLoadBalancingUtilization? loadBalancingUtilization,
            int maxReplicas,
            ImmutableArray<AutoscalarAutoscalingPolicyMetrics> metrics,
            int minReplicas)
        {
            CooldownPeriod = cooldownPeriod;
            CpuUtilization = cpuUtilization;
            LoadBalancingUtilization = loadBalancingUtilization;
            MaxReplicas = maxReplicas;
            Metrics = metrics;
            MinReplicas = minReplicas;
        }
    }

    [OutputType]
    public sealed class AutoscalarAutoscalingPolicyCpuUtilization
    {
        public readonly double Target;

        [OutputConstructor]
        private AutoscalarAutoscalingPolicyCpuUtilization(double target)
        {
            Target = target;
        }
    }

    [OutputType]
    public sealed class AutoscalarAutoscalingPolicyLoadBalancingUtilization
    {
        public readonly double Target;

        [OutputConstructor]
        private AutoscalarAutoscalingPolicyLoadBalancingUtilization(double target)
        {
            Target = target;
        }
    }

    [OutputType]
    public sealed class AutoscalarAutoscalingPolicyMetrics
    {
        public readonly string? Filter;
        public readonly string Name;
        public readonly double? SingleInstanceAssignment;
        public readonly double? Target;
        public readonly string? Type;

        [OutputConstructor]
        private AutoscalarAutoscalingPolicyMetrics(
            string? filter,
            string name,
            double? singleInstanceAssignment,
            double? target,
            string? type)
        {
            Filter = filter;
            Name = name;
            SingleInstanceAssignment = singleInstanceAssignment;
            Target = target;
            Type = type;
        }
    }
    }
}
