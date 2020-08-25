# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ForwardingRule']


class ForwardingRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_ports: Optional[pulumi.Input[bool]] = None,
                 allow_global_access: Optional[pulumi.Input[bool]] = None,
                 backend_service: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 is_mirroring_collector: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 load_balancing_scheme: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_tier: Optional[pulumi.Input[str]] = None,
                 port_range: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_label: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A ForwardingRule resource. A ForwardingRule resource specifies which pool
        of target virtual machines to forward a packet to if it matches the given
        [IPAddress, IPProtocol, portRange] tuple.

        To get more information about ForwardingRule, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/forwardingRules)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/network/forwarding-rules)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_ports: For internal TCP/UDP load balancing (i.e. load balancing scheme is
               INTERNAL and protocol is TCP/UDP), set this to true to allow packets
               addressed to any ports to be forwarded to the backends configured
               with this forwarding rule. Used with backend service. Cannot be set
               if port or portRange are set.
        :param pulumi.Input[bool] allow_global_access: If true, clients can access ILB from all regions.
               Otherwise only allows from the local region the ILB is located at.
        :param pulumi.Input[str] backend_service: A BackendService to receive the matched traffic. This is used only
               for INTERNAL load balancing.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] ip_address: The IP address that this forwarding rule is serving on behalf of.
               Addresses are restricted based on the forwarding rule's load balancing
               scheme (EXTERNAL or INTERNAL) and scope (global or regional).
               When the load balancing scheme is EXTERNAL, for global forwarding
               rules, the address must be a global IP, and for regional forwarding
               rules, the address must live in the same region as the forwarding
               rule. If this field is empty, an ephemeral IPv4 address from the same
               scope (global or regional) will be assigned. A regional forwarding
               rule supports IPv4 only. A global forwarding rule supports either IPv4
               or IPv6.
               When the load balancing scheme is INTERNAL, this can only be an RFC
               1918 IP address belonging to the network/subnet configured for the
               forwarding rule. By default, if this field is empty, an ephemeral
               internal IP address will be automatically allocated from the IP range
               of the subnet or network configured for this forwarding rule.
               An address must be specified by a literal IP address. > **NOTE:** While
               the API allows you to specify various resource paths for an address resource
               instead, this provider requires this to specifically be an IP address to
               avoid needing to fetching the IP address from resource paths on refresh
               or unnecessary diffs.
        :param pulumi.Input[str] ip_protocol: The IP protocol to which this rule applies.
               When the load balancing scheme is INTERNAL, only TCP and UDP are
               valid.
               Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
        :param pulumi.Input[bool] is_mirroring_collector: Indicates whether or not this load balancer can be used
               as a collector for packet mirroring. To prevent mirroring loops,
               instances behind this load balancer will not have their traffic
               mirrored even if a PacketMirroring rule applies to them. This
               can only be set to true for load balancers that have their
               loadBalancingScheme set to INTERNAL.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this forwarding rule.  A list of key->value pairs.
        :param pulumi.Input[str] load_balancing_scheme: This signifies what the ForwardingRule will be used for and can be
               EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
               Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
               and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
               INTERNAL is used for protocol forwarding to VMs from an internal IP address,
               and internal TCP/UDP load balancers.
               INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
               Default value is `EXTERNAL`.
               Possible values are `EXTERNAL`, `INTERNAL`, and `INTERNAL_MANAGED`.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: For internal load balancing, this field identifies the network that
               the load balanced IP should belong to for this Forwarding Rule. If
               this field is not specified, the default network will be used.
               This field is only used for INTERNAL load balancing.
        :param pulumi.Input[str] network_tier: The networking tier used for configuring this address. If this field is not
               specified, it is assumed to be PREMIUM.
               Possible values are `PREMIUM` and `STANDARD`.
        :param pulumi.Input[str] port_range: This field is used along with the target field for TargetHttpProxy,
               TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
               TargetPool, TargetInstance.
               Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
               addressed to ports in the specified range will be forwarded to target.
               Forwarding rules with the same [IPAddress, IPProtocol] pair must have
               disjoint port ranges.
               Some types of forwarding target have constraints on the acceptable
               ports:
               * TargetHttpProxy: 80, 8080
               * TargetHttpsProxy: 443
               * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetVpnGateway: 500, 4500
        :param pulumi.Input[List[pulumi.Input[str]]] ports: This field is used along with the backend_service field for internal
               load balancing.
               When the load balancing scheme is INTERNAL, a single port or a comma
               separated list of ports can be configured. Only packets addressed to
               these ports will be forwarded to the backends configured with this
               forwarding rule.
               You may specify a maximum of up to 5 ports.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: A reference to the region where the regional forwarding rule resides.
               This field is not applicable to global forwarding rules.
        :param pulumi.Input[str] service_label: An optional prefix to the service name for this Forwarding Rule.
               If specified, will be the first label of the fully qualified service
               name.
               The label must be 1-63 characters long, and comply with RFC1035.
               Specifically, the label must be 1-63 characters long and match the
               regular expression `a-z?` which means the first
               character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
               This field is only used for INTERNAL load balancing.
        :param pulumi.Input[str] subnetwork: The subnetwork that the load balanced IP should belong to for this
               Forwarding Rule.  This field is only used for INTERNAL load balancing.
               If the network specified is in auto subnet mode, this field is
               optional. However, if the network is in custom subnet mode, a
               subnetwork must be specified.
        :param pulumi.Input[str] target: The URL of the target resource to receive the matched traffic.
               The target must live in the same region as the forwarding rule.
               The forwarded traffic must be of a type appropriate to the target
               object.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['all_ports'] = all_ports
            __props__['allow_global_access'] = allow_global_access
            __props__['backend_service'] = backend_service
            __props__['description'] = description
            __props__['ip_address'] = ip_address
            __props__['ip_protocol'] = ip_protocol
            __props__['is_mirroring_collector'] = is_mirroring_collector
            __props__['labels'] = labels
            __props__['load_balancing_scheme'] = load_balancing_scheme
            __props__['name'] = name
            __props__['network'] = network
            __props__['network_tier'] = network_tier
            __props__['port_range'] = port_range
            __props__['ports'] = ports
            __props__['project'] = project
            __props__['region'] = region
            __props__['service_label'] = service_label
            __props__['subnetwork'] = subnetwork
            __props__['target'] = target
            __props__['creation_timestamp'] = None
            __props__['label_fingerprint'] = None
            __props__['self_link'] = None
            __props__['service_name'] = None
        super(ForwardingRule, __self__).__init__(
            'gcp:compute/forwardingRule:ForwardingRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_ports: Optional[pulumi.Input[bool]] = None,
            allow_global_access: Optional[pulumi.Input[bool]] = None,
            backend_service: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            ip_protocol: Optional[pulumi.Input[str]] = None,
            is_mirroring_collector: Optional[pulumi.Input[bool]] = None,
            label_fingerprint: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            load_balancing_scheme: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            network_tier: Optional[pulumi.Input[str]] = None,
            port_range: Optional[pulumi.Input[str]] = None,
            ports: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            service_label: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            subnetwork: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None) -> 'ForwardingRule':
        """
        Get an existing ForwardingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_ports: For internal TCP/UDP load balancing (i.e. load balancing scheme is
               INTERNAL and protocol is TCP/UDP), set this to true to allow packets
               addressed to any ports to be forwarded to the backends configured
               with this forwarding rule. Used with backend service. Cannot be set
               if port or portRange are set.
        :param pulumi.Input[bool] allow_global_access: If true, clients can access ILB from all regions.
               Otherwise only allows from the local region the ILB is located at.
        :param pulumi.Input[str] backend_service: A BackendService to receive the matched traffic. This is used only
               for INTERNAL load balancing.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] ip_address: The IP address that this forwarding rule is serving on behalf of.
               Addresses are restricted based on the forwarding rule's load balancing
               scheme (EXTERNAL or INTERNAL) and scope (global or regional).
               When the load balancing scheme is EXTERNAL, for global forwarding
               rules, the address must be a global IP, and for regional forwarding
               rules, the address must live in the same region as the forwarding
               rule. If this field is empty, an ephemeral IPv4 address from the same
               scope (global or regional) will be assigned. A regional forwarding
               rule supports IPv4 only. A global forwarding rule supports either IPv4
               or IPv6.
               When the load balancing scheme is INTERNAL, this can only be an RFC
               1918 IP address belonging to the network/subnet configured for the
               forwarding rule. By default, if this field is empty, an ephemeral
               internal IP address will be automatically allocated from the IP range
               of the subnet or network configured for this forwarding rule.
               An address must be specified by a literal IP address. > **NOTE:** While
               the API allows you to specify various resource paths for an address resource
               instead, this provider requires this to specifically be an IP address to
               avoid needing to fetching the IP address from resource paths on refresh
               or unnecessary diffs.
        :param pulumi.Input[str] ip_protocol: The IP protocol to which this rule applies.
               When the load balancing scheme is INTERNAL, only TCP and UDP are
               valid.
               Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
        :param pulumi.Input[bool] is_mirroring_collector: Indicates whether or not this load balancer can be used
               as a collector for packet mirroring. To prevent mirroring loops,
               instances behind this load balancer will not have their traffic
               mirrored even if a PacketMirroring rule applies to them. This
               can only be set to true for load balancers that have their
               loadBalancingScheme set to INTERNAL.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this forwarding rule.  A list of key->value pairs.
        :param pulumi.Input[str] load_balancing_scheme: This signifies what the ForwardingRule will be used for and can be
               EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
               Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
               and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
               INTERNAL is used for protocol forwarding to VMs from an internal IP address,
               and internal TCP/UDP load balancers.
               INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
               Default value is `EXTERNAL`.
               Possible values are `EXTERNAL`, `INTERNAL`, and `INTERNAL_MANAGED`.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: For internal load balancing, this field identifies the network that
               the load balanced IP should belong to for this Forwarding Rule. If
               this field is not specified, the default network will be used.
               This field is only used for INTERNAL load balancing.
        :param pulumi.Input[str] network_tier: The networking tier used for configuring this address. If this field is not
               specified, it is assumed to be PREMIUM.
               Possible values are `PREMIUM` and `STANDARD`.
        :param pulumi.Input[str] port_range: This field is used along with the target field for TargetHttpProxy,
               TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
               TargetPool, TargetInstance.
               Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
               addressed to ports in the specified range will be forwarded to target.
               Forwarding rules with the same [IPAddress, IPProtocol] pair must have
               disjoint port ranges.
               Some types of forwarding target have constraints on the acceptable
               ports:
               * TargetHttpProxy: 80, 8080
               * TargetHttpsProxy: 443
               * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetVpnGateway: 500, 4500
        :param pulumi.Input[List[pulumi.Input[str]]] ports: This field is used along with the backend_service field for internal
               load balancing.
               When the load balancing scheme is INTERNAL, a single port or a comma
               separated list of ports can be configured. Only packets addressed to
               these ports will be forwarded to the backends configured with this
               forwarding rule.
               You may specify a maximum of up to 5 ports.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: A reference to the region where the regional forwarding rule resides.
               This field is not applicable to global forwarding rules.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] service_label: An optional prefix to the service name for this Forwarding Rule.
               If specified, will be the first label of the fully qualified service
               name.
               The label must be 1-63 characters long, and comply with RFC1035.
               Specifically, the label must be 1-63 characters long and match the
               regular expression `a-z?` which means the first
               character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
               This field is only used for INTERNAL load balancing.
        :param pulumi.Input[str] service_name: The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.
        :param pulumi.Input[str] subnetwork: The subnetwork that the load balanced IP should belong to for this
               Forwarding Rule.  This field is only used for INTERNAL load balancing.
               If the network specified is in auto subnet mode, this field is
               optional. However, if the network is in custom subnet mode, a
               subnetwork must be specified.
        :param pulumi.Input[str] target: The URL of the target resource to receive the matched traffic.
               The target must live in the same region as the forwarding rule.
               The forwarded traffic must be of a type appropriate to the target
               object.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["all_ports"] = all_ports
        __props__["allow_global_access"] = allow_global_access
        __props__["backend_service"] = backend_service
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["ip_address"] = ip_address
        __props__["ip_protocol"] = ip_protocol
        __props__["is_mirroring_collector"] = is_mirroring_collector
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["load_balancing_scheme"] = load_balancing_scheme
        __props__["name"] = name
        __props__["network"] = network
        __props__["network_tier"] = network_tier
        __props__["port_range"] = port_range
        __props__["ports"] = ports
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["service_label"] = service_label
        __props__["service_name"] = service_name
        __props__["subnetwork"] = subnetwork
        __props__["target"] = target
        return ForwardingRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allPorts")
    def all_ports(self) -> Optional[bool]:
        """
        For internal TCP/UDP load balancing (i.e. load balancing scheme is
        INTERNAL and protocol is TCP/UDP), set this to true to allow packets
        addressed to any ports to be forwarded to the backends configured
        with this forwarding rule. Used with backend service. Cannot be set
        if port or portRange are set.
        """
        return pulumi.get(self, "all_ports")

    @property
    @pulumi.getter(name="allowGlobalAccess")
    def allow_global_access(self) -> Optional[bool]:
        """
        If true, clients can access ILB from all regions.
        Otherwise only allows from the local region the ILB is located at.
        """
        return pulumi.get(self, "allow_global_access")

    @property
    @pulumi.getter(name="backendService")
    def backend_service(self) -> Optional[str]:
        """
        A BackendService to receive the matched traffic. This is used only
        for INTERNAL load balancing.
        """
        return pulumi.get(self, "backend_service")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The IP address that this forwarding rule is serving on behalf of.
        Addresses are restricted based on the forwarding rule's load balancing
        scheme (EXTERNAL or INTERNAL) and scope (global or regional).
        When the load balancing scheme is EXTERNAL, for global forwarding
        rules, the address must be a global IP, and for regional forwarding
        rules, the address must live in the same region as the forwarding
        rule. If this field is empty, an ephemeral IPv4 address from the same
        scope (global or regional) will be assigned. A regional forwarding
        rule supports IPv4 only. A global forwarding rule supports either IPv4
        or IPv6.
        When the load balancing scheme is INTERNAL, this can only be an RFC
        1918 IP address belonging to the network/subnet configured for the
        forwarding rule. By default, if this field is empty, an ephemeral
        internal IP address will be automatically allocated from the IP range
        of the subnet or network configured for this forwarding rule.
        An address must be specified by a literal IP address. > **NOTE:** While
        the API allows you to specify various resource paths for an address resource
        instead, this provider requires this to specifically be an IP address to
        avoid needing to fetching the IP address from resource paths on refresh
        or unnecessary diffs.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> str:
        """
        The IP protocol to which this rule applies.
        When the load balancing scheme is INTERNAL, only TCP and UDP are
        valid.
        Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="isMirroringCollector")
    def is_mirroring_collector(self) -> Optional[bool]:
        """
        Indicates whether or not this load balancer can be used
        as a collector for packet mirroring. To prevent mirroring loops,
        instances behind this load balancer will not have their traffic
        mirrored even if a PacketMirroring rule applies to them. This
        can only be set to true for load balancers that have their
        loadBalancingScheme set to INTERNAL.
        """
        return pulumi.get(self, "is_mirroring_collector")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        The fingerprint used for optimistic locking of this resource. Used internally during updates.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        Labels to apply to this forwarding rule.  A list of key->value pairs.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="loadBalancingScheme")
    def load_balancing_scheme(self) -> Optional[str]:
        """
        This signifies what the ForwardingRule will be used for and can be
        EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
        Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
        and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
        INTERNAL is used for protocol forwarding to VMs from an internal IP address,
        and internal TCP/UDP load balancers.
        INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
        Default value is `EXTERNAL`.
        Possible values are `EXTERNAL`, `INTERNAL`, and `INTERNAL_MANAGED`.
        """
        return pulumi.get(self, "load_balancing_scheme")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource; provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        For internal load balancing, this field identifies the network that
        the load balanced IP should belong to for this Forwarding Rule. If
        this field is not specified, the default network will be used.
        This field is only used for INTERNAL load balancing.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkTier")
    def network_tier(self) -> str:
        """
        The networking tier used for configuring this address. If this field is not
        specified, it is assumed to be PREMIUM.
        Possible values are `PREMIUM` and `STANDARD`.
        """
        return pulumi.get(self, "network_tier")

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> Optional[str]:
        """
        This field is used along with the target field for TargetHttpProxy,
        TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
        TargetPool, TargetInstance.
        Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
        addressed to ports in the specified range will be forwarded to target.
        Forwarding rules with the same [IPAddress, IPProtocol] pair must have
        disjoint port ranges.
        Some types of forwarding target have constraints on the acceptable
        ports:
        * TargetHttpProxy: 80, 8080
        * TargetHttpsProxy: 443
        * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        1883, 5222
        * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
        1883, 5222
        * TargetVpnGateway: 500, 4500
        """
        return pulumi.get(self, "port_range")

    @property
    @pulumi.getter
    def ports(self) -> Optional[List[str]]:
        """
        This field is used along with the backend_service field for internal
        load balancing.
        When the load balancing scheme is INTERNAL, a single port or a comma
        separated list of ports can be configured. Only packets addressed to
        these ports will be forwarded to the backends configured with this
        forwarding rule.
        You may specify a maximum of up to 5 ports.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        A reference to the region where the regional forwarding rule resides.
        This field is not applicable to global forwarding rules.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serviceLabel")
    def service_label(self) -> Optional[str]:
        """
        An optional prefix to the service name for this Forwarding Rule.
        If specified, will be the first label of the fully qualified service
        name.
        The label must be 1-63 characters long, and comply with RFC1035.
        Specifically, the label must be 1-63 characters long and match the
        regular expression `a-z?` which means the first
        character must be a lowercase letter, and all following characters
        must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        This field is only used for INTERNAL load balancing.
        """
        return pulumi.get(self, "service_label")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def subnetwork(self) -> str:
        """
        The subnetwork that the load balanced IP should belong to for this
        Forwarding Rule.  This field is only used for INTERNAL load balancing.
        If the network specified is in auto subnet mode, this field is
        optional. However, if the network is in custom subnet mode, a
        subnetwork must be specified.
        """
        return pulumi.get(self, "subnetwork")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        The URL of the target resource to receive the matched traffic.
        The target must live in the same region as the forwarding rule.
        The forwarded traffic must be of a type appropriate to the target
        object.
        """
        return pulumi.get(self, "target")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

