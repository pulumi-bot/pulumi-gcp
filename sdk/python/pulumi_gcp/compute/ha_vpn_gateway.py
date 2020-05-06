# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class HaVpnGateway(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    An optional description of this resource.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035.  Specifically, the name must be 1-63 characters long and
    match the regular expression `a-z?` which means
    the first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    network: pulumi.Output[str]
    """
    The network this VPN gateway is accepting traffic for.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The region this gateway should sit in.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    vpn_interfaces: pulumi.Output[list]
    """
    A list of interfaces on this VPN gateway.

      * `id` (`float`) - an identifier for the resource with format `projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}`
      * `ip_address` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, network=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a VPN gateway running in GCP. This virtual device is managed
        by Google, but used only by you. This type of VPN Gateway allows for the creation
        of VPN solutions with higher availability than classic Target VPN Gateways.

        To get more information about HaVpnGateway, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/vpnGateways)
        * How-to Guides
            * [Choosing a VPN](https://cloud.google.com/vpn/docs/how-to/choosing-a-vpn)
            * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)

        ## Example Usage - Ha Vpn Gateway Basic


        {{ % example python % }}
        ```python
        import pulumi
        import pulumi_gcp as gcp

        network1 = gcp.compute.Network("network1", auto_create_subnetworks=False)
        ha_gateway1 = gcp.compute.HaVpnGateway("haGateway1",
            region="us-central1",
            network=network1.self_link)
        ```
        {{ % /example % }}
        ## Example Usage - Ha Vpn Gateway Gcp To Gcp


        {{ % example python % }}
        ```python
        import pulumi
        import pulumi_gcp as gcp

        network1 = gcp.compute.Network("network1",
            routing_mode="GLOBAL",
            auto_create_subnetworks=False)
        ha_gateway1 = gcp.compute.HaVpnGateway("haGateway1",
            region="us-central1",
            network=network1.self_link)
        network2 = gcp.compute.Network("network2",
            routing_mode="GLOBAL",
            auto_create_subnetworks=False)
        ha_gateway2 = gcp.compute.HaVpnGateway("haGateway2",
            region="us-central1",
            network=network2.self_link)
        network1_subnet1 = gcp.compute.Subnetwork("network1Subnet1",
            ip_cidr_range="10.0.1.0/24",
            region="us-central1",
            network=network1.self_link)
        network1_subnet2 = gcp.compute.Subnetwork("network1Subnet2",
            ip_cidr_range="10.0.2.0/24",
            region="us-west1",
            network=network1.self_link)
        network2_subnet1 = gcp.compute.Subnetwork("network2Subnet1",
            ip_cidr_range="192.168.1.0/24",
            region="us-central1",
            network=network2.self_link)
        network2_subnet2 = gcp.compute.Subnetwork("network2Subnet2",
            ip_cidr_range="192.168.2.0/24",
            region="us-east1",
            network=network2.self_link)
        router1 = gcp.compute.Router("router1",
            network=network1.name,
            bgp={
                "asn": 64514,
            })
        router2 = gcp.compute.Router("router2",
            network=network2.name,
            bgp={
                "asn": 64515,
            })
        tunnel1 = gcp.compute.VPNTunnel("tunnel1",
            region="us-central1",
            vpn_gateway=ha_gateway1.self_link,
            peer_gcp_gateway=ha_gateway2.self_link,
            shared_secret="a secret message",
            router=router1.self_link,
            vpn_gateway_interface=0)
        tunnel2 = gcp.compute.VPNTunnel("tunnel2",
            region="us-central1",
            vpn_gateway=ha_gateway1.self_link,
            peer_gcp_gateway=ha_gateway2.self_link,
            shared_secret="a secret message",
            router=router1.self_link,
            vpn_gateway_interface=1)
        tunnel3 = gcp.compute.VPNTunnel("tunnel3",
            region="us-central1",
            vpn_gateway=ha_gateway2.self_link,
            peer_gcp_gateway=ha_gateway1.self_link,
            shared_secret="a secret message",
            router=router2.self_link,
            vpn_gateway_interface=0)
        tunnel4 = gcp.compute.VPNTunnel("tunnel4",
            region="us-central1",
            vpn_gateway=ha_gateway2.self_link,
            peer_gcp_gateway=ha_gateway1.self_link,
            shared_secret="a secret message",
            router=router2.self_link,
            vpn_gateway_interface=1)
        router1_interface1 = gcp.compute.RouterInterface("router1Interface1",
            router=router1.name,
            region="us-central1",
            ip_range="169.254.0.1/30",
            vpn_tunnel=tunnel1.name)
        router1_peer1 = gcp.compute.RouterPeer("router1Peer1",
            router=router1.name,
            region="us-central1",
            peer_ip_address="169.254.0.2",
            peer_asn=64515,
            advertised_route_priority=100,
            interface=router1_interface1.name)
        router1_interface2 = gcp.compute.RouterInterface("router1Interface2",
            router=router1.name,
            region="us-central1",
            ip_range="169.254.1.1/30",
            vpn_tunnel=tunnel2.name)
        router1_peer2 = gcp.compute.RouterPeer("router1Peer2",
            router=router1.name,
            region="us-central1",
            peer_ip_address="169.254.1.2",
            peer_asn=64515,
            advertised_route_priority=100,
            interface=router1_interface2.name)
        router2_interface1 = gcp.compute.RouterInterface("router2Interface1",
            router=router2.name,
            region="us-central1",
            ip_range="169.254.0.1/30",
            vpn_tunnel=tunnel3.name)
        router2_peer1 = gcp.compute.RouterPeer("router2Peer1",
            router=router2.name,
            region="us-central1",
            peer_ip_address="169.254.0.2",
            peer_asn=64514,
            advertised_route_priority=100,
            interface=router2_interface1.name)
        router2_interface2 = gcp.compute.RouterInterface("router2Interface2",
            router=router2.name,
            region="us-central1",
            ip_range="169.254.1.1/30",
            vpn_tunnel=tunnel4.name)
        router2_peer2 = gcp.compute.RouterPeer("router2Peer2",
            router=router2.name,
            region="us-central1",
            peer_ip_address="169.254.1.2",
            peer_asn=64514,
            advertised_route_priority=100,
            interface=router2_interface2.name)
        ```
        {{ % /example % }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The network this VPN gateway is accepting traffic for.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region this gateway should sit in.
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
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['project'] = project
            __props__['region'] = region
            __props__['self_link'] = None
            __props__['vpn_interfaces'] = None
        super(HaVpnGateway, __self__).__init__(
            'gcp:compute/haVpnGateway:HaVpnGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, name=None, network=None, project=None, region=None, self_link=None, vpn_interfaces=None):
        """
        Get an existing HaVpnGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The network this VPN gateway is accepting traffic for.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region this gateway should sit in.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[list] vpn_interfaces: A list of interfaces on this VPN gateway.

        The **vpn_interfaces** object supports the following:

          * `id` (`pulumi.Input[float]`) - an identifier for the resource with format `projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}`
          * `ip_address` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["network"] = network
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["vpn_interfaces"] = vpn_interfaces
        return HaVpnGateway(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

