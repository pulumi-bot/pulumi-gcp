# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    """
    Creates and manages a Google Spanner Instance. For more information, see the [official documentation](https://cloud.google.com/spanner/), or the [JSON API](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances).
    """
    def __init__(__self__, __name__, __opts__=None, config=None, display_name=None, labels=None, name=None, num_nodes=None, project=None):
        """Create a Instance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not config:
            raise TypeError('Missing required property config')
        __props__['config'] = config

        if not display_name:
            raise TypeError('Missing required property display_name')
        __props__['display_name'] = display_name

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['num_nodes'] = num_nodes

        __props__['project'] = project

        __props__['state'] = None

        super(Instance, __self__).__init__(
            'gcp:spanner/instance:Instance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

