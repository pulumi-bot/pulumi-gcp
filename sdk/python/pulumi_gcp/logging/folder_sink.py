# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
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
    def __init__(__self__, resource_name, opts=None, destination=None, filter=None, folder=None, include_children=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a FolderSink resource with the given unique name, props, and options.
        
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

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_folder_sink.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if destination is None:
                raise TypeError("Missing required property 'destination'")
            __props__['destination'] = destination
            __props__['filter'] = filter
            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            __props__['include_children'] = include_children
            __props__['name'] = name
            __props__['writer_identity'] = None
        super(FolderSink, __self__).__init__(
            'gcp:logging/folderSink:FolderSink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, destination=None, filter=None, folder=None, include_children=None, name=None, writer_identity=None):
        """
        Get an existing FolderSink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[str] writer_identity: The identity associated with this sink. This identity must be granted write access to the
               configured `destination`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_folder_sink.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["destination"] = destination
        __props__["filter"] = filter
        __props__["folder"] = folder
        __props__["include_children"] = include_children
        __props__["name"] = name
        __props__["writer_identity"] = writer_identity
        return FolderSink(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

