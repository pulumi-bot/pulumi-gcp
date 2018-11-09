# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class FolderSink(pulumi.CustomResource):
    """
    Manages a folder-level logging sink. For more information see
    [the official documentation](https://cloud.google.com/logging/docs/) and
    [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
    
    Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    granted to the credentials used with terraform.
    """
    def __init__(__self__, __name__, __opts__=None, destination=None, filter=None, folder=None, include_children=None, name=None):
        """Create a FolderSink resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not destination:
            raise TypeError('Missing required property destination')
        __props__['destination'] = destination

        __props__['filter'] = filter

        if not folder:
            raise TypeError('Missing required property folder')
        __props__['folder'] = folder

        __props__['include_children'] = include_children

        __props__['name'] = name

        __props__['writer_identity'] = None

        super(FolderSink, __self__).__init__(
            'gcp:logging/folderSink:FolderSink',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

