# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Route(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, description=None, dest_range=None, name=None, network=None, next_hop_gateway=None, next_hop_instance=None, next_hop_instance_zone=None, next_hop_ip=None, next_hop_vpn_tunnel=None, priority=None, project=None, tags=None):
        """Create a Route resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        if not dest_range:
            raise TypeError('Missing required property dest_range')
        __props__['destRange'] = dest_range

        __props__['name'] = name

        if not network:
            raise TypeError('Missing required property network')
        __props__['network'] = network

        __props__['nextHopGateway'] = next_hop_gateway

        __props__['nextHopInstance'] = next_hop_instance

        __props__['nextHopInstanceZone'] = next_hop_instance_zone

        __props__['nextHopIp'] = next_hop_ip

        __props__['nextHopVpnTunnel'] = next_hop_vpn_tunnel

        __props__['priority'] = priority

        __props__['project'] = project

        __props__['tags'] = tags

        __props__['next_hop_network'] = None
        __props__['self_link'] = None

        super(Route, __self__).__init__(
            'gcp:compute/route:Route',
            __name__,
            __props__,
            __opts__)

