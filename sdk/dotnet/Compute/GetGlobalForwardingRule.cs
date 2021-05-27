// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetGlobalForwardingRule
    {
        /// <summary>
        /// Get a global forwarding rule within GCE from its name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_forwarding_rule = Output.Create(Gcp.Compute.GetGlobalForwardingRule.InvokeAsync(new Gcp.Compute.GetGlobalForwardingRuleArgs
        ///         {
        ///             Name = "forwarding-rule-global",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGlobalForwardingRuleResult> InvokeAsync(GetGlobalForwardingRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGlobalForwardingRuleResult>("gcp:compute/getGlobalForwardingRule:getGlobalForwardingRule", args ?? new GetGlobalForwardingRuleArgs(), options.WithVersion());

        public static Output<GetGlobalForwardingRuleResult> Invoke(GetGlobalForwardingRuleOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.Project.Box()
            ).Apply(a => {
                    var args = new GetGlobalForwardingRuleArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.Project));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetGlobalForwardingRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the global forwarding rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetGlobalForwardingRuleArgs()
        {
        }
    }

    public sealed class GetGlobalForwardingRuleOutputArgs
    {
        /// <summary>
        /// The name of the global forwarding rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGlobalForwardingRuleOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGlobalForwardingRuleResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpAddress;
        public readonly string IpProtocol;
        public readonly string IpVersion;
        public readonly string LabelFingerprint;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string LoadBalancingScheme;
        public readonly ImmutableArray<Outputs.GetGlobalForwardingRuleMetadataFilterResult> MetadataFilters;
        public readonly string Name;
        public readonly string Network;
        public readonly string PortRange;
        public readonly string? Project;
        public readonly string SelfLink;
        public readonly string Target;

        [OutputConstructor]
        private GetGlobalForwardingRuleResult(
            string description,

            string id,

            string ipAddress,

            string ipProtocol,

            string ipVersion,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string loadBalancingScheme,

            ImmutableArray<Outputs.GetGlobalForwardingRuleMetadataFilterResult> metadataFilters,

            string name,

            string network,

            string portRange,

            string? project,

            string selfLink,

            string target)
        {
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            IpProtocol = ipProtocol;
            IpVersion = ipVersion;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            LoadBalancingScheme = loadBalancingScheme;
            MetadataFilters = metadataFilters;
            Name = name;
            Network = network;
            PortRange = portRange;
            Project = project;
            SelfLink = selfLink;
            Target = target;
        }
    }
}
