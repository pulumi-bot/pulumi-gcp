# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetAddressResult(object):
    """
    A collection of values returned by getAddress.
    """
    def __init__(__self__, address=None, project=None, region=None, self_link=None, status=None, id=None):
        if address and not isinstance(address, str):
            raise TypeError('Expected argument address to be a str')
        __self__.address = address
        """
        The IP of the created resource.
        """
        if project and not isinstance(project, str):
            raise TypeError('Expected argument project to be a str')
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
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

def get_address(name=None, project=None, region=None):
    """
    Get the IP address from a static address. For more information see
    the official [API](https://cloud.google.com/compute/docs/reference/latest/addresses/get) documentation.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __ret__ = pulumi.runtime.invoke('gcp:compute/getAddress:getAddress', __args__)

    return GetAddressResult(
        address=__ret__.get('address'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        self_link=__ret__.get('selfLink'),
        status=__ret__.get('status'),
        id=__ret__.get('id'))
