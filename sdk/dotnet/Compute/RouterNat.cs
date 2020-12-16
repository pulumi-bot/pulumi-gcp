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
    /// A NAT service created in a router.
    /// 
    /// To get more information about RouterNat, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
    /// * How-to Guides
    ///     * [Google Cloud Router](https://cloud.google.com/router/docs/)
    /// 
    /// ## Example Usage
    /// ### Router Nat Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var net = new Gcp.Compute.Network("net", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var subnet = new Gcp.Compute.Subnetwork("subnet", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             Network = net.Id,
    ///             IpCidrRange = "10.0.0.0/16",
    ///             Region = "us-central1",
    ///         });
    ///         var router = new Gcp.Compute.Router("router", new Gcp.Compute.RouterArgs
    ///         {
    ///             Region = subnet.Region,
    ///             Network = net.Id,
    ///             Bgp = new Gcp.Compute.Inputs.RouterBgpArgs
    ///             {
    ///                 Asn = 64514,
    ///             },
    ///         });
    ///         var nat = new Gcp.Compute.RouterNat("nat", new Gcp.Compute.RouterNatArgs
    ///         {
    ///             Router = router.Name,
    ///             Region = router.Region,
    ///             NatIpAllocateOption = "AUTO_ONLY",
    ///             SourceSubnetworkIpRangesToNat = "ALL_SUBNETWORKS_ALL_IP_RANGES",
    ///             LogConfig = new Gcp.Compute.Inputs.RouterNatLogConfigArgs
    ///             {
    ///                 Enable = true,
    ///                 Filter = "ERRORS_ONLY",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Router Nat Manual Ips
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var net = new Gcp.Compute.Network("net", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var subnet = new Gcp.Compute.Subnetwork("subnet", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             Network = net.Id,
    ///             IpCidrRange = "10.0.0.0/16",
    ///             Region = "us-central1",
    ///         });
    ///         var router = new Gcp.Compute.Router("router", new Gcp.Compute.RouterArgs
    ///         {
    ///             Region = subnet.Region,
    ///             Network = net.Id,
    ///         });
    ///         var address = new List&lt;Gcp.Compute.Address&gt;();
    ///         for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///         {
    ///             var range = new { Value = rangeIndex };
    ///             address.Add(new Gcp.Compute.Address($"address-{range.Value}", new Gcp.Compute.AddressArgs
    ///             {
    ///                 Region = subnet.Region,
    ///             }));
    ///         }
    ///         var natManual = new Gcp.Compute.RouterNat("natManual", new Gcp.Compute.RouterNatArgs
    ///         {
    ///             Router = router.Name,
    ///             Region = router.Region,
    ///             NatIpAllocateOption = "MANUAL_ONLY",
    ///             NatIps = address.Select(__item =&gt; __item.SelfLink).ToList(),
    ///             SourceSubnetworkIpRangesToNat = "LIST_OF_SUBNETWORKS",
    ///             Subnetworks = 
    ///             {
    ///                 new Gcp.Compute.Inputs.RouterNatSubnetworkArgs
    ///                 {
    ///                     Name = subnet.Id,
    ///                     SourceIpRangesToNats = 
    ///                     {
    ///                         "ALL_IP_RANGES",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// RouterNat can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/routerNat:RouterNat default projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/routerNat:RouterNat default {{project}}/{{region}}/{{router}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/routerNat:RouterNat default {{region}}/{{router}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/routerNat:RouterNat default {{router}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/routerNat:RouterNat")]
    public partial class RouterNat : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of URLs of the IP resources to be drained. These IPs must be
        /// valid static external IPs that have been assigned to the NAT.
        /// </summary>
        [Output("drainNatIps")]
        public Output<ImmutableArray<string>> DrainNatIps { get; private set; } = null!;

        /// <summary>
        /// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
        /// </summary>
        [Output("icmpIdleTimeoutSec")]
        public Output<int?> IcmpIdleTimeoutSec { get; private set; } = null!;

        /// <summary>
        /// Configuration for logging on NAT
        /// Structure is documented below.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.RouterNatLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// Minimum number of ports allocated to a VM from this NAT.
        /// </summary>
        [Output("minPortsPerVm")]
        public Output<int?> MinPortsPerVm { get; private set; } = null!;

        /// <summary>
        /// Self-link of subnetwork to NAT
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// How external IPs should be allocated for this NAT. Valid values are
        /// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
        /// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
        /// Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
        /// </summary>
        [Output("natIpAllocateOption")]
        public Output<string> NatIpAllocateOption { get; private set; } = null!;

        /// <summary>
        /// Self-links of NAT IPs. Only valid if natIpAllocateOption
        /// is set to MANUAL_ONLY.
        /// </summary>
        [Output("natIps")]
        public Output<ImmutableArray<string>> NatIps { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Region where the router and NAT reside.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The name of the Cloud Router in which this NAT will be configured.
        /// </summary>
        [Output("router")]
        public Output<string> Router { get; private set; } = null!;

        /// <summary>
        /// How NAT should be configured per Subnetwork.
        /// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
        /// IP ranges in every Subnetwork are allowed to Nat.
        /// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
        /// ranges in every Subnetwork are allowed to Nat.
        /// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
        /// (specified in the field subnetwork below). Note that if this field
        /// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
        /// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
        /// other RouterNat section in any Router for this network in this region.
        /// Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
        /// </summary>
        [Output("sourceSubnetworkIpRangesToNat")]
        public Output<string> SourceSubnetworkIpRangesToNat { get; private set; } = null!;

        /// <summary>
        /// One or more subnetwork NAT configurations. Only used if
        /// `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
        /// Structure is documented below.
        /// </summary>
        [Output("subnetworks")]
        public Output<ImmutableArray<Outputs.RouterNatSubnetwork>> Subnetworks { get; private set; } = null!;

        /// <summary>
        /// Timeout (in seconds) for TCP established connections.
        /// Defaults to 1200s if not set.
        /// </summary>
        [Output("tcpEstablishedIdleTimeoutSec")]
        public Output<int?> TcpEstablishedIdleTimeoutSec { get; private set; } = null!;

        /// <summary>
        /// Timeout (in seconds) for TCP transitory connections.
        /// Defaults to 30s if not set.
        /// </summary>
        [Output("tcpTransitoryIdleTimeoutSec")]
        public Output<int?> TcpTransitoryIdleTimeoutSec { get; private set; } = null!;

        /// <summary>
        /// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
        /// </summary>
        [Output("udpIdleTimeoutSec")]
        public Output<int?> UdpIdleTimeoutSec { get; private set; } = null!;


        /// <summary>
        /// Create a RouterNat resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouterNat(string name, RouterNatArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/routerNat:RouterNat", name, args ?? new RouterNatArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouterNat(string name, Input<string> id, RouterNatState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/routerNat:RouterNat", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouterNat resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouterNat Get(string name, Input<string> id, RouterNatState? state = null, CustomResourceOptions? options = null)
        {
            return new RouterNat(name, id, state, options);
        }
    }

    public sealed class RouterNatArgs : Pulumi.ResourceArgs
    {
        [Input("drainNatIps")]
        private InputList<string>? _drainNatIps;

        /// <summary>
        /// A list of URLs of the IP resources to be drained. These IPs must be
        /// valid static external IPs that have been assigned to the NAT.
        /// </summary>
        public InputList<string> DrainNatIps
        {
            get => _drainNatIps ?? (_drainNatIps = new InputList<string>());
            set => _drainNatIps = value;
        }

        /// <summary>
        /// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
        /// </summary>
        [Input("icmpIdleTimeoutSec")]
        public Input<int>? IcmpIdleTimeoutSec { get; set; }

        /// <summary>
        /// Configuration for logging on NAT
        /// Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.RouterNatLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// Minimum number of ports allocated to a VM from this NAT.
        /// </summary>
        [Input("minPortsPerVm")]
        public Input<int>? MinPortsPerVm { get; set; }

        /// <summary>
        /// Self-link of subnetwork to NAT
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How external IPs should be allocated for this NAT. Valid values are
        /// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
        /// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
        /// Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
        /// </summary>
        [Input("natIpAllocateOption", required: true)]
        public Input<string> NatIpAllocateOption { get; set; } = null!;

        [Input("natIps")]
        private InputList<string>? _natIps;

        /// <summary>
        /// Self-links of NAT IPs. Only valid if natIpAllocateOption
        /// is set to MANUAL_ONLY.
        /// </summary>
        public InputList<string> NatIps
        {
            get => _natIps ?? (_natIps = new InputList<string>());
            set => _natIps = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the router and NAT reside.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of the Cloud Router in which this NAT will be configured.
        /// </summary>
        [Input("router", required: true)]
        public Input<string> Router { get; set; } = null!;

        /// <summary>
        /// How NAT should be configured per Subnetwork.
        /// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
        /// IP ranges in every Subnetwork are allowed to Nat.
        /// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
        /// ranges in every Subnetwork are allowed to Nat.
        /// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
        /// (specified in the field subnetwork below). Note that if this field
        /// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
        /// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
        /// other RouterNat section in any Router for this network in this region.
        /// Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
        /// </summary>
        [Input("sourceSubnetworkIpRangesToNat", required: true)]
        public Input<string> SourceSubnetworkIpRangesToNat { get; set; } = null!;

        [Input("subnetworks")]
        private InputList<Inputs.RouterNatSubnetworkArgs>? _subnetworks;

        /// <summary>
        /// One or more subnetwork NAT configurations. Only used if
        /// `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RouterNatSubnetworkArgs> Subnetworks
        {
            get => _subnetworks ?? (_subnetworks = new InputList<Inputs.RouterNatSubnetworkArgs>());
            set => _subnetworks = value;
        }

        /// <summary>
        /// Timeout (in seconds) for TCP established connections.
        /// Defaults to 1200s if not set.
        /// </summary>
        [Input("tcpEstablishedIdleTimeoutSec")]
        public Input<int>? TcpEstablishedIdleTimeoutSec { get; set; }

        /// <summary>
        /// Timeout (in seconds) for TCP transitory connections.
        /// Defaults to 30s if not set.
        /// </summary>
        [Input("tcpTransitoryIdleTimeoutSec")]
        public Input<int>? TcpTransitoryIdleTimeoutSec { get; set; }

        /// <summary>
        /// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
        /// </summary>
        [Input("udpIdleTimeoutSec")]
        public Input<int>? UdpIdleTimeoutSec { get; set; }

        public RouterNatArgs()
        {
        }
    }

    public sealed class RouterNatState : Pulumi.ResourceArgs
    {
        [Input("drainNatIps")]
        private InputList<string>? _drainNatIps;

        /// <summary>
        /// A list of URLs of the IP resources to be drained. These IPs must be
        /// valid static external IPs that have been assigned to the NAT.
        /// </summary>
        public InputList<string> DrainNatIps
        {
            get => _drainNatIps ?? (_drainNatIps = new InputList<string>());
            set => _drainNatIps = value;
        }

        /// <summary>
        /// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
        /// </summary>
        [Input("icmpIdleTimeoutSec")]
        public Input<int>? IcmpIdleTimeoutSec { get; set; }

        /// <summary>
        /// Configuration for logging on NAT
        /// Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.RouterNatLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// Minimum number of ports allocated to a VM from this NAT.
        /// </summary>
        [Input("minPortsPerVm")]
        public Input<int>? MinPortsPerVm { get; set; }

        /// <summary>
        /// Self-link of subnetwork to NAT
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How external IPs should be allocated for this NAT. Valid values are
        /// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
        /// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
        /// Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
        /// </summary>
        [Input("natIpAllocateOption")]
        public Input<string>? NatIpAllocateOption { get; set; }

        [Input("natIps")]
        private InputList<string>? _natIps;

        /// <summary>
        /// Self-links of NAT IPs. Only valid if natIpAllocateOption
        /// is set to MANUAL_ONLY.
        /// </summary>
        public InputList<string> NatIps
        {
            get => _natIps ?? (_natIps = new InputList<string>());
            set => _natIps = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the router and NAT reside.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of the Cloud Router in which this NAT will be configured.
        /// </summary>
        [Input("router")]
        public Input<string>? Router { get; set; }

        /// <summary>
        /// How NAT should be configured per Subnetwork.
        /// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
        /// IP ranges in every Subnetwork are allowed to Nat.
        /// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
        /// ranges in every Subnetwork are allowed to Nat.
        /// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
        /// (specified in the field subnetwork below). Note that if this field
        /// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
        /// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
        /// other RouterNat section in any Router for this network in this region.
        /// Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
        /// </summary>
        [Input("sourceSubnetworkIpRangesToNat")]
        public Input<string>? SourceSubnetworkIpRangesToNat { get; set; }

        [Input("subnetworks")]
        private InputList<Inputs.RouterNatSubnetworkGetArgs>? _subnetworks;

        /// <summary>
        /// One or more subnetwork NAT configurations. Only used if
        /// `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RouterNatSubnetworkGetArgs> Subnetworks
        {
            get => _subnetworks ?? (_subnetworks = new InputList<Inputs.RouterNatSubnetworkGetArgs>());
            set => _subnetworks = value;
        }

        /// <summary>
        /// Timeout (in seconds) for TCP established connections.
        /// Defaults to 1200s if not set.
        /// </summary>
        [Input("tcpEstablishedIdleTimeoutSec")]
        public Input<int>? TcpEstablishedIdleTimeoutSec { get; set; }

        /// <summary>
        /// Timeout (in seconds) for TCP transitory connections.
        /// Defaults to 30s if not set.
        /// </summary>
        [Input("tcpTransitoryIdleTimeoutSec")]
        public Input<int>? TcpTransitoryIdleTimeoutSec { get; set; }

        /// <summary>
        /// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
        /// </summary>
        [Input("udpIdleTimeoutSec")]
        public Input<int>? UdpIdleTimeoutSec { get; set; }

        public RouterNatState()
        {
        }
    }
}
