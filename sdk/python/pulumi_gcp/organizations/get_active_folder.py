# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetActiveFolderResult:
    """
    A collection of values returned by getActiveFolder.
    """
    def __init__(__self__, display_name=None, name=None, parent=None, id=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The resource name of the Folder. This uniquely identifies the folder.
        """
        if parent and not isinstance(parent, str):
            raise TypeError("Expected argument 'parent' to be a str")
        __self__.parent = parent
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_active_folder(display_name=None,parent=None,opts=None):
    """
    Get an active folder within GCP by `display_name` and `parent`.
    """
    __args__ = dict()

    __args__['displayName'] = display_name
    __args__['parent'] = parent
 .   if opts is None:
         opts = pulumi.ResourceOptions()
     if opts.version is None:
         opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('gcp:organizations/getActiveFolder:getActiveFolder', __args__, opts=opts)

    return GetActiveFolderResult(
        display_name=__ret__.get('displayName'),
        name=__ret__.get('name'),
        parent=__ret__.get('parent'),
        id=__ret__.get('id'))
