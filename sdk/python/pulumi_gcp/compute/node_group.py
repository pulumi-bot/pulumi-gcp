# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NodeGroup(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    node_template: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    size: pulumi.Output[float]
    zone: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, description=None, name=None, node_template=None, project=None, size=None, zone=None, __name__=None, __opts__=None):
        """
        Represents a NodeGroup resource to manage a group of sole-tenant nodes.
        
        
        To get more information about NodeGroup, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
        * How-to Guides
            * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
        
        > **Warning:** Due to limitations of the API, Terraform cannot update the
        number of nodes in a node group and changes to node group size either
        through Terraform config or through external changes will cause
        Terraform to delete and recreate the node group.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['name'] = name

        if node_template is None:
            raise TypeError("Missing required property 'node_template'")
        __props__['node_template'] = node_template

        __props__['project'] = project

        if size is None:
            raise TypeError("Missing required property 'size'")
        __props__['size'] = size

        __props__['zone'] = zone

        __props__['creation_timestamp'] = None
        __props__['self_link'] = None

        super(NodeGroup, __self__).__init__(
            'gcp:compute/nodeGroup:NodeGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

