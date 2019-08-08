# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetTensorflowVersionsResult:
    """
    A collection of values returned by getTensorflowVersions.
    """
    def __init__(__self__, project=None, versions=None, zone=None, id=None):
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        __self__.versions = versions
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        __self__.zone = zone
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        delattr(self, "__await__")
        delattr(self, "__iter__")
        return self

    __iter__ = __await__

def get_tensorflow_versions(project=None,zone=None,opts=None):
    __args__ = dict()

    __args__['project'] = project
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:tpu/getTensorflowVersions:getTensorflowVersions', __args__, opts=opts).value

    return GetTensorflowVersionsResult(
        project=__ret__.get('project'),
        versions=__ret__.get('versions'),
        zone=__ret__.get('zone'),
        id=__ret__.get('id'))
