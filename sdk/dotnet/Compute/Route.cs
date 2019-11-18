// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_route.html.markdown.
    /// </summary>
    public partial class Route : Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination range of outgoing packets that this route applies to. Only IPv4 is supported.
        /// </summary>
        [Output("destRange")]
        public Output<string> DestRange { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network that this route applies to.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// URL to a gateway that should handle matching packets. Currently, you can only specify the internet gateway,
        /// using a full or partial valid URL: *
        /// 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway' *
        /// 'projects/project/global/gateways/default-internet-gateway' * 'global/gateways/default-internet-gateway' *
        /// The string 'default-internet-gateway'.
        /// </summary>
        [Output("nextHopGateway")]
        public Output<string?> NextHopGateway { get; private set; } = null!;

        /// <summary>
        /// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets. You
        /// can only specify the forwarding rule as a partial or full URL. For example, the following are all valid
        /// URLs: https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
        /// regions/region/forwardingRules/forwardingRule Note that this can only be used when the destinationRange is a
        /// public (non-RFC 1918) IP CIDR range.
        /// </summary>
        [Output("nextHopIlb")]
        public Output<string?> NextHopIlb { get; private set; } = null!;

        /// <summary>
        /// URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For
        /// example: * 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance' *
        /// 'projects/project/zones/zone/instances/instance' * 'zones/zone/instances/instance' * Just the instance name,
        /// with the zone in 'next_hop_instance_zone'.
        /// </summary>
        [Output("nextHopInstance")]
        public Output<string?> NextHopInstance { get; private set; } = null!;

        /// <summary>
        /// (Optional when `next_hop_instance` is
        /// specified)  The zone of the instance specified in
        /// `next_hop_instance`.  Omit if `next_hop_instance` is specified as
        /// a URL.
        /// </summary>
        [Output("nextHopInstanceZone")]
        public Output<string?> NextHopInstanceZone { get; private set; } = null!;

        /// <summary>
        /// Network IP address of an instance that should handle matching packets.
        /// </summary>
        [Output("nextHopIp")]
        public Output<string> NextHopIp { get; private set; } = null!;

        /// <summary>
        /// URL to a Network that should handle matching packets.
        /// </summary>
        [Output("nextHopNetwork")]
        public Output<string> NextHopNetwork { get; private set; } = null!;

        /// <summary>
        /// URL to a VpnTunnel that should handle matching packets.
        /// </summary>
        [Output("nextHopVpnTunnel")]
        public Output<string?> NextHopVpnTunnel { get; private set; } = null!;

        /// <summary>
        /// The priority of this route. Priority is used to break ties in cases where there is more than one matching
        /// route of equal prefix length. In the case of two routes with equal prefix length, the one with the
        /// lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// A list of instance tags to which this route applies.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/route:Route", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/route:Route", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
        {
            return new Route(name, id, state, options);
        }
    }

    public sealed class RouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination range of outgoing packets that this route applies to. Only IPv4 is supported.
        /// </summary>
        [Input("destRange", required: true)]
        public Input<string> DestRange { get; set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network that this route applies to.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// URL to a gateway that should handle matching packets. Currently, you can only specify the internet gateway,
        /// using a full or partial valid URL: *
        /// 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway' *
        /// 'projects/project/global/gateways/default-internet-gateway' * 'global/gateways/default-internet-gateway' *
        /// The string 'default-internet-gateway'.
        /// </summary>
        [Input("nextHopGateway")]
        public Input<string>? NextHopGateway { get; set; }

        /// <summary>
        /// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets. You
        /// can only specify the forwarding rule as a partial or full URL. For example, the following are all valid
        /// URLs: https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
        /// regions/region/forwardingRules/forwardingRule Note that this can only be used when the destinationRange is a
        /// public (non-RFC 1918) IP CIDR range.
        /// </summary>
        [Input("nextHopIlb")]
        public Input<string>? NextHopIlb { get; set; }

        /// <summary>
        /// URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For
        /// example: * 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance' *
        /// 'projects/project/zones/zone/instances/instance' * 'zones/zone/instances/instance' * Just the instance name,
        /// with the zone in 'next_hop_instance_zone'.
        /// </summary>
        [Input("nextHopInstance")]
        public Input<string>? NextHopInstance { get; set; }

        /// <summary>
        /// (Optional when `next_hop_instance` is
        /// specified)  The zone of the instance specified in
        /// `next_hop_instance`.  Omit if `next_hop_instance` is specified as
        /// a URL.
        /// </summary>
        [Input("nextHopInstanceZone")]
        public Input<string>? NextHopInstanceZone { get; set; }

        /// <summary>
        /// Network IP address of an instance that should handle matching packets.
        /// </summary>
        [Input("nextHopIp")]
        public Input<string>? NextHopIp { get; set; }

        /// <summary>
        /// URL to a VpnTunnel that should handle matching packets.
        /// </summary>
        [Input("nextHopVpnTunnel")]
        public Input<string>? NextHopVpnTunnel { get; set; }

        /// <summary>
        /// The priority of this route. Priority is used to break ties in cases where there is more than one matching
        /// route of equal prefix length. In the case of two routes with equal prefix length, the one with the
        /// lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of instance tags to which this route applies.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public RouteArgs()
        {
        }
    }

    public sealed class RouteState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination range of outgoing packets that this route applies to. Only IPv4 is supported.
        /// </summary>
        [Input("destRange")]
        public Input<string>? DestRange { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network that this route applies to.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// URL to a gateway that should handle matching packets. Currently, you can only specify the internet gateway,
        /// using a full or partial valid URL: *
        /// 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway' *
        /// 'projects/project/global/gateways/default-internet-gateway' * 'global/gateways/default-internet-gateway' *
        /// The string 'default-internet-gateway'.
        /// </summary>
        [Input("nextHopGateway")]
        public Input<string>? NextHopGateway { get; set; }

        /// <summary>
        /// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets. You
        /// can only specify the forwarding rule as a partial or full URL. For example, the following are all valid
        /// URLs: https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
        /// regions/region/forwardingRules/forwardingRule Note that this can only be used when the destinationRange is a
        /// public (non-RFC 1918) IP CIDR range.
        /// </summary>
        [Input("nextHopIlb")]
        public Input<string>? NextHopIlb { get; set; }

        /// <summary>
        /// URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For
        /// example: * 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance' *
        /// 'projects/project/zones/zone/instances/instance' * 'zones/zone/instances/instance' * Just the instance name,
        /// with the zone in 'next_hop_instance_zone'.
        /// </summary>
        [Input("nextHopInstance")]
        public Input<string>? NextHopInstance { get; set; }

        /// <summary>
        /// (Optional when `next_hop_instance` is
        /// specified)  The zone of the instance specified in
        /// `next_hop_instance`.  Omit if `next_hop_instance` is specified as
        /// a URL.
        /// </summary>
        [Input("nextHopInstanceZone")]
        public Input<string>? NextHopInstanceZone { get; set; }

        /// <summary>
        /// Network IP address of an instance that should handle matching packets.
        /// </summary>
        [Input("nextHopIp")]
        public Input<string>? NextHopIp { get; set; }

        /// <summary>
        /// URL to a Network that should handle matching packets.
        /// </summary>
        [Input("nextHopNetwork")]
        public Input<string>? NextHopNetwork { get; set; }

        /// <summary>
        /// URL to a VpnTunnel that should handle matching packets.
        /// </summary>
        [Input("nextHopVpnTunnel")]
        public Input<string>? NextHopVpnTunnel { get; set; }

        /// <summary>
        /// The priority of this route. Priority is used to break ties in cases where there is more than one matching
        /// route of equal prefix length. In the case of two routes with equal prefix length, the one with the
        /// lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of instance tags to which this route applies.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public RouteState()
        {
        }
    }
}
