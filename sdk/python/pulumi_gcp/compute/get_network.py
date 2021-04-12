# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetNetworkResult',
    'AwaitableGetNetworkResult',
    'get_network',
]

@pulumi.output_type
class GetNetworkResult:
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, description=None, gateway_ipv4=None, id=None, name=None, project=None, self_link=None, subnetworks_self_links=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if gateway_ipv4 and not isinstance(gateway_ipv4, str):
            raise TypeError("Expected argument 'gateway_ipv4' to be a str")
        pulumi.set(__self__, "gateway_ipv4", gateway_ipv4)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if subnetworks_self_links and not isinstance(subnetworks_self_links, list):
            raise TypeError("Expected argument 'subnetworks_self_links' to be a list")
        pulumi.set(__self__, "subnetworks_self_links", subnetworks_self_links)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of this network.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="gatewayIpv4")
    def gateway_ipv4(self) -> str:
        """
        The IP address of the gateway.
        """
        return pulumi.get(self, "gateway_ipv4")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="subnetworksSelfLinks")
    def subnetworks_self_links(self) -> Sequence[str]:
        """
        the list of subnetworks which belong to the network
        """
        return pulumi.get(self, "subnetworks_self_links")


class AwaitableGetNetworkResult(GetNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkResult(
            description=self.description,
            gateway_ipv4=self.gateway_ipv4,
            id=self.id,
            name=self.name,
            project=self.project,
            self_link=self.self_link,
            subnetworks_self_links=self.subnetworks_self_links)


def get_network(name: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkResult:
    """
    Get a network within GCE from its name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_network = gcp.compute.get_network(name="default-us-east1")
    ```


    :param str name: The name of the network.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getNetwork:getNetwork', __args__, opts=opts, typ=GetNetworkResult).value

    return AwaitableGetNetworkResult(
        description=__ret__.description,
        gateway_ipv4=__ret__.gateway_ipv4,
        id=__ret__.id,
        name=__ret__.name,
        project=__ret__.project,
        self_link=__ret__.self_link,
        subnetworks_self_links=__ret__.subnetworks_self_links)
