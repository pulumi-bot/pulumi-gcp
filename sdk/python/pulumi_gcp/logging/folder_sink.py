# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class FolderSink(pulumi.CustomResource):
    destination: pulumi.Output[str]
    """
    The destination of the sink (or, in other words, where logs are written to). Can be a
    Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
    The writer associated with the sink must have access to write to the above resource.
    """
    filter: pulumi.Output[str]
    """
    The filter to apply when exporting logs. Only log entries that match the filter are exported.
    See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
    write a filter.
    """
    folder: pulumi.Output[str]
    """
    The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
    accepted.
    """
    include_children: pulumi.Output[bool]
    """
    Whether or not to include children folders in the sink export. If true, logs
    associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
    """
    name: pulumi.Output[str]
    """
    The name of the logging sink.
    """
    writer_identity: pulumi.Output[str]
    """
    The identity associated with this sink. This identity must be granted write access to the
    configured `destination`.
    """
    def __init__(__self__, resource_name, opts=None, destination=None, filter=None, folder=None, include_children=None, name=None, __name__=None, __opts__=None):
        """
        Manages a folder-level logging sink. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/) and
        [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
        
        Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
        granted to the credentials used with terraform.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[str] filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[str] folder: The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
               accepted.
        :param pulumi.Input[bool] include_children: Whether or not to include children folders in the sink export. If true, logs
               associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
        :param pulumi.Input[str] name: The name of the logging sink.
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
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

