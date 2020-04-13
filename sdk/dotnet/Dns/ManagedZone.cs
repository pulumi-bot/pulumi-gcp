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
    /// 
    /// To get more information about ManagedZone, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
    /// * How-to Guides
    ///     * [Managing Zones](https://cloud.google.com/dns/zones/)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_managed_zone.html.markdown.
    /// </summary>
    public partial class ManagedZone : Pulumi.CustomResource
    {
        /// <summary>
        /// A textual description field. Defaults to 'Managed by Terraform'.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// DNSSEC configuration  Structure is documented below.
        /// </summary>
        [Output("dnssecConfig")]
        public Output<Outputs.ManagedZoneDnssecConfig?> DnssecConfig { get; private set; } = null!;

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled for this zone. The value of this
        /// field contains the set of destinations to forward to.
        /// </summary>
        [Output("forwardingConfig")]
        public Output<Outputs.ManagedZoneForwardingConfig?> ForwardingConfig { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A set of key/value label pairs to assign to this ManagedZone.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
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
        /// The presence of this field indicates that DNS Peering is enabled for this zone. The value of this field
        /// contains the network to peer with.
        /// </summary>
        [Output("peeringConfig")]
        public Output<Outputs.ManagedZonePeeringConfig?> PeeringConfig { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
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
        /// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse lookup queries
        /// using automatically configured records for VPC resources. This only applies to networks listed under
        /// 'private_visibility_config'.
        /// </summary>
        [Output("reverseLookup")]
        public Output<bool?> ReverseLookup { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The zone's visibility: public zones are exposed to the Internet,
        /// while private zones are visible only to Virtual Private Cloud resources.
        /// Must be one of: `public`, `private`.
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
            : base("gcp:dns/managedZone:ManagedZone", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
        /// A textual description field. Defaults to 'Managed by Terraform'.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Input("dnsName", required: true)]
        public Input<string> DnsName { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// DNSSEC configuration  Structure is documented below.
        /// </summary>
        [Input("dnssecConfig")]
        public Input<Inputs.ManagedZoneDnssecConfigArgs>? DnssecConfig { get; set; }

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled for this zone. The value of this
        /// field contains the set of destinations to forward to.
        /// </summary>
        [Input("forwardingConfig")]
        public Input<Inputs.ManagedZoneForwardingConfigArgs>? ForwardingConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// -
        /// (Optional)
        /// A set of key/value label pairs to assign to this ManagedZone.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// User assigned name for this resource.
        /// Must be unique within the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The presence of this field indicates that DNS Peering is enabled for this zone. The value of this field
        /// contains the network to peer with.
        /// </summary>
        [Input("peeringConfig")]
        public Input<Inputs.ManagedZonePeeringConfigArgs>? PeeringConfig { get; set; }

        /// <summary>
        /// -
        /// (Optional)
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
        /// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse lookup queries
        /// using automatically configured records for VPC resources. This only applies to networks listed under
        /// 'private_visibility_config'.
        /// </summary>
        [Input("reverseLookup")]
        public Input<bool>? ReverseLookup { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The zone's visibility: public zones are exposed to the Internet,
        /// while private zones are visible only to Virtual Private Cloud resources.
        /// Must be one of: `public`, `private`.
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
        /// A textual description field. Defaults to 'Managed by Terraform'.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// DNSSEC configuration  Structure is documented below.
        /// </summary>
        [Input("dnssecConfig")]
        public Input<Inputs.ManagedZoneDnssecConfigGetArgs>? DnssecConfig { get; set; }

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled for this zone. The value of this
        /// field contains the set of destinations to forward to.
        /// </summary>
        [Input("forwardingConfig")]
        public Input<Inputs.ManagedZoneForwardingConfigGetArgs>? ForwardingConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// -
        /// (Optional)
        /// A set of key/value label pairs to assign to this ManagedZone.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Required)
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
        /// The presence of this field indicates that DNS Peering is enabled for this zone. The value of this field
        /// contains the network to peer with.
        /// </summary>
        [Input("peeringConfig")]
        public Input<Inputs.ManagedZonePeeringConfigGetArgs>? PeeringConfig { get; set; }

        /// <summary>
        /// -
        /// (Optional)
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
        /// Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse lookup queries
        /// using automatically configured records for VPC resources. This only applies to networks listed under
        /// 'private_visibility_config'.
        /// </summary>
        [Input("reverseLookup")]
        public Input<bool>? ReverseLookup { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The zone's visibility: public zones are exposed to the Internet,
        /// while private zones are visible only to Virtual Private Cloud resources.
        /// Must be one of: `public`, `private`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ManagedZoneState()
        {
            Description = "Managed by Pulumi";
        }
    }

    namespace Inputs
    {

    public sealed class ManagedZoneDnssecConfigArgs : Pulumi.ResourceArgs
    {
        [Input("defaultKeySpecs")]
        private InputList<ManagedZoneDnssecConfigDefaultKeySpecsArgs>? _defaultKeySpecs;

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies parameters that will be used for generating initial DnsKeys
        /// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
        /// you must also provide one for the other.
        /// default_key_specs can only be updated when the state is `off`.  Structure is documented below.
        /// </summary>
        public InputList<ManagedZoneDnssecConfigDefaultKeySpecsArgs> DefaultKeySpecs
        {
            get => _defaultKeySpecs ?? (_defaultKeySpecs = new InputList<ManagedZoneDnssecConfigDefaultKeySpecsArgs>());
            set => _defaultKeySpecs = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Identifies what kind of resource this is
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the mechanism used to provide authenticated denial-of-existence responses.
        /// non_existence can only be updated when the state is `off`.
        /// </summary>
        [Input("nonExistence")]
        public Input<string>? NonExistence { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies whether DNSSEC is enabled, and what mode it is in
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ManagedZoneDnssecConfigArgs()
        {
        }
    }

    public sealed class ManagedZoneDnssecConfigDefaultKeySpecsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// String mnemonic specifying the DNSSEC algorithm of this key
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Length of the keys in bits
        /// </summary>
        [Input("keyLength")]
        public Input<int>? KeyLength { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies whether this is a key signing key (KSK) or a zone
        /// signing key (ZSK). Key signing keys have the Secure Entry
        /// Point flag set and, when active, will only be used to sign
        /// resource record sets of type DNSKEY. Zone signing keys do
        /// not have the Secure Entry Point flag set and will be used
        /// to sign all other types of resource record sets.
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Identifies what kind of resource this is
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public ManagedZoneDnssecConfigDefaultKeySpecsArgs()
        {
        }
    }

    public sealed class ManagedZoneDnssecConfigDefaultKeySpecsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// String mnemonic specifying the DNSSEC algorithm of this key
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Length of the keys in bits
        /// </summary>
        [Input("keyLength")]
        public Input<int>? KeyLength { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies whether this is a key signing key (KSK) or a zone
        /// signing key (ZSK). Key signing keys have the Secure Entry
        /// Point flag set and, when active, will only be used to sign
        /// resource record sets of type DNSKEY. Zone signing keys do
        /// not have the Secure Entry Point flag set and will be used
        /// to sign all other types of resource record sets.
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Identifies what kind of resource this is
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public ManagedZoneDnssecConfigDefaultKeySpecsGetArgs()
        {
        }
    }

    public sealed class ManagedZoneDnssecConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("defaultKeySpecs")]
        private InputList<ManagedZoneDnssecConfigDefaultKeySpecsGetArgs>? _defaultKeySpecs;

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies parameters that will be used for generating initial DnsKeys
        /// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
        /// you must also provide one for the other.
        /// default_key_specs can only be updated when the state is `off`.  Structure is documented below.
        /// </summary>
        public InputList<ManagedZoneDnssecConfigDefaultKeySpecsGetArgs> DefaultKeySpecs
        {
            get => _defaultKeySpecs ?? (_defaultKeySpecs = new InputList<ManagedZoneDnssecConfigDefaultKeySpecsGetArgs>());
            set => _defaultKeySpecs = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Identifies what kind of resource this is
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the mechanism used to provide authenticated denial-of-existence responses.
        /// non_existence can only be updated when the state is `off`.
        /// </summary>
        [Input("nonExistence")]
        public Input<string>? NonExistence { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies whether DNSSEC is enabled, and what mode it is in
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ManagedZoneDnssecConfigGetArgs()
        {
        }
    }

    public sealed class ManagedZoneForwardingConfigArgs : Pulumi.ResourceArgs
    {
        [Input("targetNameServers", required: true)]
        private InputList<ManagedZoneForwardingConfigTargetNameServersArgs>? _targetNameServers;

        /// <summary>
        /// -
        /// (Required)
        /// List of target name servers to forward to. Cloud DNS will
        /// select the best available name server if more than
        /// one target is given.  Structure is documented below.
        /// </summary>
        public InputList<ManagedZoneForwardingConfigTargetNameServersArgs> TargetNameServers
        {
            get => _targetNameServers ?? (_targetNameServers = new InputList<ManagedZoneForwardingConfigTargetNameServersArgs>());
            set => _targetNameServers = value;
        }

        public ManagedZoneForwardingConfigArgs()
        {
        }
    }

    public sealed class ManagedZoneForwardingConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("targetNameServers", required: true)]
        private InputList<ManagedZoneForwardingConfigTargetNameServersGetArgs>? _targetNameServers;

        /// <summary>
        /// -
        /// (Required)
        /// List of target name servers to forward to. Cloud DNS will
        /// select the best available name server if more than
        /// one target is given.  Structure is documented below.
        /// </summary>
        public InputList<ManagedZoneForwardingConfigTargetNameServersGetArgs> TargetNameServers
        {
            get => _targetNameServers ?? (_targetNameServers = new InputList<ManagedZoneForwardingConfigTargetNameServersGetArgs>());
            set => _targetNameServers = value;
        }

        public ManagedZoneForwardingConfigGetArgs()
        {
        }
    }

    public sealed class ManagedZoneForwardingConfigTargetNameServersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Forwarding path for this TargetNameServer. If unset or `default` Cloud DNS will make forwarding
        /// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
        /// to the Internet. When set to `private`, Cloud DNS will always send queries through VPC for this target
        /// </summary>
        [Input("forwardingPath")]
        public Input<string>? ForwardingPath { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// IPv4 address of a target name server.
        /// </summary>
        [Input("ipv4Address", required: true)]
        public Input<string> Ipv4Address { get; set; } = null!;

        public ManagedZoneForwardingConfigTargetNameServersArgs()
        {
        }
    }

    public sealed class ManagedZoneForwardingConfigTargetNameServersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Forwarding path for this TargetNameServer. If unset or `default` Cloud DNS will make forwarding
        /// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
        /// to the Internet. When set to `private`, Cloud DNS will always send queries through VPC for this target
        /// </summary>
        [Input("forwardingPath")]
        public Input<string>? ForwardingPath { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// IPv4 address of a target name server.
        /// </summary>
        [Input("ipv4Address", required: true)]
        public Input<string> Ipv4Address { get; set; } = null!;

        public ManagedZoneForwardingConfigTargetNameServersGetArgs()
        {
        }
    }

    public sealed class ManagedZonePeeringConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The network with which to peer.  Structure is documented below.
        /// </summary>
        [Input("targetNetwork", required: true)]
        public Input<ManagedZonePeeringConfigTargetNetworkArgs> TargetNetwork { get; set; } = null!;

        public ManagedZonePeeringConfigArgs()
        {
        }
    }

    public sealed class ManagedZonePeeringConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The network with which to peer.  Structure is documented below.
        /// </summary>
        [Input("targetNetwork", required: true)]
        public Input<ManagedZonePeeringConfigTargetNetworkGetArgs> TargetNetwork { get; set; } = null!;

        public ManagedZonePeeringConfigGetArgs()
        {
        }
    }

    public sealed class ManagedZonePeeringConfigTargetNetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The fully qualified URL of the VPC network to forward queries to.
        /// This should be formatted like
        /// `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
        /// </summary>
        [Input("networkUrl", required: true)]
        public Input<string> NetworkUrl { get; set; } = null!;

        public ManagedZonePeeringConfigTargetNetworkArgs()
        {
        }
    }

    public sealed class ManagedZonePeeringConfigTargetNetworkGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The fully qualified URL of the VPC network to forward queries to.
        /// This should be formatted like
        /// `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
        /// </summary>
        [Input("networkUrl", required: true)]
        public Input<string> NetworkUrl { get; set; } = null!;

        public ManagedZonePeeringConfigTargetNetworkGetArgs()
        {
        }
    }

    public sealed class ManagedZonePrivateVisibilityConfigArgs : Pulumi.ResourceArgs
    {
        [Input("networks", required: true)]
        private InputList<ManagedZonePrivateVisibilityConfigNetworksArgs>? _networks;
        public InputList<ManagedZonePrivateVisibilityConfigNetworksArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<ManagedZonePrivateVisibilityConfigNetworksArgs>());
            set => _networks = value;
        }

        public ManagedZonePrivateVisibilityConfigArgs()
        {
        }
    }

    public sealed class ManagedZonePrivateVisibilityConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("networks", required: true)]
        private InputList<ManagedZonePrivateVisibilityConfigNetworksGetArgs>? _networks;
        public InputList<ManagedZonePrivateVisibilityConfigNetworksGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<ManagedZonePrivateVisibilityConfigNetworksGetArgs>());
            set => _networks = value;
        }

        public ManagedZonePrivateVisibilityConfigGetArgs()
        {
        }
    }

    public sealed class ManagedZonePrivateVisibilityConfigNetworksArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The fully qualified URL of the VPC network to forward queries to.
        /// This should be formatted like
        /// `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
        /// </summary>
        [Input("networkUrl", required: true)]
        public Input<string> NetworkUrl { get; set; } = null!;

        public ManagedZonePrivateVisibilityConfigNetworksArgs()
        {
        }
    }

    public sealed class ManagedZonePrivateVisibilityConfigNetworksGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The fully qualified URL of the VPC network to forward queries to.
        /// This should be formatted like
        /// `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
        /// </summary>
        [Input("networkUrl", required: true)]
        public Input<string> NetworkUrl { get; set; } = null!;

        public ManagedZonePrivateVisibilityConfigNetworksGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ManagedZoneDnssecConfig
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies parameters that will be used for generating initial DnsKeys
        /// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
        /// you must also provide one for the other.
        /// default_key_specs can only be updated when the state is `off`.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<ManagedZoneDnssecConfigDefaultKeySpecs> DefaultKeySpecs;
        /// <summary>
        /// -
        /// (Optional)
        /// Identifies what kind of resource this is
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the mechanism used to provide authenticated denial-of-existence responses.
        /// non_existence can only be updated when the state is `off`.
        /// </summary>
        public readonly string NonExistence;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies whether DNSSEC is enabled, and what mode it is in
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private ManagedZoneDnssecConfig(
            ImmutableArray<ManagedZoneDnssecConfigDefaultKeySpecs> defaultKeySpecs,
            string? kind,
            string nonExistence,
            string? state)
        {
            DefaultKeySpecs = defaultKeySpecs;
            Kind = kind;
            NonExistence = nonExistence;
            State = state;
        }
    }

    [OutputType]
    public sealed class ManagedZoneDnssecConfigDefaultKeySpecs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// String mnemonic specifying the DNSSEC algorithm of this key
        /// </summary>
        public readonly string? Algorithm;
        /// <summary>
        /// -
        /// (Optional)
        /// Length of the keys in bits
        /// </summary>
        public readonly int? KeyLength;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies whether this is a key signing key (KSK) or a zone
        /// signing key (ZSK). Key signing keys have the Secure Entry
        /// Point flag set and, when active, will only be used to sign
        /// resource record sets of type DNSKEY. Zone signing keys do
        /// not have the Secure Entry Point flag set and will be used
        /// to sign all other types of resource record sets.
        /// </summary>
        public readonly string? KeyType;
        /// <summary>
        /// -
        /// (Optional)
        /// Identifies what kind of resource this is
        /// </summary>
        public readonly string? Kind;

        [OutputConstructor]
        private ManagedZoneDnssecConfigDefaultKeySpecs(
            string? algorithm,
            int? keyLength,
            string? keyType,
            string? kind)
        {
            Algorithm = algorithm;
            KeyLength = keyLength;
            KeyType = keyType;
            Kind = kind;
        }
    }

    [OutputType]
    public sealed class ManagedZoneForwardingConfig
    {
        /// <summary>
        /// -
        /// (Required)
        /// List of target name servers to forward to. Cloud DNS will
        /// select the best available name server if more than
        /// one target is given.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<ManagedZoneForwardingConfigTargetNameServers> TargetNameServers;

        [OutputConstructor]
        private ManagedZoneForwardingConfig(ImmutableArray<ManagedZoneForwardingConfigTargetNameServers> targetNameServers)
        {
            TargetNameServers = targetNameServers;
        }
    }

    [OutputType]
    public sealed class ManagedZoneForwardingConfigTargetNameServers
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Forwarding path for this TargetNameServer. If unset or `default` Cloud DNS will make forwarding
        /// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
        /// to the Internet. When set to `private`, Cloud DNS will always send queries through VPC for this target
        /// </summary>
        public readonly string? ForwardingPath;
        /// <summary>
        /// -
        /// (Required)
        /// IPv4 address of a target name server.
        /// </summary>
        public readonly string Ipv4Address;

        [OutputConstructor]
        private ManagedZoneForwardingConfigTargetNameServers(
            string? forwardingPath,
            string ipv4Address)
        {
            ForwardingPath = forwardingPath;
            Ipv4Address = ipv4Address;
        }
    }

    [OutputType]
    public sealed class ManagedZonePeeringConfig
    {
        /// <summary>
        /// -
        /// (Required)
        /// The network with which to peer.  Structure is documented below.
        /// </summary>
        public readonly ManagedZonePeeringConfigTargetNetwork TargetNetwork;

        [OutputConstructor]
        private ManagedZonePeeringConfig(ManagedZonePeeringConfigTargetNetwork targetNetwork)
        {
            TargetNetwork = targetNetwork;
        }
    }

    [OutputType]
    public sealed class ManagedZonePeeringConfigTargetNetwork
    {
        /// <summary>
        /// -
        /// (Required)
        /// The fully qualified URL of the VPC network to forward queries to.
        /// This should be formatted like
        /// `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
        /// </summary>
        public readonly string NetworkUrl;

        [OutputConstructor]
        private ManagedZonePeeringConfigTargetNetwork(string networkUrl)
        {
            NetworkUrl = networkUrl;
        }
    }

    [OutputType]
    public sealed class ManagedZonePrivateVisibilityConfig
    {
        public readonly ImmutableArray<ManagedZonePrivateVisibilityConfigNetworks> Networks;

        [OutputConstructor]
        private ManagedZonePrivateVisibilityConfig(ImmutableArray<ManagedZonePrivateVisibilityConfigNetworks> networks)
        {
            Networks = networks;
        }
    }

    [OutputType]
    public sealed class ManagedZonePrivateVisibilityConfigNetworks
    {
        /// <summary>
        /// -
        /// (Required)
        /// The fully qualified URL of the VPC network to forward queries to.
        /// This should be formatted like
        /// `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
        /// </summary>
        public readonly string NetworkUrl;

        [OutputConstructor]
        private ManagedZonePrivateVisibilityConfigNetworks(string networkUrl)
        {
            NetworkUrl = networkUrl;
        }
    }
    }
}
