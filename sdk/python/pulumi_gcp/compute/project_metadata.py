# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProjectMetadata(pulumi.CustomResource):
    metadata: pulumi.Output[dict]
    """
    A series of key value pairs. Changing this resource
    updates the GCE state.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    def __init__(__self__, __name__, __opts__=None, metadata=None, project=None):
        """
        Manages metadata common to all instances for a project in GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/storing-retrieving-metadata)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/projects/setCommonInstanceMetadata).
        
        > **Note:**  If you want to manage only single key/value pairs within the project metadata
        rather than the entire set, then use
        google_compute_project_metadata_item.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] metadata: A series of key value pairs. Changing this resource
               updates the GCE state.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if metadata is None:
            raise TypeError('Missing required property metadata')
        __props__['metadata'] = metadata

        __props__['project'] = project

        super(ProjectMetadata, __self__).__init__(
            'gcp:compute/projectMetadata:ProjectMetadata',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

