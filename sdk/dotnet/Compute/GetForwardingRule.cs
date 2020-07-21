// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetForwardingRule
    {
        /// <summary>
        /// Get a forwarding rule within GCE from its name.
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
        ///         var my_forwarding_rule = Output.Create(Gcp.Compute.GetForwardingRule.InvokeAsync(new Gcp.Compute.GetForwardingRuleArgs
        ///         {
        ///             Name = "forwarding-rule-us-east1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetForwardingRuleResult> InvokeAsync(GetForwardingRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetForwardingRuleResult>("gcp:compute/getForwardingRule:getForwardingRule", args ?? new GetForwardingRuleArgs(), options.WithVersion());
    }


    public sealed class GetForwardingRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the forwarding rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the project region is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetForwardingRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetForwardingRuleResult
    {
        /// <summary>
        /// Backend service, if this forwarding rule has one.
        /// </summary>
        public readonly string BackendService;
        /// <summary>
        /// Description of this forwarding rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP address of this forwarding rule.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// IP protocol of this forwarding rule.
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// Type of load balancing of this forwarding rule.
        /// </summary>
        public readonly string LoadBalancingScheme;
        public readonly string Name;
        /// <summary>
        /// Network of this forwarding rule.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Port range, if this forwarding rule has one.
        /// </summary>
        public readonly string PortRange;
        /// <summary>
        /// List of ports to use for internal load balancing, if this forwarding rule has any.
        /// </summary>
        public readonly ImmutableArray<string> Ports;
        public readonly string Project;
        /// <summary>
        /// Region of this forwarding rule.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The URI of the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Subnetwork of this forwarding rule.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// URL of the target pool, if this forwarding rule has one.
        /// </summary>
        public readonly string Target;

        [OutputConstructor]
        private GetForwardingRuleResult(
            string backendService,

            string description,

            string id,

            string ipAddress,

            string ipProtocol,

            string loadBalancingScheme,

            string name,

            string network,

            string portRange,

            ImmutableArray<string> ports,

            string project,

            string region,

            string selfLink,

            string subnetwork,

            string target)
        {
            BackendService = backendService;
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            IpProtocol = ipProtocol;
            LoadBalancingScheme = loadBalancingScheme;
            Name = name;
            Network = network;
            PortRange = portRange;
            Ports = ports;
            Project = project;
            Region = region;
            SelfLink = selfLink;
            Subnetwork = subnetwork;
            Target = target;
        }
    }
}
