# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class FolderExclusion(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A human-readable description.
    """
    disabled: pulumi.Output[bool]
    """
    Whether this exclusion rule should be disabled or not. This defaults to
    false.
    """
    filter: pulumi.Output[str]
    """
    The filter to apply when excluding logs. Only log entries that match the filter are excluded.
    See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
    write a filter.
    """
    folder: pulumi.Output[str]
    """
    The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
    accepted.
    """
    name: pulumi.Output[str]
    """
    The name of the logging exclusion.
    """
    def __init__(__self__, resource_name, opts=None, description=None, disabled=None, filter=None, folder=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a folder-level logging exclusion. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/) and
        [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
        
        Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
        granted to the credentials used with this provider.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] folder: The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
               accepted.
        :param pulumi.Input[str] name: The name of the logging exclusion.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_folder_exclusion.html.markdown.
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

            __props__['description'] = description
            __props__['disabled'] = disabled
            if filter is None:
                raise TypeError("Missing required property 'filter'")
            __props__['filter'] = filter
            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            __props__['name'] = name
        super(FolderExclusion, __self__).__init__(
            'gcp:logging/folderExclusion:FolderExclusion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, disabled=None, filter=None, folder=None, name=None):
        """
        Get an existing FolderExclusion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] folder: The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
               accepted.
        :param pulumi.Input[str] name: The name of the logging exclusion.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_folder_exclusion.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["disabled"] = disabled
        __props__["filter"] = filter
        __props__["folder"] = folder
        __props__["name"] = name
        return FolderExclusion(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

