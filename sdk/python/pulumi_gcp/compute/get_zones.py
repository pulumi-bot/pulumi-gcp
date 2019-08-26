# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetZonesResult:
    """
    A collection of values returned by getZones.
    """
    def __init__(__self__, names=None, project=None, region=None, status=None, id=None):
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of zones available in the given region
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetZonesResult(GetZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZonesResult(
            names=self.names,
            project=self.project,
            region=self.region,
            status=self.status,
            id=self.id)

def get_zones(project=None,region=None,status=None,opts=None):
    """
    Provides access to available Google Compute zones in a region for a given project.
    See more about [regions and zones](https://cloud.google.com/compute/docs/regions-zones/regions-zones) in the upstream docs.
    
    :param str project: Project from which to list available zones. Defaults to project declared in the provider.
    :param str region: Region from which to list available zones. Defaults to region declared in the provider.
    :param str status: Allows to filter list of zones based on their current status. Status can be either `UP` or `DOWN`.
           Defaults to no filtering (all available zones - both `UP` and `DOWN`).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_zones.html.markdown.
    """
    __args__ = dict()

    __args__['project'] = project
    __args__['region'] = region
    __args__['status'] = status
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getZones:getZones', __args__, opts=opts).value

    return AwaitableGetZonesResult(
        names=__ret__.get('names'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        status=__ret__.get('status'),
        id=__ret__.get('id'))
