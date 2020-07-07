# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetVPNGatewayResult:
    """
    A collection of values returned by getVPNGateway.
    """
    def __init__(__self__, description=None, id=None, name=None, network=None, project=None, region=None, self_link=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of this VPN gateway.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
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


class AwaitableGetVPNGatewayResult(GetVPNGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVPNGatewayResult(
            description=self.description,
            id=self.id,
            name=self.name,
            network=self.network,
            project=self.project,
            region=self.region,
            self_link=self.self_link)


def get_vpn_gateway(name=None, project=None, region=None, opts=None):
    """
    Get a VPN gateway within GCE from its name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_vpn_gateway = gcp.compute.get_vpn_gateway(name="vpn-gateway-us-east1")
    ```


    :param str name: The name of the VPN gateway.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The region in which the resource belongs. If it
           is not provided, the project region is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getVPNGateway:getVPNGateway', __args__, opts=opts).value

    return AwaitableGetVPNGatewayResult(
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        network=__ret__.get('network'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        self_link=__ret__.get('selfLink'))
