# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProjectMetadataItem(pulumi.CustomResource):
    key: pulumi.Output[str]
    """
    The metadata key to set.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    value: pulumi.Output[str]
    """
    The value to set for the given metadata key.
    """
    def __init__(__self__, resource_name, opts=None, key=None, project=None, value=None, __name__=None, __opts__=None):
        """
        Manages a single key/value pair on metadata common to all instances for
        a project in GCE. Using `google_compute_project_metadata_item` lets you
        manage a single key/value setting in Terraform rather than the entire
        project metadata map.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: The metadata key to set.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] value: The value to set for the given metadata key.
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

        if key is None:
            raise TypeError("Missing required property 'key'")
        __props__['key'] = key

        __props__['project'] = project

        if value is None:
            raise TypeError("Missing required property 'value'")
        __props__['value'] = value

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ProjectMetadataItem, __self__).__init__(
            'gcp:compute/projectMetadataItem:ProjectMetadataItem',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

