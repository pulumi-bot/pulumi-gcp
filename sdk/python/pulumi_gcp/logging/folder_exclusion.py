# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class FolderExclusion(pulumi.CustomResource):
    """
    Manages a folder-level logging exclusion. For more information see
    [the official documentation](https://cloud.google.com/logging/docs/) and
    [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
    
    Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    granted to the credentials used with Terraform.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, disabled=None, filter=None, folder=None, name=None):
        """Create a FolderExclusion resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['disabled'] = disabled

        if not filter:
            raise TypeError('Missing required property filter')
        __props__['filter'] = filter

        if not folder:
            raise TypeError('Missing required property folder')
        __props__['folder'] = folder

        __props__['name'] = name

        super(FolderExclusion, __self__).__init__(
            'gcp:logging/folderExclusion:FolderExclusion',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

