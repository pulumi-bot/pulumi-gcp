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
    /// Represents a VPN gateway running in GCP. This virtual device is managed
    /// by Google, but used only by you.
    /// 
    /// To get more information about VpnGateway, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetVpnGateways)
    /// 
    /// ## Example Usage
    /// ### Target Vpn Gateway Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network1 = new Gcp.Compute.Network("network1", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var targetGateway = new Gcp.Compute.VPNGateway("targetGateway", new Gcp.Compute.VPNGatewayArgs
    ///         {
    ///             Network = network1.Id,
    ///         });
    ///         var vpnStaticIp = new Gcp.Compute.Address("vpnStaticIp", new Gcp.Compute.AddressArgs
    ///         {
    ///         });
    ///         var frEsp = new Gcp.Compute.ForwardingRule("frEsp", new Gcp.Compute.ForwardingRuleArgs
    ///         {
    ///             IpProtocol = "ESP",
    ///             IpAddress = vpnStaticIp.IPAddress,
    ///             Target = targetGateway.Id,
    ///         });
    ///         var frUdp500 = new Gcp.Compute.ForwardingRule("frUdp500", new Gcp.Compute.ForwardingRuleArgs
    ///         {
    ///             IpProtocol = "UDP",
    ///             PortRange = "500",
    ///             IpAddress = vpnStaticIp.IPAddress,
    ///             Target = targetGateway.Id,
    ///         });
    ///         var frUdp4500 = new Gcp.Compute.ForwardingRule("frUdp4500", new Gcp.Compute.ForwardingRuleArgs
    ///         {
    ///             IpProtocol = "UDP",
    ///             PortRange = "4500",
    ///             IpAddress = vpnStaticIp.IPAddress,
    ///             Target = targetGateway.Id,
    ///         });
    ///         var tunnel1 = new Gcp.Compute.VPNTunnel("tunnel1", new Gcp.Compute.VPNTunnelArgs
    ///         {
    ///             PeerIp = "15.0.0.120",
    ///             SharedSecret = "a secret message",
    ///             TargetVpnGateway = targetGateway.Id,
    ///         }, new CustomResourceOptions {
    ///             DependsOn = 
    ///             {
    ///                 frEsp,
    ///                 frUdp500,
    ///                 frUdp4500,
    ///             },
    ///         });
    ///         var route1 = new Gcp.Compute.Route("route1", new Gcp.Compute.RouteArgs
    ///         {
    ///             Network = network1.Name,
    ///             DestRange = "15.0.0.0/24",
    ///             Priority = 1000,
    ///             NextHopVpnTunnel = tunnel1.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class VPNGateway : Pulumi.CustomResource
    {
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
        /// The unique identifier for the resource.
        /// </summary>
        [Output("gatewayId")]
        public Output<int> GatewayId { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a VPNGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VPNGateway(string name, VPNGatewayArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/vPNGateway:VPNGateway", name, args ?? new VPNGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VPNGateway(string name, Input<string> id, VPNGatewayState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/vPNGateway:VPNGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VPNGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VPNGateway Get(string name, Input<string> id, VPNGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new VPNGateway(name, id, state, options);
        }
    }

    public sealed class VPNGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public VPNGatewayArgs()
        {
        }
    }

    public sealed class VPNGatewayState : Pulumi.ResourceArgs
    {
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
        /// The unique identifier for the resource.
        /// </summary>
        [Input("gatewayId")]
        public Input<int>? GatewayId { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public VPNGatewayState()
        {
        }
    }
}
