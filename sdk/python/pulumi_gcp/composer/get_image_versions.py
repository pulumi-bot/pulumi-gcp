# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetImageVersionsResult:
    """
    A collection of values returned by getImageVersions.
    """
    def __init__(__self__, image_versions=None, project=None, region=None, id=None):
        if image_versions and not isinstance(image_versions, list):
            raise TypeError("Expected argument 'image_versions' to be a list")
        __self__.image_versions = image_versions
        """
        A list of composer image versions available in the given project and location. Each `image_version` contains:
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_image_versions(project=None,region=None,opts=None):
    """
    Provides access to available Cloud Composer versions in a region for a given project.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/composer_image_versions.html.markdown.
    """
    __args__ = dict()

    __args__['project'] = project
    __args__['region'] = region
    __ret__ = await pulumi.runtime.invoke('gcp:composer/getImageVersions:getImageVersions', __args__, opts=opts)

    return GetImageVersionsResult(
        image_versions=__ret__.get('imageVersions'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        id=__ret__.get('id'))
