# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetNodeTypesResult:
    """
    A collection of values returned by getNodeTypes.
    """
    def __init__(__self__, id=None, names=None, project=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of node types available in the given zone and project.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        __self__.zone = zone
class AwaitableGetNodeTypesResult(GetNodeTypesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodeTypesResult(
            id=self.id,
            names=self.names,
            project=self.project,
            zone=self.zone)

def get_node_types(project=None,zone=None,opts=None):
    """
    Provides available node types for Compute Engine sole-tenant nodes in a zone
    for a given project. For more information, see [the official documentation](https://cloud.google.com/compute/docs/nodes/#types) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTypes).

    ## Example Usage



    ```python
    import pulumi
    import pulumi_gcp as gcp

    central1b = gcp.compute.get_node_types(zone="us-central1-b")
    tmpl = gcp.compute.NodeTemplate("tmpl",
        region="us-central1",
        node_type=data["compute.getNodeTypes"]["types"]["names"])
    ```



    :param str project: ID of the project to list available node types for.
           Should match the project the nodes of this type will be deployed to.
           Defaults to the project that the provider is authenticated with.
    :param str zone: The zone to list node types for. Should be in zone of intended node groups and region of referencing node template. If `zone` is not specified, the provider-level zone must be set and is used
           instead.
    """
    __args__ = dict()


    __args__['project'] = project
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getNodeTypes:getNodeTypes', __args__, opts=opts).value

    return AwaitableGetNodeTypesResult(
        id=__ret__.get('id'),
        names=__ret__.get('names'),
        project=__ret__.get('project'),
        zone=__ret__.get('zone'))
