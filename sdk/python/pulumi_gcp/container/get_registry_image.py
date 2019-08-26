# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRegistryImageResult:
    """
    A collection of values returned by getRegistryImage.
    """
    def __init__(__self__, digest=None, image_url=None, name=None, project=None, region=None, tag=None, id=None):
        if digest and not isinstance(digest, str):
            raise TypeError("Expected argument 'digest' to be a str")
        __self__.digest = digest
        if image_url and not isinstance(image_url, str):
            raise TypeError("Expected argument 'image_url' to be a str")
        __self__.image_url = image_url
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        __self__.tag = tag
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetRegistryImageResult(GetRegistryImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryImageResult(
            digest=self.digest,
            image_url=self.image_url,
            name=self.name,
            project=self.project,
            region=self.region,
            tag=self.tag,
            id=self.id)

def get_registry_image(digest=None,name=None,project=None,region=None,tag=None,opts=None):
    """
    This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
    
    The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/container_registry_image.html.markdown.
    """
    __args__ = dict()

    __args__['digest'] = digest
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __args__['tag'] = tag
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:container/getRegistryImage:getRegistryImage', __args__, opts=opts).value

    return AwaitableGetRegistryImageResult(
        digest=__ret__.get('digest'),
        image_url=__ret__.get('imageUrl'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        tag=__ret__.get('tag'),
        id=__ret__.get('id'))
