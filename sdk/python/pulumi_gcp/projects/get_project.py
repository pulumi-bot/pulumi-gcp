# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, filter=None, projects=None, id=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        __self__.filter = filter
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        __self__.projects = projects
        """
        A list of projects matching the provided filter. Structure is defined below.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            filter=self.filter,
            projects=self.projects,
            id=self.id)

def get_project(filter=None,opts=None):
    """
    Retrieve information about a set of projects based on a filter. See the
    [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
    for more details.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/projects.html.markdown.
    """
    __args__ = dict()

    __args__['filter'] = filter
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:projects/getProject:getProject', __args__, opts=opts).value

    return AwaitableGetProjectResult(
        filter=__ret__.get('filter'),
        projects=__ret__.get('projects'),
        id=__ret__.get('id'))
