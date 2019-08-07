# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NodeTemplate(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    node_affinity_labels: pulumi.Output[dict]
    node_type: pulumi.Output[str]
    node_type_flexibility: pulumi.Output[dict]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    server_binding: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, description=None, name=None, node_affinity_labels=None, node_type=None, node_type_flexibility=None, project=None, region=None, server_binding=None, __name__=None, __opts__=None):
        """
        Represents a NodeTemplate resource. Node templates specify properties
        for creating sole-tenant nodes, such as node type, vCPU and memory
        requirments, node affinity labels, and region.
        
        
        To get more information about NodeTemplate, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
        * How-to Guides
            * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_node_template.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description
        __props__['name'] = name
        __props__['node_affinity_labels'] = node_affinity_labels
        __props__['node_type'] = node_type
        __props__['node_type_flexibility'] = node_type_flexibility
        __props__['project'] = project
        __props__['region'] = region
        __props__['server_binding'] = server_binding
        __props__['creation_timestamp'] = None
        __props__['self_link'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(NodeTemplate, __self__).__init__(
            'gcp:compute/nodeTemplate:NodeTemplate',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

