# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetActiveFolderResult:
    """
    A collection of values returned by getActiveFolder.
    """
    def __init__(__self__, display_name=None, id=None, name=None, parent=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The resource name of the Folder. This uniquely identifies the folder.
        """
        if parent and not isinstance(parent, str):
            raise TypeError("Expected argument 'parent' to be a str")
        __self__.parent = parent
class AwaitableGetActiveFolderResult(GetActiveFolderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActiveFolderResult(
            display_name=self.display_name,
            id=self.id,
            name=self.name,
            parent=self.parent)

def get_active_folder(display_name=None,parent=None,opts=None):
    """
    Get an active folder within GCP by `display_name` and `parent`.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/active_folder.html.markdown.


    :param str display_name: The folder's display name.
    :param str parent: The resource name of the parent Folder or Organization.
    """
    __args__ = dict()


    __args__['displayName'] = display_name
    __args__['parent'] = parent
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getActiveFolder:getActiveFolder', __args__, opts=opts).value

    return AwaitableGetActiveFolderResult(
        display_name=__ret__.get('displayName'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        parent=__ret__.get('parent'))
