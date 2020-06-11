# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ForwardingRule(pulumi.CustomResource):
    all_ports: pulumi.Output[bool]
    """
    For internal TCP/UDP load balancing (i.e. load balancing scheme is
    INTERNAL and protocol is TCP/UDP), set this to true to allow packets
    addressed to any ports to be forwarded to the backends configured
    with this forwarding rule. Used with backend service. Cannot be set
    if port or portRange are set.
    """
    allow_global_access: pulumi.Output[bool]
    """
    If true, clients can access ILB from all regions.
    Otherwise only allows from the local region the ILB is located at.
    """
    backend_service: pulumi.Output[str]
    """
    A BackendService to receive the matched traffic. This is used only
    for INTERNAL load balancing.
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource. Provide this property when
    you create the resource.
    """
    ip_address: pulumi.Output[str]
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
    An address must be specified by a literal IP address. > **NOTE**: While
    the API allows you to specify various resource paths for an address resource
    instead, this provider requires this to specifically be an IP address to
    avoid needing to fetching the IP address from resource paths on refresh
    or unnecessary diffs.
    """
    ip_protocol: pulumi.Output[str]
    """
    The IP protocol to which this rule applies.
    When the load balancing scheme is INTERNAL, only TCP and UDP are
    valid.
    """
    is_mirroring_collector: pulumi.Output[bool]
    """
    Indicates whether or not this load balancer can be used
    as a collector for packet mirroring. To prevent mirroring loops,
    instances behind this load balancer will not have their traffic
    mirrored even if a PacketMirroring rule applies to them. This
    can only be set to true for load balancers that have their
    loadBalancingScheme set to INTERNAL.
    """
    label_fingerprint: pulumi.Output[str]
    """
    The fingerprint used for optimistic locking of this resource. Used internally during updates.
    """
    labels: pulumi.Output[dict]
    """
    Labels to apply to this forwarding rule.  A list of key->value pairs.
    """
    load_balancing_scheme: pulumi.Output[str]
    """
    This signifies what the ForwardingRule will be used for and can be
    EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
    Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
    and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
    INTERNAL is used for protocol forwarding to VMs from an internal IP address,
    and internal TCP/UDP load balancers.
    INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
    """
    name: pulumi.Output[str]
    """
    Name of the resource; provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035. Specifically, the name must be 1-63 characters long and match
    the regular expression `a-z?` which means the
    first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    network: pulumi.Output[str]
    """
    For internal load balancing, this field identifies the network that
    the load balanced IP should belong to for this Forwarding Rule. If
    this field is not specified, the default network will be used.
    This field is only used for INTERNAL load balancing.
    """
    network_tier: pulumi.Output[str]
    """
    The networking tier used for configuring this address. If this field is not
    specified, it is assumed to be PREMIUM.
    """
    port_range: pulumi.Output[str]
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
    ports: pulumi.Output[list]
    """
    This field is used along with the backend_service field for internal
    load balancing.
    When the load balancing scheme is INTERNAL, a single port or a comma
    separated list of ports can be configured. Only packets addressed to
    these ports will be forwarded to the backends configured with this
    forwarding rule.
    You may specify a maximum of up to 5 ports.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    A reference to the region where the regional forwarding rule resides.
    This field is not applicable to global forwarding rules.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    service_label: pulumi.Output[str]
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
    service_name: pulumi.Output[str]
    """
    The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.
    """
    subnetwork: pulumi.Output[str]
    """
    The subnetwork that the load balanced IP should belong to for this
    Forwarding Rule.  This field is only used for INTERNAL load balancing.
    If the network specified is in auto subnet mode, this field is
    optional. However, if the network is in custom subnet mode, a
    subnetwork must be specified.
    """
    target: pulumi.Output[str]
    """
    The URL of the target resource to receive the matched traffic.
    The target must live in the same region as the forwarding rule.
    The forwarded traffic must be of a type appropriate to the target
    object.
    """
    def __init__(__self__, resource_name, opts=None, all_ports=None, allow_global_access=None, backend_service=None, description=None, ip_address=None, ip_protocol=None, is_mirroring_collector=None, labels=None, load_balancing_scheme=None, name=None, network=None, network_tier=None, port_range=None, ports=None, project=None, region=None, service_label=None, subnetwork=None, target=None, __props__=None, __name__=None, __opts__=None):
        """
        A ForwardingRule resource. A ForwardingRule resource specifies which pool
        of target virtual machines to forward a packet to if it matches the given
        [IPAddress, IPProtocol, portRange] tuple.


        To get more information about ForwardingRule, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/forwardingRules)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/network/forwarding-rules)

        ## Example Usage - Forwarding Rule Global Internallb


        ```python
        import pulumi
        import pulumi_gcp as gcp

        hc = gcp.compute.HealthCheck("hc",
            check_interval_sec=1,
            timeout_sec=1,
            tcp_health_check={
                "port": "80",
            })
        backend = gcp.compute.RegionBackendService("backend",
            region="us-central1",
            health_checks=[hc.id])
        default_network = gcp.compute.Network("defaultNetwork", auto_create_subnetworks=False)
        default_subnetwork = gcp.compute.Subnetwork("defaultSubnetwork",
            ip_cidr_range="10.0.0.0/16",
            region="us-central1",
            network=default_network.id)
        # Forwarding rule for Internal Load Balancing
        default_forwarding_rule = gcp.compute.ForwardingRule("defaultForwardingRule",
            region="us-central1",
            load_balancing_scheme="INTERNAL",
            backend_service=backend.id,
            all_ports=True,
            allow_global_access=True,
            network=default_network.name,
            subnetwork=default_subnetwork.name)
        ```
        ## Example Usage - Forwarding Rule Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_target_pool = gcp.compute.TargetPool("defaultTargetPool")
        default_forwarding_rule = gcp.compute.ForwardingRule("defaultForwardingRule",
            target=default_target_pool.id,
            port_range="80")
        ```
        ## Example Usage - Forwarding Rule Internallb


        ```python
        import pulumi
        import pulumi_gcp as gcp

        hc = gcp.compute.HealthCheck("hc",
            check_interval_sec=1,
            timeout_sec=1,
            tcp_health_check={
                "port": "80",
            })
        backend = gcp.compute.RegionBackendService("backend",
            region="us-central1",
            health_checks=[hc.id])
        default_network = gcp.compute.Network("defaultNetwork", auto_create_subnetworks=False)
        default_subnetwork = gcp.compute.Subnetwork("defaultSubnetwork",
            ip_cidr_range="10.0.0.0/16",
            region="us-central1",
            network=default_network.id)
        # Forwarding rule for Internal Load Balancing
        default_forwarding_rule = gcp.compute.ForwardingRule("defaultForwardingRule",
            region="us-central1",
            load_balancing_scheme="INTERNAL",
            backend_service=backend.id,
            all_ports=True,
            network=default_network.name,
            subnetwork=default_subnetwork.name)
        ```
        ## Example Usage - Forwarding Rule Http Lb


        ```python
        import pulumi
        import pulumi_gcp as gcp

        debian_image = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        default_network = gcp.compute.Network("defaultNetwork",
            auto_create_subnetworks=False,
            routing_mode="REGIONAL")
        default_subnetwork = gcp.compute.Subnetwork("defaultSubnetwork",
            ip_cidr_range="10.1.2.0/24",
            region="us-central1",
            network=default_network.id)
        instance_template = gcp.compute.InstanceTemplate("instanceTemplate",
            machine_type="n1-standard-1",
            network_interface=[{
                "network": default_network.id,
                "subnetwork": default_subnetwork.id,
            }],
            disk=[{
                "sourceImage": debian_image.self_link,
                "autoDelete": True,
                "boot": True,
            }],
            tags=[
                "allow-ssh",
                "load-balanced-backend",
            ])
        rigm = gcp.compute.RegionInstanceGroupManager("rigm",
            region="us-central1",
            version=[{
                "instanceTemplate": instance_template.self_link,
                "name": "primary",
            }],
            base_instance_name="internal-glb",
            target_size=1)
        fw1 = gcp.compute.Firewall("fw1",
            network=default_network.id,
            source_ranges=["10.1.2.0/24"],
            allow=[
                {
                    "protocol": "tcp",
                },
                {
                    "protocol": "udp",
                },
                {
                    "protocol": "icmp",
                },
            ],
            direction="INGRESS")
        fw2 = gcp.compute.Firewall("fw2",
            network=default_network.id,
            source_ranges=["0.0.0.0/0"],
            allow=[{
                "protocol": "tcp",
                "ports": ["22"],
            }],
            target_tags=["allow-ssh"],
            direction="INGRESS")
        fw3 = gcp.compute.Firewall("fw3",
            network=default_network.id,
            source_ranges=[
                "130.211.0.0/22",
                "35.191.0.0/16",
            ],
            allow=[{
                "protocol": "tcp",
            }],
            target_tags=["load-balanced-backend"],
            direction="INGRESS")
        fw4 = gcp.compute.Firewall("fw4",
            network=default_network.id,
            source_ranges=["10.129.0.0/26"],
            target_tags=["load-balanced-backend"],
            allow=[
                {
                    "protocol": "tcp",
                    "ports": ["80"],
                },
                {
                    "protocol": "tcp",
                    "ports": ["443"],
                },
                {
                    "protocol": "tcp",
                    "ports": ["8000"],
                },
            ],
            direction="INGRESS")
        default_region_health_check = gcp.compute.RegionHealthCheck("defaultRegionHealthCheck",
            region="us-central1",
            http_health_check={
                "portSpecification": "USE_SERVING_PORT",
            })
        default_region_backend_service = gcp.compute.RegionBackendService("defaultRegionBackendService",
            load_balancing_scheme="INTERNAL_MANAGED",
            backend=[{
                "group": rigm.instance_group,
                "balancingMode": "UTILIZATION",
                "capacityScaler": 1,
            }],
            region="us-central1",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_region_health_check.id])
        default_region_url_map = gcp.compute.RegionUrlMap("defaultRegionUrlMap",
            region="us-central1",
            default_service=default_region_backend_service.id)
        default_region_target_http_proxy = gcp.compute.RegionTargetHttpProxy("defaultRegionTargetHttpProxy",
            region="us-central1",
            url_map=default_region_url_map.id)
        proxy = gcp.compute.Subnetwork("proxy",
            ip_cidr_range="10.129.0.0/26",
            region="us-central1",
            network=default_network.id,
            purpose="INTERNAL_HTTPS_LOAD_BALANCER",
            role="ACTIVE")
        # Forwarding rule for Internal Load Balancing
        default_forwarding_rule = gcp.compute.ForwardingRule("defaultForwardingRule",
            region="us-central1",
            ip_protocol="TCP",
            load_balancing_scheme="INTERNAL_MANAGED",
            port_range="80",
            target=default_region_target_http_proxy.id,
            network=default_network.id,
            subnetwork=default_subnetwork.id,
            network_tier="PREMIUM")
        ```

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
               An address must be specified by a literal IP address. > **NOTE**: While
               the API allows you to specify various resource paths for an address resource
               instead, this provider requires this to specifically be an IP address to
               avoid needing to fetching the IP address from resource paths on refresh
               or unnecessary diffs.
        :param pulumi.Input[str] ip_protocol: The IP protocol to which this rule applies.
               When the load balancing scheme is INTERNAL, only TCP and UDP are
               valid.
        :param pulumi.Input[bool] is_mirroring_collector: Indicates whether or not this load balancer can be used
               as a collector for packet mirroring. To prevent mirroring loops,
               instances behind this load balancer will not have their traffic
               mirrored even if a PacketMirroring rule applies to them. This
               can only be set to true for load balancers that have their
               loadBalancingScheme set to INTERNAL.
        :param pulumi.Input[dict] labels: Labels to apply to this forwarding rule.  A list of key->value pairs.
        :param pulumi.Input[str] load_balancing_scheme: This signifies what the ForwardingRule will be used for and can be
               EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
               Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
               and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
               INTERNAL is used for protocol forwarding to VMs from an internal IP address,
               and internal TCP/UDP load balancers.
               INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
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
        :param pulumi.Input[list] ports: This field is used along with the backend_service field for internal
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, all_ports=None, allow_global_access=None, backend_service=None, creation_timestamp=None, description=None, ip_address=None, ip_protocol=None, is_mirroring_collector=None, label_fingerprint=None, labels=None, load_balancing_scheme=None, name=None, network=None, network_tier=None, port_range=None, ports=None, project=None, region=None, self_link=None, service_label=None, service_name=None, subnetwork=None, target=None):
        """
        Get an existing ForwardingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
               An address must be specified by a literal IP address. > **NOTE**: While
               the API allows you to specify various resource paths for an address resource
               instead, this provider requires this to specifically be an IP address to
               avoid needing to fetching the IP address from resource paths on refresh
               or unnecessary diffs.
        :param pulumi.Input[str] ip_protocol: The IP protocol to which this rule applies.
               When the load balancing scheme is INTERNAL, only TCP and UDP are
               valid.
        :param pulumi.Input[bool] is_mirroring_collector: Indicates whether or not this load balancer can be used
               as a collector for packet mirroring. To prevent mirroring loops,
               instances behind this load balancer will not have their traffic
               mirrored even if a PacketMirroring rule applies to them. This
               can only be set to true for load balancers that have their
               loadBalancingScheme set to INTERNAL.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[dict] labels: Labels to apply to this forwarding rule.  A list of key->value pairs.
        :param pulumi.Input[str] load_balancing_scheme: This signifies what the ForwardingRule will be used for and can be
               EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
               Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
               and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
               INTERNAL is used for protocol forwarding to VMs from an internal IP address,
               and internal TCP/UDP load balancers.
               INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
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
        :param pulumi.Input[list] ports: This field is used along with the backend_service field for internal
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
