# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetNetworkResult:
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, description=None, gateway_ipv4=None, name=None, project=None, self_link=None, subnetworks_self_links=None, id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of this network.
        """
        if gateway_ipv4 and not isinstance(gateway_ipv4, str):
            raise TypeError("Expected argument 'gateway_ipv4' to be a str")
        __self__.gateway_ipv4 = gateway_ipv4
        """
        The IP address of the gateway.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the resource.
        """
        if subnetworks_self_links and not isinstance(subnetworks_self_links, list):
            raise TypeError("Expected argument 'subnetworks_self_links' to be a list")
        __self__.subnetworks_self_links = subnetworks_self_links
        """
        the list of subnetworks which belong to the network
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_network(name=None,project=None,opts=None):
    """
    Get a network within GCE from its name.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
 .   if opts is None:
         opts = pulumi.ResourceOptions()
     if opts.version is None:
         opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('gcp:compute/getNetwork:getNetwork', __args__, opts=opts)

    return GetNetworkResult(
        description=__ret__.get('description'),
        gateway_ipv4=__ret__.get('gatewayIpv4'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        self_link=__ret__.get('selfLink'),
        subnetworks_self_links=__ret__.get('subnetworksSelfLinks'),
        id=__ret__.get('id'))
