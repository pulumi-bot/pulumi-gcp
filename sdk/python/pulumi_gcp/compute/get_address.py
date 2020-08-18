# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetAddressResult',
    'AwaitableGetAddressResult',
    'get_address',
]



@pulumi.output_type
class GetAddressResult:
    """
    A collection of values returned by getAddress.
    """
    def __init__(__self__, address=None, id=None, name=None, project=None, region=None, self_link=None, status=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The IP of the created resource.
        """
        return pulumi.get(self, "address")

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
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Indicates if the address is used. Possible values are: RESERVED or IN_USE.
        """
        return pulumi.get(self, "status")



class AwaitableGetAddressResult(GetAddressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddressResult(
            address=self.address,
            id=self.id,
            name=self.name,
            project=self.project,
            region=self.region,
            self_link=self.self_link,
            status=self.status)


def get_address(name: Optional[str] = None,
                project: Optional[str] = None,
                region: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddressResult:
    """
    Get the IP address from a static address. For more information see
    the official [API](https://cloud.google.com/compute/docs/reference/latest/addresses/get) documentation.


    :param str name: A unique name for the resource, required by GCE.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The Region in which the created address reside.
           If it is not provided, the provider region is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getAddress:getAddress', __args__, opts=opts, typ=GetAddressResult).value

    return AwaitableGetAddressResult(
        address=__ret__.address,
        id=__ret__.id,
        name=__ret__.name,
        project=__ret__.project,
        region=__ret__.region,
        self_link=__ret__.self_link,
        status=__ret__.status)
