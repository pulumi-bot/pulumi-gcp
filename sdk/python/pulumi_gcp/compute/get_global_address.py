# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetGlobalAddressResult:
    """
    A collection of values returned by getGlobalAddress.
    """
    def __init__(__self__, address=None, name=None, project=None, self_link=None, status=None, id=None):
        if address and not isinstance(address, str):
            raise TypeError('Expected argument address to be a str')
        __self__.address = address
        """
        The IP of the created resource.
        """
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        if project and not isinstance(project, str):
            raise TypeError('Expected argument project to be a str')
        __self__.project = project
        if self_link and not isinstance(self_link, str):
            raise TypeError('Expected argument self_link to be a str')
        __self__.self_link = self_link
        """
        The URI of the created resource.
        """
        if status and not isinstance(status, str):
            raise TypeError('Expected argument status to be a str')
        __self__.status = status
        """
        Indicates if the address is used. Possible values are: RESERVED or IN_USE.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_global_address(name=None,project=None,opts=None):
    """
    Get the IP address from a static address reserved for a Global Forwarding Rule which are only used for HTTP load balancing. For more information see
    the official [API](https://cloud.google.com/compute/docs/reference/latest/globalAddresses) documentation.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __ret__ = await pulumi.runtime.invoke('gcp:compute/getGlobalAddress:getGlobalAddress', __args__, opts=opts)

    return GetGlobalAddressResult(
        address=__ret__.get('address'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        self_link=__ret__.get('selfLink'),
        status=__ret__.get('status'),
        id=__ret__.get('id'))
