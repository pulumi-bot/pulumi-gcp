# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RegionAutoscaler(pulumi.CustomResource):
    autoscaling_policy: pulumi.Output[dict]
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    region: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    target: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, autoscaling_policy=None, description=None, name=None, project=None, region=None, target=None, __name__=None, __opts__=None):
        """
        Represents an Autoscaler resource.
        
        Autoscalers allow you to automatically scale virtual machine instances in
        managed instance groups according to an autoscaling policy that you
        define.
        
        
        To get more information about RegionAutoscaler, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
        * How-to Guides
            * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_autoscaler.html.markdown.
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

        if autoscaling_policy is None:
            raise TypeError("Missing required property 'autoscaling_policy'")
        __props__['autoscaling_policy'] = autoscaling_policy

        __props__['description'] = description

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        if target is None:
            raise TypeError("Missing required property 'target'")
        __props__['target'] = target

        __props__['creation_timestamp'] = None
        __props__['self_link'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(RegionAutoscaler, __self__).__init__(
            'gcp:compute/regionAutoscaler:RegionAutoscaler',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

