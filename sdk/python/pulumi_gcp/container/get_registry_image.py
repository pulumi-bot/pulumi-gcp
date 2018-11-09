# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetRegistryImageResult(object):
    """
    A collection of values returned by getRegistryImage.
    """
    def __init__(__self__, image_url=None, project=None, id=None):
        if image_url and not isinstance(image_url, str):
            raise TypeError('Expected argument image_url to be a str')
        __self__.image_url = image_url
        if project and not isinstance(project, str):
            raise TypeError('Expected argument project to be a str')
        __self__.project = project
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_registry_image(digest=None, name=None, project=None, region=None, tag=None):
    """
    This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
    
    The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
    """
    __args__ = dict()

    __args__['digest'] = digest
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __args__['tag'] = tag
    __ret__ = pulumi.runtime.invoke('gcp:container/getRegistryImage:getRegistryImage', __args__)

    return GetRegistryImageResult(
        image_url=__ret__.get('imageUrl'),
        project=__ret__.get('project'),
        id=__ret__.get('id'))
