// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetVPNGateway
    {
        /// <summary>
        /// Get a VPN gateway within GCE from its name.
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
        ///         var my_vpn_gateway = Output.Create(Gcp.Compute.GetVPNGateway.InvokeAsync(new Gcp.Compute.GetVPNGatewayArgs
        ///         {
        ///             Name = "vpn-gateway-us-east1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVPNGatewayResult> InvokeAsync(GetVPNGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVPNGatewayResult>("gcp:compute/getVPNGateway:getVPNGateway", args ?? new GetVPNGatewayArgs(), options.WithVersion());
    }


    public sealed class GetVPNGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the VPN gateway.
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

        public GetVPNGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVPNGatewayResult
    {
        /// <summary>
        /// Description of this VPN gateway.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The network of this VPN gateway.
        /// </summary>
        public readonly string Network;
        public readonly string Project;
        /// <summary>
        /// Region of this VPN gateway.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The URI of the resource.
        /// </summary>
        public readonly string SelfLink;

        [OutputConstructor]
        private GetVPNGatewayResult(
            string description,

            string id,

            string name,

            string network,

            string project,

            string region,

            string selfLink)
        {
            Description = description;
            Id = id;
            Name = name;
            Network = network;
            Project = project;
            Region = region;
            SelfLink = selfLink;
        }
    }
}
