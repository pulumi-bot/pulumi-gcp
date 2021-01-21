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
    /// VPN tunnel resource.
    /// 
    /// To get more information about VpnTunnel, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels)
    /// * How-to Guides
    ///     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
    ///     * [Networks and Tunnel Routing](https://cloud.google.com/vpn/docs/concepts/choosing-networks-routing)
    /// 
    /// &gt; **Warning:** All arguments including `shared_secret` will be stored in the raw
    /// state as plain-text.
    /// 
    /// ## Example Usage
    /// ### Vpn Tunnel Basic
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
    ///         }, new CustomResourceOptions
    ///         {
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
    /// ### Vpn Tunnel Beta
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
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var targetGateway = new Gcp.Compute.VPNGateway("targetGateway", new Gcp.Compute.VPNGatewayArgs
    ///         {
    ///             Network = network1.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var vpnStaticIp = new Gcp.Compute.Address("vpnStaticIp", new Gcp.Compute.AddressArgs
    ///         {
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var frEsp = new Gcp.Compute.ForwardingRule("frEsp", new Gcp.Compute.ForwardingRuleArgs
    ///         {
    ///             IpProtocol = "ESP",
    ///             IpAddress = vpnStaticIp.IPAddress,
    ///             Target = targetGateway.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var frUdp500 = new Gcp.Compute.ForwardingRule("frUdp500", new Gcp.Compute.ForwardingRuleArgs
    ///         {
    ///             IpProtocol = "UDP",
    ///             PortRange = "500",
    ///             IpAddress = vpnStaticIp.IPAddress,
    ///             Target = targetGateway.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var frUdp4500 = new Gcp.Compute.ForwardingRule("frUdp4500", new Gcp.Compute.ForwardingRuleArgs
    ///         {
    ///             IpProtocol = "UDP",
    ///             PortRange = "4500",
    ///             IpAddress = vpnStaticIp.IPAddress,
    ///             Target = targetGateway.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var tunnel1 = new Gcp.Compute.VPNTunnel("tunnel1", new Gcp.Compute.VPNTunnelArgs
    ///         {
    ///             PeerIp = "15.0.0.120",
    ///             SharedSecret = "a secret message",
    ///             TargetVpnGateway = targetGateway.Id,
    ///             Labels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
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
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpnTunnel can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/vPNTunnel:VPNTunnel")]
    public partial class VPNTunnel : Pulumi.CustomResource
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
        /// Detailed status message for the VPN tunnel.
        /// </summary>
        [Output("detailedStatus")]
        public Output<string> DetailedStatus { get; private set; } = null!;

        /// <summary>
        /// IKE protocol version to use when establishing the VPN tunnel with
        /// peer VPN gateway.
        /// Acceptable IKE versions are 1 or 2. Default version is 2.
        /// </summary>
        [Output("ikeVersion")]
        public Output<int?> IkeVersion { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this VpnTunnel.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Local traffic selector to use when establishing the VPN tunnel with
        /// peer VPN gateway. The value should be a CIDR formatted string,
        /// for example `192.168.0.0/16`. The ranges should be disjoint.
        /// Only IPv4 is supported.
        /// </summary>
        [Output("localTrafficSelectors")]
        public Output<ImmutableArray<string>> LocalTrafficSelectors { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63
        /// characters long and match the regular expression
        /// `a-z?` which means the first character
        /// must be a lowercase letter, and all following characters must
        /// be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
        /// </summary>
        [Output("peerExternalGateway")]
        public Output<string?> PeerExternalGateway { get; private set; } = null!;

        /// <summary>
        /// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
        /// </summary>
        [Output("peerExternalGatewayInterface")]
        public Output<int?> PeerExternalGatewayInterface { get; private set; } = null!;

        /// <summary>
        /// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
        /// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
        /// ID in the peer GCP VPN gateway.
        /// This field must reference a `gcp.compute.HaVpnGateway` resource.
        /// </summary>
        [Output("peerGcpGateway")]
        public Output<string?> PeerGcpGateway { get; private set; } = null!;

        /// <summary>
        /// IP address of the peer VPN gateway. Only IPv4 is supported.
        /// </summary>
        [Output("peerIp")]
        public Output<string> PeerIp { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Remote traffic selector to use when establishing the VPN tunnel with
        /// peer VPN gateway. The value should be a CIDR formatted string,
        /// for example `192.168.0.0/16`. The ranges should be disjoint.
        /// Only IPv4 is supported.
        /// </summary>
        [Output("remoteTrafficSelectors")]
        public Output<ImmutableArray<string>> RemoteTrafficSelectors { get; private set; } = null!;

        /// <summary>
        /// URL of router resource to be used for dynamic routing.
        /// </summary>
        [Output("router")]
        public Output<string?> Router { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Shared secret used to set the secure session between the Cloud VPN
        /// gateway and the peer VPN gateway.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Output("sharedSecret")]
        public Output<string> SharedSecret { get; private set; } = null!;

        /// <summary>
        /// Hash of the shared secret.
        /// </summary>
        [Output("sharedSecretHash")]
        public Output<string> SharedSecretHash { get; private set; } = null!;

        /// <summary>
        /// URL of the Target VPN gateway with which this VPN tunnel is
        /// associated.
        /// </summary>
        [Output("targetVpnGateway")]
        public Output<string?> TargetVpnGateway { get; private set; } = null!;

        /// <summary>
        /// The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Output("tunnelId")]
        public Output<string> TunnelId { get; private set; } = null!;

        /// <summary>
        /// URL of the VPN gateway with which this VPN tunnel is associated.
        /// This must be used if a High Availability VPN gateway resource is created.
        /// This field must reference a `gcp.compute.HaVpnGateway` resource.
        /// </summary>
        [Output("vpnGateway")]
        public Output<string?> VpnGateway { get; private set; } = null!;

        /// <summary>
        /// The interface ID of the VPN gateway with which this VPN tunnel is associated.
        /// </summary>
        [Output("vpnGatewayInterface")]
        public Output<int?> VpnGatewayInterface { get; private set; } = null!;


        /// <summary>
        /// Create a VPNTunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VPNTunnel(string name, VPNTunnelArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/vPNTunnel:VPNTunnel", name, args ?? new VPNTunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VPNTunnel(string name, Input<string> id, VPNTunnelState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/vPNTunnel:VPNTunnel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VPNTunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VPNTunnel Get(string name, Input<string> id, VPNTunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new VPNTunnel(name, id, state, options);
        }
    }

    public sealed class VPNTunnelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IKE protocol version to use when establishing the VPN tunnel with
        /// peer VPN gateway.
        /// Acceptable IKE versions are 1 or 2. Default version is 2.
        /// </summary>
        [Input("ikeVersion")]
        public Input<int>? IkeVersion { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this VpnTunnel.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("localTrafficSelectors")]
        private InputList<string>? _localTrafficSelectors;

        /// <summary>
        /// Local traffic selector to use when establishing the VPN tunnel with
        /// peer VPN gateway. The value should be a CIDR formatted string,
        /// for example `192.168.0.0/16`. The ranges should be disjoint.
        /// Only IPv4 is supported.
        /// </summary>
        public InputList<string> LocalTrafficSelectors
        {
            get => _localTrafficSelectors ?? (_localTrafficSelectors = new InputList<string>());
            set => _localTrafficSelectors = value;
        }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63
        /// characters long and match the regular expression
        /// `a-z?` which means the first character
        /// must be a lowercase letter, and all following characters must
        /// be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
        /// </summary>
        [Input("peerExternalGateway")]
        public Input<string>? PeerExternalGateway { get; set; }

        /// <summary>
        /// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
        /// </summary>
        [Input("peerExternalGatewayInterface")]
        public Input<int>? PeerExternalGatewayInterface { get; set; }

        /// <summary>
        /// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
        /// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
        /// ID in the peer GCP VPN gateway.
        /// This field must reference a `gcp.compute.HaVpnGateway` resource.
        /// </summary>
        [Input("peerGcpGateway")]
        public Input<string>? PeerGcpGateway { get; set; }

        /// <summary>
        /// IP address of the peer VPN gateway. Only IPv4 is supported.
        /// </summary>
        [Input("peerIp")]
        public Input<string>? PeerIp { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("remoteTrafficSelectors")]
        private InputList<string>? _remoteTrafficSelectors;

        /// <summary>
        /// Remote traffic selector to use when establishing the VPN tunnel with
        /// peer VPN gateway. The value should be a CIDR formatted string,
        /// for example `192.168.0.0/16`. The ranges should be disjoint.
        /// Only IPv4 is supported.
        /// </summary>
        public InputList<string> RemoteTrafficSelectors
        {
            get => _remoteTrafficSelectors ?? (_remoteTrafficSelectors = new InputList<string>());
            set => _remoteTrafficSelectors = value;
        }

        /// <summary>
        /// URL of router resource to be used for dynamic routing.
        /// </summary>
        [Input("router")]
        public Input<string>? Router { get; set; }

        /// <summary>
        /// Shared secret used to set the secure session between the Cloud VPN
        /// gateway and the peer VPN gateway.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("sharedSecret", required: true)]
        public Input<string> SharedSecret { get; set; } = null!;

        /// <summary>
        /// URL of the Target VPN gateway with which this VPN tunnel is
        /// associated.
        /// </summary>
        [Input("targetVpnGateway")]
        public Input<string>? TargetVpnGateway { get; set; }

        /// <summary>
        /// URL of the VPN gateway with which this VPN tunnel is associated.
        /// This must be used if a High Availability VPN gateway resource is created.
        /// This field must reference a `gcp.compute.HaVpnGateway` resource.
        /// </summary>
        [Input("vpnGateway")]
        public Input<string>? VpnGateway { get; set; }

        /// <summary>
        /// The interface ID of the VPN gateway with which this VPN tunnel is associated.
        /// </summary>
        [Input("vpnGatewayInterface")]
        public Input<int>? VpnGatewayInterface { get; set; }

        public VPNTunnelArgs()
        {
        }
    }

    public sealed class VPNTunnelState : Pulumi.ResourceArgs
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
        /// Detailed status message for the VPN tunnel.
        /// </summary>
        [Input("detailedStatus")]
        public Input<string>? DetailedStatus { get; set; }

        /// <summary>
        /// IKE protocol version to use when establishing the VPN tunnel with
        /// peer VPN gateway.
        /// Acceptable IKE versions are 1 or 2. Default version is 2.
        /// </summary>
        [Input("ikeVersion")]
        public Input<int>? IkeVersion { get; set; }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this VpnTunnel.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("localTrafficSelectors")]
        private InputList<string>? _localTrafficSelectors;

        /// <summary>
        /// Local traffic selector to use when establishing the VPN tunnel with
        /// peer VPN gateway. The value should be a CIDR formatted string,
        /// for example `192.168.0.0/16`. The ranges should be disjoint.
        /// Only IPv4 is supported.
        /// </summary>
        public InputList<string> LocalTrafficSelectors
        {
            get => _localTrafficSelectors ?? (_localTrafficSelectors = new InputList<string>());
            set => _localTrafficSelectors = value;
        }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63
        /// characters long and match the regular expression
        /// `a-z?` which means the first character
        /// must be a lowercase letter, and all following characters must
        /// be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
        /// </summary>
        [Input("peerExternalGateway")]
        public Input<string>? PeerExternalGateway { get; set; }

        /// <summary>
        /// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
        /// </summary>
        [Input("peerExternalGatewayInterface")]
        public Input<int>? PeerExternalGatewayInterface { get; set; }

        /// <summary>
        /// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
        /// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
        /// ID in the peer GCP VPN gateway.
        /// This field must reference a `gcp.compute.HaVpnGateway` resource.
        /// </summary>
        [Input("peerGcpGateway")]
        public Input<string>? PeerGcpGateway { get; set; }

        /// <summary>
        /// IP address of the peer VPN gateway. Only IPv4 is supported.
        /// </summary>
        [Input("peerIp")]
        public Input<string>? PeerIp { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("remoteTrafficSelectors")]
        private InputList<string>? _remoteTrafficSelectors;

        /// <summary>
        /// Remote traffic selector to use when establishing the VPN tunnel with
        /// peer VPN gateway. The value should be a CIDR formatted string,
        /// for example `192.168.0.0/16`. The ranges should be disjoint.
        /// Only IPv4 is supported.
        /// </summary>
        public InputList<string> RemoteTrafficSelectors
        {
            get => _remoteTrafficSelectors ?? (_remoteTrafficSelectors = new InputList<string>());
            set => _remoteTrafficSelectors = value;
        }

        /// <summary>
        /// URL of router resource to be used for dynamic routing.
        /// </summary>
        [Input("router")]
        public Input<string>? Router { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Shared secret used to set the secure session between the Cloud VPN
        /// gateway and the peer VPN gateway.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("sharedSecret")]
        public Input<string>? SharedSecret { get; set; }

        /// <summary>
        /// Hash of the shared secret.
        /// </summary>
        [Input("sharedSecretHash")]
        public Input<string>? SharedSecretHash { get; set; }

        /// <summary>
        /// URL of the Target VPN gateway with which this VPN tunnel is
        /// associated.
        /// </summary>
        [Input("targetVpnGateway")]
        public Input<string>? TargetVpnGateway { get; set; }

        /// <summary>
        /// The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("tunnelId")]
        public Input<string>? TunnelId { get; set; }

        /// <summary>
        /// URL of the VPN gateway with which this VPN tunnel is associated.
        /// This must be used if a High Availability VPN gateway resource is created.
        /// This field must reference a `gcp.compute.HaVpnGateway` resource.
        /// </summary>
        [Input("vpnGateway")]
        public Input<string>? VpnGateway { get; set; }

        /// <summary>
        /// The interface ID of the VPN gateway with which this VPN tunnel is associated.
        /// </summary>
        [Input("vpnGatewayInterface")]
        public Input<int>? VpnGatewayInterface { get; set; }

        public VPNTunnelState()
        {
        }
    }
}
