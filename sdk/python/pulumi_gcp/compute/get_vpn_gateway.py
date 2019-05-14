# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetVPNGatewayResult:
    """
    A collection of values returned by getVPNGateway.
    """
    def __init__(__self__, description=None, name=None, network=None, project=None, region=None, self_link=None, id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of this VPN gateway.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        __self__.network = network
        """
        The network of this VPN gateway.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        Region of this VPN gateway.
        """
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the resource.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_vpn_gateway(name=None,project=None,region=None,opts=None):
    """
    Get a VPN gateway within GCE from its name.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('gcp:compute/getVPNGateway:getVPNGateway', __args__, opts=opts)

    return GetVPNGatewayResult(
        description=__ret__.get('description'),
        name=__ret__.get('name'),
        network=__ret__.get('network'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        self_link=__ret__.get('selfLink'),
        id=__ret__.get('id'))
