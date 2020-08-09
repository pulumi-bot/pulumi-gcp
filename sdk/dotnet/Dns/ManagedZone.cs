// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns
{
    /// <summary>
    /// A zone is a subtree of the DNS namespace under one administrative
    /// responsibility. A ManagedZone is a resource that represents a DNS zone
    /// hosted by the Cloud DNS service.
    /// 
    /// To get more information about ManagedZone, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
    /// * How-to Guides
    ///     * [Managing Zones](https://cloud.google.com/dns/zones/)
    /// 
    /// ## Example Usage
    /// ### Dns Managed Zone Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example_zone = new Gcp.Dns.ManagedZone("example-zone", new Gcp.Dns.ManagedZoneArgs
    ///         {
    ///             Description = "Example DNS zone",
    ///             DnsName = "my-domain.com.",
    ///             Labels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Dns Managed Zone Private
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network_1 = new Gcp.Compute.Network("network-1", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var network_2 = new Gcp.Compute.Network("network-2", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var private_zone = new Gcp.Dns.ManagedZone("private-zone", new Gcp.Dns.ManagedZoneArgs
    ///         {
    ///             DnsName = "private.example.com.",
    ///             Description = "Example private DNS zone",
    ///             Labels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///             Visibility = "private",
    ///             PrivateVisibilityConfig = new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigArgs
    ///             {
    ///                 Networks = 
    ///                 {
    ///                     new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigNetworkArgs
    ///                     {
    ///                         NetworkUrl = network_1.Id,
    ///                     },
    ///                     new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigNetworkArgs
    ///                     {
    ///                         NetworkUrl = network_2.Id,
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Dns Managed Zone Private Forwarding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network_1 = new Gcp.Compute.Network("network-1", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var network_2 = new Gcp.Compute.Network("network-2", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var private_zone = new Gcp.Dns.ManagedZone("private-zone", new Gcp.Dns.ManagedZoneArgs
    ///         {
    ///             DnsName = "private.example.com.",
    ///             Description = "Example private DNS zone",
    ///             Labels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///             Visibility = "private",
    ///             PrivateVisibilityConfig = new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigArgs
    ///             {
    ///                 Networks = 
    ///                 {
    ///                     new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigNetworkArgs
    ///                     {
    ///                         NetworkUrl = network_1.Id,
    ///                     },
    ///                     new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigNetworkArgs
    ///                     {
    ///                         NetworkUrl = network_2.Id,
    ///                     },
    ///                 },
    ///             },
    ///             ForwardingConfig = new Gcp.Dns.Inputs.ManagedZoneForwardingConfigArgs
    ///             {
    ///                 TargetNameServers = 
    ///                 {
    ///                     new Gcp.Dns.Inputs.ManagedZoneForwardingConfigTargetNameServerArgs
    ///                     {
    ///                         Ipv4Address = "172.16.1.10",
    ///                     },
    ///                     new Gcp.Dns.Inputs.ManagedZoneForwardingConfigTargetNameServerArgs
    ///                     {
    ///                         Ipv4Address = "172.16.1.20",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Dns Managed Zone Private Peering
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network_source = new Gcp.Compute.Network("network-source", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var network_target = new Gcp.Compute.Network("network-target", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var peering_zone = new Gcp.Dns.ManagedZone("peering-zone", new Gcp.Dns.ManagedZoneArgs
    ///         {
    ///             DnsName = "peering.example.com.",
    ///             Description = "Example private DNS peering zone",
    ///             Visibility = "private",
    ///             PrivateVisibilityConfig = new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigArgs
    ///             {
    ///                 Networks = 
    ///                 {
    ///                     new Gcp.Dns.Inputs.ManagedZonePrivateVisibilityConfigNetworkArgs
    ///                     {
    ///                         NetworkUrl = network_source.Id,
    ///                     },
    ///                 },
    ///             },
    ///             PeeringConfig = new Gcp.Dns.Inputs.ManagedZonePeeringConfigArgs
    ///             {
    ///                 TargetNetwork = new Gcp.Dns.Inputs.ManagedZonePeeringConfigTargetNetworkArgs
    ///                 {
    ///                     NetworkUrl = network_target.Id,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Dns Managed Zone Service Directory
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Gcp.ServiceDirectory.Namespace("example", new Gcp.ServiceDirectory.NamespaceArgs
    ///         {
    ///             NamespaceId = "example",
    ///             Location = "us-central1",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var sd_zone = new Gcp.Dns.ManagedZone("sd-zone", new Gcp.Dns.ManagedZoneArgs
    ///         {
    ///             DnsName = "services.example.com.",
    ///             Description = "Example private DNS Service Directory zone",
    ///             Visibility = "private",
    ///             ServiceDirectoryConfig = new Gcp.Dns.Inputs.ManagedZoneServiceDirectoryConfigArgs
    ///             {
    ///                 Namespace = new Gcp.Dns.Inputs.ManagedZoneServiceDirectoryConfigNamespaceArgs
    ///                 {
    ///                     NamespaceUrl = example.Id,
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var network = new Gcp.Compute.Network("network", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ManagedZone : Pulumi.CustomResource
    {
        /// <summary>
        /// A textual description field. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// DNSSEC configuration  Structure is documented below.
        /// </summary>
        [Output("dnssecConfig")]
        public Output<Outputs.ManagedZoneDnssecConfig?> DnssecConfig { get; private set; } = null!;

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled
        /// for this zone. The value of this field contains the set of destinations
        /// to forward to.  Structure is documented below.
        /// </summary>
        [Output("forwardingConfig")]
        public Output<Outputs.ManagedZoneForwardingConfig?> ForwardingConfig { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to this ManagedZone.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// User assigned name for this resource.
        /// Must be unique within the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Delegate your managed_zone to these virtual name servers; defined by the server
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// The presence of this field indicates that DNS Peering is enabled for this
        /// zone. The value of this field contains the network to peer with.  Structure is documented below.
        /// </summary>
        [Output("peeringConfig")]
        public Output<Outputs.ManagedZonePeeringConfig?> PeeringConfig { get; private set; } = null!;

        /// <summary>
        /// For privately visible zones, the set of Virtual Private Cloud
        /// resources that the zone is visible from.  Structure is documented below.
        /// </summary>
        [Output("privateVisibilityConfig")]
        public Output<Outputs.ManagedZonePrivateVisibilityConfig?> PrivateVisibilityConfig { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
        /// lookup queries using automatically configured records for VPC resources. This only applies
        /// to networks listed under `private_visibility_config`.
        /// </summary>
        [Output("reverseLookup")]
        public Output<bool?> ReverseLookup { get; private set; } = null!;

        /// <summary>
        /// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.  Structure is documented below.
        /// </summary>
        [Output("serviceDirectoryConfig")]
        public Output<Outputs.ManagedZoneServiceDirectoryConfig?> ServiceDirectoryConfig { get; private set; } = null!;

        /// <summary>
        /// The zone's visibility: public zones are exposed to the Internet,
        /// while private zones are visible only to Virtual Private Cloud resources.
        /// </summary>
        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedZone(string name, ManagedZoneArgs args, CustomResourceOptions? options = null)
            : base("gcp:dns/managedZone:ManagedZone", name, args ?? new ManagedZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedZone(string name, Input<string> id, ManagedZoneState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dns/managedZone:ManagedZone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManagedZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedZone Get(string name, Input<string> id, ManagedZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedZone(name, id, state, options);
        }
    }

    public sealed class ManagedZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A textual description field. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Input("dnsName", required: true)]
        public Input<string> DnsName { get; set; } = null!;

        /// <summary>
        /// DNSSEC configuration  Structure is documented below.
        /// </summary>
        [Input("dnssecConfig")]
        public Input<Inputs.ManagedZoneDnssecConfigArgs>? DnssecConfig { get; set; }

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled
        /// for this zone. The value of this field contains the set of destinations
        /// to forward to.  Structure is documented below.
        /// </summary>
        [Input("forwardingConfig")]
        public Input<Inputs.ManagedZoneForwardingConfigArgs>? ForwardingConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to this ManagedZone.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// User assigned name for this resource.
        /// Must be unique within the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The presence of this field indicates that DNS Peering is enabled for this
        /// zone. The value of this field contains the network to peer with.  Structure is documented below.
        /// </summary>
        [Input("peeringConfig")]
        public Input<Inputs.ManagedZonePeeringConfigArgs>? PeeringConfig { get; set; }

        /// <summary>
        /// For privately visible zones, the set of Virtual Private Cloud
        /// resources that the zone is visible from.  Structure is documented below.
        /// </summary>
        [Input("privateVisibilityConfig")]
        public Input<Inputs.ManagedZonePrivateVisibilityConfigArgs>? PrivateVisibilityConfig { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
        /// lookup queries using automatically configured records for VPC resources. This only applies
        /// to networks listed under `private_visibility_config`.
        /// </summary>
        [Input("reverseLookup")]
        public Input<bool>? ReverseLookup { get; set; }

        /// <summary>
        /// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.  Structure is documented below.
        /// </summary>
        [Input("serviceDirectoryConfig")]
        public Input<Inputs.ManagedZoneServiceDirectoryConfigArgs>? ServiceDirectoryConfig { get; set; }

        /// <summary>
        /// The zone's visibility: public zones are exposed to the Internet,
        /// while private zones are visible only to Virtual Private Cloud resources.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ManagedZoneArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class ManagedZoneState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A textual description field. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// DNSSEC configuration  Structure is documented below.
        /// </summary>
        [Input("dnssecConfig")]
        public Input<Inputs.ManagedZoneDnssecConfigGetArgs>? DnssecConfig { get; set; }

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled
        /// for this zone. The value of this field contains the set of destinations
        /// to forward to.  Structure is documented below.
        /// </summary>
        [Input("forwardingConfig")]
        public Input<Inputs.ManagedZoneForwardingConfigGetArgs>? ForwardingConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to this ManagedZone.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// User assigned name for this resource.
        /// Must be unique within the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameServers")]
        private InputList<string>? _nameServers;

        /// <summary>
        /// Delegate your managed_zone to these virtual name servers; defined by the server
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        /// <summary>
        /// The presence of this field indicates that DNS Peering is enabled for this
        /// zone. The value of this field contains the network to peer with.  Structure is documented below.
        /// </summary>
        [Input("peeringConfig")]
        public Input<Inputs.ManagedZonePeeringConfigGetArgs>? PeeringConfig { get; set; }

        /// <summary>
        /// For privately visible zones, the set of Virtual Private Cloud
        /// resources that the zone is visible from.  Structure is documented below.
        /// </summary>
        [Input("privateVisibilityConfig")]
        public Input<Inputs.ManagedZonePrivateVisibilityConfigGetArgs>? PrivateVisibilityConfig { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
        /// lookup queries using automatically configured records for VPC resources. This only applies
        /// to networks listed under `private_visibility_config`.
        /// </summary>
        [Input("reverseLookup")]
        public Input<bool>? ReverseLookup { get; set; }

        /// <summary>
        /// The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.  Structure is documented below.
        /// </summary>
        [Input("serviceDirectoryConfig")]
        public Input<Inputs.ManagedZoneServiceDirectoryConfigGetArgs>? ServiceDirectoryConfig { get; set; }

        /// <summary>
        /// The zone's visibility: public zones are exposed to the Internet,
        /// while private zones are visible only to Virtual Private Cloud resources.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ManagedZoneState()
        {
            Description = "Managed by Pulumi";
        }
    }
}
