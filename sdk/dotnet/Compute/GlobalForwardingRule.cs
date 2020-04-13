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
    /// Represents a GlobalForwardingRule resource. Global forwarding rules are
    /// used to forward traffic to the correct load balancer for HTTP load
    /// balancing. Global forwarding rules can only be used for HTTP load
    /// balancing.
    /// 
    /// For more information, see
    /// https://cloud.google.com/compute/docs/load-balancing/http/
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_global_forwarding_rule.html.markdown.
    /// </summary>
    public partial class GlobalForwardingRule : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the
        /// forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load
        /// balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional
        /// forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty,
        /// an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding
        /// rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing
        /// scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for
        /// the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be
        /// automatically allocated from the IP range of the subnet or network configured for this forwarding rule. An
        /// address must be specified by a literal IP address. ~&gt; **NOTE**: While the API allows you to specify
        /// various resource paths for an address resource instead, Terraform requires this to specifically be an IP
        /// address to avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The IP protocol to which this rule applies. Valid options are TCP,
        /// UDP, ESP, AH, SCTP or ICMP. When the load balancing scheme is
        /// INTERNAL_SELF_MANAGED, only TCP is valid.
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The IP Version that will be used by this global forwarding rule.
        /// Valid options are IPV4 or IPV6.
        /// </summary>
        [Output("ipVersion")]
        public Output<string?> IpVersion { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this forwarding rule. A list of key-&gt;value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// This signifies what the GlobalForwardingRule will be used for.
        /// The value of INTERNAL_SELF_MANAGED means that this will be used for
        /// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
        /// will be used for External Global Load Balancing (HTTP(S) LB,
        /// External TCP/UDP LB, SSL Proxy)
        /// NOTE: Currently global forwarding rules cannot be used for INTERNAL
        /// load balancing.
        /// </summary>
        [Output("loadBalancingScheme")]
        public Output<string?> LoadBalancingScheme { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Opaque filter criteria used by Loadbalancer to restrict routing
        /// configuration to a limited set xDS compliant clients. In their xDS
        /// requests to Loadbalancer, xDS clients present node metadata. If a
        /// match takes place, the relevant routing configuration is made available
        /// to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set
        /// to MATCH_ANY, at least one of the filterLabels must match the
        /// corresponding label provided in the metadata. If its filterMatchCriteria
        /// is set to MATCH_ALL, then all of its filterLabels must match with
        /// corresponding labels in the provided metadata.
        /// metadataFilters specified here can be overridden by those specified in
        /// the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their
        /// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
        /// </summary>
        [Output("metadataFilters")]
        public Output<ImmutableArray<Outputs.GlobalForwardingRuleMetadataFilters>> MetadataFilters { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// Name of the metadata label. The length must be between
        /// 1 and 1024 characters, inclusive.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This field is not used for external load balancing. For INTERNAL_SELF_MANAGED load balancing, this field
        /// identifies the network that the load balanced IP should belong to for this global forwarding rule. If this
        /// field is not specified, the default network will be used.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// This field is used along with the target field for TargetHttpProxy,
        /// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
        /// TargetPool, TargetInstance.
        /// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
        /// addressed to ports in the specified range will be forwarded to target.
        /// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
        /// disjoint port ranges.
        /// Some types of forwarding target have constraints on the acceptable
        /// ports:
        /// * TargetHttpProxy: 80, 8080
        /// * TargetHttpsProxy: 443
        /// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        /// 1883, 5222
        /// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        /// 1883, 5222
        /// * TargetVpnGateway: 500, 4500
        /// </summary>
        [Output("portRange")]
        public Output<string?> PortRange { get; private set; } = null!;

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
        /// -
        /// (Required)
        /// The URL of the target resource to receive the matched traffic.
        /// The forwarded traffic must be of a type appropriate to the target object.
        /// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
        /// are valid.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalForwardingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalForwardingRule(string name, GlobalForwardingRuleArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GlobalForwardingRule(string name, Input<string> id, GlobalForwardingRuleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GlobalForwardingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalForwardingRule Get(string name, Input<string> id, GlobalForwardingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new GlobalForwardingRule(name, id, state, options);
        }
    }

    public sealed class GlobalForwardingRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the
        /// forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load
        /// balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional
        /// forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty,
        /// an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding
        /// rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing
        /// scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for
        /// the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be
        /// automatically allocated from the IP range of the subnet or network configured for this forwarding rule. An
        /// address must be specified by a literal IP address. ~&gt; **NOTE**: While the API allows you to specify
        /// various resource paths for an address resource instead, Terraform requires this to specifically be an IP
        /// address to avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The IP protocol to which this rule applies. Valid options are TCP,
        /// UDP, ESP, AH, SCTP or ICMP. When the load balancing scheme is
        /// INTERNAL_SELF_MANAGED, only TCP is valid.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The IP Version that will be used by this global forwarding rule.
        /// Valid options are IPV4 or IPV6.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this forwarding rule. A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// This signifies what the GlobalForwardingRule will be used for.
        /// The value of INTERNAL_SELF_MANAGED means that this will be used for
        /// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
        /// will be used for External Global Load Balancing (HTTP(S) LB,
        /// External TCP/UDP LB, SSL Proxy)
        /// NOTE: Currently global forwarding rules cannot be used for INTERNAL
        /// load balancing.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        [Input("metadataFilters")]
        private InputList<Inputs.GlobalForwardingRuleMetadataFiltersArgs>? _metadataFilters;

        /// <summary>
        /// -
        /// (Optional)
        /// Opaque filter criteria used by Loadbalancer to restrict routing
        /// configuration to a limited set xDS compliant clients. In their xDS
        /// requests to Loadbalancer, xDS clients present node metadata. If a
        /// match takes place, the relevant routing configuration is made available
        /// to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set
        /// to MATCH_ANY, at least one of the filterLabels must match the
        /// corresponding label provided in the metadata. If its filterMatchCriteria
        /// is set to MATCH_ALL, then all of its filterLabels must match with
        /// corresponding labels in the provided metadata.
        /// metadataFilters specified here can be overridden by those specified in
        /// the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their
        /// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.GlobalForwardingRuleMetadataFiltersArgs> MetadataFilters
        {
            get => _metadataFilters ?? (_metadataFilters = new InputList<Inputs.GlobalForwardingRuleMetadataFiltersArgs>());
            set => _metadataFilters = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the metadata label. The length must be between
        /// 1 and 1024 characters, inclusive.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This field is not used for external load balancing. For INTERNAL_SELF_MANAGED load balancing, this field
        /// identifies the network that the load balanced IP should belong to for this global forwarding rule. If this
        /// field is not specified, the default network will be used.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// This field is used along with the target field for TargetHttpProxy,
        /// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
        /// TargetPool, TargetInstance.
        /// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
        /// addressed to ports in the specified range will be forwarded to target.
        /// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
        /// disjoint port ranges.
        /// Some types of forwarding target have constraints on the acceptable
        /// ports:
        /// * TargetHttpProxy: 80, 8080
        /// * TargetHttpsProxy: 443
        /// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        /// 1883, 5222
        /// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        /// 1883, 5222
        /// * TargetVpnGateway: 500, 4500
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The URL of the target resource to receive the matched traffic.
        /// The forwarded traffic must be of a type appropriate to the target object.
        /// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
        /// are valid.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        public GlobalForwardingRuleArgs()
        {
        }
    }

    public sealed class GlobalForwardingRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the
        /// forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load
        /// balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional
        /// forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty,
        /// an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding
        /// rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing
        /// scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for
        /// the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be
        /// automatically allocated from the IP range of the subnet or network configured for this forwarding rule. An
        /// address must be specified by a literal IP address. ~&gt; **NOTE**: While the API allows you to specify
        /// various resource paths for an address resource instead, Terraform requires this to specifically be an IP
        /// address to avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The IP protocol to which this rule applies. Valid options are TCP,
        /// UDP, ESP, AH, SCTP or ICMP. When the load balancing scheme is
        /// INTERNAL_SELF_MANAGED, only TCP is valid.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The IP Version that will be used by this global forwarding rule.
        /// Valid options are IPV4 or IPV6.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this forwarding rule. A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// This signifies what the GlobalForwardingRule will be used for.
        /// The value of INTERNAL_SELF_MANAGED means that this will be used for
        /// Internal Global HTTP(S) LB. The value of EXTERNAL means that this
        /// will be used for External Global Load Balancing (HTTP(S) LB,
        /// External TCP/UDP LB, SSL Proxy)
        /// NOTE: Currently global forwarding rules cannot be used for INTERNAL
        /// load balancing.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        [Input("metadataFilters")]
        private InputList<Inputs.GlobalForwardingRuleMetadataFiltersGetArgs>? _metadataFilters;

        /// <summary>
        /// -
        /// (Optional)
        /// Opaque filter criteria used by Loadbalancer to restrict routing
        /// configuration to a limited set xDS compliant clients. In their xDS
        /// requests to Loadbalancer, xDS clients present node metadata. If a
        /// match takes place, the relevant routing configuration is made available
        /// to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set
        /// to MATCH_ANY, at least one of the filterLabels must match the
        /// corresponding label provided in the metadata. If its filterMatchCriteria
        /// is set to MATCH_ALL, then all of its filterLabels must match with
        /// corresponding labels in the provided metadata.
        /// metadataFilters specified here can be overridden by those specified in
        /// the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their
        /// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.GlobalForwardingRuleMetadataFiltersGetArgs> MetadataFilters
        {
            get => _metadataFilters ?? (_metadataFilters = new InputList<Inputs.GlobalForwardingRuleMetadataFiltersGetArgs>());
            set => _metadataFilters = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the metadata label. The length must be between
        /// 1 and 1024 characters, inclusive.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This field is not used for external load balancing. For INTERNAL_SELF_MANAGED load balancing, this field
        /// identifies the network that the load balanced IP should belong to for this global forwarding rule. If this
        /// field is not specified, the default network will be used.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// This field is used along with the target field for TargetHttpProxy,
        /// TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
        /// TargetPool, TargetInstance.
        /// Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
        /// addressed to ports in the specified range will be forwarded to target.
        /// Forwarding rules with the same [IPAddress, IPProtocol] pair must have
        /// disjoint port ranges.
        /// Some types of forwarding target have constraints on the acceptable
        /// ports:
        /// * TargetHttpProxy: 80, 8080
        /// * TargetHttpsProxy: 443
        /// * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        /// 1883, 5222
        /// * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        /// 1883, 5222
        /// * TargetVpnGateway: 500, 4500
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

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

        /// <summary>
        /// -
        /// (Required)
        /// The URL of the target resource to receive the matched traffic.
        /// The forwarded traffic must be of a type appropriate to the target object.
        /// For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
        /// are valid.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public GlobalForwardingRuleState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class GlobalForwardingRuleMetadataFiltersArgs : Pulumi.ResourceArgs
    {
        [Input("filterLabels", required: true)]
        private InputList<GlobalForwardingRuleMetadataFiltersFilterLabelsArgs>? _filterLabels;

        /// <summary>
        /// -
        /// (Required)
        /// The list of label value pairs that must match labels in the
        /// provided metadata based on filterMatchCriteria
        /// This list must not be empty and can have at the most 64 entries.  Structure is documented below.
        /// </summary>
        public InputList<GlobalForwardingRuleMetadataFiltersFilterLabelsArgs> FilterLabels
        {
            get => _filterLabels ?? (_filterLabels = new InputList<GlobalForwardingRuleMetadataFiltersFilterLabelsArgs>());
            set => _filterLabels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Specifies how individual filterLabel matches within the list of
        /// filterLabels contribute towards the overall metadataFilter match.
        /// MATCH_ANY - At least one of the filterLabels must have a matching
        /// label in the provided metadata.
        /// MATCH_ALL - All filterLabels must have matching labels in the
        /// provided metadata.
        /// </summary>
        [Input("filterMatchCriteria", required: true)]
        public Input<string> FilterMatchCriteria { get; set; } = null!;

        public GlobalForwardingRuleMetadataFiltersArgs()
        {
        }
    }

    public sealed class GlobalForwardingRuleMetadataFiltersFilterLabelsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// Name of the metadata label. The length must be between
        /// 1 and 1024 characters, inclusive.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The value that the label must match. The value has a maximum
        /// length of 1024 characters.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GlobalForwardingRuleMetadataFiltersFilterLabelsArgs()
        {
        }
    }

    public sealed class GlobalForwardingRuleMetadataFiltersFilterLabelsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// Name of the metadata label. The length must be between
        /// 1 and 1024 characters, inclusive.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The value that the label must match. The value has a maximum
        /// length of 1024 characters.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GlobalForwardingRuleMetadataFiltersFilterLabelsGetArgs()
        {
        }
    }

    public sealed class GlobalForwardingRuleMetadataFiltersGetArgs : Pulumi.ResourceArgs
    {
        [Input("filterLabels", required: true)]
        private InputList<GlobalForwardingRuleMetadataFiltersFilterLabelsGetArgs>? _filterLabels;

        /// <summary>
        /// -
        /// (Required)
        /// The list of label value pairs that must match labels in the
        /// provided metadata based on filterMatchCriteria
        /// This list must not be empty and can have at the most 64 entries.  Structure is documented below.
        /// </summary>
        public InputList<GlobalForwardingRuleMetadataFiltersFilterLabelsGetArgs> FilterLabels
        {
            get => _filterLabels ?? (_filterLabels = new InputList<GlobalForwardingRuleMetadataFiltersFilterLabelsGetArgs>());
            set => _filterLabels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Specifies how individual filterLabel matches within the list of
        /// filterLabels contribute towards the overall metadataFilter match.
        /// MATCH_ANY - At least one of the filterLabels must have a matching
        /// label in the provided metadata.
        /// MATCH_ALL - All filterLabels must have matching labels in the
        /// provided metadata.
        /// </summary>
        [Input("filterMatchCriteria", required: true)]
        public Input<string> FilterMatchCriteria { get; set; } = null!;

        public GlobalForwardingRuleMetadataFiltersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GlobalForwardingRuleMetadataFilters
    {
        /// <summary>
        /// -
        /// (Required)
        /// The list of label value pairs that must match labels in the
        /// provided metadata based on filterMatchCriteria
        /// This list must not be empty and can have at the most 64 entries.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<GlobalForwardingRuleMetadataFiltersFilterLabels> FilterLabels;
        /// <summary>
        /// -
        /// (Required)
        /// Specifies how individual filterLabel matches within the list of
        /// filterLabels contribute towards the overall metadataFilter match.
        /// MATCH_ANY - At least one of the filterLabels must have a matching
        /// label in the provided metadata.
        /// MATCH_ALL - All filterLabels must have matching labels in the
        /// provided metadata.
        /// </summary>
        public readonly string FilterMatchCriteria;

        [OutputConstructor]
        private GlobalForwardingRuleMetadataFilters(
            ImmutableArray<GlobalForwardingRuleMetadataFiltersFilterLabels> filterLabels,
            string filterMatchCriteria)
        {
            FilterLabels = filterLabels;
            FilterMatchCriteria = filterMatchCriteria;
        }
    }

    [OutputType]
    public sealed class GlobalForwardingRuleMetadataFiltersFilterLabels
    {
        /// <summary>
        /// -
        /// (Required)
        /// Name of the metadata label. The length must be between
        /// 1 and 1024 characters, inclusive.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// -
        /// (Required)
        /// The value that the label must match. The value has a maximum
        /// length of 1024 characters.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GlobalForwardingRuleMetadataFiltersFilterLabels(
            string name,
            string value)
        {
            Name = name;
            Value = value;
        }
    }
    }
}
