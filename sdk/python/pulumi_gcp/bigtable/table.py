# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Table(pulumi.CustomResource):
    instance_name: pulumi.Output[str]
    """
    The name of the Bigtable instance.
    """
    name: pulumi.Output[str]
    """
    The name of the table.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    split_keys: pulumi.Output[list]
    """
    A list of predefined keys to split the table on.
    """
    def __init__(__self__, __name__, __opts__=None, instance_name=None, name=None, project=None, split_keys=None):
        """
        Creates a Google Bigtable table inside an instance. For more information see
        [the official documentation](https://cloud.google.com/bigtable/) and
        [API](https://cloud.google.com/bigtable/docs/go/reference).
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] instance_name: The name of the Bigtable instance.
        :param pulumi.Input[str] name: The name of the table.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[list] split_keys: A list of predefined keys to split the table on.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if instance_name is None:
            raise TypeError('Missing required property instance_name')
        __props__['instance_name'] = instance_name

        __props__['name'] = name

        __props__['project'] = project

        __props__['split_keys'] = split_keys

        super(Table, __self__).__init__(
            'gcp:bigtable/table:Table',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

