# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Route(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    An optional description of this resource. Provide this property
    when you create the resource.
    """
    dest_range: pulumi.Output[str]
    """
    The destination range of outgoing packets that this route applies to.
    Only IPv4 is supported.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035.  Specifically, the name must be 1-63 characters long and
    match the regular expression `a-z?` which means
    the first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the
    last character, which cannot be a dash.
    """
    network: pulumi.Output[str]
    """
    The network that this route applies to.
    """
    next_hop_gateway: pulumi.Output[str]
    """
    URL to a gateway that should handle matching packets.
    Currently, you can only specify the internet gateway, using a full or
    partial valid URL:
    * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
    * `projects/project/global/gateways/default-internet-gateway`
    * `global/gateways/default-internet-gateway`
    * The string `default-internet-gateway`.
    """
    next_hop_ilb: pulumi.Output[str]
    """
    The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets.
    You can only specify the forwarding rule as a partial or full URL. For example, the following are all valid URLs:
    https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
    regions/region/forwardingRules/forwardingRule
    Note that this can only be used when the destinationRange is a public (non-RFC 1918) IP CIDR range.
    """
    next_hop_instance: pulumi.Output[str]
    """
    URL to an instance that should handle matching packets.
    You can specify this as a full or partial URL. For example:
    * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
    * `projects/project/zones/zone/instances/instance`
    * `zones/zone/instances/instance`
    * Just the instance name, with the zone in `next_hop_instance_zone`.
    """
    next_hop_instance_zone: pulumi.Output[str]
    """
    (Optional when `next_hop_instance` is
    specified)  The zone of the instance specified in
    `next_hop_instance`.  Omit if `next_hop_instance` is specified as
    a URL.
    """
    next_hop_ip: pulumi.Output[str]
    """
    Network IP address of an instance that should handle matching packets.
    """
    next_hop_network: pulumi.Output[str]
    """
    URL to a Network that should handle matching packets.
    """
    next_hop_vpn_tunnel: pulumi.Output[str]
    """
    URL to a VpnTunnel that should handle matching packets.
    """
    priority: pulumi.Output[float]
    """
    The priority of this route. Priority is used to break ties in cases
    where there is more than one matching route of equal prefix length.
    In the case of two routes with equal prefix length, the one with the
    lowest-numbered priority value wins.
    Default value is 1000. Valid range is 0 through 65535.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    tags: pulumi.Output[list]
    """
    A list of instance tags to which this route applies.
    """
    def __init__(__self__, resource_name, opts=None, description=None, dest_range=None, name=None, network=None, next_hop_gateway=None, next_hop_ilb=None, next_hop_instance=None, next_hop_instance_zone=None, next_hop_ip=None, next_hop_vpn_tunnel=None, priority=None, project=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a Route resource.

        A route is a rule that specifies how certain packets should be handled by
        the virtual network. Routes are associated with virtual machines by tag,
        and the set of routes for a particular virtual machine is called its
        routing table. For each packet leaving a virtual machine, the system
        searches that virtual machine's routing table for a single best matching
        route.

        Routes match packets by destination IP address, preferring smaller or more
        specific ranges over larger ones. If there is a tie, the system selects
        the route with the smallest priority value. If there is still a tie, it
        uses the layer three and four packet headers to select just one of the
        remaining matching routes. The packet is then forwarded as specified by
        the next_hop field of the winning route -- either to another virtual
        machine destination, a virtual machine gateway or a Compute
        Engine-operated gateway. Packets that do not match any route in the
        sending virtual machine's routing table will be dropped.

        A Route resource must have exactly one specification of either
        nextHopGateway, nextHopInstance, nextHopIp, nextHopVpnTunnel, or
        nextHopIlb.


        To get more information about Route, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routes)
        * How-to Guides
            * [Using Routes](https://cloud.google.com/vpc/docs/using-routes)

        ## Example Usage - Route Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.Network("defaultNetwork")
        default_route = gcp.compute.Route("defaultRoute",
            dest_range="15.0.0.0/24",
            network=default_network.name,
            next_hop_ip="10.132.1.5",
            priority=100)
        ```
        ## Example Usage - Route Ilb


        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.Network("defaultNetwork", auto_create_subnetworks=False)
        default_subnetwork = gcp.compute.Subnetwork("defaultSubnetwork",
            ip_cidr_range="10.0.1.0/24",
            region="us-central1",
            network=default_network.id)
        hc = gcp.compute.HealthCheck("hc",
            check_interval_sec=1,
            timeout_sec=1,
            tcp_health_check={
                "port": "80",
            })
        backend = gcp.compute.RegionBackendService("backend",
            region="us-central1",
            health_checks=[hc.id])
        default_forwarding_rule = gcp.compute.ForwardingRule("defaultForwardingRule",
            region="us-central1",
            load_balancing_scheme="INTERNAL",
            backend_service=backend.id,
            all_ports=True,
            network=default_network.name,
            subnetwork=default_subnetwork.name)
        route_ilb = gcp.compute.Route("route-ilb",
            dest_range="0.0.0.0/0",
            network=default_network.name,
            next_hop_ilb=default_forwarding_rule.id,
            priority=2000)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property
               when you create the resource.
        :param pulumi.Input[str] dest_range: The destination range of outgoing packets that this route applies to.
               Only IPv4 is supported.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
        :param pulumi.Input[str] network: The network that this route applies to.
        :param pulumi.Input[str] next_hop_gateway: URL to a gateway that should handle matching packets.
               Currently, you can only specify the internet gateway, using a full or
               partial valid URL:
               * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
               * `projects/project/global/gateways/default-internet-gateway`
               * `global/gateways/default-internet-gateway`
               * The string `default-internet-gateway`.
        :param pulumi.Input[str] next_hop_ilb: The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets.
               You can only specify the forwarding rule as a partial or full URL. For example, the following are all valid URLs:
               https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
               regions/region/forwardingRules/forwardingRule
               Note that this can only be used when the destinationRange is a public (non-RFC 1918) IP CIDR range.
        :param pulumi.Input[str] next_hop_instance: URL to an instance that should handle matching packets.
               You can specify this as a full or partial URL. For example:
               * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
               * `projects/project/zones/zone/instances/instance`
               * `zones/zone/instances/instance`
               * Just the instance name, with the zone in `next_hop_instance_zone`.
        :param pulumi.Input[str] next_hop_instance_zone: (Optional when `next_hop_instance` is
               specified)  The zone of the instance specified in
               `next_hop_instance`.  Omit if `next_hop_instance` is specified as
               a URL.
        :param pulumi.Input[str] next_hop_ip: Network IP address of an instance that should handle matching packets.
        :param pulumi.Input[str] next_hop_vpn_tunnel: URL to a VpnTunnel that should handle matching packets.
        :param pulumi.Input[float] priority: The priority of this route. Priority is used to break ties in cases
               where there is more than one matching route of equal prefix length.
               In the case of two routes with equal prefix length, the one with the
               lowest-numbered priority value wins.
               Default value is 1000. Valid range is 0 through 65535.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] tags: A list of instance tags to which this route applies.
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

            __props__['description'] = description
            if dest_range is None:
                raise TypeError("Missing required property 'dest_range'")
            __props__['dest_range'] = dest_range
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['next_hop_gateway'] = next_hop_gateway
            __props__['next_hop_ilb'] = next_hop_ilb
            __props__['next_hop_instance'] = next_hop_instance
            __props__['next_hop_instance_zone'] = next_hop_instance_zone
            __props__['next_hop_ip'] = next_hop_ip
            __props__['next_hop_vpn_tunnel'] = next_hop_vpn_tunnel
            __props__['priority'] = priority
            __props__['project'] = project
            __props__['tags'] = tags
            __props__['next_hop_network'] = None
            __props__['self_link'] = None
        super(Route, __self__).__init__(
            'gcp:compute/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, dest_range=None, name=None, network=None, next_hop_gateway=None, next_hop_ilb=None, next_hop_instance=None, next_hop_instance_zone=None, next_hop_ip=None, next_hop_network=None, next_hop_vpn_tunnel=None, priority=None, project=None, self_link=None, tags=None):
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property
               when you create the resource.
        :param pulumi.Input[str] dest_range: The destination range of outgoing packets that this route applies to.
               Only IPv4 is supported.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
        :param pulumi.Input[str] network: The network that this route applies to.
        :param pulumi.Input[str] next_hop_gateway: URL to a gateway that should handle matching packets.
               Currently, you can only specify the internet gateway, using a full or
               partial valid URL:
               * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
               * `projects/project/global/gateways/default-internet-gateway`
               * `global/gateways/default-internet-gateway`
               * The string `default-internet-gateway`.
        :param pulumi.Input[str] next_hop_ilb: The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets.
               You can only specify the forwarding rule as a partial or full URL. For example, the following are all valid URLs:
               https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
               regions/region/forwardingRules/forwardingRule
               Note that this can only be used when the destinationRange is a public (non-RFC 1918) IP CIDR range.
        :param pulumi.Input[str] next_hop_instance: URL to an instance that should handle matching packets.
               You can specify this as a full or partial URL. For example:
               * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
               * `projects/project/zones/zone/instances/instance`
               * `zones/zone/instances/instance`
               * Just the instance name, with the zone in `next_hop_instance_zone`.
        :param pulumi.Input[str] next_hop_instance_zone: (Optional when `next_hop_instance` is
               specified)  The zone of the instance specified in
               `next_hop_instance`.  Omit if `next_hop_instance` is specified as
               a URL.
        :param pulumi.Input[str] next_hop_ip: Network IP address of an instance that should handle matching packets.
        :param pulumi.Input[str] next_hop_network: URL to a Network that should handle matching packets.
        :param pulumi.Input[str] next_hop_vpn_tunnel: URL to a VpnTunnel that should handle matching packets.
        :param pulumi.Input[float] priority: The priority of this route. Priority is used to break ties in cases
               where there is more than one matching route of equal prefix length.
               In the case of two routes with equal prefix length, the one with the
               lowest-numbered priority value wins.
               Default value is 1000. Valid range is 0 through 65535.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[list] tags: A list of instance tags to which this route applies.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["dest_range"] = dest_range
        __props__["name"] = name
        __props__["network"] = network
        __props__["next_hop_gateway"] = next_hop_gateway
        __props__["next_hop_ilb"] = next_hop_ilb
        __props__["next_hop_instance"] = next_hop_instance
        __props__["next_hop_instance_zone"] = next_hop_instance_zone
        __props__["next_hop_ip"] = next_hop_ip
        __props__["next_hop_network"] = next_hop_network
        __props__["next_hop_vpn_tunnel"] = next_hop_vpn_tunnel
        __props__["priority"] = priority
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["tags"] = tags
        return Route(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

