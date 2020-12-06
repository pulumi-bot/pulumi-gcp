# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['FolderSink']


class FolderSink(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bigquery_options: Optional[pulumi.Input[pulumi.InputType['FolderSinkBigqueryOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderSinkExclusionArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 include_children: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a folder-level logging sink. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/) and
        [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).

        Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
        granted to the credentials used with this provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        log_bucket = gcp.storage.Bucket("log-bucket")
        my_folder = gcp.organizations.Folder("my-folder",
            display_name="My folder",
            parent="organizations/123456")
        my_sink = gcp.logging.FolderSink("my-sink",
            description="some explaination on what this is",
            folder=my_folder.name,
            destination=log_bucket.name.apply(lambda name: f"storage.googleapis.com/{name}"),
            filter="resource.type = gce_instance AND severity >= WARNING")
        log_writer = gcp.projects.IAMBinding("log-writer",
            role="roles/storage.objectCreator",
            members=[my_sink.writer_identity])
        ```

        ## Import

        Folder-level logging sinks can be imported using this format

        ```sh
         $ pulumi import gcp:logging/folderSink:FolderSink my_sink folders/{{folder_id}}/sinks/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['FolderSinkBigqueryOptionsArgs']] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this sink. The maximum length of the description is 8000 characters.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[bool] disabled: If set to True, then this sink is disabled and it does not export any log entries.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderSinkExclusionArgs']]]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
               one of exclusion_filters it will not be exported.
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
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['bigquery_options'] = bigquery_options
            __props__['description'] = description
            if destination is None:
                raise TypeError("Missing required property 'destination'")
            __props__['destination'] = destination
            __props__['disabled'] = disabled
            __props__['exclusions'] = exclusions
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bigquery_options: Optional[pulumi.Input[pulumi.InputType['FolderSinkBigqueryOptionsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderSinkExclusionArgs']]]]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            include_children: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            writer_identity: Optional[pulumi.Input[str]] = None) -> 'FolderSink':
        """
        Get an existing FolderSink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['FolderSinkBigqueryOptionsArgs']] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this sink. The maximum length of the description is 8000 characters.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[bool] disabled: If set to True, then this sink is disabled and it does not export any log entries.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderSinkExclusionArgs']]]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
               one of exclusion_filters it will not be exported.
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
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bigquery_options"] = bigquery_options
        __props__["description"] = description
        __props__["destination"] = destination
        __props__["disabled"] = disabled
        __props__["exclusions"] = exclusions
        __props__["filter"] = filter
        __props__["folder"] = folder
        __props__["include_children"] = include_children
        __props__["name"] = name
        __props__["writer_identity"] = writer_identity
        return FolderSink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bigqueryOptions")
    def bigquery_options(self) -> pulumi.Output['outputs.FolderSinkBigqueryOptions']:
        """
        Options that affect sinks exporting data to BigQuery. Structure documented below.
        """
        return pulumi.get(self, "bigquery_options")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of this sink. The maximum length of the description is 8000 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[str]:
        """
        The destination of the sink (or, in other words, where logs are written to). Can be a
        Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        ```python
        import pulumi
        ```
        The writer associated with the sink must have access to write to the above resource.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to True, then this sink is disabled and it does not export any log entries.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def exclusions(self) -> pulumi.Output[Optional[Sequence['outputs.FolderSinkExclusion']]]:
        """
        Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
        one of exclusion_filters it will not be exported.
        """
        return pulumi.get(self, "exclusions")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional[str]]:
        """
        The filter to apply when exporting logs. Only log entries that match the filter are exported.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[str]:
        """
        The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
        accepted.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="includeChildren")
    def include_children(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to include children folders in the sink export. If true, logs
        associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
        """
        return pulumi.get(self, "include_children")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the logging sink.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="writerIdentity")
    def writer_identity(self) -> pulumi.Output[str]:
        """
        The identity associated with this sink. This identity must be granted write access to the
        configured `destination`.
        """
        return pulumi.get(self, "writer_identity")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

